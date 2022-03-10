package widgets

import (
	"ffvi_editor/models/consts/snes"
	"github.com/aarzilli/nucular"
)

var (
	MiscItemsHelp     nucular.TextEditor
	WeaponShieldHelp1 nucular.TextEditor
	WeaponShieldHelp2 nucular.TextEditor
	HelmetArmorHelp1  nucular.TextEditor
	HelmetArmorHelp2  nucular.TextEditor
	RelicHelp1        nucular.TextEditor
	RelicHelp2        nucular.TextEditor
)

func init() {
	InitReadOnlyText(&MiscItemsHelp, snes.EmptyText+snes.ItemsText)
	InitReadOnlyText(&WeaponShieldHelp1, snes.EmptyText+snes.WeaponShieldText1)
	InitReadOnlyText(&WeaponShieldHelp2, snes.EmptyText+snes.WeaponShieldText2)
	InitReadOnlyText(&HelmetArmorHelp1, snes.EmptyText+snes.HelmetArmorText1)
	InitReadOnlyText(&HelmetArmorHelp2, snes.EmptyText+snes.HelmetArmorText2)
	InitReadOnlyText(&RelicHelp1, snes.EmptyText+snes.RelicText1)
	InitReadOnlyText(&RelicHelp2, snes.RelicText2Header+snes.RelicText2)
}

func InitReadOnlyText(tb *nucular.TextEditor, text string) {
	tb.Flags = nucular.EditBox | nucular.EditMultiline
	tb.SingleLine = false
	tb.Text([]rune(text))
	tb.Flags |= nucular.EditReadOnly
	tb.Flags |= nucular.EditNoHorizontalScroll
}
