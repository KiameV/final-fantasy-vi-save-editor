package io

import (
	"ffvi_editor/global"
	"github.com/sqweek/dialog"
	"os"
	"path/filepath"
)

func init() {
	if global.Dir == "" {
		if dir, err := os.ReadFile(filepath.Join(global.PWD, "ff6editor.config")); err == nil {
			global.Dir = string(dir)
		}
	}
}

func createDialog(fileType global.SaveFileType) *dialog.FileBuilder {
	d := dialog.File()
	if global.Dir == "" {
		d = d.SetStartDir(".")
	} else {
		d = d.SetStartDir(global.Dir)
	}

	switch fileType {
	/*case offsets.Snes9xSaveState15, offsets.Snes9xSaveState16:
	d.Title("Select the Save State")
	d.Filters = []dialog.FileFilter{{
		Desc:       "Snes9x Save State File",
		Extensions: []string{"*.0*"},
	}}*/
	case global.SRMSlot1, global.SRMSlot2, global.SRMSlot3:
		d = d.Title("Select the Save File").Filter("SRM File", "srm")
	case global.ZnesSaveState:
		d = d.Title("Select the Save State").Filter("ZST File", "zs*").Filter("All", "*")
	}
	return d
}
