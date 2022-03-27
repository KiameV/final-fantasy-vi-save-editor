package pr

import (
	"fmt"
)

type Member struct {
	CharacterID int `json:"characterId"`
	Name        string
	//EnableEquipment bool
}

var party *Party

type Party struct {
	Members       [4]*Member
	Possible      []*Member
	PossibleNames []string
	Enabled       bool
}

func GetParty() *Party {
	if party == nil {
		party = &Party{
			Enabled: false,
		}
		party.AddPossibleMember(&Member{
			CharacterID: 0,
			Name:        "Empty",
		})
	}
	return party
}

func (p *Party) Clear() {
	for i := 0; i < len(p.Members); i++ {
		p.Members[i] = p.Possible[0]
	}
	p.Possible = p.Possible[:1]
	p.PossibleNames = p.PossibleNames[:1]
}

func (p *Party) AddPossibleMember(m *Member) {
	p.Possible = append(p.Possible, m)
	p.PossibleNames = append(p.PossibleNames, m.Name)
}

func (p *Party) SetMemberByID(slot int, characterID int) error {
	for _, m := range p.Possible {
		if characterID == m.CharacterID {
			p.Members[slot] = m
			return nil
		}
	}
	return fmt.Errorf("failed to find character %d in list of possible characters", characterID)
}

func (p *Party) SetMemberByName(slot int, name string) error {
	for _, m := range p.Possible {
		if name == m.Name {
			p.Members[slot] = m
			return nil
		}
	}
	return fmt.Errorf("failed to find %s in list of possible characters", name)
}

func (p *Party) GetPossibleIndex(m *Member) int {
	for i, po := range p.Possible {
		if m == po {
			return i
		}
	}
	return 0
}
