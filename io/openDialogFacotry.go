package io

import (
	"ffvi_editor/global"
	"github.com/aarzilli/nucular"
	"github.com/sqweek/dialog"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

func OpenDirAndFileDialog(w *nucular.Window) (dir string, files []fs.FileInfo, err error) {
	d := dialog.Directory()
	if global.Dir != "" {
		d = d.SetStartDir(global.Dir)
	} else {
		d = d.SetStartDir(".")
	}
	d = d.Title("Select Save Directory")

	if dir, err = d.Browse(); err == nil {
		updateConfig(dir)
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

func OpenFileNoDialog(file string, fileType global.SaveFileType) (err error) {
	if err = CreateSnes(fileType).Load(file); err == nil {
		updateConfig(filepath.Dir(file))
	}
	return
}

func updateConfig(dir string) {
	if f, e1 := os.Create(filepath.Join(global.PWD, "ff6editor.config")); e1 == nil {
		_, _ = f.Write([]byte(dir))
	}
}
