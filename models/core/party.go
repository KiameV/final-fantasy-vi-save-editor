package core

import (
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
)

type (
	Party struct {
		Members []*save.CorpsSlotInfo
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

func (p *Party) SetMember(slot int, characterID int) {
	p.Members[slot].CharacterID = characterID
}

func (p *Party) ToSave(ud *save.UserData) (err error) {
	cl := &save.CorpsList{}
	if cl.Target, err = save.MarshalMany[save.CorpsSlotInfo](p.Members); err == nil {
		err = ud.SetCorpsList(cl)
	}
	return
}
