module github.com/ntswamp/proof-of-kill

go 1.15

require (
	github.com/boltdb/bolt v1.3.1
	github.com/cloudflare/cfssl v1.4.0
	github.com/corgi-kx/logcustom v0.0.0-20191107084245-589ed3d08a00
	github.com/libp2p/go-libp2p v0.15.0
	github.com/libp2p/go-libp2p-circuit v0.1.4 // indirect
	github.com/libp2p/go-libp2p-core v0.9.0
	github.com/libp2p/go-libp2p-peer v0.2.0
	github.com/libp2p/go-libp2p-peerstore v0.1.4 // indirect
	github.com/libp2p/go-libp2p-secio v0.2.1 // indirect
	github.com/multiformats/go-multiaddr v0.4.0
	github.com/spf13/viper v1.5.0
)

replace (
	github.com/libp2p/go-libp2p => github.com/libp2p/go-libp2p v0.4.0
	github.com/libp2p/go-libp2p-core => github.com/libp2p/go-libp2p-core v0.2.4
)
