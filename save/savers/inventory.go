package savers

import (
	"encoding/json"

	jo "gitlab.com/c0b/go-ordered-json"
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/save/util"
)

func Inventory(data *save.Data, key string, sortKey string, inv *core.Inventory) (err error) {
	var (
		rows             = inv.RowsForPrSave()
		sl               = make([]interface{}, 0, len(rows))
		b                []byte
		slTarget         = jo.NewOrderedMap()
		found            = make(map[int]bool)
		removeDuplicates = inv.RemoveDuplicates
	)

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
