package main

import (
	"fmt"
	"time"

	"ffvi_editor/browser"
	"ffvi_editor/ui"
	"ffvi_editor/ui/forms/editors"
	"fyne.io/fyne/v2/dialog"
)

func main() {
	defer func() {
		_ = recover()
	}()
	gui := ui.New()
	editors.CreateTextBoxes()
	gui.Load()
	go func() {
		time.Sleep(1 * time.Second)
		if hasNewer, version, err := browser.CheckForUpdate(); err == nil && hasNewer {
			dialog.NewConfirm(
				"Update Available",
				fmt.Sprintf("A newer version of the editor is available: %s\nWould you like to download it now?", version), func(ok bool) {
					if ok {
						if e := browser.Update(version); e != nil {
							dialog.NewError(e, gui.Window())
						}
					}
				}, gui.Window()).Show()
		}
	}()
	gui.Run()
	gui.App().Quit()
}
