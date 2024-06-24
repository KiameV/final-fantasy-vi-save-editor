package core

import (
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
)

type (
	Transportations struct {
		Forms []*Transportation
	}
	Transportation struct {
		ID             int
		Enabled        bool
		ForcedEnabled  bool
		ForcedDisabled bool
		MapID          int
		Position       *save.Position
		Direction      int
		TimeStampTicks uint64
	}
)

func NewTransportations(game global.Game, ds *save.DataStorage) (t *Transportations, err error) {
	// TODO
	return &Transportations{
		Forms: []*Transportation{
			{},
		},
	}, nil
}

func (t *Transportations) Add(index int, form *Transportation) {
	t.Forms[index] = form
}
