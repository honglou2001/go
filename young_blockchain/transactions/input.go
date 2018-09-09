package transactions

/*type Input struct {
	TxHash  []byte  			// Tx交易的Hash值，固定为32字节
	PreviousOutTxIndex  int   	//前置交易的索引
	Signature 	 []byte   //私钥签名
	PublicKey    []byte   //公钥数据
	Sequence uint32    //序列号，UINT32, 固定4字节,默认都设成0xFFFFFFFF
}*/

//Input represents a transaction input
type Input struct {
	Txid      []byte
	Vout      int
	Signature []byte
	PubKey    []byte
	Sequence  uint32 //序列号，UINT32, 固定4字节,默认都设成0xFFFFFFFF
}
