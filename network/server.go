package network

import (
	"context"
	"crypto/rand"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	block "github.com/ntswamp/proof-of-kill/blc"

	log "github.com/corgi-kx/logcustom"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/multiformats/go-multiaddr"
)

//在P2P网络中已发现的节点池
//key:节点ID  value:节点详细信息
var peerPool = make(map[string]peer.AddrInfo)
var ctx = context.Background()
var send = Send{}

//启动本地节点
func StartNode(clier Clier) {
	//get height&kill of the latest local block
	bc := block.NewBlockchain()
	block.NEWEST_BLOCK_HEIGHT = bc.GetLastBlockHeight()
	block.NEWEST_BLOCK_KILL = bc.GetLastBlockKill()
	log.Infof("[*] Listening on IP address: %s Port: %s", ListenHost, ListenPort)
	r := rand.Reader
	// 为本地节点创建RSA密钥对
	prvKey, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, r)
	if err != nil {
		log.Panic(err)
	}
	// 创建本地节点地址信息
	sourceMultiAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%s", ListenHost, ListenPort))
	//传入地址信息，RSA密钥对信息，生成libp2p本地host信息
	host, err := libp2p.New(
		ctx,
		libp2p.ListenAddrs(sourceMultiAddr),
		libp2p.Identity(prvKey),
	)
	if err != nil {
		log.Panic(err)
	}
	//写入全局变量本地主机信息
	localHost = host
	//写入全局变量本地P2P节点地址详细信息
	localAddr = fmt.Sprintf("/ip4/%s/tcp/%s/p2p/%s", ListenHost, ListenPort, host.ID().Pretty())
	log.Infof("[*] Your P2P address: %s", localAddr)
	//启动监听本地端口，并且传入一个处理流的函数，当本地节点接收到流的时候回调处理流的函数
	host.SetStreamHandler(protocol.ID(ProtocolID), handleStream)
	//寻找p2p网络并加入到节点池里
	go findP2PPeer()
	//监测节点池,如果发现网络当中节点有变动则打印到屏幕
	go monitorP2PNodes()
	//启一个go程去向其他p2p节点发送高度信息，来进行更新区块数据
	go sendVersionToPeers()
	//启动程序的命令行输入环境
	go clier.ReceiveCMD()
	fmt.Println("local node established. check out by command: 'tail -f log<your port>.txt'")
	signalHandle()
}

//启动mdns寻找p2p网络 并等节点连接
func findP2PPeer() {
	peerChan := initMDNS(ctx, localHost, RendezvousString)
	for {
		peer := <-peerChan // will block untill we discover a peer
		//将发现的节点加入节点池
		peerPool[fmt.Sprint(peer.ID)] = peer
	}
}

//一个监测程序,监测当前网络中已发现的节点
func monitorP2PNodes() {
	currentPeerPoolNum := 0
	for {
		peerPoolNum := len(peerPool)
		if peerPoolNum != currentPeerPoolNum && peerPoolNum != 0 {
			log.Info(" ----------Network Change Detected. Existing Nodes In Current Node Pool-----------")
			for _, v := range peerPool {
				log.Info("|   ", v, "   |")
			}
			log.Info(" ---------------------------------------------------------------------------------")
			currentPeerPoolNum = peerPoolNum
		} else if peerPoolNum != currentPeerPoolNum && peerPoolNum == 0 {
			log.Info(" ------------Network Change Detected. All Nodes Left The Network------------------")
			currentPeerPoolNum = peerPoolNum
		}
		time.Sleep(time.Second)
	}
}

//向其他p2p节点发送高度信息，来进行更新区块数据
func sendVersionToPeers() {
	//如果节点池中还未存在节点的话,一直循环 直到发现已连接节点
	for {
		if len(peerPool) == 0 {
			time.Sleep(time.Second)
			continue
		} else {
			break
		}
	}
	send.SendVersionToPeers(block.NEWEST_BLOCK_HEIGHT, block.NEWEST_BLOCK_KILL)
}

//节点退出信号处理
func signalHandle() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	send.SendSignOutToPeers()
	fmt.Println("Local Node Terminated.")
	time.Sleep(time.Second)
	os.Exit(0)
}
