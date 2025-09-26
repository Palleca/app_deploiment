package core

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

type Transaction struct {
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Amount   float64 `json:"amount"`
}

type Block struct {
	Index        int           `json:"index"`
	Timestamp    string        `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
	PreviousHash string        `json:"previousHash"`
	Hash         string        `json:"hash"`
}

var (
	Blockchain   []Block
	Transactions []Transaction // mempool
	mutex        sync.Mutex
)

func calculateHash(b Block) string {
	record := fmt.Sprintf("%d|%s|%s", b.Index, b.Timestamp, b.PreviousHash)
	for _, tx := range b.Transactions {
		record += fmt.Sprintf("|%s>%s:%.8f", tx.Sender, tx.Receiver, tx.Amount)
	}
	sum := sha256.Sum256([]byte(record))
	return hex.EncodeToString(sum[:])
}

func createGenesisBlock() Block {
	genesis := Block{
		Index:        0,
		Timestamp:    time.Now().Format(time.RFC3339),
		Transactions: []Transaction{},
		PreviousHash: "0",
	}
	genesis.Hash = calculateHash(genesis)
	return genesis
}

func init() {
	ResetBlockchain()
}

func AddTransaction(tx Transaction) {
	mutex.Lock()
	defer mutex.Unlock()
	Transactions = append(Transactions, tx)
}

func AddBlock(txs []Transaction) Block {
	mutex.Lock()
	defer mutex.Unlock()

	last := Blockchain[len(Blockchain)-1]
	newBlock := Block{
		Index:        len(Blockchain),
		Timestamp:    time.Now().Format(time.RFC3339),
		Transactions: append([]Transaction(nil), txs...),
		PreviousHash: last.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)

	Blockchain = append(Blockchain, newBlock)
	Transactions = []Transaction{}
	return newBlock
}

func MineBlock() Block {
	return AddBlock(Transactions)
}

func ResetBlockchain() {
	mutex.Lock()
	defer mutex.Unlock()
	Blockchain = []Block{createGenesisBlock()}
	Transactions = []Transaction{}
}
