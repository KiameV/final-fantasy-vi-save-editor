package save

type (
	UserData struct {
		CorpsListInternal                 string  `json:"corpsList"`
		CorpsSlotsInternal                string  `json:"corpsSlots"`
		OwnedCharacterListInternal        string  `json:"ownedCharacterList"`
		ReleasedJobsInternal              string  `json:"releasedJobs"`
		OwnedGil                          int     `json:"owendGil"`
		PlayTime                          float64 `json:"playTime"`
		NormalOwnedItemListInternal       string  `json:"normalOwnedItemList"`
		ImportantOwnedItemListInternal    string  `json:"importantOwendItemList"`
		NormalOwnedItemSortIDListInternal string  `json:"normalOwnedItemSortIdList"`
		CurrentArea                       string  `json:"currentArea"`
		CurrentLocation                   string  `json:"currentLocation"`
		OwnedTransportationListInternal   string  `json:"ownedTransportationList"`
		OwnedCrystalFlagsInternal         string  `json:"owendCrystalFlags"`
		ConfigData                        string  `json:"configData"`
		WarehouseItemListInternal         string  `json:"warehouseItemList"`
		OwnedKeyWordListInternal          string  `json:"ownedKeyWaordList"`
		OwnedMagicListInternal            string  `json:"ownedMagicList"`
		LearnedAbilityListInternal        string  `json:"learnedAbilityList"`
		EscapeCount                       int     `json:"escapeCount"`
		BattleCount                       int     `json:"battleCount"`
		CorpsSlotIndex                    int     `json:"corpsSlotIndex"`
		OpenChestCount                    int     `json:"openChestCount"`
		OwnedMagicStoneListInternal       string  `json:"ownedMagicStoneList"`
		Steps                             int     `json:"steps"`
		SaveCompleteCount                 int     `json:"saveCompleteCount"`
		MonstersKilledCount               int     `json:"monstersKilledCount"`
	}
	ReleasedJobs struct {
		Target []int `json:"target"`
	}
	NormalOwnedItemList struct {
		Target []string `json:"target"`
	}
	NormalOwnedItems struct {
		ContentID int `json:"contentId"`
		Count     int `json:"count"`
	}
	ImportantOwnedItemList struct {
		Target []any `json:"target"`
	}
	NormalOwnedItemSortIDList struct {
		Target []int `json:"target"`
	}
	OwnedTransportationList struct {
		Target []string `json:"target"`
	}
	OwnedTransportation struct {
		PositionInternal string `json:"position"`
		Direction        int    `json:"direction"`
		ID               int    `json:"id"`
		MapID            int    `json:"mapId"`
		Enable           bool   `json:"enable"`
		TimeStampTicks   int    `json:"timeStampTicks"`
	}
	OwnedCrystalFlags struct {
		Target []bool `json:"target"`
	}
	WarehouseItemList struct {
		Target []any `json:"target"`
	}
	OwnedKeyWordList struct {
		Target []int `json:"target"`
	}
	OwnedMagicList struct {
		Target []any `json:"target"`
	}
	LearnedAbilityList struct {
		Target []int `json:"target"`
	}
	OwnedMagicStoneList struct {
		Target []any `json:"target"`
	}
)

func (d *UserData) CorpsList() (v *CorpsList, err error) {
	return unmarshalOne[CorpsList](d.CorpsListInternal)
}

func (d *UserData) SetCorpsList(v *CorpsList) (err error) {
	d.CorpsListInternal, err = marshalOne[CorpsList](v)
	return
}

func (d *UserData) CorpsSlots() (v *CorpsSlots, err error) {
	return unmarshalOne[CorpsSlots](d.CorpsSlotsInternal)
}

func (d *UserData) SetCorpsSlots(v *CorpsSlots) (err error) {
	d.CorpsSlotsInternal, err = marshalOne[CorpsSlots](v)
	return
}

func (d *UserData) OwnedCharacterList() (v *OwnedCharacterList, err error) {
	return unmarshalOne[OwnedCharacterList](d.OwnedCharacterListInternal)
}

func (d *UserData) SetOwnedCharacterList(v *OwnedCharacterList) (err error) {
	d.OwnedCharacterListInternal, err = marshalOne[OwnedCharacterList](v)
	return
}

func (d *UserData) ReleasedJobs() (v *ReleasedJobs, err error) {
	return unmarshalOne[ReleasedJobs](d.ReleasedJobsInternal)
}

func (d *UserData) SetReleasedJobs(v *ReleasedJobs) (err error) {
	d.ReleasedJobsInternal, err = marshalOne[ReleasedJobs](v)
	return
}

func (d *UserData) NormalOwnedItemList() (v *NormalOwnedItemList, err error) {
	return unmarshalOne[NormalOwnedItemList](d.NormalOwnedItemListInternal)
}

func (d *UserData) SetNormalOwnedItemList(v *NormalOwnedItemList) (err error) {
	d.NormalOwnedItemListInternal, err = marshalOne[NormalOwnedItemList](v)
	return
}

func (d *NormalOwnedItemList) NormalOwnedItems() (v []*NormalOwnedItems, err error) {
	return unmarshalMany[NormalOwnedItems](d.Target)
}

func (d *NormalOwnedItemList) SetNormalOwnedItems(v []*NormalOwnedItems) (err error) {
	d.Target, err = marshalMany[NormalOwnedItems](v)
	return
}

func (d *UserData) ImportantOwnedItemList() (v *ImportantOwnedItemList, err error) {
	return unmarshalOne[ImportantOwnedItemList](d.ImportantOwnedItemListInternal)
}

func (d *UserData) SetImportantOwnedItemList(v *ImportantOwnedItemList) (err error) {
	d.ImportantOwnedItemListInternal, err = marshalOne[ImportantOwnedItemList](v)
	return
}

func (d *UserData) NormalOwnedItemSortIDList() (v *NormalOwnedItemSortIDList, err error) {
	return unmarshalOne[NormalOwnedItemSortIDList](d.NormalOwnedItemSortIDListInternal)
}

func (d *UserData) SetNormalOwnedItemSortIDList(v *NormalOwnedItemSortIDList) (err error) {
	d.NormalOwnedItemSortIDListInternal, err = marshalOne[NormalOwnedItemSortIDList](v)
	return
}

func (d *UserData) OwnedTransportationList() (v *OwnedTransportationList, err error) {
	return unmarshalOne[OwnedTransportationList](d.OwnedTransportationListInternal)
}

func (d *UserData) SetOwnedTransportationList(v *OwnedTransportationList) (err error) {
	d.OwnedTransportationListInternal, err = marshalOne[OwnedTransportationList](v)
	return
}

func (d *OwnedTransportationList) OwnedTransportation() (v []*OwnedTransportation, err error) {
	return unmarshalMany[OwnedTransportation](d.Target)
}

func (d *OwnedTransportationList) SetOwnedTransportation(v []*OwnedTransportation) (err error) {
	d.Target, err = marshalMany[OwnedTransportation](v)
	return
}

func (d *OwnedTransportation) Position() (v *Position, err error) {
	return unmarshalOne[Position](d.PositionInternal)
}

func (d *OwnedTransportation) SetPosition(v *Position) (err error) {
	d.PositionInternal, err = marshalOne[Position](v)
	return
}

func (d *UserData) OwnedCrystalFlags() (v *OwnedCrystalFlags, err error) {
	return unmarshalOne[OwnedCrystalFlags](d.OwnedCrystalFlagsInternal)
}

func (d *UserData) SetOwnedCrystalFlags(v *OwnedCrystalFlags) (err error) {
	d.OwnedCrystalFlagsInternal, err = marshalOne[OwnedCrystalFlags](v)
	return
}

func (d *UserData) WarehouseItemList() (v *WarehouseItemList, err error) {
	return unmarshalOne[WarehouseItemList](d.WarehouseItemListInternal)
}

func (d *UserData) SetWarehouseItemList(v *WarehouseItemList) (err error) {
	d.WarehouseItemListInternal, err = marshalOne[WarehouseItemList](v)
	return
}

func (d *UserData) OwnedKeyWordList() (v *OwnedKeyWordList, err error) {
	return unmarshalOne[OwnedKeyWordList](d.OwnedKeyWordListInternal)
}

func (d *UserData) SetOwnedKeyWordList(v *OwnedKeyWordList) (err error) {
	d.OwnedKeyWordListInternal, err = marshalOne[OwnedKeyWordList](v)
	return
}

func (d *UserData) OwnedMagicList() (v *OwnedMagicList, err error) {
	return unmarshalOne[OwnedMagicList](d.OwnedMagicListInternal)
}

func (d *UserData) SetOwnedMagicList(v *OwnedMagicList) (err error) {
	d.OwnedMagicListInternal, err = marshalOne[OwnedMagicList](v)
	return
}

func (d *UserData) LearnedAbilityList() (v *LearnedAbilityList, err error) {
	return unmarshalOne[LearnedAbilityList](d.LearnedAbilityListInternal)
}

func (d *UserData) SetLearnedAbilityList(v *LearnedAbilityList) (err error) {
	d.LearnedAbilityListInternal, err = marshalOne[LearnedAbilityList](v)
	return
}

func (d *UserData) OwnedMagicStoneList() (v *OwnedMagicStoneList, err error) {
	return unmarshalOne[OwnedMagicStoneList](d.OwnedMagicStoneListInternal)
}

func (d *UserData) SetOwnedMagicStoneList(v *OwnedMagicStoneList) (err error) {
	d.OwnedMagicStoneListInternal, err = marshalOne[OwnedMagicStoneList](v)
	return
}
