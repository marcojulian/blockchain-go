package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// create a new blockchain with a mining difficulty of 2
	blockchain := CreateBlockchain(2)

	// record transactions on the blockchain from Alice, Bob, and John
	blockchain.addBlock("Alice", "Bob", 5)
	blockchain.addBlock("John", "Bob", 2)

	// check if the blockchain is valid; expecting true
	fmt.Println("Is valid: " + strconv.FormatBool(blockchain.isValid()))
}

func CreateBlockchain(difficulty int) BlockChain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	return BlockChain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}
