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
	three "pixel-remastered-save-editor/models/core/ff3/consts"
	four "pixel-remastered-save-editor/models/core/ff4/consts"
	five "pixel-remastered-save-editor/models/core/ff5/consts"
	"pixel-remastered-save-editor/models/finder"
)

type (
	Searches struct {
		Abilities      *Search
		Characters     *Search
		Commands       *Search
		Equipment      *Search
		Items          *Search
		ImportantItems *Search
		ItemsEquipment *Search
		Jobs           *Search
		Maps           *Search
	}
	Search struct {
		Text    *widget.TextGrid
		Input   *widget.Entry
		Results *widget.TextGrid
	}
)

var _s *Searches

func GetSearches() *Searches {
	return _s
}

func Load(game global.Game) {
	_s = &Searches{
		Characters: newSearch(finder.AllCharacters()),
	}
	if game == global.One {
		_s.Abilities = newSearch(one.Abilities)
		_s.Commands = newSearch(one.Commands)
		_s.Equipment = newSearch(one.Weapons, one.Shields, one.Armors, one.Helmets, one.Gloves)
		_s.Items = newSearch(one.Items)
		_s.ImportantItems = newSearch(one.ImportantItems)
		_s.ItemsEquipment = newSearch(one.Items, one.Weapons, one.Shields, one.Armors, one.Helmets, one.Gloves)
		_s.Jobs = newSearch(one.Jobs)
		_s.Maps = newSearch(one.Maps)
	} else if game == global.Two {
		_s.Abilities = newSearchFF2(two.Abilities)
		_s.Commands = newSearch(two.Commands)
		_s.Equipment = newSearch(two.Weapons, two.Shields, two.Armors, two.Helmets, two.Gloves)
		_s.ItemsEquipment = newSearch(two.Items, two.Weapons, two.Shields, two.Armors, two.Helmets, two.Gloves)
		_s.Items = newSearch(two.Items)
		_s.Jobs = newSearch(two.Jobs)
		_s.Maps = newSearch(two.Maps)
	} else if game == global.Three {
		_s.Abilities = newSearch(three.Abilities, three.WhiteMagic, three.BlackMagic, three.SummonMagic)
		_s.Commands = newSearch(three.Commands)
		_s.Equipment = newSearch(three.Weapons, three.Shields, three.Armors, three.Helmets, three.Hands)
		_s.Items = newSearch(three.Items)
		_s.ImportantItems = newSearch(three.ImportantItems)
		_s.ItemsEquipment = newSearch(three.Items, three.Weapons, three.Shields, three.Armors, three.Helmets, three.Hands)
		_s.Jobs = newSearch(three.Jobs)
		_s.Maps = newSearch(three.Maps)
	} else if game == global.Four {
		_s.Abilities = newSearch(four.Abilities, four.WhiteMagic, four.BlackMagic, four.SummonMagic)
		_s.Commands = newSearch(four.Commands)
		_s.Equipment = newSearch(four.Weapons, four.Shields, four.Armors, four.Helmets, four.Hands)
		_s.Items = newSearch(four.Items)
		_s.ImportantItems = newSearch(four.ImportantItems)
		_s.ItemsEquipment = newSearch(four.Items, four.Weapons, four.Shields, four.Armors, four.Helmets, four.Hands)
		_s.Jobs = newSearch(four.Jobs)
		_s.Maps = newSearch(four.Maps)
	} else if game == global.Five {
		_s.Abilities = newSearch(five.Abilities, five.WhiteMagic, five.BlackMagic, five.SummonMagic, five.TimeMagic)
		_s.Commands = newSearch(five.Commands)
		_s.Equipment = newSearch(five.Weapons, five.Shields, five.Armors, five.Helmets, five.Hands)
		_s.Items = newSearch(five.Items)
		_s.ItemsEquipment = newSearch(five.Items, five.Weapons, five.Shields, five.Armors, five.Helmets, five.Hands)
		_s.Jobs = newSearch(five.Jobs)
		_s.Maps = newSearch(five.Maps)
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
