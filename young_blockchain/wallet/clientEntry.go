package wallet

import (
	"fmt"
	"yqx_go/young_blockchain/common"
)

//ClientCreateWallet 创建钱包，可通过API的方式提供给调用
func ClientCreateWallet() {
	wallets, _ := NewWallets(common.NodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile()
	fmt.Printf("Your new address: %s\n", address)
}
