package consts

type Command NameValue

var (
	Commands = []Command{
		Command(NewNameValue("Fight", 0x0)),
		Command(NewNameValue("Item", 0x1)),
		Command(NewNameValue("Magic", 0x2)),
		Command(NewNameValue("Morph", 0x3)),
		Command(NewNameValue("Revert", 0x4)),
		Command(NewNameValue("Steal", 0x5)),
		Command(NewNameValue("Capture", 0x6)),
		Command(NewNameValue("SwdTech", 0x7)),
		Command(NewNameValue("Throw", 0x8)),
		Command(NewNameValue("Tools", 0x9)),
		Command(NewNameValue("Blitz", 0x0A)),
		Command(NewNameValue("Runic", 0x0B)),
		Command(NewNameValue("Lore", 0x0C)),
		Command(NewNameValue("Sketch", 0x0D)),
		Command(NewNameValue("Control", 0x0E)),
		Command(NewNameValue("Slot", 0x0F)),
		Command(NewNameValue("Rage", 0x10)),
		Command(NewNameValue("Leap", 0x11)),
		Command(NewNameValue("Mimic", 0x12)),
		Command(NewNameValue("Dance", 0x13)),
		Command(NewNameValue("Row", 0x14)),
		Command(NewNameValue("Def", 0x15)),
		Command(NewNameValue("Jump", 0x16)),
		Command(NewNameValue("XMagic", 0x17)),
		Command(NewNameValue("GPRain", 0x18)),
		Command(NewNameValue("Summon", 0x19)),
		Command(NewNameValue("Health", 0x1A)),
		Command(NewNameValue("Shock", 0x1B)),
		Command(NewNameValue("Possess", 0x1C)),
		Command(NewNameValue("Magitek", 0x1D)),
		Command(NewNameValue("Blank", 0x1E)),
		Command(NewNameValue("Unassigned", 0xFF)),
	}
	commandLookupByValue map[uint8]*Command
)

func init() {
	commandLookupByValue = make(map[uint8]*Command)
	for _, c := range Commands {
		commandLookupByValue[c.Value] = &c
	}
}
