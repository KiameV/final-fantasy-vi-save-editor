package consts

import (
	"pixel-remastered-save-editor/models"
)

var (
	Commands = []models.NameValue{
		models.NewNameValue("Attack", 1),
		models.NewNameValue("Defend", 2),
		models.NewNameValue("Flee", 3),
		models.NewNameValue("Row", 4),
	}
)
