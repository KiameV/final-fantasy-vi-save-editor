package consts

import (
	"pixel-remastered-save-editor/models"
)

var (
	Abilities = []models.NameValue{}

	WhiteMagic = []models.NameValue{models.NewValueName(6, "Cure"),
		models.NewValueName(7, "Poisona"),
		models.NewValueName(8, "Sight"),
		models.NewValueName(9, "Aero"),
		models.NewValueName(10, "Toad"),
		models.NewValueName(11, "Mini"),
		models.NewValueName(12, "Cura"),
		models.NewValueName(13, "Teleport"),
		models.NewValueName(14, "Blindna"),
		models.NewValueName(15, "Libra"),
		models.NewValueName(16, "Confuse"),
		models.NewValueName(17, "Silence"),
		models.NewValueName(18, "Curaga"),
		models.NewValueName(19, "Raise"),
		models.NewValueName(20, "Protect"),
		models.NewValueName(21, "Aeroga"),
		models.NewValueName(22, "Stona"),
		models.NewValueName(23, "Haste"),
		models.NewValueName(24, "Curaja"),
		models.NewValueName(25, "Esuna"),
		models.NewValueName(26, "Reflect"),
		models.NewValueName(27, "Tornado"),
		models.NewValueName(28, "Arise"),
		models.NewValueName(29, "Holy"),
	}

	BlackMagic = []models.NameValue{
		models.NewValueName(30, "Fire"),
		models.NewValueName(31, "Blizzard"),
		models.NewValueName(32, "Sleep"),
		models.NewValueName(33, "Thunder"),
		models.NewValueName(34, "Poison"),
		models.NewValueName(35, "Blind"),
		models.NewValueName(36, "Fira"),
		models.NewValueName(37, "Blizzara"),
		models.NewValueName(38, "Thundara"),
		models.NewValueName(39, "Break"),
		models.NewValueName(40, "Blizzaga"),
		models.NewValueName(41, "Shade"),
		models.NewValueName(42, "Thundaga"),
		models.NewValueName(43, "Raze"),
		models.NewValueName(44, "Erase"),
		models.NewValueName(45, "Firaga"),
		models.NewValueName(46, "Bio"),
		models.NewValueName(47, "Warp"),
		models.NewValueName(48, "Quake"),
		models.NewValueName(49, "Breakga"),
		models.NewValueName(50, "Drain"),
		models.NewValueName(51, "Flare"),
		models.NewValueName(52, "Death"),
		models.NewValueName(53, "Meteor"),
	}

	SummonMagic = []models.NameValue{
		models.NewValueName(54, "Escape"),
		models.NewValueName(55, "Icen"),
		models.NewValueName(56, "Spark"),
		models.NewValueName(57, "Heatra"),
		models.NewValueName(58, "Hyper"),
		models.NewValueName(59, "Catastro"),
		models.NewValueName(60, "Leviath"),
		models.NewValueName(61, "Bahamur"),
	}
)
