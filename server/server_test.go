package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloServer(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Create a new HTTP handler
	handler := http.HandlerFunc(HelloServer)

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "hello server!"
	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(body), expected)
	}
}
