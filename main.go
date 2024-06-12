package main

import (
	"ffvi_editor/ui"
)

func main() {
	defer func() {
		_ = recover()
	}()
	gui := ui.New()
	gui.Run()
	gui.App().Quit()
}
