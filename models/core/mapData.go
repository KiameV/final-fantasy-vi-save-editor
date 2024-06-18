package core

type (
	V3 struct {
		X, Y, Z float64
	}
	GPS struct {
		MapID  int
		AreaID int
		GpsID  int
		Width  int
		Height int
	}
	MapData struct {
		MapID             int
		PointIn           int
		TransportationID  int
		CarryingHoverShip bool
		Player            V3
		PlayerDirection   int
		Gps               GPS
		// MoveCount                int
		PlayableCharacterCorpsID int
	}
)

func NewMapData() *MapData {
	return &MapData{}
}
