package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeExistingFile(t *testing.T) {
	// Create a test server
	staticDir := "./templates/static"
	handler := http.FileServer(http.Dir(staticDir))
	server := httptest.NewServer(handler)
	defer server.Close()

	// Test a valid file
	resp, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestServeNonExistentFile(t *testing.T) {
	// Create a test server
	staticDir := "./templates/static"
	handler := http.FileServer(http.Dir(staticDir))
	server := httptest.NewServer(handler)
	defer server.Close()

	// Test a non-existent file
	resp, err := http.Get(server.URL + "/homapage.html")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", resp.StatusCode)
	}
}
