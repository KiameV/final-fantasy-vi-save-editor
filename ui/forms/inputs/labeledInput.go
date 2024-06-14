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
		hints *map[int]string
	}
	HintArgs struct {
		Align *fyne.TextAlign
		Hints *map[int]string
	}
)

func NewLabeledEntry(label string, entry fyne.CanvasObject) fyne.CanvasObject {
	return container.NewGridWithColumns(2, widget.NewLabel(label), entry)
}

func NewLabeledIntEntryWithHint(label string, entry *IntEntry, args ...HintArgs) fyne.CanvasObject {
	var a HintArgs
	if len(args) > 0 {
		a = args[0]
	}
	if a.Align == nil {
		a.Align = NewAlign(fyne.TextAlignTrailing)
	}
	h := newHinter(a)
	h.hint(entry.Text)
	entry.OnChanged = func(s string) {
		h.hint(s)
	}
	return container.NewVBox(NewLabeledEntry(label, entry), h.label)
}

func NewKeyValueIntEntryWithHint(key *IntEntry, value *IntEntry, args ...HintArgs) fyne.CanvasObject {
	var a HintArgs
	if len(args) > 0 {
		a = args[0]
	}
	if a.Align == nil {
		a.Align = NewAlign(fyne.TextAlignLeading)
	}
	h := newHinter(a)
	h.hint(key.Text)
	key.OnChanged = func(s string) {
		h.hint(s)
	}
	return container.NewVBox(container.NewGridWithColumns(2, key, value), h.label)
}

func newHinter(args HintArgs) *hinter {
	l := widget.NewLabel("")
	l.Alignment = *args.Align
	h := &hinter{
		label: l,
		hints: args.Hints,
	}
	return h
}

func (h *hinter) hint(s string) {
	if i, err := strconv.Atoi(s); err == nil {
		if i == 0 {
			h.label.SetText("[Empty]")
		} else {
			if m := h.hints; m != nil && len(*m) > 0 {
				if j, ok := (*m)[i]; ok {
					h.label.SetText(j)
				} else {
					h.label.SetText("-")
				}
			} else {
				if j, ok := pr.AllNormalItems[i]; ok {
					h.label.SetText(j)
				} else {
					h.label.SetText("-")
				}
			}
		}
	}
}

func NewAlign(a fyne.TextAlign) *fyne.TextAlign {
	return &a
}
