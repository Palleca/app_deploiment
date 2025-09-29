package tests

import (
	"app_deploiment/core"
	"app_deploiment/pkg"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFullIntegrationFlow(t *testing.T) {
	router := pkg.Router()

	// 1. Reset blockchain
	reqReset, _ := http.NewRequest("POST", "/reset", nil)
	rrReset := httptest.NewRecorder()
	router.ServeHTTP(rrReset, reqReset)
	if rrReset.Code != http.StatusOK {
		t.Fatalf("Reset attendu 200, obtenu %d", rrReset.Code)
	}

	// 2. Ajouter une transaction
	tx := core.Transaction{Sender: "Bruno", Receiver: "Toi", Amount: 10}
	txJSON, _ := json.Marshal(tx)
	reqTx, _ := http.NewRequest("POST", "/transactions", bytes.NewBuffer(txJSON))
	reqTx.Header.Set("Content-Type", "application/json")
	rrTx := httptest.NewRecorder()
	router.ServeHTTP(rrTx, reqTx)

	if rrTx.Code != http.StatusCreated {
		t.Fatalf("Transaction: statut attendu 201, obtenu %d", rrTx.Code)
	}

	// 3. Miner un bloc
	reqMine, _ := http.NewRequest("POST", "/mine", nil)
	rrMine := httptest.NewRecorder()
	router.ServeHTTP(rrMine, reqMine)

	if rrMine.Code != http.StatusOK {
		t.Fatalf("Mine: statut attendu 200, obtenu %d", rrMine.Code)
	}

	// 4. Vérifier que la transaction est dans la blockchain
	reqBlocks, _ := http.NewRequest("GET", "/blocks", nil)
	rrBlocks := httptest.NewRecorder()
	router.ServeHTTP(rrBlocks, reqBlocks)

	if rrBlocks.Code != http.StatusOK {
		t.Fatalf("Blocks: statut attendu 200, obtenu %d", rrBlocks.Code)
	}

	if !bytes.Contains(rrBlocks.Body.Bytes(), []byte("Bruno")) ||
		!bytes.Contains(rrBlocks.Body.Bytes(), []byte("Toi")) {
		t.Error("Transaction Bruno -> Toi non trouvée dans la blockchain")
	}

	// 5. Reset à nouveau et vérifier qu'il ne reste plus que le bloc genesis
	reqReset2, _ := http.NewRequest("POST", "/reset", nil)
	rrReset2 := httptest.NewRecorder()
	router.ServeHTTP(rrReset2, reqReset2)
	if rrReset2.Code != http.StatusOK {
		t.Fatalf("Reset attendu 200, obtenu %d", rrReset2.Code)
	}

	reqBlocks2, _ := http.NewRequest("GET", "/blocks", nil)
	rrBlocks2 := httptest.NewRecorder()
	router.ServeHTTP(rrBlocks2, reqBlocks2)
	if rrBlocks2.Code != http.StatusOK {
		t.Fatalf("Blocks: statut attendu 200, obtenu %d", rrBlocks2.Code)
	}

	// Le genesis block ne doit pas contenir
	if bytes.Contains(rrBlocks2.Body.Bytes(), []byte("Bruno")) ||
		bytes.Contains(rrBlocks2.Body.Bytes(), []byte("Toi")) {
		t.Error("La blockchain devrait avoir été réinitialisée au genesis uniquement")
	}
}
