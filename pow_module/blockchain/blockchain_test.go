package blockchain

import "testing"

func TestGetInstance(t *testing.T)  {
	blockchain := GetInstance()
	genesisBlock := NewBlock(0, 9, 0,0, 0,0)
	blockchain.AddBlock(genesisBlock)
}