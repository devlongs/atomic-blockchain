package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PrevBlockHash string
	Hash         string
}

// Blockchain represents a chain of blocks
var Blockchain []Block

// CalculateHash is a function that calculates the SHA256 hash of a block
func (b *Block) CalculateHash() string {
	record := strconv.Itoa(b.Index) + b.Timestamp + b.Data + b.PrevBlockHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

// GenerateBlock creates a new block in the blockchain
func GenerateBlock(oldBlock Block, data string) Block {
	var newBlock Block
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevBlockHash = oldBlock.Hash
	newBlock.Hash = newBlock.CalculateHash()
	return newBlock
}

// IsBlockValid checks if a block is valid by comparing its hash to the calculated hash
func IsBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevBlockHash {
		return false
	}

	if newBlock.CalculateHash() != newBlock.Hash {
		return false
	}

	return true
}

// ReplaceChain replaces the chain with a new chain if the new chain is longer and valid
func ReplaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

func main() {
	fmt.Println("Starting blockchain implementation in Go")
	Blockchain = append(Blockchain, Block{0, time.Now().String(), "Initial Block", "", ""})
	fmt.Println("Initial Blockchain: ", Blockchain)

	// Simulate adding new blocks to the blockchain
	for i := 1; i <= 10; i++ {
		newBlock := GenerateBlock(Blockchain[len(Blockchain)-1], "Block Data "+strconv.Itoa(i))
		if IsBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
			Blockchain = append(Blockchain, newBlock)
		}
	}

	fmt.Println("Final Blockchain: ", Blockchain)
}
