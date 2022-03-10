package pr

import (
	"encoding/json"
	"errors"
	"ffvi_editor/global"
	"ffvi_editor/models"
	pri "ffvi_editor/models/pr"
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
	if err = p.marshalTo(p.UserData, OwnedCharacterList, slTarget); err != nil {
		return
	}

	if err = p.marshalTo(p.Base, UserDatas, p.UserData); err != nil {
		return
	}

	if err = p.setValue(p.Base, "id", slot); err != nil {
		return
	}

	var data []byte
	if data, err = json.Marshal(p.Base); err != nil {
		return
	}

	//p.data = p.unfixFile(p.data)

	if _, err = os.Stat(temp); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create(temp); err != nil {
			return fmt.Errorf("failed to create save file %s: %v", toFile, err)
		}
	}
	defer os.Remove(temp)

	/*/ TODO Debug
	if _, err = os.Stat("saved.json"); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create("saved.json"); err != nil {
			return fmt.Errorf("failed to create save file %s: %v", toFile, err)
		}
	}
	s := string(p.data)
	s = strings.ReplaceAll(s, `\`, ``)
	s = strings.ReplaceAll(s, `"{`, `{`)
	s = strings.ReplaceAll(s, `}"`, `}`)
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, []byte(s), "", "\t")
	if err != nil {
		err = ioutil.WriteFile("saved.json", prettyJSON.Bytes(), 0644)
	}
	TODO END /*/

	if err = ioutil.WriteFile(temp, data, 0644); err != nil {
		return fmt.Errorf("failed to create temp file %s: %v", toFile, err)
	}

	if _, err = os.Stat(toFile); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create(toFile); err != nil {
			return fmt.Errorf("failed to create save file %s: %v", toFile, err)
		}
	}

	var out []byte
	out, err = cmd.Output()
	if err != nil && len(out) > 0 {
		err = fmt.Errorf("%s: %v", string(out), err)
	}
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
		eq := jo.NewOrderedMap()
		if err = p.unmarshalFrom(d, EquipmentList, eq); err != nil {
			return
		}
		invCounts := pri.GetInventory().GetItemLookup()
		eqIDCounts := []string{
			p.getInvCount(invCounts, c.Equipment.WeaponID),
			p.getInvCount(invCounts, c.Equipment.ShieldID),
			p.getInvCount(invCounts, c.Equipment.HelmetID),
			p.getInvCount(invCounts, c.Equipment.ArmorID),
			p.getInvCount(invCounts, c.Equipment.Relic1ID),
			p.getInvCount(invCounts, c.Equipment.Relic2ID),
		}
		eq.Set("values", eqIDCounts)

		if err = p.marshalTo(d, EquipmentList, eq); err != nil {
			return
		}

		if err = p.marshalTo(d, Parameter, params); err != nil {
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
	//lookup := map[int]bool{}
	var (
		rows     = pri.GetInventory().GetRowsForPrSave()
		sl       = make([]interface{}, len(rows))
		b        []byte
		slTarget = jo.NewOrderedMap()
	)

	for i, r := range rows {
		/*/ TODO remove start
		r = pri.Row{
			ItemID: r.ItemID,
			Count:  r.Count,
		}
		if r.ItemID > 101 && r.ItemID <= 199 {
			r.Count = r.ItemID - 100
		} else {
			r.Count = 1
		}
		lookup[r.ItemID] = true
		// TODO remove end */

		if b, err = json.Marshal(r); err != nil {
			return
		}
		sl[i] = string(b)
	}
	/*/ TODO Undo start
	for i := 101; i < 150; i++ {
		if _, ok := lookup[i]; !ok {
			r := pri.Row{
				ItemID: i,
				Count:  i - 100,
			}
			if b, err = json.Marshal(r); err != nil {
				return
			}
			sl[i] = string(b)
		}
	}
	// TODO Remove end */

	slTarget.Set(targetKey, sl)
	if err = p.marshalTo(p.UserData, NormalOwnedItemList, slTarget); err != nil {
		return
	}

	// TODO Readd
	//if pri.GetInventory().ResetSortOrder {
	slTarget = jo.NewOrderedMap()
	slTarget.Set(targetKey, make([]interface{}, 0))
	if err = p.marshalTo(p.UserData, NormalOwnedItemSortIdList, slTarget); err != nil {
		return
	}
	//}
	return
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

func (p *PR) marshalTo(to *jo.OrderedMap, key string, value *jo.OrderedMap) error {
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

func (p *PR) getInvCount(counts map[int]int, id int) string {
	var i idCount
	if count, ok := counts[id]; ok {
		i.ContentID = id
		i.Count = count
	} else {
		i.ContentID = 200
		i.Count = counts[200]
	}
	b, _ := json.Marshal(&i)
	return string(b)
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
