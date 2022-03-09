package models

import (
	"fmt"
	"strconv"
)

type Inventory struct {
	Rows []*Row
}

type Row struct {
	ItemID string
	Count  int
}

type PrRow struct {
	ContentId int `json:"contentId"`
	Count     int `json:"count"`
}

func newInventory() *Inventory {
	inv := &Inventory{
		Rows: make([]*Row, 255),
	}
	for i := 0; i < 255; i++ {
		inv.Rows[i] = &Row{ItemID: "FF", Count: 0}
	}
	return inv
}

var inventory *Inventory

func GetInventory() *Inventory {
	if inventory == nil {
		inventory = newInventory()
	}
	return inventory
}

func GetInventoryRows() []*Row {
	return GetInventory().Rows
}

func (i *Inventory) Reset() {
	for _, r := range i.Rows {
		r.ItemID = ""
		r.Count = 0
	}
}

func (i *Inventory) Set(index int, row PrRow) {
	id := fmt.Sprintf("%d", row.ContentId)
	if len(i.Rows) <= index {
		i.Rows = append(i.Rows, &Row{
			ItemID: id,
			Count:  row.Count,
		})
	} else {
		i.Rows[index].ItemID = id
		i.Rows[index].Count = row.Count
	}
}

func (i *Inventory) GetRowsForPrSave() []PrRow {
	rows := make([]PrRow, 0, len(i.Rows))
	for _, r := range i.Rows {
		if r.ItemID != "" && r.Count > 0 {
			id, _ := strconv.ParseInt(r.ItemID, 10, 8)
			if id <= 255 && id >= 0 {
				rows = append(rows, PrRow{
					ContentId: int(id),
					Count:     0,
				})
			}
		}
	}
	return rows
}
