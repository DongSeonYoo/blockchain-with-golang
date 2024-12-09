package main

import (
	"fmt"

	"github.com/DongSeonYoo/go-coin/blockchain"
)

func main() {
	chain := blockchain.GetBlockChain()
	chain.AddBlock("secondBlock")
	chain.AddBlock("thirdBlock")
	chain.AddBlock("fourBlock")

	for _, block := range chain.AllBlocks() {
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
	}
}
