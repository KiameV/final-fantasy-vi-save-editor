package finder

import (
	"pixel-remastered-save-editor/global"
	"pixel-remastered-save-editor/models"
	one "pixel-remastered-save-editor/models/core/ff1/consts"
	two "pixel-remastered-save-editor/models/core/ff2/consts"
	three "pixel-remastered-save-editor/models/core/ff3/consts"
	six "pixel-remastered-save-editor/models/core/ff6/consts"
)

type (
	Find    func(int) (string, bool)
	Finders interface {
		Abilities(int) (string, bool)
		Commands(int) (string, bool)
		Items(int) (string, bool)
		ImportantItems(int) (string, bool)
		Jobs(int) (string, bool)
		Maps(int) (string, bool)
	}
	finders struct {
		abilities map[int]string
		commands  map[int]string
		items     map[int]string
		important map[int]string
		jobs      map[int]string
		maps      map[int]string
	}
)

var (
	singletonFinder Finders
)

func Load(game global.Game) {
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
			abilities: nameLookup(),
			commands:  nameLookup(),
			important: nameLookup(),
			items:     nameLookup(two.Weapons),
			jobs:      nameLookup(),
			maps:      nameLookup(),
		}
	} else if game == global.Three {
		singletonFinder = &finders{
			abilities: nameLookup(),
			commands:  nameLookup(),
			important: nameLookup(),
			items:     nameLookup(three.Items, three.Weapons, three.Shields, three.Armors, three.Helmets, three.Hands),
			jobs:      nameLookup(),
			maps:      nameLookup(),
		}
	} else if game == global.Four {
		singletonFinder = &finders{
			abilities: nameLookup(),
			commands:  nameLookup(),
			important: nameLookup(),
			items:     nameLookup(),
			jobs:      nameLookup(),
			maps:      nameLookup(),
		}
	} else if game == global.Five {
		singletonFinder = &finders{
			abilities: nameLookup(),
			commands:  nameLookup(),
			important: nameLookup(),
			items:     nameLookup(),
			jobs:      nameLookup(),
			maps:      nameLookup(),
		}
	} else { // Six
		singletonFinder = &finders{
			abilities: nameLookup(six.Blitzes, six.Dances, six.Lores, six.Bushidos, six.Rages),
			commands:  nameLookup(),
			important: nameLookup(),
			items:     six.ItemsByID,
			jobs:      nameLookup(six.Jobs),
			maps:      nameLookup(),
		}
	}
}

func Abilities(i int) (s string, b bool) {
	return singletonFinder.Abilities(i)
}

func (f finders) Abilities(i int) (s string, b bool) {
	s, b = f.abilities[i]
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
