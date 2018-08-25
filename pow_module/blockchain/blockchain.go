package blockchain

import (
	"sync"
)

type BlockChain struct {}

var m *BlockChain
var once sync.Once
var mutex = &sync.Mutex{}
var chains []Block

func GetInstance() *BlockChain {
	once.Do(func() {
		m = &BlockChain {}
	})
	return m
}

func  (blockchain *BlockChain) AddBlock(newBlock *Block){
	val :=append(chains, *newBlock)
	replaceChain(val)
}

func replaceChain(newBlocks []Block) {
	mutex.Lock()
	if len(newBlocks) > len(chains) {
		chains = newBlocks
	}
	mutex.Unlock()
}