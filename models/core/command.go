package core

import (
	"pixel-remastered-save-editor/models"
)

type Command struct {
	models.NameValue
	SortedIndex int
}
