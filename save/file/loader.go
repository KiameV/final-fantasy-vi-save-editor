package file

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"errors"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/kiamev/ffpr-save-cypher/rijndael"
	jo "gitlab.com/c0b/go-ordered-json"
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/save/loaders"
	"pixel-remastered-save-editor/save/util"
)

func LoadSave(game global.Game, fileName string) (data *save.Data, err error) {
	var (
		out []byte
		i   interface{}
		s   string
	)
	data = save.New(game)

	if out, data.Trimmed, err = loadFile(fileName); err != nil {
		return
	}
	s = string(out)

	if err = loadBase(data.Base, s); err != nil {
		return
	}

	if err = util.UnmarshalFrom(data.Base, util.UserData, data.UserData); err != nil {
		return
	}

	if err = util.UnmarshalFrom(data.Base, util.MapData, data.MapData); err != nil {
		return
	}

	if i, err = util.FromTarget(data.UserData, util.OwnedCharacterList); err != nil {
		return
	}

	for j, c := range i.([]interface{}) {
		if data.Characters[j] == nil {
			data.Characters[j] = jo.NewOrderedMap()
		}
		s = c.(string)
		if err = data.Characters[j].UnmarshalJSON([]byte(s)); err != nil {
			return
		}
	}

	data.Save = core.NewSave()
	if data.Save.Party, err = loaders.Party(data); err == nil {
		if data.Save.Characters, err = loaders.Characters(data, data.Save.Party); err == nil {
			if data.Save.Misc, err = loaders.Misc(data); err == nil {
				if data.Save.Inventory, err = loaders.Inventory(data, util.NormalOwnedItemList); err == nil {
					if data.Save.ImportantInventory, err = loaders.Inventory(data, util.ImportantOwnedItemList); err == nil {
						if data.Save.Map, err = loaders.MapData(data); err == nil {
							data.Save.Transportations, err = loaders.Transportation(data)
						}
					}
				}
			}
		}
	}
	return
}

func loadFile(fromFile string) (out []byte, trimmed []byte, err error) {
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
	// printFile(filepath.Join(config.Dir(game), "loaded.file"), out)
	return
}

func fixEscapeCharsForLoad(s string) string {
	var (
		sb    strings.Builder
		going = false
		count int
		found = make(map[int]string)
	)
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' {
			if !going {
				going = true
			}
			sb.WriteByte('\\')
			count++
		} else {
			if going {
				going = false
				sb.WriteByte('"')
				found[count] = sb.String()
				count = 0
				sb.Reset()
			}
		}
	}

	sorted := make([]int, 0, len(found))
	for k, _ := range found {
		sorted = append(sorted, k)
	}
	sort.Ints(sorted)

	for i := len(sorted) - 1; i >= 0; i-- {
		count = sorted[i]
		v, _ := found[count]
		sb.Reset()
		for j := 0; j < count; j++ {
			sb.WriteByte('~')
		}
		sb.WriteByte('"')
		found[count] = sb.String()
		s = strings.ReplaceAll(s, v, sb.String())
	}

	for i := len(sorted) - 1; i >= 0; i-- {
		count = sorted[i]
		v, _ := found[count]
		sb.Reset()
		half := (len(v) - 1) / 2
		for j := 0; j < half; j++ {
			sb.WriteByte('\\')
		}
		if sb.Len() == 0 {
			sb.WriteByte('\\')
		}
		sb.WriteByte('"')
		s = strings.ReplaceAll(s, v, sb.String())
	}
	return s
}

// {"keys":[1,2,3,4,5,6],"values":["{\\"contentId\\":113,\\"count\\":2}","{\\"contentId\\":214,\\"count\\":1}","{\\"contentId\\":268,\\"count\\":1}","{\\"contentId\\":233,\\"count\\":8}","{\\"contentId\\":200,\\"count\\":67}","{\\"contentId\\":315,\\"count\\":1}"]}
type equipment struct {
	Keys   []int    `json:"keys"`
	Values []string `json:"values"`
}

func loadBase(base *jo.OrderedMap, s string) (err error) {
	return base.UnmarshalJSON([]byte(s))
}

func uncheckAll(rages []*core.NameValueChecked) {
	for _, r := range rages {
		r.Checked = false
	}
}
