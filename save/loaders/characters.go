package loaders

import (
	"encoding/json"

	jo "gitlab.com/c0b/go-ordered-json"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/save/util"
)

func Characters(data *save.Data, party *core.Party) (characters *core.Characters, err error) {
	characters = core.NewCharacters(len(data.Characters))
	for i, d := range data.Characters {
		if d == nil {
			continue
		}

		c := &core.Character{}
		if err = characters.SetCharacter(i, c); err != nil {
			return
		}

		if c.ID, err = util.GetInt(d, util.ID); err != nil {
			return
		}

		if c.JobID, err = util.GetInt(d, util.JobID); err != nil {
			return
		}

		if c.Name, err = util.GetString(d, util.Name); err != nil {
			return
		}

		party.AddPossibleMember(c)

		if c.IsEnabled, err = util.GetBool(d, util.IsEnableCorps); err != nil {
			return
		}

		params := jo.NewOrderedMap()
		if err = util.UnmarshalFrom(d, util.Parameter, params); err != nil {
			return
		}

		if c.Level, err = util.GetInt(params, util.AdditionalLevel); err != nil {
			return
		}

		if c.CurrentHP, err = util.GetInt(params, util.CurrentHP); err != nil {
			return
		}
		if c.MaxHP, err = util.GetInt(params, util.AdditionalMaxHp); err != nil {
			return
		}

		if c.CurrentMP, err = util.GetInt(params, util.CurrentMP); err != nil {
			return
		}
		if c.MaxMP, err = util.GetInt(params, util.AdditionalMaxMp); err != nil {
			return
		}

		if c.Exp, err = util.GetInt(d, util.CurrentExp); err != nil {
			return
		}

		// TODO Status

		if c.Power, err = util.GetInt(params, util.AdditionalPower); err != nil {
			return
		}

		if c.Vitality, err = util.GetInt(params, util.AdditionalVitality); err != nil {
			return
		}

		if c.Agility, err = util.GetInt(params, util.AdditionalAgility); err != nil {
			return
		}

		if c.Weight, err = util.GetInt(params, util.AdditionalWeight); err != nil {
			return
		}

		if c.Intelligence, err = util.GetInt(params, util.AdditionalIntelligence); err != nil {
			return
		}

		if c.Spirit, err = util.GetInt(params, util.AdditionalSpirit); err != nil {
			return
		}

		if c.Attack, err = util.GetInt(params, util.AdditionAttack); err != nil {
			return
		}

		if c.Defense, err = util.GetInt(params, util.AdditionalDefence); err != nil {
			return
		}

		if c.AbilityDefence, err = util.GetInt(params, util.AdditionAbilityDefense); err != nil {
			return
		}

		if c.AbilityEvasionRate, err = util.GetInt(params, util.AdditionAbilityEvasionRate); err != nil {
			return
		}

		if c.Magic, err = util.GetInt(params, util.AdditionMagic); err != nil {
			return
		}

		if c.Luck, err = util.GetInt(params, util.AdditionLuck); err != nil {
			return
		}

		if c.AccuracyRate, err = util.GetInt(params, util.AdditionalAccuracyRate); err != nil {
			return
		}

		if c.EvasionRate, err = util.GetInt(params, util.AdditionalEvasionRate); err != nil {
			return
		}

		if c.AbilityDisturbedRate, err = util.GetInt(params, util.AdditionAbilityDistrurbedRate); err != nil {
			return
		}

		if c.CriticalRate, err = util.GetInt(params, util.AdditionalCriticalRate); err != nil {
			return
		}

		if c.DamageDirmeter, err = util.GetInt(params, util.AdditionalDamageDirmeter); err != nil {
			return
		}

		if c.AbilityDefenseRate, err = util.GetInt(params, util.AdditionalAbilityDefenseRate); err != nil {
			return
		}

		if c.AccuracyCount, err = util.GetInt(params, util.AdditionalAccuracyCount); err != nil {
			return
		}

		if c.EvasionCount, err = util.GetInt(params, util.AdditionalEvasionCount); err != nil {
			return
		}

		if c.DefenseCount, err = util.GetInt(params, util.AdditionalDefenceCount); err != nil {
			return
		}

		if c.MagicDefenseCount, err = util.GetInt(params, util.AdditionalMagicDef); err != nil {
			return
		}

		if err = loadSpells(d, c); err != nil {
			return
		}
	}
	return
}

func loadSpells(d *jo.OrderedMap, c *core.Character) (err error) {
	var i interface{}
	if i, err = util.FromTarget(d, util.AbilityList); err != nil {
		return nil
	}
	sli := i.([]interface{})
	for j := 0; j < len(sli); j++ {
		m := jo.NewOrderedMap()
		if err = m.UnmarshalJSON([]byte(sli[j].(string))); err != nil {
			return
		}
		c.AddAbility(core.Ability{
			ID:         valueAsInt(m, "abilityId"),
			ContentID:  valueAsInt(m, "contentId"),
			SkillLevel: valueAsInt(m, "skillLevel"),
		})
	}
	return
}

func valueAsInt(om *jo.OrderedMap, key string) int {
	i, _ := om.GetValue(key)
	j, _ := i.(json.Number).Int64()
	return int(j)
}
