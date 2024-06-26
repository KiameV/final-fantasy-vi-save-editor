package consts

import (
	"pixel-remastered-save-editor/models"
)

var (
	Abilities = []models.NameValue{}

	WhiteMagic = []models.NameValue{
		models.NewValueName(6, "Cure"),
		models.NewValueName(7, "Cura"),
		models.NewValueName(8, "Curaga"),
		models.NewValueName(9, "Curaja"),
		models.NewValueName(10, "Esuna"),
		models.NewValueName(11, "Life"),
		models.NewValueName(12, "Full-Life"),
		models.NewValueName(13, "Hold"),
		models.NewValueName(14, "Silence"),
		models.NewValueName(15, "Confuse"),
		models.NewValueName(16, "Blink"),
		models.NewValueName(17, "Protect"),
		models.NewValueName(18, "Shell"),
		models.NewValueName(19, "Slow"),
		models.NewValueName(20, "Haste"),
		models.NewValueName(21, "Berserk"),
		models.NewValueName(22, "Reflect"),
		models.NewValueName(23, "Holy"),
		models.NewValueName(24, "Dispel"),
		models.NewValueName(25, "Scan"),
		models.NewValueName(26, "Mini"),
		models.NewValueName(27, "Teleport"),
		models.NewValueName(28, "Sight"),
		models.NewValueName(29, "Float"),
	}

	BlackMagic = []models.NameValue{
		models.NewValueName(33, "Fire"),
		models.NewValueName(34, "Fira"),
		models.NewValueName(35, "Firaga"),
		models.NewValueName(36, "Blizzard"),
		models.NewValueName(37, "Blizzara"),
		models.NewValueName(38, "Blizzaga"),
		models.NewValueName(39, "Thunder"),
		models.NewValueName(40, "Thundara"),
		models.NewValueName(41, "Thundaga"),
		models.NewValueName(42, "Sleep"),
		models.NewValueName(43, "Poison"),
		models.NewValueName(44, "Warp"),
		models.NewValueName(45, "Toad"),
		models.NewValueName(46, "Stop"),
		models.NewValueName(47, "Osmose"),
		models.NewValueName(48, "Bio"),
		models.NewValueName(49, "Tornado"),
		models.NewValueName(50, "Break"),
		models.NewValueName(51, "Pig"),
		models.NewValueName(52, "Quake"),
		models.NewValueName(53, "Death"),
		models.NewValueName(54, "Meteor"),
		models.NewValueName(55, "Flare"),
		models.NewValueName(56, "Drain"),
		models.NewValueName(57, "Comet"),
		models.NewValueName(58, "??? Double Meteor"),
		models.NewValueName(59, "Pyro"),
	}

	SummonMagic = []models.NameValue{
		models.NewValueName(61, "Goblin"),
		models.NewValueName(62, "Bomb"),
		models.NewValueName(63, "Cockatrice"),
		models.NewValueName(64, "Mind Flayer"),
		models.NewValueName(65, "Chocobo"),
		models.NewValueName(66, "Shiva"),
		models.NewValueName(67, "Ramuh"),
		models.NewValueName(68, "Ifrit"),
		models.NewValueName(69, "Titan"),
		models.NewValueName(70, "Dragon"),
		models.NewValueName(71, "Sylph"),
		models.NewValueName(72, "Odin"),
		models.NewValueName(73, "Leviathan"),
		models.NewValueName(74, "Asura"),
		models.NewValueName(75, "Bahamut"),
	}
)
