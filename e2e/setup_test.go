package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/grafana/k6-cloud-openapi-client-go/k6"
)

const (
	k6CloudURLEnvVar        = "K6_CLOUD_URL"
	k6CloudTokenEnvVar      = "K6_CLOUD_TOKEN"
	k6CloudStackIDEnvVar    = "K6_CLOUD_STACK_ID"
	k6StackURLEnvVar        = "K6_STACK_URL"
	k6CloudStackTokenEnvVar = "K6_CLOUD_STACK_TOKEN"
)

var (
	testCtx          context.Context
	testClient       *k6.APIClient
	testStackID      int32
	testStackToken   string
	testK6URL        string
	testStackURL     string
	testProjectID    int32
	testLoadTestID   int32
	testLoadZoneID   int32
	testLoadZoneK6ID string
)

// TestMain handles setup and teardown for all e2e tests
func TestMain(m *testing.M) {
	// Setup phase
	if err := setupTestEnvironment(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to setup test environment: %v\n", err)
		os.Exit(1)
	}

	// Run tests
	code := m.Run()

	// Teardown phase
	cleanupTestEnvironment()

	os.Exit(code)
}

// setupTestEnvironment initializes the test client and creates shared test resources
func setupTestEnvironment() error {
	var err error

	// Validate required environment variables
	if err := validateEnvironment(); err != nil {
		return err
	}

	// Initialize context with token
	testCtx, err = contextWithToken()
	if err != nil {
		return fmt.Errorf("failed to initialize context: %w", err)
	}

	// Initialize client
	testClient = initClientFromEnv()

	// Get k6 Cloud API URL from environment (required, validated in validateEnvironment)
	testK6URL = os.Getenv(k6CloudURLEnvVar)

	// Get stack URL from environment (required, validated in validateEnvironment)
	testStackURL = os.Getenv(k6StackURLEnvVar)

	// Get stack ID from environment
	testStackID, err = readStackIDFromEnv()
	if err != nil {
		return fmt.Errorf("failed to read stack ID: %w", err)
	}

	// Get stack token from environment (required, validated in validateEnvironment)
	testStackToken = os.Getenv(k6CloudStackTokenEnvVar)

	timestamp := time.Now().Unix()

	// Create private load zone first (needed for load zone tests)
	if err := createPrivateLoadZone(timestamp); err != nil {
		return fmt.Errorf("failed to create private load zone: %w", err)
	}

	// Create test project
	projectName := fmt.Sprintf("go-client-e2e-test-project-%d", timestamp)

	createProjectModel := k6.NewCreateProjectApiModel(projectName)
	projectReq := testClient.ProjectsAPI.ProjectsCreate(testCtx).
		CreateProjectApiModel(createProjectModel).
		XStackId(testStackID)

	project, _, err := projectReq.Execute()
	if err != nil {
		return fmt.Errorf("failed to create test project: %w", err)
	}
	testProjectID = project.GetId()
	fmt.Printf("Created test project: %s (ID: %d)\n", projectName, testProjectID)

	// Create test load test
	loadTestName := fmt.Sprintf("go-client-e2e-test-loadtest-%d", timestamp)
	script := io.NopCloser(bytes.NewReader([]byte(`
import http from 'k6/http';

export default function() {
	http.get('https://test.k6.io');
}
`)))

	loadTestReq := testClient.LoadTestsAPI.ProjectsLoadTestsCreate(testCtx, testProjectID).
		Name(loadTestName).
		Script(script).
		XStackId(testStackID)

	loadTest, _, err := loadTestReq.Execute()
	if err != nil {
		// If load test creation fails, clean up project
		cleanupProject()
		return fmt.Errorf("failed to create test load test: %w", err)
	}
	testLoadTestID = loadTest.GetId()
	fmt.Printf("Created test load test: %s (ID: %d)\n", loadTestName, testLoadTestID)

	return nil
}

// cleanupTestEnvironment removes all test resources
func cleanupTestEnvironment() {
	fmt.Println("Cleaning up test environment...")

	// Delete test load test
	if testLoadTestID != 0 {
		cleanupLoadTest()
	}

	// Delete test project
	if testProjectID != 0 {
		cleanupProject()
	}

	// Delete private load zone
	if testLoadZoneK6ID != "" {
		cleanupPrivateLoadZone()
	}
}

// cleanupLoadTest deletes the test load test
func cleanupLoadTest() {
	req := testClient.LoadTestsAPI.LoadTestsDestroy(testCtx, testLoadTestID).
		XStackId(testStackID)

	_, err := req.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to delete test load test %d: %v\n", testLoadTestID, err)
	} else {
		fmt.Printf("Deleted test load test (ID: %d)\n", testLoadTestID)
	}
}

// cleanupProject deletes the test project
func cleanupProject() {
	req := testClient.ProjectsAPI.ProjectsDestroy(testCtx, testProjectID).
		XStackId(testStackID)

	_, err := req.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to delete test project %d: %v\n", testProjectID, err)
	} else {
		fmt.Printf("Deleted test project (ID: %d)\n", testProjectID)
	}
}

// contextWithToken initializes a context.Context with the API token
func contextWithToken() (context.Context, error) {
	envToken, tokenDefined := os.LookupEnv(k6CloudTokenEnvVar)
	if !tokenDefined {
		return nil, fmt.Errorf("the 'API token' must be specified through the %s environment variable", k6CloudTokenEnvVar)
	}

	return context.WithValue(context.Background(), k6.ContextAccessToken, envToken), nil
}

// initClientFromEnv initializes a k6.APIClient from environment variables
func initClientFromEnv() *k6.APIClient {
	// URL is required and validated in validateEnvironment
	rootURL := os.Getenv(k6CloudURLEnvVar)

	cfg := k6.NewConfiguration()
	cfg.Servers = []k6.ServerConfiguration{
		{URL: rootURL},
	}

	return k6.NewAPIClient(cfg)
}

// readStackIDFromEnv reads the Stack ID from environment
func readStackIDFromEnv() (int32, error) {
	envStackID, stackDefined := os.LookupEnv(k6CloudStackIDEnvVar)
	if !stackDefined {
		return 0, fmt.Errorf("the 'Stack Id' must be specified through the %s environment variable", k6CloudStackIDEnvVar)
	}

	stackID, err := strconv.Atoi(envStackID)
	if err != nil {
		return 0, fmt.Errorf("the 'Stack Id' must be an integer: %s", err.Error())
	}

	return int32(stackID), nil
}

// validateEnvironment checks that all required environment variables are set
func validateEnvironment() error {
	var missing []string

	if _, exists := os.LookupEnv(k6CloudTokenEnvVar); !exists {
		missing = append(missing, k6CloudTokenEnvVar)
	}

	if _, exists := os.LookupEnv(k6CloudStackIDEnvVar); !exists {
		missing = append(missing, k6CloudStackIDEnvVar)
	}

	if _, exists := os.LookupEnv(k6CloudURLEnvVar); !exists {
		missing = append(missing, k6CloudURLEnvVar)
	}

	if _, exists := os.LookupEnv(k6StackURLEnvVar); !exists {
		missing = append(missing, k6StackURLEnvVar)
	}

	if _, exists := os.LookupEnv(k6CloudStackTokenEnvVar); !exists {
		missing = append(missing, k6CloudStackTokenEnvVar)
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required environment variables: %v\n\nPlease set:\n  export %s=<your-token>\n  export %s=<your-stack-id>\n  export %s=<api-url>        # e.g., https://api.k6.io \n  export %s=<stack-url>      # e.g., https://api.staging.k6.io\n  export %s=<stack-token>    # e.g., your-stack-token",
			missing, k6CloudTokenEnvVar, k6CloudStackIDEnvVar, k6CloudURLEnvVar, k6StackURLEnvVar, k6CloudStackTokenEnvVar)
	}

	return nil
}

// createPrivateLoadZone creates a private load zone for testing
func createPrivateLoadZone(timestamp int64) error {
	testLoadZoneK6ID = fmt.Sprintf("go-client-e2e-load-zone-%d", timestamp)
	providerID := fmt.Sprintf("go-client-e2e-provider-%d", timestamp)

	// Create load zone payload similar to TypeScript test
	payload := fmt.Sprintf(`{
		"k6_load_zone_id": "%s",
		"provider_id": "%s",
		"pod_tiers": {
			"cpu": "200m",
			"memory": "200Mi"
		},
		"config": {
			"key": "value",
			"load_runner_image": "test-image"
		}
	}`, testLoadZoneK6ID, providerID)

	// Make HTTP request to cloud-resources API
	req, err := http.NewRequestWithContext(
		testCtx,
		http.MethodPost,
		fmt.Sprintf("%s/cloud-resources/v1/load-zones", testK6URL),
		bytes.NewBufferString(payload),
	)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Stack-Id", fmt.Sprintf("%d", testStackID))
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", testStackToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to create private load zone: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create private load zone, status: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse response to get the load zone ID
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	// Get the ID from response if available
	if id, ok := result["id"].(float64); ok {
		testLoadZoneID = int32(id)
	}

	fmt.Printf("Created test private load zone: %s (ID: %d)\n", testLoadZoneK6ID, testLoadZoneID)
	return nil
}

// cleanupPrivateLoadZone deletes the private load zone
func cleanupPrivateLoadZone() {
	req, err := http.NewRequestWithContext(
		testCtx,
		http.MethodDelete,
		fmt.Sprintf("%s/cloud-resources/v1/load-zones/%s", testK6URL, testLoadZoneK6ID),
		nil,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to create delete request for load zone: %v\n", err)
		return
	}

	req.Header.Set("X-Stack-Id", fmt.Sprintf("%d", testStackID))
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", testStackToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: failed to delete private load zone %s: %v\n", testLoadZoneK6ID, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Fprintf(os.Stderr, "Warning: failed to delete private load zone %s, status: %d, body: %s\n", testLoadZoneK6ID, resp.StatusCode, string(body))
	} else {
		fmt.Printf("Deleted test private load zone (k6_load_zone_id: %s)\n", testLoadZoneK6ID)
	}
}
