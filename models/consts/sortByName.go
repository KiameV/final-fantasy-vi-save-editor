package consts

import "sort"

func SortByName(nsms []*NameSlotMask8) []*NameSlotMask8 {
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

func SortByNameChecked(nsms []*NameValueChecked) []*NameValueChecked {
	var (
		lookup      = make(map[string]*NameValueChecked)
		sortedNames = make([]string, len(nsms))
		sorted      = make([]*NameValueChecked, len(nsms))
	)
	for i, s := range nsms {
		lookup[s.Name] = s
		sortedNames[i] = s.Name
	}
	sort.Strings(sortedNames)

	for i, s := range sortedNames {
		sorted[i] = lookup[s]
	}
	lookup = nil
	sortedNames = nil
	return sorted
}

func SortByNameValue(nvs []*NameValue) []*NameValue {
	var (
		lookup      = make(map[string]*NameValue)
		sortedNames = make([]string, len(nvs))
		sorted      = make([]*NameValue, len(nvs))
	)
	for i, s := range nvs {
		lookup[s.Name] = s
		sortedNames[i] = s.Name
	}
	sort.Strings(sortedNames)

	for i, s := range sortedNames {
		sorted[i] = lookup[s]
	}
	lookup = nil
	sortedNames = nil
	return sorted
}
