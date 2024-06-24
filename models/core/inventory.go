package core

import (
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
)

type (
	Inventory struct {
		Rows             []*save.OwnedItems
		ResetSortOrder   bool
		RemoveDuplicates bool
	}
	Row struct {
		ContentID int `json:"contentId"`
		Count     int `json:"count"`
	}
)

func NewInventory(game global.Game, items []*save.OwnedItems) *Inventory {
	inv := &Inventory{
		Rows: make([]*save.OwnedItems, 0, len(items)+20),
	}
	for _, i := range items {
		inv.Rows = append(inv.Rows, i)
	}
	for i := 0; i < 20; i++ {
		inv.Rows = append(inv.Rows, &save.OwnedItems{})
	}
	return inv
}

func (i *Inventory) Reset() {
	for _, r := range i.Rows {
		r.ContentID = 0
		r.Count = 0
	}
}

func (i *Inventory) Add(row *save.OwnedItems) {
	i.Rows = append(i.Rows, row)
}

func (i *Inventory) ToSave() ([]*save.OwnedItems, error) {
	inv := make([]*save.OwnedItems, 0, len(i.Rows))
	for _, r := range i.Rows {
		if r.ContentID > 0 {
			inv = append(inv, r)
		}
	}
	return inv, nil
}
