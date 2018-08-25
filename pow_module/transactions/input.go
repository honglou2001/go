package transactions

type Input struct {
	PreTxHash  int32  			//前置交易hash
	PreTxIndex  int   			//前置交易的索引
	UnlockScriptLength int  	//解锁脚本长度
	UnlockScrip  int    		//解锁脚本
	Sequence int     			//序列
}