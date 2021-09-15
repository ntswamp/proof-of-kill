package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ntswamp/proof-of-kill/agent"
)

/// return true if the creation succeeds.
func (cli *Cli) newAg() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("people, you're looking fraught. you must heard that there will be blood when the dark falls later.\nyes, it will.\nbut I know someone hired a strong \"Agent\" to have him survived the killing last night,\nyou may want to take a chance on tavern.(y/n)")
	fmt.Print("-> ")
	yn, _ := reader.ReadString('\n')
	if yn == "n\n" {
		os.Exit(0)
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

		fmt.Printf("Such An Impressive Name.\nWelcome To The World Of PoK, %s.\n\n", name)

		fmt.Println("Now Tell Me Your Class, By A Number:")
		fmt.Println("#1 Warrior")
		fmt.Println("#2 Mage")
		fmt.Println("#3 Archer")
		fmt.Print("-> ")
		_, err := fmt.Scanf("%d", &class)
		for err != nil || (class != 1 && class != 2 && class != 3) {
			fmt.Println("Type A Number From 1 to 3:")
			fmt.Print("-> ")
			_, err = fmt.Scanf("%d", &class)
		}

		fmt.Printf("Great. Your Agent Looks Like An Experienced %s.\n", agent.CLASS_TEXT[class])
		fmt.Printf("\nNow Pick A Weapon For Your Agent:\n")

		switch class {
		case agent.CLASS.Warrior:
			fmt.Println("#1 Two-handed Sword")
			fmt.Println("#2 Buckler & Axe")
			fmt.Print("-> ")
			_, err := fmt.Scanf("%d", &weapon)
			for err != nil || (weapon != 1 && weapon != 2) {
				fmt.Println("Type A Number From 1 to 2:")
				fmt.Print("-> ")
				_, err = fmt.Scanf("%d", &weapon)
			}
		case agent.CLASS.Mage:
			fmt.Println("#3 Twilight Staff")
			fmt.Println("#4 Wand Of Dark Warlock")
			fmt.Print("-> ")
			_, err := fmt.Scanf("%d", &weapon)
			for err != nil || (weapon != 3 && weapon != 4) {
				fmt.Println("Type A Number From 3 to 4:")
				fmt.Print("-> ")
				_, err = fmt.Scanf("%d", &weapon)
			}
		case agent.CLASS.Archer:
			fmt.Println("#5 Longbow")
			fmt.Println("#6 Spitfire")
			fmt.Print("-> ")
			_, err := fmt.Scanf("%d", &weapon)
			for err != nil || (weapon != 5 && weapon != 6) {
				fmt.Println("Type A Number From 5 to 6:")
				fmt.Print("-> ")
				_, err = fmt.Scanf("%d", &weapon)
			}
		}

		fmt.Printf("%s? This Is A Wonderful Pick.\n\n", agent.WEAPON_TEXT[weapon])

		fmt.Printf("Your Agent Is Fully Primed Now.\n")
		fmt.Printf("Name  : %s\n", name)
		fmt.Printf("Class : %s\n", agent.CLASS_TEXT[class])
		fmt.Printf("Weapon: %s\n", agent.WEAPON_TEXT[weapon])
		fmt.Printf("\nGo with this perfect agent?: (y/n)\n")
		fmt.Print("-> ")
		yn, _ := reader.ReadString('\n')
		if yn == "n\n" {
			continue
		}

		a := agent.New(name, class, weapon)
		a.Save()

		fmt.Println("Congratulation, You Made A Wise Choice.")
		fmt.Println("Type `myag` To Greet Your Agent.")
		return
		/*
			if strings.Compare("hi", text) == 0 {
				fmt.Println("hello, Yourself")
			}
		*/
	}
}
