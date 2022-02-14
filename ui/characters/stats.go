package characters

import (
	f "ffvi_editor/ui/factory"
	"github.com/therecipe/qt/widgets"
)

var (
	name                                                                         *widgets.QLineEdit
	level, exp, currentHP, maxHP, currentMP, maxMP, vigor, stamina, speed, magic *widgets.QSpinBox
)

func createStatsLayout() widgets.QWidget_ITF {
	var (
		g  = widgets.NewQGridLayout2()
		gr *widgets.QGroupBox
	)
	name, gr = f.CreateTextInput("Name", func(value string) { selectedCharacter.Name = value })
	g.AddWidget2(gr, 0, 0, 0)

	level, gr = f.CreateUint8Input("Level", func(value int) { selectedCharacter.Level = uint8(value) }, 1, 99)
	g.AddWidget2(gr, 1, 0, 0)
	exp, gr = f.CreateUint32Input("Exp", func(value int) { selectedCharacter.Exp = uint32(value) }, 0, 2637112)
	g.AddWidget2(gr, 1, 1, 0)

	currentHP, gr = f.CreateUint32Input("Current HP", func(value int) { selectedCharacter.HP.Current = uint32(value) }, 0, 9999)
	g.AddWidget2(gr, 2, 0, 0)
	maxHP, gr = f.CreateUint32Input("Max HP", func(value int) { selectedCharacter.HP.Max = uint32(value) }, 1, 9999)
	g.AddWidget2(gr, 2, 1, 0)

	currentMP, gr = f.CreateUint32Input("Current MP", func(value int) { selectedCharacter.MP.Current = uint32(value) }, 0, 9999)
	g.AddWidget2(gr, 3, 0, 0)
	maxMP, gr = f.CreateUint32Input("Max MP", func(value int) { selectedCharacter.MP.Max = uint32(value) }, 1, 9999)
	g.AddWidget2(gr, 3, 1, 0)

	vigor, gr = f.CreateUint8Input("Vigor", func(value int) { selectedCharacter.Vigor = uint8(value) }, 1, 255)
	g.AddWidget2(gr, 4, 0, 0)
	stamina, gr = f.CreateUint8Input("Stamina", func(value int) { selectedCharacter.Stamina = uint8(value) }, 1, 255)
	g.AddWidget2(gr, 5, 0, 0)
	speed, gr = f.CreateUint8Input("Speed", func(value int) { selectedCharacter.Speed = uint8(value) }, 1, 255)
	g.AddWidget2(gr, 6, 0, 0)
	magic, gr = f.CreateUint8Input("Magic", func(value int) { selectedCharacter.Magic = uint8(value) }, 1, 255)
	g.AddWidget2(gr, 7, 0, 0)

	g.AddWidget3(f.CreateReadOnlyTextBox(levelText), 0, 2, 8, 1, 0)

	w := widgets.NewQWidget(nil, 0)
	w.SetLayout(g)
	return w
}

func updateStats() {
	name.SetText(selectedCharacter.Name)
	level.SetValue(int(selectedCharacter.Level))
	exp.SetValue(int(selectedCharacter.Exp))
	currentHP.SetValue(int(selectedCharacter.HP.Current))
	maxHP.SetValue(int(selectedCharacter.HP.Max))
	currentMP.SetValue(int(selectedCharacter.MP.Current))
	maxMP.SetValue(int(selectedCharacter.MP.Max))
	vigor.SetValue(int(selectedCharacter.Vigor))
	stamina.SetValue(int(selectedCharacter.Stamina))
	speed.SetValue(int(selectedCharacter.Speed))
	magic.SetValue(int(selectedCharacter.Magic))
}

const levelText = `
<b>Level - Experience</b><br>
01 - 0<br>
02 - 32<br>
03 - 96<br>
04 - 208<br>
05 - 400<br>
06 - 672<br>
07 - 1056<br>
08 - 1552<br>
09 - 2184<br>
10 - 2976<br>
11 - 3936<br>
12 - 5080<br>
13 - 6432<br>
14 - 7992<br>
15 - 9784<br>
16 - 11840<br>
17 - 14152<br>
18 - 16736<br>
19 - 19616<br>
20 - 22832<br>
21 - 26360<br>
22 - 30232<br>
23 - 24456<br>
24 - 39056<br>
25 - 44072<br>
26 - 49464<br>
27 - 55288<br>
28 - 61568<br>
29 - 68304<br>
30 - 75496<br>
31 - 93184<br>
32 - 91384<br>
33 - 100083<br>
34 - 108344<br>
35 - 119136<br>
36 - 129504<br>
37 - 140464<br>
38 - 152008<br>
39 - 164184<br>
40 - 176976<br>
41 - 190416<br>
42 - 204520<br>
43 - 219320<br>
44 - 234808<br>
45 - 251000<br>
46 - 267936<br>
47 - 285600<br>
48 - 304040<br>
49 - 323248<br>
50 - 343248<br>
51 - 364064<br>
52 - 385696<br>
53 - 408160<br>
54 - 431488<br>
55 - 455680<br>
56 - 480776<br>
57 - 506760<br>
58 - 533680<br>
59 - 561528<br>
60 - 590320<br>
61 - 620096<br>
62 - 650840<br>
63 - 682600<br>
64 - 715368<br>
65 - 749160<br>
66 - 784016<br>
67 - 819920<br>
68 - 856920<br>
69 - 895016<br>
70 - 934208<br>
71 - 974536<br>
72 - 1016000<br>
73 - 1058640<br>
74 - 1102456<br>
75 - 1147456<br>
76 - 1193648<br>
77 - 1241080<br>
78 - 1289744<br>
79 - 1339672<br>
80 - 1390872<br>
81 - 1443368<br>
82 - 1497160<br>
83 - 1553364<br>
84 - 1608712<br>
85 - 1666512<br>
86 - 1725688<br>
87 - 1786240<br>
88 - 1848184<br>
89 - 1911552<br>
90 - 1976352<br>
91 - 2042608<br>
92 - 2110320<br>
93 - 2179504<br>
94 - 2250192<br>
95 - 2322392<br>
96 - 2396128<br>
97 - 2471400<br>
98 - 2548224<br>
99 - 2637112`
