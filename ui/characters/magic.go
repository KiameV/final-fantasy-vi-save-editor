package characters

import (
	"ffvi_editor/models"
	"ffvi_editor/models/consts"
	f "ffvi_editor/ui/factory"
	"github.com/therecipe/qt/widgets"
)

var spells []*spellInput

type spellInput struct {
	*widgets.QSpinBox
	Name  string
	Spell *models.Spell
}

func (s *spellInput) OnChange(value int) {
	s.Spell.Value = uint8(value)
}

func createMagicLayout() widgets.QWidget_ITF {
	var (
		split = widgets.NewQHBoxLayout()
		left  = widgets.NewQWidget(nil, 0)
		sa    = widgets.NewQScrollArea(nil)
		ll    = widgets.NewQVBoxLayout()
		gr    *widgets.QGroupBox
		input *spellInput
	)

	sa.SetWidgetResizable(true)
	sa.SetWidget(left)
	sa.SetLayout(ll)
	left.SetLayout(ll)

	spells = make([]*spellInput, len(consts.SortedSpells))
	for i, s := range consts.SortedSpells {
		input = &spellInput{
			Name: s,
		}
		spells[i] = input

		input.QSpinBox, gr = f.CreateUint8Input(s, input.OnChange, 0, 255)
		ll.AddWidget(gr, 0, 0)
	}

	split.AddWidget(sa, 0, 0)
	split.AddWidget(f.CreateReadOnlyTextBox(magicText), 0, 0)

	w := widgets.NewQWidget(nil, 0)
	w.SetLayout(split)
	return w
}

func updateMagic() {
	for _, s := range spells {
		s.Spell = selectedCharacter.SpellsLookup[s.Name]
	}
}

const magicText = `
<b>Values:</b><br>
- 0 - 100 is the 'percent learned'<br>
- 255 means the spell is learned and can be used.`
