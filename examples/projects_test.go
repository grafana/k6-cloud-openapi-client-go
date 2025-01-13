package k6

import (
	"log"

	"github.com/grafana/k6-cloud-openapi-client-go/k6"
)

func ExampleProjectsAPI_ProjectsList() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the base request:
	req := client.ProjectsAPI.ProjectsList(ctx).
		// We specify what stack id we want to make the request for:
		XStackId(stackID).
		// Then, we specify some optional parameters, like sorting by creation date:
		Orderby("created").
		// And requesting the total amount of projects to be present in the response:
		Count(true)
	// To handle pagination, we could skip some rows (e.g. 100) with:
	// .Skip(100)

	// Finally, we execute the request.
	projectsRes, httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("ProjectsList request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)
	if projectsRes.HasCount() {
		log.Printf("Total amount of projects: %d", *projectsRes.Count)
	}

	log.Println("The list of available projects is:")
	for _, p := range projectsRes.Value {
		log.Printf("%s\n", p.GetName())
	}

	// Output:

}

func ExampleProjectsAPI_ProjectsCreate() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the project model.
	toCreate := k6.NewCreateProjectApiModel("Example GCk6 project")

	// Then, we create the base request:
	req := client.ProjectsAPI.ProjectsCreate(ctx).
		// We set the model:
		CreateProjectApiModel(toCreate).
		// And we specify what stack id we want to make the request for:
		XStackId(stackID)

	// Finally, we execute the request.
	createdRes, httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("ProjectsCreate request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)
	log.Printf("The project '%s' has been created with the id: %d", createdRes.GetName(), createdRes.GetId())

	// Output:

}

func ExampleProjectsAPI_ProjectsPartialUpdate() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the model.
	toUpdate := k6.NewPatchProjectApiModel("Example GCk6 project (Public API)")

	// Then, we create the base request:
	req := client.ProjectsAPI.ProjectsPartialUpdate(ctx, 3737039).
		// We set the model:
		PatchProjectApiModel(toUpdate).
		// And we specify what stack id we want to make the request for:
		XStackId(stackID)

	// Finally, we execute the request.
	httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("ProjectsPartialUpdate request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)

	// Output:

}

func ExampleProjectsAPI_ProjectsLimitsRetrieve() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the base request:
	req := client.ProjectsAPI.ProjectsLimitsRetrieve(ctx, 3737039).
		// And we specify what stack id we want to make the request for:
		XStackId(stackID)

	// Finally, we execute the request.
	limitsRes, httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("ProjectsLimitsRetrieve request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)
	log.Printf("Project limits for '%d'", limitsRes.GetProjectId())
	log.Printf("Max VUh per month: %d", limitsRes.GetVuhMaxPerMonth())
	log.Printf("Max VUh per test: %d", limitsRes.GetVuMaxPerTest())
	log.Printf("Max VUh browser per test: %d", limitsRes.GetVuBrowserMaxPerTest())
	log.Printf("Max duration per test: %d", limitsRes.GetDurationMaxPerTest())

	// Output:

}

func ExampleProjectsAPI_ProjectsLimitsPartialUpdate() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we prepare the update request model.
	toUpdate := k6.NewPatchProjectLimitsRequest()
	toUpdate.SetDurationMaxPerTest(7200)

	// Then, we create the base request:
	req := client.ProjectsAPI.ProjectsLimitsPartialUpdate(ctx, 3737039).
		// We set the model:
		PatchProjectLimitsRequest(toUpdate).
		// And we specify what stack id we want to make the request for:
		XStackId(stackID)

	// Finally, we execute the request.
	httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("ProjectsLimitsPartialUpdate request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)

	// Output:

}

func ExampleProjectsAPI_ProjectsDestroy() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the base request:
	req := client.ProjectsAPI.ProjectsDestroy(ctx, 3737039).
		// And we specify what stack id we want to make the request for:
		XStackId(stackID)

	// Finally, we execute the request.
	httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("ProjectsDestroy request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)

	// Output:

}
