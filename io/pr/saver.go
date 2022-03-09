package pr

import (
	"encoding/json"
	"errors"
	"ffvi_editor/global"
	"ffvi_editor/models"
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
		//slTarget = map[string]interface{}{}
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

	/*sl := slTarget[targetKey].([]interface{})
	for j, c := range sl {
		if err = p.Unmarshal(c, &p.Characters[j]); err != nil {
			return
		}
	}*/

	//if err = p.MarshalTo(p.UserData, OwnedCharacterList, &slTarget); err != nil {
	//	return
	//}

	if err = p.MarshalTo(p.Base, UserDatas, &p.UserData); err != nil {
		return
	}

	if p.data, err = json.Marshal(p.Base); err != nil {
		return
	}

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

		params := make(map[string]interface{})
		if err = p.UnmarshalFrom(d, Parameter, &params); err != nil {
			return
		}

		if err = p.setValue(params, AdditionalLevel, c.Level); err != nil {
			return
		}

		if err = p.setValue(params, CurrentHP, c.HP.Current); err != nil {
			return
		}
		if err = p.setValue(params, AdditionalMaxHp, c.HP.Max-o.HPBase); err != nil {
			return
		}

		if err = p.setValue(params, CurrentMP, c.MP.Current); err != nil {
			return
		}
		if err = p.setValue(params, AdditionalMaxMp, c.MP.Max-o.MPBase); err != nil {
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

		if err = p.setValue(params, AdditionalIntelligence, c.Magic); err != nil {
			return
		}

		// TODO Equipment

		if err = p.MarshalTo(d, Parameter, &params); err != nil {
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
}

func (p *PR) saveMiscStats() (err error) {
	// TODO
	return
}

func (p *PR) setValue(m map[string]interface{}, key string, value interface{}) (err error) {
	if _, ok := m[key]; !ok {
		err = fmt.Errorf("unable to find %s", key)
	}
	m[key] = value
	return
}

func (p *PR) MarshalTo(to map[string]interface{}, key string, value *map[string]interface{}) error {
	if _, ok := to[key]; !ok {
		return fmt.Errorf("unable to find %s", key)
	}
	if v, err := json.Marshal(value); err != nil {
		return err
	} else {
		to[key] = string(v)
	}
	return nil
}
