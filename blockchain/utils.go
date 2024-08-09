package blockchain

import (
	"fmt"
)

// PrintBlockchain prints all blocks in the blockchain
func (bc *Blockchain) PrintBlockchain() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("PreviousHash: %s\n", block.PreviousHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("Transactions:")
		for _, tx := range block.Transactions {
			fmt.Printf("  ID: %s, Sender: %s, Receiver: %s, Amount: %d\n", tx.ID, tx.Sender, tx.Receiver, tx.Amount)
		}
		fmt.Println()
	}
}
