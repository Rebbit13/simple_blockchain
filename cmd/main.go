package main

import (
	"blockchain/internal/blockchain"
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	bc := blockchain.NewChain()
	bc.AddBlock("First Block")
	bc.AddBlock("Second Block")
	bc.AddBlock("Third Block")
	elapsed := time.Since(start)
	log.Printf("Running %s", elapsed)

	for _, block := range bc.GetBlocks() {
		fmt.Printf("Data: %s\nHash: %x\nPrevious Hash: %x\n\n", block.Data, block.Hash, block.PrevHash)
		//bl, _ := json.Marshal(block)
		//bld := &blockchain.Block{}
		//_ = json.Unmarshal(bl, bld)
		//fmt.Printf("%x\n", bld.Hash)
		proof := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(proof.Validate()))
	}

}
