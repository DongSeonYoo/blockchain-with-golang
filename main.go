package main

import "github.com/DongSeonYoo/go-coin/blockchain"

func main() {
	blockchain.BlockChain().AddBlock("dongseon")
	blockchain.BlockChain().AddBlock("dongseonGod")
	blockchain.BlockChain().AddBlock("dongseonGood")
	blockchain.BlockChain().AddBlock("dongseonGoat")
}
