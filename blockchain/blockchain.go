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
	Data     string
	Hash     string
	PrevHash string
}

/*
The blockChain hash blocks (list of blocks)
*/
type blockChain struct {
	Blocks []*block
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
			b.AddBlock("GenesisBlock")
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
	blockLength := len(GetBlockChain().Blocks)

	if blockLength == 0 {
		return ""
	}

	return GetBlockChain().Blocks[blockLength-1].Hash
}

/*
Create a hash value using sha256

  - Calculate formular: block.hash + block.data
*/
func (b *block) CreateHash() {
	hash := sha256.Sum256([]byte(b.Hash + b.Data))

	b.Hash = fmt.Sprintf("%x", hash)
}

/*
Append a block to the block
*/
func (b *blockChain) AddBlock(data string) {
	b.Blocks = append(b.Blocks, CreateNewBlock(data))
}

/*
Return the all blocks in blockchain
*/
func (b *blockChain) AllBlocks() []*block {
	return b.Blocks
}
