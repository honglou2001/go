package wallet

import "fmt"

const nodeID  = "young180901"

func CreateWallet() {
	wallets, _ := NewWallets(nodeID)
	address := wallets.CreateWallet()
	wallets.SaveToFile()
	fmt.Printf("Your new address: %s\n", address)
}
