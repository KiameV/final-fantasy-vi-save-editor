package widgets

import (
	"ffvi_editor/models/consts"
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
	InitReadOnlyText(&MiscItemsHelp, consts.EmptyText+consts.ItemsText)
	InitReadOnlyText(&WeaponShieldHelp1, consts.EmptyText+consts.WeaponShieldText1)
	InitReadOnlyText(&WeaponShieldHelp2, consts.EmptyText+consts.WeaponShieldText2)
	InitReadOnlyText(&HelmetArmorHelp1, consts.EmptyText+consts.HelmetArmorText1)
	InitReadOnlyText(&HelmetArmorHelp2, consts.EmptyText+consts.HelmetArmorText2)
	InitReadOnlyText(&RelicHelp1, consts.EmptyText+consts.RelicText1)
	InitReadOnlyText(&RelicHelp2, consts.RelicText2Header+consts.RelicText2)
}

func InitReadOnlyText(tb *nucular.TextEditor, text string) {
	tb.Flags = nucular.EditBox | nucular.EditMultiline
	tb.SingleLine = false
	tb.Text([]rune(text))
	tb.Flags |= nucular.EditReadOnly
	tb.Flags |= nucular.EditNoHorizontalScroll
}
