package io

import (
	"github.com/sqweek/dialog"
	"os"
)

func createDialog(fileType SaveFileType) *dialog.FileBuilder {
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
	case SRMSlot1, SRMSlot2, SRMSlot3:
		d = d.Title("Select the Save File").Filter("SRM File", "srm")
	case ZnesSaveState:
		d = d.Title("Select the Save State").Filter("ZST File", "zs*").Filter("All", "*")
	}
	return d
}
