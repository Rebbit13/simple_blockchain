package main

import (
	"blockchain/internal/blockchain"
	"fmt"
)

func main() {
	bc := blockchain.NewChain()
	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Third Block")
	for _, block := range bc.GetBlocks() {
		fmt.Printf("Data: %s\nHash: %x\nPrevious Hash: %x\n\n", block.Data, block.Hash, block.PrevHash)
	}
}
