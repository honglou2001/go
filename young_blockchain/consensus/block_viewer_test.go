package consensus

import (
	"fmt"
	"testing"
)

func TestReadBlockChain(t *testing.T) {
	address := "2017.07.07"
	nodeID := "test"
	blockChainCreate := CreateBlockchain(address, nodeID)
	blockChainRead := ReadBlockChain(nodeID)
	fmt.Println(blockChainRead)
	deleteSuccess := DeleteBlockDBFile(nodeID)
	if (blockChainCreate == nil || len(blockChainCreate.Tip) != 32 || len(blockChainRead) != 2 || deleteSuccess == false) {
		t.Error("CreateBlockchain  or  DeleteBlockDBFile is error")
	}
}
