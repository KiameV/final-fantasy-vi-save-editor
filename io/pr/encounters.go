package pr

import (
	"bytes"
	"encoding/json"
	"errors"
	"ffvi_editor/global"
	"fmt"
	"gitlab.com/c0b/go-ordered-json"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func (p *PR) CompleteAllEncounters(saveDir string) (err error) {
	const file = "dp3fS2vqP7GDj8eF72YKqbT7FIAF=e7Shy2CsTITm2E="
	var (
		out     []byte
		s       string
		names   []unicodeNameReplace
		toFile  = path.Join(saveDir, file)
		temp    = filepath.Join(global.PWD, "temp")
		saveCmd = exec.Command("python", "./io/python/io.py", "obfuscateFile", toFile, temp)
	)
	//p.names = make([][]rune, 0, 40)
	if out, err = p.readFile(toFile); err != nil {
		return
	}
	//ioutil.WriteFile("loaded.json", out, 0644)
	/*for i := 0; i < len(out); i++ {
		if out[i] == '\\' && out[i+1] == 'x' {
			j := string(out[i-20 : i+40])
			print(j)
			i += 50
		}
	}*/
	if s, err = p.getSaveData(string(out)); err != nil {
		return
	}
	//if err == nil {
	//err = ioutil.WriteFile("loaded_pre.json", out, 0644)
	//}
	/*if strings.Contains(s, "\\x") {
		// For foreign langauge, need to double-escape the x
		p.getUnicodeNames(s)
		//for i, n := range p.names {
		//	s = strings.Replace(s, n, fmt.Sprintf(";;;%d;;;", i), 1)
		//}
		s = strings.ReplaceAll(s, "\\x", "x")
	}*/
	s = strings.ReplaceAll(s, `\\r\\n`, "")
	s = p.fixEscapeCharsForLoad(s)
	if strings.Contains(s, "\\x") {
		b := []byte(s)
		if b, err = p.replaceUnicodeNames(b, &names); err != nil {
			return
		}
		s = string(b)
	}
	//s = p.fixFile(s)

	if len(os.Args) >= 2 && os.Args[1] == "print" {
		if _, err = os.Stat("encounters.json"); errors.Is(err, os.ErrNotExist) {
			if _, err = os.Create("encounters.json"); err != nil {
			}
		}
		t := strings.ReplaceAll(s, `\`, ``)
		t = strings.ReplaceAll(t, `"{`, `{`)
		t = strings.ReplaceAll(t, `}"`, `}`)
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, []byte(t), "", "\t")
		if err == nil {
			err = ioutil.WriteFile("encounters.json", prettyJSON.Bytes(), 0644)
		}
	}

	base := ordered.NewOrderedMap()
	base.UnmarshalJSON([]byte(s))
	md := ordered.NewOrderedMap()
	if err = p.unmarshalFrom(base, "monsterDefeats", md); err != nil {
		return errors.New("unable to parse encounters file, monsterDefeats")
	}
	vs := md.Get("values")
	if vs == nil {
		return errors.New("unable to parse encounters file, monsterDefeats.values")
	}
	sl := vs.([]interface{})
	total := 2000
	for i := range sl {
		if i == 0 {
			sl[i] = json.Number("2000")
			continue
		}
		if sl[i].(json.Number) == "0" {
			sl[i] = json.Number("1")
			total++
		} else {
			j, _ := sl[i].(json.Number).Int64()
			total += int(j)
		}
	}
	md.Set("values", sl)
	base.Set("totalSubjugationCount", total)

	var sf interface{}
	if sf, err = p.getFromTarget(base, "scenarioFlags"); err != nil {
		return
	}
	sl = sf.([]interface{})
	for i := range sl {
		if sl[i].(json.Number) == "0" {
			sl[i] = json.Number("1")
		}
	}
	if err = p.setTarget(base, "scenarioFlags", sl); err != nil {
		return
	}

	if err = p.marshalTo(base, "monsterDefeats", md); err != nil {
		return
	}

	var data []byte
	if data, err = json.Marshal(base); err != nil {
		return
	}

	if _, err = os.Stat(temp); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create(temp); err != nil {
			return fmt.Errorf("failed to create save file %s: %v", toFile, err)
		}
	}
	defer os.Remove(temp)

	if err = os.WriteFile(temp, data, 0644); err != nil {
		return fmt.Errorf("failed to create temp file %s: %v", toFile, err)
	}

	if _, err = os.Stat(toFile); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create(toFile); err != nil {
			return fmt.Errorf("failed to create save file %s: %v", toFile, err)
		}
	}

	if out, err = saveCmd.Output(); err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			err = errors.New(string(ee.Stderr))
		} else {
			err = fmt.Errorf("%s: %v", string(out), err)
		}
	}
	return
}
