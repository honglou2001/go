package wallet

import "fmt"

const nodeID  = "young180901"

func ClientCreateWallet() {
	wallets, _ := NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile()
	fmt.Printf("Your new address: %s\n", address)
}
