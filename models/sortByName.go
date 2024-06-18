package models

import (
	"sort"

	"pixel-remastered-save-editor/models/core"
)

func SortByName(nsms []*core.NameSlotMask8) []*core.NameSlotMask8 {
	var (
		lookup      = make(map[string]*core.NameSlotMask8)
		sortedNames = make([]string, len(nsms))
		sorted      = make([]*core.NameSlotMask8, len(nsms))
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

func SortByNameChecked(nsms []*core.NameValueChecked) []*core.NameValueChecked {
	var (
		lookup      = make(map[string]*core.NameValueChecked)
		sortedNames = make([]string, len(nsms))
		sorted      = make([]*core.NameValueChecked, len(nsms))
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

func SortByNameValue(nvs []*core.NameValue) []*core.NameValue {
	var (
		lookup      = make(map[string]*core.NameValue)
		sortedNames = make([]string, len(nvs))
		sorted      = make([]*core.NameValue, len(nvs))
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
