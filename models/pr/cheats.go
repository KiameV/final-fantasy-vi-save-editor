package pr

type Cheats struct {
	Enabled          bool
	SetAllEncounters bool
	OpenedChestCount int
	ClearFlag        bool
	IsCompleteFlag   bool
}

var cheats *Cheats

func GetCheats() *Cheats {
	if cheats == nil {
		cheats = &Cheats{}
	}
	return cheats
}
