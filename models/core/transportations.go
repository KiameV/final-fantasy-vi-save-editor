package core

type (
	Transportations struct {
		Forms []*Transportation
	}
	Transportation struct {
		ID             int
		Enabled        bool
		ForcedEnabled  bool
		ForcedDisabled bool
		MapID          int
		Position       V3
		Direction      int
		TimeStampTicks uint64
	}
)

func NewTransportations(size int) *Transportations {
	return &Transportations{
		Forms: make([]*Transportation, size),
	}
}

func (t *Transportations) Add(index int, form *Transportation) {
	t.Forms[index] = form
}
