package wallet

import (
	"github.com/astaxie/beego"
	"testing"
)

func TestCreateWallet(t *testing.T) {
	wallets, err := NewWallets("12")
	address := wallets.CreateWallet() //1241DqYKddenzFeaxGjNabHg6BsPVLoDBPJP
	if len(address) != 36 {
		t.Error("TestCreateWallet is error")
	}
	if err != nil {
		beego.Error("TestCreateWallet fail,", err)
	}
}
