package blockchain

import (
	"encoding/json"
)

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

func (b *Block) Serialize() ([]byte, error) {
	return json.Marshal(b)
}

func BlockFromJSON(data []byte) (block *Block, err error) {
	err = json.Unmarshal(data, block)
	return
}
