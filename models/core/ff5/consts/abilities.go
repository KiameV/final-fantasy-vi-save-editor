package consts

import (
	"pixel-remastered-save-editor/models"
)

var (
	Abilities = []models.NameValue{
		models.NewValueName(6, "Guard"),
		models.NewValueName(7, "Kick"),
		models.NewValueName(8, "Focus"),
		models.NewValueName(9, "Chakra"),
		models.NewValueName(10, "Scram"),
		models.NewValueName(11, "Steal"),
		models.NewValueName(12, "Mug"),
		models.NewValueName(13, "Jump"),
		models.NewValueName(14, "Lance"),
		models.NewValueName(15, "Smoke"),
		models.NewValueName(16, "Image"),
		models.NewValueName(17, "Throw"),
		models.NewValueName(18, "Mineuchi"),
		models.NewValueName(19, "Zeninage"),
		models.NewValueName(20, "Iainuki"),
		models.NewValueName(21, "Animals"),
		models.NewValueName(22, "Aim"),
		models.NewValueName(23, "Rapid Fire"),
		models.NewValueName(24, "Call"),
		models.NewValueName(25, "Check"),
		models.NewValueName(26, "Scan"),
		models.NewValueName(27, "Calm"),
		models.NewValueName(28, "Control"),
		models.NewValueName(29, "Catch"),
		models.NewValueName(30, "Release"),
		models.NewValueName(31, "Mix"),
		models.NewValueName(32, "Drink"),
		models.NewValueName(33, "Recover"),
		models.NewValueName(34, "Revive"),
		models.NewValueName(35, "Gaia"),
		models.NewValueName(37, "Hide"),
		models.NewValueName(38, "Reveal"),
		models.NewValueName(41, "Flirt"),
		models.NewValueName(42, "Dance"),
		models.NewValueName(43, "Mimic"),
		models.NewValueName(51, "Equip Shields"),
		models.NewValueName(52, "Equip Armor"),
		models.NewValueName(53, "Equip Ribbons"),
		models.NewValueName(54, "Equip Swords"),
		models.NewValueName(55, "Equip Lances"),
		models.NewValueName(56, "Equip Katanas"),
		models.NewValueName(57, "Equip Axes"),
		models.NewValueName(58, "Equip Bows"),
		models.NewValueName(59, "Equip Whips"),
		models.NewValueName(60, "Equip Harps"),
		models.NewValueName(61, "Artful Dodger"),
		models.NewValueName(62, "HP +10%"),
		models.NewValueName(63, "HP +20%"),
		models.NewValueName(64, "HP +30%"),
		models.NewValueName(65, "MP +10%"),
		models.NewValueName(66, "MP +30%"),
		models.NewValueName(67, "Barehanded"),
		models.NewValueName(68, "Two-Handed"),
		models.NewValueName(69, "Dual-Wield"),
		models.NewValueName(70, "Pharmacology"),
		models.NewValueName(71, "Cover"),
		models.NewValueName(72, "Counter"),
		models.NewValueName(73, "Shirahadori"),
		models.NewValueName(74, "Learn"),
		models.NewValueName(75, "Magic Shell"),
		models.NewValueName(76, "Berserk"),
		models.NewValueName(77, "Vigilance"),
		models.NewValueName(78, "First Strike"),
		models.NewValueName(79, "Find Passages"),
		models.NewValueName(80, "Light Step"),
		models.NewValueName(81, "Find Pits"),
		models.NewValueName(82, "Equip Rods"),
		models.NewValueName(83, "Sprint"),
	}
	WeaponAbilities = []models.NameValue{
		models.NewValueName(90, "Fire"),
		models.NewValueName(91, "Blizzard"),
		models.NewValueName(92, "Thunder"),
		models.NewValueName(93, "Poison"),
		models.NewValueName(94, "Silence"),
		models.NewValueName(95, "Sleep"),
		models.NewValueName(96, "Fira"),
		models.NewValueName(97, "Blizzara"),
		models.NewValueName(98, "Thundara"),
		models.NewValueName(99, "Drain"),
		models.NewValueName(100, "Break"),
		models.NewValueName(101, "Bio"),
		models.NewValueName(102, "Firaga"),
		models.NewValueName(103, "Blizzaga"),
		models.NewValueName(104, "Thundaga"),
		models.NewValueName(105, "Holy"),
		models.NewValueName(106, "Flare"),
		models.NewValueName(107, "Osmose"),
	}

	WhiteMagic = []models.NameValue{
		models.NewValueName(108, "Cure"),
		models.NewValueName(109, "Libra"),
		models.NewValueName(110, "Poisona"),
		models.NewValueName(111, "Silence"),
		models.NewValueName(112, "Protect"),
		models.NewValueName(113, "Mini"),
		models.NewValueName(114, "Cura"),
		models.NewValueName(115, "Raise"),
		models.NewValueName(116, "Confuse"),
		models.NewValueName(117, "Blink"),
		models.NewValueName(118, "Shell"),
		models.NewValueName(119, "Esuna"),
		models.NewValueName(120, "Curaga"),
		models.NewValueName(121, "Reflect"),
		models.NewValueName(122, "Berserk"),
		models.NewValueName(123, "Arise"),
		models.NewValueName(124, "Holy"),
		models.NewValueName(125, "Dispel"),
	}

	BlackMagic = []models.NameValue{
		models.NewValueName(126, "Fire"),
		models.NewValueName(127, "Blizzard"),
		models.NewValueName(128, "Thunder"),
		models.NewValueName(129, "Poison"),
		models.NewValueName(130, "Sleep"),
		models.NewValueName(131, "Toad"),
		models.NewValueName(132, "Fira"),
		models.NewValueName(133, "Blizzara"),
		models.NewValueName(134, "Thundara"),
		models.NewValueName(135, "Drain"),
		models.NewValueName(136, "Break"),
		models.NewValueName(137, "Bio"),
		models.NewValueName(138, "Firaga"),
		models.NewValueName(139, "Blizzaga"),
		models.NewValueName(140, "Thundaga"),
		models.NewValueName(141, "Flare"),
		models.NewValueName(142, "Death"),
		models.NewValueName(143, "Osmose"),
	}

	SummonMagic = []models.NameValue{
		models.NewValueName(162, "Chocobo"),
		models.NewValueName(163, "Sylph"),
		models.NewValueName(164, "Remora"),
		models.NewValueName(165, "Shiva"),
		models.NewValueName(166, "Ramuh"),
		models.NewValueName(167, "Ifrit"),
		models.NewValueName(168, "Titan"),
		models.NewValueName(169, "Golem"),
		models.NewValueName(170, "Catoblepas"),
		models.NewValueName(171, "Carbuncle"),
		models.NewValueName(172, "Syldra"),
		models.NewValueName(173, "Odin"),
		models.NewValueName(174, "Phoenix"),
		models.NewValueName(175, "Leviathan"),
		models.NewValueName(176, "Bahamut"),
	}

	TimeMagic = []models.NameValue{
		models.NewValueName(144, "Speed"),
		models.NewValueName(145, "Slow"),
		models.NewValueName(146, "Regen"),
		models.NewValueName(147, "Mute"),
		models.NewValueName(148, "Haste"),
		models.NewValueName(149, "Float"),
		models.NewValueName(150, "Gravity"),
		models.NewValueName(151, "Stop"),
		models.NewValueName(152, "Teleport"),
		models.NewValueName(153, "Comet"),
		models.NewValueName(154, "Slowga"),
		models.NewValueName(155, "Return"),
		models.NewValueName(156, "Graviga"),
		models.NewValueName(157, "Hastega"),
		models.NewValueName(158, "Old"),
		models.NewValueName(159, "Meteor"),
		models.NewValueName(160, "Quick"),
		models.NewValueName(161, "Banish"),
	}
)
