package pr

type CharacterBase struct {
	ID     int
	Name   string
	HPBase int
	MPBase int
}

var (
	CharacterByName = make(map[string]*CharacterBase)
	CharacterByID   = make(map[int]*CharacterBase)
	CharacterBases  = []*CharacterBase{
		{
			ID:   1,
			Name: "Terra",
			//Character: Terra,
			HPBase: 63,
			MPBase: 138,
		},
		{
			ID:   2,
			Name: "Wedge",
			//Character: Wedge,
			HPBase: 30,
			MPBase: 0,
		},
		{
			ID:   3,
			Name: "Biggs",
			//Character: Biggs,
			HPBase: 30,
			MPBase: 0,
		},
		{
			ID:   5,
			Name: "Locke",
			//Character: Locke,
			HPBase: 48,
			MPBase: 7,
		},
		{
			ID:   16,
			Name: "Mog",
			//Character: Mog,
			HPBase: 39,
			MPBase: 16,
		},
		{
			ID:   17,
			Name: "Edgar",
			//Character: Edgar,
			HPBase: 49,
			MPBase: 6,
		},
		{
			ID:   18,
			Name: "Sabin",
			//Character: Sabin,
			HPBase: 76,
			MPBase: 3,
		},
		{
			ID:   19,
			Name: "Shadow",
			//Character: Shadow,
			HPBase: 51,
			MPBase: 6,
		},
		{
			ID:   20,
			Name: "Banon",
			//Character: Banon,
			HPBase: 44,
			MPBase: 15,
		},
		{
			ID:   22,
			Name: "Celes",
			//Character: Celes,
			HPBase: 44,
			MPBase: 15,
		},
		{
			ID:   23,
			Name: "Cyan",
			//Character: Cyan,
			HPBase: 53,
			MPBase: 5,
		},
		{
			ID:   24,
			Name: "Gau/Ghost",
			//Character: Gau,
			HPBase: 45,
			MPBase: 8,
		},
		{
			ID:   25,
			Name: "Setzer",
			//Character: Setzer,
			HPBase: 46,
			MPBase: 9,
		},
		{
			ID:   27,
			Name: "Strago",
			//Character: Strago,
			HPBase: 35,
			MPBase: 13,
		},
		{
			ID:   28,
			Name: "Relm",
			//Character: Relm,
			HPBase: 37,
			MPBase: 18,
		},
		{
			ID:   29,
			Name: "Leo",
			//Character: Leo,
			HPBase: 37,
			MPBase: 18,
		},
		{
			ID:   31,
			Name: "Umaro",
			//Character: Umaro,
			HPBase: 60,
			MPBase: 0,
		},
		{
			ID:   32,
			Name: "Gogo",
			//Character: Gogo,
			HPBase: 36,
			MPBase: 12,
		},
	}
)

func init() {
	for _, c := range CharacterBases {
		CharacterByName[c.Name] = c
		CharacterByID[c.ID] = c
	}
}
