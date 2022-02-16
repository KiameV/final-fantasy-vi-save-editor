package consts

import "sort"

func sortByName(nsms []*NameSlotMask8) []*NameSlotMask8 {
	var (
		lookup      = make(map[string]*NameSlotMask8)
		sortedNames = make([]string, len(nsms))
		sorted      = make([]*NameSlotMask8, len(nsms))
	)
	for i, s := range nsms {
		lookup[s.Name] = s
		sortedNames[i] = s.Name
	}
	sort.Strings(sortedNames)

	for i, s := range sortedNames {
		sorted[i] = lookup[s]
	}
	return sorted
}
