package inventory

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/finder"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Inventory struct {
		widget.BaseWidget
		items     []*inputs.IdCountEntry
		resetSort binding.Bool
		search    *inputs.Search
	}
)

func NewCore(inv *core.Inventory, finder finder.Find, search *inputs.Search) *Inventory {
	e := &Inventory{
		items:     make([]*inputs.IdCountEntry, len(inv.Rows)),
		resetSort: binding.BindBool(&inv.ResetSortOrder),
		search:    search,
	}
	e.ExtendBaseWidget(e)

	for i, item := range inv.Rows {
		e.items[i] = inputs.NewIdCountEntryWithDataWithHint(&item.ContentID, &item.Count, finder)
	}
	return e
}

func (e *Inventory) CreateRenderer() fyne.WidgetRenderer {
	const columns = 6
	id1 := widget.NewLabel("Item ID")
	id1.Alignment = fyne.TextAlignCenter
	id2 := widget.NewLabel("Item ID")
	id2.Alignment = fyne.TextAlignCenter
	c1 := widget.NewLabel("Count")
	c1.Alignment = fyne.TextAlignCenter
	c2 := widget.NewLabel("Count")
	c2.Alignment = fyne.TextAlignCenter
	reset := widget.NewCheckWithData("Reset Sort", e.resetSort)

	headers := container.NewPadded(container.NewGridWithColumns(columns, reset, id1, c1, container.NewStack(), id2, c2))

	l := len(e.items)
	rows := container.NewVBox()
	for i := 0; i < l; i += 2 {
		entry := e.items[i]
		g := container.NewGridWithColumns(columns, entry.Label, entry.ID, entry.Count)
		if j := i + 1; j < l {
			entry = e.items[j]
			g.Add(entry.Label)
			g.Add(entry.ID)
			g.Add(entry.Count)
		}
		rows.Add(container.NewPadded(g))
	}
	return widget.NewSimpleRenderer(
		container.NewBorder(nil, nil,
			// left
			container.NewBorder(
				headers, nil, nil, nil,
				container.NewVScroll(rows)),
			// right
			container.NewGridWithColumns(2,
				e.search.Fields(),
				e.search.Filter()),
		))
}
