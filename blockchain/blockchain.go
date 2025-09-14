package blockchain

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*Block
}

var b *blockchain

func getLastHash() string {
	blocks := GetBlockchain().blocks
	totalLen := len(blocks)
	if totalLen == 0 {
		return ""
	}
	return blocks[totalLen-1].Hash
}

func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash()}
	newBlock.calculateHash()

	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	if b == nil {
		b = &blockchain{}
		b.AddBlock("Genesis Block")
	}
	return b
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}
