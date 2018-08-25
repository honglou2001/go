package transactions

import "math/big"

type Output struct {
	Value big.Float   			//输出金额,单位是1聪
	LockScriptLength  int       //锁定脚本长度
	LockScript  int          	//锁定脚本
}