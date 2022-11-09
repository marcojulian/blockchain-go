package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	data         map[string]any
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

type BlockChain struct {
	genesisBlock Block
	chain Block[]
	difficulty int
}

func main() {

}

func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String()
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}
