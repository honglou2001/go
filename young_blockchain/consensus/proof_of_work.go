package consensus

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"yqx_go/young_blockchain/blockchain"
	"yqx_go/young_blockchain/common"
)

const targetBits = 0x1        //目标难度，此值根据网络调整
const maxNonce = math.MaxInt64 //最大随机数
//ProofOfWork 共识类，共识的属性及需要的算法
type ProofOfWork struct {
	block  *blockchain.Block //需要计算的Block
	target *big.Int
}

//NewProofOfWork 新建一个工作量证明任务
func NewProofOfWork(b *blockchain.Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}
//PrepareData 进行共识计算前的准备数据
func (pow *ProofOfWork) PrepareData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Hash,
			common.IntToHex(pow.block.Timestamp),
			common.IntToHex(int64(targetBits)),
			common.IntToHex(int64(nonce)),
		},
		[]byte{},
	)
	return data
}

//Run 执行工作量计算
func (pow *ProofOfWork) Run() (int64, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := int64(1)

	fmt.Printf("Mining a new block")
	for nonce < maxNonce {
		data := pow.PrepareData(nonce)

		hash = sha256.Sum256(data)
		if math.Remainder(float64(nonce), 100000) == 0 {
			fmt.Printf("\r%x", hash)
		}
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

//Validate 校验一个block的工作量计算是否有效
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.PrepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}
