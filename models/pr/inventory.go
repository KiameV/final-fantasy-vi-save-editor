package pr

type Inventory struct {
	Size             int
	Rows             []*Row
	ResetSortOrder   bool
	RemoveDuplicates bool
}

type Row struct {
	ItemID int `json:"contentId"`
	Count  int `json:"count"`
}

var inventory *Inventory
var importantInventory *Inventory

func GetInventory() *Inventory {
	if inventory == nil {
		inventory = &Inventory{Size: 255}
		inventory.Clear()
	}
	return inventory
}

func GetImportantInventory() *Inventory {
	if importantInventory == nil {
		importantInventory = &Inventory{Size: 100}
		importantInventory.Clear()
	}
	return importantInventory
}

func (i *Inventory) Clear() {
	// Clear existing memory
	for j := 0; j < len(i.Rows); j++ {
		i.Rows[j] = nil
	}

	i.Rows = make([]*Row, i.Size)
	for j := 0; j < i.Size; j++ {
		i.Rows[j] = &Row{ItemID: 0, Count: 0}
	}
}

func (i *Inventory) GetRows() []*Row {
	return i.Rows
}

func (i *Inventory) Reset() {
	for _, r := range i.Rows {
		r.ItemID = 0
		r.Count = 0
	}
}

func (i *Inventory) Set(index int, row Row) {
	if len(i.Rows) <= index {
		i.Rows = append(i.Rows, &row)
	} else {
		i.Rows[index] = &row
	}
}

func (i *Inventory) GetRowsForPrSave() []Row {
	rows := make([]Row, 0, len(i.Rows))
	for _, r := range i.Rows {
		if r.ItemID > 0 && r.ItemID <= 999 && r.Count > 0 && r.Count <= 99 {
			rows = append(rows, *r)
		}
	}
	return rows
}

func (i *Inventory) GetItemLookup() map[int]int {
	m := make(map[int]int)
	for _, r := range i.Rows {
		m[r.ItemID] = r.Count
	}
	return m
}

func (i *Inventory) AddNeeded(needed map[int]int) {
	for _, r := range i.Rows {
		if /*count*/ _, found := needed[r.ItemID]; found /*&& r.Count < count*/ {
			//r.Count = count
			delete(needed, r.ItemID)
		}
	}

	if len(needed) > 0 {
		var j int
		for j = len(i.Rows) - 1; j >= 0; j-- {
			if i.Rows[j] != nil && i.Rows[j].ItemID != 0 {
				j += 1
				break
			}
		}
		for id, count := range needed {
			if j >= len(i.Rows) {
				// TODO add error?
				return
			}
			i.Rows[j] = &Row{
				ItemID: id,
				Count:  count,
			}
		}
	}
}
