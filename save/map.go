package save

type (
	MapData struct {
		MapID                    int    `json:"mapId"`
		PointIn                  int    `json:"pointIn"`
		TransportationID         int    `json:"transportationId"`
		CarryingHoverShip        bool   `json:"carryingHoverShip"`
		PlayerEntityInternal     string `json:"playerEntity"`
		CompanionEntity          string `json:"companionEntity"`
		GpsDataInternal          string `json:"gpsData"`
		MoveCount                int    `json:"moveCount"`
		SubtractSteps            int    `json:"subtractSteps"`
		TelepoCacheData          string `json:"telepoCacheData"`
		PlayableCharacterCorpsID int    `json:"playableCharacterCorpsId"`
		TimerData                string `json:"timerData"`
	}
	GpsData struct {
		MapID  int `json:"mapId"`
		AreaID int `json:"areaId"`
		GpsID  int `json:"gpsId"`
		Width  int `json:"width"`
		Height int `json:"height"`
	}
	PlayerEntity struct {
		PositionInternal string `json:"position"`
		Direction        int    `json:"direction"`
	}
	Position struct {
		X int `json:"x"`
		Y int `json:"y"`
		Z int `json:"z"`
	}
)

func (d *MapData) GpsData() (v *GpsData, err error) {
	return unmarshalOne[GpsData](d.GpsDataInternal)
}

func (d *MapData) SetGpsData(v *GpsData) (err error) {
	d.GpsDataInternal, err = marshalOne(v)
	return
}

func (d *MapData) PlayerEntity() (v *PlayerEntity, err error) {
	return unmarshalOne[PlayerEntity](d.PlayerEntityInternal)
}

func (d *MapData) SetPlayerEntity(v *PlayerEntity) (err error) {
	d.PlayerEntityInternal, err = marshalOne(v)
	return
}

func (d *PlayerEntity) Position() (v *Position, err error) {
	return unmarshalOne[Position](d.PositionInternal)
}

func (d *PlayerEntity) SetPosition(v *Position) (err error) {
	d.PositionInternal, err = marshalOne(v)
	return
}
