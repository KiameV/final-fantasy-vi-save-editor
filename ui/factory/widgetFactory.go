package factory

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func CreateComboInput(label string, items []string, onChange func(string)) *widgets.QGroupBox {
	g := widgets.NewQGroupBox2(label, nil)
	i := widgets.NewQComboBox(nil)
	i.AddItems(items)
	i.ConnectCurrentTextChanged(onChange)
	l := widgets.NewQGridLayout2()
	l.AddWidget3(i, 0, 0, 1, 2, 0)
	g.SetLayout(l)
	return g
}

func CreateTextInput(label string, onChange func(text string)) (*widgets.QLineEdit, *widgets.QGroupBox) {
	g := widgets.NewQGroupBox2(label, nil)
	i := widgets.NewQLineEdit2("", nil)
	i.ConnectTextEdited(onChange)
	l := widgets.NewQGridLayout2()
	l.AddWidget3(i, 0, 0, 1, 2, 0)
	g.SetLayout(l)
	return i, g
}

func CreateUint8Input(label string, onChange func(i int), min, max uint8) (*widgets.QSpinBox, *widgets.QGroupBox) {
	g := widgets.NewQGroupBox2(label, nil)
	i := widgets.NewQSpinBox(nil)
	i.SetMinimum(int(min))
	i.SetMaximum(int(max))
	i.ConnectValueChanged(onChange)
	l := widgets.NewQGridLayout2()
	l.AddWidget3(i, 0, 0, 1, 2, 0)
	g.SetLayout(l)
	return i, g
}

func CreateUint32Input(label string, onChange func(i int), min uint32, max uint32) (*widgets.QSpinBox, *widgets.QGroupBox) {
	g := widgets.NewQGroupBox2(label, nil)
	i := widgets.NewQSpinBox(nil)
	i.SetMinimum(int(min))
	i.SetMaximum(int(max))
	i.ConnectValueChanged(onChange)
	l := widgets.NewQGridLayout2()
	l.AddWidget3(i, 0, 0, 1, 2, 0)
	g.SetLayout(l)
	return i, g
}

func CreateReadOnlyTextBox(text string) *widgets.QTextEdit {
	tb := widgets.NewQTextEdit(nil)
	tb.InsertHtmlDefault(text)
	tb.VerticalScrollBar().Minimum()
	tb.MoveCursor(gui.QTextCursor__Start, gui.QTextCursor__MoveAnchor)
	tb.SetReadOnly(true)
	tb.SetMaximumWidth(200)
	return tb
}
