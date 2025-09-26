package handlers_test

import (
	"app_deploiment/core"
	"app_deploiment/pkg"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Force la branche HTML de GetTransactionsHandler
func TestGetTransactionsHandler_HTML(t *testing.T) {
	router := pkg.Router()
	core.ResetBlockchain()

	req, _ := http.NewRequest("GET", "/transactions", nil)
	req.Header.Set("Accept", "text/html") // 👈 on force le mode HTML
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("GET /transactions (HTML) attendu 200, obtenu %d", rr.Code)
	}
	if !bytes.Contains(rr.Body.Bytes(), []byte("<html>")) {
		t.Errorf("Réponse GET /transactions (HTML) incorrecte, got: %s", rr.Body.String())
	}
}

// Force la branche Formulaire de CreateTransactionHandler
func TestCreateTransactionHandler_FormValid(t *testing.T) {
	router := pkg.Router()
	core.ResetBlockchain()

	form := "sender=Alice&receiver=Bob&amount=12.5"
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBufferString(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") // 👈 on force le Form

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther { // 303 redirection
		t.Fatalf("POST /transactions (form) attendu 303, obtenu %d", rr.Code)
	}
}
