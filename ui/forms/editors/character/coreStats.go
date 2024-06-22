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
		isEnabled:  binding.BindBool(&c.IsEnabled),
	}
	e.ExtendBaseWidget(e)
	name := widget.NewEntryWithData(binding.BindString(&c.Name))
	name.Validator = nil
	job := inputs.NewIdEntryWithDataWithHint(&c.JobID, finder.Jobs)
	e.left = []fyne.CanvasObject{
		container.NewGridWithColumns(3, widget.NewLabel("Name:"), name),
		container.NewGridWithColumns(3, widget.NewLabel("Job:"), job.ID, job.Label),
		container.NewGridWithColumns(3, widget.NewLabel("Level:"), inputs.NewIntEntryWithData(&c.Level)),
		container.NewGridWithColumns(3, widget.NewLabel("Experience:"), inputs.NewIntEntryWithData(&c.Exp)),
		container.NewGridWithColumns(3, widget.NewLabel("HP Current:"), inputs.NewIntEntryWithData(&c.CurrentHP)),
		container.NewGridWithColumns(3, widget.NewLabel("HP Max:"), inputs.NewIntEntryWithData(&c.MaxHP)),
		container.NewGridWithColumns(3, widget.NewLabel("MP Current:"), inputs.NewIntEntryWithData(&c.CurrentMP)),
		container.NewGridWithColumns(3, widget.NewLabel("MP Max:"), inputs.NewIntEntryWithData(&c.MaxMP)),
		container.NewGridWithColumns(3, widget.NewLabel("Agility:"), inputs.NewIntEntryWithData(&c.Agility)),
		container.NewGridWithColumns(3, widget.NewLabel("Intelligence:"), inputs.NewIntEntryWithData(&c.Intelligence)),
		container.NewGridWithColumns(3, widget.NewLabel("Luck:"), inputs.NewIntEntryWithData(&c.Luck)),
		container.NewGridWithColumns(3, widget.NewLabel("Magic:"), inputs.NewIntEntryWithData(&c.Magic)),
		container.NewGridWithColumns(3, widget.NewLabel("Power:"), inputs.NewIntEntryWithData(&c.Power)),
		container.NewGridWithColumns(3, widget.NewLabel("Spirit:"), inputs.NewIntEntryWithData(&c.Spirit)),
		container.NewGridWithColumns(3, widget.NewLabel("Vitality:"), inputs.NewIntEntryWithData(&c.Vitality)),
	}
	e.right = []fyne.CanvasObject{
		container.NewGridWithColumns(3, widget.NewLabel("Ability Defense:"), inputs.NewIntEntryWithData(&c.AbilityDefence)),
		container.NewGridWithColumns(3, widget.NewLabel("Ability Defense Rate:"), inputs.NewIntEntryWithData(&c.AbilityDefenseRate)),
		container.NewGridWithColumns(3, widget.NewLabel("Ability Disturbed Rate:"), inputs.NewIntEntryWithData(&c.AbilityDisturbedRate)),
		container.NewGridWithColumns(3, widget.NewLabel("Ability Evasion Rate:"), inputs.NewIntEntryWithData(&c.AbilityEvasionRate)),
		container.NewGridWithColumns(3, widget.NewLabel("Accuracy Count:"), inputs.NewIntEntryWithData(&c.AccuracyCount)),
		container.NewGridWithColumns(3, widget.NewLabel("Accuracy Rate:"), inputs.NewIntEntryWithData(&c.AccuracyRate)),
		container.NewGridWithColumns(3, widget.NewLabel("Attack:"), inputs.NewIntEntryWithData(&c.Attack)),
		container.NewGridWithColumns(3, widget.NewLabel("Critical Rate:"), inputs.NewIntEntryWithData(&c.CriticalRate)),
		container.NewGridWithColumns(3, widget.NewLabel("Damage Dirmeter:"), inputs.NewIntEntryWithData(&c.DamageDirmeter)),
		container.NewGridWithColumns(3, widget.NewLabel("Defense:"), inputs.NewIntEntryWithData(&c.Defense)),
		container.NewGridWithColumns(3, widget.NewLabel("Defense Count:"), inputs.NewIntEntryWithData(&c.DefenseCount)),
		container.NewGridWithColumns(3, widget.NewLabel("Evasion Count:"), inputs.NewIntEntryWithData(&c.EvasionCount)),
		container.NewGridWithColumns(3, widget.NewLabel("Evasion Rate:"), inputs.NewIntEntryWithData(&c.EvasionRate)),
		container.NewGridWithColumns(3, widget.NewLabel("Magic Defense Count:"), inputs.NewIntEntryWithData(&c.MagicDefenseCount)),
		container.NewGridWithColumns(3, widget.NewLabel("Weight:"), inputs.NewIntEntryWithData(&c.Weight)),
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
