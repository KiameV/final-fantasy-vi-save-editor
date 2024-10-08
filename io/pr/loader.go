package pr

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"ffvi_editor/global"
	"ffvi_editor/io/config"
	"ffvi_editor/io/file"
	"ffvi_editor/models"
	"ffvi_editor/models/consts"
	"ffvi_editor/models/consts/pr"
	pri "ffvi_editor/models/pr"
	jo "gitlab.com/c0b/go-ordered-json"
)

func (p *PR) Load(fromFile string, saveType global.SaveFileType) (err error) {
	var (
		out   []byte
		i     interface{}
		s     string
		names []unicodeNameReplace
	)

	if out, p.fileTrimmed, err = file.LoadFile(fromFile, saveType); err != nil {
		return
	}
	s = string(out)

	if err = p.loadBase(s); err != nil {
		return
	}

	if err = p.unmarshalFrom(p.Base, UserData, p.UserData); err != nil {
		return
	}

	if err = p.unmarshalFrom(p.Base, MapData, p.MapData); err != nil {
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

	pri.GetParty().Clear()

	if err = p.loadCharacters(); err != nil {
		return
	}
	if err = p.loadParty(); err != nil {
		return
	}
	if err = p.loadEspers(); err != nil {
		return
	}
	if err = p.loadMiscStats(); err != nil {
		return
	}
	if err = p.loadInventory(NormalOwnedItemList, pri.GetInventory()); err != nil {
		return
	}
	if err = p.loadInventory(importantOwnedItemList, pri.GetImportantInventory()); err != nil {
		return
	}
	if err = p.loadVeldt(); err != nil {
		return
	}
	if err = p.loadCheats(); err != nil {
		return
	}
	if err = p.loadMapData(); err != nil {
		return
	}
	if err = p.loadTransportation(); err != nil {
		return
	}

	if len(names) > 0 {
		p.names = names
	}
	return
}

func (p *PR) loadParty() (err error) {
	var (
		party = pri.GetParty()
		// id     int
		// name   string
		i      interface{}
		member pri.Member
	)
	/*for _, d := range p.Characters {
		if d == nil {
			continue
		}
		if id, err = p.getInt(d, ID); err != nil {
			return
		}
		if name, err = p.getString(d, Name); err != nil {
			return
		}
		party.AddPossibleMember(&pri.Member{
			CharacterID: id,
			Name:        name,
		})
	}*/

	if i, err = p.getFromTarget(p.UserData, CorpsList); err != nil {
		return
	}
	for slot, c := range i.([]interface{}) {
		if err = json.Unmarshal([]byte(c.(string)), &member); err != nil {
			return
		}
		if err = party.SetMemberByID(slot, member.CharacterID); err != nil {
			return
		}
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

		o, found := pri.GetCharacterBaseOffset(id, jobID)
		if !found {
			continue
		}

		c := pri.GetCharacter(o.Name)
		c.EnableCommandsSave = config.AutoEnableCmd()

		if c.Name, err = p.getString(d, Name); err != nil {
			return
		}

		// if pr.IsMainCharacter(c.Name) {
		pri.GetParty().AddPossibleMember(&pri.Member{
			CharacterID: id,
			Name:        c.Name,
		})
		// }

		if c.IsEnabled, err = p.getBool(d, IsEnableCorps); err != nil {
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

		var values interface{}
		if values, err = p.getFromTarget(d, CommandList); err != nil {
			return
		}
		for i, v := range values.([]interface{}) {
			var j int64
			if j, err = v.(json.Number).Int64(); err != nil {
				return
			}
			if i >= len(c.Commands) {
				c.Commands = append(c.Commands, pr.CommandLookupByValue[int(j)])
			} else {
				c.Commands[i] = pr.CommandLookupByValue[int(j)]
			}
		}

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

		if err = p.loadEquipment(d, c); err != nil {
			return
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

func (p *PR) loadEquipment(d *jo.OrderedMap, c *models.Character) (err error) {
	c.Equipment.WeaponID = 93
	c.Equipment.ShieldID = 93
	c.Equipment.ArmorID = 199
	c.Equipment.HelmetID = 198
	c.Equipment.Relic1ID = 200
	c.Equipment.Relic2ID = 200

	var eqIDCounts []idCount
	if eqIDCounts, err = p.unmarshalEquipment(d); err != nil {
		return
	}

	if len(eqIDCounts) > 0 {
		c.Equipment.WeaponID = eqIDCounts[0].ContentID
	}
	if len(eqIDCounts) > 1 {
		c.Equipment.ShieldID = eqIDCounts[1].ContentID
	}
	if len(eqIDCounts) > 2 {
		c.Equipment.ArmorID = eqIDCounts[2].ContentID
	}
	if len(eqIDCounts) > 3 {
		c.Equipment.HelmetID = eqIDCounts[3].ContentID
	}
	if len(eqIDCounts) > 4 {
		c.Equipment.Relic1ID = eqIDCounts[4].ContentID
	}
	if len(eqIDCounts) > 5 {
		c.Equipment.Relic2ID = eqIDCounts[5].ContentID
	}
	return
}

/*func (p *PR) getEquipmentKeys(m *jo.OrderedMap) (keys []int, err error) {
	i, ok := m.GetValue(EquipmentList)
	if !ok {
		return nil, fmt.Errorf("%s not found", EquipmentList)
	}

	eq := jo.NewOrderedMap()
	if err = eq.UnmarshalJSON([]byte(i.(string))); err != nil {
		return
	}

	if i, ok = eq.GetValue("keys"); ok && i != nil {
		keys = make([]int, len(i.([]interface{})))
		for j, v := range i.([]interface{}) {
			var k int64
			if k, err = v.(json.Number).Int64(); err != nil {
				return
			}
			keys[j] = int(k)
		}
	}
	return
}*/

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
	if models.GetMisc().MonstersKilledCount, err = p.getInt(p.UserData, MonstersKilledCount); err != nil {
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

func (p *PR) loadInventory(key string, inventory *pri.Inventory) (err error) {
	var (
		sl  interface{}
		row pri.Row
	)
	if sl, err = p.getFromTarget(p.UserData, key); err != nil {
		return
	}
	inventory.Reset()
	for i, r := range sl.([]interface{}) {
		if err = json.Unmarshal([]byte(r.(string)), &row); err != nil {
			return
		}
		inventory.Set(i, row)
	}
	return nil
}

func (p *PR) loadMapData() (err error) {
	md := pri.GetMapData()
	if md.MapID, err = p.getInt(p.MapData, MapID); err != nil {
		return
	}
	if md.PointIn, err = p.getInt(p.MapData, PointIn); err != nil {
		return
	}
	if md.TransportationID, err = p.getInt(p.MapData, TransportationID); err != nil {
		return
	}
	if md.CarryingHoverShip, err = p.getBool(p.MapData, CarryingHoverShip); err != nil {
		return
	}
	if md.PlayableCharacterCorpsID, err = p.getInt(p.MapData, PlayableCharacterCorpsId); err != nil {
		return
	}

	pe := jo.NewOrderedMap()
	if err = p.unmarshalFrom(p.MapData, PlayerEntity, pe); err != nil {
		return
	}
	var pos *jo.OrderedMap
	if pos = pe.Get(PlayerPosition).(*jo.OrderedMap); pos == nil {
		err = errors.New("unable to get transportation position")
		return
	}
	if md.Player.X, err = p.getFloat(pos, "x"); err != nil {
		return
	}
	if md.Player.Y, err = p.getFloat(pos, "y"); err != nil {
		return
	}
	if md.Player.Z, err = p.getFloat(pos, "z"); err != nil {
		return
	}
	if md.PlayerDirection, err = p.getInt(pe, PlayerDirection); err != nil {
		return
	}

	gps := jo.NewOrderedMap()
	if err = p.unmarshalFrom(p.MapData, GpsData, gps); err != nil {
		return
	}
	if md.Gps.TransportationID, err = p.getInt(gps, GpsTransportationID); err != nil {
		err = nil
	}
	if md.Gps.MapID, err = p.getInt(gps, GpsDataMapID); err != nil {
		return
	}
	if md.Gps.AreaID, err = p.getInt(gps, GpsDataAreaID); err != nil {
		return
	}
	if md.Gps.GpsID, err = p.getInt(gps, GpsDataID); err != nil {
		return
	}
	if md.Gps.Width, err = p.getInt(gps, GpsDataWidth); err != nil {
		return
	}
	if md.Gps.Height, err = p.getInt(gps, GpsDataHeight); err != nil {
		return
	}
	return
}

func (p *PR) loadTransportation() (err error) {
	var sl interface{}
	if sl, err = p.getFromTarget(p.UserData, OwnedTransportationList); err != nil {
		return
	}
	pri.Transportations = make([]*pri.Transportation, len(sl.([]interface{})))
	for index, i := range sl.([]interface{}) {
		om := jo.NewOrderedMap()
		if err = om.UnmarshalJSON([]byte(i.(string))); err != nil {
			return
		}
		t := &pri.Transportation{}
		if t.ID, err = p.getInt(om, TransID); err != nil {
			return
		}
		if t.MapID, err = p.getInt(om, TransMapID); err != nil {
			return
		}
		if t.Direction, err = p.getInt(om, TransDirection); err != nil {
			return
		}
		if t.TimeStampTicks, err = p.getUint(om, TransTimeStampTicks); err != nil {
			return
		}

		var pos *jo.OrderedMap
		if pos = om.Get(TransPosition).(*jo.OrderedMap); pos == nil {
			err = errors.New("unable to get transportation position")
			return
		}
		if t.Position.X, err = p.getFloat(pos, "x"); err != nil {
			return
		}
		if t.Position.Y, err = p.getFloat(pos, "y"); err != nil {
			return
		}
		if t.Position.Z, err = p.getFloat(pos, "z"); err != nil {
			return
		}

		t.Enabled = t.TimeStampTicks > 0 && t.MapID > 0 && t.Position.X > 0 && t.Position.Y > 0 && t.Position.Z > 0

		pri.Transportations[index] = t
	}
	return
}

func (p *PR) loadVeldt() (err error) {
	var (
		veldt = pri.GetVeldt()
		sl    []interface{}
	)
	if sl, err = p.getJsonInts(p.MapData, BeastFieldEncountExchangeFlags); err != nil {
		return
	}
	veldt.Encounters = make([]bool, len(sl))
	for i, v := range sl {
		if v.(json.Number).String() == "1" {
			veldt.Encounters[i] = true
		}
	}
	return
}

func (p *PR) loadCheats() (err error) {
	c := pri.GetCheats()
	if c.OpenedChestCount, err = p.getInt(p.UserData, OpenChestCount); err != nil {
		return
	}
	if c.IsCompleteFlag, err = p.getFlag(p.Base, IsCompleteFlag); err != nil {
		return
	}
	if c.PlayTime, err = p.getFloat(p.UserData, PlayTime); err != nil {
		return
	}
	return
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

func (p *PR) getBool(c *jo.OrderedMap, key string) (b bool, err error) {
	j, ok := c.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}
	if b, ok = j.(bool); !ok {
		err = fmt.Errorf("unable to parse field %s value %v", key, j)
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

func (p *PR) getUint(c *jo.OrderedMap, key string) (i uint64, err error) {
	j, ok := c.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}

	k := reflect.ValueOf(j)
	switch k.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(k.Int()), nil
	case reflect.Float32, reflect.Float64:
		return uint64(k.Float()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uint64(k.Uint()), nil
	case reflect.String:
		var l int64
		l, err = strconv.ParseInt(k.String(), 10, 64)
		if err == nil {
			i = uint64(l)
		}
	default:
		err = fmt.Errorf("unable to parse field %s value %v ", key, j)
	}
	return
}

func (p *PR) getFloat(c *jo.OrderedMap, key string) (f float64, err error) {
	j, ok := c.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}

	k := reflect.ValueOf(j)
	switch k.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return k.Float(), nil
	case reflect.Float32, reflect.Float64:
		return k.Float(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return k.Float(), nil
	case reflect.String:
		f, err = strconv.ParseFloat(k.String(), 64)
	default:
		err = fmt.Errorf("unable to parse field %s value %v ", key, j)
	}
	return
}

func (p *PR) getJsonInts(data *jo.OrderedMap, key string) (ints []interface{}, err error) {
	j, ok := data.GetValue(key)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
	}
	ints, ok = j.([]interface{})
	if !ok {
		err = fmt.Errorf("unable to load %s", key)
	}
	return
}

func (p *PR) getFlag(m *jo.OrderedMap, key string) (b bool, err error) {
	var i int
	if i, err = p.getInt(m, key); err != nil {
		return
	}
	if i == 0 {
		b = false
	} else {
		b = true
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
	// s := strings.ReplaceAll(i.(string), `\\"`, `\"`)
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

	if i, ok = eq.GetValue("values"); ok && i != nil {
		idCounts = make([]idCount, len(i.([]interface{})))
		for j, v := range i.([]interface{}) {
			if err = json.Unmarshal([]byte(v.(string)), &idCounts[j]); err != nil {
				return
			}
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

func (p *PR) loadBase(s string) (err error) {
	return p.Base.UnmarshalJSON([]byte(s))
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

func extractArchiveFile(dest string, f *zip.File) (err error) {
	var (
		rc   io.ReadCloser
		file *os.File
		path string
	)
	if rc, err = f.Open(); err != nil {
		return
	}
	defer func() { _ = rc.Close() }()

	path = filepath.Join(dest, f.Name)
	// Check for ZipSlip (Directory traversal)
	path = strings.ReplaceAll(path, "..", "")

	if f.FileInfo().IsDir() {
		if err = os.MkdirAll(path, f.Mode()); err != nil {
			return
		}
	} else {
		if err = os.MkdirAll(filepath.Dir(path), f.Mode()); err != nil {
			return
		}
		if file, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode()); err != nil {
			return
		}
		defer func() { _ = file.Close() }()
		if _, err = io.Copy(file, rc); err != nil {
			return
		}
	}
	return
}

func (p *PR) debug_WriteOut(out []byte) {
	var m map[string]any
	if err := json.Unmarshal(out, &m); err == nil {
		p.debug_WriteSection("configData", &m)
		p.debug_WriteSection("dataStorage", &m)
		if md, loaded := p.debug_LoadMap("mapData", &m); loaded {
			p.debug_WriteSection("gpsData", &md)
			p.debug_WriteSection("playerEntity", &md)
			m["mapData"] = md
		}
		if ud, loaded := p.debug_LoadMap("userData", &m, true); loaded {
			p.debug_WriteSection("corpsList", &ud, true)
			if cp, l := p.debug_LoadMap("corpsSlots", &ud, true); l {
				p.debug_LoadValue("values", &cp)
				ud["corpsSlots"] = cp
			}
			// p.debug_WriteSection("learnedAbilityList", &ud)
			// p.debug_WriteSection("importantOwendItemList", &ud)
			// p.debug_WriteSection("normalOwnedItemList", &ud)
			// p.debug_WriteSection("normalOwnedItemSortIdList", &ud)
			if ow, l := p.debug_LoadMap("ownedCharacterList", &ud, true); l {
				p.debug_LoadValue("target", &ow)
				sl := ow["target"].([]map[string]any)
				for _, v := range sl {
					p.debug_WriteSection("abilitySlotDataList", &v, true)
				}
				ud["ownedCharacterList"] = ow
			}
			// p.debug_WriteSection("ownedKeyWaordList", &ud)
			// p.debug_WriteSection("ownedMagicList", &ud)
			// p.debug_WriteSection("ownedMagicStoneList", &ud)
			// p.debug_WriteSection("releasedJobs", &ud)
			// p.debug_WriteSection("warehouseItemList", &ud)
			m["userData"] = ud
		}
		out, _ = json.MarshalIndent(m, "", "  ")
	}
	_ = os.WriteFile("loaded.json", out, 0755)
}

func (p *PR) debug_WriteSection(key string, parent *map[string]any, skip ...bool) {
	if m, loaded := p.debug_LoadMap(key, parent, skip...); loaded {
		(*parent)[key] = m
	}
}

func (p *PR) debug_LoadMap(key string, parent *map[string]any, skip ...bool) (m map[string]any, loaded bool) {
	var v any
	if v, loaded = (*parent)[key]; loaded {
		s := v.(string)
		// for strings.Contains(s, `\\"`) {
		if len(skip) == 0 {
			s = strings.ReplaceAll(s, `\"`, `"`)
		}
		// }
		loaded = json.Unmarshal([]byte(s), &m) == nil
	}
	return
}

func (p *PR) debug_LoadValue(key string, parent *map[string]any) {
	var v []any
	if j, ok := (*parent)[key]; ok {
		if v, ok = j.([]any); ok {
			a := make([]map[string]any, len(v))
			for i, s := range v {
				if loaded := json.Unmarshal([]byte(s.(string)), &a[i]) == nil; !loaded {
					return
				}
			}
			(*parent)[key] = a
		}
	}
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
