package savers

import (
	jo "gitlab.com/c0b/go-ordered-json"
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/save/util"
)

func Transportation(data *save.Data) (err error) {
	forms := data.Save.Transportations.Forms
	v := make([]any, len(forms))
	for i, t := range forms {
		om := jo.NewOrderedMap()
		pos := jo.NewOrderedMap()
		pos.Set("x", t.Position.X)
		pos.Set("y", t.Position.Y)
		pos.Set("z", t.Position.Z)
		om.Set(util.TransPosition, pos)
		om.Set(util.TransDirection, t.Direction)
		om.Set(util.TransID, t.ID)
		mapID := t.MapID
		if t.ForcedDisabled {
			mapID = -1
		}
		om.Set(util.TransMapID, mapID)
		om.Set(util.TransEnable, false)
		ts := t.TimeStampTicks
		if t.ForcedEnabled && ts == 0 {
			ts = global.NowToTicks()
		}
		om.Set(util.TransTimeStampTicks, ts)
		var b []byte
		if b, err = om.MarshalJSON(); err != nil {
			return
		}
		v[i] = string(b)
	}
	return util.SetTarget(data.UserData, util.OwnedTransportationList, v)
}
