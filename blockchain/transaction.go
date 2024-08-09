package blockchain

import "errors"

// Transaction represents a transaction in the blockchain
type Transaction struct {
	ID       string // Unique transaction identifier
	Sender   string // Address of the sender
	Receiver string // Address of the receiver
	Amount   int    // Amount transferred
}

// NewTransaction creates a new transaction
func NewTransaction(id, sender, receiver string, amount int) *Transaction {
	return &Transaction{
		ID:       id,
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
	}
}

// FindTransactionByID 根据交易ID查找交易
func (bc *Blockchain) FindTransactionByID(txID string) *Transaction {
	for _, block := range bc.Blocks {
		for _, tx := range block.Transactions {
			if tx.ID == txID {
				return &tx
			}
		}
	}
	return nil
}

// FromMap 将区块数据转换为 Block 结构体
func FromMap(data map[string]interface{}) (Block, error) {
	block := Block{}

	// 解析数据
	if index, ok := data["index"].(float64); ok {
		block.Index = int(index)
	} else {
		return block, errors.New("invalid index")
	}
	if timestamp, ok := data["timestamp"].(string); ok {
		block.Timestamp = timestamp
	} else {
		return block, errors.New("invalid timestamp")
	}
	if transactions, ok := data["transactions"].([]interface{}); ok {
		for _, txData := range transactions {
			txMap, ok := txData.(map[string]interface{})
			if !ok {
				return block, errors.New("invalid transaction data")
			}
			tx, err := FromTransactionMap(txMap)
			if err != nil {
				return block, err
			}
			block.Transactions = append(block.Transactions, tx)
		}
	} else {
		return block, errors.New("invalid transactions")
	}
	if previousHash, ok := data["previousHash"].(string); ok {
		block.PreviousHash = previousHash
	} else {
		return block, errors.New("invalid previousHash")
	}
	if hash, ok := data["hash"].(string); ok {
		block.Hash = hash
	} else {
		return block, errors.New("invalid hash")
	}

	return block, nil
}

// FromTransactionMap 将交易数据转换为 Transaction 结构体
func FromTransactionMap(data map[string]interface{}) (Transaction, error) {
	tx := Transaction{}

	// 解析数据
	if id, ok := data["id"].(string); ok {
		tx.ID = id
	} else {
		return tx, errors.New("invalid id")
	}
	if sender, ok := data["sender"].(string); ok {
		tx.Sender = sender
	} else {
		return tx, errors.New("invalid sender")
	}
	if receiver, ok := data["receiver"].(string); ok {
		tx.Receiver = receiver
	} else {
		return tx, errors.New("invalid receiver")
	}
	if amount, ok := data["amount"].(float64); ok {
		tx.Amount = int(amount)
	} else {
		return tx, errors.New("invalid amount")
	}

	return tx, nil
}
