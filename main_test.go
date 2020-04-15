package main

import "net/http"
import "net/http/httptest"
import "testing"

func TestHandleIndexReturnsWithStatusOK(t *testing.T) {
        request, _ := http.NewRequest("GET", "/", nil)
        response := httptest.NewRecorder()

        cityHandler(response, request)
        if response.Code != http.StatusOK {
                t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
        }
}
