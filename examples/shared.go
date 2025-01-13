package k6

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/grafana/k6-cloud-openapi-client-go/k6"
)

const (
	k6CloudUrlEnvVar     = "K6_CLOUD_URL"
	k6CloudTokenEnvVar   = "K6_CLOUD_TOKEN"
	k6CloudStackIdEnvVar = "K6_CLOUD_STACK_ID"
)

var (
	ctx     context.Context
	stackID int32
	client  *k6.APIClient
)

// NOTE: This function is only used for the sake of simplicity on the example's end, so we don't need to
// repeat the same initialization code in every example function. In a real-world scenario, you would
// initialize the context, stackID, and client as part of your application's initialization, not as an `init()`.
func init() {
	var err error

	ctx, err = contextWithToken()
	if err != nil {
		log.Fatalln(err)
	}

	stackID, err = readStackIDFromEnv()
	if err != nil {
		log.Fatalln(err)
	}

	client = initClientFromEnv()
}

// contextWithToken initializes a context.Context with the API token, ready for use.
func contextWithToken() (context.Context, error) {
	// Initialize a context.Context with the API token, which is also a requirement for the API, from environment:
	envToken, tokenDefined := os.LookupEnv(k6CloudTokenEnvVar)
	if !tokenDefined {
		return nil, fmt.Errorf("the 'API token' must be specified through the %s environment variable", k6CloudTokenEnvVar)
	}

	return context.WithValue(context.Background(), k6.ContextAccessToken, envToken), nil
}

// initClientFromEnv initializes a k6.APIClient, ready for use.
func initClientFromEnv() *k6.APIClient {
	// Use the default API URL, or the one specified in the environment.
	// You can find the reference in the Grafana Cloud k6 REST API v6 docs:
	// https://grafana.com/docs/grafana-cloud/testing/k6/reference/cloud-rest-api/v6/#api-reference
	rootURL := "https://api.k6.io"
	if envURL, urlDefined := os.LookupEnv(k6CloudUrlEnvVar); urlDefined {
		rootURL = envURL
	}

	cfg := k6.NewConfiguration()
	cfg.Servers = []k6.ServerConfiguration{
		{URL: rootURL},
	}

	return k6.NewAPIClient(cfg)
}

// readStackIDFromEnv reads the Stack ID from the environment, and returns it for use.
func readStackIDFromEnv() (int32, error) {
	// Parse the "Stack ID" as an integer, from environment, as it is a requirement for the API:
	envStackID, stackDefined := os.LookupEnv(k6CloudStackIdEnvVar)
	if !stackDefined {
		return 0, fmt.Errorf("the 'Stack Id' must be specified through the %s environment variable", k6CloudStackIdEnvVar)
	}

	stackID, err := strconv.Atoi(envStackID)
	if err != nil {
		return 0, fmt.Errorf("the 'Stack Id' must be an integer: %s", err.Error())
	}

	return int32(stackID), nil
}
