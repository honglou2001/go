package consensus

import (
	"fmt"
	"testing"
)

func TestReadBlockChain(t *testing.T) {
	address := "2017.07.07"
	nodeId := "test"
	blockChain_Create := CreateBlockchain(address, nodeId)
	blockChain_Read := ReadBlockChain(nodeId)
	fmt.Println(blockChain_Read)
	delete_success := DeleteBlockDBFile(nodeId)
	if (blockChain_Create == nil || len(blockChain_Create.Tip) != 32 || len(blockChain_Read) != 2 || delete_success == false) {
		t.Error("CreateBlockchain  or  DeleteBlockDBFile is error")
	}
}
