package blockchain

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"github.com/boltdb/bolt"
	"log"
	"sync"
	"yqx_go/young_blockchain/common"
	TxModule "yqx_go/young_blockchain/transactions"
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

// AddBlock 增加一个区块
func (bc *BlockChain) AddBlock(block *Block) {
	err := bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(common.BlocksBucket))
		blockInDb := b.Get(block.Hash)

		if blockInDb != nil {
			return nil
		}

		blockData := block.Serialize()
		err := b.Put(block.Hash, blockData)
		if err != nil {
			log.Panic(err)
		}

		lastHash := b.Get([]byte(common.KeyForLasthash))
		lastBlockData := b.Get(lastHash)
		lastBlock := DeserializeBlock(lastBlockData)

		if block.Height > lastBlock.Height {
			err = b.Put([]byte(common.KeyForLasthash), block.Hash)
			if err != nil {
				log.Panic(err)
			}
			bc.Tip = block.Hash
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

// FindUTXO finds all unspent transaction outputs and returns transactions with spent outputs removed
func (bc *BlockChain) FindUTXO() map[string]TxModule.Outputs {
	UTXO := make(map[string]TxModule.Outputs)
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()

	for {
		block := bci.Next()

		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIdx, out := range tx.TxOutput {
				// Was the output spent?
				if spentTXOs[txID] != nil {
					for _, spentOutIdx := range spentTXOs[txID] {
						if spentOutIdx == outIdx {
							continue Outputs
						}
					}
				}

				outs := UTXO[txID]
				outs.Outputs = append(outs.Outputs, out)
				UTXO[txID] = outs
			}

			if tx.IsCoinbase() == false {
				for _, in := range tx.TxInput {
					inTxID := hex.EncodeToString(in.Txid)
					spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return UTXO
}

// Iterator returns a BlockchainIterat
func (bc *BlockChain) Iterator() *ChainIterator {
	bci := &ChainIterator{bc.Tip, bc.Db}

	return bci
}
// FindTransaction finds a transaction by its ID
func (bc *BlockChain) FindTransaction(ID []byte) (TxModule.Transaction, error) {
	bci := bc.Iterator()

	for {
		block := bci.Next()

		for _, tx := range block.Transactions {
			if bytes.Compare(tx.ID, ID) == 0 {
				return *tx, nil
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return TxModule.Transaction{}, errors.New("Transaction is not found")
}

// SignTransaction signs inputs of a Transaction
func (bc *BlockChain) SignTransaction(tx *TxModule.Transaction, privKey ecdsa.PrivateKey) {
	prevTXs := make(map[string]TxModule.Transaction)

	for _, vin := range tx.TxInput {
		prevTX, err := bc.FindTransaction(vin.Txid)
		if err != nil {
			log.Panic(err)
		}
		prevTXs[hex.EncodeToString(prevTX.ID)] = prevTX
	}

	tx.Sign(privKey, prevTXs)
}


// VerifyTransaction verifies transaction input signatures
func (bc *BlockChain) VerifyTransaction(tx *TxModule.Transaction) bool {
	if tx.IsCoinbase() {
		return true
	}

	prevTXs := make(map[string]TxModule.Transaction)

	for _, vin := range tx.TxInput {
		prevTX, err := bc.FindTransaction(vin.Txid)
		if err != nil {
			log.Panic(err)
		}
		prevTXs[hex.EncodeToString(prevTX.ID)] = prevTX
	}

	return tx.Verify(prevTXs)
}

//ReadBlock 读取一个区块
func (bc *BlockChain) ReadBlock() []Block {
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
