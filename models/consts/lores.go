package consts

/*type Lore uint32

const (
	Condemned  Lore = 0x1
	Roulette   Lore = 0x2
	CleanSweep Lore = 0x4
	AquaRake   Lore = 0x8
	Aero       Lore = 0x10
	BlowFish   Lore = 0x20
	BigGuard   Lore = 0x40
	Revenge    Lore = 0x80

	PearlWind  Lore = 0x100
	L5Doom     Lore = 0x200
	L4Flare    Lore = 0x400
	L3Muddle   Lore = 0x800
	Reflect    Lore = 0x1000
	LPearl     Lore = 0x2000
	StepMine   Lore = 0x4000
	ForceField Lore = 0x8000

	Dischord   Lore = 0x10000
	SourMouth  Lore = 0x20000
	PepUp      Lore = 0x40000
	Rippler    Lore = 0x80000
	Stone      Lore = 0x100000
	Quasar     Lore = 0x200000
	GrandTrain Lore = 0x400000
	Exploder   Lore = 0x800000
)*/

var (
	Lores = NewNameSlotMask8s(
		"Condemned",
		"Roulette",
		"CleanSweep",
		"AquaRake",
		"Aero",
		"BlowFish",
		"BigGuard",
		"Revenge",
		"PearlWind",
		"L5Doom",
		"L4Flare",
		"L3Muddle",
		"Reflect",
		"LPearl",
		"StepMine",
		"ForceField",
		"Dischord",
		"SourMouth",
		"PepUp",
		"Rippler",
		"Stone",
		"Quasar",
		"GrandTrain",
		"Exploder",
	)
	SortedLores []*NameSlotMask8
)

func init() {
	SortedLores = sortByName(Lores)
}
