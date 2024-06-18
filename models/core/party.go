package core

import (
	"sort"
)

var (
	EmptyPartyMember = &Member{
		CharacterID: 0,
		Name:        "[Empty]",
	}
)

type (
	Member struct {
		CharacterID int `json:"characterId"`
		Name        string
		// EnableEquipment bool
	}
	Party struct {
		Members       []*Member
		Possible      map[string]*Member
		PossibleNames []string
		// PossibleNamesWithNPCs []string
		// Enabled bool
		// IncludeNPCs bool
	}
)

func NewParty(memberCount int) *Party {
	return &Party{
		Members: make([]*Member, memberCount),
	}
}

func (p *Party) AddPossibleMember(c *Character) {
	if p != nil {
		if _, found := p.Possible[c.Name]; !found || c.IsEnabled {
			p.Possible[c.Name] = &Member{
				CharacterID: c.ID,
				Name:        c.Name,
			}
			if !found {
				p.PossibleNames = append(p.PossibleNames, c.Name)
				sort.Strings(p.PossibleNames)
			}
		}
	}
}

func (p *Party) GetPossibleIndex(m *Member) int {
	if p != nil {
		for i, po := range p.PossibleNames {
			if m.Name == po {
				return i
			}
		}
	}
	return 0
}

func (p *Party) SetMember(slot int, member *Member) {
	if p != nil {
		p.Members[slot] = member
	}
}
