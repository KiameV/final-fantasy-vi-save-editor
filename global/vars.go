package global

import (
	"io/fs"
	"os"
)

const (
	WindowWidth  = 725
	WindowHeight = 800
)

var (
	PWD      string
	DirFiles []fs.FileInfo
	FileName string
	FileType SaveFileType
	showing  CurrentScreen
	prevShow CurrentScreen
)

type CurrentScreen byte

const (
	Blank CurrentScreen = iota
	LoadSnes
	SaveSnes
	ShowSnes
	LoadPR
	SavePR
	ShowPR
)

func RollbackShowing() CurrentScreen {
	showing = prevShow
	return showing
}

func GetCurrentShowing() CurrentScreen {
	return showing
}

func SetShowing(s CurrentScreen) {
	if showing != prevShow {
		prevShow = showing
	}
	showing = s
}

func IsShowing(s CurrentScreen) bool {
	return s == showing
}

func IsShowingPR() bool {
	return showing == LoadPR || showing == SavePR || showing == ShowPR
}

// SaveFileType Defines the supported File Types
type SaveFileType byte

const (
	// SRMSlot1 Save file slot 1
	SRMSlot1 SaveFileType = iota
	// SRMSlot2 Save file slot 2
	SRMSlot2
	// SRMSlot3 Save file slot 3
	SRMSlot3
	// ZnesSaveState ZNES save state
	ZnesSaveState
	// PixelRemastered Steam Remastered
	PixelRemastered

	/* Snes9xSaveState15 Snes9x v1.5 save state
	//Snes9xSaveState15
	// Snes9xSaveState16 Snes9x v1.6 save state offset 1
	//Snes9xSaveState16*/
)

func init() {
	var err error
	if PWD, err = os.Getwd(); err != nil {
		PWD = "."
	}
}
