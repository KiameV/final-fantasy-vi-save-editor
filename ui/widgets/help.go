package widgets

import (
	"ffvi_editor/models/consts/pr"
	"ffvi_editor/models/consts/snes"
	"github.com/aarzilli/nucular"
)

var (
	prHelperInv *Helpers
	prHelperEq  *Helpers
	snesHelper  *Helpers
)

type Helpers struct {
	MiscItemsHelp      nucular.TextEditor
	WeaponShieldHelp1  nucular.TextEditor
	WeaponShieldHelp2  nucular.TextEditor
	HelmetArmorHelp1   nucular.TextEditor
	HelmetArmorHelp2   nucular.TextEditor
	RelicHelp1         nucular.TextEditor
	RelicHelp2         nucular.TextEditor
	ImportantItemsHelp nucular.TextEditor
}

func GetSnesHelpers() *Helpers {
	if snesHelper == nil {
		snesHelper = &Helpers{}
		snesHelper.snesInit()
	}
	return snesHelper
}

func GetPrInvHelpers() *Helpers {
	if prHelperInv == nil {
		prHelperInv = &Helpers{}
		prHelperInv.prInvInit()
	}
	return prHelperInv
}

func GetPrEqHelpers() *Helpers {
	if prHelperEq == nil {
		prHelperEq = &Helpers{}
		prHelperEq.prEqInit()
	}
	return prHelperEq
}

func (h *Helpers) snesInit() {
	InitReadOnlyText(&h.MiscItemsHelp, snes.EmptyText+snes.ItemsText)
	InitReadOnlyText(&h.WeaponShieldHelp1, snes.EmptyText+snes.WeaponShieldText1)
	InitReadOnlyText(&h.WeaponShieldHelp2, snes.EmptyText+snes.WeaponShieldText2)
	InitReadOnlyText(&h.HelmetArmorHelp1, snes.EmptyText+snes.HelmetArmorText1)
	InitReadOnlyText(&h.HelmetArmorHelp2, snes.EmptyText+snes.HelmetArmorText2)
	InitReadOnlyText(&h.RelicHelp1, snes.EmptyText+snes.RelicText1)
	InitReadOnlyText(&h.RelicHelp2, snes.RelicText2Header+snes.RelicText2)
}

func (h *Helpers) prInvInit() {
	InitReadOnlyText(&h.MiscItemsHelp, pr.EmptyText+pr.ItemsText)
	InitReadOnlyText(&h.WeaponShieldHelp1, pr.EmptyText+pr.WeaponShieldText1)
	InitReadOnlyText(&h.WeaponShieldHelp2, pr.EmptyText+pr.WeaponShieldText2)
	InitReadOnlyText(&h.HelmetArmorHelp1, pr.EmptyText+pr.HelmetArmorText1)
	InitReadOnlyText(&h.HelmetArmorHelp2, pr.EmptyText+pr.HelmetArmorText2)
	InitReadOnlyText(&h.RelicHelp1, pr.EmptyText+pr.RelicText1)
	InitReadOnlyText(&h.RelicHelp2, pr.RelicText2Header+pr.RelicText2)
	InitReadOnlyText(&h.ImportantItemsHelp, pr.ImportantItemsText)
}

func (h *Helpers) prEqInit() {
	InitReadOnlyText(&h.MiscItemsHelp, pr.EmptyText+pr.ItemsText)
	InitReadOnlyText(&h.WeaponShieldHelp1, pr.EmptyWeaponShield+pr.WeaponShieldText1)
	InitReadOnlyText(&h.WeaponShieldHelp2, pr.EmptyWeaponShield+pr.WeaponShieldText2)
	InitReadOnlyText(&h.HelmetArmorHelp1, pr.EmptyHelmet+pr.HelmetArmorText1)
	InitReadOnlyText(&h.HelmetArmorHelp2, pr.EmptyArmor+pr.HelmetArmorText2)
	InitReadOnlyText(&h.RelicHelp1, pr.EmptyRelic+pr.RelicText1)
	InitReadOnlyText(&h.RelicHelp2, pr.RelicText2Header+pr.RelicText2)
}

func InitReadOnlyText(tb *nucular.TextEditor, text string) {
	tb.Flags = nucular.EditBox | nucular.EditMultiline
	tb.SingleLine = false
	tb.Text([]rune(text))
	tb.Flags |= nucular.EditReadOnly
	tb.Flags |= nucular.EditNoHorizontalScroll
}
