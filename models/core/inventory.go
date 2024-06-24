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
	return &Inventory{Rows: items}
}

func (i *Inventory) Reset() {
	for _, r := range i.Rows {
		r.ContentID = 0
		r.Count = 0
	}
}

func (i *Inventory) AddEmptyRow() {
	i.Rows = append(i.Rows, &save.OwnedItems{})
}

func (i *Inventory) Set(index int, row *save.OwnedItems) {
	if len(i.Rows) <= index {
		i.Rows = append(i.Rows, row)
	} else {
		i.Rows[index] = row
	}
}

func (i *Inventory) RowsForPrSave() []*save.OwnedItems {
	rows := make([]*save.OwnedItems, 0, len(i.Rows))
	for _, r := range i.Rows {
		if r.ContentID > 0 && r.ContentID <= 999 && r.Count > 0 && r.Count <= 99 {
			rows = append(rows, r)
		}
	}
	return rows
}

func (i *Inventory) ItemLookup() map[int]int {
	m := make(map[int]int)
	for _, r := range i.Rows {
		m[r.ContentID] = r.Count
	}
	return m
}

func (i *Inventory) Add(row *save.OwnedItems) {
	i.Rows = append(i.Rows, row)
}

func (i *Inventory) ToSave() ([]*save.OwnedItems, error) {
	return i.Rows, nil
}
