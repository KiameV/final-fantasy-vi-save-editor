package core

import (
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/models"
	"pixel-remastered-save-editor/save"
)

type (
	Save struct {
		Characters         *Characters
		Party              *Party
		Inventory          *Inventory
		ImportantInventory *Inventory
		Transportations    *Transportations
		Map                *MapData
		Misc               *Misc
		Data               *save.Data
	}
)

func NewSave(data *save.Data) (*Save, error) {
	var (
		ud, md, cd, ds, err = data.Unpack()
		s                   = &Save{Data: data}
	)
	if err == nil {
		// if s.Characters, err = ud.OwnedCharacterList(); err == nil {
		ud = ud
		md = md
		cd = cd
		ds = ds
		// }
	}
	return s, err
}

func (s *Save) ToSave() *save.Data {
	// TODO
	return s.Data
}

func (s *Save) preSave() (err error) {
	inv := make(map[int]*Row)
	for _, i := range s.Inventory.Rows {
		inv[i.ContentID] = i
	}
	eq := make(map[int][]*models.IdCount)
	for _, c := range s.Characters.All() {
		for _, e := range c.Equipment {
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
			row = &Row{ContentID: k, Count: needed}
			inv[k] = row
			s.Inventory.Add(row)
		}
		for _, e := range v {
			e.Count = row.Count
		}
	}
	return
}

func (s *Save) preCharacterSave(c *Character) (err error) {
	if s.Game() == global.One {
		owned := make(map[int]bool)
		abilities := make(map[int]bool)
		for _, i := range c.OwnedAbilityIds {
			owned[i] = true
		}
		for _, a := range c.Abilities {
			abilities[a.ContentID] = true
		}
		for i := 0; i < len(c.FF1.TrainedAbilities); i++ {
			for j := 0; j < len(c.FF1.TrainedAbilities[i]); j++ {
				t := c.FF1.TrainedAbilities[i][j]
				if t.ContentID == 0 && t.ID > 0 {
					t.ContentID = t.ID + 208
				}
				if t.ContentID > 0 {
					if _, ok := owned[t.ContentID]; !ok {
						c.OwnedAbilityIds = append(c.OwnedAbilityIds, t.ContentID)
					}
					if _, ok := abilities[t.ContentID]; !ok {
						c.Abilities = append(c.Abilities, &Ability{ID: t.ID, ContentID: t.ContentID})
					}
				}
			}
		}
		for _, a := range c.Abilities {
			if a.ID > 0 {
				if _, ok := owned[a.ContentID]; !ok {
					c.OwnedAbilityIds = append(c.OwnedAbilityIds, a.ContentID)
				}
			}
		}
	} else if s.Game() == global.Two {
		for _, a := range c.Abilities {
			if a.ContentID == 0 && a.ID > 0 {
				a.ContentID = a.ID + 207
			}
		}
	}
	return
}

func (s *Save) Game() global.Game {
	return s.Data.Game
}
