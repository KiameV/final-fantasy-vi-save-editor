package loaders

import (
	"errors"

	jo "gitlab.com/c0b/go-ordered-json"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/save/util"
)

func Transportation(data *save.Data) (tr *core.Transportations, err error) {
	var sl interface{}
	if sl, err = util.FromTarget(data.UserData, util.OwnedTransportationList); err != nil {
		return
	}
	tr = core.NewTransportations(len(sl.([]interface{})))
	for index, i := range sl.([]interface{}) {
		om := jo.NewOrderedMap()
		if err = om.UnmarshalJSON([]byte(i.(string))); err != nil {
			return
		}
		t := &core.Transportation{}
		if t.ID, err = util.GetInt(om, util.TransID); err != nil {
			return
		}
		if t.MapID, err = util.GetInt(om, util.TransMapID); err != nil {
			return
		}
		if t.Direction, err = util.GetInt(om, util.TransDirection); err != nil {
			return
		}
		if t.TimeStampTicks, err = util.GetUint(om, util.TransTimeStampTicks); err != nil {
			return
		}

		var pos *jo.OrderedMap
		if pos = om.Get(util.TransPosition).(*jo.OrderedMap); pos == nil {
			err = errors.New("unable to util.Get transportation position")
			return
		}
		if t.Position.X, err = util.GetFloat(pos, "x"); err != nil {
			return
		}
		if t.Position.Y, err = util.GetFloat(pos, "y"); err != nil {
			return
		}
		if t.Position.Z, err = util.GetFloat(pos, "z"); err != nil {
			return
		}

		t.Enabled = t.TimeStampTicks > 0 && t.MapID > 0 && t.Position.X > 0 && t.Position.Y > 0 && t.Position.Z > 0

		tr.Add(index, t)
	}
	return
}
