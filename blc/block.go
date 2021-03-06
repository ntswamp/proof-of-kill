package block

import (
	"bytes"
	"encoding/gob"
	"math/big"
	"time"

	log "github.com/corgi-kx/logcustom"
	"github.com/ntswamp/proof-of-kill/agent"
)

type Block struct {
	//hash of the provious block
	PreHash      []byte
	Transactions []Transaction
	TimeStamp    int64
	Height       int
	//hash of this block
	Hash []byte
	//for verification
	Agent agent.Agent
	//recording wins or loses for mining fights
	Proof []bool
	//how many kills this block includes
	Kill uint64
}

//生成创世区块
func newGenesisBlock(transaction []Transaction) *Block {
	//创世区块的上一个块hash默认设置成下面的样子
	preHash := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//生成创世区块
	genesisBlock, err := mineBlock(transaction, preHash, 1)
	if err != nil {
		log.Error(err)
	}
	return genesisBlock
}

//PoK
//进行挖矿来生成区块
func mineBlock(transaction []Transaction, preHash []byte, height int) (*Block, error) {
	timeStamp := time.Now().Unix()
	agent := *agent.Load()
	//hash数据+时间戳+上一个区块hash
	block := Block{preHash, transaction, timeStamp, height, nil, agent, nil, 0}
	pok := NewProofOfKill(&block)
	hash, err := pok.run()
	if err != nil {
		return nil, err
	}
	block.Hash = hash[:]
	block.Kill = pok.Kill
	log.Info("PoK verify : ", pok.Verify())
	log.Infof("Made a new block, height: %d", block.Height)
	return &block, nil
}

/**
* PoW
*
//进行挖矿来生成区块
func mineBlock(transaction []Transaction, preHash []byte, height int) (*Block, error) {
	timeStamp := time.Now().Unix()
	//hash数据+时间戳+上一个区块hash
	block := Block{preHash, transaction, timeStamp, height, 0, nil}
	pow := NewProofOfWork(&block)
	nonce, hash, err := pow.run()
	if err != nil {
		return nil, err
	}
	block.Nonce = nonce
	block.Hash = hash[:]
	log.Info("PoK verify : ", pow.Verify())
	log.Infof("Made a new block, height: %d", block.Height)
	return &block, nil
}

*/

// 将Block对象序列化成[]byte
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		panic(err)
	}
	return result.Bytes()
}

func (v *Block) Deserialize(d []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(v)
	if err != nil {
		log.Panic(err)
	}
}

func isGenesisBlock(block *Block) bool {
	var hashInt big.Int
	hashInt.SetBytes(block.PreHash)
	if big.NewInt(0).Cmp(&hashInt) == 0 {
		return true
	}
	return false
}
