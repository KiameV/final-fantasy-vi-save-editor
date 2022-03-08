package pr

type CharacterBase struct {
	Name      string
	Character CharacterKey
	HPBase    int
	MPBase    int
}

var (
	CharBaseByName        = make(map[string]*CharacterBase)
	CharBaseByCharacter   = make(map[CharacterKey]*CharacterBase)
	CharBaseByCharacterID = make(map[int]*CharacterBase)
	CharacterBases        = []*CharacterBase{
		{
			Name:      "Terra",
			Character: Terra,
			HPBase:    63,
			MPBase:    138,
		},
		{
			Name:      "WedgeLeo",
			Character: Wedge,
			HPBase:    30,
			MPBase:    0,
		},
		{
			Name:      "Vicks",
			Character: Biggs,
			HPBase:    30,
			MPBase:    0,
		},
		{
			Name:      "Locke",
			Character: Locke,
			HPBase:    48,
			MPBase:    7,
		},
		{
			Name:      "Mog",
			Character: Mog,
			HPBase:    39,
			MPBase:    16,
		},
		{
			Name:      "Edgar",
			Character: Edgar,
			HPBase:    49,
			MPBase:    6,
		},
		{
			Name:      "Sabin",
			Character: Sabin,
			HPBase:    76,
			MPBase:    3,
		},
		{
			Name:      "Shadow",
			Character: Shadow,
			HPBase:    51,
			MPBase:    6,
		},
		{
			Name:      "Celes",
			Character: Celes,
			HPBase:    44,
			MPBase:    15,
		},
		{
			Name:      "Cyan",
			Character: Cyan,
			HPBase:    53,
			MPBase:    5,
		},
		{
			Name:      "Gau",
			Character: Gau,
			HPBase:    45,
			MPBase:    8,
		},
		{
			Name:      "Setzer",
			Character: Setzer,
			HPBase:    46,
			MPBase:    9,
		},
		{
			Name:      "Strago",
			Character: Strago,
			HPBase:    35,
			MPBase:    13,
		},
		{
			Name:      "Relm",
			Character: Relm,
			HPBase:    37,
			MPBase:    18,
		},
		{
			Name:      "Umaro",
			Character: Umaro,
			HPBase:    60,
			MPBase:    0,
		},
		{
			Name:      "Gogo",
			Character: Gogo,
			HPBase:    36,
			MPBase:    12,
		},
	}
)

func init() {
	for _, c := range CharacterBases {
		CharBaseByName[c.Name] = c
		CharBaseByCharacter[c.Character] = c
		CharBaseByCharacterID[int(c.Character)] = c
	}
}
