package handlers

import (
	"app_deploiment/core"
	"fmt"
	"net/http"
)

// GET or
func ResetHandler(w http.ResponseWriter, r *http.Request) {
	core.ResetBlockchain()
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, " ")
}
