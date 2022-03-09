package pr

import (
	"encoding/json"
	"errors"
	"ffvi_editor/global"
	"ffvi_editor/models"
	"fmt"
	jo "gitlab.com/c0b/go-ordered-json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func (p *PR) Save(slot int, fileName string) (err error) {
	var (
		toFile   = filepath.Join(global.Dir, fileName)
		temp     = filepath.Join(global.PWD, "temp")
		cmd      = exec.Command("python", "./io/python/io.py", "obfuscateFile", toFile, temp)
		slTarget = jo.NewOrderedMap()
	)

	if err = p.saveCharacters(); err != nil {
		return
	}
	if err = p.saveSkills(); err != nil {
		return
	}
	if err = p.saveEspers(); err != nil {
		return
	}
	if err = p.saveInventory(); err != nil {
		return
	}
	if err = p.saveMiscStats(); err != nil {
		return
	}

	iSlice := make([]interface{}, 0, len(p.Characters))
	for _, c := range p.Characters {
		if c != nil {
			var k []byte
			if k, err = c.MarshalJSON(); err != nil {
				return
			}
			s := string(k)
			iSlice = append(iSlice, s)
		}
	}

	if err = p.unmarshalFrom(p.UserData, OwnedCharacterList, slTarget); err != nil {
		return
	}
	slTarget.Set(targetKey, iSlice)
	if err = p.MarshalTo(p.UserData, OwnedCharacterList, slTarget); err != nil {
		return
	}

	if err = p.MarshalTo(p.Base, UserDatas, p.UserData); err != nil {
		return
	}

	if err = p.setValue(p.Base, "id", slot); err != nil {
		return
	}

	if p.data, err = json.Marshal(p.Base); err != nil {
		return
	}

	p.data = p.unfixFile(p.data)

	if _, err = os.Stat(temp); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create(temp); err != nil {
			return fmt.Errorf("failed to create save file %s: %v", toFile, err)
		}
	}
	defer os.Remove(temp)

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

func (p *PR) saveCharacters() (err error) {
	for _, d := range p.Characters {
		if d == nil {
			continue
		}

		var id int
		if id, err = p.getInt(d, ID); err != nil {
			return
		}

		o, found := CharBaseByCharacterID[id-1]
		if !found {
			continue
		}

		c := models.GetCharacter(o.Name)

		if err = p.setValue(d, Name, c.Name); err != nil {
			return
		}

		params := jo.NewOrderedMap()
		if err = p.unmarshalFrom(d, Parameter, params); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalLevel, c.Level); err != nil {
			return
		}

		// TODO
		//if err = p.setValue(params, CurrentHP, c.HP.Current); err != nil {
		//	return
		//}
		if err = p.setValue(params, AdditionalMaxHp, floor0(c.HP.Max-o.HPBase)); err != nil {
			return
		}

		// TODO
		//if err = p.setValue(params, CurrentMP, c.MP.Current); err != nil {
		//	return
		//}
		if err = p.setValue(params, AdditionalMaxMp, floor0(c.MP.Max-o.MPBase)); err != nil {
			return
		}

		if err = p.setValue(d, CurrentExp, c.Exp); err != nil {
			return
		}

		// TODO Status

		// TODO Commands

		if err = p.setValue(params, AdditionalPower, c.Vigor); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalVitality, c.Stamina); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalAgility, c.Speed); err != nil {
			return
		}

		if err = p.setValue(params, AdditionMagic, c.Magic); err != nil {
			return
		}

		// TODO Equipment

		if err = p.MarshalTo(d, Parameter, params); err != nil {
			return
		}
	}
	return
}

func (p *PR) saveSkills() (err error) {
	// TODO
	return
}

func (p *PR) saveEspers() (err error) {
	// TODO
	return
}

func (p *PR) saveInventory() (err error) {
	// TODO
	return
	/*var (
		rows     = models.GetInventory().GetRowsForPrSave()
		sl       = make([]interface{}, len(rows))
		b        []byte
		slTarget = jo.NewOrderedMap()
	)

	for i, r := range rows {
		if b, err = json.Marshal(r); err != nil {
			return
		}
		sl[i] = string(b)
	}
	slTarget.Set(targetKey, sl)
	return p.MarshalTo(p.UserData, NormalOwnedItemList, slTarget)*/
}

func (p *PR) saveMiscStats() (err error) {
	// TODO
	return
}

func (p *PR) setValue(to *jo.OrderedMap, key string, value interface{}) (err error) {
	if !to.Has(key) {
		err = fmt.Errorf("unable to find %s", key)
	}
	to.Set(key, value)
	return
}

func (p *PR) MarshalTo(to *jo.OrderedMap, key string, value *jo.OrderedMap) error {
	if !to.Has(key) {
		return fmt.Errorf("unable to find %s", key)
	}
	if v, err := value.MarshalJSON(); err != nil {
		return err
	} else {
		to.Set(key, string(v))
	}
	return nil
}

func floor0(i int) int {
	if i < 0 {
		return 0
	}
	return i
}

func (p *PR) unfixFile(data []byte) []byte {
	//if p.fileEnd == "" {
	return data
	//}

	//s := string(data)
	//i := strings.Index(s, "clearFlag")
	//s = s[:i+9] + p.fileEnd
	//return []byte(s)
}
