package file

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/kiamev/ffpr-save-cypher/rijndael"
	oj "github.com/virtuald/go-ordered-json"
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/save/config"
)

func LoadSave(game global.Game, fileName string) (data *save.Data, err error) {
	var b []byte
	data = save.New(game)
	if b, data.Trimmed, err = loadFile(game, fileName); err != nil {
		return
	}
	err = oj.Unmarshal(b, data)
	return
}

func loadFile(game global.Game, fromFile string) (out []byte, trimmed []byte, err error) {
	var (
		b []byte
	)
	if b, err = os.ReadFile(fromFile); err != nil {
		return
	}
	if len(b) < 10 {
		err = errors.New("unable to load file")
		return
	}
	// Format
	if b[0] == 239 && b[1] == 187 && b[2] == 191 {
		trimmed = []byte{239, 187, 191}
		b = b[3:]
	}
	for len(b)%4 != 0 {
		b = append(b, '=')
	}
	// Decode
	b, _ = base64.StdEncoding.DecodeString(string(b))
	if len(b) == 0 {
		err = errors.New("unable to load file")
		return
	}
	// Decrypt
	if b, err = rijndael.New().Decrypt(b); err != nil {
		return
	}

	// Flate
	zr := flate.NewReader(bytes.NewReader(b))
	defer func() { _ = zr.Close() }()
	out, err = io.ReadAll(zr)
	printFile(filepath.Join(config.Dir(game), "_loaded.json"), out)
	return
}
