package core_test

import (
	"app_deploiment/core"
	"testing"
)

func TestAddTransaction(t *testing.T) {
	core.ResetBlockchain()
	tx := core.Transaction{Sender: "Alice", Receiver: "Bob", Amount: 10}
	core.AddTransaction(tx)

	if len(core.Transactions) != 1 {
		t.Errorf("Attendu 1 transaction, obtenu %d", len(core.Transactions))
	}
}

func TestMineBlock(t *testing.T) {
	core.ResetBlockchain()
	tx := core.Transaction{Sender: "Alice", Receiver: "Bob", Amount: 10}
	core.AddTransaction(tx)

	block := core.MineBlock()

	if len(block.Transactions) == 0 {
		t.Error("Le bloc miné devrait contenir une transaction")
	}
	if block.Index != 1 {
		t.Errorf("Index attendu 1, obtenu %d", block.Index)
	}
}
