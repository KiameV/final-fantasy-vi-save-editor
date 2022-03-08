package io

import (
	"ffvi_editor/global"
	"ffvi_editor/io/offsets"
)

// Create an Offset for the given save file type
func CreateSnes(saveFileType global.SaveFileType) *offsets.Offsets {
	switch saveFileType {
	//case Snes9xSaveState15:
	//	return create(0x121CF)
	//case Snes9xSaveState16:
	//	return create(0x121DB)
	case global.SRMSlot1:
		return offsets.NewOffsetsFromSrmSlot(0)
	case global.SRMSlot2:
		return offsets.NewOffsetsFromSrmSlot(1)
	case global.SRMSlot3:
		return offsets.NewOffsetsFromSrmSlot(2)
	case global.ZnesSaveState:
		return offsets.NewOffsets(0x2213)
	}
	return nil
}
