package pr

type CharacterBase struct {
	ID     int
	Name   string
	JobID  int
	HPBase int
	MPBase int
}

var (
	CharacterByName  = make(map[string]*CharacterBase)
	CharacterByID    = make(map[int]*CharacterBase)
	CharacterByJobID = make(map[int]*CharacterBase)
	CharacterBases   = []*CharacterBase{
		{
			ID:    1,
			Name:  "Terra",
			JobID: 1,
			//Character: Terra,
			HPBase: 63,
			MPBase: 138,
		},
		{
			ID:    2,
			Name:  "Wedge",
			JobID: 18,
			//Character: Wedge,
			HPBase: 30,
			MPBase: 0,
		},
		{
			ID:    3,
			Name:  "Biggs",
			JobID: 18,
			//Character: Biggs,
			HPBase: 30,
			MPBase: 0,
		},
		{
			ID:    5,
			Name:  "Locke",
			JobID: 2,
			//Character: Locke,
			HPBase: 48,
			MPBase: 7,
		},
		{
			ID:    16,
			Name:  "Mog",
			JobID: 11,
			//Character: Mog,
			HPBase: 39,
			MPBase: 16,
		},
		{
			ID:    17,
			Name:  "Edgar",
			JobID: 5,
			//Character: Edgar,
			HPBase: 49,
			MPBase: 6,
		},
		{
			ID:    18,
			Name:  "Sabin",
			JobID: 6,
			//Character: Sabin,
			HPBase: 76,
			MPBase: 3,
		},
		{
			ID:    19,
			Name:  "Shadow",
			JobID: 4,
			//Character: Shadow,
			HPBase: 51,
			MPBase: 6,
		},
		{
			ID:    20,
			Name:  "Banon",
			JobID: 15,
			//Character: Banon,
			HPBase: 44,
			MPBase: 15,
		},
		{
			ID:    22,
			Name:  "Celes",
			JobID: 7,
			//Character: Celes,
			HPBase: 44,
			MPBase: 15,
		},
		{
			ID:    23,
			Name:  "Cyan",
			JobID: 3,
			//Character: Cyan,
			HPBase: 53,
			MPBase: 5,
		},
		{
			ID:     24,
			Name:   "Ghost",
			JobID:  17,
			HPBase: 45,
			MPBase: 8,
		},
		{
			ID:    25,
			Name:  "Gau",
			JobID: 12,
			//Character: Gau,
			HPBase: 45,
			MPBase: 8,
		},
		{
			ID:    26,
			Name:  "Setzer",
			JobID: 10,
			//Character: Setzer,
			HPBase: 46,
			MPBase: 9,
		},
		{
			ID:    28,
			Name:  "Strago",
			JobID: 8,
			//Character: Strago,
			HPBase: 35,
			MPBase: 13,
		},
		{
			ID:    29,
			Name:  "Relm",
			JobID: 9,
			//Character: Relm,
			HPBase: 37,
			MPBase: 18,
		},
		{
			ID:    30,
			Name:  "Leo",
			JobID: 16,
			//Character: Leo,
			HPBase: 37,
			MPBase: 18,
		},
		{
			ID:    32,
			Name:  "Umaro",
			JobID: 14,
			//Character: Umaro,
			HPBase: 60,
			MPBase: 0,
		},
		{
			ID:    33,
			Name:  "Gogo",
			JobID: 13,
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
		CharacterByJobID[c.JobID] = c
	}
}

func GetCharacter(id, jobID int) (c *CharacterBase, ok bool) {
	if id <= 16 {
		c, ok = CharacterByID[id]
	} else {
		c, ok = CharacterByJobID[jobID]
	}
	return
}
