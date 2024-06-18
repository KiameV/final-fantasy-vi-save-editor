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
	Game byte
)

const (
	Unspecified Game = iota
	One
	Two
	Three
	Four
	Five
	Six
)

func (g Game) IsOne() bool   { return g == One }
func (g Game) IsTwo() bool   { return g == Two }
func (g Game) IsThree() bool { return g == Three }
func (g Game) IsFour() bool  { return g == Four }
func (g Game) IsFive() bool  { return g == Five }
func (g Game) IsSix() bool   { return g == Six }
func (g Game) IsAny(game ...Game) bool {
	for _, v := range game {
		if g == v {
			return true
		}
	}
	return false
}

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
