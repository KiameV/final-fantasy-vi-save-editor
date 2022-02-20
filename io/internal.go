package io

import (
	"ffvi_editor/io/save"
	"github.com/sqweek/dialog"
	"os"
)

func createDialog(fileType save.SaveFileType) *dialog.FileBuilder {
	d := dialog.File()
	if appDir != "" {
		if dir, e1 := os.ReadFile(appDir + "/ff6editor.config"); e1 != nil {
			d = d.SetStartDir(".")
		} else {
			d = d.SetStartDir(string(dir))
		}
	} else {
		d = d.SetStartDir(".")
	}
	switch fileType {
	/*case offsets.Snes9xSaveState15, offsets.Snes9xSaveState16:
	d.Title("Select the Save State")
	d.Filters = []dialog.FileFilter{{
		Desc:       "Snes9x Save State File",
		Extensions: []string{"*.0*"},
	}}*/
	case save.SRMSlot1, save.SRMSlot2, save.SRMSlot3:
		d = d.Title("Select the Save File").Filter("SRM File", "srm")
	case save.ZnesSaveState:
		d = d.Title("Select the Save State").Filter("ZST File", "zst")
	}
	return d
}
