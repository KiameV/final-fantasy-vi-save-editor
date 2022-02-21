package save

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
	// SteamRemastered Steam Remastered
	SteamRemastered

	/* Snes9xSaveState15 Snes9x v1.5 save state
	//Snes9xSaveState15
	// Snes9xSaveState16 Snes9x v1.6 save state offset 1
	//Snes9xSaveState16*/
)

var data []byte

func GetChunk(from, to int) []byte {
	return data[from:to]
}

func Get() []byte {
	return data
}

func GetAt(i int) byte {
	return data[i]
}

func GetIntAt(i int) int {
	return int(data[i])
}

func GetRuneAt(i int) rune {
	c := data[i]
	if c >= 154 && c <= 178 {
		// Lower case characters
		return rune(data[i] - 154 + 'a')
	}
	if c >= 180 && c <= 189 {
		// numbers
		return rune(data[i] - 180 + '0')
	}

	switch c {
	case 190:
		return '!'
	case 191:
		return '?'
	case 192:
		return '~'
	case 193:
		return ':'
	case 194:
		return '"'
	case 195:
		return '\''
	case 196:
		return '-'
	case 197:
		return '.'
	case 0, 255:
		return 0
	}
	return rune(data[i] - 63)
}

func Set(b []byte) {
	data = b
}

func SetAt(i int, b byte) {
	data[i] = b
}

func SetIntAt(i int, v int) {
	data[i] = byte(v)
}

func SetRuneAt(i int, r rune) {
	if r >= 'a' && r <= 'z' {
		// Lower case characters
		data[i] = byte((r - 'a') + 154)
	} else if r >= '0' && r <= '9' {
		// numbers
		data[i] = byte((r - '0') + 180)
	} else {
		switch r {
		case '!':
			data[i] = 190
		case '?':
			data[i] = 191
		case '~':
			data[i] = 192
		case ':':
			data[i] = 193
		case '"':
			data[i] = 194
		case '\'':
			data[i] = 195
		case '-':
			data[i] = 196
		case '.':
			data[i] = 197
		case 0, 255:
			data[i] = 255
		default:
			data[i] = byte(r + 63)
		}
	}
}
