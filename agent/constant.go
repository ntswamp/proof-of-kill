package agent

var CLASS = struct {
	Warrior int
	Mage    int
	Archer  int
}{
	Warrior: 1,
	Mage:    2,
	Archer:  3,
}
var CLASS_TEXT = map[int]string{
	CLASS.Warrior: "Warrior",
	CLASS.Mage:    "Mage",
	CLASS.Archer:  "Archer",
}

var WEAPON = struct {
	TwohandedSword    int
	BucklerAxe        int
	TwilightStaff     int
	WandOfDarkWarlock int
	Longbow           int
	Spitfire          int
}{
	TwohandedSword:    1,
	BucklerAxe:        2,
	TwilightStaff:     3,
	WandOfDarkWarlock: 4,
	Longbow:           5,
	Spitfire:          6,
}
var WEAPON_TEXT = map[int]string{
	WEAPON.TwohandedSword:    "Two-handed Sword",
	WEAPON.BucklerAxe:        "Buckler & Axe",
	WEAPON.TwilightStaff:     "Twilight Staff",
	WEAPON.WandOfDarkWarlock: "Wand Of Dark Warlock",
	WEAPON.Longbow:           "Longbow",
	WEAPON.Spitfire:          "Spitfire",
}
