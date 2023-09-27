package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer_handlePing(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)

	s := NewServer()
	s.mux.ServeHTTP(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expcted %s, got %s", res.Status, http.StatusText(http.StatusOK))
	}
}
