package e2e

import (
	"testing"

	"github.com/grafana/k6-cloud-openapi-client-go/k6"
)

// TestLoadZonesAPI_E2EFlow tests the complete flow of load zone operations
// Following the sequence from the TypeScript test: api_load_zone.ts
func TestLoadZonesAPI_E2EFlow(t *testing.T) {
	// Step 1: List all load zones and verify our private load zone is present
	t.Run("1. list all load zones and find private load zone", func(t *testing.T) {
		req := testClient.LoadZonesAPI.LoadZonesList(testCtx).
			XStackId(testStackID)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("LoadZonesList failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if len(res.Value) == 0 {
			t.Fatal("Expected at least one load zone")
		}

		// Find our private load zone
		found := false
		for _, lz := range res.Value {
			if lz.GetK6LoadZoneId() == testLoadZoneK6ID {
				found = true
				testLoadZoneID = lz.GetId() // Update the ID from the API response
				t.Logf("Found private load zone: %s (ID: %d)", testLoadZoneK6ID, testLoadZoneID)
				break
			}
		}

		if !found {
			t.Errorf("Private load zone %s not found in load zones list", testLoadZoneK6ID)
		}
	})

	// Step 2: Filter load zones by k6_load_zone_id
	t.Run("2. filter load zones by k6_load_zone_id", func(t *testing.T) {
		req := testClient.LoadZonesAPI.LoadZonesList(testCtx).
			XStackId(testStackID).
			K6LoadZoneId(testLoadZoneK6ID)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("LoadZonesList with filter failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if len(res.Value) != 1 {
			t.Fatalf("Expected exactly 1 load zone, got %d", len(res.Value))
		}

		if res.Value[0].GetK6LoadZoneId() != testLoadZoneK6ID {
			t.Errorf("Expected k6_load_zone_id %s, got %s", testLoadZoneK6ID, res.Value[0].GetK6LoadZoneId())
		}

		testLoadZoneID = res.Value[0].GetId()
		t.Logf("Filtered load zone ID: %d", testLoadZoneID)
	})

	// Step 3: Update load zone allowed projects (add our test project)
	t.Run("3. update load zone allowed projects", func(t *testing.T) {
		projectRef := k6.NewAllowedProjectToUpdateApiModel(testProjectID)

		updateModel := k6.NewUpdateAllowedProjectsListApiModel([]k6.AllowedProjectToUpdateApiModel{*projectRef})

		req := testClient.LoadZonesAPI.LoadZonesAllowedProjectsUpdate(testCtx, testLoadZoneID).
			UpdateAllowedProjectsListApiModel(updateModel).
			XStackId(testStackID)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("LoadZonesAllowedProjectsUpdate failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if len(res.Value) != 1 {
			t.Errorf("Expected 1 allowed project, got %d", len(res.Value))
		}

		if len(res.Value) > 0 && res.Value[0].GetId() != testProjectID {
			t.Errorf("Expected project ID %d, got %d", testProjectID, res.Value[0].GetId())
		}
	})

	// Step 4: List projects allowed to use the load zone
	t.Run("4. list projects allowed to use the load zone", func(t *testing.T) {
		req := testClient.LoadZonesAPI.LoadZonesAllowedProjectsRetrieve(testCtx, testLoadZoneID).
			XStackId(testStackID)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("LoadZonesAllowedProjectsRetrieve failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if len(res.Value) != 1 {
			t.Errorf("Expected 1 allowed project, got %d", len(res.Value))
		}

		if len(res.Value) > 0 && res.Value[0].GetId() != testProjectID {
			t.Errorf("Expected project ID %d, got %d", testProjectID, res.Value[0].GetId())
		}
	})

	// Step 5: Update project to remove the allowed load zone
	t.Run("5. update project to remove allowed load zone", func(t *testing.T) {
		updateModel := k6.NewUpdateAllowedLoadZonesListApiModel([]k6.AllowedLoadZoneToUpdateApiModel{}) // Empty list to remove all

		req := testClient.LoadZonesAPI.ProjectsAllowedLoadZonesUpdate(testCtx, testProjectID).
			UpdateAllowedLoadZonesListApiModel(updateModel).
			XStackId(testStackID)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("ProjectsAllowedLoadZonesUpdate failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if len(res.Value) != 0 {
			t.Errorf("Expected 0 allowed load zones, got %d", len(res.Value))
		}
	})

	// Step 6: List load zones allowed for project (verify it's empty)
	t.Run("6. list load zones allowed for project", func(t *testing.T) {
		req := testClient.LoadZonesAPI.ProjectsAllowedLoadZonesRetrieve(testCtx, testProjectID).
			XStackId(testStackID)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("ProjectsAllowedLoadZonesRetrieve failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if len(res.Value) != 0 {
			t.Errorf("Expected 0 allowed load zones after removal, got %d", len(res.Value))
		}
	})
}
