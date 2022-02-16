package consts

type NameValue struct {
	Name  string
	Value int
}

func NewNameValue(name string, value int) NameValue {
	return NameValue{Name: name, Value: value}
}

type NameSlotMask8 struct {
	Name    string
	Slot    uint8
	Mask    uint8
	Checked bool
}

func NewNameSlotMask8(name string, slot uint8, mask uint8) NameSlotMask8 {
	return NameSlotMask8{Name: name, Slot: slot, Mask: mask}
}

func NewNameSlotMask8s(names ...string) []*NameSlotMask8 {
	var (
		result       = make([]*NameSlotMask8, len(names))
		slot   uint8 = 0
		mask   uint8 = 0x1
	)

	for i, n := range names {
		result[i] = &NameSlotMask8{
			Name: n,
			Slot: slot,
			Mask: mask,
		}
		switch mask {
		case 0x1:
			mask = 0x2
		case 0x2:
			mask = 0x4
		case 0x4:
			mask = 0x8
		case 0x8:
			mask = 0x10
		case 0x10:
			mask = 0x20
		case 0x20:
			mask = 0x40
		case 0x40:
			mask = 0x80
		default:
			mask = 0x1
			slot++
		}
	}
	return result
}
