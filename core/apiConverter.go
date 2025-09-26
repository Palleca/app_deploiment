package core

import "encoding/json"

func BlockToJSON(b Block) ([]byte, error) {
	return json.MarshalIndent(b, "", "  ")
}

func TransactionToJSON(tx Transaction) ([]byte, error) {
	return json.MarshalIndent(tx, "", "  ")
}

func BlocksToJSON(blocks []Block) ([]byte, error) {
	return json.MarshalIndent(blocks, "", "  ")
}

func TransactionsToJSON(txs []Transaction) ([]byte, error) {
	return json.MarshalIndent(txs, "", "  ")
}
