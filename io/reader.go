package io

import (
	"github.com/aarzilli/nucular"
	"github.com/sqweek/dialog"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

func OpenDirAndFileDialog(w *nucular.Window) (dir string, files []fs.FileInfo, err error) {
	d := dialog.Directory()
	if Dir != "" {
		d = d.SetStartDir(Dir)
	} else {
		d = d.SetStartDir(".")
	}
	d = d.Title("Select Save Directory")

	if dir, err = d.Browse(); err != nil {
		if err.Error() == "Cancelled" {
			w.Close()
			return
		}
	}
	files, err = ioutil.ReadDir(dir)
	return
}

func OpenFileDialog(w *nucular.Window, fileType SaveFileType) (fileName string, err error) {
	if fileName, err = createDialog(fileType).Load(); err != nil {
		if err.Error() == "Cancelled" {
			w.Close()
			return "", nil
		}
		return
	}

	err = OpenFile(fileName, fileType)
	return
}

func OpenFile(file string, fileType SaveFileType) (err error) {
	if err = Create(fileType).Load(file); err == nil {
		if f, e1 := os.Create("./ff6editor.config"); e1 == nil {
			_, _ = f.Write([]byte(filepath.Dir(file)))
		}
	}
	return
}
