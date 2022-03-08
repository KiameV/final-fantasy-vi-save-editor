package ui

import (
	"github.com/aarzilli/nucular"
)

type UI interface {
	Draw(w *nucular.Window)
	Refresh()
	Name() string
	IsPRSupported() bool
}
