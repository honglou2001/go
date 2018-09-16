package consensus

import (
	"fmt"
	"testing"
	"yqx_go/young_blockchain/common"
)

func TestReadBlockChain(t *testing.T) {
	//address := common.GenesisAddress
	nodeID :=common.NodeID
/*	blockChainCreate := CreateBlockchain(address, nodeID)
	defer blockChainCreate.Db.Close()*/

	blockChainRead := ReadBlockChain(nodeID)
	fmt.Println(blockChainRead)
	//deleteSuccess := DeleteBlockDBFile(nodeID)
	if len(blockChainRead) != 2 {
		t.Error("TestReadBlockChain is error")
	}
}
