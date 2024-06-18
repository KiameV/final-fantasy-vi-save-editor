package savers

/*
type partyMember struct {
	ID          int `json:"id"`
	CharacterID int `json:"characterId"`
}

func Party(data *save.Data) (err error) {
	var (
		partyID = getPartyID(data)
		b       []byte
		sl      = make([]interface{}, 4)
	)
	for i, m := range  data.Save.Party.Members {
		pm := partyMember{
			ID:          partyID,
			CharacterID: m.CharacterID,
		}
		if b, err = json.Marshal(&pm); err != nil {
			return
		}
		sl[i] = string(b)
	}
	return util.SetTarget(data.UserData, util.CorpsList, sl)
}

func partyID(data *Data) int {
	i, err := getFromTarget(data.UserData, util.CorpsList)
	if err != nil {
		return 1
	}
	var m partyMember
	for _, c := range i.([]interface{}) {
		if err = json.Unmarshal([]byte(c.(string)), &m); err != nil {
			return 1
		}
	}
	return m.ID
}

func addToNeeded(needed *map[int]int, id int) {
	if count, found := (*needed)[id]; !found {
		(*needed)[id] = 1
	} else {
		(*needed)[id] = count + 1
	}
}
*/
