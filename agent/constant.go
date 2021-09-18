package agent

var CLASS = struct {
	Warrior int
	Mage    int
	Archer  int
	Paladin int
}{
	Warrior: 1,
	Mage:    2,
	Archer:  3,
	Paladin: 4,
}
var CLASS_TEXT = map[int]string{
	CLASS.Warrior: "Warrior",
	CLASS.Mage:    "Mage",
	CLASS.Archer:  "Archer",
	CLASS.Paladin: "Paladin",
}

var WEAPON = struct {
	TwohandedSword    int
	BucklerAxe        int
	TwilightStaff     int
	WandOfDarkWarlock int
	Longbow           int
	Spitfire          int
	HammerOfJudgement int
	BoneCrusher       int
}{
	TwohandedSword:    1,
	BucklerAxe:        2,
	TwilightStaff:     3,
	WandOfDarkWarlock: 4,
	Longbow:           5,
	Spitfire:          6,
	HammerOfJudgement: 7,
	BoneCrusher:       8,
}
var WEAPON_TEXT = map[int]string{
	WEAPON.TwohandedSword:    "Two-handed Sword",
	WEAPON.BucklerAxe:        "Buckler & Axe",
	WEAPON.TwilightStaff:     "Twilight Staff",
	WEAPON.WandOfDarkWarlock: "Wand Of Dark Warlock",
	WEAPON.Longbow:           "Longbow",
	WEAPON.Spitfire:          "Spitfire",
	WEAPON.HammerOfJudgement: "Hammer Of Judgement",
	WEAPON.BoneCrusher:       "Bone Crusher",
}

var GENESIS_AGENT = New("PoKGod", CLASS.Mage, WEAPON.TwilightStaff)
