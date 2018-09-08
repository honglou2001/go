package transactions

type Output struct {
	Value int   				//输出金额,单位是1聪
	PubKeyHash []byte			//锁定脚本
}


// NewTXOutput create a new TXOutput
func NewTXOutput(value int, address string) *Output {
	txo := &Output{value, nil}
	txo.Lock([]byte(address))
	return txo
}