package selections

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/core/ff1/consts"
	"pixel-remastered-save-editor/models/finder"
	"pixel-remastered-save-editor/ui/forms/editors/inventory"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Inventory struct {
		widget.BaseWidget
		save      *core.Save
		allSearch *inputs.Search
	}
)

func NewInventory(save *core.Save, allSearch *inputs.Search) *Inventory {
	s := &Inventory{
		save:      save,
		allSearch: allSearch,
	}
	s.ExtendBaseWidget(s)
	return s
}

func (s *Inventory) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		container.NewAppTabs(
			container.NewTabItem("Inventory", inventory.NewCore(s.save.Inventory, finder.Items, s.allSearch)),
			container.NewTabItem("Important", inventory.NewCore(s.save.ImportantInventory, finder.ImportantItems, inputs.NewSearch(consts.ImportantItems)))))
}
