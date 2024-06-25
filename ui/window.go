package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"pixel-remastered-save-editor/browser"
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/finder"
	"pixel-remastered-save-editor/save/config"
	"pixel-remastered-save-editor/save/file"
	"pixel-remastered-save-editor/ui/bundled"
	"pixel-remastered-save-editor/ui/forms/io"
	"pixel-remastered-save-editor/ui/forms/selections"
)

const (
	appName = "Final Fantasy Pixel Remastered Save Editor"
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
		data   *core.Save
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
			window: a.NewWindow(fmt.Sprintf("%s - v%s", appName, browser.Version)),
			canvas: container.NewStack(),
		}
	)
	g.window.SetIcon(fyne.NewStaticResource("icon", bundled.ResourceIconPng.StaticContent))
	g.window.SetContent(g.canvas)
	g.open = fyne.NewMenuItem("Open", func() {
		g.Load()
	})
	g.open.ChildMenu = fyne.NewMenu("Open",
		fyne.NewMenuItem("I", func() {
			g.gameSelected(global.One)
		}),
		fyne.NewMenuItem("II", func() {
			g.gameSelected(global.Two)
		}),
		fyne.NewMenuItem("III", func() {
			g.gameSelected(global.Three)
		}),
		fyne.NewMenuItem("IV", func() {
			g.gameSelected(global.Four)
		}),
		fyne.NewMenuItem("V", func() {
			g.gameSelected(global.Five)
		}),
		fyne.NewMenuItem("VI", func() {
			g.gameSelected(global.Six)
		}))
	g.save = fyne.NewMenuItem("Save", func() {
		g.Save()
	})
	g.save.Disabled = true
	g.window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			g.open,
			fyne.NewMenuItemSeparator(),
			g.save,
		)))
	g.window.Resize(fyne.NewSize(float32(global.WindowWidth), float32(global.WindowHeight)))
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
	g.canvas.Add(io.NewGameSelect(g.gameSelected))
}

func (g *gui) gameSelected(game global.Game) {
	g.prev = g.canvas.Objects[0]
	g.canvas.RemoveAll()
	g.canvas.Add(io.NewFileIO(io.Load, game, g.window, config.Dir(game), func(game global.Game, name, dir, f string, slot int) {
		defer func() { g.open.Disabled = false }()
		// Load file
		config.SetSaveDir(game, dir)
		if data, err := file.LoadSave(game, filepath.Join(dir, f)); err != nil {
			if g.prev != nil {
				g.canvas.RemoveAll()
				g.canvas.Add(g.prev)
			}
			dialog.NewError(err, g.window).Show()
		} else {
			if g.data, err = core.NewSave(data); err != nil {
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
				finder.Load(game, g.data.Characters.All())
				g.canvas.Add(selections.NewEditor(game, g.data))
				g.window.SetTitle(fmt.Sprintf("%s - FF%d - slot %d", appName, game, slot))
			}
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
		io.NewFileIO(io.Save, g.data.Data.Game, g.window, config.Dir(g.data.Data.Game), func(game global.Game, name, dir, f string, slot int) {
			defer func() {
				g.open.Disabled = false
				g.save.Disabled = false
			}()
			// Save file
			config.SetSaveDir(game, dir)
			d, err := g.data.ToSave(slot)
			if err == nil {
				err = file.SaveSave(game, d, slot, filepath.Join(dir, f))
			}
			if err != nil {
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
				g.window.SetTitle(fmt.Sprintf("%s - FF%d - slot %d", appName, game, slot))
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
