package inputs

import (
	"strconv"

	"ffvi_editor/io/pr"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type (
	hinter struct {
		label *widget.Label
	}
)

func NewLabeledEntry(label string, entry fyne.CanvasObject) fyne.CanvasObject {
	return container.NewGridWithColumns(2, widget.NewLabel(label), entry)
}

func NewLabeledIntEntryWithHint(label string, entry *IntEntry) fyne.CanvasObject {
	h := newHinter(fyne.TextAlignLeading)
	h.hint(entry.Text)
	entry.OnChanged = func(s string) {
		h.hint(s)
	}
	return container.NewVBox(NewLabeledEntry(label, entry), h.label)
}

func NewKeyValueIntEntryWithHint(key *IntEntry, value *IntEntry) fyne.CanvasObject {
	h := newHinter(fyne.TextAlignLeading)
	h.hint(key.Text)
	key.OnChanged = func(s string) {
		h.hint(s)
	}
	return container.NewVBox(container.NewGridWithColumns(2, key, value), h.label)
}

func newHinter(alignment fyne.TextAlign) *hinter {
	l := widget.NewLabel("")
	l.Alignment = alignment
	return &hinter{label: l}
}

func (h *hinter) hint(s string) {
	if i, err := strconv.Atoi(s); err == nil {
		if i == 0 {
			h.label.SetText("[Empty]")
		} else if j, ok := pr.AllNormalItems[i]; ok {
			h.label.SetText(j)
		} else {
			h.label.SetText("-")
		}
	}
}
