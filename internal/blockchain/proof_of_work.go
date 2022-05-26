package blockchain

import (
	"blockchain/internal/utils"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const dificulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-dificulty))
	return &ProofOfWork{block, target}
}

func (p *ProofOfWork) InitData(nonce int) []byte {
	return bytes.Join([][]byte{
		p.Block.PrevHash,
		p.Block.Data,
		utils.ToHex(int64(nonce)),
	},
		[]byte{},
	)
}

func (p *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonse := 0

	for ; nonse < math.MaxInt64; nonse++ {
		data := p.InitData(nonse)
		hash = sha256.Sum256(data)
		fmt.Printf("\n%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(p.Target) == -1 {
			break
		}
		fmt.Println("---------------")
		fmt.Printf("Block created nonse: %d", nonse)
		fmt.Println("---------------")
	}
	return nonse, hash[:]
}
