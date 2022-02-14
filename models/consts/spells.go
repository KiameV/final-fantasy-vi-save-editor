package consts

import "sort"

var (
	Spells = []string{
		"Fire",
		"Ice",
		"Bolt",
		"Poison",
		"Drain",
		"Fire2",
		"Ice2",
		"Bolt2",
		"Bio",
		"Fire3",
		"Ice3",
		"Bolt3",
		"Break",
		"Doom",
		"Pearl",
		"Flare",
		"Demi",
		"Quarter",
		"XZone",
		"Meteor",
		"Ultima",
		"Quake",
		"WWind",
		"Metron",
		"Scan",
		"Slow",
		"Rasp",
		"Mute",
		"Safe",
		"Sleep",
		"Muddle",
		"Haste",
		"Stop",
		"Berserk",
		"Float",
		"Imp",
		"Reflect",
		"Shell",
		"Vanish",
		"Haste2",
		"Slow2",
		"Osmose",
		"Warp",
		"Quick",
		"Dispel",
		"Cure",
		"Cure2",
		"Cure3",
		"Life",
		"Life2",
		"Antidote",
		"Remedy",
		"Regen",
		"Life3",
	}
	SortedSpells []string
)

func init() {
	SortedSpells = make([]string, len(Spells))
	for i, s := range Spells {
		SortedSpells[i] = s
	}
	sort.Strings(SortedSpells)
}
