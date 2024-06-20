package savers

import (
	"encoding/json"
	"sort"

	jo "gitlab.com/c0b/go-ordered-json"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/save/util"
)

func Inventory(data *save.Data, key string, sortKey string, inv *core.Inventory) (err error) {
	var (
		slTarget = jo.NewOrderedMap()
		sl       []any
	)
	if sl, err = invNormal(inv); err != nil {
		return
	}

	if inv.ResetSortOrder {
		_ = invResetSort(data, sortKey)
	}

	slTarget.Set(util.TargetKey, sl)
	if err = util.MarshalTo(data.UserData, key, slTarget); err != nil {
		return
	}

	if inv.ResetSortOrder && sortKey != "" {
		slTarget = jo.NewOrderedMap()
		slTarget.Set(util.TargetKey, make([]interface{}, 0))
		if err = util.MarshalTo(data.UserData, sortKey, slTarget); err != nil {
			return
		}
	}
	return
}

func invNormal(inv *core.Inventory) (sl []any, err error) {
	var (
		rows = inv.RowsForPrSave()
		// rows             = debugSortInv(inv.RowsForPrSave())
		b                []byte
		found            = make(map[int]bool)
		removeDuplicates = inv.RemoveDuplicates
	)
	sl = make([]any, 0, len(rows))
	for _, r := range rows {
		if removeDuplicates {
			if _, f := found[r.ItemID]; f {
				continue
			}
			found[r.ItemID] = true
		}
		// Skip Empty rows
		if r.ItemID == 0 || r.Count == 0 {
			continue
		}
		if b, err = json.Marshal(r); err != nil {
			return
		}
		sl = append(sl, string(b))
	}
	return
}

func invResetSort(data *save.Data, key string) (err error) {
	slTarget := jo.NewOrderedMap()
	slTarget.Set(util.TargetKey, make([]interface{}, 0))
	err = util.MarshalTo(data.UserData, key, slTarget)
	return

}

func invDebugSortInv(rows []core.Row) []core.Row {
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].ItemID < rows[j].ItemID
	})
	for i := 0; i < len(rows); i++ {
		c := rows[i].ItemID
		for c >= 100 {
			c -= 100
		}
		rows[i].Count = c
	}
	return rows
}
