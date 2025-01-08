package k6

import (
	"fmt"
	"io"
	"log"
	"os"
)

func ExampleLoadTestsAPI_LoadTestsList() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the base request.
	req := client.LoadTestsAPI.LoadTestsList(ctx)

	// We specify what Stack id we want to make the request for.
	req = req.XStackId(stackID)

	// Then, we specify some optional parameters, like sorting by creation date
	// and requesting the total amount of projects to be present in the response.
	req = req.Orderby("created")
	req = req.Count(true)

	// To handle pagination, we could skip some rows:
	// req = req.Skip(100)

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

	// First we create the load test script file and write some contents to it.
	f, err := os.CreateTemp("", "*.js")
	if err != nil {
		log.Fatalf("Could not create a test script file: %s", err.Error())
	}

	_, err = fmt.Fprintf(f, `
import http from 'k6/http';

export default function() {
	http.get('https://test.k6.io');
}
`)
	if err != nil {
		log.Fatalf("Could not write to the test script file: %s", err.Error())
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("Could not set the test script file ready for use: %s", err.Error())
	}

	// Then, we create the base request, and set the model.
	req := client.LoadTestsAPI.ProjectsLoadTestsCreate(ctx, 3736248)
	req = req.Name("Example GCk6 load test")
	req = req.Script(f)

	// We specify what Stack id we want to make the request for.
	req = req.XStackId(stackID)

	// Finally, we execute the request.
	loadTestRes, httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("ProjectsLoadTestsCreate request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)
	log.Printf("The load test has been created with the id: %d", loadTestRes.GetId())

	// Output:

}

func ExampleLoadTestsAPI_LoadTestsScriptRetrieve() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the base request.
	req := client.LoadTestsAPI.LoadTestsScriptRetrieve(ctx, 960271)

	// We specify what Stack id we want to make the request for.
	req = req.XStackId(stackID)

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

	// First we create the load test script file and write some contents to it.
	f, err := os.CreateTemp("", "*.js")
	if err != nil {
		log.Fatalf("Could not create a test script file: %s", err.Error())
	}

	_, err = fmt.Fprintf(f, `
import http from 'k6/http';

export default function() {
	http.get('https://test.k6.io/news.php');
}
`)
	if err != nil {
		log.Fatalf("Could not write to the test script file: %s", err.Error())
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		log.Fatalf("Could not set the test script file ready for use: %s", err.Error())
	}

	// First we initialize the base request.
	req := client.LoadTestsAPI.LoadTestsScriptUpdate(ctx, 960271)
	req = req.Body(f)

	// We specify what Stack id we want to make the request for.
	req = req.XStackId(stackID)

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

	// First we initialize the base request.
	req := client.LoadTestsAPI.LoadTestsStart(ctx, 960271)

	// We specify what Stack id we want to make the request for.
	req = req.XStackId(stackID)

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

	// First we initialize the base request.
	req := client.LoadTestsAPI.LoadTestsDestroy(ctx, 960257)

	// We specify what Stack id we want to make the request for.
	req = req.XStackId(stackID)

	// Finally, we execute the request.
	httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("LoadTestsDestroy request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)

	// Output:

}
