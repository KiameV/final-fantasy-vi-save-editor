package models

import "ffvi_editor/models/consts"

type Command struct {
	consts.NameValue
	SortedIndex int
}
