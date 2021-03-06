package cli

import (
	"fmt"

	block "github.com/ntswamp/proof-of-kill/blc"
	"github.com/ntswamp/proof-of-kill/database"
)

func (cli *Cli) generateWallet() {
	bd := database.New()
	wallets := block.NewWallets(bd)
	address, privkey, mnemonicWord := wallets.GenerateWallet(bd, block.NewBitcoinKeys, []string{})
	fmt.Println("Mnemonic Word：", mnemonicWord)
	fmt.Println("Private Key  ：", privkey)
	fmt.Println("Address      ：", address)
}
