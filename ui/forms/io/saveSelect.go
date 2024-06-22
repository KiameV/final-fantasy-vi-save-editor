package io

import (
	"io/fs"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/sqweek/dialog"
	"pixel-remastered-save-editor/global"
)

const (
	quickSave = "Quick Save"
)

type (
	OnSelect   func(game global.Game, name, dir, file string, slot int)
	saveSelect struct {
		widget.BaseWidget
		game       global.Game
		kind       Kind
		window     fyne.Window
		dir        binding.String
		buttons    *fyne.Container
		onSelected OnSelect
		onCancel   func()
	}
	Kind bool
)

const (
	Load Kind = false
	Save Kind = true
)

func NewFileIO(kind Kind, game global.Game, window fyne.Window, dir string, onSelected OnSelect, onCancel func()) *saveSelect {
	w := &saveSelect{
		kind:       kind,
		game:       game,
		window:     window,
		dir:        binding.BindString(&dir),
		onSelected: onSelected,
		onCancel:   onCancel,
		buttons:    container.NewVBox(),
	}
	w.ExtendBaseWidget(w)
	w.dir.AddListener(w)
	return w
}

func (w *saveSelect) DataChanged() {
	var (
		d        []fs.DirEntry
		dir, err = w.dir.Get()
		m        = make(map[string]fs.FileInfo)
		fi       fs.FileInfo
		found    bool
	)
	if err != nil {
		return
	}
	if d, err = os.ReadDir(dir); err == nil {
		for _, f := range d {
			if !f.IsDir() {
				if fi, err = f.Info(); err == nil {
					m[f.Name()] = fi
				}
			}
		}
	}
	if len(m) > 0 {
		var changed bool
		for _, save := range saves {
			if _, found = m[save.UUID]; found || w.kind == Save {
				if !changed {
					changed = true
					w.buttons.RemoveAll()
				}
				name := save.Name
				if found && w.kind == Save {
					name += " (replace)"
				}
				func(name string, uuid string, slot int) {
					w.buttons.Add(widget.NewButton(name, func() {
						w.onSelected(w.game, name, dir, uuid, slot)
					}))
				}(name, save.UUID, save.Slot)
			}
		}
	}
}

func (w *saveSelect) CreateRenderer() fyne.WidgetRenderer {
	top := container.NewBorder(nil, nil,
		widget.NewLabel("Directory:"),
		widget.NewButton("Change", func() {
			if dir, err := dialog.Directory().Title("Load images").Browse(); err == nil && dir != "" {
				_ = w.dir.Set(dir)
			}
		}),
		widget.NewEntryWithData(w.dir))
	bottom := widget.NewButton("Cancel", func() {
		w.onCancel()
	})
	return widget.NewSimpleRenderer(container.NewBorder(top, bottom, nil, nil, container.NewVScroll(w.buttons)))
}

var (
	saves = []struct {
		UUID string
		Name string
		Slot int
	}{
		{
			UUID: "7nCxyzTwG31W3Zlg70mo751W8ETH1n+Km0dWOzRU84Y=",
			Name: quickSave,
			Slot: 21,
		},
		{
			UUID: "ookrbATYovG3tEOXIH4HqWnsv8TrUlRWzM8AlCmW2mk=",
			Name: "slot 1",
			Slot: 1,
		},
		{
			UUID: "vgU2wnuaPje2Or53Iqs8Mp=Al6sdM+GM04Iymv229Ow=",
			Name: "slot 2",
			Slot: 2,
		},
		{
			UUID: "uhHNR4g5QL5twqCc+IhexaltjtBjJnzzcxh5RBSy4G4=",
			Name: "slot 3",
			Slot: 3,
		},
		{
			UUID: "fmsBRQ+D6YzdjCbBbl7BQuagHyg=7iX3I=EnhccyGDM=",
			Name: "slot 4",
			Slot: 4,
		},
		{
			UUID: "NXa+MQ+hiHKlPAHJ6GiVWi2Wk5JR2xQQaQxzhyCbK2E=",
			Name: "slot 5",
			Slot: 5,
		},
		{
			UUID: "UWtRedIOaeA6ig=8r6DIvxg33X92oMM9P8JBwiag4d0=",
			Name: "slot 6",
			Slot: 6,
		},
		{
			UUID: "e1gfNt2iCE2I3yucQ8zfXn0ou+P2=lREb2q7Lqm04Gc=",
			Name: "slot 7",
			Slot: 7,
		},
		{
			UUID: "6Pf6Ky7e4QBPuKH9EFJ1Iu+BUEz0zNrXdaS8866Gcq0=",
			Name: "slot 8",
			Slot: 8,
		},
		{
			UUID: "9dHjN5+9JJWfJ9xoprXo=ehwoEwJwKRYL1Hlc92UNQk=",
			Name: "slot 9",
			Slot: 9,
		},
		{
			UUID: "oY6N7KlcC4jscZnfa4ea6Nr=TUSR+I=29kwPNZe2NAo=",
			Name: "slot 10",
			Slot: 10,
		},
		{
			UUID: "NKQ3ux2pea=DqE=vXPKb8+oix5Lt467opYaG0p0brgU=",
			Name: "slot 11",
			Slot: 11,
		},
		{
			UUID: "HyhjsKWa=tCVf3TWB3qRy7NyrJbc8orciJCntDpqT=I=",
			Name: "slot 12",
			Slot: 12,
		},
		{
			UUID: "hl9YCUf633k79xePC9PiKAEOq1ajUcSZkLofQuNw2OM=",
			Name: "slot 13",
			Slot: 13,
		},
		{
			UUID: "C=ozNkSxgKEoLCgOPLJakAUUhnL820LbGlpMz0irQFI=",
			Name: "slot 14",
			Slot: 14,
		},
		{
			UUID: "z2837SldCS+oIV8y4w5LrnJK9URKYy1QrnoA9bvCg5o=",
			Name: "slot 15",
			Slot: 15,
		},
		{
			UUID: "CnvUyfaDeqDg3XbVpVWJOj=sPKcGMCV3dR=xM8Ze5jE=",
			Name: "slot 16",
			Slot: 16,
		},
		{
			UUID: "eQ9Km3NT1WoE4h0hFD90ggFIZayYxfHkIVntc7akYVo=",
			Name: "slot 17",
			Slot: 17,
		},
		{
			UUID: "Lnbq+GaFOc4ybPZaCf=llI0arXo06rJL32Eu+mCwsLg=",
			Name: "slot 18",
			Slot: 18,
		},
		{
			UUID: "9GkO1xc52WAzswcEtJxs963MkuCohOHgYj0Fhio=fPE=",
			Name: "slot 19",
			Slot: 19,
		},
		{
			UUID: "mkYfUr4Mtg0zUmF=6lw+bxRLnbnBYp9ayg1KgploDpQ=",
			Name: "slot 20",
			Slot: 20,
		},
	}
)
