package consts

/*type StatusEffect uint16

const (
	None      StatusEffect = 0x0
	Darkness  StatusEffect = 0x1
	Zombie    StatusEffect = 0x2
	Poison    StatusEffect = 0x4
	Magiteck  StatusEffect = 0x8
	Invisible StatusEffect = 0x10
	Imp       StatusEffect = 0x20
	Stone     StatusEffect = 0x40
	Wounded   StatusEffect = 0x80
	Float     StatusEffect = 0x100 // This is in another byte and will not 'flag' with the others, this is here for better code flow to the UI
)*/

var StatusEffects = []string{
	"Darkness",
	"Zombie",
	"Poison",
	"Magiteck",
	"Invisible",
	"Imp",
	"Stone",
	"Wounded",
	"Float",
}

func NewStatusEffects() (result []*NameSlotMask8) {
	s := NewNameSlotMask8s(StatusEffects...)
	result = make([]*NameSlotMask8, len(StatusEffects))
	for i, se := range s {
		result[i] = se
	}
	return
}
