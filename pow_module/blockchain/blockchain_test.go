package blockchain

import "testing"

func TestGetInstance(t *testing.T) {
	blockchain := GetInstance()
	newBlock := NewBlock(1, []byte{1}, []byte{1}, 0, 0, 0)
	blockchain.AddBlock(newBlock)

	newBlock2 := NewBlock(2, []byte{2}, []byte{3}, 2, 0, 2)
	blockchain.AddBlock(newBlock2)
	len := len(blockchain.ReadBlock())
	if len != 2 {
		t.Error("adding block is error")
	}
}
