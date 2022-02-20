package offsets

import (
	"ffvi_editor/io/save"
	"ffvi_editor/models"
	"ffvi_editor/models/consts"
)

func (o *Offsets) Save() {
	o.saveCharacters()
	o.saveSkills()
	o.saveEspers()
	o.saveMiscStats()
	o.saveInventory()
}

func (o *Offsets) saveCharacters() {
	for ci, c := range models.Characters {
		var (
			i    = ci*nextCharacterOffset + o.CharacterOffset + 2
			name = []rune(c.Name)
		)

		for j := 0; j < 6; j++ {
			if j < len(name) {
				save.SetRuneAt(i+j, name[j])
			} else {
				save.SetRuneAt(i+j, 0)
			}
		}
		i += 6

		save.SetIntAt(i, c.Level)
		i++

		save.SetMultiple(i, i+1, c.HP.Current)
		i += 2

		save.SetMultiple(i, i+1, c.HP.Max)
		i += 2

		save.SetMultiple(i, i+1, c.MP.Current)
		i += 2

		save.SetMultiple(i, i+1, c.MP.Max)
		i += 2

		save.SetMultiple(i, i+3, c.Exp)
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

func (o *Offsets) saveSkills() {

}

func (o *Offsets) saveEspers() {

}

func (o *Offsets) saveMiscStats() {

}

func (o *Offsets) saveInventory() {

}
