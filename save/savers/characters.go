package savers

import (
	jo "gitlab.com/c0b/go-ordered-json"
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/models"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/save/util"
)

func Character(game global.Game, data *save.Data) (err error) {
	var id int
	preSave(game, data.Save)
	for _, d := range data.Characters {
		if d == nil {
			continue
		}

		if id, err = util.GetInt(d, util.ID); err != nil {
			return
		}

		c, found := data.Save.Characters.GetByID(id)
		if !found {
			continue
		}
		preCharacterSave(game, c)

		if err = util.SetValue(d, util.JobID, c.JobID); err != nil {
			return
		}
		if err = util.SetValue(d, util.Name, c.Name); err != nil {
			return
		}
		if err = util.SetValue(d, util.IsEnableCorps, c.IsEnabled); err != nil {
			return
		}

		params := jo.NewOrderedMap()
		if err = util.UnmarshalFrom(d, util.Parameter, params); err != nil {
			return
		}
		if err = util.SetValue(params, util.AdditionalLevel, c.Level); err != nil {
			return
		}
		if err = util.SetValue(params, util.CurrentHP, c.CurrentHP); err != nil {
			return
		}
		if err = util.SetValue(params, util.AdditionalMaxHp, c.MaxHP); err != nil {
			return
		}
		if err = util.SetValue(params, util.CurrentMP, c.CurrentMP); err != nil {
			return
		}
		if err = util.SetValue(params, util.AdditionalMaxMp, c.MaxMP); err != nil {
			return
		}

		if err = util.SetValue(d, util.CurrentExp, c.Exp); err != nil {
			return
		}
		// TODO Status

		if err = util.SetValue(params, util.AdditionalPower, c.Power); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionalVitality, c.Vitality); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionalAgility, c.Agility); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionalWeight, c.Weight); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionalIntelligence, c.Intelligence); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionalSpirit, c.Spirit); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionAttack, c.Attack); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionalDefence, c.Defense); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionAbilityDefense, c.AbilityDefence); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionAbilityEvasionRate, c.AbilityEvasionRate); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionMagic, c.Magic); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionLuck, c.Luck); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionalAccuracyRate, c.AccuracyRate); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionalEvasionRate, c.EvasionRate); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionAbilityDistrurbedRate, c.AbilityDisturbedRate); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionalCriticalRate, c.CriticalRate); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionalDamageDirmeter, c.DamageDirmeter); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionalAbilityDefenseRate, c.AbilityDefenseRate); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionalAccuracyCount, c.AccuracyCount); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionalEvasionCount, c.EvasionCount); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionalDefenceCount, c.DefenseCount); err != nil {
			return
		}

		if err = util.SetValue(params, util.AdditionalMagicDef, c.MagicDefenseCount); err != nil {
			return
		}

		if err = util.MarshalTo(d, util.Parameter, params); err != nil {
			return
		}

		if err = saveEquipment(d, c); err != nil {
			return
		}

		if err = saveAbilities(game, d, c); err != nil {
			return
		}

		if c.Commands.EnableSave {
			if err = saveCommands(d, c); err != nil {
				return
			}
		}
	}
	return
}

func preSave(game global.Game, save *core.Save) (err error) {
	inv := make(map[int]*core.Row)
	for _, i := range save.Inventory.Rows {
		inv[i.ContentID] = i
	}
	eq := make(map[int][]*models.IdCount)
	for _, c := range save.Characters.All() {
		for _, e := range c.Equipment {
			sl, _ := eq[e.ContentID]
			sl = append(sl, e)
			eq[e.ContentID] = sl
		}
	}
	for k, v := range eq {
		needed := len(v)
		row, ok := inv[k]
		if ok {
			if row.Count < needed {
				row.Count = needed
			}
		} else {
			row = &core.Row{ContentID: k, Count: needed}
			inv[k] = row
			save.Inventory.Add(row)
		}
		for _, e := range v {
			e.Count = row.Count
		}
	}
	return
}

func preCharacterSave(game global.Game, c *core.Character) (err error) {
	if game == global.One {
		owned := make(map[int]bool)
		abilities := make(map[int]bool)
		for _, i := range c.OwnedAbilityIds {
			owned[i] = true
		}
		for _, a := range c.Abilities {
			abilities[a.ContentID] = true
		}
		for i := 0; i < len(c.FF1.TrainedAbilities); i++ {
			for j := 0; j < len(c.FF1.TrainedAbilities[i]); j++ {
				t := c.FF1.TrainedAbilities[i][j]
				if t.ContentID == 0 && t.ID > 0 {
					t.ContentID = t.ID + 208
				}
				if t.ContentID > 0 {
					if _, ok := owned[t.ContentID]; !ok {
						c.OwnedAbilityIds = append(c.OwnedAbilityIds, t.ContentID)
					}
					if _, ok := abilities[t.ContentID]; !ok {
						c.Abilities = append(c.Abilities, &core.Ability{ID: t.ID, ContentID: t.ContentID})
					}
				}
			}
		}
		for _, a := range c.Abilities {
			if a.ID > 0 {
				if _, ok := owned[a.ContentID]; !ok {
					c.OwnedAbilityIds = append(c.OwnedAbilityIds, a.ContentID)
				}
			}
		}
	} else if game == global.Two {
		for _, a := range c.Abilities {
			if a.ContentID == 0 && a.ID > 0 {
				a.ContentID = a.ID + 207
			}
		}
	}
	return
}

func saveEquipment(d *jo.OrderedMap, c *core.Character) (err error) {
	var (
		eq     = jo.NewOrderedMap()
		keys   = make([]int, len(c.Equipment))
		values = make([]string, len(c.Equipment))
		b      []byte
	)
	for i, ic := range c.Equipment {
		keys[i] = i + 1
		m := jo.NewOrderedMap()
		m.Set("contentId", ic.ContentID)
		m.Set("count", ic.Count)
		if b, err = m.MarshalJSON(); err != nil {
			return
		}
		values[i] = string(b)
	}
	eq.Set("keys", keys)
	eq.Set("values", values)
	if err = util.MarshalTo(d, util.EquipmentList, eq); err != nil {
		return
	}
	return
}

func saveAbilities(game global.Game, d *jo.OrderedMap, c *core.Character) (err error) {
	var (
		b         []byte
		abilities = make([]any, 0, len(c.Abilities))
	)
	for _, a := range c.Abilities {
		if a.ID > 0 {
			m := jo.NewOrderedMap()
			if a.ContentID == 0 {
				if game == global.One {
					a.ContentID = a.ID + 208
				}
			}
			m.Set("abilityId", a.ID)
			m.Set("contentId", a.ContentID)
			m.Set("skillLevel", a.SkillLevel)
			if b, err = m.MarshalJSON(); err != nil {
				return
			}
			abilities = append(abilities, string(b))
		}
	}
	if err = util.SetTarget(d, util.AbilityList, abilities); err != nil {
		return
	}
	if game == global.One {
		err = saveAbilitySlotDataFF1(d, c)
	}
	return
}

func saveAbilitySlotDataFF1(d *jo.OrderedMap, c *core.Character) (err error) {
	var b []byte
	slots := make([]any, 8)
	for i, lvl := range c.FF1.TrainedAbilities {
		slotInfo := jo.NewOrderedMap()
		slotInfo.Set("keys", []int{1, 2, 3})
		values := make([]string, 3)
		for j, a := range lvl {
			if a.ID > 0 {
				m := jo.NewOrderedMap()
				m.Set("abilityId", a.ID)
				m.Set("contentId", a.ContentID)
				m.Set("skillLevel", a.SkillLevel)
				if b, err = m.MarshalJSON(); err != nil {
					return
				}
				values[j] = string(b)
			} else {
				values[j] = ""
			}
		}
		slotInfo.Set("values", values)
		if b, err = slotInfo.MarshalJSON(); err != nil {
			return
		}

		level := jo.NewOrderedMap()
		level.Set("level", i+1)
		level.Set("slotInfo", string(b))

		if b, err = level.MarshalJSON(); err != nil {
			return
		}
		slots[i] = string(b)
	}
	return util.SetTarget(d, util.AbilitySlotDataList, slots)
}

func saveCommands(d *jo.OrderedMap, c *core.Character) error {
	return util.SetTarget(d, util.CommandList, c.Commands.All())
}
