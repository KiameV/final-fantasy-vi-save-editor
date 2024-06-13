package editors

import (
	"ffvi_editor/models/pr"
	"ffvi_editor/ui/forms/inputs"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type (
	InventoryImportant struct {
		widget.BaseWidget
		items *fyne.Container
	}
)

func NewInventoryImportant() *InventoryImportant {
	inv := pr.GetImportantInventory().Rows
	e := &InventoryImportant{
		items: container.NewGridWithRows(len(inv) + 1),
	}
	e.ExtendBaseWidget(e)

	for _, item := range inv {
		e.items.Add(container.NewGridWithColumns(2,
			inputs.NewIntEntryWithData(&item.ItemID),
			inputs.NewIntEntryWithData(&item.Count)))
	}
	return e
}

func (e *InventoryImportant) CreateRenderer() fyne.WidgetRenderer {
	l1 := widget.NewLabel("Item ID")
	l1.Alignment = fyne.TextAlignCenter
	l2 := widget.NewLabel("Count")
	l2.Alignment = fyne.TextAlignCenter
	return widget.NewSimpleRenderer(
		container.NewGridWithColumns(4,
			container.NewBorder(
				container.NewGridWithColumns(2, l1, l2), nil, nil, nil,
				container.NewVScroll(container.NewPadded(e.items))),
			container.NewStack(importItemsTextBox)))
}
