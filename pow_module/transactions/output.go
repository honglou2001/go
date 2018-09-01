package transactions

type Output struct {
	Value int64   				//输出金额,单位是1聪
	PubKeyScript []byte			//锁定脚本
}