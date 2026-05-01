package api

import (
	//"encoding/json"
	"net/http"
	"net/http/httptest"
	//"strings"
	"testing"
)

func newTestHandler() http.Handler {
	s := NewServer(nil)
	mux := http.NewServeMux()
	h := HandlerFromMux(s, mux)
	return h
}

func TestCreate(t *testing.T) {
	handler := newTestHandler()

	tests := []struct{
		name string
		method string
		path string
		wantCode int
		wantBody string
	}{
		{
			"GET repositories found",
			"GET",
			"/v1/repositories",
			http.StatusOK,
			"{}",
		},
	}

	for _, testInput := range tests {
		t.Run(testInput.name, func(t *testing.T) {

			req := httptest.NewRequest(testInput.method, testInput.path, nil)
			rr := httptest.NewRecorder()

			handler.ServeHTTP(rr, req)

			if rr.Code != testInput.wantCode {
				t.Errorf("expected error code '%v', got '%v'", testInput.wantCode, rr.Code)
			}
		})
	}
}
