package editors

import (
	"fmt"
	"strings"

	pr2 "ffvi_editor/io/pr"
	"ffvi_editor/models/pr"
	"ffvi_editor/ui/forms/inputs"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type (
	Inventory struct {
		widget.BaseWidget
		items   *fyne.Container
		search  *widget.Entry
		results *widget.TextGrid
	}
)

func NewInventory() *Inventory {
	inv := pr.GetInventory().Rows
	e := &Inventory{
		items:   container.NewGridWithRows(len(inv) + 1),
		search:  widget.NewEntry(),
		results: widget.NewTextGrid(),
	}
	e.ExtendBaseWidget(e)

	for _, item := range inv {
		e.items.Add(inputs.NewKeyValueIntEntryWithHint(
			inputs.NewIntEntryWithData(&item.ItemID),
			inputs.NewIntEntryWithData(&item.Count)))
	}

	e.search.OnChanged = func(s string) {
		s = strings.ToLower(s)
		if len(s) < 2 {
			e.results.SetText("")
			return
		}
		var sb strings.Builder
		if s != "" {
			for k, v := range pr2.AllNormalItems {
				if strings.Contains(strings.ToLower(v), s) {
					sb.WriteString(fmt.Sprintf("%d - %s\n", k, v))
				}
			}
		}
		e.results.SetText(sb.String())
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
				container.NewVScroll(container.NewPadded(e.items))),
			container.NewStack(itemsTextBox),
			container.NewStack(allEquipmentTextBox),
			container.NewBorder(
				inputs.NewLabeledEntry("Find By Name:", e.search), nil, nil, nil,
				container.NewVScroll(e.results))))
}
