package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/majidypd/goblueprint/internal/app"
)

func TestHealthHandler(t *testing.T) {
    // Initialize the handler
    controller := app.NewHealthHandler()

    // Create a request to pass to our handler
    req, err := http.NewRequest("GET", "/health", nil)
    if err != nil {
        t.Fatal(err)
    }

    // Create a ResponseRecorder to record the response
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(controller.Health)

    // Call the handler
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the response body is what we expect
    expected := `Hello, "/health"`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}
