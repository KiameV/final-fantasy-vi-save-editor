package inputs

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/models"
	one "pixel-remastered-save-editor/models/core/ff1/consts"
	two "pixel-remastered-save-editor/models/core/ff2/consts"
)

type (
	Searches struct {
		Abilities      *Search
		ImportantItems *Search
		Items          *Search
		Jobs           *Search
		Maps           *Search
	}
	Search struct {
		Text    *widget.TextGrid
		Input   *widget.Entry
		Results *widget.TextGrid
	}
)

var singletonSearch *Searches

func GetSearches() *Searches {
	return singletonSearch
}

func Load(game global.Game) {
	singletonSearch = &Searches{}
	if game == global.One {
		singletonSearch.Abilities = newSearch(one.Abilities)
		singletonSearch.ImportantItems = newSearch(one.ImportantItems)
		singletonSearch.Items = newSearch(one.Items, one.Weapons, one.Shields, one.Armors, one.Helmets, one.Gloves)
		singletonSearch.Jobs = newSearch(one.Jobs)
		singletonSearch.Maps = newSearch(one.Maps)
	} else if game == global.Two {
		singletonSearch.Abilities = newSearchFF2(two.Abilities)
		singletonSearch.Items = newSearch(two.Items, two.Weapons, two.Shields, two.Armors, two.Helmets, two.Gloves)
		singletonSearch.Jobs = newSearch(two.Jobs)
		singletonSearch.Maps = newSearch(two.Maps)
	} else if game == global.Three {

	} else if game == global.Four {

	} else if game == global.Five {

	} else { // Six

	}
}

func newSearch(nvss ...[]models.NameValue) *Search {
	s := &Search{
		Text:    widget.NewTextGrid(),
		Input:   widget.NewEntry(),
		Results: widget.NewTextGrid(),
	}
	s.onFilterChanged(nvss...)
	var sb strings.Builder
	for _, nvs := range nvss {
		for _, nv := range nvs {
			sb.WriteString(fmt.Sprintf("%d: %s\n", nv.Value, nv.Name))
		}
	}
	s.Text.SetText(sb.String())
	return s
}

func newSearchFF2(nvss ...[]models.NameValue) *Search {
	s := &Search{
		Text:    widget.NewTextGrid(),
		Input:   widget.NewEntry(),
		Results: widget.NewTextGrid(),
	}
	s.onFilterChanged(nvss...)
	m := make(map[string]bool)
	var r []struct {
		name     string
		from, to int
	}
	for _, nvs := range nvss {
		for _, nv := range nvs {
			n := strings.Split(nv.Name, " ")[0]
			if _, ok := m[n]; !ok {
				m[n] = true
				r = append(r, struct {
					name     string
					from, to int
				}{name: n, from: nv.Value, to: nv.Value + 15})
			}
		}
	}
	var sb strings.Builder
	for _, i := range r {
		sb.WriteString(fmt.Sprintf("%d-%d: %s\n", i.from, i.to, i.name))
	}
	s.Text.SetText(sb.String())
	return s
}

func (s *Search) onFilterChanged(nvss ...[]models.NameValue) {
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
}

func (s *Search) Filter() fyne.CanvasObject {
	if s == nil {
		return container.NewStack()
	}
	return container.NewBorder(s.Input, nil, nil, nil, container.NewVScroll(s.Results))
}

func (s *Search) Fields() fyne.CanvasObject {
	if s == nil {
		return container.NewStack()
	}
	return container.NewVScroll(s.Text)
}
