package offsets

import (
	"ffvi_editor/models"
	"ffvi_editor/models/consts"
	snes2 "ffvi_editor/models/consts/snes"
	"ffvi_editor/models/snes"
	"os"
	"strconv"
	"strings"
)

func (o *Offsets) Load(fileName string) (err error) {
	if data, err = os.ReadFile(fileName); err == nil {
		o.loadCharacters()
		o.loadSkills()
		o.loadEspers()
		o.loadMiscStats()
		o.loadInventory()
	}
	return
}

func (o *Offsets) loadCharacters() {
	for ci, c := range models.Characters {
		var (
			i = ci*nextCharacterOffset + o.CharacterOffset + 2
			//data = save.GetChunk(i, (ci+1)*nextCharacterOffset)
			r = make([]rune, 0, 6)
		)
		//data = data

		for j := 0; j < 6; j++ {
			k := GetRuneAt(i + j)
			if k != 0 && k != 255 {
				r = append(r, k)
			}
		}

		c.Name = string(r)
		i += 6

		c.Level = GetIntAt(i)
		i++

		c.HP.Current = o.reverseAndCombine(i, i+1)
		i += 2

		c.HP.Max = o.reverseAndCombine(i, i+1)
		i += 2

		c.MP.Current = o.reverseAndCombine(i, i+1)
		i += 2

		c.MP.Max = o.reverseAndCombine(i, i+1)
		i += 2

		c.Exp = o.reverseAndCombine(i, i+3)
		i += 3

		// Darkness
		c.StatusEffects[0].SetChecked(GetAt(i))
		// Imp
		c.StatusEffects[1].SetChecked(GetAt(i))
		// Invisible
		c.StatusEffects[2].SetChecked(GetAt(i))
		// Magitek
		c.StatusEffects[3].SetChecked(GetAt(i))
		// Poison
		c.StatusEffects[4].SetChecked(GetAt(i))
		// Stone
		c.StatusEffects[5].SetChecked(GetAt(i))
		// Wounded
		c.StatusEffects[6].SetChecked(GetAt(i))
		// Zombie
		c.StatusEffects[7].SetChecked(GetAt(i))
		i++
		// Float
		c.StatusEffects[8].Checked = GetAt(i) >= 128
		i++

		cmd := GetIntAt(i)
		c.Command1 = consts.CommandLookupByValue[cmd]
		i++
		c.Command2 = consts.CommandLookupByValue[GetIntAt(i)]
		i++
		c.Command3 = consts.CommandLookupByValue[GetIntAt(i)]
		i++
		c.Command4 = consts.CommandLookupByValue[GetIntAt(i)]
		i++

		c.Vigor = GetIntAt(i)
		i++
		c.Speed = GetIntAt(i)
		i++
		c.Stamina = GetIntAt(i)
		i++
		c.Magic = GetIntAt(i)
		i++
		//c.Esper = consts.EspersByValue[save.GetIntAt(i)]
		i++
		c.Equipment.WeaponID = GetIntAt(i)
		i++
		c.Equipment.ShieldID = GetIntAt(i)
		i++
		c.Equipment.HelmetID = GetIntAt(i)
		i++
		c.Equipment.ArmorID = GetIntAt(i)
		i++
		c.Equipment.Relic1ID = GetIntAt(i)
		i++
		c.Equipment.Relic2ID = GetIntAt(i)
		i++

		i = nextCharacterMagicOffset*ci + o.KnownMagicOffset
		for _, s := range c.SpellsByIndex {
			s.Value = GetIntAt(i + s.Index)
		}
	}
}

func (o *Offsets) loadSkills() {
	for _, r := range consts.SwordTech {
		r.SetChecked(GetAt(o.SwdTechOffset + r.Slot))
	}
	for _, r := range consts.Lores {
		r.SetChecked(GetAt(o.LoreOffset + r.Slot))
	}
	for _, r := range consts.Blitzes {
		r.SetChecked(GetAt(o.BlitzOffset + r.Slot))
	}
	for _, r := range consts.Dances {
		r.SetChecked(GetAt(o.DanceOffset + r.Slot))
	}
	for _, r := range snes2.Rages {
		r.SetChecked(GetAt(o.RageOffset + r.Slot))
	}
}

func (o *Offsets) loadEspers() {
	for _, r := range snes2.Espers {
		r.SetChecked(GetAt(o.EsperOffset + r.Slot))
	}
}

func (o *Offsets) loadMiscStats() {
	m := models.GetMisc()
	m.GP = o.reverseAndCombine(o.GoldOffset, o.GoldOffset+2)
	m.Steps = o.reverseAndCombine(o.StepsOffset, o.StepsOffset+2)
	m.NumberOfSaves = GetIntAt(o.NumberOfSaves)
	m.SaveCountRollOver = GetIntAt(o.NumberOfSaves + 1)
	m.MapXAxis = GetIntAt(o.MapXYOffset)
	m.MapYAxis = GetIntAt(o.MapXYOffset + 1)
	m.AirshipXAxis = GetIntAt(o.AirshipXYOffset)
	m.AirshipYAxis = GetIntAt(o.AirshipXYOffset + 1)
	m.IsAirshipVisible = (GetAt(o.AirshipSettingsOffset) & (1 << 1)) != 0
	m.CursedShieldFightCount = GetIntAt(o.CursedShieldFightOffset)
}

func (o *Offsets) loadInventory() {
	for i, r := range snes.GetInventoryRows() {
		r.ItemID = toHex(GetIntAt(i + o.InventoryItemIdOffset))
		r.Count = GetIntAt(i + o.InventoryItemCountOffset)
	}
}

func (o *Offsets) reverseAndCombine(start, end int) (r int) {
	for i := end; i >= start; i-- {
		r = r << 8
		r += GetIntAt(i)
	}
	return
}

func toHex(i int) string {
	s := strconv.FormatInt(int64(i), 16)
	if len(s) == 1 {
		s = "0" + s
	}
	return strings.ToUpper(s)
}
