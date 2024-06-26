package save

type (
	UserData struct {
		CorpsListInternal                 string             `json:"corpsList"`
		CorpsSlotsInternal                string             `json:"corpsSlots"`
		OwnedCharacterListInternal        string             `json:"ownedCharacterList"`
		ReleasedJobsInternal              string             `json:"releasedJobs"`
		OwnedGil                          int                `json:"owendGil"`
		PlayTime                          float64            `json:"playTime"`
		NormalOwnedItemListInternal       string             `json:"normalOwnedItemList"`
		ImportantOwnedItemListInternal    string             `json:"importantOwendItemList"`
		NormalOwnedItemSortIDListInternal string             `json:"normalOwnedItemSortIdList"`
		CurrentArea                       string             `json:"currentArea"`
		CurrentLocation                   string             `json:"currentLocation"`
		OwnedTransportationListInternal   string             `json:"ownedTransportationList"`
		OwnedCrystalFlagsInternal         string             `json:"owendCrystalFlags"`
		ConfigData                        string             `json:"configData"`
		WarehouseItemListInternal         string             `json:"warehouseItemList"`
		OwnedKeyWordListInternal          string             `json:"ownedKeyWaordList"`
		OwnedMagicListInternal            string             `json:"ownedMagicList"`
		LearnedAbilityListInternal        string             `json:"learnedAbilityList"`
		EscapeCount                       int                `json:"escapeCount"`
		BattleCount                       int                `json:"battleCount"`
		CorpsSlotIndex                    int                `json:"corpsSlotIndex"`
		OpenChestCount                    int                `json:"openChestCount"`
		OwnedMagicStoneListInternal       string             `json:"ownedMagicStoneList"`
		Steps                             int                `json:"steps"`
		SaveCompleteCount                 int                `json:"saveCompleteCount"`
		MonstersKilledCount               int                `json:"monstersKilledCount"`
		TotalGil                          *int               `json:"totalGil,omitempty"`
		CheatSettingsDataInternal         *CheatSettingsData `json:"cheatSettingsData,omitempty"`
		IsOpenedGameBoosterWindow         *bool              `json:"isOpenedGameBoosterWindow,omitempty"`
		WonderWandIndex                   *int               `json:"wonderWandIndex,omitempty"`
	}
	ReleasedJobs struct {
		Target []int `json:"target"`
	}
	NormalOwnedItemList struct {
		Target []string `json:"target"`
	}
	ImportantOwnedItemList struct {
		Target []string `json:"target"`
	}
	OwnedItems struct {
		ContentID int `json:"contentId"`
		Count     int `json:"count"`
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
	CheatSettingsData struct {
		IsEnableEncounters           bool    `json:"isEnableEncount"`
		ExpRate                      float64 `json:"expRate"`
		GilRate                      float64 `json:"gilRate"`
		LearningProficiencyRate      float64 `json:"learningProficiencyRate"`
		IsHpIncreaseByFightCount     bool    `json:"isHpIncreaseByFightCount"`
		AbpRate                      float64 `json:"abpRate"`
		MagicProficiencyRate         float64 `json:"magicProficiencyRate"`
		ParameterUpRate              float64 `json:"parameterUpRate"`
		MagicLearningProficiencyRate float64 `json:"magicLearningProficiencyRate"`
	}
)

func (d *UserData) CorpsList() (v *CorpsList, err error) {
	return UnmarshalOne[CorpsList](d.CorpsListInternal)
}

func (d *UserData) SetCorpsList(v *CorpsList) (err error) {
	d.CorpsListInternal, err = MarshalOne[CorpsList](v)
	return
}

func (d *UserData) CorpsSlots() (v *CorpsSlots, err error) {
	return UnmarshalOne[CorpsSlots](d.CorpsSlotsInternal)
}

func (d *UserData) SetCorpsSlots(v *CorpsSlots) (err error) {
	d.CorpsSlotsInternal, err = MarshalOne[CorpsSlots](v)
	return
}

func (d *UserData) OwnedCharacterList() (v *OwnedCharacterList, err error) {
	return UnmarshalOne[OwnedCharacterList](d.OwnedCharacterListInternal)
}

func (d *UserData) SetOwnedCharacterList(v *OwnedCharacterList) (err error) {
	d.OwnedCharacterListInternal, err = MarshalOne[OwnedCharacterList](v)
	return
}

func (d *UserData) ReleasedJobs() (v *ReleasedJobs, err error) {
	return UnmarshalOne[ReleasedJobs](d.ReleasedJobsInternal)
}

func (d *UserData) SetReleasedJobs(v *ReleasedJobs) (err error) {
	d.ReleasedJobsInternal, err = MarshalOne[ReleasedJobs](v)
	return
}

func (d *UserData) NormalOwnedItems() (v []*OwnedItems, err error) {
	var i *NormalOwnedItemList
	if i, err = UnmarshalOne[NormalOwnedItemList](d.NormalOwnedItemListInternal); err == nil {
		v, err = UnmarshalMany[OwnedItems](i.Target)
	}
	return
}

func (d *UserData) SetNormalOwnedItems(v []*OwnedItems) (err error) {
	i := &NormalOwnedItemList{}
	if i.Target, err = MarshalMany[OwnedItems](v); err == nil {
		d.NormalOwnedItemListInternal, err = MarshalOne(i)
	}
	return
}

func (d *UserData) ImportantOwnedItems() (v []*OwnedItems, err error) {
	var i *ImportantOwnedItemList
	if i, err = UnmarshalOne[ImportantOwnedItemList](d.ImportantOwnedItemListInternal); err == nil {
		v, err = UnmarshalMany[OwnedItems](i.Target)
	}
	return
}

func (d *UserData) SetImportantOwnedItems(v []*OwnedItems) (err error) {
	i := &ImportantOwnedItemList{}
	if i.Target, err = MarshalMany[OwnedItems](v); err == nil {
		d.ImportantOwnedItemListInternal, err = MarshalOne(i)
	}
	return
}

func (d *UserData) NormalOwnedItemSortIDList() (v []int, err error) {
	var l *NormalOwnedItemSortIDList
	if l, err = UnmarshalOne[NormalOwnedItemSortIDList](d.NormalOwnedItemSortIDListInternal); err == nil {
		v = l.Target
	}
	return
}

func (d *UserData) SetNormalOwnedItemSortIDList(v []int) (err error) {
	d.NormalOwnedItemSortIDListInternal, err = MarshalOne[NormalOwnedItemSortIDList](&NormalOwnedItemSortIDList{Target: v})
	return
}

func (d *UserData) OwnedTransportationList() (v *OwnedTransportationList, err error) {
	return UnmarshalOne[OwnedTransportationList](d.OwnedTransportationListInternal)
}

func (d *UserData) SetOwnedTransportationList(v *OwnedTransportationList) (err error) {
	d.OwnedTransportationListInternal, err = MarshalOne[OwnedTransportationList](v)
	return
}

func (d *OwnedTransportationList) OwnedTransportation() (v []*OwnedTransportation, err error) {
	return UnmarshalMany[OwnedTransportation](d.Target)
}

func (d *OwnedTransportationList) SetOwnedTransportation(v []*OwnedTransportation) (err error) {
	d.Target, err = MarshalMany[OwnedTransportation](v)
	return
}

func (d *OwnedTransportation) Position() (v *Position, err error) {
	return UnmarshalOne[Position](d.PositionInternal)
}

func (d *OwnedTransportation) SetPosition(v *Position) (err error) {
	d.PositionInternal, err = MarshalOne[Position](v)
	return
}

func (d *UserData) OwnedCrystalFlags() (v *OwnedCrystalFlags, err error) {
	return UnmarshalOne[OwnedCrystalFlags](d.OwnedCrystalFlagsInternal)
}

func (d *UserData) SetOwnedCrystalFlags(v *OwnedCrystalFlags) (err error) {
	d.OwnedCrystalFlagsInternal, err = MarshalOne[OwnedCrystalFlags](v)
	return
}

func (d *UserData) WarehouseItemList() (v *WarehouseItemList, err error) {
	return UnmarshalOne[WarehouseItemList](d.WarehouseItemListInternal)
}

func (d *UserData) SetWarehouseItemList(v *WarehouseItemList) (err error) {
	d.WarehouseItemListInternal, err = MarshalOne[WarehouseItemList](v)
	return
}

func (d *UserData) OwnedKeyWordList() (v *OwnedKeyWordList, err error) {
	return UnmarshalOne[OwnedKeyWordList](d.OwnedKeyWordListInternal)
}

func (d *UserData) SetOwnedKeyWordList(v *OwnedKeyWordList) (err error) {
	d.OwnedKeyWordListInternal, err = MarshalOne[OwnedKeyWordList](v)
	return
}

func (d *UserData) OwnedMagicList() (v *OwnedMagicList, err error) {
	return UnmarshalOne[OwnedMagicList](d.OwnedMagicListInternal)
}

func (d *UserData) SetOwnedMagicList(v *OwnedMagicList) (err error) {
	d.OwnedMagicListInternal, err = MarshalOne[OwnedMagicList](v)
	return
}

func (d *UserData) LearnedAbilityList() (v *LearnedAbilityList, err error) {
	return UnmarshalOne[LearnedAbilityList](d.LearnedAbilityListInternal)
}

func (d *UserData) SetLearnedAbilityList(v *LearnedAbilityList) (err error) {
	d.LearnedAbilityListInternal, err = MarshalOne[LearnedAbilityList](v)
	return
}

func (d *UserData) OwnedMagicStoneList() (v *OwnedMagicStoneList, err error) {
	return UnmarshalOne[OwnedMagicStoneList](d.OwnedMagicStoneListInternal)
}

func (d *UserData) SetOwnedMagicStoneList(v *OwnedMagicStoneList) (err error) {
	d.OwnedMagicStoneListInternal, err = MarshalOne[OwnedMagicStoneList](v)
	return
}
