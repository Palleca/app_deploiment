package handlers_test

import (
	"app_deploiment/core"
	"app_deploiment/pkg"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// GET /transactions en JSON avec une transaction
func TestGetTransactionsHandler_JSON(t *testing.T) {
	router := pkg.Router()
	core.ResetBlockchain()

	// On ajoute une transaction pour tester le retour JSON
	core.AddTransaction(core.Transaction{Sender: "Alice", Receiver: "Bob", Amount: 10})

	req, _ := http.NewRequest("GET", "/transactions", nil)
	req.Header.Set("Accept", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("GET /transactions (JSON) attendu 200, obtenu %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), `"Alice"`) {
		t.Errorf("Réponse JSON incorrecte: %s", rr.Body.String())
	}
}

// GET /transactions en JSON sans transaction
func TestGetTransactionsHandler_JSON_NoTx(t *testing.T) {
	router := pkg.Router()
	core.ResetBlockchain() // aucune transaction

	req, _ := http.NewRequest("GET", "/transactions", nil)
	req.Header.Set("Accept", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("GET /transactions (JSON sans tx) attendu 200, obtenu %d", rr.Code)
	}

	// On parse le body comme JSON pour vérifier qu'il est bien vide
	var data []core.Transaction
	if err := json.Unmarshal(rr.Body.Bytes(), &data); err != nil {
		t.Fatalf("Réponse JSON non valide: %v", err)
	}
	if len(data) != 0 {
		t.Errorf("Réponse JSON attendue vide, obtenu: %v", data)
	}
}

// POST /transactions en JSON valide
func TestCreateTransactionHandler_JSON_Valid(t *testing.T) {
	router := pkg.Router()
	core.ResetBlockchain()

	body := `{"sender":"Alice","receiver":"Bob","amount":42}`
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("POST /transactions (JSON) attendu 201, obtenu %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), `"status"`) {
		t.Errorf("Réponse JSON incorrecte: %s", rr.Body.String())
	}
}

// POST /transactions JSON invalide (mauvais payload)
func TestCreateTransactionHandler_JSON_InvalidPayload(t *testing.T) {
	router := pkg.Router()
	core.ResetBlockchain()

	body := `{"sender": "","receiver":"Bob","amount":-5}`
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("POST /transactions (JSON invalide) attendu 400, obtenu %d", rr.Code)
	}
}

// POST /transactions JSON malformé
func TestCreateTransactionHandler_JSON_Malformed(t *testing.T) {
	router := pkg.Router()
	core.ResetBlockchain()

	body := `{"sender": "Alice", "receiver": "Bob" "amount":42}` // virgule manquante
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("POST /transactions (JSON malformé) attendu 400, obtenu %d", rr.Code)
	}
}

// MineBlockHandler sans transactions
func TestMineBlockHandler_NoTransactions(t *testing.T) {
	router := pkg.Router()
	core.ResetBlockchain()

	req, _ := http.NewRequest("POST", "/mine", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("POST /mine sans transactions attendu 400, obtenu %d", rr.Code)
	}
}

// MineBlockHandler avec transactions
func TestMineBlockHandler_WithTransactions(t *testing.T) {
	router := pkg.Router()
	core.ResetBlockchain()

	// On ajoute une transaction pour pouvoir miner
	core.AddTransaction(core.Transaction{Sender: "Alice", Receiver: "Bob", Amount: 50})

	req, _ := http.NewRequest("POST", "/mine", nil)
	req.Header.Set("Accept", "application/json") // on force JSON
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("POST /mine avec transactions attendu 200, obtenu %d", rr.Code)
	}

	// On parse la réponse JSON pour vérifier la présence de "index"
	var resp map[string]any
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("Réponse JSON invalide: %v", err)
	}
	if _, ok := resp["index"]; !ok {
		t.Errorf("Réponse JSON du bloc miné incorrecte: %v", resp)
	}
}

// GetTransactionsHandler (HTML avec transactions)
func TestGetTransactionsHandler_HTML_WithTx(t *testing.T) {
	router := pkg.Router()
	core.ResetBlockchain()
	core.AddTransaction(core.Transaction{Sender: "Alice", Receiver: "Bob", Amount: 12})

	req, _ := http.NewRequest("GET", "/transactions", nil)
	req.Header.Set("Accept", "text/html")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("GET /transactions (HTML avec tx) attendu 200, obtenu %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "Alice") {
		t.Errorf("Réponse HTML ne contient pas la transaction ajoutée: %s", rr.Body.String())
	}
}

// CreateTransactionHandler formulaire avec montant invalide
func TestCreateTransactionHandler_FormInvalidAmount(t *testing.T) {
	router := pkg.Router()
	core.ResetBlockchain()

	form := "sender=Alice&receiver=Bob&amount=abc" // montant invalide
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBufferString(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("POST /transactions (form amount invalide) attendu 400, obtenu %d", rr.Code)
	}
}

// ✅ CreateTransactionHandler formulaire valide
func TestCreateTransactionHandler_FormValidFull(t *testing.T) {
	router := pkg.Router()
	core.ResetBlockchain()

	form := "sender=Alice&receiver=Bob&amount=25"
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBufferString(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther { // 303 attendu
		t.Fatalf("POST /transactions (form valide) attendu 303, obtenu %d", rr.Code)
	}

	// Vérifie que la transaction a bien été ajoutée
	if len(core.Transactions) == 0 {
		t.Errorf("Transaction non ajoutée après soumission formulaire")
	}
}

// GetTransactionsHandler (HTML sans transactions)
func TestGetTransactionsHandler_HTML_NoTx(t *testing.T) {
	router := pkg.Router()
	core.ResetBlockchain() // aucune transaction

	req, _ := http.NewRequest("GET", "/transactions", nil)
	req.Header.Set("Accept", "text/html")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("GET /transactions (HTML sans tx) attendu 200, obtenu %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), "Aucune transaction en attente") {
		t.Errorf("Réponse HTML incorrecte, obtenu: %s", rr.Body.String())
	}
}
