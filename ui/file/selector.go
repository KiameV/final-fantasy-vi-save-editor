package file

import (
	"ffvi_editor/global"
	"ffvi_editor/io"
	"ffvi_editor/io/pr"
	"github.com/aarzilli/nucular"
	"io/fs"
	"io/ioutil"
	"path"
)

type FileSelector struct{}

var prIO *pr.PR

func NewFileSelector() *FileSelector {
	// Clear slots
	fs := &FileSelector{}
	fs.updateSlots()
	return fs
}

func (fs *FileSelector) DrawLoad(w *nucular.Window) (loaded bool, err error) {
	if global.Dir == "" || len(global.DirFiles) == 0 {
		if global.Dir, global.DirFiles, err = io.OpenDirAndFileDialog(w); err != nil {
			return
		}
		fs.updateSlots()
	}

	w.Row(30).Static(400, 100, 100, 100)
	w.Label(global.Dir, "LC")
	if w.ButtonText("Change") {
		if d, f, e1 := io.OpenDirAndFileDialog(w); e1 == nil {
			global.Dir = d
			global.DirFiles = f
			fs.updateSlots()
		}
	}
	if w.ButtonText("Refresh") {
		fs.updateSlots()
	}
	if w.ButtonText("Cancel") {
		global.RollbackShowing()
		global.SetShowing(global.GetCurrentShowing())
	}

	if len(prSlots) == 0 {
		w.Row(30).Static(300)
		w.Label("No saves found in this directory.", "LC")
	} else {
		for _, s := range prSlots {
			if s.File != nil {
				w.Row(30).Static(300)
				if w.ButtonText("Load " + s.Name) {
					p := pr.NewPR()
					if err = p.Load(path.Join(global.Dir, s.File.Name())); err == nil {
						prIO = p
						global.FileName = s.File.Name()
						loaded = true
					}
				}
			}
		}
	}
	return
}

func (fs *FileSelector) DrawSave(w *nucular.Window) (saved bool, err error) {
	w.Row(30).Static(400, 100, 100, 100)
	w.Label(global.Dir, "LC")
	if w.ButtonText("Change") {
		if d, f, e1 := io.OpenDirAndFileDialog(w); e1 == nil {
			global.Dir = d
			global.DirFiles = f
		}
	}
	if w.ButtonText("Refresh") {
		fs.updateSlots()
	}
	if w.ButtonText("Cancel") {
		global.RollbackShowing()
		global.SetShowing(global.GetCurrentShowing())
	}

	for i, s := range prSlots {
		if i == 0 {
			// Don't save over QS
			continue
		}
		w.Row(30).Static(300)
		label := "Save " + s.Name
		if s.File != nil {
			label += " (replace)"
		}
		if w.ButtonText(label) {
			slot := i
			if err = prIO.Save(slot, s.UUID); err == nil {
				saved = true
				fs.updateSlots()
			}
		}
	}
	return
}

func (fs *FileSelector) updateSlots() {
	for _, s := range prSlots {
		s.File = nil
	}
	global.DirFiles, _ = ioutil.ReadDir(global.Dir)
	for _, f := range global.DirFiles {
		if s, ok := slotLookup[f.Name()]; ok {
			s.File = f
		}
	}
}

type prSlot struct {
	UUID string
	Name string
	File fs.FileInfo
}

var (
	prSlots = []*prSlot{
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
			UUID: "vgU2wnuaPje2Or53Iqs8Mp=Al6sdM+GM04Iymv229Ow=",
			Name: "slot 2",
			File: nil,
		},
		{
			UUID: "uhHNR4g5QL5twqCc+IhexaltjtBjJnzzcxh5RBSy4G4=",
			Name: "slot 3",
			File: nil,
		},
		{
			UUID: "fmsBRQ+D6YzdjCbBbl7BQuagHyg=7iX3I=EnhccyGDM=",
			Name: "slot 4",
			File: nil,
		},
		{
			UUID: "NXa+MQ+hiHKlPAHJ6GiVWi2Wk5JR2xQQaQxzhyCbK2E=",
			Name: "slot 5",
			File: nil,
		},
		{
			UUID: "UWtRedIOaeA6ig=8r6DIvxg33X92oMM9P8JBwiag4d0=",
			Name: "slot 6",
			File: nil,
		},
		{
			UUID: "e1gfNt2iCE2I3yucQ8zfXn0ou+P2=lREb2q7Lqm04Gc=",
			Name: "slot 7",
			File: nil,
		},
		{
			UUID: "6Pf6Ky7e4QBPuKH9EFJ1Iu+BUEz0zNrXdaS8866Gcq0=",
			Name: "slot 8",
			File: nil,
		},
		{
			UUID: "9dHjN5+9JJWfJ9xoprXo=ehwoEwJwKRYL1Hlc92UNQk=",
			Name: "slot 9",
			File: nil,
		},
		{
			UUID: "oY6N7KlcC4jscZnfa4ea6Nr=TUSR+I=29kwPNZe2NAo=",
			Name: "slot 10",
			File: nil,
		},
		{
			UUID: "NKQ3ux2pea=DqE=vXPKb8+oix5Lt467opYaG0p0brgU=",
			Name: "slot 11",
			File: nil,
		},
		{
			UUID: "HyhjsKWa=tCVf3TWB3qRy7NyrJbc8orciJCntDpqT=I=",
			Name: "slot 12",
			File: nil,
		},
		{
			UUID: "hl9YCUf633k79xePC9PiKAEOq1ajUcSZkLofQuNw2OM=",
			Name: "slot 13",
			File: nil,
		},
		{
			UUID: "C=ozNkSxgKEoLCgOPLJakAUUhnL820LbGlpMz0irQFI=",
			Name: "slot 14",
			File: nil,
		},
		{
			UUID: "z2837SldCS+oIV8y4w5LrnJK9URKYy1QrnoA9bvCg5o=",
			Name: "slot 15",
			File: nil,
		},
		{
			UUID: "CnvUyfaDeqDg3XbVpVWJOj=sPKcGMCV3dR=xM8Ze5jE=",
			Name: "slot 16",
			File: nil,
		},
		{
			UUID: "eQ9Km3NT1WoE4h0hFD90ggFIZayYxfHkIVntc7akYVo=",
			Name: "slot 17",
			File: nil,
		},
		{
			UUID: "Lnbq+GaFOc4ybPZaCf=llI0arXo06rJL32Eu+mCwsLg=",
			Name: "slot 18",
			File: nil,
		},
		{
			UUID: "9GkO1xc52WAzswcEtJxs963MkuCohOHgYj0Fhio=fPE=",
			Name: "slot 19",
			File: nil,
		},
		{
			UUID: "mkYfUr4Mtg0zUmF=6lw+bxRLnbnBYp9ayg1KgploDpQ=",
			Name: "slot 20",
			File: nil,
		},
	}
	slotLookup = make(map[string]*prSlot)
)

func init() {
	for _, s := range prSlots {
		slotLookup[s.UUID] = s
	}
}
