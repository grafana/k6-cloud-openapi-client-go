package k6

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// TestCallAPI_RetryReplaysBody verifies that callAPI restores the request body
// before retrying. The server fails the first attempt with 503 + Connection:
// close (as in the production incident this fixes); without the body reset the
// retry goes out with a drained body and fails with
// "http: ContentLength=N with Body length 0".
func TestCallAPI_RetryReplaysBody(t *testing.T) {
	const payload = "hello-body"

	var attempts int
	var gotBodies []string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		b, _ := io.ReadAll(r.Body)
		gotBodies = append(gotBodies, string(b))
		if attempts == 1 {
			w.Header().Set("Connection", "close")
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	cfg := NewConfiguration()
	cfg.HTTPClient = srv.Client()
	cfg.MaxRetries = 3
	cfg.RetryInterval = time.Millisecond
	c := NewAPIClient(cfg)

	req, err := http.NewRequest(http.MethodPost, srv.URL, strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	req.ContentLength = int64(len(payload))

	resp, err := c.callAPI(req)
	if err != nil {
		t.Fatalf("callAPI returned error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 after retry, got %d", resp.StatusCode)
	}
	if attempts != 2 {
		t.Fatalf("expected 2 attempts, got %d", attempts)
	}
	for i, b := range gotBodies {
		if b != payload {
			t.Fatalf("attempt %d received body %q, want %q", i+1, b, payload)
		}
	}
}
