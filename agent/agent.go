package agent

import (
	"bytes"
	"encoding/gob"

	log "github.com/corgi-kx/logcustom"
	"github.com/ntswamp/proof-of-kill/database"
)

type Agent struct {
	name   string
	class  string
	weapon string
	health int
	attack int
	luck   int //-20~20
}

func New(name string, class int, weapon int) *Agent {
	a := &Agent{}
	switch class {
	case 1:
		a.class = "Warrior"
		a.health = 200
		a.attack = 20
		a.luck = 5
	case 2:
		a.class = "Mage"
		a.health = 120
		a.attack = 35
		a.luck = 0
	//archer
	case 3:
		a.class = "Archer"
		a.health = 150
		a.attack = 20
		a.luck = 20
	default:
		a.class = "Warrior"
		a.health = 200
		a.attack = 20
		a.luck = 5
	}

	switch weapon {
	//warrior
	case 1:
		a.weapon = "Two-handed Sword"
		a.attack = a.attack + 18
		a.luck = a.luck + 5
	case 2:
		a.weapon = "Buckler & Axe"
		a.attack = a.attack + 10
		a.health = a.health + 10
		a.luck = a.luck + 1

	//mage
	case 3:
		a.weapon = "Twilight Staff"
		a.attack = a.attack + 10
		a.luck = 0
	case 4:
		a.weapon = "Wand Of Dark Warlock"
		a.attack = a.attack + 30
		a.health = a.health - 30
		a.luck = -10

	//archer
	case 5:
		a.weapon = "Longbow"
		a.attack = a.attack + 10
		a.luck = a.luck + 10
	case 6:
		a.weapon = "Spitfire"
		a.attack = a.attack + 20
		a.luck = 0

	}
	a.name = name
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

func IsAgentExist(nodeId string) bool {
	db := database.New()
	b := db.View([]byte("MYAGENT"), database.AgentBucket)
	return len(b) != 0
}

func (a *Agent) Introduce() {
	log.Infof("Introduce Agent: %v, the %v wielding a %v.\n", a.name, a.class, a.weapon)
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
