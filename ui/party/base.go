package party

import (
	"ffvi_editor/models/pr"
	"ffvi_editor/ui"
	"fmt"
	"github.com/aarzilli/nucular"
)

type partyUI struct {
	selected [4]int
}

func NewUI() ui.UI {
	return &partyUI{}
}

func (u *partyUI) Draw(w *nucular.Window) {
	p := pr.GetParty()
	w.Row(24).Static(260)
	w.CheckboxText("Enable", &p.Enabled)
	w.Row(10).Static()
	for i := 0; i < 4; i++ {
		u.drawRow(w, i, p, &u.selected[i])
	}
}

func (u *partyUI) drawRow(w *nucular.Window, slot int, p *pr.Party, selected *int) {
	w.Row(24).Static(60, 200, 200, 60)
	w.Label(fmt.Sprintf("Member: %d", slot+1), "LC")
	if i := w.ComboSimple(p.PossibleNames, *selected, 12); i != *selected {
		*selected = i
		p.Members[slot] = p.Possible[p.PossibleNames[i]]
	}
	if *selected > 0 {
		pn := p.PossibleNames[*selected]
		if pn != pr.EmptyPartyMember.Name {
			c := pr.GetCharacter(pn)
			if !c.IsEnabled {
				w.CheckboxText("Enable?", &c.IsEnabled)
			}
		}
	}
}

func (u *partyUI) Refresh() {
	p := pr.GetParty()
	for i, m := range p.Members {
		u.selected[i] = p.GetPossibleIndex(m)
	}
}

func (u *partyUI) Name() string {
	return "Party"
}

func (u *partyUI) Behavior() ui.Behavior {
	return ui.PrShow
}
