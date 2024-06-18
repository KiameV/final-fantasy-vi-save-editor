package loaders

import (
	jo "gitlab.com/c0b/go-ordered-json"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/save/util"
)

func Misc(data *save.Data) (m *core.Misc, err error) {
	m = core.NewMisc()
	if m.GP, err = util.GetInt(data.UserData, util.OwnedGil); err != nil {
		return
	}
	if m.Steps, err = util.GetInt(data.UserData, util.Steps); err != nil {
		return
	}
	if m.EscapeCount, err = util.GetInt(data.UserData, util.EscapeCount); err != nil {
		return
	}
	if m.BattleCount, err = util.GetInt(data.UserData, util.BattleCount); err != nil {
		return
	}
	if m.NumberOfSaves, err = util.GetInt(data.UserData, util.SaveCompleteCount); err != nil {
		return
	}

	if ds, ok := data.Base.GetValue(util.DataStorage); ok {
		jm := jo.NewOrderedMap()
		if err = jm.UnmarshalJSON([]byte(ds.(string))); err != nil {
			return
		}
		if m.CursedShieldFightCount, err = util.GetIntFromSlice(jm, "global"); err != nil {
			return
		}
	}

	if m.OpenedChestCount, err = util.GetInt(data.UserData, util.OpenChestCount); err != nil {
		return
	}
	if m.IsCompleteFlag, err = util.GetFlag(data.Base, util.IsCompleteFlag); err != nil {
		return
	}
	if m.PlayTime, err = util.GetFloat(data.UserData, util.PlayTime); err != nil {
		return
	}

	if data.Game.IsOne() {
		var sl interface{}
		if sl, err = util.FromTarget(data.UserData, util.OwnendCrystalFlags); err != nil {
			return
		}
		for i, j := range sl.([]interface{}) {
			m.OwnedCrystals[i] = j.(bool)
		}
	}

	return
}
