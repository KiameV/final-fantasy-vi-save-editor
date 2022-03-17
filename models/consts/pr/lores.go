package pr

import "ffvi_editor/models/consts"

const (
	LoreFrom   int64 = 140
	LoreTo     int64 = 163
	LoreOffset       = 330
)

var (
	Lores = []*consts.NameValueChecked{
		consts.NewNameValueChecked("Doom", 140),
		consts.NewNameValueChecked("Roulette", 141),
		consts.NewNameValueChecked("Tsunami", 142),
		consts.NewNameValueChecked("Aqua Breath", 143),
		consts.NewNameValueChecked("Aero", 144),
		consts.NewNameValueChecked("1000 Needles", 145),
		consts.NewNameValueChecked("Mighty Guard", 146),
		consts.NewNameValueChecked("Revenge Blast", 147),
		consts.NewNameValueChecked("White Wind", 148),
		consts.NewNameValueChecked("Lv. 5 Death", 149),
		consts.NewNameValueChecked("Lv. 4 Flare", 150),
		consts.NewNameValueChecked("Lv. 3 Confuse", 151),
		consts.NewNameValueChecked("Reflect ???", 152),
		consts.NewNameValueChecked("Lv. ? Holy", 153),
		consts.NewNameValueChecked("Traveler", 154),
		consts.NewNameValueChecked("Force Field", 155),
		consts.NewNameValueChecked("Dischord", 156),
		consts.NewNameValueChecked("Bad Breath", 157),
		consts.NewNameValueChecked("Transfusion", 158),
		consts.NewNameValueChecked("Rippler", 159),
		consts.NewNameValueChecked("Stone", 160),
		consts.NewNameValueChecked("Quasar", 161),
		consts.NewNameValueChecked("Grand Delta", 162),
		consts.NewNameValueChecked("Self-Destruct", 163),
	}
	SortedLores    []*consts.NameValueChecked
	LoreLookupByID = make(map[int]*consts.NameValueChecked)
)

func init() {
	SortedLores = consts.SortByNameChecked(Lores)
	for _, l := range Lores {
		LoreLookupByID[l.Value] = l
	}
}
