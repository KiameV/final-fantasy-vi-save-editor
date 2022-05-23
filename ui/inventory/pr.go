package inventory

import (
	"encoding/json"
	"ffvi_editor/io"
	"ffvi_editor/models/consts/pr"
	pri "ffvi_editor/models/pr"
	"ffvi_editor/ui"
	"ffvi_editor/ui/widgets"
	"github.com/aarzilli/nucular"
)

func (u *inventoryUI) drawPR(w *nucular.Window) {
	inv := pri.GetInventory()

	// Top
	w.Row(24).Static(150, 10, 150, 10, 100, 10, 100)
	w.CheckboxText("Reset Sort Order", &inv.ResetSortOrder)
	w.Spacing(1)
	w.CheckboxText("Remove Duplicates", &inv.RemoveDuplicates)
	w.Spacing(1)
	if w.ButtonText("Save") {
		if result, err := json.Marshal(inv.Rows); err != nil {
			ui.DrawError = err
		} else if err = io.SaveInvFile(w, result); err != nil {
			ui.DrawError = err
		}
	}
	w.Spacing(1)
	if w.ButtonText("Load") {
		if data, err := io.OpenInvFileDialog(w); err != nil {
			ui.DrawError = err
		} else if err = json.Unmarshal(data, &inv.Rows); err != nil {
			ui.DrawError = err
		}
		inv.ResetSortOrder = true
	}

	w.Row(500).Static(250, 450)

	// Items
	if sw := w.GroupBegin("Inv", 0); sw != nil {
		sw.Row(24).Static(100, 10, 100)
		sw.Label("Item ID", "CC")
		sw.Spacing(1)
		sw.Label("Count", "CC")

		isFirstEmptyRow := true
		for _, r := range inv.GetRows() {
			if r.ItemID == 93 || r.ItemID == 198 || r.ItemID == 199 || r.ItemID == 200 {
				continue
			}

			if r.ItemID == 0 && r.Count == 0 {
				if isFirstEmptyRow {
					isFirstEmptyRow = false
				} else {
					continue
				}
			}

			sw.Row(24).Static(100, 10, 100)
			sw.PropertyInt("", 0, &r.ItemID, 999, 1, 0)
			sw.Spacing(1)
			sw.PropertyInt("", 0, &r.Count, 999, 1, 0)

			if r.ItemID != 0 {
				if name, found := pr.ItemsByID[r.ItemID]; found {
					sw.Row(24).Static(200)
					sw.Label(name, "LC")
					sw.Row(6).Static()
				}
			}
		}
		sw.GroupEnd()
	}

	// Helpers
	if sw := w.GroupBegin("Inv Helpers", 0); sw != nil {
		helper := widgets.GetPrInvHelpers()
		sw.Row(190).Static(200, 200)
		helper.MiscItemsHelp.Edit(sw)
		widgets.AddItemFinder(sw, pr.ItemsByName)

		sw.Row(190).Static(200, 200)
		helper.WeaponShieldHelp1.Edit(sw)
		helper.WeaponShieldHelp2.Edit(sw)

		sw.Row(190).Static(200, 200)
		helper.HelmetArmorHelp1.Edit(sw)
		helper.HelmetArmorHelp2.Edit(sw)

		sw.Row(190).Static(200, 200)
		helper.RelicHelp1.Edit(sw)
		helper.RelicHelp2.Edit(sw)

		sw.GroupEnd()
	}
}
