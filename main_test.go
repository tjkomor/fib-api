package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFibHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/fibonacci", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(FibHandler)

	handler.ServeHTTP(responseRecorder, req)

    if status := responseRecorder.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := "FIBONACCI ENDPOINT"
    if responseRecorder.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            responseRecorder.Body.String(), expected)

		}
	}
