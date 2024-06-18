package inventory

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Inventory struct {
		widget.BaseWidget
		items *fyne.Container
	}
)

func NewCore(inv *core.Inventory) *Inventory {
	e := &Inventory{
		items: container.NewGridWithRows(len(inv.Rows) + 1),
	}
	e.ExtendBaseWidget(e)

	for _, item := range inv.Rows {
		e.items.Add(container.NewGridWithColumns(2,
			inputs.NewIntEntryWithData(&item.ItemID),
			inputs.NewIntEntryWithData(&item.Count)))
	}
	return e
}

func (e *Inventory) CreateRenderer() fyne.WidgetRenderer {
	l1 := widget.NewLabel("Item ID")
	l1.Alignment = fyne.TextAlignCenter
	l2 := widget.NewLabel("Count")
	l2.Alignment = fyne.TextAlignCenter
	return widget.NewSimpleRenderer(
		container.NewGridWithColumns(4,
			container.NewBorder(
				container.NewGridWithColumns(2, l1, l2), nil, nil, nil,
				container.NewVScroll(container.NewPadded(e.items)))))
}
