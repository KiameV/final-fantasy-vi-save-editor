package pr

import "ffvi_editor/models/consts"

const (
	SpellFrom   int64 = 31
	SpellTo     int64 = 84
	SpellOffset       = 330
)

var (
	Spells = []*consts.NameValue{
		consts.NewValueName(31, "Cure"),
		consts.NewValueName(32, "Cura"),
		consts.NewValueName(33, "Curaga"),
		consts.NewValueName(34, "Raise"),
		consts.NewValueName(35, "Arise"),
		consts.NewValueName(36, "Poisona"),
		consts.NewValueName(37, "Esuna"),
		consts.NewValueName(38, "Regen"),
		consts.NewValueName(39, "Reraise"),
		consts.NewValueName(40, "Fire"),
		consts.NewValueName(41, "Blizzard"),
		consts.NewValueName(42, "Thunder"),
		consts.NewValueName(43, "Poison"),
		consts.NewValueName(44, "Drain"),
		consts.NewValueName(45, "Fira"),
		consts.NewValueName(46, "Blizzara"),
		consts.NewValueName(47, "Thundara"),
		consts.NewValueName(48, "Bio"),
		consts.NewValueName(49, "Firaga"),
		consts.NewValueName(50, "Blizzaga"),
		consts.NewValueName(51, "Thundaga"),
		consts.NewValueName(52, "Break"),
		consts.NewValueName(53, "Death"),
		consts.NewValueName(54, "Holy"),
		consts.NewValueName(55, "Flare"),
		consts.NewValueName(56, "Gravity"),
		consts.NewValueName(57, "Graviga"),
		consts.NewValueName(58, "Banish"),
		consts.NewValueName(59, "Meteor"),
		consts.NewValueName(60, "Ultima"),
		consts.NewValueName(61, "Quake"),
		consts.NewValueName(62, "Tornado"),
		consts.NewValueName(63, "Meltdown"),
		consts.NewValueName(64, "Libra"),
		consts.NewValueName(65, "Slow"),
		consts.NewValueName(66, "Rasp"),
		consts.NewValueName(67, "Silence"),
		consts.NewValueName(68, "Protect"),
		consts.NewValueName(69, "Sleep"),
		consts.NewValueName(70, "Confuse"),
		consts.NewValueName(71, "Haste"),
		consts.NewValueName(72, "Stop"),
		consts.NewValueName(73, "Berserk"),
		consts.NewValueName(74, "Float"),
		consts.NewValueName(75, "Imp"),
		consts.NewValueName(76, "Reflect"),
		consts.NewValueName(77, "Shell"),
		consts.NewValueName(78, "Vanish"),
		consts.NewValueName(79, "Hastega"),
		consts.NewValueName(80, "Slowga"),
		consts.NewValueName(81, "Osmose"),
		consts.NewValueName(82, "Teleport"),
		consts.NewValueName(83, "Quick"),
		consts.NewValueName(84, "Dispel"),
	}
	SortedSpells    []*consts.NameValue
	SpellLookupByID = make(map[int]*consts.NameValue)
)

func init() {
	SortedSpells = consts.SortByNameValue(Spells)
	for _, l := range Spells {
		SpellLookupByID[l.Value] = l
	}
}
