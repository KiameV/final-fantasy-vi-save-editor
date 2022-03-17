package main

import (
	"ffvi_editor/browser"
	"ffvi_editor/global"
	"ffvi_editor/io"
	"ffvi_editor/ui"
	"ffvi_editor/ui/file"
	mm "ffvi_editor/ui/mainMenu"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/label"
	"github.com/aarzilli/nucular/rect"
	"github.com/aarzilli/nucular/style"
	"image"
	"image/color"
	"io/ioutil"
	"time"
)

const version = "1.5.0"

var (
	mainMenu     ui.UI
	status       string
	err          error
	fileSelector *file.FileSelector
	statusTimer  *time.Timer
)

func main() {
	if global.Dir != "" {
		global.DirFiles, _ = ioutil.ReadDir(global.Dir)
	}
	mainMenu = mm.NewUI()
	wnd := nucular.NewMasterWindowSize(0, "Final Fantasy VI Editor - "+version, image.Point{X: 725, Y: 500}, updateWindow)
	wnd.SetStyle(style.FromTable(customTheme, 1.2))
	wnd.Main()
}

func updateWindow(w *nucular.Window) {
	//mw := w.Master()
	var fn string
	var ft global.SaveFileType
	w.MenubarBegin()
	w.Row(12).Static(100, 100, 75, 125, 100, 100)
	if w := w.Menu(label.TA("Load Remaster Save", "CC"), 100, nil); w != nil {
		global.SetShowing(global.LoadPR)
		fileSelector = file.NewFileSelector()
		w.Close()
	}
	if w := w.Menu(label.TA("Load SNES", "CC"), 100, nil); w != nil {
		w.Row(12).Dynamic(1)
		if w.MenuItem(label.TA("SRM Slot 1", "LC")) {
			global.SetShowing(global.LoadSnes)
			ft = global.SRMSlot1
			w.Close()
		} else if w.MenuItem(label.TA("SRM Slot 2", "LC")) {
			global.SetShowing(global.LoadSnes)
			ft = global.SRMSlot2
			w.Close()
		} else if w.MenuItem(label.TA("SRM Slot 3", "LC")) {
			global.SetShowing(global.LoadSnes)
			ft = global.SRMSlot3
			w.Close()
		} else if w.MenuItem(label.TA("ZNES State", "LC")) {
			global.SetShowing(global.LoadSnes)
			ft = global.ZnesSaveState
			w.Close()
		}
	}
	if global.IsShowing(global.ShowPR) || global.IsShowing(global.ShowSnes) {
		ui.DrawError = nil
		if w := w.Menu(label.TA("Save", "CC"), 100, nil); w != nil {
			w.Close()
			if global.IsShowing(global.ShowPR) {
				fileSelector = file.NewFileSelector()
				global.SetShowing(global.SavePR)
			} else { // Snes
				if err = io.SaveFileSnesNoDialog(global.FileName); err != nil {
					if err.Error() == "Cancelled" {
						err = nil
					} else {
						popupErr(w, err)
					}
					global.RollbackShowing()
				}
				status = "Saved"
			}
		}
		if ui.DrawError != nil {
			popupErr(w, ui.DrawError)
		}
	} else {
		w.Spacing(1)
	}

	if w := w.Menu(label.TA("Check For Update", "CC"), 100, nil); w != nil {
		var hasNewer bool
		var latest string
		if hasNewer, latest, err = browser.CheckForUpdate(version); err != nil {
			popupErr(w, err)
		}
		if hasNewer {
			browser.Update(latest)
		} else {
			status = "version is current"
		}
		w.Close()
	}

	if status != "" {
		w.Spacing(1)
		w.Label("Status: "+status, "RC")
		if statusTimer != nil {
			statusTimer.Stop()
		}
		statusTimer = time.AfterFunc(2*time.Second, func() { status = "" })

	} else {
		w.Spacing(1)
	}
	w.MenubarEnd()

	switch global.GetCurrentShowing() {
	case global.LoadPR:
		var loaded bool
		if loaded, err = fileSelector.DrawLoad(w); err != nil {
			if err.Error() == "Cancelled" {
				err = nil
			} else {
				popupErr(w, err)
			}
			global.RollbackShowing()
		} else if loaded {
			global.SetShowing(global.ShowPR)
			fileSelector = nil
			mainMenu.Refresh()
		}
	case global.LoadSnes:
		if fn, err = io.OpenFileDialog(w, ft); err != nil {
			if err.Error() == "Cancelled" {
				err = nil
			} else {
				popupErr(w, err)
			}
			global.RollbackShowing()
		} else if fn != "" {
			global.DirFiles = nil
			global.FileName = fn
			global.FileType = ft
			mainMenu.Refresh()
			global.SetShowing(global.ShowSnes)
		}
	case global.SavePR, global.SaveSnes:
		var saved bool
		if saved, err = fileSelector.DrawSave(w); err != nil {
			popupErr(w, err)
			fileSelector = nil
		} else if saved {
			global.SetShowing(global.ShowPR)
			fileSelector = nil
		}
	case global.ShowPR, global.ShowSnes:
		mainMenu.Draw(w)
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
					if fileName, err = io.OpenFileDialog(w, save.SRMSlot3); err != nil {
						popupErr(w, err)
						w.Close()
					} else if fileName != "" {
						fileType = SNES
						saveFileType = save.SRMSlot3
					}
				} else if w.MenuItem(label.TA("ZNES State", "LC")) {
					if fileName, err = io.OpenFileDialog(w, save.ZnesSaveState); err != nil {
						popupErr(w, err)
						w.Close()
					} else if fileName != "" {
						fileType = SNES
						saveFileType = save.ZnesSaveState
					}
				}
			}
		}*/
	}
}

func popupErr(w *nucular.Window, err error) {
	w.Master().PopupOpen("Error", nucular.WindowMovable|nucular.WindowTitle|nucular.WindowDynamic, rect.Rect{X: 20, Y: 100, W: 600, H: 600}, true,
		func(w *nucular.Window) {
			w.Row(100).Dynamic(1)
			w.LabelWrap(err.Error())
			w.Row(25).Dynamic(1)
			if w.Button(label.T("OK"), false) {
				w.Close()
			}
		})
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
