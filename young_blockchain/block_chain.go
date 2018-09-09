package young_blockchain

import (
	log "github.com/sirupsen/logrus"
	"time"
	"yqx_go/young_blockchain/transactions"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

// Block represents a block in the blockchain
type Block struct {
	TimeStamp     int64
	Transactions  []*transactions.Transaction
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func NewBlock(transactions []*transactions.Transaction, prewBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions,
		prewBlockHash, []byte{}, 0}

	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}
