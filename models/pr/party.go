package pr

import (
	"fmt"
	"sort"
)

var EmptyPartyMember = &Member{
	CharacterID: 0,
	Name:        "[Empty]",
}

type Member struct {
	CharacterID int `json:"characterId"`
	Name        string
	//EnableEquipment bool
}

var party *Party

type Party struct {
	Members       [4]*Member
	Possible      map[string]*Member
	PossibleNames []string
	//PossibleNamesWithNPCs []string
	Enabled bool
	//IncludeNPCs bool
}

func GetParty() *Party {
	if party == nil {
		party = &Party{
			Enabled: false,
		}
		party.Clear()
	}
	return party
}

func (p *Party) Clear() {
	p.Possible = make(map[string]*Member)
	p.PossibleNames = make([]string, 0, 40)
	//p.PossibleNamesWithNPCs = make([]string, 0, 40)
	p.AddPossibleMember(EmptyPartyMember)
}

func (p *Party) AddPossibleMember(m *Member) {
	//c := GetCharacterByID(m.CharacterID)

	p.Possible[m.Name] = m

	//if !c.IsNPC {
	p.PossibleNames = append(p.PossibleNames, m.Name)
	sort.Strings(p.PossibleNames)
	//}

	//p.PossibleNamesWithNPCs = append(p.PossibleNamesWithNPCs, m.Name)
	//sort.Strings(p.PossibleNamesWithNPCs)
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
	names := p.PossibleNames
	//if p.IncludeNPCs {
	//	names = p.PossibleNamesWithNPCs
	//}

	for i, po := range names {
		if m.Name == po {
			return i
		}
	}
	return 0
}
