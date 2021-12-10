package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func startTestHTTPServer() *httptest.Server {
	pkgData := `[
		{"name":"package1", "version":"1.1"},
		{"name":"package2", "version":"1.0"}
	]`
	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-type", "application/json")
				fmt.Fprint(w, pkgData)
			}))
	return ts
}

func TestFetchRemoteResource(t *testing.T) {
	ts := startTestHTTPServer()
	defer ts.Close()

	packages, err := fetchRemoteResource(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if len(packages) != 2 {
		t.Fatalf("Expected 2 packages, Got: %d", len(packages))
	}

}
