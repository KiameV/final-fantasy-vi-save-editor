package core

type (
	Misc struct {
		GP                     int
		Steps                  int
		NumberOfSaves          int
		SaveCountRollOver      int
		MapXAxis               int
		MapYAxis               int
		AirshipXAxis           int
		AirshipYAxis           int
		IsAirshipVisible       bool
		CursedShieldFightCount int
		EscapeCount            int
		BattleCount            int
		OpenedChestCount       int
		IsCompleteFlag         bool
		PlayTime               float64
		OwnedCrystals          [4]bool
	}
)

func NewMisc() *Misc {
	return &Misc{}
}
