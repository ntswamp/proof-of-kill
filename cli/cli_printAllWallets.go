package cli

import (
	"fmt"

	block "github.com/ntswamp/proof-of-kill/blc"
	"github.com/ntswamp/proof-of-kill/database"
)

func (cli *Cli) printAllWallets() {
	bd := database.New()
	wallets := block.NewWallets(bd)
	if len(wallets.Wallets) == 0 {
		fmt.Println("No Wallets Found.")
		return
	}

	fmt.Println("My Wallets：")
	fmt.Println("==================================================================")
	for k, v := range wallets.Wallets {
		fmt.Println("Address:", k)
		fmt.Printf("Pub Key: %x\n", v.PublicKey)
		fmt.Println("Priv Key:", v.GetPrivateKey())
		fmt.Println("Mnemonic Word:", v.MnemonicWord)
		fmt.Println("==================================================================")
	}
}
