package transactions

type Transaction struct {
	Version   int32  		//交易参照的规则协议版本号
	TxInput   []Input   		//一个或多个输入交易构成的数组
	TxOutput  []Output      //一个或多个输出交易组成的数组
	LockTime uint32        	//锁定时间
}