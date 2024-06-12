package input

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewLabeledEntry(label string, entry fyne.CanvasObject) fyne.CanvasObject {
	return container.NewGridWithColumns(2, widget.NewLabel(label), entry)
}

func NewLabeledIntEntryWithHint(label string, entry *IntEntry, hints map[int]string) fyne.CanvasObject {
	hint := widget.NewLabel("")
	hint.Alignment = fyne.TextAlignTrailing
	entry.OnChanged = func(s string) {
		hint.SetText("adsfsdf")
		if i, err := strconv.Atoi(entry.Text); err == nil {
			if h, ok := hints[i]; ok {
				hint.SetText(h)
			}
		}
	}
	return container.NewVBox(NewLabeledEntry(label, entry), hint)
}
