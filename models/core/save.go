package core

type (
	Save struct {
		Characters         *Characters
		Party              *Party
		Inventory          *Inventory
		ImportantInventory *Inventory
		Transportations    *Transportations
		Map                *MapData
		Misc               *Misc
	}
)

func NewSave() *Save {
	return &Save{}
}
