package cli

import block "github.com/ntswamp/proof-of-kill/blc"

func (cli *Cli) printAllBlock() {
	bc := block.NewBlockchain()
	bc.PrintAllBlockInfo()
}
