package save

import (
	jo "gitlab.com/c0b/go-ordered-json"
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/models/core"
)

type (
	Data struct {
		Base       *jo.OrderedMap
		UserData   *jo.OrderedMap
		MapData    *jo.OrderedMap
		Characters []*jo.OrderedMap
		Trimmed    []byte
		Save       *core.Save
		Game       global.Game
	}
)

func New(game global.Game) *Data {
	return &Data{
		Base:       jo.NewOrderedMap(),
		UserData:   jo.NewOrderedMap(),
		MapData:    jo.NewOrderedMap(),
		Characters: make([]*jo.OrderedMap, 40),
		Game:       game,
	}
}
