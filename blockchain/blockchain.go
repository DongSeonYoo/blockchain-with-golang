package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

/*
The block has data, hash, prevHash
*/
type block struct {
	data     string
	hash     string
	prevHash string
}

/*
The blockChain hash blocks (list of blocks)
*/
type blockChain struct {
	blocks []*block
}

var b *blockChain
var once sync.Once

/*
Return the blockChain

  - If blockchain is empty, create a new block in the callback function
*/
func GetBlockChain() *blockChain {
	if b == nil {
		once.Do(func() {
			b = &blockChain{}
			b.blocks = append(b.blocks, CreateNewBlock("dongseonYoo"))
		})
	}

	return b
}

/*
Return the new block
*/
func CreateNewBlock(data string) *block {
	newBlock := block{data, "", GetLastHash()}
	newBlock.CreateHash()

	return &newBlock
}

/*
Return the hash value of the last block in the blockchain

  - If Blockchain is empty, return empty string
*/
func GetLastHash() string {
	blockLength := len(GetBlockChain().blocks)

	if blockLength == 0 {
		return ""
	}

	return GetBlockChain().blocks[blockLength-1].hash
}

/*
Create a hash value using sha256

  - Calculate formular: block.hash + block.data
*/
func (b *block) CreateHash() {
	hash := sha256.Sum256([]byte(b.hash + b.data))

	b.hash = fmt.Sprintf("%x", hash)
}
