package pr

import (
	jo "gitlab.com/c0b/go-ordered-json"
)

type PR struct {
	//data       []byte
	Base       *jo.OrderedMap
	UserData   *jo.OrderedMap
	MapData    *jo.OrderedMap
	Characters []*jo.OrderedMap
	names      []unicodeNameReplace
	//fileEnd    string
}

func NewPR() *PR {
	return &PR{
		Base:       jo.NewOrderedMap(),
		UserData:   jo.NewOrderedMap(),
		MapData:    jo.NewOrderedMap(),
		Characters: make([]*jo.OrderedMap, 40),
	}
}

func (p *PR) HasUnicodeNames() bool {
	return len(p.names) > 0
}
