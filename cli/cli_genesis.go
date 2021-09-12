package cli

import (
	block "github.com/ntswamp/proof-of-kill/blc"
	"github.com/ntswamp/proof-of-kill/network"
)

func (cli *Cli) genesis(address string, value int) {
	bc := block.NewBlockchain()
	bc.CreataGenesisTransaction(address, value, network.Send{})
}
