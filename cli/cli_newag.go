package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ntswamp/proof-of-kill/agent"
)

func (cli *Cli) newag() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("This Operation Will Remove Current Agent. Continue?(y/n)")
	fmt.Print("-> ")
	yn, _ := reader.ReadString('\n')
	if yn == "n" {
		return
	}
	if agent.IsAgentExist() {
		agent.Remove()
	}

	for {
		var name string
		var class int
		var weapon int

		fmt.Println("Please Name Your Agent: ")
		fmt.Print("-> ")
		name, _ = reader.ReadString('\n')
		// convert CRLF to LF
		name = strings.Replace(name, "\n", "", -1)

		fmt.Printf("%s? Such An Impressive Name.\nWelcome To The World Of PoK, %s.\n", name, name)

		fmt.Println("Now Tell Me Your Class, By The Number:")
		fmt.Println("#1 Warrior")
		fmt.Println("#2 Mage")
		fmt.Println("#3 Archer")
		fmt.Print("-> ")
		_, err := fmt.Scanf("%d", &class)
		for err != nil {
			fmt.Println("Type A Number From 1 to 3:")
			fmt.Print("-> ")
			_, err = fmt.Scanf("%d", &class)
		}

		fmt.Printf("Great. Your Agent Looks Like An Experienced %s.\n", agent.CLASS_TEXT[class])
		fmt.Printf("Now Pick A Weapon For Your Agent:")
		fmt.Print("-> ")

		switch class {
		case agent.CLASS.Warrior:
			fmt.Println("#1 Two-handed Sword")
			fmt.Println("#2 Buckler & Axe")
			fmt.Print("-> ")
			_, err := fmt.Scanf("%d", &weapon)
			for err != nil {
				fmt.Println("Type A Number From 1 to 2:")
				fmt.Print("-> ")
				_, err = fmt.Scanf("%d", &weapon)
			}
		case agent.CLASS.Mage:
			fmt.Println("#3 Twilight Staff")
			fmt.Println("#4 Wand Of Dark Warlock")
			fmt.Print("-> ")
			_, err := fmt.Scanf("%d", &weapon)
			for err != nil {
				fmt.Println("Type A Number From 3 to 4:")
				fmt.Print("-> ")
				_, err = fmt.Scanf("%d", &weapon)
			}
		case agent.CLASS.Archer:
			fmt.Println("#5 Longbow")
			fmt.Println("#6 Spitfire")
			fmt.Print("-> ")
			_, err := fmt.Scanf("%d", &weapon)
			for err != nil {
				fmt.Println("Type A Number From 5 to 6:")
				fmt.Print("-> ")
				_, err = fmt.Scanf("%d", &weapon)
			}
		}

		fmt.Printf("%s? This Is A Wonderful Pick.", agent.WEAPON_TEXT[weapon])

		fmt.Printf("Your Agent Is Fully Primed Now.")
		fmt.Printf("Name  : %s\n", name)
		fmt.Printf("Class : %s\n", agent.CLASS_TEXT[class])
		fmt.Printf("Weapon: %s\n", agent.WEAPON_TEXT[weapon])
		fmt.Printf("Go With This Perfect Agent?: (y/n)")
		fmt.Print("-> ")
		yn, _ := reader.ReadString('\n')
		if yn == "n" {
			continue
		}

		a := agent.New(name, class, weapon)
		a.Save()

		fmt.Println("Congratulation, You Made A Wise Choice.")
		fmt.Println("Type `myag` To Greet Your Agent.")
		break
		/*
			if strings.Compare("hi", text) == 0 {
				fmt.Println("hello, Yourself")
			}
		*/
	}

}
