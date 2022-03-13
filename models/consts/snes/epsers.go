package snes

import "ffvi_editor/models/consts"

var (
	Espers = consts.NewNameSlotMask8s(
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
	SortedEspers []*consts.NameSlotMask8
)

func init() {
	SortedEspers = consts.SortByName(Espers)
}
