package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Block represents a single block in the blockchain
type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction // List of transactions in the block
	PreviousHash string
	Hash         string
}

// calculateHash 计算区块哈希值
func calculateHash(index int, timestamp string, transactions []Transaction, previousHash string) string {
	record := strconv.Itoa(index) + timestamp + previousHash
	for _, tx := range transactions {
		record += tx.ID + tx.Sender + tx.Receiver + strconv.Itoa(tx.Amount)
	}
	hash := sha256.New()
	hash.Write([]byte(record))
	return hex.EncodeToString(hash.Sum(nil))
}

// NewBlock creates a new block with transactions
func NewBlock(index int, transactions []Transaction, previousHash string) *Block {
	block := &Block{
		Index:        index,
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		PreviousHash: previousHash,
	}
	block.Hash = calculateHash(0, time.Now().String(), []Transaction{}, "0")
	return block
}
