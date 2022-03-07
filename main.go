package main

import (
	"errors"
	"ffvi_editor/io"
	"ffvi_editor/ui"
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
	"io/fs"
	"io/ioutil"
	"path"
	"time"
)

type FileType byte

const (
	Unknown FileType = iota
	SNES
	PixelRemastered
)

var (
	cUI          = character.NewUI()
	iUI          = inventory.NewUI()
	sUI          = skills.NewUI()
	eUI          = espers.NewUI()
	mUI          = misc.NewUI()
	fileType     = Unknown
	saveFileType io.SaveFileType
	fn           string
	dirFiles     []fs.FileInfo
	fileName     string
	status       string
	err          error
)

func main() {
	if io.Dir != "" {
		dirFiles, _ = ioutil.ReadDir(io.Dir)
	}

	wnd := nucular.NewMasterWindowSize(0, "Final Fantasy VI Editor", image.Point{X: 725, Y: 500}, updateWindow)
	wnd.SetStyle(style.FromTable(customTheme, 1.2))
	wnd.Main()
}

func updateWindow(w *nucular.Window) {
	//mw := w.Master()
	w.MenubarBegin()
	w.Row(12).Static(100, 100, 75, 50, 100)
	if w := w.Menu(label.TA("Load Remaster Save", "CC"), 100, nil); w != nil {
		if io.Dir != "" && len(dirFiles) > 0 {
			loadPRDir(w)
			w.Close()
		} else if io.Dir, dirFiles, err = io.OpenDirAndFileDialog(w); err != nil {
			popupErr(w, err)
			w.Close()
		} else if io.Dir != "" && dirFiles != nil {
			loadPRDir(w)
			w.Close()
		}
	}
	if w := w.Menu(label.TA("Load SNES", "CC"), 100, nil); w != nil {
		w.Row(12).Dynamic(1)
		if w.MenuItem(label.TA("SRM Slot 1", "LC")) {
			if fn, err = io.OpenFileDialog(w, io.SRMSlot1); err != nil {
				popupErr(w, err)
				w.Close()
			} else if fn != "" {
				io.Dir = ""
				dirFiles = nil
				fileName = fn
				setFileType(io.SRMSlot1)
				w.Close()
				refresh()
			}
		} else if w.MenuItem(label.TA("SRM Slot 2", "LC")) {
			if fn, err = io.OpenFileDialog(w, io.SRMSlot2); err != nil {
				popupErr(w, err)
				w.Close()
			} else if fn != "" {
				fileName = fn
				setFileType(io.SRMSlot2)
				refresh()
			}
		} else if w.MenuItem(label.TA("SRM Slot 3", "LC")) {
			if fn, err = io.OpenFileDialog(w, io.SRMSlot3); err != nil {
				popupErr(w, err)
				w.Close()
			} else if fn != "" {
				fileName = fn
				setFileType(io.SRMSlot3)
				refresh()
			}
		} else if w.MenuItem(label.TA("ZNES State", "LC")) {
			if fn, err = io.OpenFileDialog(w, io.ZnesSaveState); err != nil {
				popupErr(w, err)
				w.Close()
			} else if fn != "" {
				fileName = fn
				setFileType(io.ZnesSaveState)
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
	if status != "" {
		w.Spacing(1)
		w.Label("Status: "+status, "RC")
	} else {
		w.Spacing(2)
	}
	w.MenubarEnd()

	if fileType == PixelRemastered && io.Dir != "" && len(dirFiles) > 0 && fileName == "" {
		w.Row(30).Static(400, 100)
		w.Label(io.Dir, "LC")
		if w.ButtonText("Change") {
			if d, f, e1 := io.OpenDirAndFileDialog(w); e1 == nil {
				io.Dir = d
				dirFiles = f
			}
		}
		for _, s := range prSlots {
			if s.File != nil {
				w.Row(30).Static(300)
				if w.ButtonText("Load " + s.Name) {
					if err = io.OpenFile(path.Join(io.Dir, s.File.Name()), io.PixelRemastered); err != nil {
						popupErr(w, err)
					} else {
						fileName = s.File.Name()
						refresh()
					}
				}
			}
		}
	} else {
		w.Row(5).Static(1)
		w.Row(12).Static(200, 200)
		if fileType != Unknown {
			if w.TreePush(nucular.TreeTab, "Characters", true) {
				cUI.Draw(w)
				w.TreePop()
			}
			if w.TreePush(nucular.TreeTab, "Inventory", false) {
				if fileType != PixelRemastered {
					iUI.Draw(w)
				} else {
					w.Row(30).Dynamic(1)
					w.Label("Coming soon for Pixel Remastered", "LC")
				}
				w.TreePop()
			}
			if w.TreePush(nucular.TreeTab, "Skills", false) {
				if fileType != PixelRemastered {
					sUI.Draw(w)
				} else {
					w.Row(30).Dynamic(1)
					w.Label("Coming soon for Pixel Remastered", "LC")
				}
				w.TreePop()
			}
			if w.TreePush(nucular.TreeTab, "Espers", false) {
				if fileType != PixelRemastered {
					eUI.Draw(w)
				} else {
					w.Row(30).Dynamic(1)
					w.Label("Coming soon for Pixel Remastered", "LC")
				}
				w.TreePop()
			}
			if w.TreePush(nucular.TreeTab, "Misc", false) {
				mUI.Draw(w)
				w.TreePop()
			}
		}
	}
}

func loadPRDir(w *nucular.Window) {
	fileName = ""
	setFileType(io.PixelRemastered)
	for _, s := range prSlots {
		s.File = nil
	}
	var found bool
	for _, f := range dirFiles {
		if s, ok := slotLookup[f.Name()]; ok {
			s.File = f
			found = true
		}
	}
	if !found {
		fileType = Unknown
		io.Dir = ""
		dirFiles = nil
		popupErr(w, errors.New("no save files found in that directory"))
	}
}

func setFileType(ft io.SaveFileType) {
	saveFileType = ft
	switch ft {
	case io.PixelRemastered:
		fileType = PixelRemastered
		ui.IsPR = true
	default:
		fileType = SNES
		ui.IsPR = false
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

type PRSlot struct {
	UUID string
	Name string
	File fs.FileInfo
}

var selectedPRSlot string

var (
	prSlots = []*PRSlot{
		{
			UUID: "7nCxyzTwG31W3Zlg70mo751W8ETH1n+Km0dWOzRU84Y=",
			Name: "quick save",
			File: nil,
		},
		{
			UUID: "ookrbATYovG3tEOXIH4HqWnsv8TrUlRWzM8AlCmW2mk=",
			Name: "slot 1",
			File: nil,
		},
		{
			UUID: "vgU2wnuaPje2Or53Iqs8Mp/Al6sdM+GM04Iymv229Ow=",
			Name: "slot 2",
			File: nil,
		},
		{
			UUID: "uhHNR4g5QL5twqCc+IhexaltjtBjJnzzcxh5RBSy4G4=",
			Name: "slot 3",
			File: nil,
		},
		{
			UUID: "fmsBRQ+D6YzdjCbBbl7BQuagHyg/7iX3I/EnhccyGDM=",
			Name: "slot 4",
			File: nil,
		},
		{
			UUID: "NXa+MQ+hiHKlPAHJ6GiVWi2Wk5JR2xQQaQxzhyCbK2E=",
			Name: "slot 5",
			File: nil,
		},
		{
			UUID: "UWtRedIOaeA6ig/8r6DIvxg33X92oMM9P8JBwiag4d0=",
			Name: "slot 6",
			File: nil,
		},
		{
			UUID: "e1gfNt2iCE2I3yucQ8zfXn0ou+P2/lREb2q7Lqm04Gc=",
			Name: "slot 7",
			File: nil,
		},
		{
			UUID: "6Pf6Ky7e4QBPuKH9EFJ1Iu+BUEz0zNrXdaS8866Gcq0=",
			Name: "slot 8",
			File: nil,
		},
		{
			UUID: "9dHjN5+9JJWfJ9xoprXo/ehwoEwJwKRYL1Hlc92UNQk=",
			Name: "slot 9",
			File: nil,
		},
		{
			UUID: "oY6N7KlcC4jscZnfa4ea6Nr/TUSR+I/29kwPNZe2NAo=",
			Name: "slot 10",
			File: nil,
		},
		{
			UUID: "NKQ3ux2pea/DqE/vXPKb8+oix5Lt467opYaG0p0brgU=",
			Name: "slot 11",
			File: nil,
		},
		{
			UUID: "HyhjsKWa/tCVf3TWB3qRy7NyrJbc8orciJCntDpqT/I=",
			Name: "slot 12",
			File: nil,
		},
		{
			UUID: "hl9YCUf633k79xePC9PiKAEOq1ajUcSZkLofQuNw2OM=",
			Name: "slot 13",
			File: nil,
		},
		{
			UUID: "C/ozNkSxgKEoLCgOPLJakAUUhnL820LbGlpMz0irQFI=",
			Name: "slot 14",
			File: nil,
		},
		{
			UUID: "z2837SldCS+oIV8y4w5LrnJK9URKYy1QrnoA9bvCg5o=",
			Name: "slot 15",
			File: nil,
		},
		{
			UUID: "CnvUyfaDeqDg3XbVpVWJOj/sPKcGMCV3dR/xM8Ze5jE=",
			Name: "slot 16",
			File: nil,
		},
		{
			UUID: "eQ9Km3NT1WoE4h0hFD90ggFIZayYxfHkIVntc7akYVo=",
			Name: "slot 17",
			File: nil,
		},
		{
			UUID: "Lnbq+GaFOc4ybPZaCf/llI0arXo06rJL32Eu+mCwsLg=",
			Name: "slot 18",
			File: nil,
		},
		{
			UUID: "9GkO1xc52WAzswcEtJxs963MkuCohOHgYj0Fhio/fPE=",
			Name: "slot 19",
			File: nil,
		},
		{
			UUID: "mkYfUr4Mtg0zUmF/6lw+bxRLnbnBYp9ayg1KgploDpQ=",
			Name: "slot 20",
			File: nil,
		},
	}
	slotLookup = make(map[string]*PRSlot)
)

func init() {
	for _, s := range prSlots {
		slotLookup[s.UUID] = s
	}
}
