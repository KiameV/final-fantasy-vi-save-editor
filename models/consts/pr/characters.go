package pr

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
		"Wedge",
		"Biggs",
		"Leo",
		"Banon",
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

func IsMainCharacter(n string) bool {
	return !(n == "Wedge" || n == "Biggs" || n == "Leo" || n == "Banon")
}
