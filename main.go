package main

import (
	"ffvi_editor/ui"
	"ffvi_editor/ui/forms/editors"
)

func main() {
	defer func() {
		_ = recover()
	}()
	gui := ui.New()
	editors.CreateTextBoxes()
	gui.Load()
	gui.Run()
	gui.App().Quit()
}
