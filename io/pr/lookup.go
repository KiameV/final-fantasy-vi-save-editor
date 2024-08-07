package pr

var (
	AllNormalItems = map[int]string{
		2:   "Potion",
		3:   "Hi-Potion",
		4:   "X-Potion",
		5:   "Ether",
		6:   "Hi-Ether",
		7:   "X-Ether",
		8:   "Elixir",
		9:   "Megalixir",
		10:  "Fenix Down",
		11:  "Holy Water",
		12:  "Antidote",
		13:  "Eye Drops",
		14:  "Golden Needle",
		15:  "Remedy",
		16:  "Sleeping Bag",
		17:  "Tent",
		18:  "Green Cherry",
		19:  "Magicite Shard",
		20:  "Super Ball",
		21:  "Echo Screen",
		22:  "Smoke Bomb",
		23:  "Teleport Stone",
		24:  "Dried Meat",
		25:  "Rename Card",
		31:  "Noiseblaster",
		32:  "Bioblaster",
		33:  "Flash",
		34:  "ChainSaw",
		35:  "Debilitator",
		36:  "Drill",
		37:  "Air Anchor",
		38:  "Auto Crossbow",
		26:  "Fire Scroll",
		27:  "Water Scroll",
		28:  "Lightning Scroll",
		29:  "Invisibility Scroll",
		30:  "Shadow Scroll",
		94:  "Dagger",
		95:  "Mythril Knife",
		96:  "Main Gauche",
		97:  "Air Knife",
		98:  "Thief's Knife",
		99:  "Assassin's Dagger",
		100: "Man-Eater",
		101: "Swordbreaker",
		102: "Gladius",
		103: "Valiant Knife",
		104: "Mythril Sword",
		105: "Great Sword",
		106: "Rune Blade",
		107: "Flametongue",
		108: "Icebrand",
		109: "Thunder Blade",
		110: "Bastard Sword ",
		111: "Break Blade",
		112: "Blood Sword",
		113: "Enhancer",
		114: "Crystal Sword",
		115: "Falchion",
		116: "Soul Sabre",
		117: "Organyx",
		118: "Excalibur",
		119: "Zantetsuken",
		120: "Lightbringer",
		121: "Ragnarok",
		122: "Ultima Weapon",
		123: "Mythril Spear",
		124: "Trident",
		125: "Heavy Lance",
		126: "Partisan",
		127: "Holy Lance",
		128: "Golden Lance",
		129: "Radiant Lance",
		130: "Impartisan",
		131: "Kunai",
		132: "Kodachi",
		133: "Sakura",
		134: "Sasuke",
		135: "Ichigeki",
		136: "Kagenui",
		137: "Ashura",
		138: "Kotetsu",
		139: "Kikuichimonji",
		140: "Kazekiri",
		141: "Murasame",
		142: "Masamune",
		143: "Murakumo",
		144: "Mutsunokami",
		145: "Heal Rod",
		146: "Mythril Rod",
		147: "Flame Rod",
		148: "Ice Rod",
		149: "Thunder Rod",
		150: "Poison Rod",
		151: "Holy Rod",
		152: "Gravity Rod",
		153: "Punisher",
		154: "Magus Rod",
		155: "Chocobo Brush",
		156: "Da Vinci Brush",
		157: "Magical Brush",
		158: "Rainbow Brush",
		159: "Shuriken",
		160: "Fuma Shuriken",
		161: "Pinwheel",
		162: "Chain Flail",
		163: "Moonring Blade",
		164: "Morning Star",
		165: "Boomerang",
		166: "Rising Sun",
		167: "Hawkeye",
		168: "Bone Club",
		169: "Sniper",
		170: "Wing Edge",
		171: "Cards",
		172: "Darts",
		173: "Death Tarot",
		174: "Viper Darts",
		175: "Dice",
		176: "Fixed Dice",
		177: "Metal Knuckles",
		178: "Mythril Claws",
		179: "Kaiser Knuckles",
		180: "Venom Claws",
		181: "Burning Fist",
		182: "Dragon Claws",
		183: "Tigerfangs",
		201: "Buckler",
		202: "Heavy Shield",
		203: "Mythril Shield",
		204: "Gold Shield",
		205: "Aegis Shield",
		206: "Diamond Shield",
		207: "Flame Shield",
		208: "Ice Shield",
		209: "Thunder Shield",
		210: "Crystal Shield",
		211: "Genji Shield",
		212: "Tortoise Shield",
		213: "Cursed Shield",
		214: "Paladin Shield",
		215: "Force Shield",
		216: "Leather Hat",
		217: "Hairband",
		218: "Plumed Hat",
		219: "Beret",
		220: "Magus Hat",
		221: "Bandana",
		222: "Iron Helmet",
		223: "Priest's Miter",
		224: "Bard's Hat",
		225: "Green Beret",
		226: "Head Band",
		227: "Mythril Helm",
		228: "Tiara",
		229: "Gold Helmet",
		230: "Tiger Mask",
		231: "Red Hat",
		232: "Mystery Veil",
		233: "Circlet",
		234: "Royal Crown",
		235: "Diamond Helm",
		236: "Black Hood",
		237: "Crystal Helm",
		238: "Oath Veil",
		239: "Cat-Ear Hood",
		240: "Genji Helmet",
		241: "Thornlet",
		242: "Saucer",
		244: "Leather Armor",
		245: "Cotton Robe",
		246: "Kenpo Gi",
		247: "Iron Armor",
		248: "Silk Robe",
		249: "Mythril Vest",
		250: "Ninja Gear",
		251: "White Dress",
		252: "Mythril Mail",
		253: "Gaia Gear",
		254: "Mirage Dress",
		255: "Golden Armor",
		256: "Power Sash",
		257: "Luminous Robe",
		258: "Diamond Vest",
		259: "Red Jacket",
		260: "Force Armor",
		261: "Diamond Armor",
		262: "Black Garb",
		263: "Magus Rove",
		264: "Crystal Mail",
		265: "Regal Gown",
		266: "Genji Armor",
		267: "Reed Cloak",
		268: "Minerva Bustier",
		269: "Tabby Suit",
		270: "Chocobo Suit",
		271: "Moogle Suit",
		272: "Nutkin Suit",
		273: "Behemeth Suit",
		274: "Snow Scarf",
		275: "Silver Spectacles",
		276: "Star Pendant",
		277: "Peace Ring",
		278: "Amulet",
		279: "White Cape",
		280: "Jewel Ring",
		281: "Fairy Ring",
		282: "Barrier Ring",
		283: "Mythril Glove",
		284: "Protect Ring",
		285: "Hermes Sandals",
		286: "Reflect Ring",
		287: "Angel Wings",
		288: "Angel Ring",
		289: "Knight's Code",
		290: "Dragoon Boots",
		291: "Zephyr Cloak",
		292: "Princess Ring",
		293: "Cursed Ring",
		294: "Earring",
		295: "Gigas Glove",
		296: "Blizzard Orb",
		297: "Berserker Ring",
		298: "Thief's Bracer",
		299: "Guard Bracelet",
		300: "Hero Ring",
		301: "Ribbon",
		302: "Muscle Belt",
		303: "Crystal Orb",
		304: "Gold Hairpin",
		305: "Celestriad",
		306: "Brigand's Glove",
		307: "Gauntlet",
		308: "Genji Glove",
		309: "Hyper Wrist",
		310: "Master's Scroll",
		311: "Prayer Beads",
		312: "Black Belt",
		313: "Heiki's Jitte",
		314: "Fake Mustache",
		315: "Soul of Thamasa",
		316: "Dragon Horn",
		317: "Merit Award",
		318: "Momento Ring",
		319: "Safety Bit",
		320: "Lich Ring",
		321: "Molulu's Charm",
		322: "Ward Bangle",
		323: "Miracle Shoes",
		324: "Alarm Gaurd",
		325: "Gale Hairpin",
		326: "Sniper Eye",
		327: "Growth Egg",
		328: "Tintinnabulum",
		329: "Sprint Shoes",
		0:   "Invalid",
		93:  "[Empty Weapon/Shield]",
		199: "[Empty Head]",
		198: "[Empty Armor]",
		200: "[Empty Relic]",
	}
)
