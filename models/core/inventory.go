package core

import (
	"cmp"
	"os"
	"slices"

	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
)

type (
	Inventory struct {
		Rows             []*save.OwnedItems
		order            []int
		ResetSortOrder   bool
		RemoveDuplicates bool
	}
	Row struct {
		ContentID int `json:"contentId"`
		Count     int `json:"count"`
	}
)

func NewInventory(game global.Game, items []*save.OwnedItems, order []int) *Inventory {
	inv := &Inventory{
		Rows:  make([]*save.OwnedItems, 0, len(items)+20),
		order: order,
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

func (i *Inventory) ToSave() (inv []*save.OwnedItems, order []int, err error) {
	inv = make([]*save.OwnedItems, 0, len(i.Rows))
	for _, r := range i.Rows {
		if r.ContentID > 0 {
			inv = append(inv, r)
		}
	}
	inv = i.debugToSave(inv)

	if len(i.order) > 0 && i.ResetSortOrder {
		order = make([]int, len(i.Rows))
		for j := 0; j < len(i.Rows); j++ {
			order[j] = j + 1
		}
	} else {
		order = i.order
	}
	return inv, order, nil
}

func (i *Inventory) debugToSave(inv []*save.OwnedItems) []*save.OwnedItems {
	if os.Getenv("PR_SORT_INV") == "true" {
		slices.SortFunc(inv, func(i, j *save.OwnedItems) int {
			return cmp.Compare(i.ContentID, j.ContentID)
		})
		for j, r := range inv {
			c := r.ContentID
			for c >= 100 {
				c -= 100
			}
			if c <= 0 {
				c = 50
			}
			inv[j].Count = c
		}
	}
	return inv
}
