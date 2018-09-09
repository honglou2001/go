package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
	TxModule "yqx_go/young_blockchain/transactions"
)

//Block 备份
//type Block111 struct {
//	Version int /*区块版本号，表示本区块遵守的验证规则
//	                             版本、父区块头哈希值和Merkle根采用的是小端格式编码，即低有效位放在前面。 */
//	PrevBlockHash  []byte /*前一区块的哈希值，使用SHA256(SHA256(父区块头))计算*/
//	MerkleRootHash []byte /*该区块中交易的Merkle树根的哈希值，同样采用SHA256(SHA256())计算*/
//	TimeStamp      int64  /*该区块产生的近似时间，精确到秒的UNIX时间戳，
//								   必须严格大于前11个区块时间的中值，
//								   同时全节点也会拒绝那些超出自己2个小时时间戳的区块
//								   时间戳表示的是自1970年1月1日0时0分0秒以来的秒数*/
//	DifficultyTarget int64 /*该区块工作量证明算法的难度目标，已经使用特定算法编码*/
//	Nonce            int64 /*为了找到满足难度目标所设定的随机数，
//								   为了解决32位随机数在算力飞升的情况下不够用的问题，
//								   规定时间戳和coinbase交易信息均可更改，以此扩展nonce的位数*/
//	Height       int
//	Transactions []*TxModule.Transaction
//}
//func NewBlock(version int,hashPrevBlock []byte,
//	hashMerkleRoot []byte,timeStamp int64,
//	difficultyTarget int64,Nonce int64) (block *Block) {
//	block = &Block{}
//	block.Version = version
//	block.PrevBlockHash = hashPrevBlock
//	block.MerkleRootHash = hashMerkleRoot
//	block.TimeStamp = timeStamp
//	block.DifficultyTarget = difficultyTarget
//	block.Nonce = Nonce
//	return
//}
//Block 一个区块
type Block struct {
	/*该区块产生的近似时间，精确到秒的UNIX时间戳，
	  必须严格大于前11个区块时间的中值，
	  同时全节点也会拒绝那些超出自己2个小时时间戳的区块
	  时间戳表示的是自1970年1月1日0时0分0秒以来的秒数*/
	Timestamp    int64
	Transactions []*TxModule.Transaction
	/*前一区块的哈希值，使用SHA256(SHA256(父区块头))计算*/
	PrevBlockHash []byte
	Hash          []byte

	/*为了找到满足难度目标所设定的随机数，
		  为了解决32位随机数在算力飞升的情况下不够用的问题，
	      规定时间戳和coinbase交易信息均可更改，以此扩展nonce的位数*/
	Nonce  int64
	Height int
}

//HashTransactions 获取merkle数根hash值
func (b *Block) HashTransactions() []byte {
	var transactions [][]byte

	for _, tx := range b.Transactions {
		transactions = append(transactions, tx.Serialize())
	}
	mTree := TxModule.NewMerkleTree(transactions)

	return mTree.RootNode.Data
}

// Serialize serializes the block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// DeserializeBlock deserializes a block
func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
