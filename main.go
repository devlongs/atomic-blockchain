package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Define the structure of a block
type Block struct {
    Index int
    Timestamp string
    BPM int
    Hash string
    PrevHash string
}

// Blockchain is a slice of blocks
var Blockchain []Block

// Calculate the hash of a block
func calculateHash(block Block) string {
    record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
    h := sha256.New()
    h.Write([]byte(record))
    hashed := h.Sum(nil)
    return hex.EncodeToString(hashed)
}

// Generate a new block
func generateBlock(oldBlock Block, BPM int) (Block, error) {
    var newBlock Block

    t := time.Now()

    newBlock.Index = oldBlock.Index + 1
    newBlock.Timestamp = t.String()
    newBlock.BPM = BPM
    newBlock.PrevHash = oldBlock.Hash
    newBlock.Hash = calculateHash(newBlock)

    return newBlock, nil
}

// Validate a block
func isBlockValid(newBlock, oldBlock Block) bool {
    if oldBlock.Index+1 != newBlock.Index {
        return false
    }

    if oldBlock.Hash != newBlock.PrevHash {
        return false
    }

    if calculateHash(newBlock) != newBlock.Hash {
        return false
    }

    return true
}

// Handle block chain updates
func replaceChain(newBlocks []Block) {
    if len(newBlocks) > len(Blockchain) {
        Blockchain = newBlocks
    }
}

func main() {
    genesisBlock := Block{0, time.Now().String(), 0, "", ""}
    Blockchain = []Block{genesisBlock}
}