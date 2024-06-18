package core

import (
	"pixel-remastered-save-editor/models/consts"
)

type Command struct {
	consts.NameValue
	SortedIndex int
}
