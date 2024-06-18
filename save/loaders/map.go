package loaders

import (
	"errors"

	jo "gitlab.com/c0b/go-ordered-json"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/save/util"
)

func MapData(data *save.Data) (md *core.MapData, err error) {
	md = core.NewMapData()
	if md.MapID, err = util.GetInt(data.MapData, util.MapID); err != nil {
		return
	}
	if md.PointIn, err = util.GetInt(data.MapData, util.PointIn); err != nil {
		return
	}
	if md.TransportationID, err = util.GetInt(data.MapData, util.TransportationID); err != nil {
		return
	}
	if md.CarryingHoverShip, err = util.GetBool(data.MapData, util.CarryingHoverShip); err != nil {
		return
	}
	if md.PlayableCharacterCorpsID, err = util.GetInt(data.MapData, util.PlayableCharacterCorpsId); err != nil {
		return
	}

	pe := jo.NewOrderedMap()
	if err = util.UnmarshalFrom(data.MapData, util.PlayerEntity, pe); err != nil {
		return
	}
	var pos *jo.OrderedMap
	if pos = pe.Get(util.PlayerPosition).(*jo.OrderedMap); pos == nil {
		err = errors.New("unable to get transportation position")
		return
	}
	if md.Player.X, err = util.GetFloat(pos, "x"); err != nil {
		return
	}
	if md.Player.Y, err = util.GetFloat(pos, "y"); err != nil {
		return
	}
	if md.Player.Z, err = util.GetFloat(pos, "z"); err != nil {
		return
	}
	if md.PlayerDirection, err = util.GetInt(pe, util.PlayerDirection); err != nil {
		return
	}

	gps := jo.NewOrderedMap()
	if err = util.UnmarshalFrom(data.MapData, util.GpsData, gps); err != nil {
		return
	}
	if md.Gps.MapID, err = util.GetInt(gps, util.GpsDataMapID); err != nil {
		return
	}
	if md.Gps.AreaID, err = util.GetInt(gps, util.GpsDataAreaID); err != nil {
		return
	}
	if md.Gps.GpsID, err = util.GetInt(gps, util.GpsDataID); err != nil {
		return
	}
	if md.Gps.Width, err = util.GetInt(gps, util.GpsDataWidth); err != nil {
		return
	}
	if md.Gps.Height, err = util.GetInt(gps, util.GpsDataHeight); err != nil {
		return
	}
	return
}
