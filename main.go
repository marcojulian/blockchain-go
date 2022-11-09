package main

import (
	"time"
)

func main() {

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
