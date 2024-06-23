package save

type (
	OwnedCharacterList struct {
		Target []string `json:"target"`
	}
	OwnedCharacter struct {
		ID                                   int    `json:"id"`
		CharacterStatusID                    int    `json:"characterStatusId"`
		IsEnableCorps                        bool   `json:"isEnableCorps"`
		JobID                                int    `json:"jobId"`
		Name                                 string `json:"name"`
		CurrentExp                           int    `json:"currentExp"`
		Parameter                            string `json:"parameter"`
		CommandListInternal                  string `json:"commandList"`
		AbilityListInternal                  string `json:"abilityList"`
		AbilitySlotDataListInternal          string `json:"abilitySlotDataList"`
		JobListInternal                      string `json:"jobList"`
		EquipmentListInternal                string `json:"equipmentList"`
		AdditionOrderOwnedAbilityIdsInternal string `json:"additionOrderOwnedAbilityIds"`
		SortOrderOwnedAbilityIdsInternal     string `json:"sortOrderOwnedAbilityIds"`
		AbilityDictionaryInternal            string `json:"abilityDictionary"`
		SkillLevelTargetsInternal            string `json:"skillLevelTargets"`
		LearningAbilitiesInternal            string `json:"learningAbilitys"`
		EquipmentAbilitiesInternal           string `json:"equipmentAbilitys"`
		NumberOfButtles                      int    `json:"numberOfButtles"`
		OwnedMonsterID                       int    `json:"ownedMonsterId"`
		MagicStoneID                         int    `json:"magicStoneId"`
		MagicLearningValue                   int    `json:"magicLearningValue"`
	}
	CorpsList struct {
		Target []string `json:"target"`
	}
	Corps struct {
		ID          int `json:"id"`
		CharacterID int `json:"characterId"`
	}
	CorpsSlots struct {
		Keys   []int    `json:"keys"`
		Values []string `json:"values"`
	}
	CorpsSlotsValues struct {
		Target []string `json:"target"`
	}
	CorpsSlotInfo struct {
		ID          int `json:"id"`
		CharacterID int `json:"characterId"`
	}
	CharacterParameters struct {
		CurrentHP                        int    `json:"currentHP"`
		CurrentMP                        int    `json:"currentMP"`
		CurrentMpCountListInternal       string `json:"currentMpCountList"`
		AdditionalMaxMpCountListInternal string `json:"addtionalMaxMpCountList"`
		AdditionalLevel                  int    `json:"addtionalLevel"`
		AdditionalMaxHp                  int    `json:"addtionalMaxHp"`
		AdditionalMaxMp                  int    `json:"addtionalMaxMp"`
		AdditionalPower                  int    `json:"addtionalPower"`
		AdditionalVitality               int    `json:"addtionalVitality"`
		AdditionalAgility                int    `json:"addtionalAgility"`
		AdditionalWeight                 int    `json:"addionalWeight"`
		AdditionalIntelligence           int    `json:"addtionalIntelligence"`
		AdditionalSpirit                 int    `json:"addtionalSpirit"`
		AdditionalAttack                 int    `json:"addtionalAttack"`
		AdditionalDefense                int    `json:"addtionalDefense"`
		AdditionalAbilityDefense         int    `json:"addtionalAbilityDefense"`
		AdditionalAbilityEvasionRate     int    `json:"addtionalAbilityEvasionRate"`
		AdditionalMagic                  int    `json:"addtionalMagic"`
		AdditionalLuck                   int    `json:"addtionalLuck"`
		AdditionalAccuracyRate           int    `json:"addtionalAccuracyRate"`
		AdditionalEvasionRate            int    `json:"addtionalEvasionRate"`
		AdditionalAbilityDisturbedRate   int    `json:"addtionalAbilityDisturbedRate"`
		AdditionalCriticalRate           int    `json:"addtionalCriticalRate"`
		AdditionalDamageDirmeter         int    `json:"addtionalDamageDirmeter"`
		AdditionalAbilityDefenseRate     int    `json:"addtionalAbilityDefenseRate"`
		AdditionalAccuracyCount          int    `json:"addtionalAccuracyCount"`
		AdditionalEvasionCount           int    `json:"addtionalEvasionCount"`
		AdditionalDefenseCount           int    `json:"addtionalDefenseCount"`
		AdditionalMagicDefenseCount      int    `json:"addtionalMagicDefenseCount"`
		CurrentConditionList             string `json:"currentConditionList"`
	}
	CommandList struct {
		Target []int `json:"target"`
	}
	CountList struct {
		Keys   []int `json:"keys"`
		Values []int `json:"values"`
	}
	Conditions struct {
		Target []any `json:"target"`
	}
	AbilityList struct {
		Target []string `json:"target"`
	}
	Ability struct {
		AbilityID  int `json:"abilityId"`
		ContentID  int `json:"contentId"`
		SkillLevel int `json:"skillLevel"`
	}
	AbilitySlotDataList struct {
		Target []string `json:"target"`
	}
	AbilitySlotData struct {
		Level            int    `json:"level"`
		SlotInfoInternal string `json:"slotInfo"`
	}
	AbilitySlot struct {
		Keys   []int    `json:"keys"`
		Values []string `json:"values"`
	}
	JobList struct {
		Target []string `json:"target"`
	}
	Job struct {
		ID                 int `json:"id"`
		Level              int `json:"level"`
		CurrentProficiency int `json:"currentProficiency"`
	}
	EquipmentList struct {
		Keys   []int    `json:"keys"`
		Values []string `json:"values"`
	}
	Equipment struct {
		ContentID int `json:"contentId"`
		Count     int `json:"count"`
	}
	AdditionOrderOwnedAbilityIds struct {
		Target []int `json:"target"`
	}
	SortOrderOwnedAbilityIds struct {
		Target []any `json:"target"`
	}
	AbilityDictionary struct {
		Keys   []any `json:"keys"`
		Values []any `json:"values"`
	}
	SkillLevelTargets struct {
		Keys   []int `json:"keys"`
		Values []int `json:"values"`
	}
	LearningAbilities struct {
		Target []any `json:"target"`
	}
	EquipmentAbilities struct {
		Target []int `json:"target"`
	}
)

func (d *CorpsList) Corps() (v []*Corps, err error) {
	return unmarshalMany[Corps](d.Target)
}

func (d *CorpsList) SetCorps(v []*Corps) (err error) {
	d.Target, err = marshalMany[Corps](v)
	return
}

func (d *CorpsSlots) CorpsSlotsValues() (v []*CorpsSlotsValues, err error) {
	return unmarshalMany[CorpsSlotsValues](d.Values)
}

func (d *CorpsSlots) SetCorpsSlotsValues(v []*CorpsSlotsValues) (err error) {
	d.Values, err = marshalMany[CorpsSlotsValues](v)
	return
}

func (d *CorpsSlotsValues) CorpsSlotInfo() (v []*CorpsSlotInfo, err error) {
	return unmarshalMany[CorpsSlotInfo](d.Target)
}

func (d *CorpsSlotsValues) SetCorpsSlotInfo(v []*CorpsSlotInfo) (err error) {
	d.Target, err = marshalMany[CorpsSlotInfo](v)
	return
}

func (d *OwnedCharacterList) OwnedCharacters() (v []*OwnedCharacter, err error) {
	return unmarshalMany[OwnedCharacter](d.Target)
}

func (d *OwnedCharacterList) SetOwnedCharacters(v []*OwnedCharacter) (err error) {
	d.Target, err = marshalMany[OwnedCharacter](v)
	return
}

func (d *OwnedCharacter) CharacterParameters() (v *CharacterParameters, err error) {
	return unmarshalOne[CharacterParameters](d.Parameter)
}

func (d *OwnedCharacter) SetCharacterParameters(v *CharacterParameters) (err error) {
	d.Parameter, err = marshalOne[CharacterParameters](v)
	return
}

func (d *CharacterParameters) AdditionalMaxMpCountList() (v *CountList, err error) {
	return unmarshalOne[CountList](d.AdditionalMaxMpCountListInternal)
}

func (d *CharacterParameters) SetAdditionalMaxMpCountList(v *CountList) (err error) {
	d.AdditionalMaxMpCountListInternal, err = marshalOne[CountList](v)
	return
}

func (d *CharacterParameters) CurrentMpCountList() (v *CountList, err error) {
	return unmarshalOne[CountList](d.CurrentMpCountListInternal)
}

func (d *CharacterParameters) SetCurrentMpCountList(v *CountList) (err error) {
	d.CurrentMpCountListInternal, err = marshalOne[CountList](v)
	return
}

func (d *OwnedCharacter) CommandList() (v *CommandList, err error) {
	return unmarshalOne[CommandList](d.CommandListInternal)
}

func (d *OwnedCharacter) SetCommandList(v *CommandList) (err error) {
	d.CommandListInternal, err = marshalOne[CommandList](v)
	return
}

func (d *OwnedCharacter) AbilityList() (v *AbilityList, err error) {
	return unmarshalOne[AbilityList](d.AbilityListInternal)
}

func (d *OwnedCharacter) SetAbilityList(v *AbilityList) (err error) {
	d.AbilityListInternal, err = marshalOne[AbilityList](v)
	return
}

func (d *AbilityList) Abilities() (v []*Ability, err error) {
	return unmarshalMany[Ability](d.Target)
}

func (d *AbilityList) SetAbilities(v []*Ability) (err error) {
	d.Target, err = marshalMany[Ability](v)
	return
}

func (d *OwnedCharacter) AbilitySlotDataList() (v *AbilitySlotDataList, err error) {
	return unmarshalOne[AbilitySlotDataList](d.AbilitySlotDataListInternal)
}

func (d *OwnedCharacter) SetAbilitySlotDataList(v *AbilitySlotDataList) (err error) {
	d.AbilitySlotDataListInternal, err = marshalOne[AbilitySlotDataList](v)
	return
}

func (d *AbilitySlotDataList) AbilitySlotData() (v []*AbilitySlotData, err error) {
	return unmarshalMany[AbilitySlotData](d.Target)
}

func (d *AbilitySlotDataList) SetAbilitySlotData(v []*AbilitySlotData) (err error) {
	d.Target, err = marshalMany[AbilitySlotData](v)
	return
}

func (d *AbilitySlotData) AbilitySlot() (v *AbilitySlot, err error) {
	return unmarshalOne[AbilitySlot](d.SlotInfoInternal)
}

func (d *AbilitySlotData) SetAbilitySlot(v *AbilitySlot) (err error) {
	d.SlotInfoInternal, err = marshalOne[AbilitySlot](v)
	return
}

func (d *OwnedCharacter) JobList() (v *JobList, err error) {
	return unmarshalOne[JobList](d.JobListInternal)
}

func (d *OwnedCharacter) SetJobList(v *JobList) (err error) {
	d.JobListInternal, err = marshalOne[JobList](v)
	return
}

func (d *JobList) Jobs() (v []*Job, err error) {
	return unmarshalMany[Job](d.Target)
}

func (d *JobList) SetJobs(v []*Job) (err error) {
	d.Target, err = marshalMany[Job](v)
	return
}

func (d *OwnedCharacter) EquipmentList() (v *EquipmentList, err error) {
	return unmarshalOne[EquipmentList](d.EquipmentListInternal)
}

func (d *OwnedCharacter) SetEquipmentList(v *EquipmentList) (err error) {
	d.EquipmentListInternal, err = marshalOne[EquipmentList](v)
	return
}

func (d *EquipmentList) Equipments() (v []*Equipment, err error) {
	return unmarshalMany[Equipment](d.Values)
}

func (d *EquipmentList) SetEquipments(v []*Equipment) (err error) {
	d.Values, err = marshalMany[Equipment](v)
	return
}

func (d *OwnedCharacter) AdditionOrderOwnedAbilityIds() (v *AdditionOrderOwnedAbilityIds, err error) {
	return unmarshalOne[AdditionOrderOwnedAbilityIds](d.AdditionOrderOwnedAbilityIdsInternal)
}

func (d *OwnedCharacter) SetAdditionOrderOwnedAbilityIds(v *AdditionOrderOwnedAbilityIds) (err error) {
	d.AdditionOrderOwnedAbilityIdsInternal, err = marshalOne[AdditionOrderOwnedAbilityIds](v)
	return
}

func (d *OwnedCharacter) SortOrderOwnedAbilityIds() (v *SortOrderOwnedAbilityIds, err error) {
	return unmarshalOne[SortOrderOwnedAbilityIds](d.SortOrderOwnedAbilityIdsInternal)
}

func (d *OwnedCharacter) SetSortOrderOwnedAbilityIds(v *SortOrderOwnedAbilityIds) (err error) {
	d.SortOrderOwnedAbilityIdsInternal, err = marshalOne[SortOrderOwnedAbilityIds](v)
	return
}

func (d *OwnedCharacter) AbilityDictionary() (v *AbilityDictionary, err error) {
	return unmarshalOne[AbilityDictionary](d.AbilityDictionaryInternal)
}

func (d *OwnedCharacter) SetAbilityDictionary(v *AbilityDictionary) (err error) {
	d.AbilityDictionaryInternal, err = marshalOne[AbilityDictionary](v)
	return
}

func (d *OwnedCharacter) SkillLevelTargets() (v *SkillLevelTargets, err error) {
	return unmarshalOne[SkillLevelTargets](d.SkillLevelTargetsInternal)
}

func (d *OwnedCharacter) SetSkillLevelTargets(v *SkillLevelTargets) (err error) {
	d.SkillLevelTargetsInternal, err = marshalOne[SkillLevelTargets](v)
	return
}

func (d *OwnedCharacter) LearningAbilities() (v *LearningAbilities, err error) {
	return unmarshalOne[LearningAbilities](d.LearningAbilitiesInternal)
}

func (d *OwnedCharacter) SetLearningAbilities(v *LearningAbilities) (err error) {
	d.LearningAbilitiesInternal, err = marshalOne[LearningAbilities](v)
	return
}

func (d *OwnedCharacter) EquipmentAbilities() (v *EquipmentAbilities, err error) {
	return unmarshalOne[EquipmentAbilities](d.EquipmentAbilitiesInternal)
}

func (d *OwnedCharacter) SetEquipmentAbilities(v *EquipmentAbilities) (err error) {
	d.EquipmentAbilitiesInternal, err = marshalOne[EquipmentAbilities](v)
	return
}
