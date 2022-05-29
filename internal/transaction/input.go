package transaction

import (
	"blockchain/internal/utils"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	ecdsa2 "github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/google/uuid"
)

type Input struct {
	TransactionID []byte
	OutIndex      uuid.UUID
	Sign          []byte
	PubKey        []byte
}

func NewInput(transactionID []byte, outIndex uuid.UUID, privateKey *btcec.PrivateKey) *Input {
	input := &Input{
		TransactionID: transactionID,
		OutIndex:      outIndex,
		Sign:          []byte{},
		PubKey:        privateKey.PubKey().SerializeCompressed(),
	}
	input.signByPrivateKey(privateKey)
	return input
}

func (i *Input) getDoubleHashedData() [32]byte {
	data := []byte(fmt.Sprintf("%s||%s||%x", i.TransactionID, i.OutIndex.String(), i.PubKey))
	return utils.DoubleHash(data)
}

func (i *Input) signByPrivateKey(privateKey *btcec.PrivateKey) {
	dHash := i.getDoubleHashedData()
	i.Sign = ecdsa2.Sign(privateKey, dHash[:]).Serialize()
}

func (i *Input) VerifySignature() bool {
	dHash := i.getDoubleHashedData()
	signature, err := ecdsa2.ParseSignature(i.Sign)
	if err != nil {
		return false
	}
	pubKey, err := btcec.ParsePubKey(i.PubKey)
	if err != nil {
		return false
	}
	return signature.Verify(dHash[:], pubKey)
}
