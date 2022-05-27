package transaction

type Output struct {
	Value   int64
	Address []byte
}

func NewOutput(value int64, address []byte) *Output {
	return &Output{Value: value, Address: address}
}
