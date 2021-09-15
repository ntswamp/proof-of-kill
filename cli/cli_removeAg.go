package cli

import (
	"fmt"
	"os"

	"github.com/ntswamp/proof-of-kill/agent"
)

/// return true if removed
func (cli *Cli) removeAg() {
	a := agent.Load()
	fmt.Println()
	a.Introduce()
	agent.Remove()
	fmt.Println("removed above agent. restart the app To hire a new one.")
	os.Exit(0)

}
