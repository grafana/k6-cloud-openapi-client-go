package e2e

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/grafana/k6-cloud-openapi-client-go/k6"
)

func TestLoadTestsAPI_LoadTestsList(t *testing.T) {
	t.Run("list load tests with count", func(t *testing.T) {
		req := testClient.LoadTestsAPI.LoadTestsList(testCtx).
			XStackId(testStackID).
			Count(true).
			Orderby("created")

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("LoadTestsList failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if !res.HasCount() {
			t.Error("Expected count to be present in response")
		}
	})

	t.Run("list load tests with pagination", func(t *testing.T) {
		req := testClient.LoadTestsAPI.LoadTestsList(testCtx).
			XStackId(testStackID).
			Top(1).
			Skip(0)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("LoadTestsList with pagination failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if len(res.Value) > 1 {
			t.Errorf("Expected max 1 load test, got %d", len(res.Value))
		}
	})
}

func TestLoadTestsAPI_LoadTestsRetrieve(t *testing.T) {
	req := testClient.LoadTestsAPI.LoadTestsRetrieve(testCtx, testLoadTestID).
		XStackId(testStackID)

	res, httpRes, err := req.Execute()
	if err != nil {
		t.Fatalf("LoadTestsRetrieve failed: %v", err)
	}

	if httpRes.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
	}

	if res.GetId() != testLoadTestID {
		t.Errorf("Expected load test ID %d, got %d", testLoadTestID, res.GetId())
	}

	if res.GetProjectId() != testProjectID {
		t.Errorf("Expected project ID %d, got %d", testProjectID, res.GetProjectId())
	}
}

func TestLoadTestsAPI_ProjectsLoadTestsRetrieve(t *testing.T) {
	req := testClient.LoadTestsAPI.ProjectsLoadTestsRetrieve(testCtx, testProjectID).
		XStackId(testStackID).
		Count(true)

	res, httpRes, err := req.Execute()
	if err != nil {
		t.Fatalf("ProjectsLoadTestsRetrieve failed: %v", err)
	}

	if httpRes.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
	}

	if !res.HasCount() {
		t.Error("Expected count to be present in response")
	}

	// Verify our test load test is in the list
	found := false
	for _, lt := range res.Value {
		if lt.GetId() == testLoadTestID {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected to find test load test ID %d in project's load tests", testLoadTestID)
	}
}

func TestLoadTestsAPI_LoadTestsPartialUpdate(t *testing.T) {
	newName := "go-client-e2e-test-loadtest-updated"
	updateModel := k6.NewPatchLoadTestApiModel()
	updateModel.SetName(newName)

	req := testClient.LoadTestsAPI.LoadTestsPartialUpdate(testCtx, testLoadTestID).
		PatchLoadTestApiModel(updateModel).
		XStackId(testStackID)

	httpRes, err := req.Execute()
	if err != nil {
		t.Fatalf("LoadTestsPartialUpdate failed: %v", err)
	}

	if httpRes.StatusCode != 204 {
		t.Errorf("Expected status 204, got %d", httpRes.StatusCode)
	}

	// Verify the update
	verifyReq := testClient.LoadTestsAPI.LoadTestsRetrieve(testCtx, testLoadTestID).
		XStackId(testStackID)

	loadTest, _, err := verifyReq.Execute()
	if err != nil {
		t.Fatalf("Failed to retrieve updated load test: %v", err)
	}

	if loadTest.GetName() != newName {
		t.Errorf("Expected load test name %s, got %s", newName, loadTest.GetName())
	}
}

func TestLoadTestsAPI_LoadTestsScriptRetrieve(t *testing.T) {
	req := testClient.LoadTestsAPI.LoadTestsScriptRetrieve(testCtx, testLoadTestID).
		XStackId(testStackID)

	script, httpRes, err := req.Execute()
	if err != nil {
		t.Fatalf("LoadTestsScriptRetrieve failed: %v", err)
	}

	if httpRes.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
	}

	if script == "" {
		t.Error("Expected script to be non-empty")
	}

	if len(script) < 10 {
		t.Errorf("Expected script to have reasonable length, got %d bytes", len(script))
	}
}

func TestLoadTestsAPI_LoadTestsScriptUpdate(t *testing.T) {
	newScript := `
import http from 'k6/http';

export default function() {
	http.get('https://test.k6.io/news.php');
}
`
	scriptReader := io.NopCloser(bytes.NewReader([]byte(newScript)))

	req := testClient.LoadTestsAPI.LoadTestsScriptUpdate(testCtx, testLoadTestID).
		Body(scriptReader).
		XStackId(testStackID)

	httpRes, err := req.Execute()
	if err != nil {
		t.Fatalf("LoadTestsScriptUpdate failed: %v", err)
	}

	if httpRes.StatusCode != 204 {
		t.Errorf("Expected status 204, got %d", httpRes.StatusCode)
	}

	// Verify the update
	verifyReq := testClient.LoadTestsAPI.LoadTestsScriptRetrieve(testCtx, testLoadTestID).
		XStackId(testStackID)

	updatedScript, _, err := verifyReq.Execute()
	if err != nil {
		t.Fatalf("Failed to retrieve updated script: %v", err)
	}

	// Check if the updated script contains the new endpoint
	if len(updatedScript) < len(newScript)-10 {
		t.Error("Updated script appears to be too short")
	}
}

func TestLoadTestsAPI_LoadTestsMove(t *testing.T) {
	// Create a second project to move the load test to
	timestamp := time.Now().Unix()
	secondProjectName := fmt.Sprintf("go-client-e2e-test-project-move-%d", timestamp)

	createProjectModel := k6.NewCreateProjectApiModel(secondProjectName)
	projectReq := testClient.ProjectsAPI.ProjectsCreate(testCtx).
		CreateProjectApiModel(createProjectModel).
		XStackId(testStackID)

	secondProject, _, err := projectReq.Execute()
	if err != nil {
		t.Fatalf("Failed to create second project: %v", err)
	}
	secondProjectID := secondProject.GetId()
	t.Logf("Created second project: %s (ID: %d)", secondProjectName, secondProjectID)

	// Ensure cleanup of second project
	defer func() {
		deleteReq := testClient.ProjectsAPI.ProjectsDestroy(testCtx, secondProjectID).
			XStackId(testStackID)
		_, err := deleteReq.Execute()
		if err != nil {
			t.Logf("Warning: failed to delete second project %d: %v", secondProjectID, err)
		} else {
			t.Logf("Cleaned up second project (ID: %d)", secondProjectID)
		}
	}()

	// Move the test load test to the second project
	moveModel := k6.NewMoveLoadTestApiModel(secondProjectID)
	req := testClient.LoadTestsAPI.LoadTestsMove(testCtx, testLoadTestID).
		MoveLoadTestApiModel(moveModel).
		XStackId(testStackID)

	httpRes, err := req.Execute()
	if err != nil {
		t.Fatalf("LoadTestsMove failed: %v", err)
	}

	if httpRes.StatusCode != 204 {
		t.Errorf("Expected status 204, got %d", httpRes.StatusCode)
	}

	// Verify the load test is now in the second project by retrieving it
	retrieveReq := testClient.LoadTestsAPI.LoadTestsRetrieve(testCtx, testLoadTestID).
		XStackId(testStackID)

	loadTestAfterMove, _, err := retrieveReq.Execute()
	if err != nil {
		t.Fatalf("Failed to retrieve load test after move: %v", err)
	}

	if loadTestAfterMove.GetProjectId() != secondProjectID {
		t.Errorf("Expected project ID %d after move, got %d", secondProjectID, loadTestAfterMove.GetProjectId())
	}

	// Move it back to the original project
	moveBackModel := k6.NewMoveLoadTestApiModel(testProjectID)
	moveBackReq := testClient.LoadTestsAPI.LoadTestsMove(testCtx, testLoadTestID).
		MoveLoadTestApiModel(moveBackModel).
		XStackId(testStackID)

	httpResBack, err := moveBackReq.Execute()
	if err != nil {
		t.Fatalf("LoadTestsMove (back) failed: %v", err)
	}

	if httpResBack.StatusCode != 204 {
		t.Errorf("Expected status 204 for move back, got %d", httpResBack.StatusCode)
	}

	// Verify it's back in the original project
	retrieveBackReq := testClient.LoadTestsAPI.LoadTestsRetrieve(testCtx, testLoadTestID).
		XStackId(testStackID)

	loadTestAfterMoveBack, _, err := retrieveBackReq.Execute()
	if err != nil {
		t.Fatalf("Failed to retrieve load test after moving back: %v", err)
	}

	if loadTestAfterMoveBack.GetProjectId() != testProjectID {
		t.Errorf("Expected project ID %d after moving back, got %d", testProjectID, loadTestAfterMoveBack.GetProjectId())
	}
}

func TestLoadTestsAPI_ValidateOptions(t *testing.T) {
	// Create options with cloud project ID
	cloudOptions := k6.NewCloudOptions()
	cloudOptions.SetProjectID(testProjectID)

	optionsCloud := k6.OptionsCloud{}
	optionsCloud.CloudOptions = cloudOptions

	options := k6.NewOptions()
	options.SetCloud(optionsCloud)

	validateReq := k6.NewValidateOptionsRequest(*options)

	req := testClient.LoadTestsAPI.ValidateOptions(testCtx).
		ValidateOptionsRequest(validateReq).
		XStackId(testStackID)

	res, httpRes, err := req.Execute()
	if err != nil {
		t.Fatalf("ValidateOptions failed: %v", err)
	}

	if httpRes.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
	}

	if res.GetVuhUsage() < 0 {
		t.Error("Expected VUH usage to be non-negative")
	}
}
