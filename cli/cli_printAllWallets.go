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

	fmt.Println("已生成的钱包信息：")
	fmt.Println("==================================================================")
	for k, v := range wallets.Wallets {
		fmt.Println("地址:", k)
		fmt.Printf("公钥:%x\n", v.PublicKey)
		fmt.Println("私钥:", v.GetPrivateKey())
		fmt.Println("助记词:", v.MnemonicWord)
		fmt.Println("==================================================================")
	}
}
