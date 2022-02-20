package offsets

import (
	"ffvi_editor/io/save"
	"ffvi_editor/models"
	"ffvi_editor/models/consts"
	"strconv"
	"strings"
)

func (o *Offsets) Load() {
	o.loadCharacters()
	o.loadSkills()
	o.loadEspers()
	o.loadMiscStats()
	o.loadInventory()
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
			k := save.GetRuneAt(i + j)
			if k != 0 && k != 255 {
				r = append(r, k)
			}
		}

		c.Name = string(r)
		i += 6

		c.Level = save.GetIntAt(i)
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
		c.StatusEffects[0].SetChecked(save.GetAt(i))
		// Imp
		c.StatusEffects[1].SetChecked(save.GetAt(i))
		// Invisible
		c.StatusEffects[2].SetChecked(save.GetAt(i))
		// Magitek
		c.StatusEffects[3].SetChecked(save.GetAt(i))
		// Poison
		c.StatusEffects[4].SetChecked(save.GetAt(i))
		// Stone
		c.StatusEffects[5].SetChecked(save.GetAt(i))
		// Wounded
		c.StatusEffects[6].SetChecked(save.GetAt(i))
		// Zombie
		c.StatusEffects[7].SetChecked(save.GetAt(i))
		i++
		// Float
		c.StatusEffects[8].Checked = save.GetAt(i) >= 128
		i++

		cmd := save.GetIntAt(i)
		c.Command1 = consts.CommandLookupByValue[cmd]
		i++
		c.Command2 = consts.CommandLookupByValue[save.GetIntAt(i)]
		i++
		c.Command3 = consts.CommandLookupByValue[save.GetIntAt(i)]
		i++
		c.Command4 = consts.CommandLookupByValue[save.GetIntAt(i)]
		i++

		c.Vigor = save.GetIntAt(i)
		i++
		c.Speed = save.GetIntAt(i)
		i++
		c.Stamina = save.GetIntAt(i)
		i++
		c.Magic = save.GetIntAt(i)
		i++
		//c.Esper = consts.EspersByValue[save.GetIntAt(i)]
		i++
		c.Equipment.WeaponID = save.GetIntAt(i)
		i++
		c.Equipment.ShieldID = save.GetIntAt(i)
		i++
		c.Equipment.HelmetID = save.GetIntAt(i)
		i++
		c.Equipment.ArmorID = save.GetIntAt(i)
		i++
		c.Equipment.Relic1ID = save.GetIntAt(i)
		i++
		c.Equipment.Relic2ID = save.GetIntAt(i)
		i++

		i = nextCharacterMagicOffset*ci + o.KnownMagicOffset
		for _, s := range c.SpellsByIndex {
			s.Value = save.GetIntAt(i + s.Index)
		}
	}
}

func (o *Offsets) loadSkills() {
	for _, r := range consts.SwordTech {
		r.SetChecked(save.GetAt(o.SwdTechOffset + r.Slot))
	}
	for _, r := range consts.Lores {
		r.SetChecked(save.GetAt(o.LoreOffset + r.Slot))
	}
	for _, r := range consts.Blitzes {
		r.SetChecked(save.GetAt(o.BlitzOffset + r.Slot))
	}
	for _, r := range consts.Dances {
		r.SetChecked(save.GetAt(o.DanceOffset + r.Slot))
	}
	for _, r := range consts.Rages {
		r.SetChecked(save.GetAt(o.RageOffset + r.Slot))
	}
}

func (o *Offsets) loadEspers() {
	for _, r := range consts.Espers {
		r.SetChecked(save.GetAt(o.EsperOffset + r.Slot))
	}
}

func (o *Offsets) loadMiscStats() {
	m := models.GetMisc()

	m.GP = o.reverseAndCombine(o.GoldOffset, o.GoldOffset+2)
	m.Steps = o.reverseAndCombine(o.StepsOffset, o.StepsOffset+2)
	m.NumberOfSaves = save.GetIntAt(o.NumberOfSaves)
	m.SaveCountRollOver = save.GetIntAt(o.NumberOfSaves + 1)
	m.MapXAxis = save.GetIntAt(o.MapXYOffset)
	m.MapYAxis = save.GetIntAt(o.MapXYOffset + 1)
	m.AirshipXAxis = save.GetIntAt(o.AirshipXYOffset)
	m.AirshipYAxis = save.GetIntAt(o.AirshipXYOffset + 1)
	m.IsAirshipVisible = (save.GetAt(o.AirshipSettingsOffset) & (1 << 1)) != 0
	m.CursedShieldFightCount = save.GetIntAt(o.CursedShieldFightOffset)
}

func (o *Offsets) loadInventory() {
	for i, r := range models.GetInventoryRows() {
		r.ItemID = toHex(save.GetIntAt(i + o.InventoryItemIdOffset))
		r.Count = save.GetIntAt(i + o.InventoryItemCountOffset)
	}
}

func (o *Offsets) reverseAndCombine(start, end int) (r int) {
	for i := end; i >= start; i-- {
		r = r << 8
		r += save.GetIntAt(i)
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
