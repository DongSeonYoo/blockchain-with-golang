package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

/*
The Block has data, hash, prevHash
*/
type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
}

/*
The blockChain hash blocks (list of blocks)
*/
type blockChain struct {
	Blocks []*Block
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
func CreateNewBlock(data string) *Block {
	newBlock := Block{data, "", GetLastHash(), len(GetBlockChain().AllBlocks()) + 1}
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
func (b *Block) CreateHash() {
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
func (b *blockChain) AllBlocks() []*Block {
	return b.Blocks
}

var ErrBlockNotFound = errors.New("block not found")

/*
Get block details using id (path parameter)
*/
func (b *blockChain) GetBlockById(height int) (*Block, error) {
	if height > len(b.Blocks) {
		return nil, ErrBlockNotFound
	}

	return b.Blocks[height-1], nil
}
