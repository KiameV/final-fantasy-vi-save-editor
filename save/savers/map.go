package savers

import (
	jo "gitlab.com/c0b/go-ordered-json"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/save/util"
)

func MapData(data *save.Data) (err error) {
	m := data.Save.Map
	if err = util.SetValue(data.MapData, util.MapID, m.MapID); err != nil {
		return
	}
	if err = util.SetValue(data.MapData, util.PointIn, m.PointIn); err != nil {
		return
	}
	if err = util.SetValue(data.MapData, util.TransportationID, m.TransportationID); err != nil {
		return
	}
	if err = util.SetValue(data.MapData, util.CarryingHoverShip, m.CarryingHoverShip); err != nil {
		return
	}
	if data.Game.IsSix() {
		if err = util.SetValue(data.MapData, util.PlayableCharacterCorpsId, m.PlayableCharacterCorpsID); err != nil {
			return
		}
	}

	pe := jo.NewOrderedMap()
	pos := jo.NewOrderedMap()
	pos.Set("x", m.Player.X)
	pos.Set("y", m.Player.Y)
	pos.Set("z", m.Player.Z)
	pe.Set(util.PlayerPosition, pos)
	pe.Set(util.PlayerDirection, m.PlayerDirection)
	if err = util.MarshalTo(data.MapData, util.PlayerEntity, pe); err != nil {
		return
	}

	gps := jo.NewOrderedMap()
	gps.Set(util.GpsDataMapID, m.Gps.MapID)
	gps.Set(util.GpsDataAreaID, m.Gps.AreaID)
	gps.Set(util.GpsDataID, m.Gps.GpsID)
	gps.Set(util.GpsDataWidth, m.Gps.Width)
	gps.Set(util.GpsDataHeight, m.Gps.Height)
	if err = util.MarshalTo(data.MapData, util.GpsData, gps); err != nil {
		return
	}
	return
}
