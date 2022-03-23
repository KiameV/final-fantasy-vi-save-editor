package ui

import (
	"github.com/aarzilli/nucular"
)

var DrawError error

type Behavior byte

const (
	Hide Behavior = iota
	SnesShow
	PrShow
	Show
)

type UI interface {
	Draw(w *nucular.Window)
	Refresh()
	Name() string
	Behavior() Behavior
}
