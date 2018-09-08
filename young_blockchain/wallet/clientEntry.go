package wallet

import (
	"fmt"
	"yqx_go/young_blockchain/common"
)


func ClientCreateWallet() {
	wallets, _ := NewWallets(common.NodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile()
	fmt.Printf("Your new address: %s\n", address)
}
