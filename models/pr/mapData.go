package pr

type V3 struct {
	X, Y, Z float64
}

type GPS struct {
	TransportationID int
	MapID            int
	AreaID           int
	GpsID            int
	Width            int
	Height           int
}

type MapData struct {
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

var mapData *MapData

func GetMapData() *MapData {
	if mapData == nil {
		mapData = &MapData{}
	}
	return mapData
}
