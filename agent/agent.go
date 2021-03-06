package agent

import (
	"bytes"
	"encoding/gob"
	"fmt"

	log "github.com/corgi-kx/logcustom"
	"github.com/ntswamp/proof-of-kill/database"
)

type Agent struct {
	Name   string
	Class  string
	Weapon string
	Health int
	Attack int
	Luck   int //-20~20
}

func New(name string, class int, weapon int) *Agent {
	a := &Agent{}
	switch class {
	case CLASS.Warrior:
		a.Health = 200
		a.Attack = 16
		a.Luck = 5
	case CLASS.Mage:
		a.Health = 120
		a.Attack = 36
		a.Luck = 0
	case CLASS.Archer:
		a.Health = 150
		a.Attack = 16
		a.Luck = 20
	case CLASS.Paladin:
		a.Health = 220
		a.Attack = 15
		a.Luck = 11
	}

	switch weapon {
	//warrior
	case WEAPON.TwohandedSword:
		a.Attack = a.Attack + 18
		a.Luck = a.Luck + 5
	case WEAPON.BucklerAxe:
		a.Attack = a.Attack + 10
		a.Health = a.Health + 10
		a.Luck = a.Luck + 1

	//mage
	case WEAPON.TwilightStaff:
		a.Attack = a.Attack + 10
		a.Luck = 0
	case WEAPON.WandOfDarkWarlock:
		a.Attack = a.Attack + 30
		a.Health = a.Health - 30
		a.Luck = -10

	//archer
	case WEAPON.Longbow:
		a.Attack = a.Attack + 10
		a.Luck = a.Luck + 10
	case WEAPON.Spitfire:
		a.Attack = a.Attack + 20
		a.Luck = 0

		//paladin
	case WEAPON.HammerOfJudgement:
		a.Attack = a.Attack + 20
		a.Luck = 3
	case WEAPON.BoneCrusher:
		a.Attack = a.Attack + 0
		a.Luck = a.Luck + 50

	}
	a.Name = name
	a.Class = CLASS_TEXT[class]
	a.Weapon = WEAPON_TEXT[weapon]
	return a
}

func (a *Agent) Save() {
	db := database.New()
	b := db.View([]byte("MYAGENT"), database.AgentBucket)
	if len(b) != 0 {
		log.Warn("Agent Already Exists.")
		return
	}
	db.Put([]byte("MYAGENT"), a.Serliazle(), database.AgentBucket)

}

func Load() *Agent {
	db := database.New()
	b := db.View([]byte("MYAGENT"), database.AgentBucket)
	if len(b) == 0 {
		log.Warn("Agent Not Found. Type 'newag' To Create One.")
		return nil
	}
	a := &Agent{}
	a.Deserialize(b)
	return a

}

func Remove() {
	db := database.New()
	db.Delete([]byte("MYAGENT"), database.AgentBucket)
}

func IsAgentExist() bool {
	db := database.New()
	b := db.View([]byte("MYAGENT"), database.AgentBucket)
	return len(b) != 0
}

func (a *Agent) Introduce() {
	fmt.Printf("greeting to %v: the great %v wielding with %v.\n", a.Name, a.Class, a.Weapon)
}

func (a *Agent) Serliazle() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(a)
	if err != nil {
		panic(err)
	}
	return result.Bytes()
}

func (a *Agent) Deserialize(b []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(a)
	if err != nil {
		log.Panic(err)
	}
}

/**
*
*   for duel
*
**/

//deal a.Attak + rand damage
func (a *Agent) DealDamage(rand int) int {
	damage := a.Attack + rand
	return damage
}

func (a *Agent) TakeDamage(damage int) {
	a.Health = a.Health - damage
}

func (a *Agent) IsDied() bool {
	return a.Health <= 0
}
