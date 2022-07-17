package io

import (
	"ffvi_editor/global"
	"github.com/ncruces/zenity"
)

func createDialog(fileType global.SaveFileType) (dir string, err error) {
	var (
		title  string
		filter zenity.FileFilters
	)
	dir = GetConfig().SaveDir
	if dir == "" {
		dir = "."
	}
	switch fileType {
	/*case offsets.Snes9xSaveState15, offsets.Snes9xSaveState16:
	d.Title("Select the Save State")
	d.Filters = []dialog.FileFilter{{
		Desc:       "Snes9x Save State File",
		Extensions: []string{"*.0*"},
	}}*/
	case global.SRMSlot1, global.SRMSlot2, global.SRMSlot3:
		title = "Select the Save File"
		filter = zenity.FileFilters{
			{
				Name:     "SRM File",
				Patterns: []string{"*.srm"},
			},
		}
	case global.ZnesSaveState:
		title = "Select the Save State"
		filter = zenity.FileFilters{
			{
				Name:     "ZST File",
				Patterns: []string{"*.zs*"},
			},
			{
				Name:     "All",
				Patterns: []string{"*.*"},
			},
		}
	}
	return zenity.SelectFile(
		zenity.Title(title),
		zenity.Filename(dir),
		filter)
}

func openInvDialog() (string, error) {
	return zenity.SelectFile(
		zenity.Title("Select the Inventory Save File"),
		zenity.Filename("."),
		zenity.FileFilter{
			Name:     "FF6INV File",
			Patterns: []string{getExt()},
		})
}

func saveDialogInv() (string, error) {
	return zenity.SelectFileSave(
		zenity.Title("Select the Inventory Save File"),
		zenity.Filename("."),
		zenity.FileFilter{
			Name:     "FF6INV File",
			Patterns: []string{getExt()},
		})
}

func getExt() string {
	return "*.ff6inv"
}
