package cli

import (
	"github.com/ntswamp/proof-of-kill/network"
)

func (cli Cli) startNode() {
	network.StartNode(cli)
}
