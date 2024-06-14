package inputs

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

type (
	FloatEntry struct {
		widget.Entry
	}
	FloatEntryBinding struct {
		s binding.String
		i binding.Float
	}
)

func NewFloatEntry() *FloatEntry {
	entry := &FloatEntry{}
	entry.ExtendBaseWidget(entry)
	entry.Entry.Validator = nil
	return entry
}

func NewFloatEntryWithBinding(data FloatEntryBinding) *FloatEntry {
	entry := NewFloatEntry()
	entry.Bind(data.s)
	entry.Entry.Validator = nil
	return entry
}

func NewFloatEntryWithData(f *float64) *FloatEntry {
	entry := NewFloatEntry()
	data := NewFloatEntryBinding(f)
	entry.Bind(data.s)
	entry.Entry.Validator = nil
	return entry
}

func (e *FloatEntry) Float() float64 {
	f, err := strconv.ParseFloat(e.Text, 64)
	if err != nil {
		f = 0
	}
	return f
}

func (e *FloatEntry) SetFloat(f float64) {
	e.Entry.SetText(fmt.Sprintf("%f", f))
}

func (e *FloatEntry) TypedRune(r rune) {
	if (r >= '0' && r <= '9') || r == '.' || r == ',' {
		e.Entry.TypedRune(r)
	}
}

func (e *FloatEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

func (e *FloatEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}

func NewFloatEntryBinding(f *float64) FloatEntryBinding {
	s := binding.NewString()
	_ = s.Set(fmt.Sprintf("%f", *f))
	b := FloatEntryBinding{
		s: s,
		i: binding.BindFloat(f),
	}
	b.s.AddListener(b)
	return b
}

func (b FloatEntryBinding) DataChanged() {
	if s, err := b.s.Get(); err == nil {
		if s == "" {
			_ = b.i.Set(0)
		} else {
			var f float64
			if f, err = strconv.ParseFloat(s, 64); err == nil {
				_ = b.i.Set(f)
			}
		}
	} else {

	}
}

func (b FloatEntryBinding) Set(f float64) {
	_ = b.s.Set(fmt.Sprintf("%f", f))
}
