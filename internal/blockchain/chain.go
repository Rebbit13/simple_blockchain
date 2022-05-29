package blockchain

import (
	"blockchain/internal/transaction"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/google/uuid"
	"log"
)

type Chain struct {
	blocks   []*Block
	MemePool []*transaction.Transaction
	// UTXO: unspent transaction output
	UTXO []*transaction.Output
}

func firstTransaction() *transaction.Transaction {
	data, err := hex.DecodeString("7f7a75d22e5fcac3200cd0a275912964b3ce5423ea5300a86cf8e0b3c097c9f4")
	if err != nil {
		log.Fatal(err)
	}
	_, public := btcec.PrivKeyFromBytes(data)
	inputs := []*transaction.Input{&transaction.Input{
		TransactionID: []byte{},
		OutIndex:      uuid.UUID{},
		Sign:          []byte{},
		PubKey:        []byte{},
	}}
	outputs := []*transaction.Output{&transaction.Output{
		Value:   100,
		Address: public.SerializeCompressed(),
	}}
	return transaction.NewTransaction([]byte(uuid.New().String()), inputs, outputs)
}

func NewChain() (*Chain, error) {
	firstTransactions := []*transaction.Transaction{firstTransaction()}
	data, err := json.Marshal(firstTransactions)
	if err != nil {
		return nil, err
	}
	genesisBlock := NewBlock(data, []byte{})
	firstOutputs := []*transaction.Output{}
	for _, trx := range firstTransactions {
		firstOutputs = append(firstOutputs, trx.Output...)
	}
	return &Chain{
		[]*Block{genesisBlock},
		[]*transaction.Transaction{},
		firstOutputs,
	}, nil
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

func (c *Chain) findUTXOByInput(input *transaction.Input) *transaction.Output {
	for _, UTXO := range c.UTXO {
		if input.OutIndex == UTXO.ID {
			return UTXO
		}
	}
	return nil
}

func (c *Chain) deleteUTXO(input *transaction.Input) {
	i := -1
	for j := 0; j < len(c.UTXO); j++ {
		candidate := c.UTXO[j]
		if input.OutIndex == candidate.ID {
			i = j
			break
		}
	}
	if i != -1 {
		c.UTXO = append(c.UTXO[:i], c.UTXO[i+1:]...)
	}
}

func (c *Chain) verifyTransaction(transaction *transaction.Transaction) bool {
	for _, input := range transaction.Inputs {
		UTXO := c.findUTXOByInput(input)
		if UTXO == nil {
			return false
		}
	}
	return true
}

func (c *Chain) AddToMemePool(transaction *transaction.Transaction) error {
	if !c.verifyTransaction(transaction) {
		return errors.New("transaction is not valid")
	}
	for _, input := range transaction.Inputs {
		c.deleteUTXO(input)
	}
	c.MemePool = append(c.MemePool, transaction)
	return nil
}

func (c *Chain) GetBlocks() []*Block {
	return c.blocks
}

func (c *Chain) Run() {
	for {
		switch len(c.MemePool) > 0 {
		case true:
			transactions := c.MemePool
			c.MemePool = []*transaction.Transaction{}
			err := c.AddBlock(transactions)
			for _, trx := range transactions {
				c.UTXO = append(c.UTXO, trx.Output...)
			}
			if err != nil {
				log.Fatal(err)
			}
			message := "UTXO LEFT: "
			for _, UTXO := range c.UTXO {
				message += fmt.Sprintf(
					"\nID: %s | Value: %d | Vallet: %x",
					UTXO.ID.String(), UTXO.Value, UTXO.Address,
				)
			}
			log.Println(message)
		}
	}
}
