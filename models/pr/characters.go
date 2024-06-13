package pr

import (
	"ffvi_editor/models"
	"ffvi_editor/models/consts"
	"ffvi_editor/models/consts/pr"
)

var Characters []*models.Character

func init() {
	Characters = make([]*models.Character, len(pr.Characters))
	// defaultCommand := consts.CommandLookupByValue[0xFF]
	for i, name := range pr.Characters {
		o, ok := CharacterOffsetByName[name]
		if !ok {
			panic("failed to load character " + name)
		}
		defaultCommand := pr.CommandLookupByValue[4]
		c := &models.Character{
			ID:                 o.ID,
			RootName:           name,
			Name:               name,
			StatusEffects:      consts.NewStatusEffects(),
			IsNPC:              o.IsNPC,
			EnableCommandsSave: false,
			Commands: []*models.Command{
				defaultCommand,
				defaultCommand,
				defaultCommand,
				defaultCommand,
				defaultCommand,
				defaultCommand,
				defaultCommand,
				defaultCommand,
				defaultCommand,
			},
		}
		c.SpellsByIndex, c.SpellsSorted, c.SpellsByID = NewSpells()
		Characters[i] = c
	}
}

func GetCharacter(name string) (c *models.Character) {
	for _, c = range Characters {
		if c.RootName == name {
			break
		}
	}
	return
}

func GetCharacterByID(id int) (c *models.Character) {
	for _, c = range Characters {
		if c.ID == id {
			break
		}
	}
	return
}

func CharacterNamesHumanSelect() []string {
	return []string{
		"Celes",
		"Cyan",
		"Edgar",
		"Gau",
		"Gogo",
		"Locke",
		"Mog",
		"Relm",
		"Sabin",
		"Setzer",
		"Shadow",
		"Strago",
		"Terra",
		"Umaro",
		"??????",
		"Banon",
		"Biggs",
		"Cosmog",
		"Leo",
		"Maduin",
		"Moggie",
		"Moghan",
		"Moglin",
		"Mogret",
		"Mogsy",
		"Moguel",
		"Mogwin",
		"Molulu",
		"Mugmug",
		"Wedge",
	}
}
