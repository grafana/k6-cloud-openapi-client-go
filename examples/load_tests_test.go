package k6

import (
	"bytes"
	"io"
	"log"
)

func ExampleLoadTestsAPI_LoadTestsList() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the base request:
	req := client.LoadTestsAPI.LoadTestsList(ctx).
		// We specify what stack id we want to make the request for:
		XStackId(stackID).
		// Then, we specify some optional parameters, like sorting by creation date:
		Orderby("created").
		// And requesting the total amount of projects to be present in the response:
		Count(true)
	// To handle pagination, we could skip some rows (e.g. 100) with:
	// .Skip(100)

	// Finally, we execute the request.
	loadTestsRes, httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("LoadTestsList request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)
	if loadTestsRes.HasCount() {
		log.Printf("Total amount of load tests: %d", *loadTestsRes.Count)
	}

	log.Println("The list of available load tests is:")
	for _, lt := range loadTestsRes.Value {
		log.Printf("%s (%d)\n", lt.GetName(), lt.GetId())
	}

	// Output:

}

func ExampleLoadTestsAPI_ProjectsLoadTestsCreate() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we prepare an [io.ReadCloser] with the load test script contents.
	// We define it inline for simplicity, but we could also use an [*os.File] here.
	f := io.NopCloser(bytes.NewReader([]byte(`
import http from 'k6/http';

export default function() {
	http.get('https://test.k6.io');
}
`)))

	// Then, we create the base request:
	req := client.LoadTestsAPI.ProjectsLoadTestsCreate(ctx, 3737039).
		// We set the load test name:
		Name("Example GCk6 load test").
		// We set the script:
		Script(f).
		// And we specify what stack id we want to make the request for:
		XStackId(stackID)

	// Finally, we execute the request.
	loadTestRes, httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("ProjectsLoadTestsCreate request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)
	log.Printf("The load test has been created with the id: %d", loadTestRes.GetId())

	// Output:

}

func ExampleLoadTestsAPI_ProjectsLoadTestsRetrieve() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we prepare an [io.ReadCloser] with the load test script contents.
	// We define it inline for simplicity, but we could also use an [*os.File] here.

	// Then, we create the base request:
	req := client.LoadTestsAPI.ProjectsLoadTestsRetrieve(ctx, 3760801).
		XStackId(stackID)

	// Finally, we execute the request.
	loadTestsRes, httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("ProjectsLoadTestsRetrieve request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)
	log.Printf("Total amount of load tests: %d", len(loadTestsRes.Value))

	// Output:

}

func ExampleLoadTestsAPI_LoadTestsScriptRetrieve() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the base request:
	req := client.LoadTestsAPI.LoadTestsScriptRetrieve(ctx, 960826).
		// And we specify what stack id we want to make the request for:
		XStackId(stackID)

	// Finally, we execute the request.
	ltsRes, httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("LoadTestsScriptRetrieve request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)
	log.Printf("The script is:\n%s", ltsRes)

	// Output:

}

func ExampleLoadTestsAPI_LoadTestsScriptUpdate() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we prepare an [io.ReadCloser] with the load test script contents.
	// We define it inline for simplicity, but we could also use an [*os.File] here.
	f := io.NopCloser(bytes.NewReader([]byte(`
import http from 'k6/http';

export default function() {
	http.get('https://test.k6.io/news.php');
}
`)))

	// Then, we initialize the base request:
	req := client.LoadTestsAPI.LoadTestsScriptUpdate(ctx, 960826).
		// We set the script:
		Body(f).
		// And we specify what stack id we want to make the request for:
		XStackId(stackID)

	// Finally, we execute the request.
	httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("LoadTestsScriptUpdate request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)

	// Output:

}

func ExampleLoadTestsAPI_LoadTestsStart() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the base request:
	req := client.LoadTestsAPI.LoadTestsStart(ctx, 960826).
		// And we specify what stack id we want to make the request for:
		XStackId(stackID)

	// Finally, we execute the request.
	testRunRes, httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("LoadTestsStart request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)
	log.Printf("The load test run has been created with the id: %d", testRunRes.GetId())

	// Output:

}

func ExampleLoadTestsAPI_LoadTestsDestroy() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the base request:
	req := client.LoadTestsAPI.LoadTestsDestroy(ctx, 960257).
		// And we specify what stack id we want to make the request for:
		XStackId(stackID)

	// Finally, we execute the request.
	httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("LoadTestsDestroy request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)

	// Output:

}
