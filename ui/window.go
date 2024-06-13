package ui

import (
	"os"
	"path/filepath"
	"strings"

	"ffvi_editor/io/config"
	"ffvi_editor/io/pr"
	"ffvi_editor/ui/forms"
	"ffvi_editor/ui/forms/selections"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
)

type (
	Gui interface {
		App() fyne.App
		Window() fyne.Window
		Canvas() *fyne.Container
		Load()
		Save()
		Run()
	}
	gui struct {
		app    fyne.App
		window fyne.Window
		canvas *fyne.Container
		open   *fyne.MenuItem
		save   *fyne.MenuItem
		prev   fyne.CanvasObject
		pr     *pr.PR
	}
	MenuItem interface {
		Item() *fyne.MenuItem
		Disable()
		Enable()
		OnSelected(func())
	}
	menuItem struct {
		*fyne.MenuItem
		onSelected func()
	}
)

var (
	_ MenuItem = menuItem{}
)

func New() Gui {
	if wd, err := os.Getwd(); err == nil {
		var dir []os.DirEntry
		if dir, err = os.ReadDir(wd); err == nil {
			for _, f := range dir {
				if !f.IsDir() && strings.HasSuffix(f.Name(), ".ttf") {
					_ = os.Setenv("FYNE_FONT", f.Name())
					break
				}
			}
		}
	}
	var (
		a = app.New()
		g = &gui{
			app:    a,
			window: a.NewWindow("Final Fantasy VI Save Editor"),
			canvas: container.NewStack(),
		}
	)
	g.window.SetContent(g.canvas)
	g.open = fyne.NewMenuItem("Open", func() {
		g.Load()
	})
	g.save = fyne.NewMenuItem("Save", func() {
		g.Save()
	})
	g.save.Disabled = true
	x, y := config.WindowSize()
	g.window.Resize(fyne.NewSize(x, y))
	g.window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			g.open,
			fyne.NewMenuItemSeparator(),
			g.save,
		)))
	return g
}

func (g *gui) App() fyne.App {
	return g.app
}

func (g *gui) Window() fyne.Window {
	return g.window
}

func (g *gui) Canvas() *fyne.Container {
	return g.canvas
}

func (g *gui) Load() {
	if len(g.canvas.Objects) > 0 {
		g.prev = g.canvas.Objects[0]
	}
	g.open.Disabled = true
	g.canvas.RemoveAll()
	g.canvas.Add(
		forms.NewFileIO(forms.Load, g.window, config.SaveDir(), func(name, dir, file string, _ int) {
			defer func() { g.open.Disabled = false }()
			// Load file
			config.SetSaveDir(dir)
			p := pr.New()
			if err := p.Load(filepath.Join(dir, file)); err != nil {
				if g.prev != nil {
					g.canvas.RemoveAll()
					g.canvas.Add(g.prev)
				}
				dialog.NewError(err, g.window).Show()
			} else {
				// Success
				g.canvas.RemoveAll()
				g.prev = nil
				g.save.Disabled = false
				g.pr = p
				g.canvas.Add(selections.NewEditor())
			}
		}, func() {
			defer func() { g.open.Disabled = false }()
			// Cancel
			if g.prev != nil {
				g.canvas.RemoveAll()
				g.canvas.Add(g.prev)
			}
		}))
}

func (g *gui) Save() {
	if len(g.canvas.Objects) > 0 {
		g.prev = g.canvas.Objects[0]
	}
	g.open.Disabled = true
	g.save.Disabled = true
	g.canvas.RemoveAll()
	g.canvas.Add(
		forms.NewFileIO(forms.Save, g.window, config.SaveDir(), func(name, dir, file string, slot int) {
			defer func() {
				g.open.Disabled = false
				g.save.Disabled = false
			}()
			// Save file
			config.SetSaveDir(dir)
			if err := g.pr.Save(slot, filepath.Join(dir, file)); err != nil {
				if g.prev != nil {
					g.canvas.RemoveAll()
					g.canvas.Add(g.prev)
				}
				dialog.NewError(err, g.window).Show()
			} else {
				// Success
				if g.prev != nil {
					g.canvas.RemoveAll()
					g.canvas.Add(g.prev)
				}
			}
		}, func() {
			defer func() {
				g.open.Disabled = false
				g.save.Disabled = false
			}()
			// Cancel
			if g.prev != nil {
				g.canvas.RemoveAll()
				g.canvas.Add(g.prev)
			}
		}))
}

func (g *gui) Run() {
	g.window.ShowAndRun()
}

func (m menuItem) Item() *fyne.MenuItem {
	return m.MenuItem
}

func (m menuItem) Disable() {
	m.MenuItem.Disabled = true
}

func (m menuItem) Enable() {
	m.MenuItem.Disabled = false
}

func (m menuItem) OnSelected(f func()) {
	m.onSelected = f
}
