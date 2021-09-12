package cli

import (
	"fmt"

	block "github.com/corgi-kx/blockchain_golang/blc"
	"github.com/corgi-kx/blockchain_golang/database"
)

func (cli *Cli) generateWallet() {
	bd := database.New()
	wallets := block.NewWallets(bd)
	address, privkey, mnemonicWord := wallets.GenerateWallet(bd, block.NewBitcoinKeys, []string{})
	fmt.Println("MNEMONIC WORD：", mnemonicWord)
	fmt.Println("PRIVATE KEY  ：", privkey)
	fmt.Println("ADDRESS      ：", address)
}
