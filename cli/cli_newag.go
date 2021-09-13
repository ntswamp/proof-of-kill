package cli

import "github.com/ntswamp/proof-of-kill/agent"

func (cli *Cli) newag() {
	if agent.IsAgentExist() {
		agent.Remove()
	}

}
