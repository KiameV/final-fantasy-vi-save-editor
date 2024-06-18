package savers

import (
	jo "gitlab.com/c0b/go-ordered-json"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/save/util"
)

func Misc(data *save.Data) (err error) {
	misc := data.Save.Misc
	if err = util.SetValue(data.UserData, util.OwnedGil, misc.GP); err != nil {
		return
	}
	if err = util.SetValue(data.UserData, util.Steps, misc.Steps); err != nil {
		return
	}
	if err = util.SetValue(data.UserData, util.EscapeCount, misc.EscapeCount); err != nil {
		return
	}
	if err = util.SetValue(data.UserData, util.BattleCount, misc.BattleCount); err != nil {
		return
	}
	if err = util.SetValue(data.UserData, util.SaveCompleteCount, misc.NumberOfSaves); err != nil {
		return
	}

	if ds, ok := data.Base.GetValue(util.DataStorage); ok {
		m := jo.NewOrderedMap()
		if err = m.UnmarshalJSON([]byte(ds.(string))); err != nil {
			return
		}
		if err = util.SetIntInSlice(m, "global", misc.CursedShieldFightCount); err != nil {
			return
		}
		var b []byte
		if b, err = m.MarshalJSON(); err != nil {
			return
		}
		data.Base.Set(util.DataStorage, string(b))
	}

	if err = util.SetValue(data.UserData, util.OpenChestCount, misc.OpenedChestCount); err != nil {
		return
	}
	if err = util.SetFlag(data.Base, util.IsCompleteFlag, misc.IsCompleteFlag); err != nil {
		return
	}
	if data.Game.IsOne() {
		sl := make([]interface{}, len(misc.OwnedCrystals))
		for i, b := range misc.OwnedCrystals {
			sl[i] = b
		}
		if err = util.SetTarget(data.UserData, util.OwnendCrystalFlags, sl); err != nil {
			return
		}
	}
	return util.SetValue(data.UserData, util.PlayTime, misc.PlayTime)
}
