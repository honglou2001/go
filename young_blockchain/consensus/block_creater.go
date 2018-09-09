package consensus

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
	"time"
	BlModule "yqx_go/young_blockchain/blockchain"
	TxModule "yqx_go/young_blockchain/transactions"
)

/*type BlockChain struct{}*/
const dbFile = "young_blockchain_%s.db"
const blocksBucket = "blocks"
const keyForLasthash = "1"
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

// CreateBlockchain creates a new blockchain DB
func CreateBlockchain(address, nodeID string) *BlModule.BlockChain {
	dbFile := fmt.Sprintf(dbFile, nodeID)
	if dbExists(dbFile) {
		fmt.Println("Blockchain already exists.")
		os.Exit(1)
	}

	var tip []byte

	cbtx := TxModule.NewCoinbaseTX(address, genesisCoinbaseData)
	genesis := NewGenesisBlock(cbtx)

	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucket([]byte(blocksBucket))
		if err != nil {
			log.Panic(err)
		}

		err = b.Put(genesis.Hash, genesis.Serialize())
		if err != nil {
			log.Panic(err)
		}
		//用特定key，保留最后的Hash
		err = b.Put([]byte(keyForLasthash), genesis.Hash)
		if err != nil {
			log.Panic(err)
		}
		tip = genesis.Hash

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	bc := BlModule.BlockChain{tip, db}

	defer db.Close()

	return &bc
}
// NewBlock 新产生一个区块，需要经过共识
func NewBlock(transactions []*TxModule.Transaction, prevBlockHash []byte, height int) *BlModule.Block {
	block := &BlModule.Block{time.Now().Unix(), transactions, prevBlockHash,
	[]byte{}, 0, height}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
//NewGenesisBlock 生成创世区块
func NewGenesisBlock(coinbase *TxModule.Transaction) *BlModule.Block {
	return NewBlock([]*TxModule.Transaction{coinbase}, []byte{}, 0)
}

//dbExists 区块链数据库是否存在
func dbExists(dbFile string) bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}
	return true
}
//DeleteBlockDBFile 删除区块链数据库文件，用于测试用
func DeleteBlockDBFile(nodeID string) bool{
	dbFile := fmt.Sprintf(dbFile, nodeID)
	errRemove := os.Remove(dbFile)
	if errRemove != nil {
		//如果删除失败则输出 file remove Error!
		fmt.Println("file remove Error!")
		//输出错误详细信息
		fmt.Printf("%s", errRemove)
		return false
	}
	//如果删除成功则输出 file remove OK!
	fmt.Print("file remove OK!")
	return true
}
