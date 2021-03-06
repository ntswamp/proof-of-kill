package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/corgi-kx/logcustom"
	"github.com/ntswamp/proof-of-kill/agent"
)

type Cli struct {
}

func printUsage() {
	fmt.Println("----------------------------------------------------------------------------- ")
	fmt.Println("Usage:")
	fmt.Println("\thelp                                           check help message")
	fmt.Println("\tmyag                                           check current agent")
	fmt.Println("\tremoveag                                       remove the current agent permanently")
	fmt.Println("\tgen -a <address> -v <amount>                   make genesis block with money funded")
	fmt.Println("\tsetmineaddr -a <address>                       set the address for mining")
	fmt.Println("\tnewwal                                         make a new wallet")
	fmt.Println("\timportwal -m <mnemonic word>                   import wallets by mnemonics")
	fmt.Println("\tmywal                                          print all local wallets")
	fmt.Println("\tmyaddr                                         print all local addresses")
	fmt.Println("\tbal -a <address>                               check balance")
	fmt.Println("\tsend -from <addr> -to <addr> -amount <amount>  make transfer")
	fmt.Println("\tchain                                          print current chain")
	fmt.Println("\tresetutxo                                      reset UTXO data")
	fmt.Println("------------------------------------------------------------------------------")
}

func New() *Cli {
	return &Cli{}
}

func (cli *Cli) Run() {
	//create an agent if no existing one found
	if !agent.IsAgentExist() {
		cli.newAg()
	}
	printUsage()
	go cli.startNode()
	cli.ReceiveCMD()
}

//获取用户输入
func (cli Cli) ReceiveCMD() {
	stdReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error Reading From Stdin")
			panic(err)
		}
		cli.userCmdHandle(sendData)
	}
}

//parse the cmd
func (cli Cli) userCmdHandle(data string) {
	//trim spaces
	data = strings.TrimSpace(data)
	var cmd string
	var context string
	if strings.Contains(data, " ") {
		cmd = data[:strings.Index(data, " ")]
		context = data[strings.Index(data, " ")+1:]
	} else {
		cmd = data
	}
	switch cmd {
	case "help":
		printUsage()
	case "removeag":
		cli.removeAg()
	case "myag":
		cli.myAg()
	case "gen":
		address := getSpecifiedContent(data, "-a", "-v")
		value := getSpecifiedContent(data, "-v", "")
		v, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		cli.genesis(address, v)
	case "newwal":
		cli.generateWallet()
	case "setmineaddr":
		addrss := getSpecifiedContent(data, "-a", "")
		cli.setRewardAddress(addrss)
	case "importwal":
		mnemonicword := getSpecifiedContent(data, "-m", "")
		cli.importWalletByMnemonicword(mnemonicword)
	case "myaddr":
		cli.printAllAddress()
	case "mywal":
		cli.printAllWallets()
	case "chain":
		cli.printAllBlock()
	case "bal":
		address := getSpecifiedContent(data, "-a", "")
		cli.getBalance(address)
	case "resetutxo":
		cli.resetUTXODB()
	case "send":
		fromString := (context[strings.Index(context, "-from")+len("-from") : strings.Index(context, "-to")])
		toString := strings.TrimSpace(context[strings.Index(context, "-to")+len("-to") : strings.Index(context, "-amount")])
		amountString := strings.TrimSpace(context[strings.Index(context, "-amount")+len("-amount"):])
		cli.transfer(fromString, toString, amountString)
	default:
		fmt.Println("invalid Command.")
		printUsage()
	}
}

//返回data字符串中,标签为tag的内容
func getSpecifiedContent(data, tag, end string) string {
	if end != "" {
		return strings.TrimSpace(data[strings.Index(data, tag)+len(tag) : strings.Index(data, end)])
	}
	return strings.TrimSpace(data[strings.Index(data, tag)+len(tag):])
}
