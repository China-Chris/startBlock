package p2p

import (
	"encoding/json"
	"log"
	"net"
	"startblock/blockchain"
)

// SendBlock 向其他节点发送区块
func SendBlock(address string, block blockchain.Block) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Error connecting to node: %v", err)
	}
	defer conn.Close()

	message := map[string]interface{}{
		"type":  "block",
		"block": block,
	}

	encoder := json.NewEncoder(conn)
	if err := encoder.Encode(message); err != nil {
		log.Printf("Error sending block: %v", err)
	}
}

// SendTransaction 向其他节点发送交易
func SendTransaction(address string, tx blockchain.Transaction) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("Error connecting to node: %v", err)
	}
	defer conn.Close()

	message := map[string]interface{}{
		"type":        "transaction",
		"transaction": tx,
	}

	encoder := json.NewEncoder(conn)
	if err := encoder.Encode(message); err != nil {
		log.Printf("Error sending transaction: %v", err)
	}
}
