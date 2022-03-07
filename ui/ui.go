package ui

import (
	"github.com/aarzilli/nucular"
)

var IsPR = false

type UI interface {
	Draw(w *nucular.Window)
	Refresh()
}
