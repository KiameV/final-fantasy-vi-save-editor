package consts

/*type SwdTech uint8

const (
	Dispatch    SwdTech = 0x1
	Retort      SwdTech = 0x2
	Slash       SwdTech = 0x4
	QuadraSlam  SwdTech = 0x8
	Empowerer   SwdTech = 0x10
	Stunner     SwdTech = 0x20
	QuadraSlice SwdTech = 0x40
	Cleave      SwdTech = 0x80
)*/

var SwordTech = NewNameSlotMask8s(
	"Dispatch",
	"Retort",
	"Slash",
	"QuadraSlam",
	"Empowerer",
	"Stunner",
	"QuadraSlice",
	"Cleave",
)
