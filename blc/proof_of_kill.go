package block

import (
	"bytes"
	"crypto/sha256"
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
	Round        uint64
	verifyAmount uint64
}

//return PoK instance
func NewProofOfKill(block *Block) *proofOfKill {
	var round uint64 = util.Uint64Pow(2, ROUND_BIT)
	verifyAmount := util.Uint64Pow(2, VERIFY_BIT)
	pok := &proofOfKill{block, round, verifyAmount}
	return pok
}

//进行hash运算,获取到当前区块的hash值
func (p *proofOfKill) run() ([]byte, error) {
	var hashByte [32]byte
	log.Info("Start Mining...")

	//generate random seed from the hash of latest block.
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

	var round uint64 = 0
	for round < p.Round {
		for _, tx := range p.Transactions {
			//other nodes mined a block already?
			if p.Height <= NEWEST_BLOCK_HEIGHT {
				//stop ticker
				ticker1.Stop()
				return nil, errors.New("***MINING STOPPED***Received The Latest Block From Another Node")
			}
			//generate random part of damage
			myRandom := util.RandomInRange(0, p.Agent.Luck)
			enemyRandom := util.RandomInRange(0, tx.Agent.Luck)

			duelResult := p.isKilledOpponent(&tx.Agent, myRandom, enemyRandom)
			if duelResult {
				p.Kill = p.Kill + 1
			} else {
				//death punishment
				time.Sleep(time.Millisecond * DEATH_PUNISHMENT)
			}
			if round < p.verifyAmount {
				p.Proof = append(p.Proof, duelResult)
			}
			//a duel is done. round + 1
			round = round + 1
		}
	}
	data := p.jointData(seed)
	hashByte = sha256.Sum256(data)

	//结束计数器
	ticker1.Stop()
	log.Infof("***AGENT MINED A BLOCK***HEIGHT: %d, SEED: %d, KILLS: %d\nHASH: %x", p.Height, seed, p.Kill, hashByte)
	return hashByte[:], nil
}

//verify PoK
func (p *proofOfKill) Verify() bool {
	//generate seed locally
	var seed int64 = GENESIS_SEED
	if p.Height != 1 {
		bc := NewBlockchain()
		LocalLatestBlockHash := bc.GetBlockHashByHeight(p.Height - 1)
		if LocalLatestBlockHash == nil {
			log.Infof("Seed(hash) used in incomming block not found in local chain.")
			return false
		}
		seed = int64(p.generateSeedByHash(LocalLatestBlockHash))
	}
	rand.Seed(seed)
	log.Debugf("seed: %d", seed)
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

//making hash
func (p *proofOfKill) jointData(seed int64) (data []byte) {
	preHash := p.Block.PreHash
	timeStampByte := util.Int64ToBytes(p.Block.TimeStamp)
	heightByte := util.Int64ToBytes(int64(p.Block.Height))
	seedByte := util.Int64ToBytes(int64(seed))
	killByte := util.Uint64ToBytes(p.Kill)
	agentByte := p.Agent.Serliazle()
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
		seedByte,
		killByte,
		agentByte},
		[]byte(""))
	return
}

func (p *proofOfKill) generateSeedByHash(hash []byte) uint64 {
	var seed uint64 = binary.BigEndian.Uint64(hash)
	return seed
}

//return true if win
func (p *proofOfKill) isKilledOpponent(opponent *agent.Agent, myRandom int, enemyRandom int) bool {
	//DEBUG
	//log.Debugf("miner random:%d, enemy random:%d", myRandom, enemyRandom)

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
