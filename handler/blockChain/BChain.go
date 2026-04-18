package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	Data      []byte
	PrevHash  []byte
	Hash      []byte
}

type BlockChain struct {
	Blocks []*Block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]

}

func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlocks := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlocks)
}
func NewBlock(data string, prevBlock []byte) *Block {
	block := &Block{
		Timestamp: time.Now().Unix(),
		Data:      []byte(data),
		PrevHash:  prevBlock,
	}

	block.SetHash()
	return block
}

func NewGennissBlock() *Block {
	return NewBlock("Gennies Block", []byte{})

}

func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGennissBlock()}}
}
