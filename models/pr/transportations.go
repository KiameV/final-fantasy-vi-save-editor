package pr

var Transportations []*Transportation

type Transportation struct {
	ID             int
	Enabled        bool
	ForcedEnabled  bool
	ForcedDisabled bool
	MapID          int
	Position       V3
	Direction      int
	TimeStampTicks uint64
}
