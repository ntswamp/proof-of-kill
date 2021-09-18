package block

//用于network包向对等节点发送信息
type Sender interface {
	SendVersionToPeers(height int, kill uint64)
	SendTransToPeers(tss []Transaction)
}
