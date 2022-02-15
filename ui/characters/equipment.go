package characters

import (
	"ffvi_editor/ui"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/rect"
	"strings"
)

type equipmentUI struct {
	weaponID nucular.TextEditor
	shieldID nucular.TextEditor
	helmetID nucular.TextEditor
	armorID  nucular.TextEditor
	relic1ID nucular.TextEditor
	relic2ID nucular.TextEditor

	weaponShieldHelp1 nucular.TextEditor
	weaponShieldHelp2 nucular.TextEditor
	helmetArmorHelp1  nucular.TextEditor
	helmetArmorHelp2  nucular.TextEditor
	relicHelp1        nucular.TextEditor
	relicHelp2        nucular.TextEditor
}

func newEquipmentUI() widget {
	u := &equipmentUI{}

	u.initEquipmentTB(&u.weaponID)
	u.initEquipmentTB(&u.shieldID)
	u.initEquipmentTB(&u.helmetID)
	u.initEquipmentTB(&u.armorID)
	u.initEquipmentTB(&u.relic1ID)
	u.initEquipmentTB(&u.relic2ID)

	ui.InitReadOnlyText(&u.weaponShieldHelp1, weaponShieldText1)
	ui.InitReadOnlyText(&u.weaponShieldHelp2, weaponShieldText2)
	ui.InitReadOnlyText(&u.helmetArmorHelp1, helmetArmorText1)
	ui.InitReadOnlyText(&u.helmetArmorHelp2, helmetArmorText2)
	ui.InitReadOnlyText(&u.relicHelp1, relicText1)
	ui.InitReadOnlyText(&u.relicHelp2, relicText2)

	return u
}

func (u *equipmentUI) Draw(w *nucular.Window) {
	w.Row(610).SpaceBegin(18)
	u.drawPair(w, 0, "Weapon ID:", &u.weaponID, "Shield ID:", &u.shieldID, &u.weaponShieldHelp1, &u.weaponShieldHelp2)
	u.drawPair(w, 200, "Helmet ID:", &u.helmetID, "Armor ID:", &u.armorID, &u.helmetArmorHelp1, &u.helmetArmorHelp2)
	u.drawPair(w, 400, "Relic 1 ID:", &u.relic1ID, "Relic 2 ID:", &u.relic2ID, &u.relicHelp1, &u.relicHelp2)
}

func (u *equipmentUI) Update() {

}

func (u *equipmentUI) drawPair(w *nucular.Window, y int, label1 string, tb1 *nucular.TextEditor, label2 string, tb2 *nucular.TextEditor, helpTB1 *nucular.TextEditor, helpTB2 *nucular.TextEditor) {
	w.LayoutSpacePush(rect.Rect{
		X: 0,
		Y: y,
		W: 80,
		H: 22,
	})
	w.Label(label1, "LC")
	w.LayoutSpacePush(rect.Rect{
		X: 90,
		Y: y,
		W: 70,
		H: 22,
	})
	u.drawAndValidateInput(w, tb1)

	w.LayoutSpacePush(rect.Rect{
		X: 0,
		Y: y + 24,
		W: 80,
		H: 22,
	})
	w.Label(label2, "LC")
	w.LayoutSpacePush(rect.Rect{
		X: 90,
		Y: y + 24,
		W: 70,
		H: 22,
	})
	u.drawAndValidateInput(w, tb2)

	w.LayoutSpacePush(rect.Rect{
		X: 170,
		Y: y,
		W: 170,
		H: 190,
	})
	helpTB1.Edit(w)

	w.LayoutSpacePush(rect.Rect{
		X: 360,
		Y: y,
		W: 170,
		H: 190,
	})
	helpTB2.Edit(w)
}

func (u *equipmentUI) initEquipmentTB(tb *nucular.TextEditor) {
	tb.Flags = nucular.EditField
	tb.Maxlen = 2
	tb.SingleLine = true
}

func (u *equipmentUI) drawAndValidateInput(w *nucular.Window, tb *nucular.TextEditor) {
	if e := tb.Edit(w); e == nucular.EditActive {
		s := strings.ToUpper(string(tb.Buffer))
		r := FixHex2(s)
		tb.SelectAll()
		tb.DeleteSelection()
		tb.Text(r)
	}
}

func FixHex2(s string) []rune {
	l := len(s)
	if l == 1 {
		if isValidHexRune([]rune(s)[0]) {
			return []rune(s)
		}
		return make([]rune, 0)
	} else if l == 2 {
		r := []rune(s)
		if isValidHexRune(r[0]) {
			if isValidHexRune(r[1]) {
				return r
			}
			return []rune{r[0]}
		}
		if isValidHexRune(r[1]) {
			return []rune{r[1]}
		}
	}
	return []rune(s)
}

func isValidHexRune(r rune) bool {
	return (r >= '0' && r <= '9') || (r >= 'A' && r <= 'F')
}

const (
	weaponShieldText1 = `FF - Empty
Dirk
00 - Dirk
01 - MithrilKnife
02 - Guardian
03 - Air Lancet
04 - ThiefKnife
05 - Assassin
06 - Man Eater
07 - SwordBreaker
08 - Graedus
09 - ValiantKnife

Sword
0A - MithrilBlade
0B - RegalCutlass
0C - Rune Edge
0D - Flame Sabre
0E - Blizzard
0F - ThunderBlade
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
3D - Chocobo Brsh
3E - DaVinci Brsh
3F - Magical Brsh
40 - Rainbow Brsh

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
53 - MetalKnuckle
54 - Mithril Claw
55 - Kaiser
56 - Poison Claw
57 - Fire Knuckle
58 - Dragon Claw
59 - Tiger Fangs`
	weaponShieldText2 = `FF - Empty
Shield
5A - Buckler
5B - Heavy Shld
5C - Mithril Shld
5D - Gold Shld
5E - Aegis Shld
5F - Diamond Shld
60 - Flame Shld
61 - Ice Shld
62 - Thunder Shld
63 - Crystal Shld
64 - Genji Shld
65 - TortoiseShld
66 - Cursed Shld
67 - Paladin Shld
68 - Force Shld`
	helmetArmorText1 = `FF - Empty
Helmet
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
	helmetArmorText2 = `FF - Empty
Armor
84 - LeatherArmor
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
95 - DiamondArmor
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
A1 - BehemethSuit
A2 - Snow Muffler`
	relicText1 = `FF - Empty
Relic
B0 - Goggles
B1 - Star Pendant
B2 - Peace Ring
B3 - Amulet
B4 - White Cape
B5 - Jewel Ring
B6 - Fair Ring
B7 - Barrier Ring
B8 - MithrilGlove
B9 - Guard Ring
BA - RunningShoes
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
C5 - BlizzardRing
C6 - Rage Ring
C7 - Sneak Ring`
	relicText2 = `Relic (continued)
C8 - Pod Bracelet
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
D7 - FakeMustache
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
