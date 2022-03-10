package widgets

import (
	"github.com/aarzilli/nucular"
	"strconv"
	"strings"
)

func DrawAndValidateHexInput(w *nucular.Window, tb *nucular.TextEditor) nucular.EditEvents {
	e := tb.Edit(w)
	if e == nucular.EditActive || e == nucular.EditCommitted {
		s := strings.ToUpper(string(tb.Buffer))
		r := fixHex2(s)
		tb.SelectAll()
		tb.DeleteSelection()
		tb.Text(r)
	}
	return e
}

func DrawAndValidateNumberInput(w *nucular.Window, tb *nucular.TextEditor) nucular.EditEvents {
	e := tb.Edit(w)
	if e == nucular.EditActive || e == nucular.EditCommitted {
		s := string(tb.Buffer)
		if s != "" {
			if i, err := strconv.ParseInt(s, 10, 16); err != nil || i < 0 {
				tb.SelectAll()
				tb.DeleteSelection()
				tb.Text([]rune(`0`))
			} else if i > 999 {
				tb.SelectAll()
				tb.DeleteSelection()
				tb.Text([]rune(`999`))
			}
		}
	}
	return e
}

func fixHex2(s string) []rune {
	l := len(s)
	if l == 1 {
		if isValidHexRune([]rune(s)[0]) {
			return []rune(s)
		}
		return make([]rune, 0)
	} else if l == 2 {
		r := []rune(s)
		if isValidHexRune(r[0]) {
			if isValidHexRune(r[1]) {
				return r
			}
			return []rune{r[0]}
		}
		if isValidHexRune(r[1]) {
			return []rune{r[1]}
		}
	}
	return []rune(s)
}

func isValidHexRune(r rune) bool {
	return (r >= '0' && r <= '9') || (r >= 'A' && r <= 'F')
}
