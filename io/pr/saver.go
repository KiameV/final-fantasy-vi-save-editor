package pr

import (
	"encoding/json"
	"errors"
	"ffvi_editor/global"
	"ffvi_editor/models/consts"
	"ffvi_editor/models/consts/pr"
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
		needed   = make(map[int]int)
		slTarget = jo.NewOrderedMap()
	)

	/*/ TODO Test bulk item override
	j := 0
	for i := 30; i <= 42; i++ {
		pri.GetInventory().Set(j, pri.Row{
			ItemID: 200 + i,
			Count:  i,
		})
		j++
	}
	//*/

	p.populateNeeded(&needed)
	pri.GetInventory().AddNeeded(needed)

	if err = p.saveCharacters(); err != nil {
		return
	}
	if err = p.saveInventory(); err != nil {
		return
	}
	if err = p.saveEspers(); err != nil {
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

	if err = p.marshalTo(p.Base, UserData, p.UserData); err != nil {
		return
	}

	if err = p.setValue(p.Base, "id", slot); err != nil {
		return
	}

	var data []byte
	if data, err = json.Marshal(p.Base); err != nil {
		return
	}

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
	// TODO END /*/

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

		var jobID int
		if jobID, err = p.getInt(d, JobID); err != nil {
			return
		}

		o, found := GetCharacter(id, jobID)
		if !found {
			continue
		}

		c := pri.GetCharacter(o.Name)

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

		if err = p.setValue(params, CurrentHP, c.HP.Current); err != nil {
			return
		}
		if err = p.setValue(params, AdditionalMaxHp, floor0(c.HP.Max /*-o.HPBase*/)); err != nil {
			return
		}

		if err = p.setValue(params, CurrentMP, c.MP.Current); err != nil {
			return
		}
		if err = p.setValue(params, AdditionalMaxMp, floor0(c.MP.Max /*-o.MPBase*/)); err != nil {
			return
		}

		if err = p.setValue(d, CurrentExp, c.Exp); err != nil {
			return
		}

		// TODO Status

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

		eq := jo.NewOrderedMap()
		if err = p.unmarshalFrom(d, EquipmentList, eq); err != nil {
			return
		}
		invCounts := pri.GetInventory().GetItemLookup()
		var eqIDCounts []string
		p.tryGetInvCount(&eqIDCounts, invCounts, c.Equipment.WeaponID)
		p.tryGetInvCount(&eqIDCounts, invCounts, c.Equipment.ShieldID)
		p.tryGetInvCount(&eqIDCounts, invCounts, c.Equipment.HelmetID)
		p.tryGetInvCount(&eqIDCounts, invCounts, c.Equipment.ArmorID)
		p.tryGetInvCount(&eqIDCounts, invCounts, c.Equipment.Relic1ID)
		p.tryGetInvCount(&eqIDCounts, invCounts, c.Equipment.Relic2ID)
		eq.Set("values", eqIDCounts)

		if err = p.marshalTo(d, EquipmentList, eq); err != nil {
			return
		}

		if err = p.marshalTo(d, Parameter, params); err != nil {
			return
		}

		// Cyan
		if jobID == 3 {
			if err = p.saveSkills(d, pr.BushidoFrom, pr.BushidoTo, pr.BushidoLookupByID); err != nil {
				return
			}
		}
		// Sabin
		if jobID == 6 {
			if err = p.saveSkills(d, pr.BlitzFrom, pr.BlitzTo, pr.BlitzLookupByID); err != nil {
				return
			}
		}
		// Mog
		if id == 16 {
			if err = p.saveSkills(d, pr.DanceFrom, pr.DanceTo, pr.DanceLookupByID); err != nil {
				return
			}
		}
		// Strago
		if jobID == 8 {
			if err = p.saveSkills(d, pr.LoreFrom, pr.LoreTo, pr.LoreLookupByID); err != nil {
				return
			}
		}
		// Gau
		if jobID == 12 {
			if err = p.saveSkills(d, pr.RageFrom, pr.RageTo, pr.RageLookupByID); err != nil {
				return
			}
		}
	}
	return
}

func (p *PR) populateNeeded(needed *map[int]int) {
	for _, c := range pri.Characters {
		p.addToNeeded(needed, c.Equipment.WeaponID)
	}
}

func (p *PR) addToNeeded(needed *map[int]int, id int) {
	if count, found := (*needed)[id]; !found {
		(*needed)[id] = 1
	} else {
		(*needed)[id] = count + 1
	}
}

func (p *PR) saveSkills(d *jo.OrderedMap, from int64, to int64, nvcLookup map[int]*consts.NameValueChecked) (err error) {
	var (
		b      []byte
		i      interface{}
		found  bool
		nvc    *consts.NameValueChecked
		lookup = make(map[int]bool)
	)
	if i, err = p.getFromTarget(d, AbilityList); err != nil {
		return
	}
	isl := i.([]interface{})
	for j := 0; j < len(isl); j++ {
		m := jo.NewOrderedMap()
		if err = m.UnmarshalJSON([]byte(isl[j].(string))); err != nil {
			return
		}
		if v, ok := m.GetValue("abilityId"); ok {
			if iv, _ := v.(json.Number).Int64(); iv >= from && iv <= to {
				if nvc, found = nvcLookup[int(iv)]; found {
					if nvc.Checked {
						m.Set("skillLevel", 100)
					} else {
						m.Set("skillLevel", 0)
					}
				}
				lookup[int(iv)] = true
				if b, err = m.MarshalJSON(); err != nil {
					return
				}
				isl[j] = string(b)
			}
		}
	}
	for _, r := range pr.Rages {
		if r.Checked {
			if _, found = lookup[r.Value]; !found {
				m := jo.NewOrderedMap()
				m.Set("abilityId", r.Value)
				m.Set("contentId", r.Value+168)
				m.Set("skillLevel", 100)
				if b, err = m.MarshalJSON(); err != nil {
					return
				}
				isl = append(isl, string(b))
			}
		}
	}
	if err = p.setTarget(d, AbilityList, isl); err != nil {
		return
	}
	return
}

func (p *PR) saveEspers() (err error) {
	var sl []int
	for _, e := range pr.Espers {
		if e.Checked {
			sl = append(sl, e.Value)
		}
	}
	return p.setTarget(p.UserData, OwnedMagicStoneList, sl)
}

func (p *PR) saveInventory() (err error) {
	var (
		rows             = pri.GetInventory().GetRowsForPrSave()
		sl               = make([]interface{}, 0, len(rows))
		b                []byte
		slTarget         = jo.NewOrderedMap()
		found            = make(map[int]bool)
		removeDuplicates = pri.GetInventory().RemoveDuplicates
	)

	for _, r := range rows {
		if removeDuplicates {
			if _, f := found[r.ItemID]; f {
				continue
			}
			found[r.ItemID] = true
		}
		// Skip known crashing items
		if r.ItemID == 184 || r.ItemID == 243 {
			continue
		}
		// Skip Empty rows
		if r.ItemID == 0 || r.Count == 0 {
			continue
		}
		if b, err = json.Marshal(r); err != nil {
			return
		}
		sl = append(sl, string(b))
	}

	slTarget.Set(targetKey, sl)
	if err = p.marshalTo(p.UserData, NormalOwnedItemList, slTarget); err != nil {
		return
	}

	if pri.GetInventory().ResetSortOrder {
		slTarget = jo.NewOrderedMap()
		slTarget.Set(targetKey, make([]interface{}, 0))
		if err = p.marshalTo(p.UserData, NormalOwnedItemSortIdList, slTarget); err != nil {
			return
		}
	}
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

func (p *PR) tryGetInvCount(eq *[]string, counts map[int]int, id int) {
	var i idCount
	if id == 0 {
		return
	}
	if count, ok := counts[id]; ok {
		i.ContentID = id
		i.Count = count
	} else {
		i.ContentID = 200
		i.Count = counts[200]
	}
	b, _ := json.Marshal(&i)
	*eq = append(*eq, string(b))
}

func (p *PR) setTarget(d *jo.OrderedMap, key string, value interface{}) (err error) {
	var (
		t = jo.NewOrderedMap()
		b []byte
	)
	t.Set(targetKey, value)
	b, err = t.MarshalJSON()
	return p.setValue(d, key, string(b))
}
