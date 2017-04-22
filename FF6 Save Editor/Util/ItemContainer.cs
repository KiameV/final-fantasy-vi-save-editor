using System.Collections.Generic;

namespace FF6_Save_State_Editor.Util
{
    /// <summary>
    /// Interface for finding Items
    /// </summary>
    public static class ItemContainer
    {
        /// <summary>
        /// Defines an item's name and it's id
        /// </summary>
        private class ItemDto
        {
            /// <summary>
            /// The name of the item
            /// </summary>
            public readonly string Name;
            /// <summary>
            /// The id of the item
            /// </summary>
            public readonly byte ItemId;

            internal ItemDto(byte itemId, string name)
            {
                ItemId = itemId;
                Name = name;
            }
        }

        /// <summary>
        /// List of all items
        /// </summary>
        private static List<ItemDto> items = new List<ItemDto>(255);

        /// <summary>
        /// Get the item defined by the given id
        /// </summary>
        /// <param name="id">The id of the item to find</param>
        /// <returns>The item which corresponds to the given id</returns>
        public static string GetMatchFor(byte id)
        {
            // The Item list will have 255 values so no check is needed for out of bounds values
            return items[id].Name;
        }

        /// <summary>
        /// Finds an item which starts with the given string.
        /// </summary>
        /// <param name="s">The string to use to find items that start with it</param>
        /// <returns>If more than one Item is found an empty string is returned. If no items are found, "none" is returned. If only one item starts with the given string, that item's name is returned</returns>
        public static string GetMatchFor(string s)
        {
            string lower = s.ToLower();
            ItemDto found = null;
            foreach (ItemDto i in items)
            {
                if (i.Name.ToLower().StartsWith(lower))
                {
                    if (found != null)
                    {
                        return "";
                    }
                    found = i;
                }
            }

            if (found == null)
                return "None";
            return found.ItemId.ToString("X2");
        }

        /// <summary>
        /// Create an ItemDTO for each item and add it to the items list
        /// </summary>
        static ItemContainer()
        {
            items.Add(new ItemDto(0x00, "Dirk"));
            items.Add(new ItemDto(0x01, "MithrilKnife"));
            items.Add(new ItemDto(0x02, "Guardian"));
            items.Add(new ItemDto(0x03, "Air Lancet"));
            items.Add(new ItemDto(0x04, "ThiefKnife"));
            items.Add(new ItemDto(0x05, "Assassin"));
            items.Add(new ItemDto(0x06, "Man Eater"));
            items.Add(new ItemDto(0x07, "SwordBreaker"));
            items.Add(new ItemDto(0x08, "Graedus"));
            items.Add(new ItemDto(0x09, "ValiantKnife"));
            items.Add(new ItemDto(0x0A, "MithrilBlade"));
            items.Add(new ItemDto(0x0B, "RegalCutlass"));
            items.Add(new ItemDto(0x0C, "Rune Edge"));
            items.Add(new ItemDto(0x0D, "Flame Sabre"));
            items.Add(new ItemDto(0x0E, "Blizzard"));
            items.Add(new ItemDto(0x0F, "ThunderBlade"));
            items.Add(new ItemDto(0x10, "Epee"));
            items.Add(new ItemDto(0x11, "Break Blade"));
            items.Add(new ItemDto(0x12, "Drainer"));
            items.Add(new ItemDto(0x13, "Enhancer"));
            items.Add(new ItemDto(0x14, "Crystal"));
            items.Add(new ItemDto(0x15, "Falchion"));
            items.Add(new ItemDto(0x16, "Soul Sabre"));
            items.Add(new ItemDto(0x17, "Ogre Nix"));
            items.Add(new ItemDto(0x18, "Excalibur"));
            items.Add(new ItemDto(0x19, "Scimiter"));
            items.Add(new ItemDto(0x1A, "Illumina"));
            items.Add(new ItemDto(0x1B, "Ragnarok"));
            items.Add(new ItemDto(0x1C, "Atma Weapon"));
            items.Add(new ItemDto(0x1D, "Mithril Pike"));
            items.Add(new ItemDto(0x1E, "Trident"));
            items.Add(new ItemDto(0x1F, "Stout Spear"));
            items.Add(new ItemDto(0x20, "Partisan"));
            items.Add(new ItemDto(0x21, "Pearl Lance"));
            items.Add(new ItemDto(0x22, "Gold Lance"));
            items.Add(new ItemDto(0x23, "Aura Lance"));
            items.Add(new ItemDto(0x24, "Imp Halberd"));
            items.Add(new ItemDto(0x25, "Imperial"));
            items.Add(new ItemDto(0x26, "Kodachi"));
            items.Add(new ItemDto(0x27, "Blossom"));
            items.Add(new ItemDto(0x28, "Hardened"));
            items.Add(new ItemDto(0x29, "Striker"));
            items.Add(new ItemDto(0x2A, "Stunner"));
            items.Add(new ItemDto(0x2B, "Ashura"));
            items.Add(new ItemDto(0x2C, "Kotetsu"));
            items.Add(new ItemDto(0x2D, "Forged"));
            items.Add(new ItemDto(0x2E, "Tempest"));
            items.Add(new ItemDto(0x2F, "Murasame"));
            items.Add(new ItemDto(0x30, "Aura"));
            items.Add(new ItemDto(0x31, "Strato"));
            items.Add(new ItemDto(0x32, "Sky Render"));
            items.Add(new ItemDto(0x33, "Heal Rod"));
            items.Add(new ItemDto(0x34, "Mithril Rod"));
            items.Add(new ItemDto(0x35, "Fire Rod"));
            items.Add(new ItemDto(0x36, "Ice Rod"));
            items.Add(new ItemDto(0x37, "Thunder Rod"));
            items.Add(new ItemDto(0x38, "Poison Rod"));
            items.Add(new ItemDto(0x39, "Pearl Rod"));
            items.Add(new ItemDto(0x3A, "Gravity Rod"));
            items.Add(new ItemDto(0x3B, "Punisher"));
            items.Add(new ItemDto(0x3C, "Magus Rod"));
            items.Add(new ItemDto(0x3D, "Chocobo Brsh"));
            items.Add(new ItemDto(0x3E, "DaVinci Brsh"));
            items.Add(new ItemDto(0x3F, "Magical Brsh"));
            items.Add(new ItemDto(0x40, "Rainbow Brsh"));
            items.Add(new ItemDto(0x41, "Shuriken"));
            items.Add(new ItemDto(0x42, "Ninja Star"));
            items.Add(new ItemDto(0x43, "Tack Star"));
            items.Add(new ItemDto(0x44, "Flail"));
            items.Add(new ItemDto(0x45, "Full Moon"));
            items.Add(new ItemDto(0x46, "Morning Star"));
            items.Add(new ItemDto(0x47, "Boomerang"));
            items.Add(new ItemDto(0x48, "Rising Sun"));
            items.Add(new ItemDto(0x49, "Hawk Eye"));
            items.Add(new ItemDto(0x4A, "Bone Club"));
            items.Add(new ItemDto(0x4B, "Sniper"));
            items.Add(new ItemDto(0x4C, "Wing Edge"));
            items.Add(new ItemDto(0x4D, "Cards"));
            items.Add(new ItemDto(0x4E, "Darts"));
            items.Add(new ItemDto(0x4F, "Doom Darts"));
            items.Add(new ItemDto(0x50, "Trump"));
            items.Add(new ItemDto(0x51, "Dice"));
            items.Add(new ItemDto(0x52, "Fixed Dice"));
            items.Add(new ItemDto(0x53, "MetalKnuckle"));
            items.Add(new ItemDto(0x54, "Mithril Claw"));
            items.Add(new ItemDto(0x55, "Kaiser"));
            items.Add(new ItemDto(0x56, "Poison Claw"));
            items.Add(new ItemDto(0x57, "Fire Knuckle"));
            items.Add(new ItemDto(0x58, "Dragon Claw"));
            items.Add(new ItemDto(0x59, "Tiger Fangs"));
            items.Add(new ItemDto(0x5A, "Buckler"));
            items.Add(new ItemDto(0x5B, "Heavy Shld"));
            items.Add(new ItemDto(0x5C, "Mithril Shld"));
            items.Add(new ItemDto(0x5D, "Gold Shld"));
            items.Add(new ItemDto(0x5E, "Aegis Shld"));
            items.Add(new ItemDto(0x5F, "Diamond Shld"));
            items.Add(new ItemDto(0x60, "Flame Shld"));
            items.Add(new ItemDto(0x61, "Ice Shld"));
            items.Add(new ItemDto(0x62, "Thunder Shld"));
            items.Add(new ItemDto(0x63, "Crystal Shld"));
            items.Add(new ItemDto(0x64, "Genji Shld"));
            items.Add(new ItemDto(0x65, "TortoiseShld"));
            items.Add(new ItemDto(0x66, "Cursed Shld"));
            items.Add(new ItemDto(0x67, "Paladin Shld"));
            items.Add(new ItemDto(0x68, "Force Shld"));
            items.Add(new ItemDto(0x69, "Leather Hat"));
            items.Add(new ItemDto(0x6A, "Hair Band"));
            items.Add(new ItemDto(0x6B, "Plumed Hat"));
            items.Add(new ItemDto(0x6C, "Beret"));
            items.Add(new ItemDto(0x6D, "Magus Hat"));
            items.Add(new ItemDto(0x6E, "Bandana"));
            items.Add(new ItemDto(0x6F, "Iron Helmet"));
            items.Add(new ItemDto(0x70, "Coronet"));
            items.Add(new ItemDto(0x71, "Bard's Hat"));
            items.Add(new ItemDto(0x72, "Green Beret"));
            items.Add(new ItemDto(0x73, "Head Band"));
            items.Add(new ItemDto(0x74, "Mithril Helm"));
            items.Add(new ItemDto(0x75, "Tiara"));
            items.Add(new ItemDto(0x76, "Gold Helmet"));
            items.Add(new ItemDto(0x77, "Tiger Mask"));
            items.Add(new ItemDto(0x78, "Red Hat"));
            items.Add(new ItemDto(0x79, "Mystery Veil"));
            items.Add(new ItemDto(0x7A, "Circlet"));
            items.Add(new ItemDto(0x7B, "Regal Crown"));
            items.Add(new ItemDto(0x7C, "Diamond Helm"));
            items.Add(new ItemDto(0x7D, "Dark Hood"));
            items.Add(new ItemDto(0x7E, "Crystal Helm"));
            items.Add(new ItemDto(0x7F, "Oath Veil"));
            items.Add(new ItemDto(0x80, "Cat Hood"));
            items.Add(new ItemDto(0x81, "Genji Helmet"));
            items.Add(new ItemDto(0x82, "Thornlet"));
            items.Add(new ItemDto(0x83, "Titanium"));
            items.Add(new ItemDto(0x84, "LeatherArmor"));
            items.Add(new ItemDto(0x85, "Cotton Robe"));
            items.Add(new ItemDto(0x86, "Kung Fu Suit"));
            items.Add(new ItemDto(0x87, "Iron Armor"));
            items.Add(new ItemDto(0x88, "Silk Robe"));
            items.Add(new ItemDto(0x89, "Mithril Vest"));
            items.Add(new ItemDto(0x8A, "Ninja Gear"));
            items.Add(new ItemDto(0x8B, "White Dress"));
            items.Add(new ItemDto(0x8C, "Mithril Mail"));
            items.Add(new ItemDto(0x8D, "Gaia Gear"));
            items.Add(new ItemDto(0x8E, "Mirage Dress"));
            items.Add(new ItemDto(0x8F, "Gold Armor"));
            items.Add(new ItemDto(0x90, "Power Sash"));
            items.Add(new ItemDto(0x91, "Light Robe"));
            items.Add(new ItemDto(0x92, "Diamond Vest"));
            items.Add(new ItemDto(0x93, "Red Jacket"));
            items.Add(new ItemDto(0x94, "Force Armor"));
            items.Add(new ItemDto(0x95, "DiamondArmor"));
            items.Add(new ItemDto(0x96, "Dark Gear"));
            items.Add(new ItemDto(0x97, "Tao Robe"));
            items.Add(new ItemDto(0x98, "Crystal Mail"));
            items.Add(new ItemDto(0x99, "Czarina Gown"));
            items.Add(new ItemDto(0x9A, "Genji Armor"));
            items.Add(new ItemDto(0x9B, "Imp's Armor"));
            items.Add(new ItemDto(0x9C, "Minerva"));
            items.Add(new ItemDto(0x9D, "Tabby Suit"));
            items.Add(new ItemDto(0x9E, "Chocobo Suit"));
            items.Add(new ItemDto(0x9F, "Moogle Suit"));
            items.Add(new ItemDto(0xA0, "Nutkin Suit"));
            items.Add(new ItemDto(0xA1, "BehemethSuit"));
            items.Add(new ItemDto(0xA2, "Snow Muffler"));
            items.Add(new ItemDto(0xA3, "NoiseBlaster"));
            items.Add(new ItemDto(0xA4, "Bio Blaster"));
            items.Add(new ItemDto(0xA5, "Flash"));
            items.Add(new ItemDto(0xA6, "Chain Saw"));
            items.Add(new ItemDto(0xA7, "Debilitator"));
            items.Add(new ItemDto(0xA8, "Drill"));
            items.Add(new ItemDto(0xA9, "Air Anchor"));
            items.Add(new ItemDto(0xAA, "AutoCrossbow"));
            items.Add(new ItemDto(0xAB, "Fire Skean"));
            items.Add(new ItemDto(0xAC, "Water Edge"));
            items.Add(new ItemDto(0xAD, "Bolt Edge"));
            items.Add(new ItemDto(0xAE, "Inviz Edge"));
            items.Add(new ItemDto(0xAF, "Shadow Edge"));
            items.Add(new ItemDto(0xB0, "Goggles"));
            items.Add(new ItemDto(0xB1, "Star Pendant"));
            items.Add(new ItemDto(0xB2, "Peace Ring"));
            items.Add(new ItemDto(0xB3, "Amulet"));
            items.Add(new ItemDto(0xB4, "White Cape"));
            items.Add(new ItemDto(0xB5, "Jewel Ring"));
            items.Add(new ItemDto(0xB6, "Fair Ring"));
            items.Add(new ItemDto(0xB7, "Barrier Ring"));
            items.Add(new ItemDto(0xB8, "MithrilGlove"));
            items.Add(new ItemDto(0xB9, "Guard Ring"));
            items.Add(new ItemDto(0xBA, "RunningShoes"));
            items.Add(new ItemDto(0xBB, "Wall Ring"));
            items.Add(new ItemDto(0xBC, "Cherub Down"));
            items.Add(new ItemDto(0xBD, "Cure Ring"));
            items.Add(new ItemDto(0xBE, "True Knight"));
            items.Add(new ItemDto(0xBF, "DragoonBoots"));
            items.Add(new ItemDto(0xC0, "Zephyr Cape"));
            items.Add(new ItemDto(0xC1, "Czarina Ring"));
            items.Add(new ItemDto(0xC2, "Cursed Cing"));
            items.Add(new ItemDto(0xC3, "Earrings"));
            items.Add(new ItemDto(0xC4, "Atlas Armlet"));
            items.Add(new ItemDto(0xC5, "BlizzardRing"));
            items.Add(new ItemDto(0xC6, "Rage Ring"));
            items.Add(new ItemDto(0xC7, "Sneak Ring"));
            items.Add(new ItemDto(0xC8, "Pod Bracelet"));
            items.Add(new ItemDto(0xC9, "Hero Ring"));
            items.Add(new ItemDto(0xCA, "Ribbon"));
            items.Add(new ItemDto(0xCB, "Muscle Belt"));
            items.Add(new ItemDto(0xCC, "Crystal Orb"));
            items.Add(new ItemDto(0xCD, "Gold Hairpin"));
            items.Add(new ItemDto(0xCE, "Economizer"));
            items.Add(new ItemDto(0xCF, "Thief Glove"));
            items.Add(new ItemDto(0xD0, "Gauntlet"));
            items.Add(new ItemDto(0xD1, "Genji Glove"));
            items.Add(new ItemDto(0xD2, "Hyper Wrist"));
            items.Add(new ItemDto(0xD3, "Offering"));
            items.Add(new ItemDto(0xD4, "Beads"));
            items.Add(new ItemDto(0xD5, "Black Belt"));
            items.Add(new ItemDto(0xD6, "Coin Toss"));
            items.Add(new ItemDto(0xD7, "FakeMustache"));
            items.Add(new ItemDto(0xD8, "Gem Box"));
            items.Add(new ItemDto(0xD9, "Dragon Horn"));
            items.Add(new ItemDto(0xDA, "Merit Award"));
            items.Add(new ItemDto(0xDB, "Momento Ring"));
            items.Add(new ItemDto(0xDC, "Safety Bit"));
            items.Add(new ItemDto(0xDD, "Relic Ring"));
            items.Add(new ItemDto(0xDE, "Moogle Charm"));
            items.Add(new ItemDto(0xDF, "Charm Bangle"));
            items.Add(new ItemDto(0xE0, "Marvel Shoes"));
            items.Add(new ItemDto(0xE1, "Back Gaurd"));
            items.Add(new ItemDto(0xE2, "Gale Hairpin"));
            items.Add(new ItemDto(0xE3, "Sniper Sight"));
            items.Add(new ItemDto(0xE4, "Exp. Egg"));
            items.Add(new ItemDto(0xE5, "Tintinabar"));
            items.Add(new ItemDto(0xE6, "Sprint Shoes"));
            items.Add(new ItemDto(0xE7, "Rename Card"));
            items.Add(new ItemDto(0xE8, "Tonic"));
            items.Add(new ItemDto(0xE9, "Potion"));
            items.Add(new ItemDto(0xEA, "X-Potion"));
            items.Add(new ItemDto(0xEB, "Tincture"));
            items.Add(new ItemDto(0xEC, "Ether"));
            items.Add(new ItemDto(0xED, "X-Ether"));
            items.Add(new ItemDto(0xEE, "Elixir"));
            items.Add(new ItemDto(0xEF, "Megalixir"));
            items.Add(new ItemDto(0xF0, "Fenix Down"));
            items.Add(new ItemDto(0xF1, "Revivify"));
            items.Add(new ItemDto(0xF2, "Antidote"));
            items.Add(new ItemDto(0xF3, "Eydrop"));
            items.Add(new ItemDto(0xF4, "Soft"));
            items.Add(new ItemDto(0xF5, "Remedy"));
            items.Add(new ItemDto(0xF6, "Sleeping Bag"));
            items.Add(new ItemDto(0xF7, "Tent"));
            items.Add(new ItemDto(0xF8, "Green Cherry"));
            items.Add(new ItemDto(0xF9, "Magicite"));
            items.Add(new ItemDto(0xFA, "Super Ball"));
            items.Add(new ItemDto(0xFB, "Echo Screen"));
            items.Add(new ItemDto(0xFC, "Smoke Bomb"));
            items.Add(new ItemDto(0xFD, "Warp Stone"));
            items.Add(new ItemDto(0xFE, "Dried Meat"));
            items.Add(new ItemDto(0xFF, "[Nothing]"));
        }
    }
}
