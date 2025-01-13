package blockchain

import (
	"crypto/sha256"
	"fmt"

	"github.com/DongSeonYoo/go-coin/db"
)

/*
The Block has data, hash, prevHash
*/
type Block struct {
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prevHash,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
}

func createBlock(data string, prevhash string, height int) *Block {
	block := Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevhash,
		Height:   height,
	}

	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist()
	return &block
}

func (b *Block) toBytes() []byte {
	return []byte(b.Data)
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, b.toBytes())
}
