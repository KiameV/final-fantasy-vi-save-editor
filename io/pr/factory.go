package pr

import (
	j "gitlab.com/c0b/go-ordered-json"
)

type PR struct {
	data       []byte
	Base       *j.OrderedMap
	UserData   *j.OrderedMap
	Characters []*j.OrderedMap
}

func NewPR() *PR {
	return &PR{
		Base:       j.NewOrderedMap(),
		UserData:   j.NewOrderedMap(),
		Characters: make([]*j.OrderedMap, 40),
	}
}
