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
	fmt.Println("People, you're looking fraught. you must heard that there will be blood when the dark falls later.\nyes, it will.\nbut I know someone hired a brave \"Agent\" to have him survived the killing last night,\nyou may want to take a chance on tavern.(y/n)")
	fmt.Print("-> ")
	yn, _ := reader.ReadString('\n')
	if yn == "n\n" {
		os.Exit(0)
	}

	for {
		var name string
		var class int
		var weapon int

		fmt.Println("Incredibly, a job-hunter-looking agent noticed your visit. you made it with no hesitation.\n\nWhat's his name?")
		fmt.Print("-> ")
		name, _ = reader.ReadString('\n')
		// convert CRLF to LF
		name = strings.Replace(name, "\n", "", -1)

		fmt.Printf("\n%s may help you surviving the world of PoK.\n\n", name)

		fmt.Println("Now tell me his class, by a number:")
		fmt.Println("#1 Warrior")
		fmt.Println("#2 Mage")
		fmt.Println("#3 Archer")
		fmt.Println("#4 Paladin")
		fmt.Print("-> ")
		_, err := fmt.Scanf("%d", &class)
		for err != nil || (class != 1 && class != 2 && class != 3 && class != 4) {
			fmt.Println("Type a number from 1 to 4:")
			fmt.Print("-> ")
			_, err = fmt.Scanf("%d", &class)
		}

		fmt.Printf("Your agent looks like an experienced %s, great.\n", agent.CLASS_TEXT[class])
		fmt.Printf("\nMaybe go acquire a weapon for your man?\n")

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
		case agent.CLASS.Paladin:
			fmt.Println("#7 Hammer Of Judgement")
			fmt.Println("#8 Bone Crusher")
			fmt.Print("-> ")
			_, err := fmt.Scanf("%d", &weapon)
			for err != nil || (weapon != 7 && weapon != 8) {
				fmt.Println("Type A Number From 7 to 8:")
				fmt.Print("-> ")
				_, err = fmt.Scanf("%d", &weapon)
			}
		}

		fmt.Printf("%s? definitely a wonderful pick.\n\n", agent.WEAPON_TEXT[weapon])

		fmt.Printf("%s is fully primed now.\n", name)
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

		fmt.Println("Congratulation, you made a wise choice.")
		fmt.Println("Type `myag` to greet your agent now.")
		return
		/*
			if strings.Compare("hi", text) == 0 {
				fmt.Println("hello, Yourself")
			}
		*/
	}
}
