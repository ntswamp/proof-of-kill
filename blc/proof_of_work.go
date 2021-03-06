package block

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"math"
	"math/big"
	"time"

	"github.com/ntswamp/proof-of-kill/util"

	log "github.com/corgi-kx/logcustom"
)

//工作量证明(pow)结构体
type proofOfWork struct {
	*Block
	Target *big.Int
}

//获取POW实例
func NewProofOfWork(block *Block) *proofOfWork {
	target := big.NewInt(1)
	//返回一个大数(1 << 256-TargetBits)
	target.Lsh(target, uint(256-ROUND_BIT))
	pow := &proofOfWork{block, target}
	return pow
}

//进行hash运算,获取到当前区块的hash值
func (p *proofOfWork) run() (int64, []byte, error) {
	var nonce int64 = 0
	var hashByte [32]byte
	var hashInt big.Int
	log.Info("Start Mining...")
	//开启一个计数器,每隔五秒打印一下当前挖矿,用来直观展现挖矿情况
	times := 0
	ticker1 := time.NewTicker(5 * time.Second)
	go func(t *time.Ticker) {
		for {
			<-t.C
			times += 5
			log.Infof("Mining on Height:%d, had been running for %ds.\nNonce:%d, Current Hash:%x", p.Height, times, nonce, hashByte)
		}
	}(ticker1)
	for nonce < MAXINT {
		//检测网络上其他节点是否已经挖出了区块
		if p.Height <= NEWEST_BLOCK_HEIGHT {
			//结束计数器
			ticker1.Stop()
			return 0, nil, errors.New("***MINING STOPPED***Received The Latest Block From Another Node.")
		}
		data := p.jointData(nonce)
		hashByte = sha256.Sum256(data)
		//fmt.Printf("\r current hash : %x", hashByte)
		//convert hash to a big Int
		hashInt.SetBytes(hashByte[:])
		//如果hash后的data值小于设置的挖矿难度大数字,则代表挖矿成功!
		if hashInt.Cmp(p.Target) == -1 {
			break
		} else {
			//nonce++
			bigInt, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
			if err != nil {
				log.Panic("rand:", err)
			}
			nonce = bigInt.Int64()
		}
	}
	//结束计数器
	ticker1.Stop()
	log.Infof("***AGENT MINED A BLOCK***HEIGHT:%d, NONCE:%d, HASH: %x", p.Height, nonce, hashByte)
	return nonce, hashByte[:], nil
}

//检验区块是否有效
func (p *proofOfWork) Verify() bool {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-ROUND_BIT))
	//data := p.jointData(p.Block.Nonce)
	data := p.jointData(9999999999)
	hash := sha256.Sum256(data)
	var hashInt big.Int
	hashInt.SetBytes(hash[:])
	if hashInt.Cmp(target) == -1 {
		return true
	}
	return false
}

//将上一区块hash、数据、时间戳、难度位数、随机数 拼接成字节数组
func (p *proofOfWork) jointData(nonce int64) (data []byte) {
	preHash := p.Block.PreHash
	timeStampByte := util.Int64ToBytes(p.Block.TimeStamp)
	heightByte := util.Int64ToBytes(int64(p.Block.Height))
	nonceByte := util.Int64ToBytes(int64(nonce))
	targetBitsByte := util.Int64ToBytes(int64(ROUND_BIT))
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
