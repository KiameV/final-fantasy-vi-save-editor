package snes

const (
	EmptyText = `FF - Empty
`
	ItemsText = `Miscellaneous
2 - Potion
3 - Hi-Potion
4 - X-Potion
5 - Ether
6 - Hi-Ether
7 - X-Ether
8 - Elixir
9 - Megalixir
10 - Fenix Down
11 - Holy Water
12 - Antidote
13 - Eye Drops
14 - Golden Needle
15 - Remedy
16 - Sleeping Bag
17 - Tent
18 - Green Cherry
19 - Magicite Shard
20 - Super Ball
21 - Echo Screen
22 - Smoke Bomb
23 - Teleport Stone
24 - Dried Meat
25 - Rename Card

Tool
31 - Noiseblaster
32 - Bioblaster
33 - Flash
34 - ChainSaw
35 - Debilitator
36 - Drill
37 - Air Anchor
38 - Auto Crossbow

Scroll
26 - Fire Scroll
27 - Water Scroll
28 - Lightning Scroll
29 - Invisibility Scroll
30 - Shadow Scroll`

	WeaponShieldText1 = `Dirk
94 - Dagger
95 - Mythril Knife
96 - Main Gauche
97 - Air Knife
98 - Thief's Knife
99 - Assassin's Dagger
100 - Man-Eater
101 - Swordbreaker
102 - Gladius
103 - Valiant Knife

Sword
104 - Mythril Sword
105 - Great Sword
106 - Rune Blade
107 - Flametongue
108 - Icebrand
109 - Thunder Blade
110 - Bastard Sword 
111 - Break Blade
112 - Blood Sword
113 - Enhancer
114 - Crystal Sword
115 - Falchion
116 - Soul Sabre
117 - Organyx
118 - Excalibur
119 - Zantetsuken
120 - Lightbringer
121 - Ragnarok
122 - Ultima Weapon

Lance
123 - Mythril Spear
124 - Trident
125 - Heavy Lance
126 - Partisan
127 - Holy Lance
128 - Golden Lance
129 - Radiant Lance
130 - Impartisan

Dirk
131 - Kunai
132 - Kodachi
133 - Sakura
134 - Sasuke
135 - Ichigeki
136 - Kagenui

Katana
137 - Ashura
138 - Kotetsu
139 - Kikuichimonji
140 - Kazekiri
141 - Masamane
142 - Masamune
143 - Murakumo
144 - Mutsunokami

Rod
145 - Heal Rod
146 - Mythril Rod
147 - Flame Rod
148 - Ice Rod
149 - Thunder Rod
150 - Poison Rod
151 - Holy Rod
152 - Gravity Rod
153 - Punisher
154 - Magus Rod

Brush
155 - Chocobo Brush
156 - Da Vinci Brush
157 - Magical Brush
158 - Rainbow Brush

Stars
159 - Shuriken
160 - Fuma Shuriken
161 - Pinwheel

Special
162 - Chain Flail
163 - Moonring Blade
164 - Morning Star
165 - Boomerang
166 - Rising Sun
167 - Hawkeye
168 - Bone Club
169 - Sniper
170 - Wing Edge

Gambler
171 - Cards
172 - Darts
173 - Death Tarot
174 - Viper Darts
175 - Dice
176 - Fixed Dice

Claw
177 - Metal Knuckles
178 - Mythril Claws
179 - Kaiser Knuckles
180 - Venom Claws
181 - Burning Fist
182 - Dragon Claws
183 - Tigerfangs`

	WeaponShieldText2 = `Shield
 - Buckler
 - Heavy Shield
 - Mithril Shield
 - Gold Shield
 - Aegis Shield
 - Diamond Shield
 - Flame Shield
 - Ice Shield
 - Thunder Shield
 - Crystal Shield
 - Genji Shield
 - Tortoise Shield
 - Cursed Shield
 - Paladin Shield
 - Force Shield`

	HelmetArmorText1 = `Helmet
 - Leather Hat
 - Hair Band
 - Plumed Hat
 - Beret
 - Magus Hat
 - Bandana
 - Iron Helmet
 - Coronet
 - Bard's Hat
 - Green Beret
 - Head Band
 - Mithril Helm
 - Tiara
 - Gold Helmet
 - Tiger Mask
 - Red Hat
 - Mystery Veil
 - Circlet
 - Regal Crown
 - Diamond Helm
 - Dark Hood
 - Crystal Helm
 - Oath Veil
 - Cat Hood
 - Genji Helmet
 - Thornlet
 - Titanium`

	HelmetArmorText2 = `Armor
 - Leather Armor
 - Cotton Robe
 - Kung Fu Suit
 - Iron Armor
 - Silk Robe
 - Mithril Vest
 - Ninja Gear
 - White Dress
 - Mithril Mail
 - Gaia Gear
 - Mirage Dress
 - Gold Armor
 - Power Sash
 - Light Robe
 - Diamond Vest
 - Red Jacket
 - Force Armor
 - Diamond Armor
 - Dark Gear
 - Tao Robe
 - Crystal Mail
 - Czarina Gown
 - Genji Armor
 - Imp's Armor
 - Minerva
 - Tabby Suit
 - Chocobo Suit
 - Moogle Suit
 - Nutkin Suit
 - Behemeth Suit
 - Snow Muffler`

	RelicText1 = `Relic
 - Goggles
 - Star Pendant
 - Peace Ring
 - Amulet
 - White Cape
 - Jewel Ring
 - Fair Ring
 - Barrier Ring
 - Mithril Glove
 - Guard Ring
 - Running Shoes
 - Wall Ring
 - Cherub Down
 - Cure Ring
 - True Knight
 - DragoonBoots
 - Zephyr Cape
 - Czarina Ring
 - Cursed Cing
 - Earrings
 - Atlas Armlet
 - Blizzard Ring
 - Rage Ring
 - Sneak Ring
`
	RelicText2Header = `Relic (continued)
`
	RelicText2 = `C8 - Pod Bracelet
 - Hero Ring
 - Ribbon
 - Muscle Belt
 - Crystal Orb
 - Gold Hairpin
 - Economizer
 - Thief Glove
 - Gauntlet
 - Genji Glove
 - Hyper Wrist
 - Offering
 - Beads
 - Black Belt
 - Coin Toss
 - Fake Mustache
 - Gem Box
 - Dragon Horn
 - Merit Award
 - Momento Ring
 - Safety Bit
 - Relic Ring
 - Moogle Charm
 - Charm Bangle
 - Marvel Shoes
 - Back Gaurd
 - Gale Hairpin
 - Sniper Sight
 - Exp.Egg
 - Tintinabar
 - Sprint Shoes`
)

var (
	ItemsByName = make(map[string]string)
	ItemsByID   = map[string]string{}
)

func init() {
	for k, v := range ItemsByID {
		ItemsByName[v] = k
	}
}
