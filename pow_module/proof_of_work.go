package pow_module

import (
	"math/big"
	"bytes"
	"fmt"
	"math"
	"crypto/sha256"
	common "yqx_go/pow_module/common"
)

const targetBits = 24

var (
	maxNonce = math.MaxInt64
)

type ProofOfWork struct {
	Block *Block
	Target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork  {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}
	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte  {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.HashMerkleRoot(),
			common.IntToHex(pow.Block.TimeStamp),
			common.IntToHex(int64(targetBits)),
			common.IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte

	nonce := 0

	fmt.Printf("Minning the block containning %s\n", pow.Block.Data)
	for nonce < maxNonce{
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)

		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.Target) == -1 {
			break
		}else {
			nonce ++
		}
	}

	fmt.Print("\n\n")
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate()bool  {
	var hashInt big.Int

	data := pow.prepareData(pow.Block.Nonce)
	hash := sha256.Sum256(data)

	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.Target) == - 1

	return isValid
}
