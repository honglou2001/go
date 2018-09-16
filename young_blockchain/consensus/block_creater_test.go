package consensus

import (
	"testing"
	"yqx_go/young_blockchain/common"
)
import BlModule "yqx_go/young_blockchain/blockchain"

func TestGetInstance(t *testing.T) {
	blockchain := BlModule.GetInstance()
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
	blockChain := CreateBlockchain(common.GenesisAddress, common.NodeID)
	defer blockChain.Db.Close()

	UTXOSet := UTXO{blockChain}
	UTXOSet.Reindex()
	if blockChain == nil || len(blockChain.Tip) != 32{
		t.Error("CreateBlockchain  or  DeleteBlockDBFile is error")
	}
}

func TestDeleteBlockDBFile(t *testing.T) {
	blockChain := CreateBlockchain("2017.07.07", "test")
	result := DeleteBlockDBFile("test")
	if blockChain == nil || len(blockChain.Tip) != 32 || result == false {
		t.Error("CreateBlockchain  or  DeleteBlockDBFile is error")
	}
}
