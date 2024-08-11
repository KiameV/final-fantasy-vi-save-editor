package pr

import (
	"encoding/json"
	"fmt"
	"strings"

	"ffvi_editor/global"
	"ffvi_editor/io/file"
	"ffvi_editor/models"
	"ffvi_editor/models/consts"
	"ffvi_editor/models/consts/pr"
	pri "ffvi_editor/models/pr"
	jo "gitlab.com/c0b/go-ordered-json"
)

func (p *PR) Save(slot int, toFile string, saveType global.SaveFileType) (err error) {
	var (
		// needed   = make(map[int]int)
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

	// p.populateNeeded(&needed)
	// pri.GetInventory().AddNeeded(needed)

	var addedItems []int
	if err = p.saveCharacters(&addedItems); err != nil {
		return
	}
	if err = p.saveInventory(NormalOwnedItemList, NormalOwnedItemSortIdList, pri.GetInventory(), addedItems); err != nil {
		return
	}
	if err = p.saveInventory(importantOwnedItemList, "", pri.GetImportantInventory(), nil); err != nil {
		return
	}
	if err = p.saveEspers(); err != nil {
		return
	}
	if err = p.saveMiscStats(); err != nil {
		return
	}
	if err = p.saveTransportation(); err != nil {
		return
	}
	if err = p.saveMapData(); err != nil {
		return
	}
	if pri.GetParty().Enabled {
		if err = p.saveParty(); err != nil {
			return
		}
	}
	if err = p.saveVeldt(); err != nil {
		return
	}
	if err = p.saveCheats(); err != nil {
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

	if err = p.marshalTo(p.Base, MapData, p.MapData); err != nil {
		return
	}

	if err = p.setValue(p.Base, "id", slot); err != nil {
		return
	}

	var data []byte
	if data, err = json.Marshal(p.Base); err != nil {
		return
	}

	return file.SaveFile(data, toFile, p.fileTrimmed, saveType)
}

func (p *PR) saveCharacters(addedItems *[]int) (err error) {
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

		if err = p.setValue(d, Name, c.Name); err != nil {
			return
		}

		if err = p.setValue(d, IsEnableCorps, c.IsEnabled); err != nil {
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

		maxHp := c.HP.Max - o.HPBase
		if err = p.setValue(params, AdditionalMaxHp, maxHp); err != nil {
			return
		}
		/*var v int
		if v, err = p.getInt(params, AdditionalMaxHp); err == nil {
			if v != c.HP.Max {
				if err = p.setValue(params, AdditionalMaxHp, v+(c.HP.Max-v)); err != nil {
					return
				}
			}
		}*/

		if err = p.setValue(params, CurrentMP, c.MP.Current); err != nil {
			return
		}

		maxMp := c.MP.Max - o.MPBase
		if err = p.setValue(params, AdditionalMaxMp, maxMp); err != nil {
			return
		}
		/*if v, err = p.getInt(params, AdditionalMaxMp); err == nil {
			if v != c.MP.Max {
				if err = p.setValue(params, AdditionalMaxMp, v+(c.MP.Max-v)); err != nil {
					return
				}
			}
		}*/

		if err = p.setValue(d, CurrentExp, c.Exp); err != nil {
			return
		}

		if c.EnableCommandsSave {
			sl := make([]interface{}, len(c.Commands))
			for i, cmd := range c.Commands {
				sl[i] = cmd.Value
			}
			if err = p.setTarget(d, CommandList, sl); err != nil {
				return
			}
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
		p.getInvCount(&eqIDCounts, invCounts, addedItems, c.Equipment.WeaponID, 93)
		p.getInvCount(&eqIDCounts, invCounts, addedItems, c.Equipment.ShieldID, 93)
		p.getInvCount(&eqIDCounts, invCounts, addedItems, c.Equipment.ArmorID, 199)
		p.getInvCount(&eqIDCounts, invCounts, addedItems, c.Equipment.HelmetID, 198)
		p.getInvCount(&eqIDCounts, invCounts, addedItems, c.Equipment.Relic1ID, 200)
		p.getInvCount(&eqIDCounts, invCounts, addedItems, c.Equipment.Relic2ID, 200)
		eq.Set("values", eqIDCounts)

		if err = p.marshalTo(d, EquipmentList, eq); err != nil {
			return
		}

		if err = p.marshalTo(d, Parameter, params); err != nil {
			return
		}

		if err = p.saveSpells(d, c); err != nil {
			return
		}

		// Cyan
		if jobID == 3 {
			if err = p.saveSkills(d, pr.BushidoFrom, pr.BushidoTo, pr.BushidoOffset, pr.BushidoLookupByID); err != nil {
				return
			}
		}
		// Sabin
		if jobID == 6 {
			if err = p.saveSkills(d, pr.BlitzFrom, pr.BlitzTo, pr.BlitzOffset, pr.BlitzLookupByID); err != nil {
				return
			}
		}
		// Mog
		if id == 16 {
			if err = p.saveSkills(d, pr.DanceFrom, pr.DanceTo, pr.DanceOffset, pr.DanceLookupByID); err != nil {
				return
			}
		}
		// Strago
		if jobID == 8 {
			if err = p.saveSkills(d, pr.LoreFrom, pr.LoreTo, pr.LoreOffset, pr.LoreLookupByID); err != nil {
				return
			}
		}
		// Gau
		if jobID == 12 {
			if err = p.saveSkills(d, pr.RageFrom, pr.RageTo, pr.RageOffset, pr.RageLookupByID); err != nil {
				return
			}
		}
	}
	return
}

func (p *PR) populateNeeded(needed *map[int]int) {
	for _, c := range pri.Characters {
		if c.IsEnabled { // pr.IsMainCharacter(c.RootName) {
			p.addToNeeded(needed, c.Equipment.WeaponID)
		}
	}
}

func (p *PR) addToNeeded(needed *map[int]int, id int) {
	if count, found := (*needed)[id]; !found {
		(*needed)[id] = 1
	} else {
		(*needed)[id] = count + 1
	}
}

type partyMember struct {
	ID          int `json:"id"`
	CharacterID int `json:"characterId"`
}

func (p *PR) saveParty() (err error) {
	var (
		party   = pri.GetParty()
		partyID = p.getPartyID()
		b       []byte
		sl      = make([]interface{}, 4)
	)
	for i, m := range party.Members {
		pm := partyMember{
			ID:          partyID,
			CharacterID: m.CharacterID,
		}
		if b, err = json.Marshal(&pm); err != nil {
			return
		}
		sl[i] = string(b)
	}
	return p.setTarget(p.UserData, CorpsList, sl)
}

func (p *PR) getPartyID() int {
	i, err := p.getFromTarget(p.UserData, CorpsList)
	if err != nil {
		return 1
	}
	var m partyMember
	for _, c := range i.([]interface{}) {
		if err = json.Unmarshal([]byte(c.(string)), &m); err != nil {
			return 1
		}
	}
	return m.ID
}

func (p *PR) saveSpells(d *jo.OrderedMap, c *models.Character) (err error) {
	var (
		b           []byte
		i           interface{}
		found       bool
		spell       *models.Spell
		lookup      = make(map[int]bool)
		knownSpells = make(map[int]bool)
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
			if iv, _ := v.(json.Number).Int64(); iv >= pr.SpellFrom && iv <= pr.SpellTo {
				if spell, found = c.SpellsByID[int(iv)]; found {
					if spell.Value == 100 {
						knownSpells[int(iv)] = true
					}
					m.Set("skillLevel", spell.Value)
					if b, err = m.MarshalJSON(); err != nil {
						return
					}
					isl[j] = string(b)
				}
				lookup[int(iv)] = true
			}
		}
	}
	for _, s := range c.SpellsByID {
		if s.Value > 0 {
			if _, found = lookup[s.Index]; !found {
				m := jo.NewOrderedMap()
				m.Set("abilityId", s.Value)
				m.Set("contentId", s.Value+pr.SpellOffset)
				m.Set("skillLevel", s.Value)
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

	return p.unlearnSpellsForCharacter(d, knownSpells)
}

func (p *PR) unlearnSpellsForCharacter(d *jo.OrderedMap, spells map[int]bool) (err error) {
	var (
		kv      = jo.NewOrderedMap()
		i       interface{}
		sl      []interface{}
		b       []byte
		ok      bool
		changed bool
	)
	if err = p.unmarshalFrom(d, AbilityDictionary, kv); err != nil {
		return
	}
	if i, ok = kv.GetValue("values"); ok {
		if sl, ok = i.([]interface{}); ok && len(sl) >= 2 {
			m := jo.NewOrderedMap()
			if err = m.UnmarshalJSON([]byte(sl[1].(string))); err != nil {
				return
			}
			if changed, err = p.unlearnSpells(m, spells); err != nil {
				return
			} else if changed {
				if b, err = m.MarshalJSON(); err != nil {
					return
				}
				sl[1] = string(b)
				kv.Set("value", sl)
			}
		}
	}
	if changed {
		if err = p.marshalTo(d, AbilityDictionary, kv); err != nil {
			return
		}
	}
	return
}

func (p *PR) unlearnSpells(om *jo.OrderedMap, knownSpells map[int]bool) (changed bool, err error) {
	changed = false
	if i, ok := om.GetValue("target"); ok {
		sl := i.([]interface{})
		var result []interface{}
		for _, r := range sl {
			m := jo.NewOrderedMap()
			if err = m.UnmarshalJSON([]byte(r.(string))); err != nil {
				return
			}
			if v, ok := m.GetValue("abilityId"); ok {
				if spellID, _ := v.(json.Number).Int64(); spellID >= pr.SpellFrom && spellID <= pr.SpellTo {
					if !knownSpells[int(spellID)] {
						changed = true
						continue
					}
				}
			}
			result = append(result, r)
		}
		if changed {
			if len(result) == 0 {
				om.Set("target", make([]interface{}, 0))
			} else {
				om.Set("target", result)
			}
		}
	}
	return
}

func (p *PR) saveSkills(d *jo.OrderedMap, from int64, to int64, offset int, nvcLookup map[int]*consts.NameValueChecked) (err error) {
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
					if b, err = m.MarshalJSON(); err != nil {
						return
					}
					isl[j] = string(b)
				}
				lookup[int(iv)] = true
			}
		}
	}
	for _, v := range nvcLookup {
		if v.Checked {
			if _, found = lookup[v.Value]; !found {
				m := jo.NewOrderedMap()
				m.Set("abilityId", v.Value)
				m.Set("contentId", v.Value+offset)
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
	var sl []interface{}
	for _, e := range pr.Espers {
		if e.Checked {
			sl = append(sl, e.Value)
		}
	}
	return p.setTarget(p.UserData, OwnedMagicStoneList, sl)
}

func (p *PR) saveInventory(baseKey string, sortKey string, inventory *pri.Inventory, addedItems []int) (err error) {
	var (
		rows             = inventory.GetRows()
		sl               = make([]interface{}, 0, len(rows))
		b                []byte
		slTarget         = jo.NewOrderedMap()
		found            = make(map[int]bool)
		addedCountLookup = make(map[int]int)
	)

	for _, r := range rows {
		if inventory.RemoveDuplicates {
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

	for _, i := range addedItems {
		addedCountLookup[i]++
	}
	for k, v := range addedCountLookup {
		if b, err = json.Marshal(&pri.Row{
			ItemID: k,
			Count:  v,
		}); err != nil {
			return
		}
		sl = append(sl, string(b))
	}

	slTarget.Set(targetKey, sl)
	if err = p.marshalTo(p.UserData, baseKey, slTarget); err != nil {
		return
	}

	if pri.GetInventory().ResetSortOrder && sortKey != "" {
		slTarget = jo.NewOrderedMap()
		slTarget.Set(targetKey, make([]interface{}, 0))
		if err = p.marshalTo(p.UserData, sortKey, slTarget); err != nil {
			return
		}
	}
	return
}

func (p *PR) saveMiscStats() (err error) {
	if err = p.setValue(p.UserData, OwnedGil, models.GetMisc().GP); err != nil {
		return
	}
	if err = p.setValue(p.UserData, Steps, models.GetMisc().Steps); err != nil {
		return
	}
	if err = p.setValue(p.UserData, EscapeCount, models.GetMisc().EscapeCount); err != nil {
		return
	}
	if err = p.setValue(p.UserData, BattleCount, models.GetMisc().BattleCount); err != nil {
		return
	}
	if err = p.setValue(p.UserData, SaveCompleteCount, models.GetMisc().NumberOfSaves); err != nil {
		return
	}
	if err = p.setValue(p.UserData, MonstersKilledCount, models.GetMisc().MonstersKilledCount); err != nil {
		return
	}
	if ds, ok := p.Base.GetValue(DataStorage); ok {
		m := jo.NewOrderedMap()
		if err = m.UnmarshalJSON([]byte(ds.(string))); err != nil {
			return
		}
		if err = p.SetIntInSlice(m, "global", models.GetMisc().CursedShieldFightCount); err != nil {
			return
		}
		var b []byte
		if b, err = m.MarshalJSON(); err != nil {
			return
		}
		p.Base.Set(DataStorage, string(b))
	}
	return
}

func (p *PR) saveTransportation() (err error) {
	v := make([]interface{}, len(pri.Transportations))
	for i, t := range pri.Transportations {
		om := jo.NewOrderedMap()
		pos := jo.NewOrderedMap()
		pos.Set("x", t.Position.X)
		pos.Set("y", t.Position.Y)
		pos.Set("z", t.Position.Z)
		om.Set(TransPosition, pos)
		om.Set(TransDirection, t.Direction)
		om.Set(TransID, t.ID)
		mapID := t.MapID
		if t.ForcedDisabled {
			mapID = -1
		}
		om.Set(TransMapID, mapID)
		om.Set(TransEnable, false)
		ts := t.TimeStampTicks
		if t.ForcedEnabled && ts == 0 {
			ts = global.NowToTicks()
		}
		om.Set(TransTimeStampTicks, ts)
		var b []byte
		if b, err = om.MarshalJSON(); err != nil {
			return
		}
		v[i] = string(b)
	}
	return p.setTarget(p.UserData, OwnedTransportationList, v)
}

func (p *PR) saveMapData() (err error) {
	md := pri.GetMapData()
	if err = p.setValue(p.MapData, MapID, md.MapID); err != nil {
		return
	}
	if err = p.setValue(p.MapData, PointIn, md.PointIn); err != nil {
		return
	}
	if err = p.setValue(p.MapData, TransportationID, md.TransportationID); err != nil {
		return
	}
	if err = p.setValue(p.MapData, CarryingHoverShip, md.CarryingHoverShip); err != nil {
		return
	}
	if err = p.setValue(p.MapData, PlayableCharacterCorpsId, md.PlayableCharacterCorpsID); err != nil {
		return
	}

	pe := jo.NewOrderedMap()
	pos := jo.NewOrderedMap()
	pos.Set("x", md.Player.X)
	pos.Set("y", md.Player.Y)
	pos.Set("z", md.Player.Z)
	pe.Set(PlayerPosition, pos)
	pe.Set(PlayerDirection, md.PlayerDirection)
	if err = p.marshalTo(p.MapData, PlayerEntity, pe); err != nil {
		return
	}

	gps := jo.NewOrderedMap()
	gps.Set(GpsDataMapID, md.Gps.MapID)
	gps.Set(GpsDataAreaID, md.Gps.AreaID)
	gps.Set(GpsDataID, md.Gps.GpsID)
	gps.Set(GpsDataWidth, md.Gps.Width)
	gps.Set(GpsDataHeight, md.Gps.Height)
	if err = p.marshalTo(p.MapData, GpsData, gps); err != nil {
		return
	}
	return
}

func (p *PR) saveVeldt() (err error) {
	var (
		veldt = pri.GetVeldt()
		set   = make([]int, len(veldt.Encounters))
	)
	for i, v := range veldt.Encounters {
		if v {
			set[i] = 1
		}
	}
	err = p.setValue(p.MapData, BeastFieldEncountExchangeFlags, set)
	return
}

func (p *PR) saveCheats() (err error) {
	c := pri.GetCheats()
	if err = p.setValue(p.UserData, OpenChestCount, c.OpenedChestCount); err != nil {
		return
	}
	if err = p.setFlag(p.Base, IsCompleteFlag, c.IsCompleteFlag); err != nil {
		return
	}
	err = p.setValue(p.UserData, PlayTime, c.PlayTime)
	return
}

func (p *PR) SetIntInSlice(to *jo.OrderedMap, key string, value int) (err error) {
	var (
		i, ok = to.GetValue(key)
		sl    []interface{}
	)
	if !ok {
		err = fmt.Errorf("unable to find %s", key)
		return
	}

	if sl, ok = i.([]interface{}); !ok || len(sl) < 9 {
		err = fmt.Errorf("unable to load cursed shield battle count")
		return
	}
	sl[9] = value
	to.Set("global", sl)
	return
}

func (p *PR) setValue(to *jo.OrderedMap, key string, value interface{}) (err error) {
	if !to.Has(key) {
		err = fmt.Errorf("unable to find %s", key)
	}
	to.Set(key, value)
	return
}

func (p *PR) setFlag(to *jo.OrderedMap, key string, value bool) error {
	var i int
	if value {
		i = 1
	}
	return p.setValue(to, key, i)
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

func (p *PR) getInvCount(eq *[]string, counts map[int]int, addedItems *[]int, id int, emptyID int) {
	var i idCount
	if id == 0 {
		i.ContentID = emptyID
		i.Count = counts[emptyID]
	}

	if count, ok := counts[id]; ok {
		i.ContentID = id
		i.Count = count
	} else {
		// *addedItems = append(*addedItems, id)
		i.ContentID = id
		i.Count = 1
	}
	b, _ := json.Marshal(&i)
	*eq = append(*eq, string(b))
}

func (p *PR) setTarget(d *jo.OrderedMap, key string, value []interface{}) (err error) {
	var (
		t = jo.NewOrderedMap()
		b []byte
	)
	if value != nil {
		t.Set(targetKey, value)
	} else {
		t.Set(targetKey, make([]interface{}, 0))
	}
	b, err = t.MarshalJSON()
	return p.setValue(d, key, string(b))
}

func (p *PR) revertUnicodeNames(b []byte) []byte {
	s := string(b)
	for _, r := range p.names {
		s = strings.Replace(s, r.Replaced, r.Original, 1)
	}
	return []byte(s)
	// strconv.Unquote(strings.Replace(strconv.Quote(string(original)), `\\x`, `\x`, -1));
	/*i := 0
	for j := 0; j < len(p.names); j++ {
		original := p.names[j].Original
		replaced := p.names[j].Replaced
		for ; i < len(b)-10; i++ {
			if b[i] == replaced[0] {
				matched := true
				for k := 1; k < len(replaced); k++ {
					if b[i+k] != replaced[k] {
						matched = false
						break
					}
				}
				if matched {
					for k := 0; k < len(replaced); k++ {
						b[i+k] = original[k]
					}
					i += len(replaced)
					break
				}
			}
		}
	}
	return b*/
}
