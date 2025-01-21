package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"

	"github.com/DongSeonYoo/go-coin/db"
	"github.com/DongSeonYoo/go-coin/utils"
)

/*
The Block has data, hash, prevHash
*/
type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
	Nonce    int    `json:"nonce"`
}

/*
Converting block to bytes
*/
func (b *Block) toBytes() []byte {
	var blockBuffer bytes.Buffer
	encoder := gob.NewEncoder(&blockBuffer)
	utils.HandleErr(encoder.Encode(b))

	return blockBuffer.Bytes()
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, b.toBytes())
}

/*
Create my block
*/
func createBlock(data string, prevhash string, height int) *Block {
	block := Block{
		Data:     data,
		Hash:     "",
		PrevHash: prevhash,
		Height:   height,
		Nonce:    0,
	}

	payload := block.Data + block.PrevHash + string(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist()

	return &block
}
