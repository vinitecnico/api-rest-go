package main

import (
	"api-rest-go/controllers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetPeopleShouldReturnSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/contato", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetPeople)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"id":"1","firstname":"John","lastname":"Doe","address":{"city":"City X","state":"State X"}}]`
	result := rr.Body.String()
	result = strings.TrimSuffix(result, "\n")
	if result != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", result, expected)
	}
}

func TestGetPersonShouldReturnSuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/contato/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(controllers.GetPerson)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{}`
	result := rr.Body.String()
	result = strings.TrimSuffix(result, "\n")
	if result != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", result, expected)
	}
}
