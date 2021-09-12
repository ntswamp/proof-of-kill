package cli

import (
	"fmt"

	block "github.com/ntswamp/proof-of-kill/blc"
)

func (cli *Cli) setRewardAddress(address string) {
	bc := block.NewBlockchain()
	bc.SetRewardAddress(address)
	fmt.Printf("Receiving Mining Reward On Address [%s].\n", address)
}
