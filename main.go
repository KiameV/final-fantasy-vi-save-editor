package main

import (
	"ffvi_editor/io"
	"ffvi_editor/io/save"
	"ffvi_editor/ui/character"
	"ffvi_editor/ui/espers"
	"ffvi_editor/ui/inventory"
	"ffvi_editor/ui/misc"
	"ffvi_editor/ui/skills"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/label"
	"github.com/aarzilli/nucular/style"
	"image"
	"image/color"
)

var (
	cUI    = character.NewUI()
	iUI    = inventory.NewUI()
	sUI    = skills.NewUI()
	eUI    = espers.NewUI()
	mUI    = misc.NewUI()
	loaded = true
)

func main() {
	wnd := nucular.NewMasterWindowSize(0, "Final Fantasy VI Editor", image.Point{X: 725, Y: 500}, updateWindow)
	wnd.SetStyle(style.FromTable(customTheme, 1.2))
	wnd.Main()
}

func updateWindow(w *nucular.Window) {
	mw := w.Master()

	w.MenubarBegin()
	w.Row(12).Static(45)
	if w := w.Menu(label.TA("File", "LC"), 120, nil); w != nil {
		w.Row(12).Dynamic(1)
		if w.MenuItem(label.TA("Open SRM Slot 1", "LC")) {
			if err := io.OpenFile(save.SRMSlot1); err == nil {
				refresh()
			}
		} else if w.MenuItem(label.TA("Open SRM Slot 2", "LC")) {
			if err := io.OpenFile(save.SRMSlot2); err == nil {
				refresh()
			}
		} else if w.MenuItem(label.TA("Open SRM Slot 3", "LC")) {
			if err := io.OpenFile(save.SRMSlot3); err == nil {
				refresh()
			}
		} else if w.MenuItem(label.TA("Open ZNES State", "LC")) {
			if err := io.OpenFile(save.ZnesSaveState); err == nil {
				refresh()
			}
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
			iUI.Draw(w)
			w.TreePop()
		}
		if w.TreePush(nucular.TreeTab, "Skills", false) {
			sUI.Draw(w)
			w.TreePop()
		}
		if w.TreePush(nucular.TreeTab, "Espers", false) {
			eUI.Draw(w)
			w.TreePop()
		}
		if w.TreePush(nucular.TreeTab, "Misc", false) {
			mUI.Draw(w)
			w.TreePop()
		}
	}
}

func refresh() {
	cUI.Refresh()
	iUI.Refresh()
	sUI.Refresh()
	eUI.Refresh()
	mUI.Refresh()
}

var customTheme = style.ColorTable{
	ColorText:                  color.RGBA{0, 0, 0, 255},
	ColorWindow:                color.RGBA{255, 255, 255, 255},
	ColorHeader:                color.RGBA{242, 242, 242, 255},
	ColorHeaderFocused:         color.RGBA{0xc3, 0x9a, 0x9a, 255},
	ColorBorder:                color.RGBA{0, 0, 0, 255},
	ColorButton:                color.RGBA{185, 185, 185, 255},
	ColorButtonHover:           color.RGBA{215, 215, 215, 255},
	ColorButtonActive:          color.RGBA{200, 200, 200, 255},
	ColorToggle:                color.RGBA{225, 225, 225, 255},
	ColorToggleHover:           color.RGBA{200, 200, 200, 255},
	ColorToggleCursor:          color.RGBA{30, 30, 30, 255},
	ColorSelect:                color.RGBA{175, 175, 175, 255},
	ColorSelectActive:          color.RGBA{190, 190, 190, 255},
	ColorSlider:                color.RGBA{190, 190, 190, 255},
	ColorSliderCursor:          color.RGBA{215, 215, 215, 255},
	ColorSliderCursorHover:     color.RGBA{235, 235, 235, 255},
	ColorSliderCursorActive:    color.RGBA{225, 225, 225, 255},
	ColorProperty:              color.RGBA{225, 225, 225, 255},
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
