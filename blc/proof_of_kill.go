package block

import (
	"encoding/binary"
	"errors"
	"math/rand"
	"time"

	log "github.com/corgi-kx/logcustom"
	"github.com/ntswamp/proof-of-kill/agent"
	"github.com/ntswamp/proof-of-kill/util"
)

type proofOfKill struct {
	*Block
	//target kills need to be reached
	Target       uint64
	verifyAmount uint64
}

//return PoK instance
func NewProofOfKill(block *Block) *proofOfKill {
	var target uint64 = util.Uint64Pow(uint64(2), uint64(TARGET_BIT))
	verifyAmount := util.Uint64Pow(uint64(2), uint64(VERIFY_BIT))
	pok := &proofOfKill{block, target, verifyAmount}
	return pok
}

//进行hash运算,获取到当前区块的hash值
func (p *proofOfKill) run() (int64, []byte, error) {
	var nonce int64 = 0
	var hashByte [32]byte
	log.Info("Start Mining...")

	//generate random seed by the hash of latest block.
	//the case of genesis block
	var seed int64 = GENESIS_SEED
	if p.Height != 1 {
		seed = int64(p.generateSeedByHash(p.PreHash))
	}
	rand.Seed(seed)

	//print mining progress every 5 seconds
	times := 0
	ticker1 := time.NewTicker(5 * time.Second)
	go func(t *time.Ticker) {
		for {
			<-t.C
			times += 5
			log.Infof("Mining on Height:%d, had been running for %ds.\nCurrent kill:%d, Seed:%d.", p.Height, times, p.Kill, seed)
		}
	}(ticker1)

OUTER:
	for {
		for _, tx := range p.Transactions {
			//are other nodes mined a block already?
			if p.Height <= NEWEST_BLOCK_HEIGHT {
				//stop ticker
				ticker1.Stop()
				return 0, nil, errors.New("***MINING STOPPED***Received The Latest Block Already")
			}
			//generate random part of damage
			myRandom := util.RandomInRange(0, p.Agent.Luck)
			enemyRandom := util.RandomInRange(0, tx.Agent.Luck)

			duelResult := p.isKilledOpponent(&tx.Agent, myRandom, enemyRandom)
			if duelResult {
				p.Kill = p.Kill + 1
			}
			//a duel is done.
			p.Attempt = p.Attempt + 1

			if p.Attempt < p.verifyAmount {
				p.Proof = append(p.Proof, duelResult)
			}

			if p.Kill >= p.Target {
				break OUTER
			}
		}
	}

	//结束计数器
	ticker1.Stop()
	log.Infof("***AGENT MINED A BLOCK***HEIGHT: %d, SEED: %d, ATTEMPTS: %d\nHASH: %x", p.Height, seed, p.Attempt, hashByte)
	return nonce, hashByte[:], nil
}

//verify PoK
func (p *proofOfKill) Verify() bool {
	if p.Kill != p.Target {
		log.Infof("Kill target not achieved.")
		return false
	}
	//generate seed locally
	var seed int64 = GENESIS_SEED
	if p.Height != 1 {
		bc := NewBlockchain()
		LocalLatestBlockHash := bc.GetBlockHashByHeight(p.Height - 1)
		if LocalLatestBlockHash == nil {
			log.Infof("Seed(hash) used in incomming block not equals to our latest hash in local chain.")
			return false
		}
		seed = int64(p.generateSeedByHash(LocalLatestBlockHash))
	}
	rand.Seed(seed)

	var i uint64
	for i = 0; i < p.verifyAmount-1; i++ {
		for _, tx := range p.Transactions {

			//generate random part of damage
			minerRandom := util.RandomInRange(0, p.Agent.Luck)
			enemyRandom := util.RandomInRange(0, tx.Agent.Luck)
			duelResult := p.isKilledOpponent(&tx.Agent, minerRandom, enemyRandom)
			if duelResult != p.Proof[i] {
				log.Infof("Block duel result: %v, local result:%v, number of duel:%d", p.Proof[i], duelResult, i)
				return false
			}
		}
	}
	return true
}

/*    PoW
//将上一区块hash、数据、时间戳、难度位数、随机数 拼接成字节数组
func (p *proofOfKill) jointData(nonce int64) (data []byte) {
	preHash := p.Block.PreHash
	timeStampByte := util.Int64ToBytes(p.Block.TimeStamp)
	heightByte := util.Int64ToBytes(int64(p.Block.Height))
	nonceByte := util.Int64ToBytes(int64(nonce))
	targetBitsByte := util.Int64ToBytes(int64(TARGET_BIT))
	//拼接成交易数组
	transData := [][]byte{}
	for _, v := range p.Block.Transactions {
		tBytes := v.getTransBytes() //这里为什么要用到自己写的方法，而不是gob序列化，是因为gob同样的数据序列化后的字节数组有可能不一致，无法用于hash验证
		transData = append(transData, tBytes)
	}
	//获取交易数据的根默克尔节点
	mt := util.NewMerkelTree(transData)

	data = bytes.Join([][]byte{
		preHash,
		timeStampByte,
		heightByte,
		mt.MerkelRootNode.Data,
		nonceByte,
		targetBitsByte},
		[]byte(""))
	return
}
PoW */

func (p *proofOfKill) generateSeedByHash(hash []byte) uint64 {
	var seed uint64 = binary.BigEndian.Uint64(hash)
	return seed
}

//return true if win
func (p *proofOfKill) isKilledOpponent(opponent *agent.Agent, myRandom int, enemyRandom int) bool {
	//DEBUG
	log.Infof("miner random:%d, enemy random:%d", myRandom, enemyRandom)

	me := p.Agent
	enemy := *opponent

	//round continues until one died
	for me.Health > 0 && enemy.Health > 0 {

		//decide first mover
		if rand.Intn(2) == 0 {
			//my turn
			enemy.TakeDamage(me.DealDamage(myRandom))
			if enemy.IsDied() {
				return true
			}
			me.TakeDamage(enemy.DealDamage(enemyRandom))
			if me.IsDied() {
				return false
			}

		} else {
			//enemy's turn
			me.TakeDamage(enemy.DealDamage(enemyRandom))
			if me.IsDied() {
				return false
			}
			enemy.TakeDamage(me.DealDamage(myRandom))
			if enemy.IsDied() {
				return true
			}
		}
	}
	//never reach here
	return false
}
