package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Block struct {
	Timestamp    int64
	Data         interface{}
	PreviousHash string
	Hash         string
	Nonce        int
}

// CalculateHash calculates the hash of the block
func (b *Block) CalculateHash() string {
	data, _ := json.Marshal(b.Data)
	blockData := fmt.Sprintf("%s%d%s%d", b.PreviousHash, b.Timestamp, data, b.Nonce)
	hash := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hash[:])
}

// Mine mines the block by finding a hash with the required number of leading zeros
func (b *Block) Mine(difficulty int) {
	for !strings.HasPrefix(b.Hash, strings.Repeat("0", difficulty)) {
		b.Nonce++
		b.Hash = b.CalculateHash()
	}
	fmt.Println("Block mined:", b.Hash)
}

// Blockchain struct represents the chain of blocks
type Blockchain struct {
	Chain      []Block
	Difficulty int
}

// NewBlockchain initializes a blockchain with a genesis block
func NewBlockchain() *Blockchain {
	bc := &Blockchain{Difficulty: 2}
	bc.Chain = append(bc.Chain, bc.createGenesisBlock())
	return bc
}

// createGenesisBlock creates the initial block in the blockchain
func (bc *Blockchain) createGenesisBlock() Block {
	return Block{
		Timestamp:    time.Now().Unix(),
		Data:         "Genesis Block",
		PreviousHash: "0",
		Hash:         "",
	}
}

// GetLatestBlock returns the last block in the chain
func (bc *Blockchain) GetLatestBlock() *Block {
	if len(bc.Chain) == 0 {
		return nil
	}
	return &bc.Chain[len(bc.Chain)-1]
}

// AddBlock adds a new block to the chain
func (bc *Blockchain) AddBlock(data interface{}) error {
	latestBlock := bc.GetLatestBlock()
	if latestBlock == nil {
		return fmt.Errorf("Cannot add block to empty blockchain")
	}

	newBlock := Block{
		Timestamp:    time.Now().Unix(),
		Data:         data,
		PreviousHash: latestBlock.Hash,
		Nonce:        0,
	}
	newBlock.Hash = newBlock.CalculateHash()
	newBlock.Mine(bc.Difficulty)
	bc.Chain = append(bc.Chain, newBlock)
	return nil
}

// IsValid checks if the blockchain is valid
func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		previousBlock := bc.Chain[i-1]

		if currentBlock.Hash != currentBlock.CalculateHash() || currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

// Main function to demonstrate the blockchain
func main() {
	myBlockchain := NewBlockchain()
	myBlockchain.AddBlock(map[string]interface{}{"amount": 10})
	myBlockchain.AddBlock(map[string]interface{}{"amount": 20})

	chainData, _ := json.MarshalIndent(myBlockchain, "", "  ")
	fmt.Println(string(chainData))
	fmt.Println("Blockchain valid?", myBlockchain.IsValid())
}
