package transaction

import "github.com/google/uuid"

type Output struct {
	ID      uuid.UUID
	Value   int64
	Address []byte
}

func NewOutput(value int64, address []byte) *Output {
	return &Output{ID: uuid.New(), Value: value, Address: address}
}
