package blockchain

import (
	"encoding/json"
	"fmt"
	"os"
)

type Storage struct {
	chainPath  string
	configPath string
}

func NewStorage(chainPath string, configPath string) *Storage {
	return &Storage{chainPath: chainPath, configPath: configPath}
}

func (s *Storage) Get(hash string) (block *Block, err error) {
	f, err := os.Open(fmt.Sprintf("%s/%s", s.chainPath, hash))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	data := []byte{}
	_, err = f.Read(data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, block)
	return block, err
}

func (s *Storage) Create(block *Block) error {
	f, err := os.Open(fmt.Sprintf("%s/%s", s.chainPath, block.Hash))
	if err != nil {
		return err
	}
	defer f.Close()
	data, err := json.Marshal(block)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	return err
}
