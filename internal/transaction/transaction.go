package transaction

type Transaction struct {
	ID     []byte
	Inputs []*Input
	Output []*Output
}

func NewTransaction(ID []byte, inputs []*Input, output []*Output) *Transaction {
	return &Transaction{ID: ID, Inputs: inputs, Output: output}
}
