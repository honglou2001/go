package wallet

import (
	"fmt"
	"github.com/astaxie/beego"
	"testing"
	"yqx_go/young_blockchain/common"
)

func TestCreateWallet(t *testing.T) {
	wallets, err := NewWallets("12")
	address := wallets.CreateWallet() //1241DqYKddenzFeaxGjNabHg6BsPVLoDBPJP
	if len(address) != 36 {
		t.Error("TestCreateWallet is error")
	}
	if err != nil {
		beego.Error("TestCreateWallet fail", err)
	}
}

func TestWallets_LoadFromFile(t *testing.T) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)
	wallets.nodeID =common.NodeID
	wallets.LoadFromFile(wallets.nodeID )
	fmt.Println(wallets.ToString())
}
