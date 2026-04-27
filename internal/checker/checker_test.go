package checker

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCheckURLsReturnsStatusCode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	}))
	defer server.Close()

	results := CheckURLs([]string{server.URL}, time.Second)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].StatusCode != http.StatusAccepted {
		t.Fatalf("expected status %d, got %d", http.StatusAccepted, results[0].StatusCode)
	}

	if results[0].Error != "" {
		t.Fatalf("expected empty error, got %q", results[0].Error)
	}
}

func TestCheckURLsReportsError(t *testing.T) {
	results := CheckURLs([]string{"http://127.0.0.1:1"}, 50*time.Millisecond)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}

	if results[0].Error == "" {
		t.Fatal("expected an error for unreachable host")
	}
}
