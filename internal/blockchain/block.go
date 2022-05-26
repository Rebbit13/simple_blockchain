package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func NewBlock(data string, prevHash []byte) *Block {
	block := &Block{Data: []byte(data), PrevHash: prevHash}
	block.DeriveHash()
	return block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}
