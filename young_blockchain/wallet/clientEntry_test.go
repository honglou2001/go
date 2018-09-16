package wallet

import (
	"fmt"
	"testing"

)

func TestClientCreateWallet(t *testing.T) {
	ClientCreateWallet()
	fmt.Printf("test createwallet from client side: %s\n", "nodeID")
}
