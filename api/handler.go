package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"startblock/blockchain"
	"strconv"
)

// APIHandler 结构体包含区块链实例
type APIHandler struct {
	Blockchain *blockchain.Blockchain
}

// NewAPIHandler 创建新的 APIHandler
func NewAPIHandler(bc *blockchain.Blockchain) *APIHandler {
	return &APIHandler{Blockchain: bc}
}

// GetTransactionHandler 处理获取交易的请求
func (h *APIHandler) GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	txID := vars["id"]
	transaction := h.Blockchain.FindTransactionByID(txID)
	if transaction != nil {
		fmt.Fprintf(w, "Found transaction: ID: %s, Sender: %s, Receiver: %s, Amount: %d\n", transaction.ID, transaction.Sender, transaction.Receiver, transaction.Amount)
	} else {
		http.Error(w, "Transaction not found", http.StatusNotFound)
	}
}

// GetBlockByIndexHandler 根据区块索引获取区块信息
func (h *APIHandler) GetBlockByIndexHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	index, err := strconv.Atoi(vars["index"])
	if err != nil {
		http.Error(w, "Invalid block index", http.StatusBadRequest)
		return
	}
	block := h.Blockchain.GetBlockByIndex(index)
	if block != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(block)
	} else {
		http.Error(w, "Block not found", http.StatusNotFound)
	}
}

// GetBalanceHandler 查询账户余额
func (h *APIHandler) GetBalanceHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	balance := h.Blockchain.GetBalance(address)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"balance": balance})
}
