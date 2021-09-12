package cli

import (
	"fmt"

	block "github.com/ntswamp/proof-of-kill/blc"
)

func (cli *Cli) getBalance(address string) {
	bc := block.NewBlockchain()
	balance := bc.GetBalance(address)
	fmt.Printf("address: %s\nbalance：%d\n", address, balance)
}
