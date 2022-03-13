package ui

import (
	"github.com/aarzilli/nucular"
)

var DrawError error

type UI interface {
	Draw(w *nucular.Window)
	Refresh()
	Name() string
	IsPRSupported() bool
}
