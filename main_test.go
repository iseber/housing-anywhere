package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNavigationHandler_ShouldReturn_StatusMethodNotAllowed_WhenRequestMethodIsNot_Post(t *testing.T){
	req, err := http.NewRequest("GET", "", strings.NewReader("{ \"x\":\"10.25\", \"y\":\"12.23\", \"z\":\"15.12\", \"vel\":\"40\"}"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(navigation)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestNavigationHandler_ShouldReturn_StatusBadRequest_WhenRequestBodyIsInvalid(t *testing.T){
	req, err := http.NewRequest("POST", "", strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(navigation)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestNavigationHandler(t *testing.T){
	req, err := http.NewRequest("POST", "", strings.NewReader("{ \"x\":\"10.25\", \"y\":\"12.23\", \"z\":\"15.12\", \"vel\":\"40\"}"))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(navigation)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"loc": 77.60}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthcheck)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}