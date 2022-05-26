package blockchain

import (
	"blockchain/internal/utils"
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const dificulty = 24

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-dificulty))
	return &ProofOfWork{block, target}
}

func (p *ProofOfWork) initData(nonce int) []byte {
	return bytes.Join([][]byte{
		p.Block.PrevHash,
		p.Block.Data,
		utils.ToHex(int64(nonce)),
	},
		[]byte{},
	)
}

func (p *ProofOfWork) Validate() bool {
	//TODO: move it to a block
	var intHash big.Int
	data := p.initData(p.Block.Nounse)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])
	return intHash.Cmp(p.Target) == -1
}

func (p *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonse := 0

	for ; nonse < math.MaxInt64; nonse++ {
		data := p.initData(nonse)
		hash = sha256.Sum256(data)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(p.Target) == -1 {
			break
		}
	}
	fmt.Println("\n---------------")
	fmt.Printf("Block created \nHash: %x\nNonse: %d", hash, nonse)
	fmt.Println("\n---------------")
	return nonse, hash[:]
}
