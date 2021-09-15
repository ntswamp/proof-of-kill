package cli

import (
	"fmt"

	block "github.com/ntswamp/proof-of-kill/blc"
	"github.com/ntswamp/proof-of-kill/network"
)

func (cli Cli) transfer(from, to, amount string) {
	blc := block.NewBlockchain()
	blc.CreateTransaction(from, to, amount, network.Send{})
	fmt.Println("transaction has been broadcast.")
}
