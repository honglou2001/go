package transactions

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"yqx_go/young_blockchain/crypto"
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
	ID       []byte   //交易参照的规则协议版本号
	TxInput  []Input  //一个或多个输入交易构成的数组
	TxOutput []Output //一个或多个输出交易组成的数组
	LockTime uint32   //锁定时间
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
// IsCoinbase checks whether the transaction is coinbase
func (tx Transaction) IsCoinbase() bool {
	return len(tx.TxInput) == 1 && len(tx.TxInput[0].Txid) == 0 && tx.TxInput[0].Vout == -1
}

// Hash returns the hash of the Transaction
func (tx *Transaction) Hash() []byte {
	var hash [32]byte

	txCopy := *tx
	txCopy.ID = []byte{}

	hash = sha256.Sum256(txCopy.Serialize())

	return hash[:]
}

// Verify verifies signatures of Transaction inputs
func (tx *Transaction) Verify(prevTXs map[string]Transaction) bool {
	if tx.IsCoinbase() {
		return true
	}

	for _, vin := range tx.TxInput {
		if prevTXs[hex.EncodeToString(vin.Txid)].ID == nil {
			log.Panic("ERROR: Previous transaction is not correct")
		}
	}

	txCopy := tx.TrimmedCopy()
	curve := elliptic.P256()

	for inID, vin := range tx.TxInput {
		prevTx := prevTXs[hex.EncodeToString(vin.Txid)]
		txCopy.TxInput[inID].Signature = nil
		txCopy.TxInput[inID].PubKey = prevTx.TxOutput[vin.Vout].PubKeyHash

		r := big.Int{}
		s := big.Int{}
		sigLen := len(vin.Signature)
		r.SetBytes(vin.Signature[:(sigLen / 2)])
		s.SetBytes(vin.Signature[(sigLen / 2):])

		x := big.Int{}
		y := big.Int{}
		keyLen := len(vin.PubKey)
		x.SetBytes(vin.PubKey[:(keyLen / 2)])
		y.SetBytes(vin.PubKey[(keyLen / 2):])

		dataToVerify := fmt.Sprintf("%x\n", txCopy)

		rawPubKey := ecdsa.PublicKey{Curve: curve, X: &x, Y: &y}
		if ecdsa.Verify(&rawPubKey, []byte(dataToVerify), &r, &s) == false {
			return false
		}
		txCopy.TxInput[inID].PubKey = nil
	}

	return true
}


// Sign signs each input of a Transaction
func (tx *Transaction) Sign(privKey ecdsa.PrivateKey, prevTXs map[string]Transaction) {
	if tx.IsCoinbase() {
		return
	}

	for _, vin := range tx.TxInput {
		if prevTXs[hex.EncodeToString(vin.Txid)].ID == nil {
			log.Panic("ERROR: Previous transaction is not correct")
		}
	}

	txCopy := tx.TrimmedCopy()

	for inID, vin := range txCopy.TxInput {
		prevTx := prevTXs[hex.EncodeToString(vin.Txid)]
		txCopy.TxInput[inID].Signature = nil
		txCopy.TxInput[inID].PubKey = prevTx.TxOutput[vin.Vout].PubKeyHash

		dataToSign := fmt.Sprintf("%x\n", txCopy)

		r, s, err := ecdsa.Sign(rand.Reader, &privKey, []byte(dataToSign))
		if err != nil {
			log.Panic(err)
		}
		signature := append(r.Bytes(), s.Bytes()...)

		tx.TxInput[inID].Signature = signature
		txCopy.TxInput[inID].PubKey = nil
	}
}

// TrimmedCopy creates a trimmed copy of Transaction to be used in signing
func (tx *Transaction) TrimmedCopy() Transaction {
	var inputs []Input
	var outputs []Output

	for _, vin := range tx.TxInput {
		inputs = append(inputs, Input{vin.Txid, vin.Vout, nil, nil,0xFFFFFFFF})
	}

	for _, vout := range tx.TxOutput {
		outputs = append(outputs, Output{vout.Value, vout.PubKeyHash})
	}

	txCopy := Transaction{tx.ID, inputs, outputs, 0}

	return txCopy
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
