package finder

import (
	"cmp"
	"slices"

	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/models"
	"pixel-remastered-save-editor/models/core"
	one "pixel-remastered-save-editor/models/core/ff1/consts"
	two "pixel-remastered-save-editor/models/core/ff2/consts"
	three "pixel-remastered-save-editor/models/core/ff3/consts"
	four "pixel-remastered-save-editor/models/core/ff4/consts"
	five "pixel-remastered-save-editor/models/core/ff5/consts"
	six "pixel-remastered-save-editor/models/core/ff6/consts"
)

type (
	Find    func(int) (string, bool)
	Finders interface {
		Abilities(int) (string, bool)
		Characters(int) (string, bool)
		Commands(int) (string, bool)
		Items(int) (string, bool)
		ImportantItems(int) (string, bool)
		Jobs(int) (string, bool)
		Maps(int) (string, bool)
	}
	finders struct {
		abilities  map[int]string
		characters map[int]string
		commands   map[int]string
		items      map[int]string
		important  map[int]string
		jobs       map[int]string
		maps       map[int]string
	}
)

var (
	singletonFinder *finders
)

func Load(game global.Game, characters []*core.Character) {
	if game == global.One {
		singletonFinder = &finders{
			abilities: nameLookup(one.Abilities),
			commands:  nameLookup(one.Commands),
			important: nameLookup(one.ImportantItems),
			items:     nameLookup(one.Items, one.Weapons, one.Shields, one.Armors, one.Helmets, one.Gloves),
			jobs:      nameLookup(one.Jobs),
			maps:      nameLookup(),
		}
	} else if game == global.Two {
		singletonFinder = &finders{
			abilities: nameLookup(two.Abilities),
			commands:  nameLookup(two.Commands),
			important: nameLookup(),
			items:     nameLookup(two.Items, two.Weapons, two.Shields, two.Armors, two.Helmets, two.Gloves),
			jobs:      nameLookup(two.Jobs),
			maps:      nameLookup(two.Maps),
		}
	} else if game == global.Three {
		singletonFinder = &finders{
			abilities: nameLookup(three.Abilities, three.WhiteMagic, three.BlackMagic, three.SummonMagic),
			commands:  nameLookup(three.Commands),
			important: nameLookup(three.ImportantItems),
			items:     nameLookup(three.Items, three.Weapons, three.Shields, three.Armors, three.Helmets, three.Hands),
			jobs:      nameLookup(three.Jobs),
			maps:      nameLookup(three.Maps),
		}
	} else if game == global.Four {
		singletonFinder = &finders{
			abilities: nameLookup(four.Abilities, four.WhiteMagic, four.BlackMagic, four.SummonMagic),
			commands:  nameLookup(four.Commands),
			important: nameLookup(four.ImportantItems),
			items:     nameLookup(four.Items, four.Weapons, four.Shields, four.Armors, four.Helmets, four.Hands),
			jobs:      nameLookup(four.Jobs),
			maps:      nameLookup(four.Maps),
		}
	} else if game == global.Five {
		singletonFinder = &finders{
			abilities: nameLookup(five.Abilities, five.WhiteMagic, five.BlackMagic, five.SummonMagic, five.TimeMagic),
			commands:  nameLookup(five.Commands),
			important: nameLookup(five.ImportantItems),
			items:     nameLookup(five.Items, five.Weapons, five.Shields, five.Armors, five.Helmets, five.Hands),
			jobs:      nameLookup(five.Jobs),
			maps:      nameLookup(five.Maps),
		}
	} else { // Six
		var temp []models.NameValue
		singletonFinder = &finders{
			abilities: nameLookup(six.Blitzes, six.Dances, six.Lores, six.Bushidos, six.Rages),
			commands:  nameLookup(temp),
			important: nameLookup(temp),
			items:     six.ItemsByID,
			jobs:      nameLookup(six.Jobs),
			maps:      nameLookup(temp),
		}
	}
	singletonFinder.characters = make(map[int]string)
	for _, c := range characters {
		singletonFinder.characters[c.Base.ID] = c.Base.Name
	}
}

func Abilities(i int) (s string, b bool) {
	return singletonFinder.Abilities(i)
}

func (f finders) Abilities(i int) (s string, b bool) {
	s, b = f.abilities[i]
	return
}

func Characters(i int) (string, bool) {
	return singletonFinder.Characters(i)
}

func (f finders) Characters(i int) (s string, b bool) {
	s, b = f.characters[i]
	return
}

func Commands(i int) (s string, b bool) {
	return singletonFinder.Commands(i)
}

func (f finders) Commands(i int) (s string, b bool) {
	s, b = f.commands[i]
	return
}

func Items(i int) (s string, b bool) {
	return singletonFinder.Items(i)
}

func (f finders) Items(i int) (s string, b bool) {
	s, b = f.items[i]
	return
}

func ImportantItems(i int) (s string, b bool) {
	return singletonFinder.ImportantItems(i)
}

func (f finders) ImportantItems(i int) (s string, b bool) {
	s, b = f.important[i]
	return
}

func Jobs(i int) (s string, b bool) {
	return singletonFinder.Jobs(i)
}

func (f finders) Jobs(i int) (s string, b bool) {
	s, b = f.jobs[i]
	return
}

func Maps(i int) (s string, b bool) {
	return singletonFinder.Maps(i)
}

func (f finders) Maps(i int) (s string, b bool) {
	s, b = f.maps[i]
	return
}

func Get() Finders {
	return singletonFinder
}

func nameLookup(args ...[]models.NameValue) map[int]string {
	m := make(map[int]string)
	for _, a := range args {
		for _, i := range a {
			m[i.Value] = i.Name
		}
	}
	return m
}

func AllCharacters() (v []models.NameValue) {
	v = make([]models.NameValue, 0, len(singletonFinder.characters))
	for i, n := range singletonFinder.characters {
		v = append(v, models.NewNameValue(n, i))
	}
	slices.SortFunc(v, func(i, j models.NameValue) int {
		return cmp.Compare(i.Value, j.Value)
	})
	return
}

func AllJobs() (v map[int]string) {
	return singletonFinder.jobs
}
