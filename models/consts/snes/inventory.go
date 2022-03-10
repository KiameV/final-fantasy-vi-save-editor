package snes

const (
	EmptyText = `FF - Empty
`
	ItemsText = `Miscellaneous
E7 - Rename Card
E8 - Tonic
E9 - Potion
EA - X-Potion
EB - Tincture
EC - Ether
ED - X-Ether
EE - Elixir
EF - Megalixir
F0 - Fenix Down
F1 - Revivify
F2 - Antidote
F3 - Eydrop
F4 - Soft
F5 - Remedy
F6 - Sleeping Bag
F7 - Tent
F8 - Green Cherry
F9 - Magicite
FA - Super Ball
FB - Echo Screen
FC - Smoke Bomb
FD - Warp Stone
FE - Dried Meat

Tool
A3 - NoiseBlaster
A4 - Bio Blaster
A5 - Flash
A6 - Chain Saw
A7 - Debilitator
A8 - Drill
A9 - Air Anchor
AA - AutoCrossbow

Skean/Edge
AB - Fire Skean
AC - Water Edge
AD - Bolt Edge
AE - Inviz Edge
AF - Shadow Edge`

	WeaponShieldText1 = `Dirk
00 - Dirk
01 - Mithril Knife
02 - Guardian
03 - Air Lancet
04 - ThiefKnife
05 - Assassin
06 - Man Eater
07 - Sword Breaker
08 - Graedus
09 - ValiantKnife

Sword
0A - MithrilBlade
0B - RegalCutlass
0C - Rune Edge
0D - Flame Sabre
0E - Blizzard
0F - Thunder Blade
10 - Epee
11 - Break Blade
12 - Drainer
13 - Enhancer
14 - Crystal
15 - Falchion
16 - Soul Sabre
17 - Ogre Nix
18 - Excalibur
19 - Scimiter
1A - Illumina
1B - Ragnarok
1C - Atma Weapon

Lance
1D - Mithril Pike
1E - Trident
1F - Stout Spear
20 - Partisan
21 - Pearl Lance
22 - Gold Lance
23 - Aura Lance
24 - Imp Halberd

Dirk
25 - Imperial
26 - Kodachi
27 - Blossom
28 - Hardened
29 - Striker
2A - Stunner

Knife
2B - Ashura
2C - Kotetsu
2D - Forged
2E - Tempest
2F - Murasame
30 - Aura
31 - Strato
32 - Sky Render

Rod
33 - Heal Rod
34 - Mithril Rod
35 - Fire Rod
36 - Ice Rod
37 - Thunder Rod
38 - Poison Rod
39 - Pearl Rod
3A - Gravity Rod
3B - Punisher
3C - Magus Rod

Brush
3D - Chocobo Brush
3E - DaVinci Brush
3F - Magical Brush
40 - Rainbow Brush

Stars
41 - Shuriken
42 - Ninja Star
43 - Tack Star

Special
44 - Flail
45 - Full Moon
46 - Morning Star
47 - Boomerang
48 - Rising Sun
49 - Hawk Eye
4A - Bone Club
4B - Sniper
4C - Wing Edge

Gambler
4D - Cards
4E - Darts
4F - Doom Darts
50 - Trump
51 - Dice
52 - Fixed Dice

Claw
53 - Metal Knuckle
54 - Mithril Claw
55 - Kaiser
56 - Poison Claw
57 - Fire Knuckle
58 - Dragon Claw
59 - Tiger Fangs`

	WeaponShieldText2 = `Shield
5A - Buckler
5B - Heavy Shield
5C - Mithril Shield
5D - Gold Shield
5E - Aegis Shield
5F - Diamond Shield
60 - Flame Shield
61 - Ice Shield
62 - Thunder Shield
63 - Crystal Shield
64 - Genji Shield
65 - Tortoise Shield
66 - Cursed Shield
67 - Paladin Shield
68 - Force Shield`

	HelmetArmorText1 = `Helmet
69 - Leather Hat
6A - Hair Band
6B - Plumed Hat
6C - Beret
6D - Magus Hat
6E - Bandana
6F - Iron Helmet
70 - Coronet
71 - Bard's Hat
72 - Green Beret
73 - Head Band
74 - Mithril Helm
75 - Tiara
76 - Gold Helmet
77 - Tiger Mask
78 - Red Hat
79 - Mystery Veil
7A - Circlet
7B - Regal Crown
7C - Diamond Helm
7D - Dark Hood
7E - Crystal Helm
7F - Oath Veil
80 - Cat Hood
81 - Genji Helmet
82 - Thornlet
83 - Titanium`

	HelmetArmorText2 = `Armor
84 - Leather Armor
85 - Cotton Robe
86 - Kung Fu Suit
87 - Iron Armor
88 - Silk Robe
89 - Mithril Vest
8A - Ninja Gear
8B - White Dress
8C - Mithril Mail
8D - Gaia Gear
8E - Mirage Dress
8F - Gold Armor
90 - Power Sash
91 - Light Robe
92 - Diamond Vest
93 - Red Jacket
94 - Force Armor
95 - Diamond Armor
96 - Dark Gear
97 - Tao Robe
98 - Crystal Mail
99 - Czarina Gown
9A - Genji Armor
9B - Imp's Armor
9C - Minerva
9D - Tabby Suit
9E - Chocobo Suit
9F - Moogle Suit
A0 - Nutkin Suit
A1 - Behemeth Suit
A2 - Snow Muffler`

	RelicText1 = `Relic
B0 - Goggles
B1 - Star Pendant
B2 - Peace Ring
B3 - Amulet
B4 - White Cape
B5 - Jewel Ring
B6 - Fair Ring
B7 - Barrier Ring
B8 - Mithril Glove
B9 - Guard Ring
BA - Running Shoes
BB - Wall Ring
BC - Cherub Down
BD - Cure Ring
BE - True Knight
BF - DragoonBoots
C0 - Zephyr Cape
C1 - Czarina Ring
C2 - Cursed Cing
C3 - Earrings
C4 - Atlas Armlet
C5 - Blizzard Ring
C6 - Rage Ring
C7 - Sneak Ring
`
	RelicText2Header = `Relic (continued)
`
	RelicText2 = `C8 - Pod Bracelet
C9 - Hero Ring
CA - Ribbon
CB - Muscle Belt
CC - Crystal Orb
CD - Gold Hairpin
CE - Economizer
CF - Thief Glove
D0 - Gauntlet
D1 - Genji Glove
D2 - Hyper Wrist
D3 - Offering
D4 - Beads
D5 - Black Belt
D6 - Coin Toss
D7 - Fake Mustache
D8 - Gem Box
D9 - Dragon Horn
DA - Merit Award
DB - Momento Ring
DC - Safety Bit
DD - Relic Ring
DE - Moogle Charm
DF - Charm Bangle
E0 - Marvel Shoes
E1 - Back Gaurd
E2 - Gale Hairpin
E3 - Sniper Sight
E4 - Exp.Egg
E5 - Tintinabar
E6 - Sprint Shoes`
)

var (
	ItemsByName = make(map[string]string)
	ItemsByID   = map[string]string{
		"00": "Dirk",
		"01": "MithrilKnife",
		"02": "Guardian",
		"03": "Air Lancet",
		"04": "Thief Knife",
		"05": "Assassin",
		"06": "Man Eater",
		"07": "Sword Breaker",
		"08": "Graedus",
		"09": "Valiant Knife",
		"0A": "Mithril Blade",
		"0B": "Regal Cutlass",
		"0C": "Rune Edge",
		"0D": "Flame Sabre",
		"0E": "Blizzard",
		"0F": "Thunder Blade",
		"10": "Epee",
		"11": "Break Blade",
		"12": "Drainer",
		"13": "Enhancer",
		"14": "Crystal",
		"15": "Falchion",
		"16": "Soul Sabre",
		"17": "Ogre Nix",
		"18": "Excalibur",
		"19": "Scimiter",
		"1A": "Illumina",
		"1B": "Ragnarok",
		"1C": "Atma Weapon",
		"1D": "Mithril Pike",
		"1E": "Trident",
		"1F": "Stout Spear",
		"20": "Partisan",
		"21": "Pearl Lance",
		"22": "Gold Lance",
		"23": "Aura Lance",
		"24": "Imp Halberd",
		"25": "Imperial",
		"26": "Kodachi",
		"27": "Blossom",
		"28": "Hardened",
		"29": "Striker",
		"2A": "Stunner",
		"2B": "Ashura",
		"2C": "Kotetsu",
		"2D": "Forged",
		"2E": "Tempest",
		"2F": "Murasame",
		"30": "Aura",
		"31": "Strato",
		"32": "Sky Render",
		"33": "Heal Rod",
		"34": "Mithril Rod",
		"35": "Fire Rod",
		"36": "Ice Rod",
		"37": "Thunder Rod",
		"38": "Poison Rod",
		"39": "Pearl Rod",
		"3A": "Gravity Rod",
		"3B": "Punisher",
		"3C": "Magus Rod",
		"3D": "Chocobo Brush",
		"3E": "DaVinci Brush",
		"3F": "Magical Brush",
		"40": "Rainbow Brush",
		"41": "Shuriken",
		"42": "Ninja Star",
		"43": "Tack Star",
		"44": "Flail",
		"45": "Full Moon",
		"46": "Morning Star",
		"47": "Boomerang",
		"48": "Rising Sun",
		"49": "Hawk Eye",
		"4A": "Bone Club",
		"4B": "Sniper",
		"4C": "Wing Edge",
		"4D": "Cards",
		"4E": "Darts",
		"4F": "Doom Darts",
		"50": "Trump",
		"51": "Dice",
		"52": "Fixed Dice",
		"53": "Metal Knuckle",
		"54": "Mithril Claw",
		"55": "Kaiser",
		"56": "Poison Claw",
		"57": "Fire Knuckle",
		"58": "Dragon Claw",
		"59": "Tiger Fangs",
		"5A": "Buckler",
		"5B": "Heavy Shield",
		"5C": "Mithril Shield",
		"5D": "Gold Shield",
		"5E": "Aegis Shield",
		"5F": "Diamond Shield",
		"60": "Flame Shield",
		"61": "Ice Shield",
		"62": "Thunder Shield",
		"63": "Crystal Shield",
		"64": "Genji Shield",
		"65": "Tortoise Shield",
		"66": "Cursed Shield",
		"67": "Paladin Shield",
		"68": "Force Shield",
		"69": "Leather Hat",
		"6A": "Hair Band",
		"6B": "Plumed Hat",
		"6C": "Beret",
		"6D": "Magus Hat",
		"6E": "Bandana",
		"6F": "Iron Helmet",
		"70": "Coronet",
		"71": "Bard's Hat",
		"72": "Green Beret",
		"73": "Head Band",
		"74": "Mithril Helm",
		"75": "Tiara",
		"76": "Gold Helmet",
		"77": "Tiger Mask",
		"78": "Red Hat",
		"79": "Mystery Veil",
		"7A": "Circlet",
		"7B": "Regal Crown",
		"7C": "Diamond Helm",
		"7D": "Dark Hood",
		"7E": "Crystal Helm",
		"7F": "Oath Veil",
		"80": "Cat Hood",
		"81": "Genji Helmet",
		"82": "Thornlet",
		"83": "Titanium",
		"84": "Leather Armor",
		"85": "Cotton Robe",
		"86": "Kung Fu Suit",
		"87": "Iron Armor",
		"88": "Silk Robe",
		"89": "Mithril Vest",
		"8A": "Ninja Gear",
		"8B": "White Dress",
		"8C": "Mithril Mail",
		"8D": "Gaia Gear",
		"8E": "Mirage Dress",
		"8F": "Gold Armor",
		"90": "Power Sash",
		"91": "Light Robe",
		"92": "Diamond Vest",
		"93": "Red Jacket",
		"94": "Force Armor",
		"95": "DiamondArmor",
		"96": "Dark Gear",
		"97": "Tao Robe",
		"98": "Crystal Mail",
		"99": "Czarina Gown",
		"9A": "Genji Armor",
		"9B": "Imp's Armor",
		"9C": "Minerva",
		"9D": "Tabby Suit",
		"9E": "Chocobo Suit",
		"9F": "Moogle Suit",
		"A0": "Nutkin Suit",
		"A1": "Behemeth Suit",
		"A2": "Snow Muffler",
		"B0": "Goggles",
		"B1": "Star Pendant",
		"B2": "Peace Ring",
		"B3": "Amulet",
		"B4": "White Cape",
		"B5": "Jewel Ring",
		"B6": "Fair Ring",
		"B7": "Barrier Ring",
		"B8": "Mithril Glove",
		"B9": "Guard Ring",
		"BA": "Running Shoes",
		"BB": "Wall Ring",
		"BC": "Cherub Down",
		"BD": "Cure Ring",
		"BE": "True Knight",
		"BF": "Dragoon Boots",
		"C0": "Zephyr Cape",
		"C1": "Czarina Ring",
		"C2": "Cursed Cing",
		"C3": "Earrings",
		"C4": "Atlas Armlet",
		"C5": "BlizzardRing",
		"C6": "Rage Ring",
		"C7": "Sneak Ring",
		"C8": "Pod Bracelet",
		"C9": "Hero Ring",
		"CA": "Ribbon",
		"CB": "Muscle Belt",
		"CC": "Crystal Orb",
		"CD": "Gold Hairpin",
		"CE": "Economizer",
		"CF": "Thief Glove",
		"D0": "Gauntlet",
		"D1": "Genji Glove",
		"D2": "Hyper Wrist",
		"D3": "Offering",
		"D4": "Beads",
		"D5": "Black Belt",
		"D6": "Coin Toss",
		"D7": "Fake Mustache",
		"D8": "Gem Box",
		"D9": "Dragon Horn",
		"DA": "Merit Award",
		"DB": "Momento Ring",
		"DC": "Safety Bit",
		"DD": "Relic Ring",
		"DE": "Moogle Charm",
		"DF": "Charm Bangle",
		"E0": "Marvel Shoes",
		"E1": "Back Gaurd",
		"E2": "Gale Hairpin",
		"E3": "Sniper Sight",
		"E4": "Exp.Egg",
		"E5": "Tintinabar",
		"E6": "Sprint Shoes",
		"A3": "NoiseBlaster",
		"A4": "Bio Blaster",
		"A5": "Flash",
		"A6": "Chain Saw",
		"A7": "Debilitator",
		"A8": "Drill",
		"A9": "Air Anchor",
		"AA": "AutoCrossbow",
		"AB": "Fire Skean",
		"AC": "Water Edge",
		"AD": "Bolt Edge",
		"AE": "Inviz Edge",
		"AF": "Shadow Edge",
		"E7": "Rename Card",
		"E8": "Tonic",
		"E9": "Potion",
		"EA": "X-Potion",
		"EB": "Tincture",
		"EC": "Ether",
		"ED": "X-Ether",
		"EE": "Elixir",
		"EF": "Megalixir",
		"F0": "Fenix Down",
		"F1": "Revivify",
		"F2": "Antidote",
		"F3": "Eydrop",
		"F4": "Soft",
		"F5": "Remedy",
		"F6": "Sleeping Bag",
		"F7": "Tent",
		"F8": "Green Cherry",
		"F9": "Magicite",
		"FA": "Super Ball",
		"FB": "Echo Screen",
		"FC": "Smoke Bomb",
		"FD": "Warp Stone",
		"FE": "Dried Meat",
		"FF": "Empty",
	}
)

func init() {
	for k, v := range ItemsByID {
		ItemsByName[v] = k
	}
}
