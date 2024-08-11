package global

import (
	"os"
	"time"
)

const (
	WindowWidth  = 725
	WindowHeight = 800
)

type (
	SaveFileType byte
)

const (
	PC SaveFileType = iota
	PS
)

var (
	PWD string
)

func init() {
	var err error
	if PWD, err = os.Getwd(); err != nil {
		PWD = "."
	}
}

func NowToTicks() uint64 {
	return uint64(float64(time.Now().UnixNano())*0.01) + uint64(60*60*24*365*1970*10000000)
}
