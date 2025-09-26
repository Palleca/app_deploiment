package core_test

import (
	"app_deploiment/core"
	"strings"
	"testing"
)

func TestBlockToJSON(t *testing.T) {
	block := core.Block{
		Index:        1,
		Timestamp:    "2025-09-26T12:00:00Z",
		Transactions: []core.Transaction{{Sender: "Alice", Receiver: "Bob", Amount: 42}},
		PreviousHash: "prevhash",
		Hash:         "hash",
	}

	data, err := core.BlockToJSON(block)
	if err != nil {
		t.Fatalf("BlockToJSON returned error: %v", err)
	}
	if !strings.Contains(string(data), "Alice") {
		t.Errorf("BlockToJSON incorrect, got: %s", string(data))
	}
}

func TestTransactionToJSON(t *testing.T) {
	tx := core.Transaction{Sender: "Alice", Receiver: "Bob", Amount: 99}
	data, err := core.TransactionToJSON(tx)
	if err != nil {
		t.Fatalf("TransactionToJSON returned error: %v", err)
	}
	if !strings.Contains(string(data), "Alice") {
		t.Errorf("TransactionToJSON incorrect, got: %s", string(data))
	}
}

func TestBlocksToJSON(t *testing.T) {
	blocks := []core.Block{
		{Index: 1, Timestamp: "2025-09-26T12:00:00Z", PreviousHash: "0", Hash: "hash1"},
		{Index: 2, Timestamp: "2025-09-26T12:01:00Z", PreviousHash: "hash1", Hash: "hash2"},
	}
	data, err := core.BlocksToJSON(blocks)
	if err != nil {
		t.Fatalf("BlocksToJSON returned error: %v", err)
	}
	if !strings.Contains(string(data), "hash2") {
		t.Errorf("BlocksToJSON incorrect, got: %s", string(data))
	}
}

func TestTransactionsToJSON(t *testing.T) {
	txs := []core.Transaction{
		{Sender: "Alice", Receiver: "Bob", Amount: 10},
		{Sender: "Charlie", Receiver: "Dave", Amount: 20},
	}
	data, err := core.TransactionsToJSON(txs)
	if err != nil {
		t.Fatalf("TransactionsToJSON returned error: %v", err)
	}
	if !strings.Contains(string(data), "Charlie") {
		t.Errorf("TransactionsToJSON incorrect, got: %s", string(data))
	}
}
