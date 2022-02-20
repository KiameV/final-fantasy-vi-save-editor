package io

import (
	"ffvi_editor/io/save"
	"os"
)

func SaveFile(fileType save.SaveFileType) (err error) {
	var file string
	if file, err = createDialog(fileType).Save(); err != nil {
		return
	}

	/*if fileType == offsets.Snes9xSaveState15 ||
		fileType == offsets.Snes9xSaveState16 {
		return openGzipFile(file)
	}*/
	return saveFile(file)
}

/*func openGzipFile(file string) error {
	archive.
	return nil
}*/

func saveFile(file string) error {
	// TODO Update save
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	if _, err = f.Write(save.Get()); err != nil {
		return err
	}
	return f.Close()
}
