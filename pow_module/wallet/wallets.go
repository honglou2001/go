package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"os"
)

const walletFile = "wallet_%s.txt"

/*多个钱包*/
type Wallets struct {
	Wallets map[string]*Wallet
	nodeID string
}

func NewWallets(nodeID string) (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)
	wallets.nodeID = nodeID
	err := wallets.LoadFromFile(nodeID)
	return &wallets, err
}

/*create a unique wallet and assign to wallets map*/
func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := fmt.Sprintf("%s", wallet.GetAddress())
	ws.Wallets[address] = wallet
	return address
}

/*加载当前节点的所有钱包*/
func (ws *Wallets) LoadFromFile(nodeID string) error {
	walletFile := fmt.Sprintf(walletFile, nodeID)
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		beego.Error("LoadFromFile ReadFile fail,", err)
	}

	var wallets Wallets
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if err != nil {
		beego.Error("LoadFromFile NewDecoder fail,", err)
	}
	ws.Wallets = wallets.Wallets
	return nil
}

// 保存所有钱包到一个文件
func (ws *Wallets) SaveToFile() {
	var content bytes.Buffer
	walletFile := fmt.Sprintf(walletFile, ws.nodeID)
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(ws)
	if err != nil {
		beego.Error("SaveToFile NewDecoder fail,", err)
	}

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		beego.Error("SaveToFile WriteFile fail,", err)
	}
}
