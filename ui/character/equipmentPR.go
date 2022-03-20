package character

import (
	"ffvi_editor/models/consts/pr"
	"ffvi_editor/ui/widgets"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/rect"
)

func (u *equipmentUI) drawPR(w *nucular.Window) {
	var (
		count   = 18
		helpers = widgets.GetPrEqHelpers()
	)

	w.Row(610).SpaceBegin(u.countLast)
	count += u.drawPrPair(w, 0,
		"Weapon ID:", &character.Equipment.WeaponID, 93, 94, 183,
		"Shield ID:", &character.Equipment.ShieldID, 93, 201, 215,
		&helpers.WeaponShieldHelp1, &helpers.WeaponShieldHelp2)
	count += widgets.DrawItemFinder(w, 550, 0)
	count += u.drawPrPair(w, 200,
		"Helmet ID:", &character.Equipment.HelmetID, 198, 216, 242,
		"Armor ID:", &character.Equipment.ArmorID, 199, 244, 274,
		&helpers.HelmetArmorHelp1, &helpers.HelmetArmorHelp2)
	count += u.drawPrPair(w, 400,
		"Relic 1 ID:", &character.Equipment.Relic1ID, 200, 275, 299,
		"Relic 2 ID:", &character.Equipment.Relic2ID, 200, 275, 299,
		&helpers.RelicHelp1, &helpers.RelicHelp2)
	u.countLast = count
}

func (u *equipmentUI) drawPrPair(w *nucular.Window, y int,
	label1 string, v1 *int, default1, min1, max1 int,
	label2 string, v2 *int, default2, min2, max2 int,
	helpTB1 *nucular.TextEditor, helpTB2 *nucular.TextEditor) (count int) {
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
		W: 80,
		H: 24,
	})
	w.PropertyInt("", default1, v1, max1, 1, 0)
	if item, found := pr.ItemsByID[*v1]; found {
		w.LayoutSpacePush(rect.Rect{
			X: 90,
			Y: y + 24,
			W: 80,
			H: 22,
		})
		w.Label(item, "LC")
		count++
	}

	w.LayoutSpacePush(rect.Rect{
		X: 0,
		Y: y + 48,
		W: 80,
		H: 22,
	})
	w.Label(label2, "LC")
	w.LayoutSpacePush(rect.Rect{
		X: 90,
		Y: y + 48,
		W: 80,
		H: 24,
	})
	w.PropertyInt("", default2, v2, max2, 1, 0)
	if item, found := pr.ItemsByID[*v2]; found {
		w.LayoutSpacePush(rect.Rect{
			X: 90,
			Y: y + 76,
			W: 80,
			H: 22,
		})
		w.Label(item, "LC")
		count++
	}

	w.LayoutSpacePush(rect.Rect{
		X: 180,
		Y: y,
		W: 170,
		H: 190,
	})
	helpTB1.Edit(w)

	w.LayoutSpacePush(rect.Rect{
		X: 370,
		Y: y,
		W: 170,
		H: 190,
	})
	helpTB2.Edit(w)
	return
}
