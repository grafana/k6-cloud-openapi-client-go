package k6

import (
	"log"
)

func ExampleTestRunsAPI_TestRunsList() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the base request:
	req := client.TestRunsAPI.TestRunsList(ctx).
		// We specify what stack id we want to make the request for:
		XStackId(stackID).
		// And we specify some optional parameters, like requesting the total amount of
		// projects to be present in the response:
		Count(true)
	// To handle pagination, we could skip some rows (e.g. 100) with:
	// .Skip(100)

	// Finally, we execute the request.
	testRunsRes, httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("TestRunsList request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)
	if testRunsRes.HasCount() {
		log.Printf("Total amount of load tests: %d", *testRunsRes.Count)
	}

	log.Println("The list of available test runs is:")
	for _, tr := range testRunsRes.Value {
		log.Printf("%d - status: %s, result: %s\n", tr.GetId(), tr.GetStatus(), tr.GetResult())
	}

	// Output:

}

func ExampleTestRunsAPI_TestRunsRetrieve() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the base request:
	req := client.TestRunsAPI.TestRunsRetrieve(ctx, 3763757).
		// And we specify what stack id we want to make the request for:
		XStackId(stackID)

	// Finally, we execute the request.
	testRunRes, httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("TestRunsRetrieve request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)
	log.Printf("Test run id: %d (%d)", testRunRes.GetId(), testRunRes.GetProjectId())
	log.Printf("Test run status: %s", testRunRes.GetStatus())
	log.Printf("Test run result: %s", testRunRes.GetResult())
	log.Printf("Test run started at: %s", testRunRes.GetCreated().Format("Monday, January 2, 2006 at 3:04 PM"))
	if endedAt := testRunRes.GetEnded(); !endedAt.IsZero() {
		log.Printf("Test run ended at: %s", endedAt.Format("Monday, January 2, 2006 at 3:04 PM"))
	}

	// Output:

}

func ExampleTestRunsAPI_TestRunsAbort() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the base request:
	req := client.TestRunsAPI.TestRunsAbort(ctx, 2065455).
		// And we specify what stack id we want to make the request for:
		XStackId(stackID)

	// Finally, we execute the request.
	httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("TestRunsAbort request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)

	// Output:

}

func ExampleTestRunsAPI_TestRunsDestroy() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the base request:
	req := client.TestRunsAPI.TestRunsDestroy(ctx, 2065455).
		// And we specify what stack id we want to make the request for:
		XStackId(stackID)

	// Finally, we execute the request.
	httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("TestRunsDestroy request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)

	// Output:

}
