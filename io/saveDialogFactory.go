package io

import (
	"ffvi_editor/global"
	"ffvi_editor/io/offsets"
	"github.com/aarzilli/nucular"
	"os"
)

func SaveFileSnes(w *nucular.Window) error {
	fn, err := createDialog(global.FileType).Save()
	if err != nil {
		return err
	}
	return SaveFileSnesNoDialog(fn)
}

func SaveFileSnesNoDialog(fileName string) error {
	CreateSnes(global.FileType).Save()
	f, err := os.OpenFile(global.FileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	if _, err = f.Write(offsets.Get()); err != nil {
		return err
	}
	global.FileName = fileName
	return f.Close()
}
