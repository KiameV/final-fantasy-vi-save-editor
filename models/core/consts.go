package core

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
		Items(int) (string, bool)
		ImportantItems(int) (string, bool)
		Map(int) (string, bool)
	}
	finders struct {
		abilities map[int]string
		items     map[int]string
		important map[int]string
		maps      map[int]string
	}
)

var (
	singletonFinder Finders
)

func LoadFinders(game global.Game) {
	if game == global.One {
		singletonFinder = &finders{
			abilities: nameLookup(one.Abilities),
			important: nameLookup(one.ImportantItems),
			items:     nameLookup(one.Items, one.Weapons, one.Shields, one.Armors, one.Helmets, one.Gloves),
			maps:      nameLookup(),
		}
	} else if game == global.Two {
		singletonFinder = &finders{
			abilities: nameLookup(),
			important: nameLookup(),
			items:     nameLookup(two.Weapons),
			maps:      nameLookup(),
		}
	} else if game == global.Three {
		singletonFinder = &finders{
			abilities: nameLookup(),
			important: nameLookup(),
			items:     nameLookup(three.Items, three.Weapons, three.Shields, three.Armors, three.Helmets, three.Hands),
			maps:      nameLookup(),
		}
	} else if game == global.Four {
		singletonFinder = &finders{
			abilities: nameLookup(),
			important: nameLookup(),
			items:     nameLookup(),
			maps:      nameLookup(),
		}
	} else if game == global.Five {
		singletonFinder = &finders{
			abilities: nameLookup(),
			important: nameLookup(),
			items:     nameLookup(),
			maps:      nameLookup(),
		}
	} else { // Six
		singletonFinder = &finders{
			abilities: nameLookup(six.Blitzes, six.Dances, six.Lores, six.Bushidos, six.Rages),
			important: nameLookup(),
			items:     six.ItemsByID,
			maps:      nameLookup(),
		}
	}
}

func FindAbilities(i int) (s string, b bool) {
	return singletonFinder.Abilities(i)
}

func (f finders) Abilities(i int) (s string, b bool) {
	s, b = f.abilities[i]
	return
}

func FindItems(i int) (s string, b bool) {
	return singletonFinder.Items(i)
}

func (f finders) Items(i int) (s string, b bool) {
	s, b = f.items[i]
	return
}

func FindImportantItems(i int) (s string, b bool) {
	return singletonFinder.ImportantItems(i)
}

func (f finders) ImportantItems(i int) (s string, b bool) {
	s, b = f.important[i]
	return
}

func FindMap(i int) (s string, b bool) {
	return singletonFinder.Map(i)
}

func (f finders) Map(i int) (s string, b bool) {
	s, b = f.maps[i]
	return
}

func GetFinders() Finders {
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
