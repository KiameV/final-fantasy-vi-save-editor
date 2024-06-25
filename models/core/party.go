package core

import (
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
)

type (
	Party struct {
		Members []*save.CorpsSlotInfo
	}
	Parties struct {
		slots   *save.CorpsSlots
		Parties [][]*save.CorpsSlotInfo
	}
)

func NewParty(game global.Game, ud *save.UserData) (p *Party, err error) {
	var cl *save.CorpsList
	if cl, err = ud.CorpsList(); err != nil {
		return
	}
	p = &Party{}
	p.Members, err = save.UnmarshalMany[save.CorpsSlotInfo](cl.Target)
	return
}

func (p *Party) ToSave(ud *save.UserData) (err error) {
	cl := &save.CorpsList{}
	if cl.Target, err = save.MarshalMany[save.CorpsSlotInfo](p.Members); err == nil {
		err = ud.SetCorpsList(cl)
	}
	return
}

func NewParties(game global.Game, ud *save.UserData) (p *Parties, err error) {
	p = &Parties{}
	if p.slots, err = ud.CorpsSlots(); err != nil {
		return
	}
	var v []*save.CorpsSlotsValues
	if v, err = p.slots.CorpsSlotsValues(); err != nil {
		return
	}
	p.Parties = make([][]*save.CorpsSlotInfo, len(v))
	for i, c := range v {
		if p.Parties[i], err = c.CorpsSlotInfo(); err != nil {
			return
		}
	}
	return
}

func (p *Parties) ToSave(ud *save.UserData) (err error) {
	v := make([]*save.CorpsSlotsValues, len(p.Parties))
	for i, j := range p.Parties {
		var csv save.CorpsSlotsValues
		csv.Target, err = save.MarshalMany[save.CorpsSlotInfo](j)
		v[i] = &csv
	}
	if err = p.slots.SetCorpsSlotsValues(v); err != nil {
		return
	}
	return ud.SetCorpsSlots(p.slots)
}
