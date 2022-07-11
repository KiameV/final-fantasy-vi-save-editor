package pr

type CharacterBase struct {
	ID     int
	Name   string
	JobID  int
	HPBase int
	MPBase int
	IsNPC  bool
}

var (
	CharacterOffsetByName  = make(map[string]*CharacterBase)
	CharacterOffsetByID    = make(map[int]*CharacterBase)
	CharacterOffsetByJobID = make(map[int]*CharacterBase)
	CharacterOffsetBases   = []*CharacterBase{
		{
			ID:    1,
			Name:  "Terra",
			JobID: 1,
			//Character: Terra,
			HPBase: 29,
			MPBase: 11,
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
			ID:     6,
			Name:   "Moglin",
			JobID:  11,
			HPBase: 20,
			MPBase: 20,
			IsNPC:  true,
		},
		{
			ID:     7,
			Name:   "Mogret",
			JobID:  11,
			HPBase: 20,
			MPBase: 20,
			IsNPC:  true,
		},
		{
			ID:     8,
			Name:   "Moggie",
			JobID:  11,
			HPBase: 20,
			MPBase: 20,
			IsNPC:  true,
		},
		{
			ID:     9,
			Name:   "Molulu",
			JobID:  11,
			HPBase: 20,
			MPBase: 20,
			IsNPC:  true,
		},
		{
			ID:     10,
			Name:   "Moghan",
			JobID:  11,
			HPBase: 20,
			MPBase: 20,
			IsNPC:  true,
		},
		{
			ID:     11,
			Name:   "Moguel",
			JobID:  11,
			HPBase: 20,
			MPBase: 20,
			IsNPC:  true,
		},
		{
			ID:     12,
			Name:   "Mogsy",
			JobID:  11,
			HPBase: 20,
			MPBase: 20,
			IsNPC:  true,
		},
		{
			ID:     13,
			Name:   "Mogwin",
			JobID:  11,
			HPBase: 20,
			MPBase: 20,
			IsNPC:  true,
		},
		{
			ID:     14,
			Name:   "Mugmug",
			JobID:  11,
			HPBase: 20,
			MPBase: 20,
			IsNPC:  true,
		},
		{
			ID:     15,
			Name:   "Cosmog",
			JobID:  11,
			HPBase: 20,
			MPBase: 20,
			IsNPC:  true,
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
			HPBase: 58,
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
			Name:  "Cyan",
			JobID: 3,
			//Character: Cyan,
			HPBase: 53,
			MPBase: 5,
		},
		{
			ID:     23,
			Name:   "??????",
			JobID:  17,
			HPBase: 20,
			MPBase: 20,
			IsNPC:  true,
		},
		{
			ID:    24,
			Name:  "Gau",
			JobID: 12,
			//Character: Gau,
			HPBase: 45,
			MPBase: 10,
		},
		{
			ID:    25,
			Name:  "Celes",
			JobID: 7,
			//Character: Celes,
			HPBase: 44,
			MPBase: 15,
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
			ID:     27,
			Name:   "Maduin",
			JobID:  20,
			HPBase: 20,
			MPBase: 20,
			IsNPC:  true,
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
	for _, c := range CharacterOffsetBases {
		CharacterOffsetByName[c.Name] = c
		CharacterOffsetByID[c.ID] = c
		CharacterOffsetByJobID[c.JobID] = c
	}
}

func GetCharacterBaseOffset(id, jobID int) (c *CharacterBase, ok bool) {
	if id <= 16 {
		c, ok = CharacterOffsetByID[id]
	} else if id != 21 {
		c, ok = CharacterOffsetByJobID[jobID]
	}
	return
}
