package io

import (
	"ffvi_editor/io/offsets"
	"github.com/aarzilli/nucular"
	"os"
)

func SaveFile(w *nucular.Window, fileType SaveFileType) (fileName string, err error) {
	if fileName, err = createDialog(fileType).Save(); err != nil {
		if err.Error() == "Cancelled" {
			w.Close()
			return "", nil
		}
		return
	}
	err = SaveFileNoDialog(fileName, fileType)
	return
}

func SaveFileNoDialog(file string, fileType SaveFileType) error {
	Create(fileType).Save()
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	if _, err = f.Write(offsets.Get()); err != nil {
		return err
	}
	return f.Close()
}
