package wallet

import (
	"fmt"
	"testing"
)
func TestCreateWallet(t *testing.T){
	CreateWallet()
	fmt.Printf("test createwallet from client side: %s\n", nodeID)
}