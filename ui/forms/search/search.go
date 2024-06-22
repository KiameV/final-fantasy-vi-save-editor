package search

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type (
	TextBoxes struct {
		Abilities      *widget.RichText
		Items          *widget.RichText
		ImportantItems *widget.RichText
		Maps           *widget.RichText
	}
	Search struct {
		Abilities fyne.CanvasObject
		Items     fyne.CanvasObject
		Maps      fyne.CanvasObject
	}
)
