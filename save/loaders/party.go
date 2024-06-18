package loaders

import (
	"pixel-remastered-save-editor/models/core"
	"pixel-remastered-save-editor/save"
)

func Party(data *save.Data) (p *core.Party, err error) {
	/*count := 4
	if data.Game.IsFive() {
		count = 5
	}
	var (
		party  = core.NewParty(count)
		member core.Member
		i      any
	)
	if i, err = util.FromTarget(data.UserData, util.CorpsList); err != nil {
		return
	}
	for slot, c := range i.([]interface{}) {
		if err = json.Unmarshal([]byte(c.(string)), &member); err != nil {
			return
		}
		party.SetMember(slot, &member)
	}*/
	return
}
