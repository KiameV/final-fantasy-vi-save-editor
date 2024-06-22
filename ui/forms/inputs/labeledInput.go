package inputs

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/finder"
)

type (
	hinter struct {
		label *widget.Label
		hints finder.Find
	}
	HintArgs struct {
		Align *fyne.TextAlign
		Hints finder.Find
	}
	IdCountEntry struct {
		ID    *IntEntry
		Count *IntEntry
		Label *widget.Label
	}
	IdEntry struct {
		ID    *IntEntry
		Label *widget.Label
	}
)

func NewLabeledEntry(label string, entry fyne.CanvasObject) fyne.CanvasObject {
	return container.NewGridWithColumns(2, widget.NewLabel(label), entry)
}

func NewLabeledIntEntryWithHint(label string, i *int, args HintArgs) fyne.CanvasObject {
	if args.Align == nil {
		args.Align = NewAlign(fyne.TextAlignTrailing)
	}
	h := newHinter(args)
	entry := NewIntEntryWithData(i)
	h.hint(entry.Text)
	entry.OnChanged = func(s string) {
		h.hint(s)
	}
	return container.NewVBox(NewLabeledEntry(label, entry), h.label)
}

func NewIdCountEntryWithHint(key *IntEntry, value *IntEntry, find finder.Find) *IdCountEntry {
	label := widget.NewLabel("")
	key.OnChanged = func(s string) {
		hint(s, label, find)
	}
	return &IdCountEntry{
		ID:    key,
		Count: value,
		Label: label,
	}
}

func NewIdCountEntryWithDataWithHint(key *int, value *int, find finder.Find) *IdCountEntry {
	id := NewIntEntryWithData(key)
	count := NewIntEntryWithData(value)
	label := widget.NewLabel("")
	id.OnChanged = func(s string) {
		hint(s, label, find)
	}
	return &IdCountEntry{
		ID:    id,
		Count: count,
		Label: label,
	}
}

func NewIdEntryWithDataWithHint(key *int, find finder.Find) *IdEntry {
	id := NewIntEntryWithData(key)
	label := widget.NewLabel("")
	id.OnChanged = func(s string) {
		hint(s, label, find)
	}
	return &IdEntry{
		ID:    id,
		Label: label,
	}
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
	hint(s, h.label, h.hints)
}

func hint(s string, label *widget.Label, find finder.Find) {
	if find != nil {
		if i, err := strconv.Atoi(s); err == nil {
			if i == 0 {
				label.SetText("[Empty]")
			} else {
				if j, ok := find(i); ok {
					label.SetText(j)
				} else {
					label.SetText("-")
				}
			}
		}
	}
}

func NewAlign(a fyne.TextAlign) *fyne.TextAlign {
	return &a
}
