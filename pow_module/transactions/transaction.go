package transactions

type Transaction struct {
	Version int32  			//交易参照的规则协议版本号
	InputCount int 			//输入交易的数量
	Input  []Input   		//一个或多个输入交易构成的数组
	OutputCount int			//输出交易的数量
	Output  []Output        //一个或多个输出交易组成的数组
	TimeStamp int        	//锁定时间
}