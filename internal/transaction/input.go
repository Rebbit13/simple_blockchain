package transaction

type Input struct {
	TransactionID []byte
	OutIndex      int64
	Sign          []byte
	PubKey        []byte
}

func NewInput(transactionID []byte, outIndex int64, sign []byte, pubKey []byte) *Input {
	return &Input{TransactionID: transactionID, OutIndex: outIndex, Sign: sign, PubKey: pubKey}
}
