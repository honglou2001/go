package consensus



import (
	"encoding/hex"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	BLModule "yqx_go/young_blockchain/blockchain"
	"yqx_go/young_blockchain/common"
	"yqx_go/young_blockchain/crypto"
	TxModule "yqx_go/young_blockchain/transactions"
	WAModule "yqx_go/young_blockchain/wallet"

)

const utxoBucket = "chainstate"
// UTXOSet represents UTXO set
type UTXO struct {
	BlockChain *BLModule.BlockChain
}


// FindUTXO finds UTXO for a public key hash
func (u UTXO) FindUTXO(pubKeyHash []byte) []TxModule.Output {
	var UTXOs []TxModule.Output
	db := u.BlockChain.Db

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utxoBucket))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			outs := TxModule.DeserializeOutputs(v)

			for _, out := range outs.Outputs {
				if out.IsLockedWithKey(pubKeyHash) {
					UTXOs = append(UTXOs, out)
				}
			}
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return UTXOs
}

// FindSpendableOutputs finds and returns unspent outputs to reference in inputs
func (u UTXO) FindSpendableOutputs(pubkeyHash []byte, amount int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)
	accumulated := 0
	db := u.BlockChain.Db

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(utxoBucket))
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			txID := hex.EncodeToString(k)
			outs := TxModule.DeserializeOutputs(v)

			for outIdx, out := range outs.Outputs {
				if out.IsLockedWithKey(pubkeyHash) && accumulated < amount {
					accumulated += out.Value
					unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)
				}
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return accumulated, unspentOutputs
}

// Reindex rebuilds the UTXO set
func (u UTXO) Reindex() {
	db := u.BlockChain.Db
	bucketName := []byte(utxoBucket)

	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket(bucketName)
		if err != nil && err != bolt.ErrBucketNotFound {
			log.Panic(err)
		}

		_, err = tx.CreateBucket(bucketName)
		if err != nil {
			log.Panic(err)
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	utxos := u.BlockChain.FindUTXO()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)

		for txID, outs := range utxos {
			key, err := hex.DecodeString(txID)
			if err != nil {
				log.Panic(err)
			}

			err = b.Put(key, outs.Serialize())
			if err != nil {
				log.Panic(err)
			}
		}

		return nil
	})
}
// NewUTXOTransaction creates a new transaction
func NewUTXOTransaction(wallet *WAModule.Wallet, to string, amount int, UTXOSet *UTXO) *TxModule.Transaction {
	var inputs []TxModule.Input
	var outputs []TxModule.Output

	pubKeyHash := crypto.HashPubKey(wallet.PublicKey)
	walletsText := fmt.Sprintf("pubKeyHash=\n%s,PublicKey=%s\n",common.ByteToHexString(pubKeyHash),
		common.ByteToHexString(wallet.PublicKey))
	fmt.Println(walletsText)

	acc, validOutputs := UTXOSet.FindSpendableOutputs(pubKeyHash, amount)

	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}

	// Build a list of inputs
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		for _, out := range outs {
			input := TxModule.Input{txID, out, nil, wallet.PublicKey, 0xFFFFFFFF}
			inputs = append(inputs, input)
		}
	}

	// Build a list of outputs
	from := fmt.Sprintf("%s", wallet.GetAddress())
	outputs = append(outputs, *TxModule.NewTXOutput(amount, to))
	if acc > amount {
		outputs = append(outputs, *TxModule.NewTXOutput(acc-amount, from)) // a change
	}

	tx := TxModule.Transaction{nil, inputs, outputs,0}
	tx.ID = tx.Hash()
	UTXOSet.BlockChain.SignTransaction(&tx, wallet.PrivateKey)

	return &tx
}

