package importantInventory

import (
	"ffvi_editor/models/consts/pr"
	pri "ffvi_editor/models/pr"
	"ffvi_editor/ui"
	"ffvi_editor/ui/widgets"
	"github.com/aarzilli/nucular"
)

type importantInventoryUI struct {
	ids    []*nucular.TextEditor
	counts []*nucular.TextEditor
}

func NewUI() ui.UI {
	inv := &importantInventoryUI{
		ids:    make([]*nucular.TextEditor, 255),
		counts: make([]*nucular.TextEditor, 255),
	}
	for i := 0; i < 50; i++ {
		tb := &nucular.TextEditor{}
		inv.ids[i] = tb
		tb.Flags = nucular.EditField
		tb.SingleLine = true
		tb.Maxlen = 3

		tb = &nucular.TextEditor{}
		inv.counts[i] = tb
		tb.Flags = nucular.EditField
		tb.SingleLine = true
		tb.Maxlen = 3
	}
	return inv
}

func (u *importantInventoryUI) Draw(w *nucular.Window) {
	inv := pri.GetImportantInventory()

	w.Row(300).Static(250, 450)
	// Items
	if sw := w.GroupBegin("Imp Inv", 0); sw != nil {
		sw.Row(24).Static(100, 10, 100)
		sw.Label("Item ID", "LC")
		sw.Spacing(1)
		sw.Label("Count", "LC")

		isFirstEmptyRow := true
		for _, r := range inv.GetRows() {
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
				if name, found := pr.ImportantItemsByID[r.ItemID]; found {
					sw.Row(24).Static(200)
					sw.Label(name, "LC")
					sw.Row(6).Static()
				}
			}
		}
		sw.GroupEnd()
	}
	// Helpers
	if sw := w.GroupBegin("Imp Inv Helper", 0); sw != nil {
		helper := widgets.GetPrInvHelpers()
		sw.Row(280).Static(200)
		helper.ImportantItemsHelp.Edit(sw)
		sw.GroupEnd()
	}
}

func (u *importantInventoryUI) Refresh() {

}

func (u *importantInventoryUI) Name() string {
	return "Important Inventory"
}

func (u *importantInventoryUI) Behavior() ui.Behavior {
	return ui.PrShow
}
