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
		EncountEnable            *bool  `json:"encountEnable,omitempty"`
		ViewType                 *int   `json:"viewType,omitempty"`
	}
	GpsData struct {
		MapID  int `json:"mapId"`
		AreaID int `json:"areaId"`
		GpsID  int `json:"gpsId"`
		Width  int `json:"width"`
		Height int `json:"height"`
	}
	PlayerEntity struct {
		Position  Position `json:"position"`
		Direction int      `json:"direction"`
	}
)

func (d *MapData) GpsData() (v *GpsData, err error) {
	return UnmarshalOne[GpsData](d.GpsDataInternal)
}

func (d *MapData) SetGpsData(v *GpsData) (err error) {
	d.GpsDataInternal, err = MarshalOne(v)
	return
}

func (d *MapData) PlayerEntity() (v *PlayerEntity, err error) {
	return UnmarshalOne[PlayerEntity](d.PlayerEntityInternal)
}

func (d *MapData) SetPlayerEntity(v *PlayerEntity) (err error) {
	d.PlayerEntityInternal, err = MarshalOne(v)
	return
}
