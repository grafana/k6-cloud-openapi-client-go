package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grafana/k6-cloud-openapi-client/k6"
)

const (
	k6CloudTokenEnvVar   = "K6_CLOUD_TOKEN"
	k6CloudStackIdEnvVar = "K6_CLOUD_STACK_ID"
)

func main() {
	// Initialize a context.Context, with the API token:
	ctx := context.Background()
	if token, ok := os.LookupEnv(k6CloudTokenEnvVar); ok {
		ctx = context.WithValue(ctx, k6.ContextAccessToken, token)
	}

	// Initialize the client configuration:
	cfg := k6.NewConfiguration()
	cfg.Servers = []k6.ServerConfiguration{
		// Let's assume you have the service forwarded to a local port:
		// kubectl --context=dev-us-east-0 --namespace k6-cloud port-forward service/general-api 8080:80
		{URL: "http://127.0.0.1:8080/public"},
	}

	// Initialize the client:
	client := k6.NewAPIClient(cfg)

	// Prepare the request, and execute
	req := client.LoadTestsAPI.ProjectsLoadTestsList(ctx)
	req = req.XStackId(os.Getenv(k6CloudStackIdEnvVar))
	loadTests, httpRes, err := req.Execute()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", httpRes)
	fmt.Printf("%+v\n", loadTests)
}
