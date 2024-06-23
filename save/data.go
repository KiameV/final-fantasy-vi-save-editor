package save

import (
	oj "github.com/virtuald/go-ordered-json"
	"pixel-remastered-save-editor/global"
)

type (
	Data struct {
		ID                  int         `json:"id"`
		PictureData         string      `json:"pictureData"`
		InternalUserData    string      `json:"userData"`
		InternalConfigData  string      `json:"configData"`
		InternalDataStorage string      `json:"dataStorage"`
		InternalMapData     string      `json:"mapData"`
		TimeStamp           string      `json:"timeStamp"`
		PlayTime            float64     `json:"playTime"`
		ClearFlag           int         `json:"clearFlag"`
		IsCompleteFlag      int         `json:"isCompleteFlag"`
		Trimmed             []byte      `json:"-"`
		Game                global.Game `json:"-"`
	}
	ConfigData struct {
		ButtonType              int     `json:"buttonType"`
		BattleModeIndex         int     `json:"battleModeIndex"`
		BattleSpeedIndex        int     `json:"battleSpeedIndex"`
		BattleMessageSpeedIndex int     `json:"battleMessageSpeedIndex"`
		IsCursorMemory          bool    `json:"isCursorMemory"`
		IsKeepAutoBattle        bool    `json:"isKeepAutoBattle"`
		MessageSpeed            int     `json:"messageSpeed"`
		Brightness              float64 `json:"brightness"`
		MasterVolume            float64 `json:"masterVolume"`
		BgmVolume               float64 `json:"bgmVolume"`
		SeVolume                float64 `json:"seVolume"`
		IsLeftActionIcon        bool    `json:"isLeftActionIcon"`
		IsLeftVirtualPad        bool    `json:"isLeftVirtualPad"`
		IsLeftMenuCommand       bool    `json:"isLeftMenuCommand"`
		IsLeftBattleCommand     bool    `json:"isLeftBattleCommand"`
		ReequipIndex            int     `json:"reequipIndex"`
		AnalogModeIndex         int     `json:"analogModeIndex"`
	}
	DataStorage struct {
		Scenario       []int `json:"scenario"`
		Treasure       []int `json:"treasure"`
		Global         []int `json:"global"`
		Area           []int `json:"area"`
		Map            []int `json:"map"`
		Selected       int   `json:"selected"`
		ItemSelected   int   `json:"itemSelected"`
		Transportation []any `json:"transportation"`
	}
)

func New(game global.Game) *Data {
	return &Data{
		Game: game,
	}
}

func (d *Data) Unpack() (ud *UserData, md *MapData, cd *ConfigData, ds *DataStorage, err error) {
	if ud, err = d.UserData(); err == nil {
		if md, err = d.MapData(); err == nil {
			if cd, err = d.ConfigData(); err == nil {
				ds, err = d.DataStorage()
			}
		}
	}
	return
}

func (d *Data) ConfigData() (v *ConfigData, err error) {
	return unmarshalOne[ConfigData](d.InternalConfigData)
}

func (d *Data) SetConfigData(v *ConfigData) (err error) {
	d.InternalConfigData, err = marshalOne[ConfigData](v)
	return
}

func (d *Data) DataStorage() (v *DataStorage, err error) {
	return unmarshalOne[DataStorage](d.InternalDataStorage)
}

func (d *Data) SetDataStorage(v *DataStorage) (err error) {
	d.InternalDataStorage, err = marshalOne[DataStorage](v)
	return
}

func (d *Data) MapData() (v *MapData, err error) {
	return unmarshalOne[MapData](d.InternalMapData)
}

func (d *Data) SetMapData(v *MapData) (err error) {
	d.InternalMapData, err = marshalOne[MapData](v)
	return
}

func (d *Data) UserData() (v *UserData, err error) {
	return unmarshalOne[UserData](d.InternalUserData)
}

func (d *Data) SetUserData(v *UserData) (err error) {
	d.InternalUserData, err = marshalOne[UserData](v)
	return
}

func unmarshalOne[T any](from string) (to *T, err error) {
	to = new(T)
	err = oj.Unmarshal([]byte(from), to)
	return to, err
}

func unmarshalMany[T any](from []string) (to []*T, err error) {
	to = make([]*T, len(from))
	for i, j := range from {
		if to[i], err = unmarshalOne[T](j); err != nil {
			return
		}
	}
	return
}

func marshalOne[T any](from *T) (string, error) {
	to, err := oj.Marshal(from)
	return string(to), err
}

func marshalMany[T any](from []*T) (to []string, err error) {
	to = make([]string, len(from))
	for i, j := range from {
		if to[i], err = marshalOne[T](j); err != nil {
			return
		}
	}
	return
}
