package consts

import (
	"bufio"
	"strconv"
	"strings"
)

var (
	LevelToExp = make([]uint, 100)
)

func init() {
	var (
		sr        = strings.NewReader(LevelText)
		br        = bufio.NewReader(sr)
		b, _, err = br.ReadLine()
		i         = 1
	)
	LevelToExp = make([]uint, 100)
	for err == nil && len(b) > 0 && i <= 99 {
		b, _, err = br.ReadLine()
		if i == 1 {
			i++
			continue
		}
		s := string(b)
		if strings.Contains(s, " - ") {
			sl := strings.Split(s, " - ")
			j, _ := strconv.ParseInt(sl[1], 10, 64)
			LevelToExp[i] = uint(j)
			i++
		}
	}
	i = 0
	i++
}

const LevelText = `Level - Experience
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
