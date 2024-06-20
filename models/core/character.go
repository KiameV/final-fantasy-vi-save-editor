package core

import (
	"errors"
	"slices"

	"pixel-remastered-save-editor/models"
)

type (
	AdditionalParams struct {
		MaxHP                int
		MaxMP                int
		Power                int
		Vitality             int
		Agility              int
		Weight               int
		Intelligence         int
		Spirit               int
		Attack               int
		Defense              int
		AbilityDefence       int
		AbilityEvasionRate   int
		Magic                int
		Luck                 int
		AccuracyRate         int
		EvasionRate          int
		AbilityDisturbedRate int
		CriticalRate         int
		DamageDirmeter       int
		AbilityDefenseRate   int
		AccuracyCount        int
		EvasionCount         int
		DefenseCount         int
		MagicDefenseCount    int
		Abilities            []Ability
	}
	Character struct {
		AdditionalParams
		ID        int
		JobID     int
		IsEnabled bool
		Name      string
		Level     int
		Exp       int
		CurrentHP int
		CurrentMP int
		Equipment []*models.IdCount
	}
	Ability struct {
		ID         int
		ContentID  int
		SkillLevel int
	}
	Characters struct {
		characters []*Character
	}
)

func NewCharacters(size int) *Characters {
	return &Characters{
		characters: make([]*Character, size),
	}
}

func (c *Character) AddAbility(ability Ability) {
	c.Abilities = append(c.Abilities, ability)
}

func (c *Characters) Characters() []*Character {
	return c.characters
}

func (c *Characters) Names() []string {
	names := make([]string, 0, len(c.characters))
	for _, character := range c.characters {
		if character != nil {
			names = append(names, character.Name)
		}
	}
	slices.Sort(names)
	return names
}

func (c *Characters) SetCharacter(i int, character *Character) error {
	if i < 0 || i >= len(c.characters) {
		return errors.New("index out of range")
	}
	c.characters[i] = character
	return nil
}

func (c *Characters) GetByID(id int) (character *Character, found bool) {
	for _, character = range c.characters {
		if character != nil && character.ID == id {
			return character, true
		}
	}
	return nil, false
}

func (c *Characters) GetByName(name string) (character *Character, found bool) {
	for _, character = range c.characters {
		if character != nil && character.Name == name {
			return character, true
		}
	}
	return nil, false
}
