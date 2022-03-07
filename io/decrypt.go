package io

import (
	"bytes"
	"compress/flate"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"os/exec"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const (
	BlockSize = 32
	password  = "TKX73OHHK1qMonoICbpVT0hIDGe7SkW0"
	salt      = "71Ba2p0ULBGaE6oJ7TjCqwsls1jBKmRL"
)

type BaseKey string

const (
	PictureData                    BaseKey = "pictureData"
	UserDatas                      BaseKey = "userData"
	DataStorage                    BaseKey = "dataStorage"
	CompanionEntity                BaseKey = "companionEntity"
	MoveCount                      BaseKey = "moveCount"
	SubtractSteps                  BaseKey = "SubtractSteps"
	TeleportCacheData              BaseKey = "telepoCacheData"
	PlayableCharacterCorpsId       BaseKey = "playableCharacterCorpsId"
	CurrentSelectedPartyId         BaseKey = "currentSelectedPartyId"
	TimerData                      BaseKey = "timerData"
	ViewType                       BaseKey = "viewType"
	OtherPartyDataList             BaseKey = "otherPartyDataList"
	PartyPlayableCharacterCorpsId  BaseKey = "partyPlayableCharacterCorpsId"
	FieldDefenseNpcEntityIDList    BaseKey = "fieldDefenseNpcEntityIDList"
	BeastFieldEncountExchangeFlags BaseKey = "beastFieldEncountExchangeFlags"
	TimeStamp                      BaseKey = "timeStamp"
	ClearFlag                      BaseKey = "clearFlag"
	IsCompleteFlag                 BaseKey = "isCompleteFlag"
	//ConfigData                     BaseKey = "configData"
	//PlayTime                       BaseKey = "playTime"
)

type UserDataKey string

const (
	CorpsList                 UserDataKey = "corpsList"
	CorpsSlots                UserDataKey = "corpsSlots"
	OwnedCharacterList        UserDataKey = "ownedCharacterList"
	ReleasedJobs              UserDataKey = "releasedJobs"
	OwnedGil                  UserDataKey = "owendGil"
	NormalOwnedItemList       UserDataKey = "normalOwnedItemList"
	importantOwnedItemList    UserDataKey = "importantOwendItemList"
	NormalOwnedItemSortIdList UserDataKey = "normalOwnedItemSortIdList"
	CurrentArea               UserDataKey = "currentArea"
	CurrentLocation           UserDataKey = "currentLocation"
	OwnedTransportationList   UserDataKey = "ownedTransportationList"
	OwendCrystalFlags         UserDataKey = "owendCrystalFlags"
	WarehouseItemList         UserDataKey = "warehouseItemList"
	OwnedKeyWaordList         UserDataKey = "ownedKeyWaordList"
	OwnedMagicList            UserDataKey = "ownedMagicList"
	LearnedAbilityList        UserDataKey = "learnedAbilityList"
	EscapeCount               UserDataKey = "escapeCount"
	BattleCount               UserDataKey = "battleCount"
	CorpsSlotIndex            UserDataKey = "corpsSlotIndex"
	OpenChestCount            UserDataKey = "openChestCount"
	OwnedMagicStoneList       UserDataKey = "ownedMagicStoneList"
	Steps                     UserDataKey = "steps"
	SaveCompleteCount         UserDataKey = "saveCompleteCount"
	MonstersKilledCount       UserDataKey = "monstersKilledCount"
	TotalGil                  UserDataKey = "totalGil"
	//PlayTime                  UserDataKey = "playTime"
	//ConfigData                UserDataKey = "configData"
)

type CharacterDataKey string

const (
	ID                           CharacterDataKey = "id"
	CharacterStatusID            CharacterDataKey = "characterStatusId"
	IsEnableCorps                CharacterDataKey = "isEnableCorps"
	JobID                        CharacterDataKey = "jobId"
	Name                         CharacterDataKey = "name"
	CurrentExp                   CharacterDataKey = "currentExp"
	Parameter                    CharacterDataKey = "parameter"
	CommandList                  CharacterDataKey = "commandList"
	AbilityList                  CharacterDataKey = "abilityList"
	AbilitySlotDataList          CharacterDataKey = "abilitySlotDataList"
	JobList                      CharacterDataKey = "jobList"
	EquipmentList                CharacterDataKey = "equipmentList"
	AdditionOrderOwnedAbilityIds CharacterDataKey = "additionOrderOwnedAbilityIds"
	SortOrderOwnedAbilityIds     CharacterDataKey = "sortOrderOwnedAbilityIds"
	AbilityDictionary            CharacterDataKey = "abilityDictionary"
	SkillLevelTargets            CharacterDataKey = "skillLevelTargets"
	LearningAbilities            CharacterDataKey = "learningAbilitys"
	EquipmentAbilities           CharacterDataKey = "equipmentAbilitys"
	NumberOfBattles              CharacterDataKey = "numberOfButtles"
	OwnedMonsterId               CharacterDataKey = "ownedMonsterId"
	MagicStoneId                 CharacterDataKey = "magicStoneId"
	MagicLearningValue           CharacterDataKey = "magicLearningValue"
)

const targetKey = "target"

var (
	Base       = map[string]interface{}{}
	UserData   = map[string]interface{}{}
	Characters []map[string]interface{}
)

func ParsePixelRemasteredSave(fileName string) (err error) {
	var (
		cmd = exec.Command("python", "./io/python/reader.py", "deobfuscateFile", fileName)
		out []byte

		base       = map[string]interface{}{}
		userData   = map[string]interface{}{}
		characters = make([]map[string]interface{}, 40)
		slTarget   = map[string][]interface{}{}
	)
	if out, err = cmd.Output(); err != nil {
		return
	}

	ol := len(out)
	ol = ol
	s := string(out[2 : len(out)-3])
	s = strings.ReplaceAll(s, `\\r\\n`, "")
	s = strings.ReplaceAll(s, `\\"`, `\"`)
	//s = strings.ReplaceAll(s, "\"{", "{")
	//s = strings.ReplaceAll(s, "}\"", "}")
	if err = json.Unmarshal([]byte(s), &base); err != nil {
		return
	}
	// {"Foo":"{\"Bar\":\"{\\\"Cul\\\":\\\"acul\\\"}\"}"}

	i, _ := base[string(UserDatas)]
	s = strings.ReplaceAll(i.(string), `\\"`, `\"`)
	if err = json.Unmarshal([]byte(s), &userData); err != nil {
		return
	}

	i, _ = userData[string(OwnedCharacterList)]
	s = strings.ReplaceAll(i.(string), `\\"`, `\"`)
	if err = json.Unmarshal([]byte(s), &slTarget); err != nil {
		return
	}

	i, _ = slTarget[targetKey]
	sl := i.([]interface{})
	for j, c := range sl {
		s = strings.ReplaceAll(c.(string), `\\"`, `\"`)
		if err = json.Unmarshal([]byte(s), &characters[j]); err != nil {
			return
		}
	}

	/*
		parameter: {
			"currentHP":5094,
			"currentMP":704,
			"currentMpCountList":"{\\"keys\\":[1,2,3,4,5,6,7,8],\\"values\\":[0,0,0,0,0,0,0,0]}",
			"addtionalMaxMpCountList":"{\\"keys\\":[1,2,3,4,5,6,7,8],\\"values\\":[0,0,0,0,0,0,0,0]}",
			"addtionalLevel":60,
			"addtionalMaxHp":5050,
			"addtionalMaxMp":689,
			"addtionalPower":0,
			"addtionalVitality":0,
			"addtionalAgility":0,
			"addionalWeight":0,
			"addtionalIntelligence":0,
			"addtionalSpirit":0,
			"addtionalAttack":0,
			"addtionalDefense":0,
			"addtionalAbilityDefense":0,
			"addtionalAbilityEvasionRate":0,
			"addtionalMagic":0,
			"addtionalLuck":0,
			"addtionalAccuracyRate":0,
			"addtionalEvasionRate":0,
			"addtionalAbilityDisturbedRate":0,
			"addtionalCriticalRate":0,
			"addtionalDamageDirmeter":0,
			"addtionalAbilityDefenseRate":0,
			"addtionalAccuracyCount":0,
			"addtionalEvasionCount":0,
			"addtionalDefenseCount":0,
			"addtionalMagicDefenseCount":0,
			"currentConditionList":"{\\"target\\":[]}"}
	*/

	/*encrypted, err := base64.StdEncoding.DecodeString(string(f))
	if err != nil {
		return err
	}

	for len(encrypted)%32 != 0 {
		encrypted = append(encrypted, 0)
	}*/

	//b, err := mcrypt.Decrypt([]byte(password), encrypted[:32], []byte(encrypted[32:]))
	//print(string(result))
	Base = base
	UserData = userData
	Characters = characters

	return nil
}

func decompress(deflated []byte) ([]byte, error) {
	var b bytes.Buffer
	r := flate.NewReader(bytes.NewReader(deflated))
	_, err := b.ReadFrom(r)
	if err != nil {
		return nil, err
	}
	r.Close()
	return b.Bytes(), nil
}

func deriveKey() []byte {
	return pbkdf2.Key([]byte(password), []byte(salt), 10, 32, sha1.New)
}

func addBase64Padding(value string) string {
	m := len(value) % 4
	if m != 0 {
		value += strings.Repeat("=", 4-m)
	}

	return value
}

func removeBase64Padding(value string) string {
	return strings.Replace(value, "=", "", -1)
}

func pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(src, padtext...)
}

func unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - unpadding)], nil
}

func encrypt(text string) (string, error) {
	key := deriveKey()
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	msg := pad([]byte(text))
	ciphertext := make([]byte, aes.BlockSize+len(msg))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(msg))
	finalMsg := removeBase64Padding(base64.URLEncoding.EncodeToString(ciphertext))

	return finalMsg, nil
}

func decode(text string) ([]byte, error) {
	decodedMsg, err := base64.StdEncoding.DecodeString(addBase64Padding(text))
	if err != nil {
		return nil, err
	}

	if (len(decodedMsg) % aes.BlockSize) != 0 {
		return nil, errors.New("blocksize must be multipe of decoded message length")
	}
	return decodedMsg, nil
}

func decrypt(text string) ([]byte, error) {
	keyiv := deriveKey()
	key := keyiv[:32]
	iv := keyiv[:16]
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decodedMsg, err := decode(text)
	if err != nil {
		return nil, err
	}

	msg := decodedMsg
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(msg, msg)

	unpadMsg, err := unpad(msg)
	if err != nil {
		return nil, err
	}

	return unpadMsg, nil
}
