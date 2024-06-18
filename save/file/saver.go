package file

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kiamev/ffpr-save-cypher/rijndael"
	jo "gitlab.com/c0b/go-ordered-json"
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/save/config"
	"pixel-remastered-save-editor/save/savers"
	"pixel-remastered-save-editor/save/util"
)

func SaveSave(data *save.Data, slot int, fileName string) (err error) {
	var (
		toFile = filepath.Join(config.Dir(data.Game), fileName)
		temp   = filepath.Join(global.PWD, "temp")
		cmd    = exec.Command("cmd", "/C", "pr_io.exe", "obfuscateFile", toFile, temp)
		// needed   = make(map[int]int)
		slTarget = jo.NewOrderedMap()
	)
	cmd.Dir = strings.ReplaceAll(filepath.Join(global.PWD, "pr_io"), "\\", "/")

	if err = savers.Character(data); err == nil {
		if err = savers.Inventory(data, util.NormalOwnedItemList, util.NormalOwnedItemSortIdList, data.Save.Inventory); err == nil {
			if err = savers.Inventory(data, util.ImportantOwnedItemList, "", data.Save.ImportantInventory); err == nil {
				if err = savers.MapData(data); err == nil {
					if err = savers.Misc(data); err == nil {
						if err = savers.Transportation(data); err == nil {
							// err = savers.Party(data)
						}
					}
				}
			}
		}
	}
	iSlice := make([]interface{}, 0, len(data.Characters))
	for _, c := range data.Characters {
		if c != nil {
			var k []byte
			if k, err = c.MarshalJSON(); err != nil {
				return
			}
			s := string(k)
			iSlice = append(iSlice, s)
		}
	}

	if err = util.UnmarshalFrom(data.UserData, util.OwnedCharacterList, slTarget); err != nil {
		return
	}
	slTarget.Set(util.TargetKey, iSlice)
	if err = util.MarshalTo(data.UserData, util.OwnedCharacterList, slTarget); err != nil {
		return
	}

	if err = util.MarshalTo(data.Base, util.UserData, data.UserData); err != nil {
		return
	}

	if err = util.MarshalTo(data.Base, util.MapData, data.MapData); err != nil {
		return
	}

	if err = util.SetValue(data.Base, "id", slot); err != nil {
		return
	}

	var out []byte
	if out, err = json.Marshal(data.Base); err != nil {
		return
	}

	if _, err = os.Stat(temp); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create(temp); err != nil {
			return fmt.Errorf("failed to create save file %s: %v", toFile, err)
		}
	}
	defer func() { _ = os.Remove(temp) }()

	return saveFile(out, fileName, data.Trimmed)
}

func saveFile(data []byte, toFile string, trimmed []byte) (err error) {
	var (
		b  bytes.Buffer
		zw *flate.Writer
	)
	// printFile("save.file", data)
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
