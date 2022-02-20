package io

import (
	"ffvi_editor/io/offsets"
	"ffvi_editor/io/save"
	"os"
	"path/filepath"
)

var appDir string

func init() {
	appDir, _ = os.Getwd()
}

func OpenFile(fileType save.SaveFileType) (err error) {
	var file string
	if file, err = createDialog(fileType).Load(); err != nil {
		return
	}

	/*if fileType == offsets.Snes9xSaveState15 ||
		fileType == offsets.Snes9xSaveState16 {
		return openGzipFile(file)
	}*/
	return openFile(file, fileType)
}

/*func openGzipFile(file string) error {
	archive.
	return nil
}*/

func openFile(file string, fileType save.SaveFileType) (err error) {
	var data []byte
	data, err = os.ReadFile(file)
	if appDir != "" {
		if f, e1 := os.Create(appDir + "/ff6editor.config"); e1 == nil {
			_, _ = f.Write([]byte(filepath.Dir(file)))
		}
	}
	save.Set(data)
	offsets.Create(fileType).Load()
	return
}
