package cli

import (
	"fmt"

	block "github.com/corgi-kx/blockchain_golang/blc"
)

func (cli *Cli) setRewardAddress(address string) {
	bc := block.NewBlockchain()
	bc.SetRewardAddress(address)
	fmt.Printf("Using address [%s] for receiving mining reward.\n", address)
}
