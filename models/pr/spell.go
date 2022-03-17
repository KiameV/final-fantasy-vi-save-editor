package pr

import (
	"ffvi_editor/models"
	"ffvi_editor/models/consts/pr"
)

func NewSpells() (byIndex []*models.Spell, alpha []*models.Spell, byID map[int]*models.Spell) {
	byIndex = make([]*models.Spell, len(pr.Spells))
	alpha = make([]*models.Spell, len(pr.Spells))
	byID = make(map[int]*models.Spell)
	lookup := make(map[string]*models.Spell)
	for i, s := range pr.Spells {
		spell := &models.Spell{
			Index: s.Value,
			Name:  s.Name,
			Value: 0,
		}
		byIndex[i] = spell
		lookup[s.Name] = spell
		byID[s.Value] = spell
	}

	for i, s := range pr.SortedSpells {
		alpha[i] = lookup[s.Name]
	}
	return
}
