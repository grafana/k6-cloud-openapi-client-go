package e2e

import (
	"testing"

	"github.com/grafana/k6-cloud-openapi-client-go/k6"
)

func TestProjectsAPI_ProjectsList(t *testing.T) {
	t.Run("list projects with count", func(t *testing.T) {
		req := testClient.ProjectsAPI.ProjectsList(testCtx).
			XStackId(testStackID).
			Count(true).
			Orderby("created")

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("ProjectsList failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if !res.HasCount() {
			t.Error("Expected count to be present in response")
		}

		if len(res.Value) == 0 {
			t.Error("Expected at least one project in list")
		}
	})

	t.Run("list projects with pagination", func(t *testing.T) {
		req := testClient.ProjectsAPI.ProjectsList(testCtx).
			XStackId(testStackID).
			Top(1).
			Skip(0)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("ProjectsList with pagination failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if len(res.Value) > 1 {
			t.Errorf("Expected max 1 project, got %d", len(res.Value))
		}
	})

	t.Run("list projects with name filter", func(t *testing.T) {
		// First, get the test project to get its name
		retrieveReq := testClient.ProjectsAPI.ProjectsRetrieve(testCtx, testProjectID).
			XStackId(testStackID)

		project, _, err := retrieveReq.Execute()
		if err != nil {
			t.Fatalf("Failed to retrieve test project: %v", err)
		}

		projectName := project.GetName()

		// Now filter by that name
		req := testClient.ProjectsAPI.ProjectsList(testCtx).
			XStackId(testStackID).
			Name(projectName)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("ProjectsList with name filter failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		// Verify all returned projects have the filtered name
		found := false
		for _, p := range res.Value {
			if p.GetName() != projectName {
				t.Errorf("Expected project name %s, got %s", projectName, p.GetName())
			}
			if p.GetId() == testProjectID {
				found = true
			}
		}

		if !found {
			t.Errorf("Expected to find test project ID %d in filtered results", testProjectID)
		}
	})
}

func TestProjectsAPI_ProjectsRetrieve(t *testing.T) {
	req := testClient.ProjectsAPI.ProjectsRetrieve(testCtx, testProjectID).
		XStackId(testStackID)

	res, httpRes, err := req.Execute()
	if err != nil {
		t.Fatalf("ProjectsRetrieve failed: %v", err)
	}

	if httpRes.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
	}

	if res.GetId() != testProjectID {
		t.Errorf("Expected project ID %d, got %d", testProjectID, res.GetId())
	}

	if res.GetName() == "" {
		t.Error("Expected project name to be non-empty")
	}
}

func TestProjectsAPI_ProjectsPartialUpdate(t *testing.T) {
	newName := "go-client-e2e-test-project-updated"
	updateModel := k6.NewPatchProjectApiModel(newName)

	req := testClient.ProjectsAPI.ProjectsPartialUpdate(testCtx, testProjectID).
		PatchProjectApiModel(updateModel).
		XStackId(testStackID)

	httpRes, err := req.Execute()
	if err != nil {
		t.Fatalf("ProjectsPartialUpdate failed: %v", err)
	}

	if httpRes.StatusCode != 204 {
		t.Errorf("Expected status 204, got %d", httpRes.StatusCode)
	}

	// Verify the update
	verifyReq := testClient.ProjectsAPI.ProjectsRetrieve(testCtx, testProjectID).
		XStackId(testStackID)

	project, _, err := verifyReq.Execute()
	if err != nil {
		t.Fatalf("Failed to retrieve updated project: %v", err)
	}

	if project.GetName() != newName {
		t.Errorf("Expected project name %s, got %s", newName, project.GetName())
	}
}

func TestProjectsAPI_ProjectsLimitsRetrieve(t *testing.T) {
	req := testClient.ProjectsAPI.ProjectsLimitsRetrieve(testCtx, testProjectID).
		XStackId(testStackID)

	res, httpRes, err := req.Execute()
	if err != nil {
		t.Fatalf("ProjectsLimitsRetrieve failed: %v", err)
	}

	if httpRes.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
	}

	if res.GetProjectId() != testProjectID {
		t.Errorf("Expected project ID %d, got %d", testProjectID, res.GetProjectId())
	}
}

func TestProjectsAPI_ProjectLimitsRetrieve(t *testing.T) {
	t.Run("retrieve project limits with count", func(t *testing.T) {
		req := testClient.ProjectsAPI.ProjectLimitsRetrieve(testCtx).
			XStackId(testStackID).
			Count(true)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("ProjectLimitsRetrieve failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if !res.HasCount() {
			t.Error("Expected count to be present in response")
		}
	})

	t.Run("retrieve project limits with pagination", func(t *testing.T) {
		req := testClient.ProjectsAPI.ProjectLimitsRetrieve(testCtx).
			XStackId(testStackID).
			Top(5).
			Skip(0)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("ProjectLimitsRetrieve with pagination failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if len(res.Value) > 5 {
			t.Errorf("Expected max 5 project limits, got %d", len(res.Value))
		}
	})
}

func TestProjectsAPI_ProjectsLimitsPartialUpdate(t *testing.T) {
	updateModel := k6.NewPatchProjectLimitsRequest()
	updateModel.SetDurationMaxPerTest(7200)

	req := testClient.ProjectsAPI.ProjectsLimitsPartialUpdate(testCtx, testProjectID).
		PatchProjectLimitsRequest(updateModel).
		XStackId(testStackID)

	httpRes, err := req.Execute()
	if err != nil {
		t.Fatalf("ProjectsLimitsPartialUpdate failed: %v", err)
	}

	if httpRes.StatusCode != 204 {
		t.Errorf("Expected status 204, got %d", httpRes.StatusCode)
	}

	// Verify the update
	verifyReq := testClient.ProjectsAPI.ProjectsLimitsRetrieve(testCtx, testProjectID).
		XStackId(testStackID)

	limits, _, err := verifyReq.Execute()
	if err != nil {
		t.Fatalf("Failed to retrieve updated limits: %v", err)
	}

	if limits.GetDurationMaxPerTest() != 7200 {
		t.Errorf("Expected duration max per test 7200, got %d", limits.GetDurationMaxPerTest())
	}
}
