using FF6_Save_Editor.Util;
using System;
using FF6_Save_State_Editor.Enums;
using FF6_Save_Editor.Enums;

namespace FF6_Save_Editor.DTOs
{
    /// <summary>
    /// Contains the data from the loaded save slot/save state
    /// </summary>
    class SaveDto
    {
        /// <summary>
        /// The Character-specific data
        /// </summary>
        public readonly CharacterContainer Characters;
        /// <summary>
        /// The current item inventory
        /// </summary>
        public readonly InventoryDto Inventory;
        /// <summary>
        /// The known skills including Blitzes, Lore, SwdTechs, Rages, and Dances
        /// </summary>
        public readonly SkillsDto Skills;
        /// <summary>
        /// Misc global stats like player airship locaiton, player location, save count, GP and steps 
        /// </summary>
        public readonly OtherStatsDto OtherStats;
        /// <summary>
        /// All aquired Espers
        /// </summary>
        public readonly KnownSkillsDto<EsperEnum> KnownEspers;

        /// <summary>
        /// The type of file that is opened
        /// </summary>
        public SaveFileTypeEnum SaveFileType { get; private set; }

        /// <summary>
        /// Byte array pf tje loaded save slot/save state
        /// </summary>
        public Byte[] SaveByteArray { get; private set; }

        public SaveDto()
        {
            this.SaveByteArray = new Byte[0];

            this.Characters = new CharacterContainer();
            this.Inventory = new InventoryDto();
            this.Skills = new SkillsDto();
            this.OtherStats = new OtherStatsDto();
            this.KnownEspers = new KnownSkillsDto<EsperEnum>(typeof(EsperEnum));
        }

        /// <summary>
        /// Set the save file's data and type
        /// </summary>
        /// <param name="saveFileType">The type of file which was opened</param>
        /// <param name="saveByteArray">The file's data</param>
        internal void SetSaveState(SaveFileTypeEnum saveFileType, byte[] saveByteArray)
        {
            this.SaveFileType = saveFileType;
            this.SaveByteArray = saveByteArray;
        }
    }
}
