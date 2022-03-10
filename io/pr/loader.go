package pr

import (
	"encoding/json"
	"ffvi_editor/models"
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
		out []byte
		i   interface{}
	)
	if out, err = p.execPythonLoad(fileName, true); err != nil {
		if out, err = p.execPythonLoad(fileName, false); err != nil {
			return
		}
	}

	var end int
	end = len(out)
	for end = len(out) - 1; end > 0 && out[end-1] != '}'; end-- {
	}
	p.data = make([]byte, len(out))
	for a := 2; a < end; a++ {
		p.data[a] = out[a]
	}
	s := string(out[2:end])
	s = strings.ReplaceAll(s, `\\r\\n`, "")
	s = p.fixEscapeCharsForLoad(s)
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
	TODO END /*/

	if err = p.Base.UnmarshalJSON([]byte(s)); err != nil {
		return
	}

	if err = p.unmarshalFrom(p.Base, UserDatas, p.UserData); err != nil {
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
	if err = p.loadSkills(); err != nil {
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

		o, found := CharBaseByCharacterID[id-1]
		if !found {
			continue
		}

		c := models.GetCharacter(o.Name)

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

		if c.MP.Current, err = p.getInt(params, CurrentHP); err != nil {
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

		// TODO Commands

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

		// TODO
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
	}
	return
}

func (p *PR) loadSkills() (err error) {
	return
}

func (p *PR) loadEspers() (err error) {
	return
}

func (p *PR) loadMiscStats() (err error) {
	if models.GetMisc().GP, err = p.getInt(p.UserData, OwnedGil); err != nil {
		return
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
	if err = p.unmarshalFrom(p.UserData, key, slTarget); err != nil {
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

func (p *PR) loadBase(s string) (err error) {
	if err = p.Base.UnmarshalJSON([]byte(s)); err != nil {
		/*i := strings.Index(s, "clearFlag")
		if i != -1 {
			c := s[i+9]
			if c != ':' {
				cc := s[i+10]
				if cc >= 48 && c <= 57 {
					cc -= 48
				} else {
					cc = 0
				}
				s = s[:i+9] + fmt.Sprintf(`":%d}`, cc)
				return p.Base.UnmarshalJSON([]byte(s))
			}
		}*/
		return
	}

	p.data = []byte(s)
	return
}

func (p *PR) fixFile(s string) string {
	/*	i := strings.Index(s, "clearFlag")
		if i != -1 {
			c := s[i+9]
			if c != ':' {
				cc := s[i+10]
				if cc >= 48 && c <= 57 {
					cc -= 48
				} else {
					cc = 0
				}
				p.fileEnd = s[i+9:]
				s = s[:i+9] + fmt.Sprintf(`":%d}`, cc)
			}
		}*/
	return s
}
