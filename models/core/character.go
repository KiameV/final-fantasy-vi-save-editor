package core

import (
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
		Abilities            []*Ability
	}
	Character struct {
		AdditionalParams
		ID              int
		JobID           int
		IsEnabled       bool
		Name            string
		Level           int
		Exp             int
		CurrentHP       int
		CurrentMP       int
		Equipment       []*models.IdCount
		OwnedAbilityIds []int
		FF1             struct {
			TrainedAbilities [8][3]*Ability `json:"trainedAbilities"`
		}
		Commands *Commands
	}
	Ability struct {
		ID         int `json:"abilityId"`
		ContentID  int `json:"contentID"`
		SkillLevel int `json:"skillLevel"`
	}
	Characters struct {
		characters []*Character
	}
	Commands struct {
		EnableSave bool
		commands   []*int
	}
)

func NewCharacters(size int) *Characters {
	return &Characters{
		characters: make([]*Character, 0, size),
	}
}

func NewCharacter() *Character {
	return &Character{}
}

func (c *Character) AddAbility(ability *Ability) {
	c.Abilities = append(c.Abilities, ability)
}

func (c *Character) RemoveAbility(index int) {
	c.Abilities = append(c.Abilities[:index], c.Abilities[index+1:]...)
}

func (c *Character) AbilitiesToSave() []*Ability {
	a := make([]*Ability, 0, len(c.Abilities))
	for _, i := range c.Abilities {
		if i.ID != 0 {
			a = append(a, i)
		}
	}
	return a
}

func (c *Characters) All() []*Character {
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

func (c *Characters) AddCharacter(character *Character) {
	c.characters = append(c.characters, character)
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

func NewCommands(size int) *Commands {
	c := &Commands{
		commands: make([]*int, size),
	}
	for i := 0; i < size; i++ {
		c.commands[i] = new(int)
	}
	return c
}

func (c *Commands) All() []any {
	commands := make([]any, len(c.commands))
	for i, j := range c.commands {
		commands[i] = *j
	}
	return commands
}

func (c *Commands) Get(i int) *int {
	return c.commands[i]
}

func (c *Commands) Len() int {
	return len(c.commands)
}

func (c *Commands) Set(i int, value int) {
	*c.commands[i] = value
}
