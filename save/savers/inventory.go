package savers

import (
	"encoding/json"
	"os"
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

	slTarget.Set(util.TargetKey, sl)
	if err = util.MarshalTo(data.UserData, key, slTarget); err != nil {
		return
	}

	if inv.ResetSortOrder && sortKey != "" {
		err = invResetSort(data, sortKey)
	}
	return
}

func invNormal(inv *core.Inventory) (sl []any, err error) {
	var (
		rows             = inv.RowsForPrSave()
		b                []byte
		found            = make(map[int]bool)
		removeDuplicates = inv.RemoveDuplicates
	)
	if os.Getenv("PR_SORT_INV") == "true" {
		rows = debugSortInv(inv.RowsForPrSave())
	}
	sl = make([]any, 0, len(rows))
	for _, r := range rows {
		if removeDuplicates {
			if _, f := found[r.ContentID]; f {
				continue
			}
			found[r.ContentID] = true
		}
		// Skip Empty rows
		if r.ContentID == 0 || r.Count == 0 {
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
	// slTarget.Set(util.TargetKey, make([]interface{}, 0))
	// err = util.MarshalTo(data.UserData, key, slTarget)
	l := len(data.Save.Inventory.Rows)
	order := make([]int, l)
	for i := 0; i < l; i++ {
		order[i] = i + 1
	}
	slTarget.Set(util.TargetKey, order)
	_ = util.MarshalTo(data.UserData, key, slTarget)
	return

}

func debugSortInv(rows []core.Row) []core.Row {
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].ContentID < rows[j].ContentID
	})
	for i := 0; i < len(rows); i++ {
		c := rows[i].ContentID
		for c >= 100 {
			c -= 100
		}
		rows[i].Count = c
	}
	return rows
}
