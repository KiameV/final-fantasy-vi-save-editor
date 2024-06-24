package character

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/models/finder"
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
		isEnabled:  binding.BindBool(&c.Base.IsEnableCorps),
	}
	e.ExtendBaseWidget(e)
	name := widget.NewEntryWithData(binding.BindString(&c.Base.Name))
	name.Validator = nil
	job := inputs.NewIdEntryWithDataWithHint(&c.Base.JobID, finder.Jobs)
	e.left = []fyne.CanvasObject{
		container.NewGridWithColumns(3, widget.NewLabel("Name:"), name),
		container.NewGridWithColumns(3, widget.NewLabel("Job:"), job.ID, job.Label),
		container.NewGridWithColumns(3, widget.NewLabel("Level:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalLevel)),
		container.NewGridWithColumns(3, widget.NewLabel("Experience:"), inputs.NewIntEntryWithData(&c.Base.CurrentExp)),
		container.NewGridWithColumns(3, widget.NewLabel("HP Current:"), inputs.NewIntEntryWithData(&c.Parameters.CurrentHP)),
		container.NewGridWithColumns(3, widget.NewLabel("HP Max:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalMaxHp)),
		container.NewGridWithColumns(3, widget.NewLabel("MP Current:"), inputs.NewIntEntryWithData(&c.Parameters.CurrentMP)),
		container.NewGridWithColumns(3, widget.NewLabel("MP Max:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalMaxMp)),
		container.NewGridWithColumns(3, widget.NewLabel("Agility:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalAgility)),
		container.NewGridWithColumns(3, widget.NewLabel("Intelligence:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalIntelligence)),
		container.NewGridWithColumns(3, widget.NewLabel("Luck:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalLuck)),
		container.NewGridWithColumns(3, widget.NewLabel("Magic:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalMagic)),
		container.NewGridWithColumns(3, widget.NewLabel("Power:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalPower)),
		container.NewGridWithColumns(3, widget.NewLabel("Spirit:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalSpirit)),
		container.NewGridWithColumns(3, widget.NewLabel("Vitality:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalVitality)),
	}
	e.right = []fyne.CanvasObject{
		container.NewGridWithColumns(3, widget.NewLabel("Ability Defense:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalAbilityDefense)),
		container.NewGridWithColumns(3, widget.NewLabel("Ability Defense Rate:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalAbilityDefenseRate)),
		container.NewGridWithColumns(3, widget.NewLabel("Ability Disturbed Rate:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalAbilityDisturbedRate)),
		container.NewGridWithColumns(3, widget.NewLabel("Ability Evasion Rate:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalAbilityEvasionRate)),
		container.NewGridWithColumns(3, widget.NewLabel("Accuracy Count:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalAccuracyCount)),
		container.NewGridWithColumns(3, widget.NewLabel("Accuracy Rate:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalAccuracyRate)),
		container.NewGridWithColumns(3, widget.NewLabel("Attack:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalAttack)),
		container.NewGridWithColumns(3, widget.NewLabel("Critical Rate:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalCriticalRate)),
		container.NewGridWithColumns(3, widget.NewLabel("Damage Dirmeter:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalDamageDirmeter)),
		container.NewGridWithColumns(3, widget.NewLabel("Defense:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalDefense)),
		container.NewGridWithColumns(3, widget.NewLabel("Defense Count:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalDefenseCount)),
		container.NewGridWithColumns(3, widget.NewLabel("Evasion Count:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalEvasionCount)),
		container.NewGridWithColumns(3, widget.NewLabel("Evasion Rate:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalEvasionRate)),
		container.NewGridWithColumns(3, widget.NewLabel("Magic Defense Count:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalMagicDefenseCount)),
		container.NewGridWithColumns(3, widget.NewLabel("Weight:"), inputs.NewIntEntryWithData(&c.Parameters.AdditionalWeight)),
	}
	return e
}

func (e *Character) CreateRenderer() fyne.WidgetRenderer {
	search := inputs.GetSearches().Jobs
	return widget.NewSimpleRenderer(
		container.NewBorder(
			container.NewHBox(widget.NewCheckWithData("Enabled", e.isEnabled)),
			nil, nil, nil,
			container.NewVScroll(container.NewGridWithColumns(3,
				container.NewVBox(e.left...),
				container.NewVBox(e.right...),
				search.Fields()))))
}
