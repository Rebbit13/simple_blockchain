package transaction

type Input struct {
	TransactionID []byte
	OutIndex      int64
	Sign          []byte
	PubKey        []byte
}
