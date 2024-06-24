package save

import (
	oj "github.com/virtuald/go-ordered-json"
	"pixel-remastered-save-editor/global"
)

type (
	Data struct {
		ID                  int         `json:"id"`
		PictureData         string      `json:"pictureData"`
		UserDataInternal    string      `json:"userData"`
		ConfigDataInternal  string      `json:"configData"`
		DataStorageInternal string      `json:"dataStorage"`
		MapDataInternal     string      `json:"mapData"`
		TimeStamp           string      `json:"timeStamp"`
		PlayTime            float64     `json:"playTime"`
		ClearFlag           int         `json:"clearFlag"`
		IsCompleteFlag      int         `json:"isCompleteFlag"`
		Trimmed             []byte      `json:"-"`
		Game                global.Game `json:"-"`
	}
	ConfigData struct {
		ButtonType              int      `json:"buttonType"`
		BattleModeIndex         int      `json:"battleModeIndex"`
		BattleSpeedIndex        int      `json:"battleSpeedIndex"`
		BattleMessageSpeedIndex int      `json:"battleMessageSpeedIndex"`
		IsCursorMemory          bool     `json:"isCursorMemory"`
		IsKeepAutoBattle        bool     `json:"isKeepAutoBattle"`
		MessageSpeed            int      `json:"messageSpeed"`
		Brightness              *float64 `json:"brightness,omitempty"`
		MasterVolume            *float64 `json:"masterVolume,omitempty"`
		BgmVolume               *float64 `json:"bgmVolume,omitempty"`
		SeVolume                *float64 `json:"seVolume,omitempty"`
		IsLeftActionIcon        bool     `json:"isLeftActionIcon"`
		IsLeftVirtualPad        bool     `json:"isLeftVirtualPad"`
		IsLeftMenuCommand       bool     `json:"isLeftMenuCommand"`
		IsLeftBattleCommand     bool     `json:"isLeftBattleCommand"`
		ReequipIndex            int      `json:"reequipIndex"`
		AnalogModeIndex         int      `json:"analogModeIndex"`
		IsAutoDash              *int     `json:"isAutoDash,omitempty"`
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
	Position struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	}
)

func New(game global.Game) *Data {
	return &Data{
		Game: game,
	}
}

func (d *Data) Unpack() (ud *UserData, md *MapData, ds *DataStorage, err error) {
	if ud, err = d.UserData(); err == nil {
		if md, err = d.MapData(); err == nil {
			ds, err = d.DataStorage()
		}
	}
	return
}

func (d *Data) Pack(slot int, ud *UserData, md *MapData, ds *DataStorage) (data *Data, err error) {
	if err = d.SetUserData(ud); err == nil {
		if err = d.SetMapData(md); err == nil {
			if err = d.SetDataStorage(ds); err == nil {
				data = d
				data.ID = slot
			}
		}
	}
	return
}

func (d *Data) ConfigData() (v *ConfigData, err error) {
	return UnmarshalOne[ConfigData](d.ConfigDataInternal)
}

func (d *Data) SetConfigData(v *ConfigData) (err error) {
	d.ConfigDataInternal, err = MarshalOne[ConfigData](v)
	return
}

func (d *Data) DataStorage() (v *DataStorage, err error) {
	return UnmarshalOne[DataStorage](d.DataStorageInternal)
}

func (d *Data) SetDataStorage(v *DataStorage) (err error) {
	d.DataStorageInternal, err = MarshalOne[DataStorage](v)
	return
}

func (d *Data) MapData() (v *MapData, err error) {
	return UnmarshalOne[MapData](d.MapDataInternal)
}

func (d *Data) SetMapData(v *MapData) (err error) {
	d.MapDataInternal, err = MarshalOne[MapData](v)
	return
}

func (d *Data) UserData() (v *UserData, err error) {
	return UnmarshalOne[UserData](d.UserDataInternal)
}

func (d *Data) SetUserData(v *UserData) (err error) {
	d.UserDataInternal, err = MarshalOne[UserData](v)
	return
}

func UnmarshalOne[T any](from string) (to *T, err error) {
	to = new(T)
	err = oj.Unmarshal([]byte(from), to)
	return to, err
}

func UnmarshalMany[T any](from []string) (to []*T, err error) {
	to = make([]*T, len(from))
	for i, j := range from {
		if to[i], err = UnmarshalOne[T](j); err != nil {
			return
		}
	}
	return
}

func MarshalOne[T any](from *T) (string, error) {
	to, err := oj.Marshal(from)
	return string(to), err
}

func MarshalMany[T any](from []*T) (to []string, err error) {
	to = make([]string, len(from))
	for i, j := range from {
		if to[i], err = MarshalOne[T](j); err != nil {
			return
		}
	}
	return
}
