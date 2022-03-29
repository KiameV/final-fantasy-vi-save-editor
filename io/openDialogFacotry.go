package io

import (
	"ffvi_editor/global"
	"github.com/aarzilli/nucular"
	"github.com/sqweek/dialog"
	"io/fs"
	"io/ioutil"
	"path/filepath"
)

func OpenDirAndFileDialog(w *nucular.Window) (dir string, files []fs.FileInfo, err error) {
	d := dialog.Directory()

	if GetConfig().SaveDir != "" {
		d = d.SetStartDir(GetConfig().SaveDir)
	} else {
		d = d.SetStartDir(".")
	}
	d = d.Title("Select Save Directory")

	if dir, err = d.Browse(); err == nil {
		GetConfig().SaveDir = dir
		SaveConfig()
		files, err = ioutil.ReadDir(dir)
	}
	return
}

func OpenFileDialog(w *nucular.Window, fileType global.SaveFileType) (fileName string, err error) {
	if fileName, err = createDialog(fileType).Load(); err != nil {
		return
	}

	err = OpenFileNoDialog(fileName, fileType)
	return
}

func OpenInvFileDialog(w *nucular.Window) (data []byte, err error) {
	var fn string
	if fn, err = createInvDialog().Load(); err != nil {
		return
	}
	return ioutil.ReadFile(fn)
}

func OpenFileNoDialog(file string, fileType global.SaveFileType) (err error) {
	if err = CreateSnes(fileType).Load(file); err == nil {
		GetConfig().SaveDir = filepath.Dir(file)
		SaveConfig()
	}
	return
}
