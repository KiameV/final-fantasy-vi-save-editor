package pr

import "ffvi_editor/models/consts"

const (
	BushidoFrom   int64 = 124
	BushidoTo     int64 = 131
	BushidoOffset       = 330
)

var (
	Bushidos = []*consts.NameValueChecked{
		consts.NewNameValueChecked("Fang", 124),
		consts.NewNameValueChecked("Sky", 125),
		consts.NewNameValueChecked("Tiger", 126),
		consts.NewNameValueChecked("Flurry", 127),
		consts.NewNameValueChecked("Dragon", 128),
		consts.NewNameValueChecked("Eclipse", 129),
		consts.NewNameValueChecked("Tempest", 130),
		consts.NewNameValueChecked("Oblivion", 131),
	}
	BushidoLookupByID = make(map[int]*consts.NameValueChecked)
)

func init() {
	for _, b := range Bushidos {
		BushidoLookupByID[b.Value] = b
	}
}
