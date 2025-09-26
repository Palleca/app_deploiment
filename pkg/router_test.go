package pkg_test

import (
	"app_deploiment/pkg"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouterRoutes(t *testing.T) {
	router := pkg.Router()

	tests := []struct {
		method string
		path   string
		status int
	}{
		{"GET", "/health", http.StatusOK},
		{"GET", "/", http.StatusOK},
		{"POST", "/reset", http.StatusOK},
	}

	for _, tt := range tests {
		req, _ := http.NewRequest(tt.method, tt.path, nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code != tt.status {
			t.Errorf("%s %s attendu %d, obtenu %d", tt.method, tt.path, tt.status, rr.Code)
		}
	}
}
