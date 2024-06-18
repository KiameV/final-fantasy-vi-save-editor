package loaders

import (
	"encoding/json"

	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/save"
	"pixel-remastered-save-editor/save/util"
)

func Inventory(data *save.Data, key string) (inv *core.Inventory, err error) {
	var (
		sl  interface{}
		row core.Row
	)
	inv = core.NewInventory()
	if sl, err = util.FromTarget(data.UserData, key); err != nil {
		return
	}
	for i, r := range sl.([]interface{}) {
		if err = json.Unmarshal([]byte(r.(string)), &row); err != nil {
			return
		}
		inv.Set(i, row)
	}
	return
}
