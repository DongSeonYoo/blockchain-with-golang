package blockchain

import (
	"sync"
)

type blockchain struct {
	NewstHash string `json:"newstHash"`
	Height    int    `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewstHash, b.Height+1)
	b.NewstHash = block.Hash
	b.Height = block.Height
}

/*
Return the blockChain

  - If blockchain is empty, create a new block in the callback function
*/
func BlockChain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			b.AddBlock("GenesisBlock")
		})
	}

	return b
}
