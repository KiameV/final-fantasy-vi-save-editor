package pr

import "ffvi_editor/models/consts"

const (
	DanceFrom   int64 = 164
	DanceTo     int64 = 171
	DanceOffset       = 330
)

var (
	Dances = []*consts.NameValueChecked{
		consts.NewNameValueChecked("Wind Rhapsody", 164),
		consts.NewNameValueChecked("Forest Nocturne", 165),
		consts.NewNameValueChecked("Desert Lullaby", 166),
		consts.NewNameValueChecked("Love Serenade", 167),
		consts.NewNameValueChecked("Earth Blues", 168),
		consts.NewNameValueChecked("Water Harmony", 169),
		consts.NewNameValueChecked("Twilight Requiem", 170),
		consts.NewNameValueChecked("Snowman Rondo", 171),
	}
	DanceLookupByID = make(map[int]*consts.NameValueChecked)
)

func init() {
	for _, d := range Dances {
		DanceLookupByID[d.Value] = d
	}
}
