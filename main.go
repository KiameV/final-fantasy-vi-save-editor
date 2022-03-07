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
	"github.com/aarzilli/nucular/rect"
	"github.com/aarzilli/nucular/style"
	"image"
	"image/color"
	"time"
)

type FileType byte

const (
	Unknown FileType = iota
	SNES
	SteamRemastered
)

var (
	cUI          = character.NewUI()
	iUI          = inventory.NewUI()
	sUI          = skills.NewUI()
	eUI          = espers.NewUI()
	mUI          = misc.NewUI()
	fileType     = Unknown
	saveFileType save.SaveFileType
	fn           string
	fileName     string
	status       string
	err          error
)

func main() {
	io.ParsePixelRemasteredSave(`C:\Users\travis\Desktop\bu\start\uhHNR4g5QL5twqCc+IhexaltjtBjJnzzcxh5RBSy4G4=`)
	return

	wnd := nucular.NewMasterWindowSize(0, "Final Fantasy VI Editor", image.Point{X: 725, Y: 500}, updateWindow)
	wnd.SetStyle(style.FromTable(customTheme, 1.2))
	wnd.Main()
}

func updateWindow(w *nucular.Window) {
	//mw := w.Master()
	w.MenubarBegin()
	w.Row(12).Static(100, 100, 75, 50, 100)
	/*if w := w.Menu(label.TA("Load Remaster Save", "CC"), 100, nil); w != nil {
		if fn, err = io.OpenFile(w, save.SteamRemastered); err != nil {
			popupErr(w, err)
			w.Close()
		} else if fn != "" {
			fileName = fn
			saveFileType = save.SteamRemastered
			fileType = SteamRemastered
			refresh()
		}
	} else */
	if w := w.Menu(label.TA("Load SNES", "CC"), 100, nil); w != nil {
		w.Row(12).Dynamic(1)
		if w.MenuItem(label.TA("SRM Slot 1", "LC")) {
			if fn, err = io.OpenFile(w, save.SRMSlot1); err != nil {
				popupErr(w, err)
				w.Close()
			} else if fn != "" {
				fileName = fn
				saveFileType = save.SRMSlot1
				fileType = SNES
				refresh()
			}
		} else if w.MenuItem(label.TA("SRM Slot 2", "LC")) {
			if fn, err = io.OpenFile(w, save.SRMSlot2); err != nil {
				popupErr(w, err)
				w.Close()
			} else if fn != "" {
				fileName = fn
				saveFileType = save.SRMSlot2
				fileType = SNES
				refresh()
			}
		} else if w.MenuItem(label.TA("SRM Slot 3", "LC")) {
			if fn, err = io.OpenFile(w, save.SRMSlot3); err != nil {
				popupErr(w, err)
				w.Close()
			} else if fn != "" {
				fileName = fn
				saveFileType = save.SRMSlot3
				fileType = SNES
				refresh()
			}
		} else if w.MenuItem(label.TA("ZNES State", "LC")) {
			if fn, err = io.OpenFile(w, save.ZnesSaveState); err != nil {
				popupErr(w, err)
				w.Close()
			} else if fn != "" {
				fileName = fn
				saveFileType = save.ZnesSaveState
				fileType = SNES
				refresh()
			}
		}
	}
	if fileType != Unknown && fileName != "" {
		if w := w.Menu(label.TA("Save", "CC"), 100, nil); w != nil {
			if err = io.SaveFileNoDialog(fileName, saveFileType); err != nil {
				popupErr(w, err)
				w.Close()
			}
			status = "Saved"
			time.AfterFunc(2*time.Second, func() {
				status = ""
			})
			w.Close()
		}
	}
	/* else if fileType == SNES {
		if w := w.Menu(label.TA("Save As...", "LC"), 100, nil); w != nil {
			w.Row(12).Dynamic(1)
			if fileName != "" {
				if w.MenuItem(label.TA("Save", "LC")) {
					if err = io.SaveFileNoDialog(fileName, saveFileType); err != nil {
				popupErr(w, err)
				w.Close()
					}
				}
			}
			if w.MenuItem(label.TA("SRM Slot 1", "LC")) {
				if fileName, err = io.SaveFile(w, save.SRMSlot1); err != nil {
					popupErr(w, err)
					w.Close()
				} else if fileName != "" {
					fileType = SNES
					saveFileType = save.SRMSlot1
				}
			} else if w.MenuItem(label.TA("SRM Slot 2", "LC")) {
				if fileName, err = io.SaveFile(w, save.SRMSlot2); err != nil {
					popupErr(w, err)
					w.Close()
				} else if fileName != "" {
					fileType = SNES
					saveFileType = save.SRMSlot2
				}
			} else if w.MenuItem(label.TA("SRM Slot 3", "LC")) {
				if fileName, err = io.OpenFile(w, save.SRMSlot3); err != nil {
					popupErr(w, err)
					w.Close()
				} else if fileName != "" {
					fileType = SNES
					saveFileType = save.SRMSlot3
				}
			} else if w.MenuItem(label.TA("ZNES State", "LC")) {
				if fileName, err = io.OpenFile(w, save.ZnesSaveState); err != nil {
					popupErr(w, err)
					w.Close()
				} else if fileName != "" {
					fileType = SNES
					saveFileType = save.ZnesSaveState
				}
			}
		}
	}*/
	if status != "" {
		w.Spacing(1)
		w.Label("Status: "+status, "RC")
	} else {
		w.Spacing(2)
	}
	w.MenubarEnd()

	w.Row(5).Static(1)
	w.Row(12).Static(200, 200)
	if fileType != Unknown {
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

func popupErr(w *nucular.Window, err error) {
	w.Master().PopupOpen("Error", nucular.WindowMovable|nucular.WindowTitle|nucular.WindowDynamic|nucular.WindowNoScrollbar, rect.Rect{X: 20, Y: 100, W: 230, H: 150}, true,
		func(w *nucular.Window) {
			w.Row(25).Dynamic(1)
			w.Label(err.Error(), "LC")
			w.Row(25).Dynamic(1)
			if w.Button(label.T("OK"), false) {
				w.Close()
			}
		})
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

/*
Slots:
b'ookrbATYovG3tEOXIH4HqWnsv8TrUlRWzM8AlCmW2mk=' :  slot1.sav
b'vgU2wnuaPje2Or53Iqs8Mp/Al6sdM+GM04Iymv229Ow=' :  slot2.sav
b'uhHNR4g5QL5twqCc+IhexaltjtBjJnzzcxh5RBSy4G4=' :  slot3.sav
b'fmsBRQ+D6YzdjCbBbl7BQuagHyg/7iX3I/EnhccyGDM=' :  slot4.sav
b'NXa+MQ+hiHKlPAHJ6GiVWi2Wk5JR2xQQaQxzhyCbK2E=' :  slot5.sav
b'UWtRedIOaeA6ig/8r6DIvxg33X92oMM9P8JBwiag4d0=' :  slot6.sav
b'e1gfNt2iCE2I3yucQ8zfXn0ou+P2/lREb2q7Lqm04Gc=' :  slot7.sav
b'6Pf6Ky7e4QBPuKH9EFJ1Iu+BUEz0zNrXdaS8866Gcq0=' :  slot8.sav
b'9dHjN5+9JJWfJ9xoprXo/ehwoEwJwKRYL1Hlc92UNQk=' :  slot9.sav
b'oY6N7KlcC4jscZnfa4ea6Nr/TUSR+I/29kwPNZe2NAo=' :  slot10.sav
b'NKQ3ux2pea/DqE/vXPKb8+oix5Lt467opYaG0p0brgU=' :  slot11.sav
b'HyhjsKWa/tCVf3TWB3qRy7NyrJbc8orciJCntDpqT/I=' :  slot12.sav
b'hl9YCUf633k79xePC9PiKAEOq1ajUcSZkLofQuNw2OM=' :  slot13.sav
b'C/ozNkSxgKEoLCgOPLJakAUUhnL820LbGlpMz0irQFI=' :  slot14.sav
b'z2837SldCS+oIV8y4w5LrnJK9URKYy1QrnoA9bvCg5o=' :  slot15.sav
b'CnvUyfaDeqDg3XbVpVWJOj/sPKcGMCV3dR/xM8Ze5jE=' :  slot16.sav
b'eQ9Km3NT1WoE4h0hFD90ggFIZayYxfHkIVntc7akYVo=' :  slot17.sav
b'Lnbq+GaFOc4ybPZaCf/llI0arXo06rJL32Eu+mCwsLg=' :  slot18.sav
b'9GkO1xc52WAzswcEtJxs963MkuCohOHgYj0Fhio/fPE=' :  slot19.sav
b'mkYfUr4Mtg0zUmF/6lw+bxRLnbnBYp9ayg1KgploDpQ=' :  slot20.sav
b'7nCxyzTwG31W3Zlg70mo751W8ETH1n+Km0dWOzRU84Y=' :  slot21.sav
*/
