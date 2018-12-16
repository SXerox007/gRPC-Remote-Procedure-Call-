//simple blockchain
//nothing its simple as Linked list previous Data hash always there with data

//Refrence :- Alex Pliutau

package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Data          string
	PrevBlockHash string
	Hash          string
}

type Blockchain struct {
	Blocks []*Block
}

func (b *Block) setHash() {
	hash := sha256.Sum256([]byte(b.PrevBlockHash + b.Data))
	b.Hash = hex.EncodeToString(hash[:])
}

func NewBlock(data string, prevBlockHash string) *Block {
	block := &Block{data, prevBlockHash, ""}
	block.setHash()

	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", "")
}

func (bc *Blockchain) AddBlock(data string) *Block {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)

	return newBlock
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
