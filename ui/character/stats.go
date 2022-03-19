package character

import (
	"ffvi_editor/global"
	"ffvi_editor/models"
	"ffvi_editor/models/consts/snes"
	"ffvi_editor/ui/file"
	"ffvi_editor/ui/widgets"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/rect"
)

var (
	minValue   = [2]int{1, 0}
	vigorLabel = [2]string{"Vigor", "Strength"}
	speedLabel = [2]string{"Speed", "Agility"}
)

type statsUI struct {
	name       nucular.TextEditor
	levelExpTB nucular.TextEditor
	yLast      int
}

func newStatsUI() widget {
	u := &statsUI{}

	u.name.Flags = nucular.EditField
	u.name.Maxlen = 8
	u.name.SingleLine = true
	u.name.Text([]rune(snes.Characters[0]))

	widgets.InitReadOnlyText(&u.levelExpTB, levelText)
	return u
}

func (u *statsUI) Draw(w *nucular.Window) {
	var y = 0
	w.Row(u.yLast).SpaceBegin(23)
	w.LayoutSpacePush(rect.Rect{
		X: 0,
		Y: 0,
		W: 50,
		H: 22,
	})
	if !global.IsShowingPR() || (file.PrIO != nil && !file.PrIO.HasUnicodeNames()) {
		w.Label("Name:", "LC")
		w.LayoutSpacePush(rect.Rect{
			X: 60,
			Y: 0,
			W: 100,
			H: 22,
		})
		u.name.Edit(w)
		t := string(u.name.Buffer)
		if t != character.Name {
			character.Name = t
		}
		y += 24
	}

	addPropertyInt(w, 0, y, "Level", &character.Level, 1, 99, 1)
	addPropertyInt(w, 180, y, "Exp", &character.Exp, 1, 2637112, 1000)
	y += 26

	addPropertyInt(w, 0, y, "Current HP", &character.HP.Current, 0, 9999, 100)
	addPropertyInt(w, 180, y, "Max HP", &character.HP.Max, 1, 9999, 100)
	y += 26

	addPropertyInt(w, 0, y, "Current MP", &character.MP.Current, 0, 999, 10)
	addPropertyInt(w, 180, y, "Max MP", &character.MP.Max, 1, 999, 10)
	y += 26

	i := 0
	if global.IsShowing(global.ShowPR) {
		i = 1
	}

	addPropertyInt(w, 0, y, vigorLabel[i], &character.Vigor, minValue[i], 255, 1)
	addPropertyInt(w, 180, y, "Stamina", &character.Stamina, minValue[i], 255, 1)
	y += 26

	addPropertyInt(w, 0, y, speedLabel[i], &character.Speed, minValue[i], 255, 1)
	addPropertyInt(w, 180, y, "Magic", &character.Magic, minValue[i], 255, 1)
	y += 26

	w.LayoutSpacePush(rect.Rect{
		X: 360,
		Y: 0,
		W: 200,
		H: u.yLast,
	})
	u.levelExpTB.Edit(w)

	character.HP.Fix()
	character.MP.Fix()

	u.yLast = y
}

func (u *statsUI) Update(character *models.Character) {
	u.name.SelectAll()
	u.name.DeleteSelection()
	u.name.Text([]rune(character.Name))
}

func addPropertyInt(w *nucular.Window, x int, y int, name string, value *int, min int, max int, step int) {
	w.LayoutSpacePush(rect.Rect{
		X: x,
		Y: y,
		W: 160,
		H: 24,
	})
	w.PropertyInt(name, min, value, max, step, 0)
}

const levelText = `Level - Experience
01 - 0
02 - 32
03 - 96
04 - 208
05 - 400
06 - 672
07 - 1056
08 - 1552
09 - 2184
10 - 2976
11 - 3936
12 - 5080
13 - 6432
14 - 7992
15 - 9784
16 - 11840
17 - 14152
18 - 16736
19 - 19616
20 - 22832
21 - 26360
22 - 30232
23 - 24456
24 - 39056
25 - 44072
26 - 49464
27 - 55288
28 - 61568
29 - 68304
30 - 75496
31 - 93184
32 - 91384
33 - 100083
34 - 108344
35 - 119136
36 - 129504
37 - 140464
38 - 152008
39 - 164184
40 - 176976
41 - 190416
42 - 204520
43 - 219320
44 - 234808
45 - 251000
46 - 267936
47 - 285600
48 - 304040
49 - 323248
50 - 343248
51 - 364064
52 - 385696
53 - 408160
54 - 431488
55 - 455680
56 - 480776
57 - 506760
58 - 533680
59 - 561528
60 - 590320
61 - 620096
62 - 650840
63 - 682600
64 - 715368
65 - 749160
66 - 784016
67 - 819920
68 - 856920
69 - 895016
70 - 934208
71 - 974536
72 - 1016000
73 - 1058640
74 - 1102456
75 - 1147456
76 - 1193648
77 - 1241080
78 - 1289744
79 - 1339672
80 - 1390872
81 - 1443368
82 - 1497160
83 - 1553364
84 - 1608712
85 - 1666512
86 - 1725688
87 - 1786240
88 - 1848184
89 - 1911552
90 - 1976352
91 - 2042608
92 - 2110320
93 - 2179504
94 - 2250192
95 - 2322392
96 - 2396128
97 - 2471400
98 - 2548224
99 - 2637112`
