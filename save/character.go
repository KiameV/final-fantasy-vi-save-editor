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
		Level            int          `json:"level"`
		SlotInfoInternal string       `json:"slotInfo"`
		SlotInfo         *AbilitySlot `json:"-"`
	}
	AbilitySlot struct {
		Keys           []int      `json:"keys"`
		ValuesInternal []string   `json:"values"`
		Values         []*Ability `json:"-"`
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
		Keys           []int        `json:"keys"`
		ValuesInternal []string     `json:"values"`
		Values         []*Equipment `json:"-"`
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
	return UnmarshalMany[Corps](d.Target)
}

func (d *CorpsList) SetCorps(v []*Corps) (err error) {
	d.Target, err = MarshalMany[Corps](v)
	return
}

func (d *CorpsSlots) CorpsSlotsValues() (v []*CorpsSlotsValues, err error) {
	return UnmarshalMany[CorpsSlotsValues](d.Values)
}

func (d *CorpsSlots) SetCorpsSlotsValues(v []*CorpsSlotsValues) (err error) {
	d.Values, err = MarshalMany[CorpsSlotsValues](v)
	return
}

func (d *CorpsSlotsValues) CorpsSlotInfo() (v []*CorpsSlotInfo, err error) {
	return UnmarshalMany[CorpsSlotInfo](d.Target)
}

func (d *CorpsSlotsValues) SetCorpsSlotInfo(v []*CorpsSlotInfo) (err error) {
	d.Target, err = MarshalMany[CorpsSlotInfo](v)
	return
}

func (d *OwnedCharacterList) OwnedCharacters() (v []*OwnedCharacter, err error) {
	return UnmarshalMany[OwnedCharacter](d.Target)
}

func (d *OwnedCharacterList) SetOwnedCharacters(v []*OwnedCharacter) (err error) {
	d.Target, err = MarshalMany[OwnedCharacter](v)
	return
}

func (d *OwnedCharacter) CharacterParameters() (v *CharacterParameters, err error) {
	return UnmarshalOne[CharacterParameters](d.Parameter)
}

func (d *OwnedCharacter) SetCharacterParameters(v *CharacterParameters) (err error) {
	d.Parameter, err = MarshalOne[CharacterParameters](v)
	return
}

func (d *CharacterParameters) AdditionalMaxMpCountList() (v *CountList, err error) {
	return UnmarshalOne[CountList](d.AdditionalMaxMpCountListInternal)
}

func (d *CharacterParameters) SetAdditionalMaxMpCountList(v *CountList) (err error) {
	d.AdditionalMaxMpCountListInternal, err = MarshalOne[CountList](v)
	return
}

func (d *CharacterParameters) CurrentMpCountList() (v *CountList, err error) {
	return UnmarshalOne[CountList](d.CurrentMpCountListInternal)
}

func (d *CharacterParameters) SetCurrentMpCountList(v *CountList) (err error) {
	d.CurrentMpCountListInternal, err = MarshalOne[CountList](v)
	return
}

func (d *OwnedCharacter) CommandList() (v *CommandList, err error) {
	return UnmarshalOne[CommandList](d.CommandListInternal)
}

func (d *OwnedCharacter) SetCommandList(v []int) (err error) {
	d.CommandListInternal, err = MarshalOne[CommandList](&CommandList{Target: v})
	return
}

func (d *OwnedCharacter) Abilities() (v []*Ability, err error) {
	var al *AbilityList
	if al, err = UnmarshalOne[AbilityList](d.AbilityListInternal); err == nil {
		v, err = UnmarshalMany[Ability](al.Target)
	}
	return
}

func (d *OwnedCharacter) SetAbilityList(v []*Ability) (err error) {
	var s []string
	if s, err = MarshalMany[Ability](v); err == nil {
		d.AbilityListInternal, err = MarshalOne[AbilityList](&AbilityList{Target: s})
	}
	return
}

func (d *OwnedCharacter) AbilitySlotData() (v []*AbilitySlotData, err error) {
	var (
		asd *AbilitySlotDataList
		i   *Ability
	)
	if asd, err = UnmarshalOne[AbilitySlotDataList](d.AbilitySlotDataListInternal); err == nil {
		if v, err = UnmarshalMany[AbilitySlotData](asd.Target); err == nil {
			for _, s := range v {
				if s.SlotInfo, err = UnmarshalOne[AbilitySlot](s.SlotInfoInternal); err == nil {
					for _, a := range s.SlotInfo.ValuesInternal {
						if a == "" {
							i = &Ability{}
						} else if i, err = UnmarshalOne[Ability](a); err != nil {
							return
						}
						s.SlotInfo.Values = append(s.SlotInfo.Values, i)
					}
				}
			}
		}
	}
	return
}

func (d *OwnedCharacter) SetAbilitySlotData(v []*AbilitySlotData) (err error) {
	for _, asd := range v {
		for i, a := range asd.SlotInfo.Values {
			if a == nil || a.AbilityID == 0 {
				asd.SlotInfo.ValuesInternal[i] = ""
			} else if asd.SlotInfo.ValuesInternal[i], err = MarshalOne(a); err != nil {
				return
			}
		}
		if asd.SlotInfoInternal, err = MarshalOne(asd.SlotInfo); err != nil {
			return
		}
	}
	l := &AbilitySlotDataList{}
	l.Target, err = MarshalMany(v)
	d.AbilitySlotDataListInternal, err = MarshalOne[AbilitySlotDataList](l)
	return
}

func (d *AbilitySlotDataList) SetAbilitySlotData(v []*AbilitySlotData) (err error) {
	d.Target, err = MarshalMany[AbilitySlotData](v)
	return
}

func (d *AbilitySlotData) SetAbilitySlot(v *AbilitySlot) (err error) {
	d.SlotInfoInternal, err = MarshalOne[AbilitySlot](v)
	return
}

func (d *OwnedCharacter) Jobs() (v []*Job, err error) {
	var l *JobList
	if l, err = UnmarshalOne[JobList](d.JobListInternal); err == nil {
		v, err = UnmarshalMany[Job](l.Target)
	}
	return
}

func (d *OwnedCharacter) SetJobs(v []*Job) (err error) {
	l := &JobList{}
	if l.Target, err = MarshalMany[Job](v); err == nil {
		d.JobListInternal, err = MarshalOne[JobList](l)
	}
	return
}

func (d *OwnedCharacter) Equipment() (v *EquipmentList, err error) {
	if v, err = UnmarshalOne[EquipmentList](d.EquipmentListInternal); err == nil {
		v.Values, err = UnmarshalMany[Equipment](v.ValuesInternal)
	}
	return
}

func (d *OwnedCharacter) SetEquipment(v *EquipmentList) (err error) {
	if v.ValuesInternal, err = MarshalMany(v.Values); err == nil {
		d.EquipmentListInternal, err = MarshalOne(v)
	}
	return
}

func (d *OwnedCharacter) AdditionOrderOwnedAbilityIds() (v []int, err error) {
	var i *AdditionOrderOwnedAbilityIds
	if i, err = UnmarshalOne[AdditionOrderOwnedAbilityIds](d.AdditionOrderOwnedAbilityIdsInternal); err == nil {
		v = i.Target
	}
	return
}

func (d *OwnedCharacter) SetAdditionOrderOwnedAbilityIds(v []int) (err error) {
	d.AdditionOrderOwnedAbilityIdsInternal, err = MarshalOne[AdditionOrderOwnedAbilityIds](&AdditionOrderOwnedAbilityIds{Target: v})
	return
}

func (d *OwnedCharacter) SortOrderOwnedAbilityIds() (v *SortOrderOwnedAbilityIds, err error) {
	return UnmarshalOne[SortOrderOwnedAbilityIds](d.SortOrderOwnedAbilityIdsInternal)
}

func (d *OwnedCharacter) SetSortOrderOwnedAbilityIds(v *SortOrderOwnedAbilityIds) (err error) {
	d.SortOrderOwnedAbilityIdsInternal, err = MarshalOne[SortOrderOwnedAbilityIds](v)
	return
}

func (d *OwnedCharacter) AbilityDictionary() (v *AbilityDictionary, err error) {
	return UnmarshalOne[AbilityDictionary](d.AbilityDictionaryInternal)
}

func (d *OwnedCharacter) SetAbilityDictionary(v *AbilityDictionary) (err error) {
	d.AbilityDictionaryInternal, err = MarshalOne[AbilityDictionary](v)
	return
}

func (d *OwnedCharacter) SkillLevelTargets() (v *SkillLevelTargets, err error) {
	return UnmarshalOne[SkillLevelTargets](d.SkillLevelTargetsInternal)
}

func (d *OwnedCharacter) SetSkillLevelTargets(v *SkillLevelTargets) (err error) {
	d.SkillLevelTargetsInternal, err = MarshalOne[SkillLevelTargets](v)
	return
}

func (d *OwnedCharacter) LearningAbilities() (v *LearningAbilities, err error) {
	return UnmarshalOne[LearningAbilities](d.LearningAbilitiesInternal)
}

func (d *OwnedCharacter) SetLearningAbilities(v *LearningAbilities) (err error) {
	d.LearningAbilitiesInternal, err = MarshalOne[LearningAbilities](v)
	return
}

func (d *OwnedCharacter) EquipmentAbilities() (v *EquipmentAbilities, err error) {
	return UnmarshalOne[EquipmentAbilities](d.EquipmentAbilitiesInternal)
}

func (d *OwnedCharacter) SetEquipmentAbilities(v *EquipmentAbilities) (err error) {
	d.EquipmentAbilitiesInternal, err = MarshalOne[EquipmentAbilities](v)
	return
}
