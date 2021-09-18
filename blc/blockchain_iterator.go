package block

import "github.com/ntswamp/proof-of-kill/database"

//区块迭代器
type blockchainIterator struct {
	LastBlockHash []byte
	BD            *database.BlockchainDB
}

//获取区块迭代器实例
func NewBlockchainIterator(bc *blockchain) *blockchainIterator {
	blockchainIterator := &blockchainIterator{bc.BD.View([]byte(LATEST_BLOCK_HASH_KEY), database.BlockBucket), bc.BD}
	return blockchainIterator
}

//迭代下一个区块信息
func (bi *blockchainIterator) Next() *Block {
	currentByte := bi.BD.View(bi.LastBlockHash, database.BlockBucket)
	if len(currentByte) == 0 {
		return nil
	}
	block := Block{}
	block.Deserialize(currentByte)
	bi.LastBlockHash = block.PreHash
	return &block
}
