package inputs

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models"
)

type (
	Search struct {
		Text    *widget.TextGrid
		Input   *widget.Entry
		Results *widget.TextGrid
	}
)

func NewSearch(nvss ...[]models.NameValue) *Search {
	s := &Search{
		Text:    widget.NewTextGrid(),
		Input:   widget.NewEntry(),
		Results: widget.NewTextGrid(),
	}
	s.Input.OnChanged = func(text string) {
		if len(text) >= 2 {
			text = strings.ToLower(text)
			var sb strings.Builder
			for _, nvs := range nvss {
				for _, nv := range nvs {
					if strings.Contains(strings.ToLower(nv.Name), text) {
						sb.WriteString(fmt.Sprintf("%d: %s\n", nv.Value, nv.Name))
					}
				}
			}
			s.Results.SetText(sb.String())
		} else {
			s.Results.SetText("")
		}
	}
	var sb strings.Builder
	for _, nvs := range nvss {
		for _, nv := range nvs {
			sb.WriteString(fmt.Sprintf("%d: %s\n", nv.Value, nv.Name))
		}
	}
	s.Text.SetText(sb.String())
	return s
}

func (s *Search) Filter() fyne.CanvasObject {
	return container.NewBorder(s.Input, nil, nil, nil, container.NewVScroll(s.Results))
}

func (s *Search) Fields() fyne.CanvasObject {
	return container.NewVScroll(s.Text)
}
