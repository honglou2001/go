package transactions

import (
	"bytes"
	"encoding/gob"
	"log"
)

//Output 交易的输出
type Output struct {
	Value      int    //输出金额,单位是1聪
	PubKeyHash []byte //锁定脚本
}

// NewTXOutput create a new TXOutput
func NewTXOutput(value int, address string) *Output {
	txo := &Output{value, nil}
	txo.Lock([]byte(address))
	return txo
}

// Outputs collects Outputs
type Outputs struct {
	Outputs []Output
}


// Serialize serializes Outputs
func (outs Outputs) Serialize() []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(outs)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// IsLockedWithKey checks if the output can be used by the owner of the pubkey
func (out *Output) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// DeserializeOutputs deserializes TXOutputs
func DeserializeOutputs(data []byte) Outputs {
	var outputs Outputs

	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&outputs)
	if err != nil {
		log.Panic(err)
	}

	return outputs
}

