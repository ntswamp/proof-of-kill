package cli

import (
	"fmt"

	block "github.com/ntswamp/proof-of-kill/blc"
)

func (cli *Cli) resetUTXODB() {
	bc := block.NewBlockchain()
	utxos := block.UTXOHandle{bc}
	utxos.ResetUTXODataBase()
	fmt.Println("UTXO Database Has Been Reset.")
}
