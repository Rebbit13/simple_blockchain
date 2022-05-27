package blockchain

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonse    int
}

func NewBlock(data []byte, prevHash []byte) *Block {
	block := &Block{Data: data, PrevHash: prevHash}
	proof := NewProof(block)
	block.Nonse, block.Hash = proof.Run()
	return block
}
