package handlers_test

import (
	"app_deploiment/core"
	"app_deploiment/pkg"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	router := pkg.Router()

	req, _ := http.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}
	if rr.Body.String() != "OK" {
		t.Errorf("Expected body 'OK', got %s", rr.Body.String())
	}
}

func TestHomeHandler(t *testing.T) {
	router := pkg.Router()

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rr.Code)
	}
}

func TestBlocksHandler(t *testing.T) {
	router := pkg.Router()

	// Reset blockchain pour avoir au moins le bloc genesis
	core.ResetBlockchain()

	req, _ := http.NewRequest("GET", "/blocks", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rr.Code)
	}
	if rr.Body.Len() == 0 {
		t.Error("Expected response body with blockchain data")
	}
}

func TestTransactionsHandler(t *testing.T) {
	router := pkg.Router()

	tx := core.Transaction{Sender: "Bruno", Receiver: "toto", Amount: 42}
	txJSON, _ := json.Marshal(tx)
	req, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(txJSON))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", rr.Code)
	}
}

func TestMineHandler(t *testing.T) {
	router := pkg.Router()

	// Ajoutons d'abord une transaction pour que le bloc miné ait du contenu
	core.ResetBlockchain()
	tx := core.Transaction{Sender: "Bruno", Receiver: "toto", Amount: 99}
	txJSON, _ := json.Marshal(tx)
	reqTx, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(txJSON))
	reqTx.Header.Set("Content-Type", "application/json")
	rrTx := httptest.NewRecorder()
	router.ServeHTTP(rrTx, reqTx)

	// Miner un bloc
	req, _ := http.NewRequest("POST", "/mine", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rr.Code)
	}
	if rr.Body.Len() == 0 {
		t.Error("Expected response body with mined block data")
	}
}

func TestResetHandler(t *testing.T) {
	router := pkg.Router()

	req, _ := http.NewRequest("POST", "/reset", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rr.Code)
	}
}
