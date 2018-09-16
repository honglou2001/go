package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"log"
	"os"
)

const walletFile = "wallet_%s.txt"

//Wallets 所有钱包信息，需要持久化
type Wallets struct {
	Wallets map[string]*Wallet
	nodeID  string
}

// GetWallet returns a Wallet by its address
func (ws Wallets) GetWallet(address string) Wallet {
	return *ws.Wallets[address]
}

//NewWallets 根据节点信息来建立一个钱包
func NewWallets(nodeID string) (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)
	wallets.nodeID = nodeID
	err := wallets.LoadFromFile(nodeID)
	return &wallets, err
}

//CreateWallet a unique wallet and assign to wallets map
func (ws *Wallets) CreateWallet() string {
	wallet := NewWallet()
	address := fmt.Sprintf("%s", wallet.GetAddress())
	ws.Wallets[address] = wallet
	return address
}

//LoadFromFile 加载当前节点的所有钱包
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

//SaveToFile 保存所有钱包到一个文件
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

func (ws *Wallets) ToString() string {
	jsonStr, err := json.Marshal(ws)
	if err != nil {
		log.Panic(err)
	}
	walletsText := fmt.Sprintf("wallets=\n%s\n", string(jsonStr))
	return walletsText
}