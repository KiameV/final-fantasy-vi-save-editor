package editors

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	weaponsTextBox      fyne.CanvasObject
	shieldsTextBox      fyne.CanvasObject
	helmetTextBox       fyne.CanvasObject
	armorTextBox        fyne.CanvasObject
	relic1TextBox       fyne.CanvasObject
	relic2TextBox       fyne.CanvasObject
	itemsTextBox        fyne.CanvasObject
	allEquipmentTextBox fyne.CanvasObject
	importItemsTextBox  fyne.CanvasObject
	mapsTextBox         fyne.CanvasObject
)

func CreateTextBoxes() {
	if weaponsTextBox == nil {
		weaponsTextBox = container.NewVScroll(widget.NewRichTextWithText(weaponsText))
		shieldsTextBox = container.NewVScroll(widget.NewRichTextWithText(shieldsText))
		helmetTextBox = container.NewVScroll(widget.NewRichTextWithText(helmetText))
		armorTextBox = container.NewVScroll(widget.NewRichTextWithText(armorText))
		relic1TextBox = container.NewVScroll(widget.NewRichTextWithText(relicText1))
		relic2TextBox = container.NewVScroll(widget.NewRichTextWithText(relicText2Header + relicText2))
		itemsTextBox = container.NewVScroll(widget.NewRichTextWithText(itemsText))
		allEquipmentTextBox = container.NewVScroll(widget.NewRichTextWithText(weaponsText + shieldsText + helmetText + armorText + relicText1 + relicText2))
		importItemsTextBox = container.NewVScroll(widget.NewRichTextWithText(importantItemsText))
		mapsTextBox = container.NewVScroll(widget.NewRichTextWithText(mapText))
	}
}

var (
	mapLookup = map[int]string{
		1:   "World of Balance",
		2:   "World of Ruin",
		3:   "Narshe (Ruin)",
		4:   "Narshe",
		5:   "Narshe",
		6:   "Narshe",
		7:   "Narshe",
		8:   "Narshe",
		9:   "Narshe",
		10:  "Narshe",
		11:  "Narshe",
		12:  "Narshe",
		13:  "Narshe",
		14:  "Narshe",
		15:  "Narshe",
		16:  "Narshe",
		17:  "Narshe",
		18:  "Narshe (Ruin)",
		19:  "Narshe (Ruin)",
		20:  "Narshe (Ruin)",
		21:  "Figaro Castle",
		22:  "Figaro Castle",
		23:  "Figaro Castle",
		24:  "Figaro Castle",
		25:  "Figaro Castle",
		26:  "Figaro Castle",
		27:  "Figaro Castle",
		28:  "Figaro Castle",
		29:  "Figaro Castle",
		30:  "Figaro Castle",
		31:  "Figaro Castle",
		32:  "Figaro Castle",
		33:  "Figaro Castle",
		34:  "Figaro Castle",
		35:  "Figaro Castle",
		36:  "Duncan's Cabin",
		37:  "Duncan's Cabin",
		38:  "South Figaro",
		39:  "South Figaro",
		40:  "South Figaro",
		41:  "South Figaro",
		42:  "South Figaro",
		43:  "South Figaro",
		44:  "South Figaro",
		45:  "South Figaro",
		46:  "South Figaro",
		47:  "South Figaro",
		48:  "South Figaro",
		49:  "South Figaro",
		50:  "South Figaro",
		51:  "South Figaro",
		52:  "South Figaro",
		53:  "South Figaro (Balance)",
		54:  "South Figaro",
		55:  "South Figaro",
		56:  "South Figaro",
		57:  "South Figaro",
		58:  "South Figaro",
		59:  "South Figaro",
		60:  "South Figaro",
		61:  "South Figaro",
		62:  "South Figaro",
		63:  "South Figaro",
		64:  "South Figaro",
		65:  "South Figaro",
		66:  "The Returners' Hideout",
		67:  "The Returners' Hideout",
		68:  "The Returners' Hideout",
		69:  "The Returners' Hideout",
		70:  "The Returners' Hideout",
		71:  "The Returners' Hideout",
		72:  "The Returners' Hideout",
		73:  "Cabin",
		74:  "Cabin",
		75:  "Doma Castle",
		76:  "Doma Castle",
		77:  "Doma Castle",
		78:  "Doma Castle",
		79:  "Doma Castle",
		80:  "Doma Castle",
		81:  "Doma Castle",
		82:  "Doma Castle",
		83:  "Doma Castle",
		84:  "Mobliz",
		85:  "Mobliz",
		86:  "Mobliz",
		87:  "Mobliz",
		88:  "Mobliz",
		89:  "Mobliz",
		90:  "Mobliz",
		91:  "Mobliz",
		92:  "Nikeah",
		93:  "Nikeah",
		94:  "Nikeah",
		95:  "Nikeah",
		96:  "Nikeah",
		97:  "Kohlingen",
		98:  "Kohlingen",
		99:  "Kohlingen",
		100: "Kohlingen",
		101: "Kohlingen",
		102: "Kohlingen",
		103: "Dragon's Neck Cabin",
		104: "Dragon's Neck Cabin",
		105: "Jidoor",
		106: "Jidoor",
		107: "Jidoor",
		108: "Jidoor",
		109: "Jidoor",
		110: "Jidoor",
		111: "Jidoor",
		112: "Opera House",
		113: "Opera House",
		114: "Opera House",
		115: "Opera House",
		116: "Opera House",
		117: "Opera House",
		118: "Opera House",
		119: "Blackjack",
		120: "Blackjack",
		121: "Blackjack",
		122: "Blackjack",
		123: "Blackjack",
		124: "Blackjack",
		125: "Albrook",
		126: "Albrook",
		127: "Albrook",
		128: "Albrook",
		129: "Albrook",
		130: "Albrook",
		131: "Albrook",
		132: "Tzen",
		133: "Tzen",
		134: "Tzen",
		135: "Tzen",
		136: "Tzen",
		137: "Tzen",
		138: "Maranda",
		139: "Maranda",
		140: "Maranda",
		141: "Maranda",
		142: "Maranda",
		143: "Vector (Balance)",
		144: "Vector",
		145: "Vector",
		146: "Vector",
		147: "Vector",
		148: "Vector",
		149: "Vector",
		150: "Vector",
		151: "Vector",
		152: "Vector",
		153: "Vector",
		154: "Vector",
		155: "Esper World",
		156: "Esper World",
		157: "Esper World",
		158: "Esper World",
		159: "Esper World",
		160: "Esper World",
		161: "Esper World",
		162: "Imperial Observation Post",
		163: "Imperial Observation Post",
		164: "Imperial Observation Post",
		165: "Imperial Palace",
		166: "Imperial Palace",
		167: "Imperial Palace",
		168: "Imperial Palace",
		169: "Imperial Palace",
		170: "Imperial Palace",
		171: "Imperial Palace",
		172: "Imperial Palace",
		173: "Imperial Palace",
		174: "Imperial Palace",
		175: "Imperial Palace",
		176: "Imperial Palace",
		177: "Imperial Palace",
		178: "Imperial Palace",
		179: "Imperial Palace",
		180: "Thamasa",
		181: "Thamasa",
		182: "Thamasa",
		183: "Thamasa",
		184: "Thamasa",
		185: "Thamasa",
		186: "Thamasa",
		187: "Thamasa",
		188: "Chocobo Stable",
		189: "Chocobo Stable",
		190: "Solitary Island",
		191: "Solitary Island",
		192: "Solitary Island",
		193: "Solitary Island",
		194: "Albrook",
		195: "Albrook",
		196: "Albrook",
		197: "Albrook",
		198: "Albrook",
		199: "Albrook",
		200: "Albrook",
		201: "Tzen",
		202: "Tzen",
		203: "Tzen",
		204: "Tzen",
		205: "Tzen",
		206: "Tzen",
		207: "Mobliz",
		208: "Mobliz",
		209: "Mobliz",
		210: "Mobliz",
		211: "Mobliz",
		212: "Mobliz",
		213: "Mobliz",
		214: "Nikeah",
		215: "Nikeah",
		216: "Nikeah",
		217: "Nikeah",
		218: "Nikeah",
		219: "South Figaro",
		220: "South Figaro",
		221: "South Figaro",
		222: "South Figaro",
		223: "South Figaro",
		224: "South Figaro",
		225: "South Figaro",
		226: "South Figaro",
		227: "South Figaro",
		228: "South Figaro",
		229: "South Figaro",
		230: "South Figaro",
		231: "South Figaro",
		232: "South Figaro",
		233: "South Figaro",
		234: "South Figaro (Balance)",
		235: "South Figaro",
		236: "South Figaro",
		237: "South Figaro",
		238: "South Figaro",
		239: "South Figaro",
		240: "South Figaro",
		241: "South Figaro",
		242: "South Figaro",
		243: "South Figaro",
		244: "South Figaro",
		245: "South Figaro",
		246: "South Figaro",
		247: "Figaro Castle",
		248: "Figaro Castle",
		249: "Figaro Castle",
		250: "Figaro Castle",
		251: "Figaro Castle",
		252: "Figaro Castle",
		253: "Figaro Castle",
		254: "Figaro Castle",
		255: "Figaro Castle",
		256: "Figaro Castle",
		257: "Figaro Castle",
		258: "Figaro Castle",
		259: "Figaro Castle",
		260: "Figaro Castle",
		261: "Figaro Castle",
		262: "Kohlingen",
		263: "Kohlingen",
		264: "Kohlingen",
		265: "Kohlingen",
		266: "Kohlingen",
		267: "Kohlingen",
		268: "Falcon",
		269: "Falcon",
		270: "Falcon",
		271: "Falcon",
		272: "Duncan's Cabin",
		273: "Duncan's Cabin",
		274: "Maranda",
		275: "Maranda",
		276: "Maranda",
		277: "Maranda",
		278: "Maranda",
		279: "Coliseum",
		280: "Coliseum",
		281: "Coliseum",
		282: "Jidoor",
		283: "Jidoor",
		284: "Jidoor",
		285: "Jidoor",
		286: "Jidoor",
		287: "Jidoor",
		288: "Jidoor",
		289: "Jidoor",
		290: "Opera House",
		291: "Opera House",
		292: "Opera House",
		293: "Opera House",
		294: "Opera House",
		295: "Opera House",
		296: "Doma Castle",
		297: "Doma Castle",
		298: "Doma Castle",
		299: "Doma Castle",
		300: "Doma Castle",
		301: "Doma Castle",
		302: "Doma Castle",
		303: "Doma Castle",
		304: "Doma Castle",
		305: "Thamasa",
		306: "Thamasa",
		307: "Thamasa",
		308: "Thamasa",
		309: "Thamasa",
		310: "Thamasa",
		311: "Thamasa",
		312: "Thamasa",
		313: "Cabin",
		314: "Cabin",
		315: "Chocobo Stable",
		316: "Chocobo Stable",
		317: "Narshe Mines (Balance)",
		318: "Narshe Mines (Ruin)",
		319: "Narshe Mines",
		320: "Narshe Mines (Ruin)",
		321: "Narshe Mines (Balance)",
		322: "Narshe Mines (Balance)",
		323: "Narshe Mines (Balance)",
		324: "Narshe Mines (Balance)",
		325: "Narshe Mines (Balance)",
		326: "Narshe Mines (Balance)",
		327: "Narshe Mines (Balance)",
		328: "South Figaro Cave",
		329: "South Figaro Cave (Balance)",
		330: "South Figaro Cave (Balance)",
		331: "South Figaro Cave (Balance)",
		332: "South Figaro Cave (Balance)",
		333: "South Figaro Cave (Balance)",
		334: "Mt. Kolts (Balance)",
		335: "Mt. Kolts (Balance)",
		336: "Mt. Kolts (Balance)",
		337: "Mt. Kolts (Balance)",
		338: "Mt. Kolts (Balance)",
		339: "Mt. Kolts (Balance)",
		340: "Mt. Kolts",
		341: "Mt. Kolts (Balance)",
		342: "Mt. Kolts",
		343: "Mt. Kolts (Balance)",
		344: "Mt. Kolts",
		345: "Lethe River",
		346: "Lethe River",
		347: "Lethe River",
		348: "Lethe River",
		349: "Lethe River",
		350: "Imperial Camp",
		351: "Phantom Forest (Balance)",
		352: "Phantom Forest (Balance)",
		353: "Phantom Forest (Balance)",
		354: "Phantom Forest (Balance)",
		355: "Platform",
		356: "Phantom Train (Balance)",
		357: "Phantom Train (Balance)",
		358: "Phantom Train (Balance)",
		359: "Phantom Train",
		360: "Phantom Train (Balance)",
		361: "Phantom Train (Balance)",
		362: "Phantom Train (Balance)",
		363: "Phantom Train",
		364: "Phantom Train (Balance)",
		365: "Phantom Train",
		366: "Phantom Train",
		367: "Phantom Train (Balance)",
		368: "Phantom Train",
		369: "Phantom Train",
		370: "Phantom Train",
		371: "Phantom Train",
		372: "Platform",
		373: "Cave to Baren Falls",
		374: "Cave to Baren Falls",
		375: "Baren Falls",
		376: "Riverside",
		377: "Crescent Mountain Cave",
		378: "Crescent Mountain Cave",
		379: "Serpent Trench",
		380: "Serpent Trench",
		381: "Serpent Trench",
		382: "Serpent Trench",
		383: "Zozo (Balance)",
		384: "Zozo (Balance)",
		385: "Zozo (Balance)",
		386: "Zozo (Balance)",
		387: "Zozo (Balance)",
		388: "Zozo (Balance)",
		389: "Zozo (Balance)",
		390: "Zozo (Balance)",
		391: "Zozo (Balance)",
		392: "Zozo (Balance)",
		393: "Zozo (Balance)",
		394: "Zozo (Balance)",
		395: "Zozo",
		396: "Magitek Factory (Balance)",
		397: "Magitek Factory (Balance)",
		398: "Magitek Factory (Balance)",
		399: "Magitek Factory",
		400: "Magitek Factory (Balance)",
		401: "Magitek Research Facility (Balance)",
		402: "Magitek Research Facility (Balance)",
		403: "Magitek Research Facility",
		404: "Magitek Research Facility",
		405: "Cave to the Sealed Gate (Balance)",
		406: "Cave to the Sealed Gate (Balance)",
		407: "Cave to the Sealed Gate (Balance)",
		408: "Cave to the Sealed Gate (Balance)",
		409: "Cave to the Sealed Gate",
		410: "Cave to the Sealed Gate (Balance)",
		411: "Cave to the Sealed Gate (Balance)",
		412: "Burning Home",
		413: "Burning Home",
		414: "Burning Home",
		415: "Burning Home",
		416: "Burning Home",
		417: "Burning Home",
		418: "Burning Home",
		419: "Burning Home",
		420: "Burning Home",
		421: "Esper Cave (Balance)",
		422: "Esper Cave (Balance)",
		423: "Esper Cave (Balance)",
		424: "Esper Cave (Balance)",
		425: "Esper Cave",
		426: "Esper Cave (Balance)",
		427: "Esper Cave (Balance)",
		428: "Esper Cave (Balance)",
		429: "Esper Cave (Balance)",
		430: "Esper Cave (Balance)",
		431: "Floating Continent (Balance)",
		432: "Floating Continent",
		433: "Crumbling House (Ruin)",
		434: "Crumbling House (Ruin)",
		435: "Figaro Castle (Ruin)",
		436: "Figaro Castle (Ruin)",
		437: "Figaro Castle (Ruin)",
		438: "Figaro Castle (Ruin)",
		439: "Figaro Castle",
		440: "Figaro Castle (Ruin)",
		441: "Figaro Castle",
		442: "Owzer's Mansion",
		443: "Owzer's Mansion (Ruin)",
		444: "Owzer's Mansion (Ruin)",
		445: "Owzer's Mansion (Ruin)",
		446: "Owzer's Mansion",
		447: "Scenario Selection",
		448: "South Figaro Cave (Ruin)",
		449: "South Figaro Cave (Ruin)",
		450: "South Figaro Cave (Ruin)",
		451: "South Figaro Cave (Ruin)",
		452: "South Figaro Cave (Ruin)",
		453: "South Figaro Cave (Ruin)",
		454: "South Figaro Cave (Ruin)",
		455: "Darill's Tomb",
		456: "Darill's Tomb (Ruin)",
		457: "Darill's Tomb (Ruin)",
		458: "Darill's Tomb (Ruin)",
		459: "Darill's Tomb (Ruin)",
		460: "Darill's Tomb (Ruin)",
		461: "Darill's Tomb (Ruin)",
		462: "Darill's Tomb (Ruin)",
		463: "Darill's Tomb (Ruin)",
		464: "Darill's Tomb (Ruin)",
		465: "Darill's Tomb (Ruin)",
		466: "Darill's Tomb (Ruin)",
		467: "Darill's Tomb (Ruin)",
		468: "Darill's Tomb (Ruin)",
		469: "Mt. Zozo (Ruin)",
		470: "Mt. Zozo (Ruin)",
		471: "Mt. Zozo (Ruin)",
		472: "Mt. Zozo (Ruin)",
		473: "Mt. Zozo (Ruin)",
		474: "Mt. Zozo",
		475: "Mt. Zozo",
		476: "Cultists' Tower",
		477: "Cultists' Tower (Ruin)",
		478: "Cultists' Tower (Ruin)",
		479: "Cultists' Tower (Ruin)",
		480: "Cultists' Tower (Ruin)",
		481: "Cultists' Tower (Ruin)",
		482: "Cultists' Tower (Ruin)",
		483: "Cultists' Tower (Ruin)",
		484: "Cultists' Tower (Ruin)",
		485: "Cultists' Tower (Ruin)",
		486: "Cultists' Tower (Ruin)",
		487: "Cultists' Tower (Ruin)",
		488: "Cave on the Veldt (Ruin)",
		489: "Cave on the Veldt (Ruin)",
		490: "Cave on the Veldt (Ruin)",
		491: "Cave on the Veldt (Ruin)",
		492: "Cave on the Veldt (Ruin)",
		493: "Yeti's Cave (Ruin)",
		494: "Yeti's Cave (Ruin)",
		495: "Yeti's Cave (Ruin)",
		496: "Yeti's Cave (Ruin)",
		497: "Phoenix Cave",
		498: "Phoenix Cave (Ruin)",
		499: "Phoenix Cave (Ruin)",
		500: "Zone Eater's Belly (Ruin)",
		501: "Zone Eater's Belly (Ruin)",
		502: "Zone Eater's Belly (Ruin)",
		503: "Zone Eater's Belly",
		504: "Zone Eater's Belly",
		505: "Zone Eater's Belly (Ruin)",
		506: "Zone Eater's Belly",
		507: "Zone Eater's Belly (Ruin)",
		508: "Cave to the Ancient Castle (Ruin)",
		509: "Cave to the Ancient Castle (Ruin)",
		510: "Cave to the Ancient Castle (Ruin)",
		511: "Cave to the Ancient Castle",
		512: "Ancient Castle (Ruin)",
		513: "Ancient Castle (Ruin)",
		514: "Ancient Castle (Ruin)",
		515: "Ancient Castle (Ruin)",
		516: "Ancient Castle (Ruin)",
		517: "Ancient Castle (Ruin)",
		518: "Ancient Castle",
		519: "Ancient Castle",
		520: "Dreamscape (Ruin)",
		521: "Dreamscape (Ruin)",
		522: "Dreamscape - Phantom Train",
		523: "Dreamscape - Phantom Train (Ruin)",
		524: "Dreamscape - Phantom Train (Ruin)",
		525: "Dreamscape - Phantom Train (Ruin)",
		526: "Dreamscape - Phantom Train",
		527: "Dreamscape - Phantom Train",
		528: "Dreamscape - Cave (Ruin)",
		529: "Dreamscape - Cave (Ruin)",
		530: "Dreamscape - Doma Castle (Ruin)",
		531: "Dreamscape - Doma Castle (Ruin)",
		532: "Dreamscape - Doma Castle (Ruin)",
		533: "Dreamscape - Doma Castle (Ruin)",
		534: "Dreamscape - Doma Castle (Ruin)",
		535: "Dreamscape - Doma Castle (Ruin)",
		536: "Dreamscape - Doma Castle (Ruin)",
		537: "Dreamscape - Doma Castle (Ruin)",
		538: "Dreamscape - Doma Castle (Ruin)",
		539: "Ebot's Rock (Ruin)",
		540: "Ebot's Rock (Ruin)",
		541: "Ebot's Rock (Ruin)",
		542: "Kefka's Tower (Ruin)",
		543: "Kefka's Tower (Ruin)",
		544: "Kefka's Tower (Ruin)",
		545: "Kefka's Tower",
		546: "Kefka's Tower (Ruin)",
		547: "Kefka's Tower (Ruin)",
		548: "Kefka's Tower (Ruin)",
		549: "Kefka's Tower",
		550: "Kefka's Tower (Ruin)",
		551: "Kefka's Tower (Ruin)",
		552: "Kefka's Tower (Ruin)",
		553: "Kefka's Tower (Ruin)",
		554: "Kefka's Tower (Ruin)",
		555: "Kefka's Tower (Ruin)",
		556: "Kefka's Tower (Ruin)",
		557: "Kefka's Tower (Ruin)",
		558: "Kefka's Tower (Ruin)",
		559: "Kefka's Tower (Ruin)",
		560: "Kefka's Tower (Ruin)",
		561: "Kefka's Tower (Ruin)",
		562: "Kefka's Tower (Ruin)",
		563: "Kefka's Tower (Ruin)",
		564: "Kefka's Tower (Ruin)",
		565: "Kefka's Tower (Ruin)",
		566: "Kefka's Tower (Ruin)",
		567: "Kefka's Tower",
		568: "Kefka's Tower",
		569: "Kefka's Tower",
		570: "Kefka's Tower",
		571: "Kefka's Tower",
		572: "Overworld",
		573: "Overworld",
		574: "Opera House",
		575: "Overworld",
		576: "Opera House",
		577: "Darill's Tomb",
		578: "Overworld",
		579: "Darill's Tomb",
		580: "Narshe Mines",
		581: "Narshe",
		582: "Blackjack",
		583: "Mobliz",
		584: "Scenario Selection",
		585: "Falcon",
		586: "Overworld",
		587: "Figaro Castle",
		588: "Figaro Castle",
		589: "Figaro Castle",
		590: "Kefka's Tower",
		591: "Scenario Selection",
		592: "Scenario Selection",
		593: "Scenario Selection",
		594: "Scenario Selection",
		595: "Scenario Selection",
		596: "Scenario Selection",
		597: "Solitary Island",
		598: "Overworld",
		599: "Overworld",
		600: "Overworld",
		601: "Scenario Selection",
		602: "Scenario Selection",
		603: "Scenario Selection",
		604: "Esper World",
		605: "Scenario Selection",
		606: "",
		607: "",
		608: "",
		609: "Narshe",
		610: "Jidoor",
		611: "Jidoor",
		612: "Chocobo Stable",
		613: "Chocobo Stable",
		614: "Chocobo Stable",
		615: "Chocobo Stable",
		616: "Chocobo Stable",
		617: "Chocobo Stable",
		618: "Chocobo Stable",
		619: "Chocobo Stable",
		620: "Chocobo Stable",
		621: "Chocobo Stable",
		622: "Nikeah",
		623: "Albrook",
		624: "Phantom Train (Balance)",
		625: "Magitek Research Facility",
		626: "South Figaro Cave (Balance)",
		627: "South Figaro Cave (Balance)",
		628: "South Figaro Cave (Balance)",
		629: "South Figaro Cave (Balance)",
		630: "South Figaro Cave (Balance)",
		631: "Floating Continent (Balance)",
		632: "Narshe Mines (Ruin)",
		633: "Narshe Mines (Ruin)",
		634: "Narshe Mines (Ruin)",
		635: "Narshe Mines (Ruin)",
		636: "Narshe Mines (Ruin)",
		637: "Narshe Mines (Ruin)",
		638: "Narshe Mines (Ruin)",
		639: "Narshe Mines (Ruin)",
		640: "",
		641: "Floating Continent",
		642: "Mt. Kolts",
		643: "Magitek Factory (Balance)",
		800: "Overworld",
		801: "Overworld",
		802: "Overworld",
		803: "Overworld",
		804: "Overworld",
		805: "Overworld",
	}
)

const (
	emptyText = `0 - Invalid
`
	emptyWeaponShield = `93 - Empty
`
	emptyHelmet = `199 - Empty
`
	emptyArmor = `198 - Empty
`
	emptyRelic = `200 - Empty
`
	itemsText = `Miscellaneous
2 - Potion
3 - Hi-Potion
4 - X-Potion
5 - Ether
6 - Hi-Ether
7 - X-Ether
8 - Elixir
9 - Megalixir
10 - Fenix Down
11 - Holy Water
12 - Antidote
13 - Eye Drops
14 - Golden Needle
15 - Remedy
16 - Sleeping Bag
17 - Tent
18 - Green Cherry
19 - Magicite Shard
20 - Super Ball
21 - Echo Screen
22 - Smoke Bomb
23 - Teleport Stone
24 - Dried Meat
25 - Rename Card

Tool
31 - Noiseblaster
32 - Bioblaster
33 - Flash
34 - ChainSaw
35 - Debilitator
36 - Drill
37 - Air Anchor
38 - Auto Crossbow

Scroll
26 - Fire Scroll
27 - Water Scroll
28 - Lightning Scroll
29 - Invisibility Scroll
30 - Shadow Scroll`

	weaponsText = `Dirk
94 - Dagger
95 - Mythril Knife
96 - Main Gauche
97 - Air Knife
98 - Thief's Knife
99 - Assassin's Dagger
100 - Man-Eater
101 - Swordbreaker
102 - Gladius
103 - Valiant Knife

Sword
104 - Mythril Sword
105 - Great Sword
106 - Rune Blade
107 - Flametongue
108 - Icebrand
109 - Thunder Blade
110 - Bastard Sword 
111 - Break Blade
112 - Blood Sword
113 - Enhancer
114 - Crystal Sword
115 - Falchion
116 - Soul Sabre
117 - Organyx
118 - Excalibur
119 - Zantetsuken
120 - Lightbringer
121 - Ragnarok
122 - Ultima Weapon

Lance
123 - Mythril Spear
124 - Trident
125 - Heavy Lance
126 - Partisan
127 - Holy Lance
128 - Golden Lance
129 - Radiant Lance
130 - Impartisan

Dirk
131 - Kunai
132 - Kodachi
133 - Sakura
134 - Sasuke
135 - Ichigeki
136 - Kagenui

Katana
137 - Ashura
138 - Kotetsu
139 - Kikuichimonji
140 - Kazekiri
141 - Murasame
142 - Masamune
143 - Murakumo
144 - Mutsunokami

Rod
145 - Heal Rod
146 - Mythril Rod
147 - Flame Rod
148 - Ice Rod
149 - Thunder Rod
150 - Poison Rod
151 - Holy Rod
152 - Gravity Rod
153 - Punisher
154 - Magus Rod

Brush
155 - Chocobo Brush
156 - Da Vinci Brush
157 - Magical Brush
158 - Rainbow Brush

Stars
159 - Shuriken
160 - Fuma Shuriken
161 - Pinwheel

Special
162 - Chain Flail
163 - Moonring Blade
164 - Morning Star
165 - Boomerang
166 - Rising Sun
167 - Hawkeye
168 - Bone Club
169 - Sniper
170 - Wing Edge

Gambler
171 - Cards
172 - Darts
173 - Death Tarot
174 - Viper Darts
175 - Dice
176 - Fixed Dice

Claw
177 - Metal Knuckles
178 - Mythril Claws
179 - Kaiser Knuckles
180 - Venom Claws
181 - Burning Fist
182 - Dragon Claws
183 - Tigerfangs`

	shieldsText = `Shield
201 - Buckler
202 - Heavy Shield
203 - Mythril Shield
204 - Gold Shield
205 - Aegis Shield
206 - Diamond Shield
207 - Flame Shield
208 - Ice Shield
209 - Thunder Shield
210 - Crystal Shield
211 - Genji Shield
212 - Tortoise Shield
213 - Cursed Shield
214 - Paladin Shield
215 - Force Shield`

	helmetText = `Helmet
216 - Leather Hat
217 - Hairband
218 - Plumed Hat
219 - Beret
220 - Magus Hat
221 - Bandana
222 - Iron Helmet
223 - Priest's Miter
224 - Bard's Hat
225 - Green Beret
226 - Head Band
227 - Mythril Helm
228 - Tiara
229 - Gold Helmet
230 - Tiger Mask
231 - Red Hat
232 - Mystery Veil
233 - Circlet
234 - Royal Crown
235 - Diamond Helm
236 - Black Hood
237 - Crystal Helm
238 - Oath Veil
239 - Cat-Ear Hood
240 - Genji Helmet
241 - Thornlet
242 - Saucer`

	armorText = `Armor
244 - Leather Armor
245 - Cotton Robe
246 - Kenpo Gi
247 - Iron Armor
248 - Silk Robe
249 - Mythril Vest
250 - Ninja Gear
251 - White Dress
252 - Mythril Mail
253 - Gaia Gear
254 - Mirage Dress
255 - Golden Armor
256 - Power Sash
257 - Luminous Robe
258 - Diamond Vest
259 - Red Jacket
260 - Force Armor
261 - Diamond Armor
262 - Black Garb
263 - Magus Rove
264 - Crystal Mail
265 - Regal Gown
266 - Genji Armor
267 - Reed Cloak
268 - Minerva Bustier
269 - Tabby Suit
270 - Chocobo Suit
271 - Moogle Suit
272 - Nutkin Suit
273 - Behemeth Suit
274 - Snow Scarf`

	relicText1 = `Relic
275 - Silver Spectacles
276 - Star Pendant
277 - Peace Ring
278 - Amulet
279 - White Cape
280 - Jewel Ring
281 - Fairy Ring
282 - Barrier Ring
283 - Mythril Glove
284 - Protect Ring
285 - Hermes Sandals
286 - Reflect Ring
287 - Angel Wings
288 - Angel Ring
289 - Knight's Code
290 - Dragoon Boots
291 - Zephyr Cloak
292 - Princess Ring
293 - Cursed Ring
294 - Earring
295 - Gigas Glove
296 - Blizzard Orb
297 - Berserker Ring
298 - Thief's Bracer
299 - Guard Bracelet
`
	relicText2Header = `Relic (continued)
`
	relicText2 = `300 - Hero Ring
301 - Ribbon
302 - Muscle Belt
303 - Crystal Orb
304 - Gold Hairpin
305 - Celestriad
306 - Brigand's Glove
307 - Gauntlet
308 - Genji Glove
309 - Hyper Wrist
310 - Master's Scroll
311 - Prayer Beads
312 - Black Belt
313 - Heiki's Jitte
314 - Fake Mustache
315 - Soul of Thamasa
316 - Dragon Horn
317 - Merit Award
318 - Momento Ring
319 - Safety Bit
320 - Lich Ring
321 - Molulu's Charm
322 - Ward Bangle
323 - Miracle Shoes
324 - Alarm Gaurd
325 - Gale Hairpin
326 - Sniper Eye
327 - Growth Egg
328 - Tintinnabulum
329 - Sprint Shoes`
	importantItemsText = `39 - Cider
40 - Old Clock Key
41 - Fish
42 - Fish
43 - Fish
44 - Fish
45 - Lump of Metal
46 - Lola's Letter
47 - Coral
48 - Books
49 - Emperor's Letter
50 - Rush-Rid
58 - Pendant`
	mapText = `1: World of Balance
2: World of Ruin
3: Narshe (Ruin)
4: Narshe
5: Narshe
6: Narshe
7: Narshe
8: Narshe
9: Narshe
10: Narshe
11: Narshe
12: Narshe
13: Narshe
14: Narshe
15: Narshe
16: Narshe
17: Narshe
18: Narshe (Ruin)
19: Narshe (Ruin)
20: Narshe (Ruin)
21: Figaro Castle
22: Figaro Castle
23: Figaro Castle
24: Figaro Castle
25: Figaro Castle
26: Figaro Castle
27: Figaro Castle
28: Figaro Castle
29: Figaro Castle
30: Figaro Castle
31: Figaro Castle
32: Figaro Castle
33: Figaro Castle
34: Figaro Castle
35: Figaro Castle
36: Duncan's Cabin
37: Duncan's Cabin
38: South Figaro
39: South Figaro
40: South Figaro
41: South Figaro
42: South Figaro
43: South Figaro
44: South Figaro
45: South Figaro
46: South Figaro
47: South Figaro
48: South Figaro
49: South Figaro
50: South Figaro
51: South Figaro
52: South Figaro
53: South Figaro (Balance)
54: South Figaro
55: South Figaro
56: South Figaro
57: South Figaro
58: South Figaro
59: South Figaro
60: South Figaro
61: South Figaro
62: South Figaro
63: South Figaro
64: South Figaro
65: South Figaro
66: The Returners' Hideout
67: The Returners' Hideout
68: The Returners' Hideout
69: The Returners' Hideout
70: The Returners' Hideout
71: The Returners' Hideout
72: The Returners' Hideout
73: Cabin
74: Cabin
75: Doma Castle
76: Doma Castle
77: Doma Castle
78: Doma Castle
79: Doma Castle
80: Doma Castle
81: Doma Castle
82: Doma Castle
83: Doma Castle
84: Mobliz
85: Mobliz
86: Mobliz
87: Mobliz
88: Mobliz
89: Mobliz
90: Mobliz
91: Mobliz
92: Nikeah
93: Nikeah
94: Nikeah
95: Nikeah
96: Nikeah
97: Kohlingen
98: Kohlingen
99: Kohlingen
100: Kohlingen
101: Kohlingen
102: Kohlingen
103: Dragon's Neck Cabin
104: Dragon's Neck Cabin
105: Jidoor
106: Jidoor
107: Jidoor
108: Jidoor
109: Jidoor
110: Jidoor
111: Jidoor
112: Opera House
113: Opera House
114: Opera House
115: Opera House
116: Opera House
117: Opera House
118: Opera House
119: Blackjack
120: Blackjack
121: Blackjack
122: Blackjack
123: Blackjack
124: Blackjack
125: Albrook
126: Albrook
127: Albrook
128: Albrook
129: Albrook
130: Albrook
131: Albrook
132: Tzen
133: Tzen
134: Tzen
135: Tzen
136: Tzen
137: Tzen
138: Maranda
139: Maranda
140: Maranda
141: Maranda
142: Maranda
143: Vector (Balance)
144: Vector
145: Vector
146: Vector
147: Vector
148: Vector
149: Vector
150: Vector
151: Vector
152: Vector
153: Vector
154: Vector
155: Esper World
156: Esper World
157: Esper World
158: Esper World
159: Esper World
160: Esper World
161: Esper World
162: Imperial Observation Post
163: Imperial Observation Post
164: Imperial Observation Post
165: Imperial Palace
166: Imperial Palace
167: Imperial Palace
168: Imperial Palace
169: Imperial Palace
170: Imperial Palace
171: Imperial Palace
172: Imperial Palace
173: Imperial Palace
174: Imperial Palace
175: Imperial Palace
176: Imperial Palace
177: Imperial Palace
178: Imperial Palace
179: Imperial Palace
180: Thamasa
181: Thamasa
182: Thamasa
183: Thamasa
184: Thamasa
185: Thamasa
186: Thamasa
187: Thamasa
188: Chocobo Stable
189: Chocobo Stable
190: Solitary Island
191: Solitary Island
192: Solitary Island
193: Solitary Island
194: Albrook
195: Albrook
196: Albrook
197: Albrook
198: Albrook
199: Albrook
200: Albrook
201: Tzen
202: Tzen
203: Tzen
204: Tzen
205: Tzen
206: Tzen
207: Mobliz
208: Mobliz
209: Mobliz
210: Mobliz
211: Mobliz
212: Mobliz
213: Mobliz
214: Nikeah
215: Nikeah
216: Nikeah
217: Nikeah
218: Nikeah
219: South Figaro
220: South Figaro
221: South Figaro
222: South Figaro
223: South Figaro
224: South Figaro
225: South Figaro
226: South Figaro
227: South Figaro
228: South Figaro
229: South Figaro
230: South Figaro
231: South Figaro
232: South Figaro
233: South Figaro
234: South Figaro (Balance)
235: South Figaro
236: South Figaro
237: South Figaro
238: South Figaro
239: South Figaro
240: South Figaro
241: South Figaro
242: South Figaro
243: South Figaro
244: South Figaro
245: South Figaro
246: South Figaro
247: Figaro Castle
248: Figaro Castle
249: Figaro Castle
250: Figaro Castle
251: Figaro Castle
252: Figaro Castle
253: Figaro Castle
254: Figaro Castle
255: Figaro Castle
256: Figaro Castle
257: Figaro Castle
258: Figaro Castle
259: Figaro Castle
260: Figaro Castle
261: Figaro Castle
262: Kohlingen
263: Kohlingen
264: Kohlingen
265: Kohlingen
266: Kohlingen
267: Kohlingen
268: Falcon
269: Falcon
270: Falcon
271: Falcon
272: Duncan's Cabin
273: Duncan's Cabin
274: Maranda
275: Maranda
276: Maranda
277: Maranda
278: Maranda
279: Coliseum
280: Coliseum
281: Coliseum
282: Jidoor
283: Jidoor
284: Jidoor
285: Jidoor
286: Jidoor
287: Jidoor
288: Jidoor
289: Jidoor
290: Opera House
291: Opera House
292: Opera House
293: Opera House
294: Opera House
295: Opera House
296: Doma Castle
297: Doma Castle
298: Doma Castle
299: Doma Castle
300: Doma Castle
301: Doma Castle
302: Doma Castle
303: Doma Castle
304: Doma Castle
305: Thamasa
306: Thamasa
307: Thamasa
308: Thamasa
309: Thamasa
310: Thamasa
311: Thamasa
312: Thamasa
313: Cabin
314: Cabin
315: Chocobo Stable
316: Chocobo Stable
317: Narshe Mines (Balance)
318: Narshe Mines (Ruin)
319: Narshe Mines
320: Narshe Mines (Ruin)
321: Narshe Mines (Balance)
322: Narshe Mines (Balance)
323: Narshe Mines (Balance)
324: Narshe Mines (Balance)
325: Narshe Mines (Balance)
326: Narshe Mines (Balance)
327: Narshe Mines (Balance)
328: South Figaro Cave
329: South Figaro Cave (Balance)
330: South Figaro Cave (Balance)
331: South Figaro Cave (Balance)
332: South Figaro Cave (Balance)
333: South Figaro Cave (Balance)
334: Mt. Kolts (Balance)
335: Mt. Kolts (Balance)
336: Mt. Kolts (Balance)
337: Mt. Kolts (Balance)
338: Mt. Kolts (Balance)
339: Mt. Kolts (Balance)
340: Mt. Kolts
341: Mt. Kolts (Balance)
342: Mt. Kolts
343: Mt. Kolts (Balance)
344: Mt. Kolts
345: Lethe River
346: Lethe River
347: Lethe River
348: Lethe River
349: Lethe River
350: Imperial Camp
351: Phantom Forest (Balance)
352: Phantom Forest (Balance)
353: Phantom Forest (Balance)
354: Phantom Forest (Balance)
355: Platform
356: Phantom Train (Balance)
357: Phantom Train (Balance)
358: Phantom Train (Balance)
359: Phantom Train
360: Phantom Train (Balance)
361: Phantom Train (Balance)
362: Phantom Train (Balance)
363: Phantom Train
364: Phantom Train (Balance)
365: Phantom Train
366: Phantom Train
367: Phantom Train (Balance)
368: Phantom Train
369: Phantom Train
370: Phantom Train
371: Phantom Train
372: Platform
373: Cave to Baren Falls
374: Cave to Baren Falls
375: Baren Falls
376: Riverside
377: Crescent Mountain Cave
378: Crescent Mountain Cave
379: Serpent Trench
380: Serpent Trench
381: Serpent Trench
382: Serpent Trench
383: Zozo (Balance)
384: Zozo (Balance)
385: Zozo (Balance)
386: Zozo (Balance)
387: Zozo (Balance)
388: Zozo (Balance)
389: Zozo (Balance)
390: Zozo (Balance)
391: Zozo (Balance)
392: Zozo (Balance)
393: Zozo (Balance)
394: Zozo (Balance)
395: Zozo
396: Magitek Factory (Balance)
397: Magitek Factory (Balance)
398: Magitek Factory (Balance)
399: Magitek Factory
400: Magitek Factory (Balance)
401: Magitek Research Facility (Balance)
402: Magitek Research Facility (Balance)
403: Magitek Research Facility
404: Magitek Research Facility
405: Cave to the Sealed Gate (Balance)
406: Cave to the Sealed Gate (Balance)
407: Cave to the Sealed Gate (Balance)
408: Cave to the Sealed Gate (Balance)
409: Cave to the Sealed Gate
410: Cave to the Sealed Gate (Balance)
411: Cave to the Sealed Gate (Balance)
412: Burning Home
413: Burning Home
414: Burning Home
415: Burning Home
416: Burning Home
417: Burning Home
418: Burning Home
419: Burning Home
420: Burning Home
421: Esper Cave (Balance)
422: Esper Cave (Balance)
423: Esper Cave (Balance)
424: Esper Cave (Balance)
425: Esper Cave
426: Esper Cave (Balance)
427: Esper Cave (Balance)
428: Esper Cave (Balance)
429: Esper Cave (Balance)
430: Esper Cave (Balance)
431: Floating Continent (Balance)
432: Floating Continent
433: Crumbling House (Ruin)
434: Crumbling House (Ruin)
435: Figaro Castle (Ruin)
436: Figaro Castle (Ruin)
437: Figaro Castle (Ruin)
438: Figaro Castle (Ruin)
439: Figaro Castle
440: Figaro Castle (Ruin)
441: Figaro Castle
442: Owzer's Mansion
443: Owzer's Mansion (Ruin)
444: Owzer's Mansion (Ruin)
445: Owzer's Mansion (Ruin)
446: Owzer's Mansion
447: Scenario Selection
448: South Figaro Cave (Ruin)
449: South Figaro Cave (Ruin)
450: South Figaro Cave (Ruin)
451: South Figaro Cave (Ruin)
452: South Figaro Cave (Ruin)
453: South Figaro Cave (Ruin)
454: South Figaro Cave (Ruin)
455: Darill's Tomb
456: Darill's Tomb (Ruin)
457: Darill's Tomb (Ruin)
458: Darill's Tomb (Ruin)
459: Darill's Tomb (Ruin)
460: Darill's Tomb (Ruin)
461: Darill's Tomb (Ruin)
462: Darill's Tomb (Ruin)
463: Darill's Tomb (Ruin)
464: Darill's Tomb (Ruin)
465: Darill's Tomb (Ruin)
466: Darill's Tomb (Ruin)
467: Darill's Tomb (Ruin)
468: Darill's Tomb (Ruin)
469: Mt. Zozo (Ruin)
470: Mt. Zozo (Ruin)
471: Mt. Zozo (Ruin)
472: Mt. Zozo (Ruin)
473: Mt. Zozo (Ruin)
474: Mt. Zozo
475: Mt. Zozo
476: Cultists' Tower
477: Cultists' Tower (Ruin)
478: Cultists' Tower (Ruin)
479: Cultists' Tower (Ruin)
480: Cultists' Tower (Ruin)
481: Cultists' Tower (Ruin)
482: Cultists' Tower (Ruin)
483: Cultists' Tower (Ruin)
484: Cultists' Tower (Ruin)
485: Cultists' Tower (Ruin)
486: Cultists' Tower (Ruin)
487: Cultists' Tower (Ruin)
488: Cave on the Veldt (Ruin)
489: Cave on the Veldt (Ruin)
490: Cave on the Veldt (Ruin)
491: Cave on the Veldt (Ruin)
492: Cave on the Veldt (Ruin)
493: Yeti's Cave (Ruin)
494: Yeti's Cave (Ruin)
495: Yeti's Cave (Ruin)
496: Yeti's Cave (Ruin)
497: Phoenix Cave
498: Phoenix Cave (Ruin)
499: Phoenix Cave (Ruin)
500: Zone Eater's Belly (Ruin)
501: Zone Eater's Belly (Ruin)
502: Zone Eater's Belly (Ruin)
503: Zone Eater's Belly
504: Zone Eater's Belly
505: Zone Eater's Belly (Ruin)
506: Zone Eater's Belly
507: Zone Eater's Belly (Ruin)
508: Cave to the Ancient Castle (Ruin)
509: Cave to the Ancient Castle (Ruin)
510: Cave to the Ancient Castle (Ruin)
511: Cave to the Ancient Castle
512: Ancient Castle (Ruin)
513: Ancient Castle (Ruin)
514: Ancient Castle (Ruin)
515: Ancient Castle (Ruin)
516: Ancient Castle (Ruin)
517: Ancient Castle (Ruin)
518: Ancient Castle
519: Ancient Castle
520: Dreamscape (Ruin)
521: Dreamscape (Ruin)
522: Dreamscape - Phantom Train
523: Dreamscape - Phantom Train (Ruin)
524: Dreamscape - Phantom Train (Ruin)
525: Dreamscape - Phantom Train (Ruin)
526: Dreamscape - Phantom Train
527: Dreamscape - Phantom Train
528: Dreamscape - Cave (Ruin)
529: Dreamscape - Cave (Ruin)
530: Dreamscape - Doma Castle (Ruin)
531: Dreamscape - Doma Castle (Ruin)
532: Dreamscape - Doma Castle (Ruin)
533: Dreamscape - Doma Castle (Ruin)
534: Dreamscape - Doma Castle (Ruin)
535: Dreamscape - Doma Castle (Ruin)
536: Dreamscape - Doma Castle (Ruin)
537: Dreamscape - Doma Castle (Ruin)
538: Dreamscape - Doma Castle (Ruin)
539: Ebot's Rock (Ruin)
540: Ebot's Rock (Ruin)
541: Ebot's Rock (Ruin)
542: Kefka's Tower (Ruin)
543: Kefka's Tower (Ruin)
544: Kefka's Tower (Ruin)
545: Kefka's Tower
546: Kefka's Tower (Ruin)
547: Kefka's Tower (Ruin)
548: Kefka's Tower (Ruin)
549: Kefka's Tower
550: Kefka's Tower (Ruin)
551: Kefka's Tower (Ruin)
552: Kefka's Tower (Ruin)
553: Kefka's Tower (Ruin)
554: Kefka's Tower (Ruin)
555: Kefka's Tower (Ruin)
556: Kefka's Tower (Ruin)
557: Kefka's Tower (Ruin)
558: Kefka's Tower (Ruin)
559: Kefka's Tower (Ruin)
560: Kefka's Tower (Ruin)
561: Kefka's Tower (Ruin)
562: Kefka's Tower (Ruin)
563: Kefka's Tower (Ruin)
564: Kefka's Tower (Ruin)
565: Kefka's Tower (Ruin)
566: Kefka's Tower (Ruin)
567: Kefka's Tower
568: Kefka's Tower
569: Kefka's Tower
570: Kefka's Tower
571: Kefka's Tower
572: Overworld
573: Overworld
574: Opera House
575: Overworld
576: Opera House
577: Darill's Tomb
578: Overworld
579: Darill's Tomb
580: Narshe Mines
581: Narshe
582: Blackjack
583: Mobliz
584: Scenario Selection
585: Falcon
586: Overworld
587: Figaro Castle
588: Figaro Castle
589: Figaro Castle
590: Kefka's Tower
591: Scenario Selection
592: Scenario Selection
593: Scenario Selection
594: Scenario Selection
595: Scenario Selection
596: Scenario Selection
597: Solitary Island
598: Overworld
599: Overworld
600: Overworld
601: Scenario Selection
602: Scenario Selection
603: Scenario Selection
604: Esper World
605: Scenario Selection
606: 
607: 
608: 
609: Narshe
610: Jidoor
611: Jidoor
612: Chocobo Stable
613: Chocobo Stable
614: Chocobo Stable
615: Chocobo Stable
616: Chocobo Stable
617: Chocobo Stable
618: Chocobo Stable
619: Chocobo Stable
620: Chocobo Stable
621: Chocobo Stable
622: Nikeah
623: Albrook
624: Phantom Train (Balance)
625: Magitek Research Facility
626: South Figaro Cave (Balance)
627: South Figaro Cave (Balance)
628: South Figaro Cave (Balance)
629: South Figaro Cave (Balance)
630: South Figaro Cave (Balance)
631: Floating Continent (Balance)
632: Narshe Mines (Ruin)
633: Narshe Mines (Ruin)
634: Narshe Mines (Ruin)
635: Narshe Mines (Ruin)
636: Narshe Mines (Ruin)
637: Narshe Mines (Ruin)
638: Narshe Mines (Ruin)
639: Narshe Mines (Ruin)
640: 
641: Floating Continent
642: Mt. Kolts
643: Magitek Factory (Balance)
800: Overworld
801: Overworld
802: Overworld
803: Overworld
804: Overworld
805: Overworld
`
)
