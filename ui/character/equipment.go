package character

import (
	"ffvi_editor/global"
	"ffvi_editor/models"
	"ffvi_editor/models/consts/snes"
	"ffvi_editor/ui/widgets"
	"fmt"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/rect"
	"strconv"
	"strings"
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
	if global.IsShowingPR() {
		u.drawPR(w)
	} else {
		var (
			count   = 18
			helpers = widgets.GetSnesHelpers()
		)

		w.Row(610).SpaceBegin(u.countLast)
		count += u.drawPair(w, 0, "Weapon ID:", &u.weaponID, "Shield ID:", &u.shieldID, &helpers.WeaponShieldHelp1, &helpers.WeaponShieldHelp2)
		count += widgets.DrawItemFinder(w, 550, 0)
		count += u.drawPair(w, 200, "Helmet ID:", &u.helmetID, "Armor ID:", &u.armorID, &helpers.HelmetArmorHelp1, &helpers.HelmetArmorHelp2)
		count += u.drawPair(w, 400, "Relic 1 ID:", &u.relic1ID, "Relic 2 ID:", &u.relic2ID, &helpers.RelicHelp1, &helpers.RelicHelp2)
		u.countLast = count
	}
}

func (u *equipmentUI) Update(character *models.Character) {
	u.weaponID.Text(toHex(character.Equipment.WeaponID))
	u.shieldID.Text(toHex(character.Equipment.ShieldID))
	u.helmetID.Text(toHex(character.Equipment.HelmetID))
	u.armorID.Text(toHex(character.Equipment.ArmorID))
	u.relic1ID.Text(toHex(character.Equipment.Relic1ID))
	u.relic2ID.Text(toHex(character.Equipment.Relic2ID))
}

func toNum(i int) []rune {
	return []rune(fmt.Sprintf("%d", i))
}

func toHex(i int) []rune {
	s := strconv.FormatInt(int64(i), 16)
	s = strings.ToUpper(s)
	return []rune(s)
}

func (u *equipmentUI) drawPair(w *nucular.Window, y int, label1 string, tb1 *nucular.TextEditor, label2 string, tb2 *nucular.TextEditor, helpTB1 *nucular.TextEditor, helpTB2 *nucular.TextEditor) (count int) {
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
		H: 22,
	})
	_ = widgets.DrawAndValidateHexInput(w, tb1)

	if item, found := snes.ItemsByID[string(tb1.Buffer)]; found {
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
		H: 22,
	})
	_ = widgets.DrawAndValidateHexInput(w, tb2)
	if item, found := snes.ItemsByID[string(tb2.Buffer)]; found {
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

func (u *equipmentUI) initEquipmentTB(tb *nucular.TextEditor) {
	tb.Flags = nucular.EditField
	tb.Maxlen = 2
	tb.SingleLine = true
}
