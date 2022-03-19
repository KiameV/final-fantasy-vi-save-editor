package pr

import (
	"encoding/json"
	"errors"
	"ffvi_editor/models"
	"ffvi_editor/models/consts"
	"ffvi_editor/models/consts/pr"
	pri "ffvi_editor/models/pr"
	"fmt"
	jo "gitlab.com/c0b/go-ordered-json"
	"os/exec"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func (p *PR) Load(fileName string) (err error) {
	var (
		out   []byte
		i     interface{}
		s     string
		names []unicodeNameReplace
	)
	//p.names = make([][]rune, 0, 40)
	if out, err = p.readFile(fileName); err != nil {
		return
	}
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

	/*/ TODO Debug
	if _, err = os.Stat("loaded.json"); errors.Is(err, os.ErrNotExist) {
		if _, err = os.Create("loaded.json"); err != nil {
		}
	}
	t := strings.ReplaceAll(s, `\`, ``)
	t = strings.ReplaceAll(t, `"{`, `{`)
	t = strings.ReplaceAll(t, `}"`, `}`)
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, []byte(t), "", "\t")
	if err == nil {
		err = ioutil.WriteFile("loaded.json", prettyJSON.Bytes(), 0644)
	}
	///TODO END /*/

	if err = p.loadBase(s); err != nil {
		return
	}

	//if err = p.Base.UnmarshalJSON([]byte(s)); err != nil {
	//	return
	//}

	if err = p.unmarshalFrom(p.Base, UserData, p.UserData); err != nil {
		return
	}

	if i, err = p.getFromTarget(p.UserData, OwnedCharacterList); err != nil {
		return
	}

	for j, c := range i.([]interface{}) {
		if p.Characters[j] == nil {
			p.Characters[j] = jo.NewOrderedMap()
		}
		s = c.(string)
		if err = p.Characters[j].UnmarshalJSON([]byte(s)); err != nil {
			return
		}
	}

	if err = p.loadCharacters(); err != nil {
		return
	}
	if err = p.loadEspers(); err != nil {
		return
	}
	if err = p.loadMiscStats(); err != nil {
		return
	}
	if err = p.loadInventory(); err != nil {
		return
	}

	if len(names) > 0 {
		p.names = names
	}
	return
}

func (p *PR) loadCharacters() (err error) {
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

		if c.Name, err = p.getString(d, Name); err != nil {
			return
		}

		params := jo.NewOrderedMap()
		if err = p.unmarshalFrom(d, Parameter, params); err != nil {
			return
		}

		if c.Level, err = p.getInt(params, AdditionalLevel); err != nil {
			return
		}

		if c.HP.Current, err = p.getInt(params, CurrentHP); err != nil {
			return
		}
		if c.HP.Max, err = p.getInt(params, AdditionalMaxHp); err != nil {
			return
		}
		c.HP.Max += o.HPBase

		if c.MP.Current, err = p.getInt(params, CurrentMP); err != nil {
			return
		}
		if c.MP.Max, err = p.getInt(params, AdditionalMaxMp); err != nil {
			return
		}
		c.MP.Max += o.MPBase

		if c.Exp, err = p.getInt(d, CurrentExp); err != nil {
			return
		}

		// TODO Status

		// Command changing does not work in PR

		if c.Vigor, err = p.getInt(params, AdditionalPower); err != nil {
			return
		}

		if c.Stamina, err = p.getInt(params, AdditionalVitality); err != nil {
			return
		}

		if c.Speed, err = p.getInt(params, AdditionalAgility); err != nil {
			return
		}

		if c.Magic, err = p.getInt(params, AdditionMagic); err != nil {
			return
		}

		var eqIDCounts []idCount
		if eqIDCounts, err = p.unmarshalEquipment(d); err != nil {
			return
		}
		if len(eqIDCounts) > 0 {
			c.Equipment.WeaponID = eqIDCounts[0].ContentID
		} else {
			c.Equipment.WeaponID = 0
		}
		if len(eqIDCounts) > 1 {
			c.Equipment.ShieldID = eqIDCounts[1].ContentID
		} else {
			c.Equipment.ShieldID = 0
		}
		if len(eqIDCounts) > 2 {
			c.Equipment.HelmetID = eqIDCounts[2].ContentID
		} else {
			c.Equipment.HelmetID = 0
		}
		if len(eqIDCounts) > 3 {
			c.Equipment.ArmorID = eqIDCounts[3].ContentID
		} else {
			c.Equipment.ArmorID = 0
		}
		if len(eqIDCounts) > 4 {
			c.Equipment.Relic1ID = eqIDCounts[4].ContentID
		} else {
			c.Equipment.Relic1ID = 0
		}
		if len(eqIDCounts) > 5 {
			c.Equipment.Relic2ID = eqIDCounts[5].ContentID
		} else {
			c.Equipment.Relic2ID = 0
		}

		if err = p.loadSpells(d, c); err != nil {
			return
		}

		// Cyan
		if jobID == 3 {
			p.uncheckAll(pr.Bushidos)
			if err = p.loadSkills(d, pr.BushidoFrom, pr.BushidoTo, pr.BushidoLookupByID); err != nil {
				return
			}
		}
		// Sabin
		if jobID == 6 {
			p.uncheckAll(pr.Blitzes)
			if err = p.loadSkills(d, pr.BlitzFrom, pr.BlitzTo, pr.BlitzLookupByID); err != nil {
				return
			}
		}
		// Mog
		if id == 16 {
			p.uncheckAll(pr.Dances)
			if err = p.loadSkills(d, pr.DanceFrom, pr.DanceTo, pr.DanceLookupByID); err != nil {
				return
			}
		}
		// Strago
		if jobID == 8 {
			p.uncheckAll(pr.Lores)
			if err = p.loadSkills(d, pr.LoreFrom, pr.LoreTo, pr.LoreLookupByID); err != nil {
				return
			}
		}
		// Gau
		if jobID == 12 {
			p.uncheckAll(pr.Rages)
			if err = p.loadSkills(d, pr.RageFrom, pr.RageTo, pr.RageLookupByID); err != nil {
				return
			}
		}
	}
	return
}

func (p *PR) loadSpells(d *jo.OrderedMap, c *models.Character) (err error) {
	var i interface{}
	if i, err = p.getFromTarget(d, AbilityList); err != nil {
		return
	}
	sli := i.([]interface{})
	for j := 0; j < len(sli); j++ {
		m := jo.NewOrderedMap()
		if err = m.UnmarshalJSON([]byte(sli[j].(string))); err != nil {
			return
		}
		if v, ok := m.GetValue("abilityId"); ok {
			if iv, _ := v.(json.Number).Int64(); iv >= pr.SpellFrom && iv <= pr.SpellTo {
				if err = m.UnmarshalJSON([]byte(sli[j].(string))); err != nil {
					return
				}
				var spell *models.Spell
				if spell, ok = c.SpellsByID[int(iv)]; ok {
					k := m.Get("skillLevel")
					iv, _ = k.(json.Number).Int64()
					spell.Value = int(iv)
				}
			}
		}
	}
	return
}

func (p *PR) loadSkills(d *jo.OrderedMap, from int64, to int64, nvcLookup map[int]*consts.NameValueChecked) (err error) {
	var (
		i     interface{}
		nvc   *consts.NameValueChecked
		found bool
	)
	if i, err = p.getFromTarget(d, AbilityList); err != nil {
		return
	}
	sli := i.([]interface{})
	for j := 0; j < len(sli); j++ {
		m := jo.NewOrderedMap()
		if err = m.UnmarshalJSON([]byte(sli[j].(string))); err != nil {
			return
		}
		if v, ok := m.GetValue("abilityId"); ok {
			if iv, _ := v.(json.Number).Int64(); iv >= from && iv <= to {
				if err = m.UnmarshalJSON([]byte(sli[j].(string))); err != nil {
					return
				}
				k := m.Get("skillLevel")
				l, _ := k.(json.Number).Int64()
				if l == 100 {
					if nvc, found = nvcLookup[int(iv)]; found {
						nvc.Checked = true
					}
				}
			}
		}
	}
	return
}

func (p *PR) loadEspers() (err error) {
	var espers interface{}
	if espers, err = p.getFromTarget(p.UserData, OwnedMagicStoneList); err != nil {
		return
	}
	var id int64
	for _, e := range pr.Espers {
		e.Checked = false
	}
	if espers != nil {
		for _, n := range espers.([]interface{}) {
			if id, err = n.(json.Number).Int64(); err != nil {
				return
			}
			if e, found := pr.EspersByValue[int(id)]; found {
				e.Checked = true
			}
		}
	}
	return
}

func (p *PR) loadMiscStats() (err error) {
	if models.GetMisc().GP, err = p.getInt(p.UserData, OwnedGil); err != nil {
		return
	}
	if models.GetMisc().Steps, err = p.getInt(p.UserData, Steps); err != nil {
		return
	}
	if models.GetMisc().EscapeCount, err = p.getInt(p.UserData, EscapeCount); err != nil {
		return
	}
	if models.GetMisc().BattleCount, err = p.getInt(p.UserData, BattleCount); err != nil {
		return
	}
	if models.GetMisc().NumberOfSaves, err = p.getInt(p.UserData, SaveCompleteCount); err != nil {
		return
	}

	if ds, ok := p.Base.GetValue(DataStorage); ok {
		m := jo.NewOrderedMap()
		if err = m.UnmarshalJSON([]byte(ds.(string))); err != nil {
			return
		}
		if models.GetMisc().CursedShieldFightCount, err = p.getIntFromSlice(m, "global"); err != nil {
			return
		}
	}
	return
}

func (p *PR) loadInventory() (err error) {
	var (
		sl  interface{}
		inv = pri.GetInventory()
		row pri.Row
	)
	if sl, err = p.getFromTarget(p.UserData, NormalOwnedItemList); err != nil {
		return
	}
	inv.Reset()
	for i, r := range sl.([]interface{}) {
		if err = json.Unmarshal([]byte(r.(string)), &row); err != nil {
			return
		}
		inv.Set(i, row)
	}
	return nil
}

func (p *PR) getString(c *jo.OrderedMap, key string) (s string, err error) {
	j, ok := c.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}
	if s, ok = j.(string); !ok {
		err = fmt.Errorf("unable to parse field %s value %v ", key, j)
	}
	return
}

func (p *PR) getInt(c *jo.OrderedMap, key string) (i int, err error) {
	j, ok := c.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}

	k := reflect.ValueOf(j)
	switch k.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(k.Int()), nil
	case reflect.Float32, reflect.Float64:
		return int(k.Float()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int(k.Uint()), nil
	case reflect.String:
		var l int64
		l, err = strconv.ParseInt(k.String(), 10, 32)
		if err == nil {
			i = int(l)
		}
	default:
		err = fmt.Errorf("unable to parse field %s value %v ", key, j)
	}
	return
}

func (p *PR) getIntFromSlice(from *jo.OrderedMap, key string) (v int, err error) {
	var (
		i, ok = from.GetValue(key)
		sl    []interface{}
		i64   int64
	)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
		return
	}

	if sl, ok = i.([]interface{}); !ok || len(sl) < 9 {
		err = fmt.Errorf("unable to load cursed shield battle count")
		return
	}

	if i64, err = sl[9].(json.Number).Int64(); err != nil {
		return
	}
	v = int(i64)
	return
}

func (p *PR) unmarshalFrom(from *jo.OrderedMap, key string, m *jo.OrderedMap) (err error) {
	i, ok := from.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}
	return m.UnmarshalJSON([]byte(i.(string)))
}

func (p *PR) unmarshal(i interface{}, m *map[string]interface{}) error {
	//s := strings.ReplaceAll(i.(string), `\\"`, `\"`)
	return json.Unmarshal([]byte(i.(string)), m)
}

func (p *PR) unmarshalEquipment(m *jo.OrderedMap) (idCounts []idCount, err error) {
	i, ok := m.GetValue(EquipmentList)
	if !ok {
		return nil, fmt.Errorf("%s not found", EquipmentList)
	}

	eq := jo.NewOrderedMap()
	if err = eq.UnmarshalJSON([]byte(i.(string))); err != nil {
		return
	}

	i, _ = eq.GetValue("values")
	idCounts = make([]idCount, len(i.([]interface{})))
	for j, v := range i.([]interface{}) {
		if err = json.Unmarshal([]byte(v.(string)), &idCounts[j]); err != nil {
			return
		}
	}
	return
}

func (p *PR) getFromTarget(data *jo.OrderedMap, key string) (i interface{}, err error) {
	var (
		slTarget = jo.NewOrderedMap()
		ok       bool
	)
	if err = p.unmarshalFrom(data, key, slTarget); err != nil {
		return
	}
	if i, ok = slTarget.GetValue(targetKey); !ok {
		err = fmt.Errorf("unable to find %s", targetKey)
	}
	return
}

func (p *PR) fixEscapeCharsForLoad(s string) string {
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

type idCount struct {
	ContentID int `json:"contentId"`
	Count     int `json:"count"`
}

func (p *PR) execPythonLoad(fileName string, omitFirstBytes bool) ([]byte, error) {
	s := "0"
	if omitFirstBytes {
		s = "1"
	}
	cmd := exec.Command("python", "./io/python/io.py", "deobfuscateFile", fileName, s)
	return cmd.Output()
}

func (p *PR) initialSetupCmd() error {
	cmd := exec.Command("python", "-m", "pip", "install", "--upgrade", "pip")
	if _, err := cmd.Output(); err != nil {
		return err
	}

	cmd = exec.Command("pip", "install", "py3rijndael")
	if _, err := cmd.Output(); err != nil {
		return err
	}

	cmd = exec.Command("pip", "install", "pycryptodome")
	if _, err := cmd.Output(); err != nil {
		return err
	}
	return nil
}

func handleCmdError(err error) error {
	if e, ok := err.(*exec.ExitError); ok {
		return fmt.Errorf("Python failed to load file: " + string(e.Stderr))
	}
	return fmt.Errorf("Python failed to load file: " + err.Error())
}

func (p *PR) loadBase(s string) (err error) {
	return p.Base.UnmarshalJSON([]byte(s))
}

func (p *PR) getSaveData(s string) (string, error) {
	var (
		start int
		end   = strings.Index(s, `,"clearFlag`)
	)
	if end == -1 {
		return "", errors.New("unable to load file. Please try resaving to a new unused game slot and try loading that slot instead")
	}
	for start < len(s) && s[start] != '{' {
		start++
	}
	end = len(s) - 1
	for end >= 0 && s[end] != '}' {
		end--
	}
	if end+1 >= len(s) {
		return "", errors.New("unable to load file. Please try resaving to a new unused game slot and try loading that slot instead")
	}
	return s[start : end+1], nil // + `,"playTime":0.0,"clearFlag":0}`, nil
}

func (p *PR) readFile(fileName string) (out []byte, err error) {
	if out, err = p.execPythonLoad(fileName, true); err == nil {
		return
	}
	if out, err = p.execPythonLoad(fileName, false); err == nil {
		return
	}
	if err = p.initialSetupCmd(); err != nil {
		return
	}
	if out, err = p.execPythonLoad(fileName, true); err != nil {
		e1 := handleCmdError(err)
		if out, err = p.execPythonLoad(fileName, false); err != nil {
			err = e1
			return
		}
	}
	return
}

func (p *PR) uncheckAll(rages []*consts.NameValueChecked) {
	for _, r := range rages {
		r.Checked = false
	}
}

func (p *PR) replaceUnicodeNames(b []byte, names *[]unicodeNameReplace) ([]byte, error) {
	for i := 0; i < len(b)-10; i++ {
		if b[i] == '"' && b[i+1] == 'n' && b[i+2] == 'a' && b[i+3] == 'm' && b[i+4] == 'e' && b[i+5] == '\\' {
			i += 5
			for i < len(b)-1 && (b[i] == '\\' || b[i] == ':' || b[i] == ' ' || b[i] == '"') {
				i++
			}
			if b[i-1] != '"' {
				i--
			}

			// Start of name is i

			isUnicode := false
			for j := i; j < len(b)-50 && b[j] != '"' && !(b[j] == '\\' && b[j+1] == '\\'); j++ {
				if b[j] == '\\' && b[j+1] == 'x' {
					isUnicode = true
				}
			}
			if !isUnicode {
				continue
			}

			var original, replaced []byte
			for i < len(b)-1 && b[i] != '"' && !(b[i] == '\\' && b[i+1] == '\\') {
				original = append(original, b[i])
				replaced = append(replaced, '~')
				b[i] = '~'
				i++
			}
			r := unicodeNameReplace{
				Replaced: string(replaced),
			}
			var err error
			sb := strings.Builder{}
			for j := 0; j < len(original)-1; j++ {
				var char []byte
				if original[j] == '\\' && original[j+1] == 'x' {
					char = append(char, original[j])
					for j++; j < len(original); j++ {
						if original[j] == '\\' {
							break
						}
						char = append(char, original[j])
					}
					j--
					var s string
					if s, err = strconv.Unquote(strings.Replace(strconv.Quote(string(char)), `\\x`, `\x`, -1)); err != nil {
						return nil, err
					}
					sb.WriteString(s)
				} else {
					sb.WriteString(string(original[j]))
				}
			}
			r.Original = sb.String()
			*names = append(*names, r)
		}
	}
	return b, nil
}

type unicodeNameReplace struct {
	Original string
	Replaced string
}

/*
d := []byte(s)
		for i, c := range d {
			if c == 'x' && d[i-1] == '\\' {
				d[i] = '^'
				d[i-1] = '^'
			}
		}

func (p *PR) fixFile(s string) (bool, string) {
	if i := strings.Index(s, "clearFlag"); i != -1 {
		c := s[i+9]
		if c != ':' {
			cc := s[i+10]
			if cc >= 48 && c <= 57 {
				cc -= 48
			} else {
				cc = 0
			}
			s = s[:i+9] + fmt.Sprintf(`":%d}`, cc)
		}
		return true, s
	} else if i = strings.Index(s, `"playTime`); i != -1 && s[i+4] >= 48 && s[i+4] <= 57 {
		s = s[0:i] + `"playTime":0.0,"clearFlag":0}`
		return true, s
	}
	return false, s
}*/
