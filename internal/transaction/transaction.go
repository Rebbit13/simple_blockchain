package transaction

type Transaction struct {
	ID     []byte
	Inputs []*Input
	Output []*Output
}
