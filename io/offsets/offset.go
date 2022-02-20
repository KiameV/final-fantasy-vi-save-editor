package offsets

type Offsets struct {
	// Offset for the start of the character information
	CharacterOffset int
	// Offset to the known-magic information
	KnownMagicOffset int
	// Offset to the obtained espers
	EsperOffset int
	// Offset to the SwdTech flag
	SwdTechOffset int
	// Offset to the start of the Lore flags
	LoreOffset int
	// Offset to the Blitz flag
	BlitzOffset int
	// Offset to the Dance flag
	DanceOffset int
	// Offset to the start of the Rage flags
	RageOffset int
	// Offset to the start of the Item Id for each item
	InventoryItemIdOffset int
	// Offset to the start of the Item Count for each item
	InventoryItemCountOffset int
	// Offset to the GP definition
	GoldOffset int
	// Offset to the Number of Steps definition
	StepsOffset int
	// Offset to the Map X/Y Axis definition
	MapXYOffset int
	// Offset to the Airship X/Y Axis definition
	AirshipXYOffset int
	// bit 0: WoR: 0=sad song, 1=norm.song
	// bit 1: airship visibility
	// bit 2: 1==coin toss cutscean/crashed
	// bit 3: start w/ Veldt Music
	// bit 4: ?
	// bit 5: ?
	// bit 6: ?
	// bit 7: 1=is a save point
	AirshipSettingsOffset int
	// Offset to the Cursed Shield fight count
	CursedShieldFightOffset int
	// Offset to the Number of Saves definition
	NumberOfSaves int
}
