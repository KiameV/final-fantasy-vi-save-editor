package pr

import "ffvi_editor/models/consts"

const (
	BlitzFrom   int64 = 132
	BlitzTo     int64 = 139
	BlitzOffset       = 330
)

var (
	Blitzes = []*consts.NameValueChecked{
		consts.NewNameValueChecked("Raging Fist", 132),
		consts.NewNameValueChecked("Aura Cannon", 133),
		consts.NewNameValueChecked("Meteor Strike", 134),
		consts.NewNameValueChecked("Rising Phoenix", 135),
		consts.NewNameValueChecked("Chakra", 136),
		consts.NewNameValueChecked("Razor Gale", 137),
		consts.NewNameValueChecked("Soul Spiral", 138),
		consts.NewNameValueChecked("Phantom Rush", 139),
	}
	BlitzLookupByID = make(map[int]*consts.NameValueChecked)
)

func init() {
	for _, b := range Blitzes {
		BlitzLookupByID[b.Value] = b
	}
}
