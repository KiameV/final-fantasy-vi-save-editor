package io

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/ui/bundled"
)

type (
	GameSelect struct {
		widget.BaseWidget
		selections []fyne.CanvasObject
	}
)

func NewGameSelect(onSelect func(game global.Game)) *GameSelect {
	w := &GameSelect{
		selections: make([]fyne.CanvasObject, 6),
	}
	w.ExtendBaseWidget(w)
	for i, j := range []*fyne.StaticResource{bundled.Resource1Png, bundled.Resource2Png, bundled.Resource3Png, bundled.Resource4Png, bundled.Resource5Png, bundled.Resource6Png} {
		func(i int, r *fyne.StaticResource) {
			img := canvas.NewImageFromResource(r)
			img.FillMode = canvas.ImageFillContain
			w.selections[i] = container.NewPadded(container.NewStack(widget.NewButton("", func() {
				onSelect(global.Game(i + 1))
			}), img))
		}(i, j)
	}
	return w
}

func (w *GameSelect) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		container.NewGridWithRows(6, w.selections...))
}
