package block

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math/rand"
	"time"

	"github.com/ntswamp/proof-of-kill/agent"
	"github.com/ntswamp/proof-of-kill/util"

	log "github.com/corgi-kx/logcustom"
)

type proofOfKill struct {
	*Block
	//target kills need to be reached
	Target uint64
	agent  *agent.Agent
	//current kills
	Kill uint64
}

//return PoK instance
func NewProofOfKill(block *Block) *proofOfKill {
	var target uint64 = util.Uint64Pow(uint64(10), uint64(TargetBits))
	var kill uint64 = 0
	agent := agent.Load()
	pok := &proofOfKill{block, target, agent, kill}
	return pok
}

//进行hash运算,获取到当前区块的hash值
func (p *proofOfKill) run() (int64, []byte, error) {
	var nonce int64 = 0
	var hashByte [32]byte
	log.Info("Start Mining...")

	//initialize random seed
	//genesis block
	seed := time.Now().UnixNano()
	if p.Height != 1 {
		seed = int64(p.generateSeedByHash())
	}
	rand.Seed(seed)

	//开启一个计数器,每隔五秒打印一下当前挖矿,用来直观展现挖矿情况
	times := 0
	ticker1 := time.NewTicker(5 * time.Second)
	go func(t *time.Ticker) {
		for {
			<-t.C
			times += 5
			log.Infof("Mining on Height:%d, had been running for %ds.\nSeed:%d", p.Height, times, seed)
		}
	}(ticker1)

OUTER:
	for {
		for _, tx := range p.Transactions {
			//检测网络上其他节点是否已经挖出了区块
			if p.Height <= NewestBlockHeight {
				//结束计数器
				ticker1.Stop()
				return 0, nil, errors.New("***MINING STOPPED***Received the Latest Block")
			}
			//generate random part of damage
			myRandom := util.RandomInRange(0, p.agent.Luck)
			enemyRandom := util.RandomInRange(0, tx.Agent.Luck)
			if p.isKilledOpponent(&tx.Agent, myRandom, enemyRandom) {
				p.Kill = p.Kill + 1
			}

			if p.Kill >= p.Target {
				break OUTER
			}
		}
	}

	//结束计数器
	ticker1.Stop()
	log.Infof("***AGENT MINED A BLOCK***HEIGHT:%d, SEED:%d, HASH: %x", p.Height, seed, hashByte)
	return nonce, hashByte[:], nil
}

//检验区块是否有效
func (p *proofOfKill) Verify() bool {
	return p.Kill == p.Target
}

//将上一区块hash、数据、时间戳、难度位数、随机数 拼接成字节数组
func (p *proofOfKill) jointData(nonce int64) (data []byte) {
	preHash := p.Block.PreHash
	timeStampByte := util.Int64ToBytes(p.Block.TimeStamp)
	heightByte := util.Int64ToBytes(int64(p.Block.Height))
	nonceByte := util.Int64ToBytes(int64(nonce))
	targetBitsByte := util.Int64ToBytes(int64(TargetBits))
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

func (p *proofOfKill) generateSeedByHash() uint64 {
	var seed uint64 = binary.BigEndian.Uint64(p.Hash)
	return seed
}

//return true if win
func (p *proofOfKill) isKilledOpponent(opponent *agent.Agent, myRandom int, enemyRandom int) bool {

	me := *p.agent
	enemy := *opponent

	//continue until one died
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
