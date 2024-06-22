package core

type (
	Inventory struct {
		Rows             []*Row
		ResetSortOrder   bool
		RemoveDuplicates bool
	}
	Row struct {
		ContentID int `json:"contentId"`
		Count     int `json:"count"`
	}
)

func NewInventory() *Inventory {
	return &Inventory{}
}

func (i *Inventory) Reset() {
	for _, r := range i.Rows {
		r.ContentID = 0
		r.Count = 0
	}
}

func (i *Inventory) AppendEmptyRows() {
	for j := 0; j < 20; j++ {
		i.Rows = append(i.Rows, &Row{})
	}
}

func (i *Inventory) Set(index int, row Row) {
	if len(i.Rows) <= index {
		i.Rows = append(i.Rows, &row)
	} else {
		i.Rows[index] = &row
	}
}

func (i *Inventory) RowsForPrSave() []Row {
	rows := make([]Row, 0, len(i.Rows))
	for _, r := range i.Rows {
		if r.ContentID > 0 && r.ContentID <= 999 && r.Count > 0 && r.Count <= 99 {
			rows = append(rows, *r)
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

func (i *Inventory) Add(row *Row) {
	i.Rows = append(i.Rows, row)
}
