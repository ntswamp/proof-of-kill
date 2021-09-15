package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ntswamp/proof-of-kill/agent"
)

/// return true if removed
func (cli *Cli) removeAg() {
	a := agent.Load()
	a.Introduce()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nthis operation will remove above agent permanently. continue?(y/n)")
	fmt.Print("-> ")
	yn, _ := reader.ReadString('\n')
	if yn == "y\n" {
		agent.Remove()
		fmt.Println("agent removed. restart the app To hire a new one.")
		os.Exit(0)
	}
}
