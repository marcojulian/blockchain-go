package main

import "time"

type Block struct {
	data         map[string]any
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

func main() {

}
