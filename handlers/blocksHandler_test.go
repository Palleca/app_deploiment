package handlers

import (
	"app_deploiment/core"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Couvre la branche JSON (Accept: application/json)
func TestGetBlocksHandler_JSON(t *testing.T) {
	core.ResetBlockchain()
	core.Blockchain = []core.Block{
		{
			Index:        1,
			Hash:         "hash1",
			PreviousHash: "genesis",
			Transactions: []core.Transaction{
				{Sender: "Bruno", Receiver: "TOTO", Amount: 10},
			},
		},
	}

	req := httptest.NewRequest("GET", "/blocks", nil)
	req.Header.Set("Accept", "application/json")
	w := httptest.NewRecorder()

	GetBlocksHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "Bruno") {
		t.Fatalf("expected transaction in JSON, got: %s", w.Body.String())
	}
}

// Couvre la branche HTML + cas sans transactions
func TestGetBlocksHandler_HTML_NoTx(t *testing.T) {
	core.ResetBlockchain()
	core.Blockchain = []core.Block{
		{Index: 1, Hash: "hash1", PreviousHash: "genesis", Transactions: nil},
	}

	req := httptest.NewRequest("GET", "/blocks", nil) // pas d'Accept JSON => HTML
	w := httptest.NewRecorder()

	GetBlocksHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if !strings.Contains(w.Body.String(), "Aucune transaction") {
		t.Fatalf("expected empty tx message, got: %s", w.Body.String())
	}
	if !strings.Contains(w.Body.String(), "hash1") {
		t.Fatalf("expected block hash in HTML, got: %s", w.Body.String())
	}
}

// Couvre la branche HTML + cas avec transactions (boucle for sur les tx)
func TestGetBlocksHandler_HTML_WithTx(t *testing.T) {
	core.ResetBlockchain()
	core.Blockchain = []core.Block{
		{
			Index:        2,
			Hash:         "hash2",
			PreviousHash: "hash1",
			Transactions: []core.Transaction{
				{Sender: "Bruno", Receiver: "toto", Amount: 12.5},
			},
		},
	}

	req := httptest.NewRequest("GET", "/blocks", nil) // HTML
	w := httptest.NewRecorder()

	GetBlocksHandler(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	// Vérifie que la transaction s'affiche dans le HTML
	if !strings.Contains(w.Body.String(), "Bruno") || !strings.Contains(w.Body.String(), "toto") {
		t.Fatalf("expected rendered transaction, got: %s", w.Body.String())
	}
}
