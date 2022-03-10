package widgets

import (
	"ffvi_editor/models/consts/snes"
	"fmt"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/rect"
	"sort"
	"strings"
)

var (
	//id       nucular.TextEditor
	//idResult string

	name       nucular.TextEditor
	prevName   string
	nameResult []string
)

func init() {
	/*id.Flags = nucular.EditField
	id.Maxlen = 2
	id.SingleLine = true*/

	name.Flags = nucular.EditField
	name.Maxlen = 8
	name.SingleLine = true
}

func DrawItemFinder(w *nucular.Window, x, y int) (count int) {
	count = 6
	/*w.LayoutSpacePush(rect.Rect{
		X: x,
		Y: y,
		W: 80,
		H: 18,
	})
	w.Label("Find By ID:", "LC")
	y += 24

	w.LayoutSpacePush(rect.Rect{
		X: x,
		Y: y,
		W: 40,
		H: 22,
	})
	if e := DrawAndValidateHexInput(w, &id); e == nucular.EditActive || e == nucular.EditCommitted {
		if len(id.Buffer) == 2 {
			idResult = consts.ItemsByID[string(id.Buffer)]
		} else {
			idResult = ""
		}
	}
	w.LayoutSpacePush(rect.Rect{
		X: x + 50,
		Y: y,
		W: 100,
		H: 22,
	})
	w.Label(idResult, "LC")
	y += 24
	*/

	w.LayoutSpacePush(rect.Rect{
		X: x,
		Y: y,
		W: 120,
		H: 18,
	})
	w.Label("Find By Name:", "LC")
	y += 24

	w.LayoutSpacePush(rect.Rect{
		X: x,
		Y: y,
		W: 80,
		H: 22,
	})
	if e := name.Edit(w); e == nucular.EditActive || e == nucular.EditCommitted {
		l := len(name.Buffer)
		if l == 0 || l >= 2 {
			nameResult = nameResult[:0]
		}
		if l >= 2 {
			s := strings.ToLower(string(name.Buffer))
			for n, v := range snes.ItemsByName {
				if strings.Index(strings.ToLower(n), s) != -1 {
					nameResult = append(nameResult, fmt.Sprintf("%s - %s", v, n))
					count++
				}
			}
			sort.Strings(nameResult)
		}
	}
	w.LayoutSpacePush(rect.Rect{
		X: x + 90,
		Y: y,
		W: 40,
		H: 22,
	})
	y += 24

	for _, s := range nameResult {
		w.LayoutSpacePush(rect.Rect{
			X: x + 5,
			Y: y,
			W: 150,
			H: 22,
		})
		w.Label(s, "LC")
		y += 24
	}
	return
}
