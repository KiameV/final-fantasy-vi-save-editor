package file

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"

	"ffvi_editor/global"
	"github.com/kiamev/ffpr-save-cypher/rijndael"
)

func LoadFile(fromFile string, saveType global.SaveFileType) (out []byte, trimmed []byte, err error) {
	var (
		b []byte
	)
	if b, err = os.ReadFile(fromFile); err != nil {
		return
	}
	if saveType == global.PS {
		return b, nil, nil
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
	printFile("loaded.json", out)
	return
}

func SaveFile(data []byte, toFile string, trimmed []byte, saveType global.SaveFileType) (err error) {
	var (
		b  bytes.Buffer
		zw *flate.Writer
	)
	printFile("save.json", data)
	if saveType == global.PC {
		// Flate
		if zw, err = flate.NewWriter(&b, 6); err != nil {
			return
		}
		if _, err = zw.Write(data); err != nil {
			return
		}
		_ = zw.Flush()
		_ = zw.Close()

		// Encrypt
		if data, err = rijndael.New().Encrypt(b.Bytes()); err != nil {
			return
		}

		// Encode
		s := base64.StdEncoding.EncodeToString(data)

		// Format
		if len(trimmed) > 0 {
			data = make([]byte, 0, len(trimmed)+len(s))
			data = append(data, trimmed...)
			data = append(data, []byte(s)...)
		} else {
			data = []byte(s)
		}
	}
	// Write to file
	if _, err = os.Stat(toFile); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create(toFile); err != nil {
			return fmt.Errorf("failed to create save file %s: %v", toFile, err)
		}
	}
	if err = os.WriteFile(toFile, data, 0777); err != nil {
		return fmt.Errorf("failed to write save file %s: %v", toFile, err)
	}
	return
}

func printFile(name string, b []byte) {
	// s := strings.ReplaceAll(string(b), "\"", "\n\"")
	// s = strings.ReplaceAll(s, "\\", "")
	// os.WriteFile(name, []byte(s), 0755)
}
