package blockchain

import "testing"

func TestGetInstance(t *testing.T) {
	blockchain := GetInstance()
	newBlock := NewBlock(1, 9, 0, 0, 0, 0)
	blockchain.AddBlock(newBlock)

	newBlock2 := NewBlock(2, 7, 1, 2, 0, 2)
	blockchain.AddBlock(newBlock2)
	len := len(blockchain.ReadBlock())
	if len != 2 {
		t.Error("adding block is error")
	}
}
