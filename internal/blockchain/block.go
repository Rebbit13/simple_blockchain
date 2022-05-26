package blockchain

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nounse   int
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{Data: []byte(data), PrevHash: prevHash}
	proof := NewProof(block)
	block.Nounse, block.Hash = proof.Run()
	return block
}
