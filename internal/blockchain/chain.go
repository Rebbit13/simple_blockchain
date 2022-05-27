package blockchain

import (
	"blockchain/internal/transaction"
	"encoding/json"
	"log"
)

type Chain struct {
	blocks   []*Block
	MemePool []*transaction.Transaction
}

func NewChain() *Chain {
	genesisBlock := NewBlock([]byte("Genesis"), []byte{})
	return &Chain{[]*Block{genesisBlock}, []*transaction.Transaction{}}
}

func (c *Chain) AddBlock(transactions []*transaction.Transaction) error {
	prevBlock := c.blocks[len(c.blocks)-1]
	data, err := json.Marshal(transactions)
	if err != nil {
		return err
	}
	block := NewBlock(data, prevBlock.Hash)
	c.blocks = append(c.blocks, block)
	return nil
}

func (c *Chain) AddToMemePool(transaction *transaction.Transaction) {
	c.MemePool = append(c.MemePool, transaction)
}

func (c *Chain) GetBlocks() []*Block {
	return c.blocks
}

func (c *Chain) Run() {
	for {
		switch len(c.MemePool) > 3 {
		case true:
			transactions := c.MemePool
			c.MemePool = []*transaction.Transaction{}
			err := c.AddBlock(transactions)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
