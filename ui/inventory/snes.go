package inventory

import (
	"ffvi_editor/models/consts/snes"
	snes2 "ffvi_editor/models/snes"
	"ffvi_editor/ui/widgets"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/rect"
)

func (u *inventoryUI) drawSnes(w *nucular.Window) {
	var (
		e     nucular.EditEvents
		tb    *nucular.TextEditor
		y     int
		count = 2
	)

	w.Row(u.yLast).SpaceBegin(u.countLast)
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

	for i, r := range snes2.GetInventoryRows() {
		w.LayoutSpacePush(rect.Rect{
			X: 0,
			Y: y,
			W: 75,
			H: 22,
		})
		tb = u.ids[i]
		if e = widgets.DrawAndValidateHexInput(w, tb); e == nucular.EditActive || e == nucular.EditCommitted {
			if len(tb.Buffer) == 2 {
				r.ItemID = string(tb.Buffer)
			}
		}

		w.LayoutSpacePush(rect.Rect{
			X: 85,
			Y: y,
			W: 75,
			H: 24,
		})
		w.PropertyInt("", 0, &r.Count, 99, 1, 0)
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
	helper := widgets.GetSnesHelpers()
	w.LayoutSpacePush(rect.Rect{
		X: 170,
		Y: 0,
		W: 170,
		H: 190,
	})
	helper.MiscItemsHelp.Edit(w)

	w.LayoutSpacePush(rect.Rect{
		X: 170,
		Y: 200,
		W: 170,
		H: 190,
	})
	helper.WeaponShieldHelp1.Edit(w)

	w.LayoutSpacePush(rect.Rect{
		X: 360,
		Y: 200,
		W: 170,
		H: 190,
	})
	helper.WeaponShieldHelp2.Edit(w)

	w.LayoutSpacePush(rect.Rect{
		X: 170,
		Y: 400,
		W: 170,
		H: 190,
	})
	helper.HelmetArmorHelp1.Edit(w)

	w.LayoutSpacePush(rect.Rect{
		X: 360,
		Y: 400,
		W: 170,
		H: 190,
	})
	helper.HelmetArmorHelp2.Edit(w)

	w.LayoutSpacePush(rect.Rect{
		X: 170,
		Y: 600,
		W: 170,
		H: 190,
	})
	helper.RelicHelp1.Edit(w)

	w.LayoutSpacePush(rect.Rect{
		X: 360,
		Y: 600,
		W: 170,
		H: 190,
	})
	helper.RelicHelp2.Edit(w)
	count += 7

	// Finder
	u.countLast = count + widgets.DrawItemFinder(w, 540, 0)
}
