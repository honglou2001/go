package blockchain

import (
	"github.com/boltdb/bolt"
	"sync"
)

/*type BlockChain struct{}*/
/*const dbFile = "young_blockchain_%s.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"*/

//BlockChain implements interactions with a DB
type BlockChain struct {
	Tip []byte
	Db  *bolt.DB
}

var m *BlockChain
var once sync.Once
var mutex = &sync.Mutex{}
var chains []Block

//GetInstance 单例实例化
func GetInstance() *BlockChain {
	once.Do(func() {
		m = &BlockChain{}
	})
	return m
}

//AddBlock 增加一个区块
func (blockchain *BlockChain) AddBlock(newBlock *Block) {
	val := append(chains, *newBlock)
	ReplaceChain(val)
}

//ReadBlock 读取一个区块
func (blockchain *BlockChain) ReadBlock() []Block {
	return chains
}

//ReplaceChain 取最长的区块链
func ReplaceChain(newBlocks []Block) {
	mutex.Lock()
	if len(newBlocks) > len(chains) {
		chains = newBlocks
	}
	mutex.Unlock()
}
