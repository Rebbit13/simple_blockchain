package main

import (
	"blockchain/internal/blockchain"
	"blockchain/internal/transaction"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"log"
	"strconv"
	"time"
)

var n int64 = 0

func randomTransaction() *transaction.Transaction {
	n++
	input := transaction.NewInput([]byte("0"), n, []byte("0"), []byte("0"))
	output := transaction.NewOutput(n*100, []byte("0"))
	return transaction.NewTransaction(
		[]byte(fmt.Sprintf("%d", n)),
		[]*transaction.Input{input},
		[]*transaction.Output{output},
	)
}

func main() {
	privKey, err := btcec.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}
	pubKey := privKey.PubKey().SerializeCompressed()
	pubKeyHash := btcutil.Hash160(pubKey)

	addrPKH, err := btcutil.NewAddressPubKeyHash(pubKeyHash, &chaincfg.Params{})

	addr := addrPKH.EncodeAddress()
	fmt.Println(pubKey)
	fmt.Println(addr)

	bc := blockchain.NewChain()
	go bc.Run()

	for i := 0; i < 30; i++ {
		tr := randomTransaction()
		bc.AddToMemePool(tr)
		log.Print("ADD TRANSACTION")
		time.Sleep(1 * time.Second)
	}
	time.Sleep(20 * time.Second)

	for _, block := range bc.GetBlocks() {
		fmt.Printf("Data: %s\nHash: %x\nPrevious Hash: %x\n\n", block.Data, block.Hash, block.PrevHash)
		proof := blockchain.NewProof(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(proof.Validate()))
	}

}
