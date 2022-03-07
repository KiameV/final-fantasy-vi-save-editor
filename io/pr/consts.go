package pr

// base keys
const (
	PictureData                    = "pictureData"
	UserDatas                      = "userData"
	DataStorage                    = "dataStorage"
	CompanionEntity                = "companionEntity"
	MoveCount                      = "moveCount"
	SubtractSteps                  = "SubtractSteps"
	TeleportCacheData              = "telepoCacheData"
	PlayableCharacterCorpsId       = "playableCharacterCorpsId"
	CurrentSelectedPartyId         = "currentSelectedPartyId"
	TimerData                      = "timerData"
	ViewType                       = "viewType"
	OtherPartyDataList             = "otherPartyDataList"
	PartyPlayableCharacterCorpsId  = "partyPlayableCharacterCorpsId"
	FieldDefenseNpcEntityIDList    = "fieldDefenseNpcEntityIDList"
	BeastFieldEncountExchangeFlags = "beastFieldEncountExchangeFlags"
	TimeStamp                      = "timeStamp"
	ClearFlag                      = "clearFlag"
	IsCompleteFlag                 = "isCompleteFlag"
	//ConfigData                   = "configData"
	//PlayTime                     = "playTime"
)

// user data keys
const (
	CorpsList                 = "corpsList"
	CorpsSlots                = "corpsSlots"
	OwnedCharacterList        = "ownedCharacterList"
	ReleasedJobs              = "releasedJobs"
	OwnedGil                  = "owendGil"
	NormalOwnedItemList       = "normalOwnedItemList"
	importantOwnedItemList    = "importantOwendItemList"
	NormalOwnedItemSortIdList = "normalOwnedItemSortIdList"
	CurrentArea               = "currentArea"
	CurrentLocation           = "currentLocation"
	OwnedTransportationList   = "ownedTransportationList"
	OwendCrystalFlags         = "owendCrystalFlags"
	WarehouseItemList         = "warehouseItemList"
	OwnedKeyWaordList         = "ownedKeyWaordList"
	OwnedMagicList            = "ownedMagicList"
	LearnedAbilityList        = "learnedAbilityList"
	EscapeCount               = "escapeCount"
	BattleCount               = "battleCount"
	CorpsSlotIndex            = "corpsSlotIndex"
	OpenChestCount            = "openChestCount"
	OwnedMagicStoneList       = "ownedMagicStoneList"
	Steps                     = "steps"
	SaveCompleteCount         = "saveCompleteCount"
	MonstersKilledCount       = "monstersKilledCount"
	TotalGil                  = "totalGil"
	//PlayTime                = "playTime"
	//ConfigData              = "configData"
)

// character data keys
const (
	ID                           = "id"
	CharacterStatusID            = "characterStatusId"
	IsEnableCorps                = "isEnableCorps"
	JobID                        = "jobId"
	Name                         = "name"
	CurrentExp                   = "currentExp"
	Parameter                    = "parameter"
	CommandList                  = "commandList"
	AbilityList                  = "abilityList"
	AbilitySlotDataList          = "abilitySlotDataList"
	JobList                      = "jobList"
	EquipmentList                = "equipmentList"
	AdditionOrderOwnedAbilityIds = "additionOrderOwnedAbilityIds"
	SortOrderOwnedAbilityIds     = "sortOrderOwnedAbilityIds"
	AbilityDictionary            = "abilityDictionary"
	SkillLevelTargets            = "skillLevelTargets"
	LearningAbilities            = "learningAbilitys"
	EquipmentAbilities           = "equipmentAbilitys"
	NumberOfBattles              = "numberOfButtles"
	OwnedMonsterId               = "ownedMonsterId"
	MagicStoneId                 = "magicStoneId"
	MagicLearningValue           = "magicLearningValue"
)

type CharacterKey byte

const (
	Terra  CharacterKey = 0
	Wedge  CharacterKey = 1
	Biggs  CharacterKey = 2
	Locke  CharacterKey = 4
	Mog    CharacterKey = 15
	Edgar  CharacterKey = 16
	Sabin  CharacterKey = 17
	Shadow CharacterKey = 18
	Celes  CharacterKey = 21
	Cyan   CharacterKey = 22
	Gau    CharacterKey = 24
	Setzer CharacterKey = 25
	Strago CharacterKey = 27
	Relm   CharacterKey = 28
	Umaro  CharacterKey = 31
	Gogo   CharacterKey = 32
)

const targetKey = "target"

// character param fields
const (
	CurrentHP                     = "currentHP"
	CurrentMP                     = "currentMP"
	AdditionalMaxMpCountList      = "addtionalMaxMpCountList"
	AdditionalLevel               = "addtionalLevel"
	AdditionalMaxHp               = "addtionalMaxHp"
	AdditionalMaxMp               = "addtionalMaxMp"
	AdditionalPower               = "addtionalPower"
	AdditionalVitality            = "addtionalVitality"
	AdditionalAgility             = "addtionalAgility"
	AdditionalWeight              = "addionalWeight"
	AdditionalIntelligence        = "addtionalIntelligence"
	AdditionalSpirit              = "addtionalSpirit"
	AdditionAttack                = "addtionalAttack"
	AdditionalDefence             = "addtionalDefense"
	AdditionAbilityDefense        = "addtionalAbilityDefense"
	AdditionAbilityEvasionRate    = "addtionalAbilityEvasionRate"
	AdditionMagic                 = "addtionalMagic"
	AdditionLuck                  = "addtionalLuck"
	AdditionalAccuracyRate        = "addtionalAccuracyRate"
	AdditionalEvasionRate         = "addtionalEvasionRate"
	AdditionAbilityDistrurbedRate = "addtionalAbilityDisturbedRate"
	AdditionalCriticalRate        = "addtionalCriticalRate"
	AdditionalDamageDirmeter      = "addtionalDamageDirmeter"
	AdditionalAbilityDefenseRate  = "addtionalAbilityDefenseRate"
	AdditionalAccuracyCount       = "addtionalAccuracyCount"
	AdditionalEvasionCount        = "addtionalEvasionCount"
	AdditionalDefenceCount        = "addtionalDefenseCount"
	AdditionalMagicDef            = "addtionalMagicDefenseCount"
	CurrentConditionList          = "currentConditionList"
)
