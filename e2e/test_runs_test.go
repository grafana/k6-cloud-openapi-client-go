package e2e

import (
	"testing"
	"time"

	"github.com/grafana/k6-cloud-openapi-client-go/k6"
)

func TestTestRunsAPI_TestRunsList(t *testing.T) {
	t.Run("list test runs with count", func(t *testing.T) {
		req := testClient.TestRunsAPI.TestRunsList(testCtx).
			XStackId(testStackID).
			Count(true)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("TestRunsList failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if !res.HasCount() {
			t.Error("Expected count to be present in response")
		}
	})

	t.Run("list test runs with date filters", func(t *testing.T) {
		now := time.Now()
		yesterday := now.Add(-24 * time.Hour)

		req := testClient.TestRunsAPI.TestRunsList(testCtx).
			XStackId(testStackID).
			CreatedAfter(yesterday).
			CreatedBefore(now)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("TestRunsList with date filters failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		// Verify dates are within range
		for _, tr := range res.Value {
			created := tr.GetCreated()
			if created.Before(yesterday) || created.After(now) {
				t.Errorf("Test run created date %v is outside range %v - %v", created, yesterday, now)
			}
		}
	})

	t.Run("list test runs with pagination", func(t *testing.T) {
		req := testClient.TestRunsAPI.TestRunsList(testCtx).
			XStackId(testStackID).
			Top(1).
			Skip(0)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("TestRunsList with pagination failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if len(res.Value) > 1 {
			t.Errorf("Expected max 1 test run, got %d", len(res.Value))
		}
	})
}

func TestTestRunsAPI_LoadTestsTestRunsRetrieve(t *testing.T) {
	t.Run("retrieve load test test runs with count", func(t *testing.T) {
		req := testClient.TestRunsAPI.LoadTestsTestRunsRetrieve(testCtx, testLoadTestID).
			XStackId(testStackID).
			Count(true)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("LoadTestsTestRunsRetrieve failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if !res.HasCount() {
			t.Error("Expected count to be present in response")
		}

		// All returned test runs should belong to our test load test
		for _, tr := range res.Value {
			if tr.GetTestId() != testLoadTestID {
				t.Errorf("Expected test ID %d, got %d", testLoadTestID, tr.GetTestId())
			}
		}
	})

	t.Run("retrieve load test test runs with pagination", func(t *testing.T) {
		req := testClient.TestRunsAPI.LoadTestsTestRunsRetrieve(testCtx, testLoadTestID).
			XStackId(testStackID).
			Top(5).
			Skip(0)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("LoadTestsTestRunsRetrieve with pagination failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if len(res.Value) > 5 {
			t.Errorf("Expected max 5 test runs, got %d", len(res.Value))
		}
	})

	t.Run("retrieve load test test runs with date filters", func(t *testing.T) {
		now := time.Now()
		yesterday := now.Add(-24 * time.Hour)

		req := testClient.TestRunsAPI.LoadTestsTestRunsRetrieve(testCtx, testLoadTestID).
			XStackId(testStackID).
			CreatedAfter(yesterday).
			CreatedBefore(now)

		res, httpRes, err := req.Execute()
		if err != nil {
			t.Fatalf("LoadTestsTestRunsRetrieve with date filters failed: %v", err)
		}

		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		// Verify dates are within range
		for _, tr := range res.Value {
			created := tr.GetCreated()
			if created.Before(yesterday) || created.After(now) {
				t.Errorf("Test run created date %v is outside range %v - %v", created, yesterday, now)
			}
		}
	})
}

func TestTestRunsAPI_TestRunsPartialUpdate(t *testing.T) {
	// First, we need to start a test run to have something to update
	startReq := testClient.LoadTestsAPI.LoadTestsStart(testCtx, testLoadTestID).
		XStackId(testStackID)

	testRun, httpRes, err := startReq.Execute()
	if err != nil {
		t.Fatalf("Failed to start test run: %v", err)
	}

	if httpRes.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
	}

	testRunID := testRun.GetId()

	// Wait a moment for the test run to initialize
	time.Sleep(2 * time.Second)

	// Abort the test run so we don't consume resources
	abortReq := testClient.TestRunsAPI.TestRunsAbort(testCtx, testRunID).
		XStackId(testStackID)

	_, err = abortReq.Execute()
	if err != nil {
		t.Logf("Warning: failed to abort test run: %v", err)
	}

	// Now update the note
	updateModel := k6.NewPatchTestRunApiModel("E2E test run note")

	updateReq := testClient.TestRunsAPI.TestRunsPartialUpdate(testCtx, testRunID).
		PatchTestRunApiModel(updateModel).
		XStackId(testStackID)

	httpRes, err = updateReq.Execute()
	if err != nil {
		t.Fatalf("TestRunsPartialUpdate failed: %v", err)
	}

	if httpRes.StatusCode != 204 {
		t.Errorf("Expected status 204, got %d", httpRes.StatusCode)
	}

	// Verify the update
	verifyReq := testClient.TestRunsAPI.TestRunsRetrieve(testCtx, testRunID).
		XStackId(testStackID)

	updated, _, err := verifyReq.Execute()
	if err != nil {
		t.Fatalf("Failed to retrieve updated test run: %v", err)
	}

	if updated.GetNote() != "E2E test run note" {
		t.Errorf("Expected note 'E2E test run note', got '%s'", updated.GetNote())
	}

	// Clean up: delete the test run
	deleteReq := testClient.TestRunsAPI.TestRunsDestroy(testCtx, testRunID).
		XStackId(testStackID)

	_, err = deleteReq.Execute()
	if err != nil {
		t.Logf("Warning: failed to delete test run: %v", err)
	}
}

func TestTestRunsAPI_TestRunsSaveUnsave(t *testing.T) {
	// Start a test run
	startReq := testClient.LoadTestsAPI.LoadTestsStart(testCtx, testLoadTestID).
		XStackId(testStackID)

	testRun, _, err := startReq.Execute()
	if err != nil {
		t.Fatalf("Failed to start test run: %v", err)
	}

	testRunID := testRun.GetId()

	// Wait a moment
	time.Sleep(2 * time.Second)

	// Abort it
	abortReq := testClient.TestRunsAPI.TestRunsAbort(testCtx, testRunID).
		XStackId(testStackID)

	_, err = abortReq.Execute()
	if err != nil {
		t.Logf("Warning: failed to abort test run: %v", err)
	}

	// Wait for abort to complete
	time.Sleep(2 * time.Second)

	// Save the test run
	saveReq := testClient.TestRunsAPI.TestRunsSave(testCtx, testRunID).
		XStackId(testStackID)

	httpRes, err := saveReq.Execute()
	if err != nil {
		t.Fatalf("TestRunsSave failed: %v", err)
	}

	if httpRes.StatusCode != 204 {
		t.Errorf("Expected status 204, got %d", httpRes.StatusCode)
	}

	// Verify it's saved by checking retention expiry
	verifyReq := testClient.TestRunsAPI.TestRunsRetrieve(testCtx, testRunID).
		XStackId(testStackID)

	saved, _, err := verifyReq.Execute()
	if err != nil {
		t.Fatalf("Failed to retrieve saved test run: %v", err)
	}

	retentionExpiry := saved.GetRetentionExpiry()
	if !retentionExpiry.IsZero() {
		t.Error("Expected retention expiry to be zero (null) for saved test run")
	}

	// Unsave the test run
	unsaveReq := testClient.TestRunsAPI.TestRunsUnsave(testCtx, testRunID).
		XStackId(testStackID)

	httpRes, err = unsaveReq.Execute()
	if err != nil {
		t.Fatalf("TestRunsUnsave failed: %v", err)
	}

	if httpRes.StatusCode != 204 {
		t.Errorf("Expected status 204, got %d", httpRes.StatusCode)
	}

	// Clean up: delete the test run
	deleteReq := testClient.TestRunsAPI.TestRunsDestroy(testCtx, testRunID).
		XStackId(testStackID)

	_, err = deleteReq.Execute()
	if err != nil {
		t.Logf("Warning: failed to delete test run: %v", err)
	}
}

func TestTestRunsAPI_TestRunsScriptRetrieve(t *testing.T) {
	// Start a test run
	startReq := testClient.LoadTestsAPI.LoadTestsStart(testCtx, testLoadTestID).
		XStackId(testStackID)

	testRun, _, err := startReq.Execute()
	if err != nil {
		t.Fatalf("Failed to start test run: %v", err)
	}

	testRunID := testRun.GetId()

	// Wait a moment
	time.Sleep(2 * time.Second)

	// Abort it
	abortReq := testClient.TestRunsAPI.TestRunsAbort(testCtx, testRunID).
		XStackId(testStackID)

	_, err = abortReq.Execute()
	if err != nil {
		t.Logf("Warning: failed to abort test run: %v", err)
	}

	// Retrieve the script
	req := testClient.TestRunsAPI.TestRunsScriptRetrieve(testCtx, testRunID).
		XStackId(testStackID)

	script, httpRes, err := req.Execute()
	if err != nil {
		t.Fatalf("TestRunsScriptRetrieve failed: %v", err)
	}

	if httpRes.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
	}

	if script == "" {
		t.Error("Expected script to be non-empty")
	}

	// Clean up
	deleteReq := testClient.TestRunsAPI.TestRunsDestroy(testCtx, testRunID).
		XStackId(testStackID)

	_, err = deleteReq.Execute()
	if err != nil {
		t.Logf("Warning: failed to delete test run: %v", err)
	}
}

func TestTestRunsAPI_TestRunsDistributionRetrieve(t *testing.T) {
	// Start a test run
	startReq := testClient.LoadTestsAPI.LoadTestsStart(testCtx, testLoadTestID).
		XStackId(testStackID)

	testRun, _, err := startReq.Execute()
	if err != nil {
		t.Fatalf("Failed to start test run: %v", err)
	}

	testRunID := testRun.GetId()

	// Wait for distribution to be available
	time.Sleep(3 * time.Second)

	// Retrieve distribution
	req := testClient.TestRunsAPI.TestRunsDistributionRetrieve(testCtx, testRunID).
		XStackId(testStackID)

	dist, httpRes, err := req.Execute()
	if err != nil {
		// Distribution might not be available immediately, that's okay
		t.Logf("TestRunsDistributionRetrieve returned error (may be expected): %v", err)
	} else {
		if httpRes.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
		}

		if dist.GetDistribution() == nil {
			t.Error("Expected distribution to be non-nil")
		}
	}

	// Abort and clean up
	abortReq := testClient.TestRunsAPI.TestRunsAbort(testCtx, testRunID).
		XStackId(testStackID)

	_, err = abortReq.Execute()
	if err != nil {
		t.Logf("Warning: failed to abort test run: %v", err)
	}

	time.Sleep(2 * time.Second)

	deleteReq := testClient.TestRunsAPI.TestRunsDestroy(testCtx, testRunID).
		XStackId(testStackID)

	_, err = deleteReq.Execute()
	if err != nil {
		t.Logf("Warning: failed to delete test run: %v", err)
	}
}
