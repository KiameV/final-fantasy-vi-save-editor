package consts

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

// 243 will fail
const (
	EmptyText = `0 - Invalid
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

	ImportantItemsText = `39 - Cider
40 - Old Clock Key
41 - Fish
42 - Fish
43 - Fish
44 - Fish
45 - Lump of Metal
46 - Lola's Letter
47 - Coral
48 - Books
49 - Emperor's Letter
50 - Rush-Rid
58 - Pendant`
)

var (
	ItemsByName          = make(map[string]int)
	ItemsByID            = make(map[int]string)
	ImportantItemsByName = make(map[string]int)
	ImportantItemsByID   = make(map[int]string)
)

func init() {
	loadItems(EmptyText, ItemsByName, ItemsByID)
	loadItems(ItemsText, ItemsByName, ItemsByID)
	loadItems(WeaponShieldText1, ItemsByName, ItemsByID)
	loadItems(WeaponShieldText2, ItemsByName, ItemsByID)
	loadItems(HelmetArmorText1, ItemsByName, ItemsByID)
	loadItems(HelmetArmorText2, ItemsByName, ItemsByID)
	loadItems(RelicText1, ItemsByName, ItemsByID)
	loadItems(RelicText2, ItemsByName, ItemsByID)
	loadItems(ImportantItemsText, ImportantItemsByName, ImportantItemsByID)
}

func loadItems(s string, byName map[string]int, byID map[int]string) {
	var in []int
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		sl := strings.Split(scanner.Text(), " - ")
		if len(sl) == 2 {
			if i, err := strconv.ParseInt(sl[0], 10, 32); err != nil {
				panic(sl[1])
			} else {
				byName[sl[1]] = int(i)
				byID[int(i)] = sl[1]
				in = append(in, int(i))
			}
		}
	}
	byID[93] = "Empty"
	byID[198] = "Empty"
	byID[199] = "Empty"
	byID[200] = "Empty"
	sort.Ints(in)
}
