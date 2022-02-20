package consts

var (
	Espers = NewNameSlotMask8s(
		"Ramuh",
		"Ifrit",
		"Shiva",
		"Siren",
		"Terrato",
		"Maduin",
		"Shoat",
		"Bismark",
		"Stray",
		"Palidor",
		"Tritoch",
		"Odin",
		"Raiden",
		"Bahamut",
		"Alexander",
		"Crusader",
		"Ragnarok",
		"Kirin",
		"Zoneseek",
		"Carbunkl",
		"Phantom",
		"Sraphim",
		"Golem",
		"Unicorn",
		"Fenrir",
		"Startlet",
		"Phoenix",
	)
	SortedEspers []*NameSlotMask8
)

func init() {
	SortedEspers = sortByName(Espers)
}
