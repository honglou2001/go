package transactions

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"yqx_go/young_blockchain/crypto"
	"crypto/rand"
	"crypto/sha256"
)

const subsidy = 12

// Lock signs the output
func (out *Output) Lock(address []byte) {
	pubKeyHash := crypto.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

//Transaction 一个交易
type Transaction struct {
	ID   []byte  		//交易参照的规则协议版本号
	TxInput   []Input   		//一个或多个输入交易构成的数组
	TxOutput  []Output      //一个或多个输出交易组成的数组
	LockTime uint32        	//锁定时间
}


// NewCoinbaseTX creates a new coinbase transaction
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		randData := make([]byte, 20)
		_, err := rand.Read(randData)
		if err != nil {
			log.Panic(err)
		}

		data = fmt.Sprintf("%x", randData)
	}

	txin := Input{[]byte{}, -1, nil, []byte(data),0xFFFFFFFF}
	txout := NewTXOutput(subsidy, to)
	tx := Transaction{nil, []Input{txin}, []Output{*txout},0}
	tx.ID = tx.Hash()

	return &tx
}

// Hash returns the hash of the Transaction
func (tx *Transaction) Hash() []byte {
	var hash [32]byte

	txCopy := *tx
	txCopy.ID = []byte{}

	hash = sha256.Sum256(txCopy.Serialize())

	return hash[:]
}


//Serialize 返回交易的序列化
func (tx Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}
