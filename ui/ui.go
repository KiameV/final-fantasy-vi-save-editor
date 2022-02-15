package ui

import "github.com/aarzilli/nucular"

type UI interface {
	Draw(w *nucular.Window)
}

func InitReadOnlyText(tb *nucular.TextEditor, text string) {
	tb.Flags = nucular.EditBox | nucular.EditMultiline
	tb.SingleLine = false
	tb.Text([]rune(text))
	tb.Flags |= nucular.EditReadOnly
	tb.Flags |= nucular.EditNoHorizontalScroll
}
