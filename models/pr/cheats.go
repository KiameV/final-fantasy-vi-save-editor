package pr

type Cheats struct {
	Enabled          bool
	OpenedChestCount int
	ClearFlag        bool
	IsCompleteFlag   bool
	PlayTime         float64
}

var cheats *Cheats

func GetCheats() *Cheats {
	if cheats == nil {
		cheats = &Cheats{}
	}
	return cheats
}
