package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"

	"github.com/DongSeonYoo/go-coin/db"
	"github.com/DongSeonYoo/go-coin/utils"
)

type blockchain struct {
	NewstHash string `json:"newstHash"`
	Height    int    `json:"height"`
}

var b *blockchain
var once sync.Once

/*
Add a new Block.
save the block on the database (persist)
*/
func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewstHash, b.Height+1)
	b.NewstHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func (b *blockchain) persist() {
	db.SaveBlockChain(utils.ToBytes(b))
}

/*
Return the blockChain

  - If blockchain is empty, create a new block in the callback function
*/
func BlockChain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			persistedBlockChain := db.CheckPoint()
			if persistedBlockChain == nil {
				b.AddBlock("GenesisBlock")
			} else {
				// resotre b from bytes
				fmt.Println("restored data")
				b.restore(persistedBlockChain)
			}
		})
	}

	fmt.Printf("NewestHash: %s, Height: %d\n", b.NewstHash, b.Height)
	return b
}

func (b *blockchain) restore(data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	decoder.Decode(b)
}
