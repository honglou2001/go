package blockchain

import "testing"

func TestGetInstance(t *testing.T) {
	blockchain := GetInstance()
	newBlock := NewBlock(nil, []byte{1}, 1)
	blockchain.AddBlock(newBlock)

	newBlock2 := NewBlock(nil, []byte{1}, 2)
	blockchain.AddBlock(newBlock2)
	len := len(blockchain.ReadBlock())
	if len != 2 {
		t.Error("adding block is error")
	}
}

func TestCreateBlockchain(t *testing.T) {
	blockChain := CreateBlockchain("2017,07,07", "localhost")
	if blockChain.tip == nil {
		t.Error("CreateBlockchain is error")
	}
}
