package io

import (
	"ffvi_editor/io/offsets"
)

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

// Create an Offset for the given save file type
func Create(saveFileType SaveFileType) IOer {
	switch saveFileType {
	//case Snes9xSaveState15:
	//	return create(0x121CF)
	//case Snes9xSaveState16:
	//	return create(0x121DB)
	case SRMSlot1:
		return offsets.NewOffsetsFromSrmSlot(0)
	case SRMSlot2:
		return offsets.NewOffsetsFromSrmSlot(1)
	case SRMSlot3:
		return offsets.NewOffsetsFromSrmSlot(2)
	case ZnesSaveState:
		return offsets.NewOffsets(0x2213)
	case PixelRemastered:
		panic("Pixel Remastered not supported")
		//return pr.NewPR()
	}
	return nil
}
