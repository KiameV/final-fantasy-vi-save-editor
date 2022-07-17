package io

import (
	"ffvi_editor/global"
	"github.com/aarzilli/nucular"
	"github.com/ncruces/zenity"
	"io/fs"
	"io/ioutil"
	"path/filepath"
)

func OpenDirAndFileDialog(w *nucular.Window) (dir string, files []fs.FileInfo, err error) {
	dir = GetConfig().SaveDir
	if dir == "" {
		dir = "."
	}
	dir, err = zenity.SelectFile(
		zenity.Title("Select Save Directory"),
		zenity.Directory(),
		zenity.Filename(dir))
	if err == nil {
		GetConfig().SaveDir = dir
		SaveConfig()
		files, err = ioutil.ReadDir(dir)
	}
	return
}

func OpenFileDialog(w *nucular.Window, fileType global.SaveFileType) (fileName string, err error) {
	if fileName, err = createDialog(fileType); err != nil {
		return
	}
	err = OpenFileNoDialog(fileName, fileType)
	return
}

func OpenInvFileDialog(w *nucular.Window) (data []byte, err error) {
	var fn string
	if fn, err = openInvDialog(); err != nil {
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
