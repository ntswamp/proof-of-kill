package agent

import (
	"bytes"
	"encoding/gob"

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
	case 1:
		a.Class = "Warrior"
		a.Health = 200
		a.Attack = 20
		a.Luck = 5
	case 2:
		a.Class = "Mage"
		a.Health = 120
		a.Attack = 35
		a.Luck = 0
	//archer
	case 3:
		a.Class = "Archer"
		a.Health = 150
		a.Attack = 20
		a.Luck = 20
	default:
		a.Class = "Warrior"
		a.Health = 200
		a.Attack = 20
		a.Luck = 5
	}

	switch weapon {
	//warrior
	case 1:
		a.Weapon = "Two-handed Sword"
		a.Attack = a.Attack + 18
		a.Luck = a.Luck + 5
	case 2:
		a.Weapon = "Buckler & Axe"
		a.Attack = a.Attack + 10
		a.Health = a.Health + 10
		a.Luck = a.Luck + 1

	//mage
	case 3:
		a.Weapon = "Twilight Staff"
		a.Attack = a.Attack + 10
		a.Luck = 0
	case 4:
		a.Weapon = "Wand Of Dark Warlock"
		a.Attack = a.Attack + 30
		a.Health = a.Health - 30
		a.Luck = -10

	//archer
	case 5:
		a.Weapon = "Longbow"
		a.Attack = a.Attack + 10
		a.Luck = a.Luck + 10
	case 6:
		a.Weapon = "Spitfire"
		a.Attack = a.Attack + 20
		a.Luck = 0

	}
	a.Name = name
	return a
}

func (a *Agent) Save() {
	db := database.New()
	b := db.View([]byte("MYAGENT"), database.AgentBucket)
	if len(b) != 0 {
		log.Warn("Agent Already Exists.")
		return
	}
	db.Put([]byte("MYAGENT"), a.serliazle(), database.AgentBucket)

}

func Load() *Agent {
	db := database.New()
	b := db.View([]byte("MYAGENT"), database.AgentBucket)
	if len(b) == 0 {
		log.Warn("Agent Not Found. Type 'newag' To Create One.")
		return nil
	}
	a := &Agent{}
	a.deserialize(b)
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
	log.Infof("Introduce Agent: %v, the %v wielding a %v.\n", a.Name, a.Class, a.Weapon)
}

func (a *Agent) serliazle() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(a)
	if err != nil {
		panic(err)
	}
	return result.Bytes()
}

func (a *Agent) deserialize(b []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(a)
	if err != nil {
		log.Panic(err)
	}
}
