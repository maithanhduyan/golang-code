package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// Block đại diện cho một khối trong blockchain
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// Blockchain đại diện cho chuỗi khối
type Blockchain struct {
	Blocks []*Block
}

// Tạo một khối mới
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
	}
	block.SetHash()
	return block
}

// Tính toán và thiết lập giá trị hash cho khối
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

// Thêm một khối mới vào chuỗi khối
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// In thông tin của các khối trong chuỗi khối
func (bc *Blockchain) PrintBlocks() {
	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}

func main() {
	// Khởi tạo blockchain và thêm các khối vào
	genesisBlock := NewBlock("Genesis Block", []byte{})
	blockchain := Blockchain{Blocks: []*Block{genesisBlock}}
	blockchain.AddBlock("Data of Block 1")
	blockchain.AddBlock("Data of Block 2")

	// In thông tin của các khối trong blockchain
	blockchain.PrintBlocks()
}
