package pr

import (
	"errors"
	"ffvi_editor/global"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func (p *PR) Save(fileName string) (err error) {
	var (
		toFile = filepath.Join(global.Dir, fileName)
		temp   = filepath.Join(global.PWD, "temp")
		cmd    = exec.Command("python", "./io/python/io.py", "obfuscateFile", toFile, temp)
	)

	if _, err = os.Stat(temp); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create(temp); err != nil {
			return fmt.Errorf("failed to create save file %s: %v", toFile, err)
		}
	}
	defer os.Remove(temp)

	// TODO Remove
	//s := p.fixEscapeCharsForLoad(string(p.data))
	if err = ioutil.WriteFile(temp, p.data, 0644); err != nil {
		return fmt.Errorf("failed to create temp file %s: %v", toFile, err)
	}

	if _, err = os.Stat(toFile); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create(toFile); err != nil {
			return fmt.Errorf("failed to create save file %s: %v", toFile, err)
		}
	}

	_, err = cmd.Output()
	return
}
