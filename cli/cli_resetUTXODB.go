package cli

import (
	"fmt"

	block "github.com/corgi-kx/blockchain_golang/blc"
)

func (cli *Cli) resetUTXODB() {
	bc := block.NewBlockchain()
	utxos := block.UTXOHandle{bc}
	utxos.ResetUTXODataBase()
	fmt.Println("UTXO database is reset")
}
