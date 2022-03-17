package consts

type NameValue struct {
	Name  string
	Value int
}

type NameValueChecked struct {
	NameValue
	Checked bool
}

func NewNameValue(name string, value int) NameValue {
	return NameValue{Name: name, Value: value}
}

func NewValueName(value int, name string) *NameValue {
	return &NameValue{Name: name, Value: value}
}

func NewNameValueChecked(name string, value int) *NameValueChecked {
	return &NameValueChecked{NameValue: NameValue{Name: name, Value: value}}
}

func NewValueNameChecked(value int, name string) *NameValueChecked {
	return &NameValueChecked{NameValue: NameValue{Name: name, Value: value}}
}

func NewNameValues(names ...string) []NameValue {
	result := make([]NameValue, len(names))
	for i, n := range names {
		result[i] = NewNameValue(n, i)
	}
	return result
}

type NameSlotMask8 struct {
	Name    string
	Slot    int
	Mask    uint8
	Checked bool
}

func NewNameSlotMask8(name string, slot int, mask uint8) NameSlotMask8 {
	return NameSlotMask8{Name: name, Slot: slot, Mask: mask}
}

func NewNameSlotMask8s(names ...string) []*NameSlotMask8 {
	var (
		result       = make([]*NameSlotMask8, len(names))
		slot   int   = 0
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

func (n *NameSlotMask8) SetChecked(b byte) {
	n.Checked = (n.Mask & b) != 0
}

func GenerateBytes(v []*NameSlotMask8) []byte {
	result := make([]byte, v[len(v)-1].Slot+1)
	for _, n := range v {
		if n.Checked {
			result[n.Slot] |= n.Mask
		}
	}
	return result
}
