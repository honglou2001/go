package consensus

import (
	"log"
	"testing"
	"yqx_go/young_blockchain/common"
	"yqx_go/young_blockchain/transactions"
	"yqx_go/young_blockchain/wallet"
)

func TestMineBlock(t *testing.T) {
	nodeID := common.NodeID
	fromAddress :=common.GenesisAddress
	toAddress := common.ToAddress
	amount := 1

	bc := NewBlockchain(nodeID)
	UTXOSet := UTXO{bc}
	defer bc.Db.Close()

	wallets, err := wallet.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(fromAddress)

	tx := NewUTXOTransaction(&wallet, toAddress, amount, &UTXOSet)
	cbTx := transactions.NewCoinbaseTX(fromAddress, "")
	txs := []*transactions.Transaction{cbTx, tx}

	newBlock := MineBlock(bc,txs)

	if newBlock == nil || len(newBlock.Hash) <= 30 {
		t.Error("TestMineBlock is error")
	}

}
