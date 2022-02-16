package character

import (
	"ffvi_editor/ui/widgets"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/rect"
)

type equipmentUI struct {
	weaponID nucular.TextEditor
	shieldID nucular.TextEditor
	helmetID nucular.TextEditor
	armorID  nucular.TextEditor
	relic1ID nucular.TextEditor
	relic2ID nucular.TextEditor

	countLast int
}

func newEquipmentUI() widget {
	u := &equipmentUI{}

	u.initEquipmentTB(&u.weaponID)
	u.initEquipmentTB(&u.shieldID)
	u.initEquipmentTB(&u.helmetID)
	u.initEquipmentTB(&u.armorID)
	u.initEquipmentTB(&u.relic1ID)
	u.initEquipmentTB(&u.relic2ID)

	return u
}

func (u *equipmentUI) Draw(w *nucular.Window) {
	count := 18
	w.Row(610).SpaceBegin(u.countLast)
	u.drawPair(w, 0, "Weapon ID:", &u.weaponID, "Shield ID:", &u.shieldID, &widgets.WeaponShieldHelp1, &widgets.WeaponShieldHelp2)
	count += widgets.DrawItemFinder(w, 540, 0)
	u.drawPair(w, 200, "Helmet ID:", &u.helmetID, "Armor ID:", &u.armorID, &widgets.HelmetArmorHelp1, &widgets.HelmetArmorHelp2)
	u.drawPair(w, 400, "Relic 1 ID:", &u.relic1ID, "Relic 2 ID:", &u.relic2ID, &widgets.RelicHelp1, &widgets.RelicHelp2)
	u.countLast = count
}

func (u *equipmentUI) Update() {

}

func (u *equipmentUI) drawPair(w *nucular.Window, y int, label1 string, tb1 *nucular.TextEditor, label2 string, tb2 *nucular.TextEditor, helpTB1 *nucular.TextEditor, helpTB2 *nucular.TextEditor) {
	w.LayoutSpacePush(rect.Rect{
		X: 0,
		Y: y,
		W: 80,
		H: 22,
	})
	w.Label(label1, "LC")
	w.LayoutSpacePush(rect.Rect{
		X: 90,
		Y: y,
		W: 70,
		H: 22,
	})
	_ = widgets.DrawAndValidateHexInput(w, tb1)

	w.LayoutSpacePush(rect.Rect{
		X: 0,
		Y: y + 24,
		W: 80,
		H: 22,
	})
	w.Label(label2, "LC")
	w.LayoutSpacePush(rect.Rect{
		X: 90,
		Y: y + 24,
		W: 70,
		H: 22,
	})
	_ = widgets.DrawAndValidateHexInput(w, tb2)

	w.LayoutSpacePush(rect.Rect{
		X: 170,
		Y: y,
		W: 170,
		H: 190,
	})
	helpTB1.Edit(w)

	w.LayoutSpacePush(rect.Rect{
		X: 360,
		Y: y,
		W: 170,
		H: 190,
	})
	helpTB2.Edit(w)
}

func (u *equipmentUI) initEquipmentTB(tb *nucular.TextEditor) {
	tb.Flags = nucular.EditField
	tb.Maxlen = 2
	tb.SingleLine = true
}
