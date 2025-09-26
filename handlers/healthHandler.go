package handlers

import (
	"fmt"
	"net/http"
)

// GET /health
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}
