package offsets

const (
	// srmSlotSize Size of the SRM's save slot
	srmSlotSize = 0xA00
	// nextCharacterOffset Number of bytes between each character
	nextCharacterOffset = 37
	// nextCharacterMagicOffset Number of bytes between each characters known-magic set
	nextCharacterMagicOffset = 54

	// Character Information

	// characterOffset Offset for the start of the character information
	characterOffset = 0x0
	// knownMagicOffset Offset to the known-magic information
	knownMagicOffset = 0x46E

	// Obtained Espers

	// esperOffset Offset to the obtained espers
	esperOffset = 0x469

	// Skills

	// swdTechOffset Offset to the SwdTech flag
	swdTechOffset = 0x6F7
	// blitzOffset Offset to the Blitz flag
	blitzOffset = 0x728
	// loreOffset Offset to the start of the Lore flags
	loreOffset = 0x729
	// danceOffset Offset to the Dance flag
	danceOffset = 0x74C
	// rageOffset Offset to the start of the Rage flags
	rageOffset = 0x72C

	// Inventory

	// inventoryItemIDOffset Offset to the start of the Item Id for each item
	inventoryItemIDOffset = 0x269
	// inventoryItemCountOffset Offset to the start of the Item Count for each item
	inventoryItemCountOffset = 0x369

	// Misc Data

	// goldOffset Offset to the GP definition
	goldOffset = 0x260
	// stepsOffset Offset to the Number of Steps definition
	stepsOffset = 0x266
	// numberOfSavesOffset Offset to the Number of Saves definition
	numberOfSavesOffset = 0x7C7
	// mapXYOffset Offset to the Map X/Y Axis definition
	mapXYOffset = 0x960
	// airshipXYOffset Offset to the Airship X/Y Axis definition
	airshipXYOffset = 0x962
	// cursedShieldFightsOffset Offset to the Cursed Shield fight counter
	cursedShieldFightsOffset = 0x7D5

	// TODO - Investigate Airship Settings Further
	// airshipSettingsOffset Offset to Airship settings flag
	//  bit 0: WoR: 0=sad song, 1=norm.song
	//  bit 1: airship visibility
	//  bit 2: 1=coin toss cutscean/crashed
	//  bit 3: start w/ Veldt Music
	//  bits 4, 5, 6: ?
	//  bit 7: 1=is a save point
	airshipSettingsOffset = 0x8B7

	// GameTimeOffset Offset to the Played Time definition
	gameTimeOffset = 0x263
)

// CreateFromOffset an Offset for the given save file type (snes)
//func CreateFromSnesOffset(offset int) *Offsets {
//	return Create(offset - 2)
//}

// NewOffsetsFromSrmSlot an Offset for the given save slot
func NewOffsetsFromSrmSlot(offset int) *Offsets {
	return NewOffsets(offset * srmSlotSize)
}

// NewOffsets an Offset instance with the given base offset for each defined offset
func NewOffsets(baseOffset int) *Offsets {
	return &Offsets{
		CharacterOffset:          characterOffset + baseOffset,
		KnownMagicOffset:         knownMagicOffset + baseOffset,
		EsperOffset:              esperOffset + baseOffset,
		SwdTechOffset:            swdTechOffset + baseOffset,
		LoreOffset:               loreOffset + baseOffset,
		BlitzOffset:              blitzOffset + baseOffset,
		DanceOffset:              danceOffset + baseOffset,
		RageOffset:               rageOffset + baseOffset,
		InventoryItemIdOffset:    inventoryItemIDOffset + baseOffset,
		InventoryItemCountOffset: inventoryItemCountOffset + baseOffset,
		GoldOffset:               goldOffset + baseOffset,
		StepsOffset:              stepsOffset + baseOffset,
		MapXYOffset:              mapXYOffset + baseOffset,
		AirshipXYOffset:          airshipXYOffset + baseOffset,
		AirshipSettingsOffset:    airshipSettingsOffset + baseOffset,
		CursedShieldFightOffset:  cursedShieldFightsOffset + baseOffset,
		NumberOfSaves:            numberOfSavesOffset + baseOffset,
	}
}

/*
looks like active party and positions are handled in the area 2646 and 2469
const unsigned char BTL_MN1 =0;
const unsigned char BTL_MN2 =1;
const unsigned char BTL_MN3 =2;
const unsigned char BTL_MN4 =3;
const unsigned char EQ_SWRD =0;
const unsigned char EQ_SHLD =1;
const unsigned char EQ_HELM =2;
const unsigned char EQ_ARMR =3;
const unsigned char EQ_REL1 =4;
const unsigned char EQ_REL2 =5;
const unsigned char STAT_A  =0;
const unsigned char STAT_B  =1;
const unsigned char VIGOR   =2;
const unsigned char SPEED   =3;
const unsigned char STAMINA =4;
const unsigned char MAGPOW  =5;
const unsigned char ESPER   =6;
const unsigned char PROF1   =7;
const unsigned char PROF2   =8;
typedef unsigned char data;
const long maxdata = 36000lu;
const int  max_slots = 3;
const int  max_magic = 54;          // 54 spells
const int  slot1_offset = 0x0;
const int  slot2_offset = 0xA00;
const int  slot3_offset = 0x1400;
const int  slot_lenght = 0xA00;
// 16 official players + 2 customizable players
const int  nb_players_menu = 18;
const int  nb_players = 16;         // 16 official players
const int  nb_mgc_users = 12;       // 12 magic users
// the first 592 bytes cover character's abilities
const int  car_info_lenght = 37;    // 37 bytes X 16 players
const int  raster_offset  = 592;    // setup of active party(ies)
const int  gold_offset  = 608;      // gold pieces
const int  gtime_offset = 611;      // game time
const int  steps_offset = 614;      // steps
const int  itm_lst_offset = 617;    // possessed items
const int  itm_qty_offset = 873;    // quantity of items
// possessed espers
const int  esp_list_offset = 1129;  // four bytes, switch on/off
const int  magic_offset = 1134;     // learn % of magic X 12
// not decoded 1782, 1784 thru 1831
const int  swrdt_offset = 1783;     // swordtech abil.
const int  blitz_offset = 1832;     // blitz abil.
const int  lores_offset = 1833;     // lores abil.
const int  rages_offset = 1836;     // rages abil.
const int  dance_offset = 1868;     // dance abil.
const int  nb_saves_offset = 1991;  // number of saves, CzarDragon credit
const int  act_memb_offset = 2268;  // active members (shop selection), 2 bytes, CzarDragon credits, lsb byte1 is Terra
const int  act_memb2_offset = 2270; // active members (airship selection), 2 bytes, lsb byte1 is Terra
const int  map_xy_offset = 2400; // XY location on overworld, 2 bytes, CzarDragon credit
const int  ship_xy_offset = 2402;   // XY location of airship, 2 bytes, CzarDragon credit
// not decoded 1869 thru end (2560)
// decoded but not implemented yet:
const int  actgrp_offset = 1133;    // active group, CzarDragon credit
// sword tech name for FF6 CzarDragon credit
const int  swtk_name_offset = 1784; // 48 bytes
const int  btlspeed_offset = 1869;  // mess. and btl speed, 3bits: 6->4 lsb=btl, 2->0 msb=mess
                        // value from 1 to 6 = 0 to 5
                        // bit 7 = cmd.set, 1 = short; bit 3 = bat. mode, 1 = wait
const int  wallpaper_offset = 1870; // forest, chocobo, stone, ect., bit 0 -> 3
                        // gauge off, sound mono, cursor memory, reequip empty, bit 4->7 set to 1
const int  dummy1_offset = 1871;    // ????     when dummy1 to 5 + config1 is set to ff, all key becomes down key
const int  dummy2_offset = 1872;    // ????
const int  dummy3_offset = 1873;    // ????
const int  dummy4_offset = 1874;    // ????
const int  dummy5_offset = 1875;    // ????
const int  config1_offset = 1876;   // about single/multiple cntrl, mag. order, color "arrow" pointer
const int  fontcolor1_offset = 1877; // 00 = blue, red doesnt work
const int  fontcolor2_offset = 1878; //
// this is probably coded on 12 bits, it won't be tested yet...
const int  border1wp1_offset = 1879; // 1st color pntr, 00 = is blue (border part 1, wallpaper no 1)
const int  border2wp1_offset = 1880; // " 00 = is red
const int  border3wp1_offset = 1881; // 2nd color pointer
const int  border4wp1_offset = 1882; // 2nd color pointer
const int  border5wp1_offset = 1883; // 3nd color pointer
const int  border6wp1_offset = 1884; // 3rd color pointer
const int  border7wp1_offset = 1885; // 4rd color pointer
const int  border8wp1_offset = 1886; // 4rd color pointer
const int  border9wp1_offset = 1887; // 5rd color pointer
const int  border10wp1_offset = 1888; // 5rd color pointer
const int  border11wp1_offset = 1889; // 6rd color pointer
const int  border12wp1_offset = 1890; // 6rd color pointer
const int  border13wp1_offset = 1891; // 7th color pointer
const int  border14wp1_offset = 1892; // 7th color pointer
// 1893 thru 1906 wallparer 2
// 1907 thru 1920 wallparer 3
// 1921 thru 1934 wallparer 4
// 1935 thru 1948 wallparer 5
// 1949 thru 1962 wallparer 6
// 1963 thru 1976 wallparer 7
// 1977 thru 1990 wallparer 8
// those "illegal" wp may just be data representing something else in the game...
const int  test_offset = 2404;  // 2403,2418 sp warp ; 2418,1428 sprites
const int  test2_offset = 2405; // 2495,2559
const unsigned char tester = 0x76;
// random, semi-decoded data or not implemented
const int  airship_offset = 2231;   // Lord J, CzarDragon credit
                        // bit 0: WoR: 0=sad song, 1=norm.song
                        // bit 1: airship visibility
                        // bit 2: 1==coin toss cutscean/crashed
                        // bit 3: start w/ Veldt Music
                        // bit 4: ?
                        // bit 5: ?
                        // bit 6: ?
                        // bit 7: 1=is a save point
const int  sound_at_startup_offset = 2230; // 99$ = save pnt snd
const int  rare_items_offset = 2234;       // 3 bytes, rare items: 24 items, list is 20 long in game, 4 items are "illegal"
const int  world_offset = 2404;     // offset 2405 was set to 32
                        // 00 is in balance world, 01 is in world of ruin,
                        // 02 air ship, serpent trench????
                        // 03 biz, enemies, tent south, warp to floating isl. cinema, ect., airship battle
                        // 04 tomb battle bkgnd, no warp, 05 no nothing, 06 "semi" stalled airship
                        // 07 airship casino deck,
                        // 09 team splitup (banon) scenario, 0a, "fake" airship
                        // 0b overwold balance, sabin scenario or doesnt affect?
                        // 0c casino deck fake airship,
                        // 0d end of world setup on airship
                        // 0e,10 chocobo stable, warp to last caverm
                        // 0f s.fig house warp, blocked
                        // 75 imperial camp
const int  overworld_palette = 2416;// not confirmed yet
const int  special_warp_offset = 2495;
                        // 00 = normal, 01 = lete river scenario w/ Banon!,
                        // 03,02,3f,80,8f,99 = freeze,
                        // 88 = walk around in spiral!,9f = air ship takeoff +freeze
                        // 7f = crash
const int  x1_sp_offset = 2496;     // x offset at save point, add = right off., subs = left off.
const int  y1_sp_offset = 2497;     // y offset at save point, add = down off., subs = up off.
offset 0x1cb9(slot 3), bit 4 ON force Narshe song
offset 0x1cad(slot 3), bit 7 ON make control of AS avail.
offset 0x1cd8(slot 3), bit 1 ON make "cinema black bar" in town
offset 0x1c94(slot 3), bit 7 ON make airship change to FALCON
offset 0x1c93(slot 3), bit 6 ON make continent disapear (WOB)
offset 0x1c93(slot 3), bit 7 ON Floating continent unselectable from AS
           but if bit 6 is OFF and 7 OFF, still unselectable
offset 0x08ed(slot 1), the way Cid's going to bed!!
offset 0x09d0(slot 1), cid's health on deserted island
offset 0x0896(slot 1), bit 3 ON cid's fine, stay on bed
offset 0x08ec(slot 1), bit 7 OFF cid is gone (hacking), cid is healthy (normal)
*/
