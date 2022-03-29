package io

import (
	"ffvi_editor/global"
	"github.com/sqweek/dialog"
)

func createDialog(fileType global.SaveFileType) *dialog.FileBuilder {
	d := dialog.File()
	dir := GetConfig().SaveDir
	if dir == "" {
		d = d.SetStartDir(".")
	} else {
		d = d.SetStartDir(dir)
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

func createInvDialog() *dialog.FileBuilder {
	d := dialog.File()
	d = d.SetStartDir(".")
	d = d.Title("Select the Inventory Save File").Filter("FF6INV File", "ff6inv")
	return d
}

func createDialogInv() *dialog.FileBuilder {
	d := dialog.File()
	d = d.SetStartDir(".")
	d = d.Title("Select the Inventory Save File").Filter("FF6INV File", "ff6inv")
	return d
}
