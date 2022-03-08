package consts

var (
	Characters = []string{
		"Terra",
		"Locke",
		"Cyan",
		"Shadow",
		"Edgar",
		"Sabin",
		"Celes",
		"Strago",
		"Relm",
		"Setzer",
		"Mog",
		"Gau",
		"Gogo",
		"Umaro",
		"WedgeLeo",
		"Vicks",
	}
	CharacterLookup = make(map[string]int)
)

func init() {
	for i, c := range Characters {
		CharacterLookup[c] = i
	}
}

func GetCharacterIndex(name string) int {
	if i, b := CharacterLookup[name]; !b {
		panic("did not find character " + name)
	} else {
		return i
	}
}

/*
const (
	Terra    Character = 0x0
	Locke    Character = 0x1
	Cyan     Character = 0x2
	Shadow   Character = 0x3
	Edgar    Character = 0x4
	Sabin    Character = 0x5
	Celes    Character = 0x6
	Strago   Character = 0x7
	Relm     Character = 0x8
	Setzer   Character = 0x9
	Mog      Character = 0xA
	Gau      Character = 0xB
	Gogo     Character = 0xC
	Umaru    Character = 0xD
	WedgeLeo Character = 0xE
	Vicks    Character = 0xF
)
*/
