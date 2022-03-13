package offsets

import (
	"ffvi_editor/models"
	"ffvi_editor/models/consts"
	snes2 "ffvi_editor/models/consts/snes"
	"ffvi_editor/models/snes"
	"strconv"
	"unsafe"
)

func (o *Offsets) Save() error {
	o.saveCharacters()
	o.saveSkills()
	o.saveEspers()
	o.saveMiscStats()
	o.saveInventory()
	return nil
}

func (o *Offsets) saveCharacters() {
	for ci, c := range models.Characters {
		var (
			i    = ci*nextCharacterOffset + o.CharacterOffset + 2
			name = []rune(c.Name)
		)

		for j := 0; j < 6; j++ {
			if j < len(name) {
				SetRuneAt(i+j, name[j])
			} else {
				SetRuneAt(i+j, 0)
			}
		}
		i += 6

		SetIntAt(i, c.Level)
		i++

		o.setMultiple(c.HP.Current, i, i+1)
		i += 2

		o.setMultiple(c.HP.Max, i, i+1)
		i += 2

		o.setMultiple(c.MP.Current, i, i+1)
		i += 2

		o.setMultiple(c.MP.Max, i, i+1)
		i += 2

		o.setMultiple(c.Exp, i, i+1, i+2)
		i += 3

		i = o.setFromNameSlotMasks(i, c.StatusEffects)
		// Float
		if c.StatusEffects[8].Checked {
			SetAt(i-1, 0xFF)
		} else {
			SetAt(i-1, 0)
		}

		SetIntAt(i, c.Command1.Value)
		i++
		SetIntAt(i, c.Command2.Value)
		i++
		SetIntAt(i, c.Command3.Value)
		i++
		SetIntAt(i, c.Command4.Value)
		i++

		SetIntAt(i, c.Vigor)
		i++
		SetIntAt(i, c.Speed)
		i++
		SetIntAt(i, c.Stamina)
		i++
		SetIntAt(i, c.Magic)
		i++
		// Skip Esper
		i++
		SetIntAt(i, c.Equipment.WeaponID)
		i++
		SetIntAt(i, c.Equipment.ShieldID)
		i++
		SetIntAt(i, c.Equipment.HelmetID)
		i++
		SetIntAt(i, c.Equipment.ArmorID)
		i++
		SetIntAt(i, c.Equipment.Relic1ID)
		i++
		SetIntAt(i, c.Equipment.Relic2ID)
		i++

		i = nextCharacterMagicOffset*ci + o.KnownMagicOffset
		for _, s := range c.SpellsByIndex {
			SetIntAt(i+s.Index, s.Value)
		}
	}
}

func (o *Offsets) saveSkills() {
	o.setFromNameSlotMasks(o.SwdTechOffset, consts.SwordTech)
	o.setFromNameSlotMasks(o.LoreOffset, consts.Lores)
	o.setFromNameSlotMasks(o.BlitzOffset, consts.Blitzes)
	o.setFromNameSlotMasks(o.DanceOffset, consts.Dances)
}

func (o *Offsets) saveEspers() {
	o.setFromNameSlotMasks(o.EsperOffset, snes2.Espers)
}

func (o *Offsets) saveMiscStats() {
	m := models.GetMisc()
	o.setMultiple(m.GP, o.GoldOffset, o.GoldOffset+1, o.GoldOffset+2)
	o.setMultiple(m.Steps, o.StepsOffset, o.StepsOffset+1, o.StepsOffset+2)
	SetIntAt(o.NumberOfSaves, m.NumberOfSaves)
	SetIntAt(o.NumberOfSaves+1, m.SaveCountRollOver)
	SetIntAt(o.MapXYOffset, m.MapXAxis)
	SetIntAt(o.MapXYOffset+1, m.MapYAxis)
	SetIntAt(o.AirshipXYOffset, m.AirshipXAxis)
	SetIntAt(o.AirshipXYOffset+1, m.AirshipYAxis)
	s := GetIntAt(o.AirshipSettingsOffset)
	if m.IsAirshipVisible {
		s |= 1 << 1
	} else {
		mask := ^(1 << 1)
		s &= mask
	}
	SetIntAt(o.AirshipSettingsOffset, s)
	SetIntAt(o.CursedShieldFightOffset, m.CursedShieldFightCount)
}

func (o *Offsets) saveInventory() {
	for i, r := range snes.GetInventoryRows() {
		SetIntAt(i+o.InventoryItemIdOffset, fromHex(r.ItemID))
		SetIntAt(i+o.InventoryItemCountOffset, r.Count)
	}
}

func (o *Offsets) setMultiple(value int, indexes ...int) {
	for i := 0; i < len(indexes) && i < int(unsafe.Sizeof(value)); i++ {
		b := *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&value)) + uintptr(i)))
		SetAt(indexes[i], b)
	}
}

func (o *Offsets) setFromNameSlotMasks(index int, nsms []*consts.NameSlotMask8) int {
	for _, v := range consts.GenerateBytes(nsms) {
		SetAt(index, v)
		index++
	}
	return index
}

func fromHex(hex string) int {
	if value, err := strconv.ParseInt(hex, 16, 32); err != nil {
		return 0xFF
	} else {
		return int(value)
	}
}
