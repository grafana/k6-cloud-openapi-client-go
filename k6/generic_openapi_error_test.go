package k6

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGenericOpenAPIError_StatusCode verifies that GenericOpenAPIError exposes
// the HTTP status code of the response that produced the error, across several
// distinct generated Api services (the field/accessor is template-driven, so
// this guards against the template change only reaching some operations).
func TestGenericOpenAPIError_StatusCode(t *testing.T) {
	const code = http.StatusForbidden // 403

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		_, _ = w.Write([]byte(`{"error":{"message":"nope"}}`))
	}))
	defer srv.Close()

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: srv.URL}}
	cfg.HTTPClient = srv.Client()
	client := NewAPIClient(cfg)
	ctx := context.Background()

	assertCode := func(t *testing.T, name string, err error) {
		t.Helper()
		if err == nil {
			t.Fatalf("%s: expected error, got nil", name)
		}
		var apiErr *GenericOpenAPIError
		if !errors.As(err, &apiErr) {
			t.Fatalf("%s: error is not *GenericOpenAPIError: %T", name, err)
		}
		if got := apiErr.StatusCode(); got != code {
			t.Fatalf("%s: StatusCode() = %d, want %d", name, got, code)
		}
	}

	_, _, err := client.ProjectsAPI.ProjectsRetrieve(ctx, 1).XStackId(1).Execute()
	assertCode(t, "ProjectsAPI.ProjectsRetrieve", err)

	_, _, err = client.ProjectsAPI.ProjectsList(ctx).XStackId(1).Execute()
	assertCode(t, "ProjectsAPI.ProjectsList", err)

	_, _, err = client.LoadZonesAPI.LoadZonesList(ctx).XStackId(1).Execute()
	assertCode(t, "LoadZonesAPI.LoadZonesList", err)
}

// TestGenericOpenAPIError_StatusCodeZeroBeforeRequest verifies that an error
// produced before the request is sent (here, an unresolvable server URL)
// yields a GenericOpenAPIError with StatusCode() == 0.
func TestGenericOpenAPIError_StatusCodeZeroBeforeRequest(t *testing.T) {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{} // no server to resolve => pre-request error
	client := NewAPIClient(cfg)

	_, _, err := client.ProjectsAPI.ProjectsList(context.Background()).XStackId(1).Execute()
	if err == nil {
		t.Fatal("expected a pre-request error, got nil")
	}
	var apiErr *GenericOpenAPIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("error is not *GenericOpenAPIError: %T", err)
	}
	if got := apiErr.StatusCode(); got != 0 {
		t.Fatalf("StatusCode() = %d, want 0 for pre-request failure", got)
	}
}
