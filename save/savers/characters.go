package savers

import (
	jo "gitlab.com/c0b/go-ordered-json"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/save/util"
)

func Character(data *save.Data) (err error) {
	var id int
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

		if len(c.Equipment) > 0 {
			if err = saveEquipment(d, c); err != nil {
				return
			}
		}
	}
	return
}

func saveEquipment(d *jo.OrderedMap, c *core.Character) (err error) {
	// eq := jo.NewOrderedMap()
	// if err = util.UnmarshalFrom(d, util.EquipmentList, eq); err != nil {
	// 	return
	// }
	// invCounts := pri.GetInventory().GetItemLookup()
	// var eqIDCounts []string
	// util.InvCount(&eqIDCounts, invCounts, addedItems, c.Equipment.WeaponID, 93)
	// util.InvCount(&eqIDCounts, invCounts, addedItems, c.Equipment.ShieldID, 93)
	// util.InvCount(&eqIDCounts, invCounts, addedItems, c.Equipment.ArmorID, 199)
	// util.InvCount(&eqIDCounts, invCounts, addedItems, c.Equipment.HelmetID, 198)
	// util.InvCount(&eqIDCounts, invCounts, addedItems, c.Equipment.Relic1ID, 200)
	// util.InvCount(&eqIDCounts, invCounts, addedItems, c.Equipment.Relic2ID, 200)
	// eq.Set("values", eqIDCounts)
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

type equipment struct {
	Keys   []int            `json:"keys"`
	Values []*jo.OrderedMap `json:"values"`
}
