package consensus

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
	BlModule "yqx_go/young_blockchain/blockchain"
	"yqx_go/young_blockchain/common"
)

//ReadBlockChain 读区块链信息
func ReadBlockChain(nodeID string) []string {
	dbFile := fmt.Sprintf(common.DbFile, nodeID)
	if !dbExists(dbFile) {
		fmt.Println("Blockchain not exists.")
		os.Exit(1)
	}
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	var data []string
	db.View(func(tx *bolt.Tx) error {
		var vData string
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(common.BlocksBucket))
		b.ForEach(func(k, v []byte) error {
			if string(k) != common.KeyForLasthash {
				block := BlModule.DeserializeBlock(v)
				jsonStr, err := json.Marshal(block)
				if err != nil {
					log.Panic(err)
				}
				vData = string(jsonStr)
			} else {
				vData = hex.EncodeToString(v)
			}
			blockString := fmt.Sprintf("key=%s, value=%s\n", hex.EncodeToString(k), vData)
			data = append(data, blockString)
			return nil
		})
		return nil
	})

	defer db.Close()
	return data
}
