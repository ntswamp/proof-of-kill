package network

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	blc "github.com/ntswamp/proof-of-kill/blc"

	log "github.com/corgi-kx/logcustom"
	"github.com/libp2p/go-libp2p-core/network"
)

//对接收到的数据解析出命令,然后对不同的命令分别进行处理
func handleStream(stream network.Stream) {
	data, err := ioutil.ReadAll(stream)
	if err != nil {
		log.Panic(err)
	}
	//取信息的前十二位得到命令
	cmd, content := splitMessage(data)
	log.Tracef("Request received：%s", cmd)
	switch command(cmd) {
	case cVersion:
		go handleVersion(content)
	case cGetHash:
		go handleGetHash(content)
	case cHashMap:
		go handleHashMap(content)
	case cGetBlock:
		go handleGetBlock(content)
	case cBlock:
		go handleBlock(content)
	case cTransaction:
		go handleTransaction(content)
	case cMyError:
		go handleMyError(content)
	}
}

//打印接收到的错误信息
func handleMyError(content []byte) {
	e := myerror{}
	e.deserialize(content)
	log.Warn(e.Error)
	peer := buildPeerInfoByAddr(e.Addrfrom)
	delete(peerPool, fmt.Sprint(peer.ID))
}

//接收交易信息，满足条件后进行挖矿
func handleTransaction(content []byte) {
	t := Transactions{}
	t.Deserialize(content)
	if len(t.Ts) == 0 {
		log.Error("没有满足条件的转账信息，顾不存入交易池")
		return
	}
	//交易池中只能存在每个地址的一笔交易信息
	//判断当前交易池中是否已有该地址发起的交易
	if len(tradePool.Ts) != 0 {
	circle:
		for i := range t.Ts {
			for _, v := range tradePool.Ts {
				if bytes.Equal(t.Ts[i].Vint[0].PublicKey, v.Vint[0].PublicKey) {
					s := fmt.Sprintf("Transaction involved address[%s] already exists in current pool. retry after that transaction get mined.", blc.GetAddressFromPublicKey(t.Ts[i].Vint[0].PublicKey))
					log.Error(s)
					t.Ts = append(t.Ts[:i], t.Ts[i+1:]...)
					goto circle
				}
			}
		}
	}

	if len(t.Ts) == 0 {
		return
	}
	mineBlock(t)
}

//调用区块模块进行挖矿操作
var lock = sync.Mutex{}

func mineBlock(t Transactions) {
	//锁上,等待上一个挖矿结束后才进行挖矿!
	lock.Lock()
	defer lock.Unlock()
	//将临时交易池的交易添加进交易池
	tradePool.Ts = append(tradePool.Ts, t.Ts...)

	for {
		//满足交易池规定的大小后进行挖矿
		if len(tradePool.Ts) >= TradePoolLength {
			log.Debugf("Trade pool is full now:%d, ready to mine.", TradePoolLength)
			mineTrans := Transactions{make([]Transaction, TradePoolLength)}
			copy(mineTrans.Ts, tradePool.Ts[:TradePoolLength])

			bc := blc.NewBlockchain()
			//如果当前节点区块高度小于网络最新高度，则等待节点更新区块后在进行挖矿
			for {
				currentHeight := bc.GetLastBlockHeight()
				if currentHeight >= blc.NEWEST_BLOCK_HEIGHT {
					break
				}
				time.Sleep(time.Second * 1)
			}
			//将network下的transaction转换为blc下的transaction
			nTs := make([]blc.Transaction, len(mineTrans.Ts))
			for i := range mineTrans.Ts {
				nTs[i].TxHash = mineTrans.Ts[i].TxHash
				nTs[i].Vint = mineTrans.Ts[i].Vint
				nTs[i].Vout = mineTrans.Ts[i].Vout
			}
			//进行转帐挖矿
			bc.Transfer(nTs, send)
			//剔除已打包进区块的交易
			newTrans := []Transaction{}
			newTrans = append(newTrans, tradePool.Ts[TradePoolLength:]...)
			tradePool.Ts = newTrans
		} else {
			log.Infof("Transaction collected: %d，Pool size: %d，waiting for more to come", len(tradePool.Ts), TradePoolLength)
			break
		}
	}
}

//接收到区块数据,进行验证后加入数据库
func handleBlock(content []byte) {
	block := &blc.Block{}
	block.Deserialize(content)
	log.Infof("Received a block from another node, Hash of that block：%x", block.Hash)
	bc := blc.NewBlockchain()
	//pow := blc.NewProofOfWork(block)
	pok := blc.NewProofOfKill(block)
	//重新计算本块hash,进行pok验证
	if pok.Verify() {
		log.Infof("PoK verified, Block height：%d", block.Height)
		//如果是创世区块则直接添加进本地库
		currentHash := bc.GetBlockHashByHeight(block.Height)
		if block.Height == 1 && currentHash == nil {
			bc.AddBlock(block)
			utxos := blc.UTXOHandle{bc}
			utxos.ResetUTXODataBase() //重置utxo数据库
			log.Info("Genesis block verified, updated local chain.")
		}
		//验证上一个区块的hash与本块中prehash是否一致
		lastBlockHash := bc.GetBlockHashByHeight(block.Height - 1)
		if lastBlockHash == nil {
			//如果找不到上一个区块,可能是还未同步,建立个循环等待同步
			for {
				time.Sleep(time.Second)
				lastBlockHash = bc.GetBlockHashByHeight(block.Height - 1)
				if lastBlockHash != nil {
					log.Debugf("Block of height %d not found, try to resynchronize...", block.Height-1)
					break
				}
			}
		}
		//如果上一块的hash等于本块prehash则通过存入本地库
		if bytes.Equal(lastBlockHash, block.PreHash) {
			bc.AddBlock(block)
			utxos := blc.UTXOHandle{bc}
			//重置utxo数据库
			utxos.ResetUTXODataBase()
			log.Infof("Prehash verified, block height:%d,", block.Height)
			log.Infof("Received block has passed validation, local chain updated:\nBlock height: %d,\nHash: %x", block.Height, block.Hash)
		} else {
			log.Infof("上一个块高度为%d的hash值为:%x,与本块中的prehash值:%x不一致,固不存入区块链中", block.Height-1, lastBlockHash, block.Hash)
		}
	} else {
		log.Errorf("Failed PoK verification. block[%x] won't be added to local chain\nkill:%d, target:%d", block.Hash, pok.Kill, pok.Target)
	}
}

//接收到获取区块命令,通过hash值 找到该区块 然后把该区块发送过去
func handleGetBlock(content []byte) {
	g := getBlock{}
	g.deserialize(content)
	bc := blc.NewBlockchain()
	blockBytes := bc.GetBlockByHash(g.BlockHash)
	data := jointMessage(cBlock, blockBytes)
	log.Debugf("Sent requested block to %s, Block Hash: %x", g.AddrFrom, g.BlockHash)
	send.SendMessage(buildPeerInfoByAddr(g.AddrFrom), data)
}

//从对面节点处获取到本地区块链所没有的区块hash列表,然后依次发送"获取区块命令"到该节点
func handleHashMap(content []byte) {
	h := hash{}
	h.deserialize(content)
	hm := h.HashMap
	bc := blc.NewBlockchain()
	lastHeight := bc.GetLastBlockHeight()
	targetHeight := lastHeight + 1
	for {
		hash := hm[targetHeight]
		if hash == nil {
			break
		}
		g := getBlock{hash, localAddr}
		data := jointMessage(cGetBlock, g.serialize())
		send.SendMessage(buildPeerInfoByAddr(h.AddrFrom), data)
		log.Debugf("Requsting block... target height：%d", targetHeight)
		targetHeight++
	}
}

//接收到"获取hash列表"命令,返回对面节点所没有的区块的hash信息(两条链的高度差)
func handleGetHash(content []byte) {
	g := getHash{}
	g.deserialize(content)
	bc := blc.NewBlockchain()
	lastHeight := bc.GetLastBlockHeight()
	hm := hashMap{}
	for i := g.Height + 1; i <= lastHeight; i++ {
		hm[i] = bc.GetBlockHashByHeight(i)
	}
	h := hash{hm, localAddr}
	data := jointMessage(cHashMap, h.serialize())
	send.SendMessage(buildPeerInfoByAddr(g.AddrFrom), data)
	log.Debug("Sent requested hashes.")
}

//接收到其他节点的区块高度信息,与本地区块高度进行对比
func handleVersion(content []byte) {
	var lock sync.Mutex
	lock.Lock()
	defer lock.Unlock()
	v := version{}
	v.deserialize(content)
	bc := blc.NewBlockchain()
	if blc.NEWEST_BLOCK_HEIGHT > v.Height {
		log.Info("Other nodes have a smaller height,sending version messages to update them...")
		for {
			currentHeight := bc.GetLastBlockHeight()
			if currentHeight < blc.NEWEST_BLOCK_HEIGHT {
				log.Info("Updating blocks data, about to send the version message.")
				time.Sleep(time.Second)
			} else {
				newV := version{versionInfo, currentHeight, localAddr}
				data := jointMessage(cVersion, newV.serialize())
				send.SendMessage(buildPeerInfoByAddr(v.AddrFrom), data)
				break
			}
		}
	} else if blc.NEWEST_BLOCK_HEIGHT < v.Height {
		log.Debugf("Other nodes have the newer version:%v,requesting hash of new blocks...", v)
		gh := getHash{blc.NEWEST_BLOCK_HEIGHT, localAddr}
		blc.NEWEST_BLOCK_HEIGHT = v.Height
		data := jointMessage(cGetHash, gh.serialize())
		send.SendMessage(buildPeerInfoByAddr(v.AddrFrom), data)
	} else {
		log.Debug("Our node keeps the same height as others, nothing to update.")
	}
}
