package pr

import "ffvi_editor/models/consts"

var (
	Espers = []*consts.NameValueChecked{
		consts.NewNameValueChecked("Ramuh", 62),
		consts.NewNameValueChecked("Kirin", 63),
		consts.NewNameValueChecked("Siren", 64),
		consts.NewNameValueChecked("Cait Sith", 65),
		consts.NewNameValueChecked("Ifrit", 66),
		consts.NewNameValueChecked("Shiva", 67),
		consts.NewNameValueChecked("Unicorn", 68),
		consts.NewNameValueChecked("Maduin", 69),
		consts.NewNameValueChecked("Catoblepas", 70),
		consts.NewNameValueChecked("Phantom", 71),
		consts.NewNameValueChecked("Carbuncle", 72),
		consts.NewNameValueChecked("Bismark", 73),
		consts.NewNameValueChecked("Golem", 74),
		consts.NewNameValueChecked("Zona Seeker", 75),
		consts.NewNameValueChecked("Seraph", 76),
		consts.NewNameValueChecked("Quetzalli", 77),
		consts.NewNameValueChecked("Fenrir", 78),
		consts.NewNameValueChecked("Valigarmanda", 79),
		consts.NewNameValueChecked("Midgardsormr", 80),
		consts.NewNameValueChecked("Lakshmi", 81),
		consts.NewNameValueChecked("Alexander", 82),
		consts.NewNameValueChecked("Phoenix", 83),
		consts.NewNameValueChecked("Odin", 84),
		consts.NewNameValueChecked("Bahamut", 85),
		consts.NewNameValueChecked("Ragnarok", 86),
		consts.NewNameValueChecked("Crusader", 87),
		consts.NewNameValueChecked("Raiden", 88),
	}
	SortedEspers  = make([]*consts.NameValueChecked, 0, len(Espers))
	EspersByValue = make(map[int]*consts.NameValueChecked)
)

func init() {
	SortedEspers = consts.SortByNameChecked(Espers)
	for _, i := range Espers {
		EspersByValue[i.Value] = i
	}
}
