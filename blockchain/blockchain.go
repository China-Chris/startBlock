package blockchain

import (
	"sync"
	"time"
)

// TransactionPool 交易池
type TransactionPool struct {
	Transactions []Transaction
	mutex        sync.Mutex
}

// AddTransaction 添加交易到交易池
func (tp *TransactionPool) AddTransaction(tx Transaction) {
	tp.mutex.Lock()
	defer tp.mutex.Unlock()

	// 可以添加一些验证逻辑，比如交易重复性检查
	tp.Transactions = append(tp.Transactions, tx)
}

// Blockchain 结构体包含区块链的所有区块和总供应量
type Blockchain struct {
	Blocks          []Block
	TotalSupply     int // 总供应量
	TransactionPool TransactionPool
}

// NewBlockchain 创建一个新的区块链，并设置总供应量
func NewBlockchain(totalSupply int) *Blockchain {
	bc := &Blockchain{
		TotalSupply:     totalSupply, // 设置总供应量
		TransactionPool: TransactionPool{},
	}
	bc.CreateGenesisBlock()
	return bc
}

// CreateGenesisBlock 创建创世区块
func (bc *Blockchain) CreateGenesisBlock() {
	genesisBlock := Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Transactions: []Transaction{},
		PreviousHash: "0",
		Hash:         calculateHash(0, time.Now().String(), []Transaction{}, "0"),
	}
	bc.Blocks = append(bc.Blocks, genesisBlock)
}

//// AddBlock 添加新区块到区块链
//func (bc *Blockchain) AddBlock(transactions []Transaction) {
//	index := len(bc.Blocks)
//	timestamp := time.Now().String()
//	var previousHash string
//	if index == 0 {
//		previousHash = "0"
//	} else {
//		previousHash = bc.Blocks[index-1].Hash
//	}
//	block := Block{
//		Index:        index,
//		Timestamp:    timestamp,
//		Transactions: transactions,
//		PreviousHash: previousHash,
//		Hash:         calculateHash(index, timestamp, transactions, previousHash),
//	}
//	bc.Blocks = append(bc.Blocks, block)
//}
// AddBlock 将区块添加到区块链中
func (bc *Blockchain) AddBlock(block Block) {
	// 验证区块的哈希和前一个区块的哈希
	if len(bc.Blocks) > 0 && bc.Blocks[len(bc.Blocks)-1].Hash != block.PreviousHash {
		// 不合法的区块，返回或记录错误
		return
	}

	// 将区块添加到区块链
	bc.Blocks = append(bc.Blocks, block)
}

// GetBlockByIndex 根据区块索引获取区块信息
func (bc *Blockchain) GetBlockByIndex(index int) *Block {
	for _, block := range bc.Blocks {
		if block.Index == index {
			return &block
		}
	}
	return nil
}

// GetBalance 根据地址获取余额
func (bc *Blockchain) GetBalance(address string) int {
	balance := 0
	for _, block := range bc.Blocks {
		for _, tx := range block.Transactions {
			if tx.Receiver == address {
				balance += tx.Amount
			}
			if tx.Sender == address {
				balance -= tx.Amount
			}
		}
	}
	return balance
}

// GetBlockByHash 根据区块哈希获取区块信息
func (bc *Blockchain) GetBlockByHash(hash string) *Block {
	for _, block := range bc.Blocks {
		if block.Hash == hash {
			return &block
		}
	}
	return nil
}

// AddTransaction 将交易添加到交易池中
func (bc *Blockchain) AddTransaction(tx Transaction) {
	bc.TransactionPool.AddTransaction(tx)
}
