package consensus

import (
	"github.com/boltdb/bolt"

	"log"
	BLModule "yqx_go/young_blockchain/blockchain"
	"yqx_go/young_blockchain/common"
	TxModule "yqx_go/young_blockchain/transactions"
)

// MineBlock mines a new block with the provided transactions
func MineBlock(bc *BLModule.BlockChain,transactions []*TxModule.Transaction) *BLModule.Block {
	var lastHash []byte
	var lastHeight int

	for _, tx := range transactions {
		// TODO: ignore transaction if it's not valid
		if bc.VerifyTransaction(tx) != true {
			log.Panic("ERROR: Invalid transaction")
		}
	}

	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(common.BlocksBucket))
		lastHash = b.Get([]byte(common.KeyForLasthash))

		blockData := b.Get(lastHash)
		block :=BLModule.DeserializeBlock(blockData)

		lastHeight = block.Height

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(transactions, lastHash, lastHeight+1)

	err = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(common.BlocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}

		err = b.Put([]byte(common.KeyForLasthash), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.Tip = newBlock.Hash

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	return newBlock
}

