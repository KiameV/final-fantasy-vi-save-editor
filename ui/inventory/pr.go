package inventory

import (
	"ffvi_editor/models/consts/snes"
	pri "ffvi_editor/models/pr"
	"ffvi_editor/ui/widgets"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/rect"
)

func (u *inventoryUI) drawPR(w *nucular.Window) {
	var (
		y     int
		count = 3
	)

	w.Row(u.yLast).SpaceBegin(u.countLast)
	w.LayoutSpacePush(rect.Rect{
		X: 0,
		Y: y,
		W: 150,
		H: 22,
	})
	w.CheckboxText("Reset Sort Order", &pri.GetInventory().ResetSortOrder)
	y += 24

	w.LayoutSpacePush(rect.Rect{
		X: 0,
		Y: y,
		W: 75,
		H: 22,
	})
	w.Label("Item ID", "LC")

	w.LayoutSpacePush(rect.Rect{
		X: 85,
		Y: y,
		W: 75,
		H: 22,
	})
	w.Label("Count", "LC")
	y += 24

	for i, r := range pri.GetInventoryRows() {
		w.LayoutSpacePush(rect.Rect{
			X: 0,
			Y: y,
			W: 75,
			H: 24,
		})
		w.PropertyInt("", 0, &r.ItemID, 999, 1, 0)

		w.LayoutSpacePush(rect.Rect{
			X: 85,
			Y: y,
			W: 75,
			H: 24,
		})
		w.PropertyInt("", 0, &r.Count, 999, 1, 0)
		y += 24
		count += 2
		if item, found := snes.ItemsByID[string(u.ids[i].Buffer)]; found && item != "Empty" {
			w.LayoutSpacePush(rect.Rect{
				X: 0,
				Y: y,
				W: 160,
				H: 22,
			})
			w.Label(item, "LC")
			count++
			y += 30
		}
	}
	u.yLast = y

	// Helpers
	w.LayoutSpacePush(rect.Rect{
		X: 170,
		Y: 0,
		W: 170,
		H: 190,
	})
	widgets.MiscItemsHelp.Edit(w)

	w.LayoutSpacePush(rect.Rect{
		X: 170,
		Y: 200,
		W: 170,
		H: 190,
	})
	widgets.WeaponShieldHelp1.Edit(w)

	w.LayoutSpacePush(rect.Rect{
		X: 360,
		Y: 200,
		W: 170,
		H: 190,
	})
	widgets.WeaponShieldHelp2.Edit(w)

	w.LayoutSpacePush(rect.Rect{
		X: 170,
		Y: 400,
		W: 170,
		H: 190,
	})
	widgets.HelmetArmorHelp1.Edit(w)

	w.LayoutSpacePush(rect.Rect{
		X: 360,
		Y: 400,
		W: 170,
		H: 190,
	})
	widgets.HelmetArmorHelp2.Edit(w)

	w.LayoutSpacePush(rect.Rect{
		X: 170,
		Y: 600,
		W: 170,
		H: 190,
	})
	widgets.RelicHelp1.Edit(w)

	w.LayoutSpacePush(rect.Rect{
		X: 360,
		Y: 600,
		W: 170,
		H: 190,
	})
	widgets.RelicHelp2.Edit(w)
	count += 7

	// Finder
	u.countLast = count + widgets.DrawItemFinder(w, 540, 0)
}
