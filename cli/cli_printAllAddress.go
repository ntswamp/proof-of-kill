package cli

import (
	"fmt"

	block "github.com/ntswamp/proof-of-kill/blc"
	"github.com/ntswamp/proof-of-kill/database"

	log "github.com/corgi-kx/logcustom"
)

func (cli *Cli) printAllAddress() {
	bd := database.New()
	addressList := block.GetAllAddress(bd)
	if addressList == nil {
		log.Fatal("No Wallets Found At Current Node.")
	}
	fmt.Println("===================================")
	fmt.Println("Existed Addresses：")
	for _, v := range *addressList {
		fmt.Println(string(v))
	}
	fmt.Println("===================================")
}
