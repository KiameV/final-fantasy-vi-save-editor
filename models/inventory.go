package models

type Row struct {
	ItemID uint16
	Count  uint8
}

type Inventory struct {
	Rows []Row
}

func NewInventory() Inventory {
	return Inventory{
		Rows: make([]Row, 255),
	}
}
