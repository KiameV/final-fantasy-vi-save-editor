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
		left      []*inputs.IdEntry
		right     []*inputs.IdEntry
	}
)

func NewCoreStats(c *core.Character) *Character {
	e := &Character{
		BaseWidget: widget.BaseWidget{},
		isEnabled:  binding.BindBool(&c.Base.IsEnableCorps),
	}
	e.ExtendBaseWidget(e)
	name := widget.NewEntryWithData(binding.BindString(&c.Base.Name))
	name.Validator = nil
	e.left = []*inputs.IdEntry{
		inputs.NewLabeledEntryWithData("Name:", name),
		inputs.NewLabeledIntEntryWithData("Level:", &c.Parameters.AdditionalLevel),
		inputs.NewLabeledIntEntryWithData("Experience:", &c.Base.CurrentExp),
		inputs.NewLabeledIntEntryWithData("HP Current:", &c.Parameters.CurrentHP),
		inputs.NewLabeledIntEntryWithData("HP Max:", &c.Parameters.AdditionalMaxHp),
		inputs.NewLabeledIntEntryWithData("MP Current:", &c.Parameters.CurrentMP),
		inputs.NewLabeledIntEntryWithData("MP Max:", &c.Parameters.AdditionalMaxMp),
		inputs.NewLabeledIntEntryWithData("Agility:", &c.Parameters.AdditionalAgility),
		inputs.NewLabeledIntEntryWithData("Intelligence:", &c.Parameters.AdditionalIntelligence),
		inputs.NewLabeledIntEntryWithData("Luck:", &c.Parameters.AdditionalLuck),
		inputs.NewLabeledIntEntryWithData("Magic:", &c.Parameters.AdditionalMagic),
		inputs.NewLabeledIntEntryWithData("Power:", &c.Parameters.AdditionalPower),
		inputs.NewLabeledIntEntryWithData("Spirit:", &c.Parameters.AdditionalSpirit),
		inputs.NewLabeledIntEntryWithData("Vitality:", &c.Parameters.AdditionalVitality),
	}
	e.right = []*inputs.IdEntry{
		inputs.NewLabeledIntEntryWithData("Ability Defense:", &c.Parameters.AdditionalAbilityDefense),
		inputs.NewLabeledIntEntryWithData("Ability Defense Rate:", &c.Parameters.AdditionalAbilityDefenseRate),
		inputs.NewLabeledIntEntryWithData("Ability Disturbed Rate:", &c.Parameters.AdditionalAbilityDisturbedRate),
		inputs.NewLabeledIntEntryWithData("Ability Evasion Rate:", &c.Parameters.AdditionalAbilityEvasionRate),
		inputs.NewLabeledIntEntryWithData("Accuracy Count:", &c.Parameters.AdditionalAccuracyCount),
		inputs.NewLabeledIntEntryWithData("Accuracy Rate:", &c.Parameters.AdditionalAccuracyRate),
		inputs.NewLabeledIntEntryWithData("Attack:", &c.Parameters.AdditionalAttack),
		inputs.NewLabeledIntEntryWithData("Critical Rate:", &c.Parameters.AdditionalCriticalRate),
		inputs.NewLabeledIntEntryWithData("Damage Dirmeter:", &c.Parameters.AdditionalDamageDirmeter),
		inputs.NewLabeledIntEntryWithData("Defense:", &c.Parameters.AdditionalDefense),
		inputs.NewLabeledIntEntryWithData("Defense Count:", &c.Parameters.AdditionalDefenseCount),
		inputs.NewLabeledIntEntryWithData("Evasion Count:", &c.Parameters.AdditionalEvasionCount),
		inputs.NewLabeledIntEntryWithData("Evasion Rate:", &c.Parameters.AdditionalEvasionRate),
		inputs.NewLabeledIntEntryWithData("Magic Defense Count:", &c.Parameters.AdditionalMagicDefenseCount),
		inputs.NewLabeledIntEntryWithData("Weight:", &c.Parameters.AdditionalWeight),
	}
	return e
}

func (e *Character) CreateRenderer() fyne.WidgetRenderer {
	l := len(e.left)
	r := len(e.right)
	s := l
	if r > l {
		s = r
	}
	rows := make([]fyne.CanvasObject, 0, s)
	for i := 0; i < s; i++ {
		row := container.NewGridWithColumns(4)
		if i < l {
			row.Add(e.left[i].Label)
			row.Add(e.left[i].ID)
		} else {
			row.Add(container.NewStack())
			row.Add(container.NewStack())
		}
		if i < r {
			row.Add(e.right[i].Label)
			row.Add(e.right[i].ID)
		} else {
			row.Add(container.NewStack())
			row.Add(container.NewStack())
		}
		rows = append(rows, container.NewPadded(row))
	}
	return widget.NewSimpleRenderer(
		container.NewBorder(
			widget.NewCheckWithData("Is Available", e.isEnabled),
			nil, container.NewVScroll(container.NewVBox(rows...)), nil))
}
