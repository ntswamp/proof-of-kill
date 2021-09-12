package cli

import (
	"fmt"

	block "github.com/corgi-kx/blockchain_golang/blc"
)

func (cli *Cli) getBalance(address string) {
	bc := block.NewBlockchain()
	balance := bc.GetBalance(address)
	fmt.Printf("address: %s\nbalance：%d\n", address, balance)
}
