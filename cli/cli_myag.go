package cli

import (
	"fmt"

	"github.com/ntswamp/proof-of-kill/agent"
)

func (cli *Cli) myAg() {
	if agent.IsAgentExist() {
		a := agent.Load()
		a.Introduce()
	} else {
		fmt.Println("agent not found.")
	}
}
