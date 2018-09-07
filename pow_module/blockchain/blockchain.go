package blockchain

import (
	"sync"
)

type BlockChain struct{}

var m *BlockChain
var once sync.Once
var mutex = &sync.Mutex{}
var chains []Block

func GetInstance() *BlockChain {
	once.Do(func() {
		m = &BlockChain{}
	})
	return m
}

func (blockchain *BlockChain) AddBlock(newBlock *Block) {
	val := append(chains, *newBlock)
	ReplaceChain(val)
}

func (blockchain *BlockChain) ReadBlock() []Block {
	return chains
}

func ReplaceChain(newBlocks []Block) {
	mutex.Lock()
	if len(newBlocks) > len(chains) {
		chains = newBlocks
	}
	mutex.Unlock()
}
