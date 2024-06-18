package character

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/ui/forms/inputs"
)

type (
	Character struct {
		widget.BaseWidget
		isEnabled binding.Bool
		left      []fyne.CanvasObject
		right     []fyne.CanvasObject
	}
)

func NewCoreStats(c *core.Character) *Character {
	e := &Character{
		BaseWidget: widget.BaseWidget{},
		isEnabled:  binding.BindBool(&c.IsEnabled),
	}
	e.ExtendBaseWidget(e)
	name := widget.NewEntryWithData(binding.BindString(&c.Name))
	name.Validator = nil
	e.left = []fyne.CanvasObject{
		inputs.NewLabeledEntry("Name:", name),
		inputs.NewLabeledEntry("Level:", inputs.NewIntEntryWithData(&c.Level)),
		inputs.NewLabeledEntry("Experience:", inputs.NewIntEntryWithData(&c.Exp)),
		inputs.NewLabeledEntry("HP Current:", inputs.NewIntEntryWithData(&c.CurrentHP)),
		inputs.NewLabeledEntry("HP Max:", inputs.NewIntEntryWithData(&c.MaxHP)),
		inputs.NewLabeledEntry("MP Current:", inputs.NewIntEntryWithData(&c.CurrentMP)),
		inputs.NewLabeledEntry("MP Max:", inputs.NewIntEntryWithData(&c.MaxMP)),
		inputs.NewLabeledEntry("Agility:", inputs.NewIntEntryWithData(&c.Agility)),
		inputs.NewLabeledEntry("Intelligence:", inputs.NewIntEntryWithData(&c.Intelligence)),
		inputs.NewLabeledEntry("Luck:", inputs.NewIntEntryWithData(&c.Luck)),
		inputs.NewLabeledEntry("Magic:", inputs.NewIntEntryWithData(&c.Magic)),
		inputs.NewLabeledEntry("Power:", inputs.NewIntEntryWithData(&c.Power)),
		inputs.NewLabeledEntry("Spirit:", inputs.NewIntEntryWithData(&c.Spirit)),
		inputs.NewLabeledEntry("Vitality:", inputs.NewIntEntryWithData(&c.Vitality)),
	}
	e.right = []fyne.CanvasObject{
		inputs.NewLabeledEntry("Ability Defense:", inputs.NewIntEntryWithData(&c.AbilityDefence)),
		inputs.NewLabeledEntry("Ability Defense Rate:", inputs.NewIntEntryWithData(&c.AbilityDefenseRate)),
		inputs.NewLabeledEntry("Ability Disturbed Rate:", inputs.NewIntEntryWithData(&c.AbilityDisturbedRate)),
		inputs.NewLabeledEntry("Ability Evasion Rate:", inputs.NewIntEntryWithData(&c.AbilityEvasionRate)),
		inputs.NewLabeledEntry("Accuracy Count:", inputs.NewIntEntryWithData(&c.AccuracyCount)),
		inputs.NewLabeledEntry("Accuracy Rate:", inputs.NewIntEntryWithData(&c.AccuracyRate)),
		inputs.NewLabeledEntry("Attack:", inputs.NewIntEntryWithData(&c.Attack)),
		inputs.NewLabeledEntry("Critical Rate:", inputs.NewIntEntryWithData(&c.CriticalRate)),
		inputs.NewLabeledEntry("Damage Dirmeter:", inputs.NewIntEntryWithData(&c.DamageDirmeter)),
		inputs.NewLabeledEntry("Defense:", inputs.NewIntEntryWithData(&c.Defense)),
		inputs.NewLabeledEntry("Defense Count:", inputs.NewIntEntryWithData(&c.DefenseCount)),
		inputs.NewLabeledEntry("Evasion Count:", inputs.NewIntEntryWithData(&c.EvasionCount)),
		inputs.NewLabeledEntry("Evasion Rate:", inputs.NewIntEntryWithData(&c.EvasionRate)),
		inputs.NewLabeledEntry("Magic Defense Count:", inputs.NewIntEntryWithData(&c.MagicDefenseCount)),
		inputs.NewLabeledEntry("Weight:", inputs.NewIntEntryWithData(&c.Weight)),
	}
	return e
}

func (e *Character) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		container.NewBorder(
			container.NewHBox(widget.NewCheckWithData("Enabled", e.isEnabled)),
			nil, nil, nil,
			container.NewVScroll(container.NewGridWithColumns(2,
				container.NewVBox(e.left...),
				container.NewVBox(e.right...)))))
}
