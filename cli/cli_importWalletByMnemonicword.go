package cli

import (
	"encoding/json"
	"fmt"

	block "github.com/ntswamp/proof-of-kill/blc"
	"github.com/ntswamp/proof-of-kill/database"

	log "github.com/corgi-kx/logcustom"
)

func (cli *Cli) importWalletByMnemonicword(mnemonicword string) {
	mnemonicwords := []string{}
	err := json.Unmarshal([]byte(mnemonicword), &mnemonicwords)
	if err != nil {
		log.Error("json err:", err)
	}

	bd := database.New()
	wallets := block.NewWallets(bd)
	address, privkey, mnemonicWord := wallets.GenerateWallet(bd, block.CreateBitcoinKeysByMnemonicWord, mnemonicwords)
	fmt.Println("Mnemonic Word：", mnemonicWord)
	fmt.Println("Private Key  ：", privkey)
	fmt.Println("Address      ：", address)
}
