package main

import (
	"ffvi_editor/ui/characters"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/label"
	"github.com/aarzilli/nucular/style"
	"image/color"
)

var (
	cUI    = characters.NewCharacterUI()
	loaded = true
)

func main() {
	wnd := nucular.NewMasterWindow(0, "Final Fantasy VI Editor", updateWindow)
	wnd.SetStyle(style.FromTable(customTheme, 1.2))
	wnd.Main()
}

func updateWindow(w *nucular.Window) {
	mw := w.Master()

	w.MenubarBegin()
	w.Row(12).Static(45)
	if w := w.Menu(label.TA("File", "LC"), 120, nil); w != nil {
		w.Row(12).Dynamic(1)
		if w.MenuItem(label.TA("Open", "LC")) {
			loaded = true
		}
		w.Row(12).Dynamic(1)
		if w.MenuItem(label.TA("Exit", "LC")) {
			mw.Close()
		}
	}
	w.MenubarEnd()

	w.Row(5).Static(1)
	w.Row(12).Static(200, 200)
	if loaded {
		if w.TreePush(nucular.TreeTab, "Characters", true) {
			cUI.Draw(w)
			w.TreePop()
		}
		if w.TreePush(nucular.TreeTab, "Inventory", false) {
			w.TreePop()
		}
		if w.TreePush(nucular.TreeTab, "Skills", false) {
			w.TreePop()
		}
		if w.TreePush(nucular.TreeTab, "Espers", false) {
			w.TreePop()
		}
		if w.TreePush(nucular.TreeTab, "Misc", false) {
			w.TreePop()
		}
	}

	//w.Row(32).Dynamic(1)
	//if w.ButtonText(fmt.Sprintf("increment: %d", count)) {
	//	count++
	//}
}

var customTheme = style.ColorTable{
	ColorText:                  color.RGBA{0, 0, 0, 255},
	ColorWindow:                color.RGBA{255, 255, 255, 255},
	ColorHeader:                color.RGBA{242, 242, 242, 255},
	ColorHeaderFocused:         color.RGBA{0xc3, 0x9a, 0x9a, 255},
	ColorBorder:                color.RGBA{0, 0, 0, 255},
	ColorButton:                color.RGBA{185, 185, 185, 255},
	ColorButtonHover:           color.RGBA{170, 170, 170, 255},
	ColorButtonActive:          color.RGBA{160, 160, 160, 255},
	ColorToggle:                color.RGBA{150, 150, 150, 255},
	ColorToggleHover:           color.RGBA{120, 120, 120, 255},
	ColorToggleCursor:          color.RGBA{175, 175, 175, 255},
	ColorSelect:                color.RGBA{175, 175, 175, 255},
	ColorSelectActive:          color.RGBA{190, 190, 190, 255},
	ColorSlider:                color.RGBA{190, 190, 190, 255},
	ColorSliderCursor:          color.RGBA{80, 80, 80, 255},
	ColorSliderCursorHover:     color.RGBA{70, 70, 70, 255},
	ColorSliderCursorActive:    color.RGBA{60, 60, 60, 255},
	ColorProperty:              color.RGBA{175, 175, 175, 255},
	ColorEdit:                  color.RGBA{245, 245, 245, 255},
	ColorEditCursor:            color.RGBA{0, 0, 0, 255},
	ColorCombo:                 color.RGBA{225, 225, 225, 255},
	ColorChart:                 color.RGBA{160, 160, 160, 255},
	ColorChartColor:            color.RGBA{45, 45, 45, 255},
	ColorChartColorHighlight:   color.RGBA{255, 0, 0, 255},
	ColorScrollbar:             color.RGBA{180, 180, 180, 255},
	ColorScrollbarCursor:       color.RGBA{140, 140, 140, 255},
	ColorScrollbarCursorHover:  color.RGBA{150, 150, 150, 255},
	ColorScrollbarCursorActive: color.RGBA{160, 160, 160, 255},
	ColorTabHeader:             color.RGBA{210, 210, 210, 255},
}
