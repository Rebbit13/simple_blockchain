package transaction

import "github.com/google/uuid"

type Input struct {
	TransactionID []byte
	OutIndex      uuid.UUID
	Sign          []byte
	PubKey        []byte
}

func NewInput(transactionID []byte, outIndex uuid.UUID, sign []byte, pubKey []byte) *Input {
	return &Input{TransactionID: transactionID, OutIndex: outIndex, Sign: sign, PubKey: pubKey}
}
