package main

import (
	"blockchain/internal/blockchain"
	"blockchain/internal/keys"
	"blockchain/internal/transaction"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/google/uuid"
	"log"
	"strconv"
	"time"
)

var n int = 0

const prv1Compressed = "a44126a18978ff6f744b9d854eaf42457f1fc634e4d0bd9992f43ad1eb431625"

func randomTransaction(output *transaction.Output, private *btcec.PrivateKey, address []byte) *transaction.Transaction {
	n++
	input := transaction.NewInput([]byte("0"), output.ID, []byte("0"), private.PubKey().SerializeCompressed())
	newOutput := transaction.NewOutput(100, address)
	return transaction.NewTransaction(
		[]byte(uuid.New().String()),
		[]*transaction.Input{input},
		[]*transaction.Output{newOutput},
	)
}

func main() {
	s, err := hex.DecodeString(prv1Compressed)
	if err != nil {
		panic(err)
	}
	prv1, pub1 := btcec.PrivKeyFromBytes(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("FIRST: %x\n", pub1.SerializeCompressed())

	prv2, pub2 := keys.GetOrCreateKeyPair()
	fmt.Printf("SECOND: %x\n", pub2.SerializeCompressed())

	bc, err := blockchain.NewChain()
	if err != nil {
		log.Fatal(err)
	}
	go bc.Run()

	for i := 0; i < 30; i++ {
		time.Sleep(3 * time.Second)
		out := bc.UTXO[0]
		tr := &transaction.Transaction{}
		if i%2 == 0 {
			tr = randomTransaction(out, prv2, pub1.SerializeCompressed())
		} else {
			tr = randomTransaction(out, prv1, pub2.SerializeCompressed())
		}
		err = bc.AddToMemePool(tr)
		if err != nil {
			log.Println(err)
		}
		log.Print("ADD TRANSACTION")
	}
	time.Sleep(10 * time.Second)

	for _, block := range bc.GetBlocks() {
		fmt.Printf("Data: %s\nHash: %x\nPrevious Hash: %x\n\n", block.Data, block.Hash, block.PrevHash)
		proof := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(proof.Validate()))
	}

}
