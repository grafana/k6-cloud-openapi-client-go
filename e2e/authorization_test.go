package e2e

import (
	"testing"
)

func TestAuthorizationAPI_Auth(t *testing.T) {
	req := testClient.AuthorizationAPI.Auth(testCtx).
		XStackUrl(testStackURL)

	res, httpRes, err := req.Execute()
	if err != nil {
		t.Fatalf("Auth failed: %v", err)
	}

	if httpRes.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", httpRes.StatusCode)
	}

	if res.GetStackId() == 0 {
		t.Error("Expected stack ID to be non-zero")
	}

	if res.GetDefaultProjectId() == 0 {
		t.Error("Expected default project ID to be non-zero")
	}

	// Verify the returned stack ID matches our test stack ID
	if res.GetStackId() != testStackID {
		t.Logf("Warning: Auth returned stack ID %d, but we're using %d", res.GetStackId(), testStackID)
	}
}
