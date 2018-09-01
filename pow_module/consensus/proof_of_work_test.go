package consensus

import (
	"fmt"
	"testing"
)
import block "yqx_go/pow_module/blockchain"

func MockProofOfWork() (*ProofOfWork){
	block := block.Block{}
	block.HashPrevBlock =[]byte{1, 2, 3}
	block.HashMerkleRoot=[]byte{2, 3, 4}
	block.TimeStamp = 1000;
	block.Version=1
	block.Nonce = 1
	proofOfWork := NewProofOfWork(&block)
	return proofOfWork
}
func TestProofOfWork_PrepareData(t *testing.T) {
	proofOfWork := MockProofOfWork()
	result :=proofOfWork.PrepareData(0x12)
	if len(result) !=30 {
		t.Error("TestPrepareData is error")
	}
}


func TestProofOfWork_Run(t *testing.T) {
	proofOfWork := MockProofOfWork()
	nonce, hash := proofOfWork.Run()
	fmt.Println(nonce)
	fmt.Println(hash)
	if (nonce !=2 && len(hash) <= 30) {
		t.Error("TestProofOfWork_Run is error")
	}
}
