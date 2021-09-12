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
		log.Fatal("当前节点没有生成或导入的钱包信息！")
	}
	fmt.Println("===================================")
	fmt.Println("已生成地址：")
	for _, v := range *addressList {
		fmt.Println(string(v))
	}
	fmt.Println("===================================")
}
