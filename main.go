package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string
}

type blockChain struct {
	blocks []block
}

// hash is one-way function
// "dongsen" + "yoo"  = h_func(x) => "qwasadtwqerkeqwjhr12343jk4111"

// The block number one has a block number one hash.
/**
B1
	b1Hash = data + prevHash
*/
// func main() {
// 	genesisBlock := block{data: "genesis block", hash: "", prevHash: ""}

// 	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.prevHash))

// 	hexHash := fmt.Sprintf("%x", hash)
// 	genesisBlock.hash = hexHash
// 	fmt.Println(genesisBlock)
// }

// If block is a first block.
func (b *blockChain) addBlock(data string) {
	newBlock := block{data: data, hash: "", prevHash: b.getLastHash()}
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
	newBlock.hash = fmt.Sprintf("%x", hash)

	b.blocks = append(b.blocks, newBlock)
}

func (b *blockChain) getLastHash() string {
	// If not newblock.
	if len(b.blocks) > 0 {
		// block should have prev hash.
		return b.blocks[len(b.blocks)-1].hash
	}

	return ""
}

func (b *blockChain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Println(block)
	}
}

func main() {
	chain := blockChain{}

	chain.addBlock("firstBlock")
	chain.addBlock("secondBlock")
	chain.addBlock("thirdBlock")
	chain.addBlock("fourBlock")
	chain.listBlocks()
}
