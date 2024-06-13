package editors

import (
	"fmt"
	"strings"

	"ffvi_editor/io/pr"
	"ffvi_editor/models"
	"ffvi_editor/ui/forms/inputs"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type (
	Equipment struct {
		widget.BaseWidget
		c       *models.Character
		search  *widget.Entry
		results *widget.TextGrid
	}
)

func NewEquipment(c *models.Character) *Equipment {
	e := &Equipment{
		c:       c,
		search:  widget.NewEntry(),
		results: widget.NewTextGrid(),
	}
	e.ExtendBaseWidget(e)
	e.search.OnChanged = func(s string) {
		s = strings.ToLower(s)
		if len(s) < 2 {
			e.results.SetText("")
			return
		}
		var sb strings.Builder
		if s != "" {
			for k, v := range pr.AllNormalItems {
				if strings.Contains(strings.ToLower(v), s) {
					sb.WriteString(fmt.Sprintf("%d - %s\n", k, v))
				}
			}
		}
		e.results.SetText(sb.String())
	}
	return e
}

func (e *Equipment) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(
		container.NewGridWithColumns(4,
			container.NewGridWithRows(3,
				container.NewVBox(
					inputs.NewLabeledIntEntryWithHint("Weapon:", inputs.NewIntEntryWithData(&e.c.Equipment.WeaponID)),
					inputs.NewLabeledIntEntryWithHint("Shield:", inputs.NewIntEntryWithData(&e.c.Equipment.ShieldID)),
				),
				container.NewVBox(
					inputs.NewLabeledIntEntryWithHint("Helmet:", inputs.NewIntEntryWithData(&e.c.Equipment.HelmetID)),
					inputs.NewLabeledIntEntryWithHint("Armor:", inputs.NewIntEntryWithData(&e.c.Equipment.ArmorID)),
				),
				container.NewVBox(
					inputs.NewLabeledIntEntryWithHint("Relic 1:", inputs.NewIntEntryWithData(&e.c.Equipment.Relic1ID)),
					inputs.NewLabeledIntEntryWithHint("Relic 2:", inputs.NewIntEntryWithData(&e.c.Equipment.Relic2ID)),
				),
			),
			container.NewGridWithRows(3,
				weaponsTextBox,
				helmetTextBox,
				relic1TextBox),
			container.NewGridWithRows(3,
				shieldsTextBox,
				armorTextBox,
				relic2TextBox),
			container.NewBorder(
				inputs.NewLabeledEntry("Find By Name:", e.search), nil, nil, nil,
				container.NewVScroll(e.results))))
}
