package blockchain

type Chain struct {
	blocks []*Block
}

func NewChain() *Chain {
	genesisBlock := NewBlock("Genesis", []byte{})
	return &Chain{[]*Block{genesisBlock}}
}

func (c *Chain) AddBlock(data string) {
	prevBlock := c.blocks[len(c.blocks)-1]
	block := NewBlock(data, prevBlock.Hash)
	c.blocks = append(c.blocks, block)
}

func (c *Chain) GetBlocks() []*Block {
	return c.blocks
}
