package core

import (
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
)

type (
	Save struct {
		Characters         *Characters
		Party              *Party
		Parties            *Parties
		Inventory          *Inventory
		ImportantInventory *Inventory
		Transportations    *Transportations
		Map                *MapData
		// Misc               *Misc
		Data *save.Data
	}
)

func NewSave(data *save.Data) (s *Save, err error) {
	var (
		ud *save.UserData
		md *save.MapData
		ds *save.DataStorage
		cl *save.OwnedCharacterList
		oi []*save.OwnedItems
		so []int
	)
	s = &Save{Data: data}
	if ud, md, ds, err = data.Unpack(); err != nil {
		return
	}
	if cl, err = ud.OwnedCharacterList(); err != nil {
		return
	}
	if s.Characters, err = NewCharacters(data.Game, cl); err != nil {
		return
	}
	if s.Party, err = NewParty(data.Game, ud); err != nil {
		return
	}
	if s.Parties, err = NewParties(data.Game, ud); err != nil {
		return
	}
	if oi, err = ud.NormalOwnedItems(); err != nil {
		return
	}
	if so, err = ud.NormalOwnedItemSortIDList(); err != nil {
		return
	}
	s.Inventory = NewInventory(data.Game, oi, so)
	if oi, err = ud.ImportantOwnedItems(); err != nil {
		return
	}
	s.ImportantInventory = NewInventory(data.Game, oi, nil)
	if s.Transportations, err = NewTransportations(data.Game, ds); err != nil {
		return
	}
	if s.Map, err = NewMapData(data.Game, md); err != nil {
		return
	}
	return s, err
}

func (s *Save) ToSave(slot int) (d *save.Data, err error) {
	if err = s.preSave(); err != nil {
		return
	}
	var (
		ud  *save.UserData
		md  *save.MapData
		ds  *save.DataStorage
		all = s.Characters.All()
		cl  = &save.OwnedCharacterList{Target: make([]string, len(all))}
		oi  []*save.OwnedItems
		so  []int
	)
	if ud, md, ds, err = s.Data.Unpack(); err != nil {
		return
	}

	for i, c := range all {
		if cl.Target[i], err = c.ToSave(); err != nil {
			return
		}
	}
	if err = ud.SetOwnedCharacterList(cl); err != nil {
		return
	}
	if err = s.Party.ToSave(ud); err != nil {
		return
	}
	if err = s.Parties.ToSave(ud); err != nil {
		return
	}
	if oi, so, err = s.Inventory.ToSave(); err != nil {
		return
	}
	if err = ud.SetNormalOwnedItems(oi); err != nil {
		return
	}
	if err = ud.SetNormalOwnedItemSortIDList(so); err != nil {
		return
	}
	if oi, _, err = s.ImportantInventory.ToSave(); err != nil {
		return
	}
	if err = ud.SetImportantOwnedItems(oi); err != nil {
		return
	}
	d, err = s.Data.Pack(slot, ud, md, ds)
	return
}

func (s *Save) preSave() (err error) {
	game := s.Game()

	inv := make(map[int]*save.OwnedItems)
	for _, i := range s.Inventory.Rows {
		inv[i.ContentID] = i
	}
	eq := make(map[int][]*save.Equipment)
	for _, c := range s.Characters.All() {
		for _, e := range c.Equipment.Values {
			sl, _ := eq[e.ContentID]
			sl = append(sl, e)
			eq[e.ContentID] = sl
		}
	}
	for k, v := range eq {
		needed := len(v)
		row, ok := inv[k]
		if ok {
			if row.Count < needed {
				row.Count = needed
			}
		} else {
			row = &save.OwnedItems{ContentID: k, Count: needed}
			inv[k] = row
			s.Inventory.Add(row)
		}
		for _, e := range v {
			e.Count = row.Count
		}
	}
	for _, c := range s.Characters.All() {
		if !c.Base.IsEnableCorps {
			continue
		}
		if game == global.One {
			err = s.preCharacterSaveFF1(c)
		} else if game == global.Two {
			err = s.preCharacterSaveFF2(c)
		} else if game == global.Three {
			err = s.preCharacterSaveFF3(c)
		} else if game == global.Four {
			err = s.preCharacterSaveFF4(c)
		}
		if err != nil {
			return
		}
	}
	return
}

func (s *Save) preCharacterSaveFF1(c *Character) (err error) {
	owned := make(map[int]bool)
	abilities := make(map[int]bool)
	for _, i := range c.AdditionAbilityIds {
		owned[i] = true
	}
	for _, a := range c.Abilities {
		abilities[a.ContentID] = true
	}
	for _, asd := range c.AbilitySlotData {
		for _, sd := range asd.SlotInfo.Values {
			if sd.AbilityID > 0 {
				sd.ContentID = sd.AbilityID + 208
			}
			if sd.ContentID > 0 {
				if _, ok := owned[sd.ContentID]; !ok {
					c.AdditionAbilityIds = append(c.AdditionAbilityIds, sd.ContentID)
				}
				if _, ok := abilities[sd.ContentID]; !ok {
					c.Abilities = append(c.Abilities, &save.Ability{AbilityID: sd.AbilityID, ContentID: sd.ContentID})
				}
			}
		}
	}
	for _, a := range c.Abilities {
		if a.AbilityID > 0 {
			if _, ok := owned[a.ContentID]; !ok {
				c.AdditionAbilityIds = append(c.AdditionAbilityIds, a.ContentID)
			}
		}
	}
	return
}

func (s *Save) preCharacterSaveFF2(c *Character) (err error) {
	for _, a := range c.Abilities {
		if a.ContentID == 0 && a.AbilityID > 0 {
			a.ContentID = a.AbilityID + 207
		}
	}
	return
}

func (s *Save) preCharacterSaveFF3(c *Character) (err error) {
	for _, a := range c.Abilities {
		if a.ContentID == 0 && a.AbilityID > 0 {
			a.ContentID = a.AbilityID + 201
		}
	}
	return
}

func (s *Save) preCharacterSaveFF4(c *Character) (err error) {
	for _, a := range c.Abilities {
		if a.ContentID == 0 && a.AbilityID > 0 {
			a.ContentID = a.AbilityID + 310
		}
	}
	return
}

func (s *Save) Game() global.Game {
	return s.Data.Game
}
