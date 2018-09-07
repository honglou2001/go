package transactions

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Transaction struct {
	Version   int32  		//交易参照的规则协议版本号
	TxInput   []Input   		//一个或多个输入交易构成的数组
	TxOutput  []Output      //一个或多个输出交易组成的数组
	LockTime uint32        	//锁定时间
}

//返回交易的序列化
func (tx Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}
