package k6

import (
	"log"

	"github.com/grafana/k6-cloud-openapi-client/k6"
)

func ExampleProjectsAPI_ProjectsList() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the base request.
	req := client.ProjectsAPI.ProjectsList(ctx)

	// We specify what Stack id we want to make the request for.
	req = req.XStackId(stackID)

	// Then, we specify some optional parameters, like sorting by creation date
	// and requesting the total amount of projects to be present in the response.
	req = req.Orderby("created")
	req = req.Count(true)

	// To handle pagination, we could skip some rows:
	// req = req.Skip(100)

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
		log.Printf("%s\n", p.Name)
	}

	// Output:

}

func ExampleProjectsAPI_ProjectsCreate() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the project model.
	toCreate := k6.NewCreateProjectApiModel("Example GCk6 project")

	// Then, we create the base request, and set the model.
	req := client.ProjectsAPI.ProjectsCreate(ctx)
	req = req.CreateProjectApiModel(*toCreate)

	// We specify what Stack id we want to make the request for.
	req = req.XStackId(stackID)

	// Finally, we execute the request.
	createdRes, httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("ProjectsCreate request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)
	log.Printf("The project '%s' has been created with the id: %d", createdRes.Name, createdRes.Id)

	// Output:

}

func ExampleProjectsAPI_ProjectsPartialUpdate() {
	// The following example assumes that there is a k6 client initialized and ready to use.
	// Have a look at the `shared.go` file to see how to initialize the client.

	// First we initialize the model.
	toUpdate := k6.NewPatchProjectApiModel("Example GCk6 project (Public API)")

	// Then, we create the base request, and set the model.
	req := client.ProjectsAPI.ProjectsPartialUpdate(ctx, 3736248)
	req = req.PatchProjectApiModel(*toUpdate)

	// We specify what Stack id we want to make the request for.
	req = req.XStackId(stackID)

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

	// First we initialize the base request.
	req := client.ProjectsAPI.ProjectsLimitsRetrieve(ctx, 3736248)

	// We specify what Stack id we want to make the request for.
	req = req.XStackId(stackID)

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

	// Then, we create the base request, and set the model.
	req := client.ProjectsAPI.ProjectsLimitsPartialUpdate(ctx, 3736248)
	req = req.PatchProjectLimitsRequest(*toUpdate)

	// We specify what Stack id we want to make the request for.
	req = req.XStackId(stackID)

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

	// First we initialize the base request.
	req := client.ProjectsAPI.ProjectsDestroy(ctx, 3736246)

	// We specify what Stack id we want to make the request for.
	req = req.XStackId(stackID)

	// Finally, we execute the request.
	httpRes, err := req.Execute()
	if err != nil {
		log.Fatalf("ProjectsDestroy request failed: %s", err.Error())
	}

	log.Printf("Status code: %d", httpRes.StatusCode)

	// Output:

}
