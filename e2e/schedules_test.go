package e2e

import (
	"testing"
	"time"

	"github.com/grafana/k6-cloud-openapi-client-go/k6"
)

func TestSchedulesAPI_LoadTestsScheduleCreate(t *testing.T) {
	// Create a schedule
	starts := time.Now().Add(24 * time.Hour) // Schedule for tomorrow
	createReq := k6.NewCreateScheduleRequest(starts)

	req := testClient.SchedulesAPI.LoadTestsScheduleCreate(testCtx, testLoadTestID).
		CreateScheduleRequest(createReq).
		XStackId(testStackID)

	schedule, httpRes, err := req.Execute()
	if err != nil {
		t.Fatalf("LoadTestsScheduleCreate failed: %v", err)
	}

	if httpRes.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
	}

	scheduleID := schedule.GetId()

	if schedule.GetLoadTestId() != testLoadTestID {
		t.Errorf("Expected load test ID %d, got %d", testLoadTestID, schedule.GetLoadTestId())
	}

	// Retrieve the schedule
	retrieveReq := testClient.SchedulesAPI.LoadTestsScheduleRetrieve(testCtx, testLoadTestID).
		XStackId(testStackID)

	retrieved, httpRes, err := retrieveReq.Execute()
	if err != nil {
		t.Fatalf("LoadTestsScheduleRetrieve failed: %v", err)
	}

	if httpRes.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
	}

	if retrieved.GetId() != scheduleID {
		t.Errorf("Expected schedule ID %d, got %d", scheduleID, retrieved.GetId())
	}

	// Clean up
	deleteReq := testClient.SchedulesAPI.SchedulesDestroy(testCtx, scheduleID).
		XStackId(testStackID)

	_, err = deleteReq.Execute()
	if err != nil {
		t.Logf("Warning: failed to delete schedule: %v", err)
	}
}

func TestSchedulesAPI_SchedulesList(t *testing.T) {
	t.Run("list schedules with count", func(t *testing.T) {
		req := testClient.SchedulesAPI.SchedulesList(testCtx).
			XStackId(testStackID).
			Count(true)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("SchedulesList failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if !res.HasCount() {
			t.Error("Expected count to be present in response")
		}
	})

	t.Run("list schedules with pagination", func(t *testing.T) {
		req := testClient.SchedulesAPI.SchedulesList(testCtx).
			XStackId(testStackID).
			Top(1).
			Skip(0)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("SchedulesList with pagination failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if len(res.Value) > 1 {
			t.Errorf("Expected max 1 schedule, got %d", len(res.Value))
		}
	})
}

func TestSchedulesAPI_SchedulesRetrieve(t *testing.T) {
	// Create a schedule first
	starts := time.Now().Add(24 * time.Hour)
	createReq := k6.NewCreateScheduleRequest(starts)

	createScheduleReq := testClient.SchedulesAPI.LoadTestsScheduleCreate(testCtx, testLoadTestID).
		CreateScheduleRequest(createReq).
		XStackId(testStackID)

	schedule, _, err := createScheduleReq.Execute()
	if err != nil {
		t.Fatalf("Failed to create schedule: %v", err)
	}

	scheduleID := schedule.GetId()

	// Retrieve by schedule ID
	req := testClient.SchedulesAPI.SchedulesRetrieve(testCtx, scheduleID).
		XStackId(testStackID)

	res, httpRes, err := req.Execute()
	if err != nil {
		t.Fatalf("SchedulesRetrieve failed: %v", err)
	}

	if httpRes.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
	}

	if res.GetId() != scheduleID {
		t.Errorf("Expected schedule ID %d, got %d", scheduleID, res.GetId())
	}

	// Clean up
	deleteReq := testClient.SchedulesAPI.SchedulesDestroy(testCtx, scheduleID).
		XStackId(testStackID)

	_, err = deleteReq.Execute()
	if err != nil {
		t.Logf("Warning: failed to delete schedule: %v", err)
	}
}

func TestSchedulesAPI_ScheduleActivateDeactivate(t *testing.T) {
	// Create a schedule
	starts := time.Now().Add(24 * time.Hour)
	createReq := k6.NewCreateScheduleRequest(starts)

	createScheduleReq := testClient.SchedulesAPI.LoadTestsScheduleCreate(testCtx, testLoadTestID).
		CreateScheduleRequest(createReq).
		XStackId(testStackID)

	schedule, _, err := createScheduleReq.Execute()
	if err != nil {
		t.Fatalf("Failed to create schedule: %v", err)
	}

	scheduleID := schedule.GetId()

	// Deactivate the schedule
	deactivateReq := testClient.SchedulesAPI.ScheduleDeactivate(testCtx, scheduleID).
		XStackId(testStackID)

	httpRes, err := deactivateReq.Execute()
	if err != nil {
		t.Fatalf("ScheduleDeactivate failed: %v", err)
	}

	if httpRes.StatusCode != 204 {
		t.Errorf("Expected status 204, got %d", httpRes.StatusCode)
	}

	// Verify it's deactivated
	verifyReq := testClient.SchedulesAPI.SchedulesRetrieve(testCtx, scheduleID).
		XStackId(testStackID)

	deactivated, _, err := verifyReq.Execute()
	if err != nil {
		t.Fatalf("Failed to retrieve deactivated schedule: %v", err)
	}

	if !deactivated.GetDeactivated() {
		t.Error("Expected schedule to be deactivated")
	}

	// Activate the schedule
	activateReq := testClient.SchedulesAPI.ScheduleActivate(testCtx, scheduleID).
		XStackId(testStackID)

	httpRes, err = activateReq.Execute()
	if err != nil {
		t.Fatalf("ScheduleActivate failed: %v", err)
	}

	if httpRes.StatusCode != 204 {
		t.Errorf("Expected status 204, got %d", httpRes.StatusCode)
	}

	// Verify it's activated
	verifyReq = testClient.SchedulesAPI.SchedulesRetrieve(testCtx, scheduleID).
		XStackId(testStackID)

	activated, _, err := verifyReq.Execute()
	if err != nil {
		t.Fatalf("Failed to retrieve activated schedule: %v", err)
	}

	if activated.GetDeactivated() {
		t.Error("Expected schedule to be activated (not deactivated)")
	}

	// Clean up
	deleteReq := testClient.SchedulesAPI.SchedulesDestroy(testCtx, scheduleID).
		XStackId(testStackID)

	_, err = deleteReq.Execute()
	if err != nil {
		t.Logf("Warning: failed to delete schedule: %v", err)
	}
}
