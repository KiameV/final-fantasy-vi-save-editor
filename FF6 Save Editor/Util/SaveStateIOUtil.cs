using FF6_Save_Editor.DTOs;
using FF6_Save_Editor.Enums;
using FF6_Save_State_Editor.Enums;
using FF6_Save_State_Editor.Util;
using System;
using System.Collections;
using System.Collections.Generic;
using System.Text;
using System.Windows.Forms;
using static FF6_Save_State_Editor.Util.EsperContainer;
using static FF6_Save_State_Editor.Util.OffsetFactory;
using static FF6_Save_State_Editor.Util.RagesContainer;
using Microsoft.VisualBasic;

namespace FF6_Save_Editor.Util
{
    /// <summary>
    /// Utility class for sync'ing SaveDataDto information with the save's byte array ad vice versa
    /// </summary>
    class SaveStateIOUtil
    {
        /// <summary>
        /// Utility method for sync'ing the SaveStateDto's fields with the save byte array's data
        /// </summary>
        /// <param name="saveStateDto">SaveStateDto which contains the fields and the save's byte array to sync</param>
        public static void LoadSaveStateInformation(SaveDto saveStateDto)
        {
            Offsets offsets = null;
            if (saveStateDto.SaveFileType == SaveFileTypeEnum.Snes9xSaveState16)
            {
                if (!TryFindDynamicOffset(saveStateDto, out int offset))
                    return;

                offsets = OffsetFactory.CreateOffsets(offset);
            }
            else {
                offsets = OffsetFactory.CreateOffsets(saveStateDto.SaveFileType);
            }

            try
            {
                LoadCharacters(saveStateDto.Characters, saveStateDto.SaveByteArray, offsets);
                LoadInventory(saveStateDto.Inventory, saveStateDto.SaveByteArray, offsets);
                LoadSkills(saveStateDto.Skills, saveStateDto.SaveByteArray, offsets);
                LoadEspers(saveStateDto.KnownEspers, saveStateDto.SaveByteArray, offsets);
                LoadOtherStats(saveStateDto.OtherStats, saveStateDto.SaveByteArray, offsets);
            }
            catch
            {
                MessageBox.Show("Failed to load.");
            }
        }

        /// <summary>
        /// Utility method for sync'ing the save byte array's data with the SaveStateDto's fields
        /// </summary>
        /// <param name="saveStateDto">SaveStateDto which contains the save's byte array and the fields to sync</param>
        public static void SaveSaveStateInformation(SaveDto saveStateDto)
        {
            Offsets offsets = null;
            if (saveStateDto.SaveFileType == SaveFileTypeEnum.Snes9xSaveState16)
            {
                if (!TryFindDynamicOffset(saveStateDto, out int offset, saveStateDto.Characters.getCharacter(CharacterEnum.Terra).Name))
                    return;

                offsets = OffsetFactory.CreateOffsets(offset);
            }
            else
            {
                offsets = OffsetFactory.CreateOffsets(saveStateDto.SaveFileType);
            }

            try
            {
                SaveCharacters(saveStateDto.Characters, saveStateDto.SaveByteArray, offsets);
                SaveInventory(saveStateDto.Inventory, saveStateDto.SaveByteArray, offsets);
                SaveSkills(saveStateDto.Skills, saveStateDto.SaveByteArray, offsets);
                SaveEspers(saveStateDto.KnownEspers, saveStateDto.SaveByteArray, offsets);
                SaveOtherStats(saveStateDto.OtherStats, saveStateDto.SaveByteArray, offsets);

                if (saveStateDto.SaveFileType == SaveFileTypeEnum.SRMSlot1 ||
                    saveStateDto.SaveFileType == SaveFileTypeEnum.SRMSlot2 ||
                    saveStateDto.SaveFileType == SaveFileTypeEnum.SRMSlot3)
                {
                    int baseOffset = 0;
                    if (saveStateDto.SaveFileType == SaveFileTypeEnum.SRMSlot2)
                        baseOffset = OffsetFactory.SRM_SLOT_SIZE;
                    else if (saveStateDto.SaveFileType == SaveFileTypeEnum.SRMSlot3)
                        baseOffset = OffsetFactory.SRM_SLOT_SIZE * 2;

                    uint newchecksum = 0;

                    for (int i = baseOffset; i < baseOffset + OffsetFactory.SRM_SLOT_SIZE - 2; i++)
                    {
                        newchecksum += saveStateDto.SaveByteArray[i];
                        newchecksum &= 0xFFFF;
                    }

                    byte[] bytes = BitConverter.GetBytes(newchecksum);

                    saveStateDto.SaveByteArray[baseOffset + OffsetFactory.SRM_SLOT_SIZE - 2] = bytes[0];
                    saveStateDto.SaveByteArray[baseOffset + OffsetFactory.SRM_SLOT_SIZE - 1] = bytes[1];
                }
            }
            catch(Exception e)
            {
                MessageBox.Show($"Failed to save state.\n{e.Message}");
            }
        }

        /// <summary>
        /// Load the characters' data from the byte array
        /// </summary>
        /// <param name="characters">The CharactersContainer to populate with the saveState's data</param>
        /// <param name="saveState">The save's byte array</param>
        /// <param name="offsets">Contians the offsets for the data in the save's byte array</param>
        private static void LoadCharacters(CharacterContainer characters, Byte[] saveState, Offsets offsets)
        {
            foreach (CharacterEnum cEnum in Enum.GetValues(typeof(CharacterEnum)))
            {
                CharacterDto character = characters.getCharacter(cEnum);
                int i = ((int)cEnum) * NEXT_CHARACTER_OFFSET + offsets.CharacterOffset; // Character

                // Character Id
                i += 2;

                // Name
                StringBuilder name = new StringBuilder();
                for (uint n = 0; n < 6; ++n)
                {
                    byte b = saveState[i + n];
                    if (b == 0 || b == 255)
                    {
                        break;
                    }

                    if (b == 191)
                    {
                        name.Append('?');
                    }
                    else
                    {
                        name.Append((char)(b - 63));
                    }
                }
                i += 6;
                character.Name = name.ToString();

                // Level
                character.Level = saveState[i];
                ++i;

                // Current HP
                character.CurrentHP = ReverseAndCombine(saveState[i + 1], saveState[i]);
                i += 2;

                // Max HP
                character.MaxHP = ReverseAndCombine(saveState[i + 1], saveState[i]);
                i += 2;

                // Current MP
                character.CurrentMP = ReverseAndCombine(saveState[i + 1], saveState[i]);
                i += 2;

                // Max MP
                character.MaxMP = ReverseAndCombine(saveState[i + 1], saveState[i]);
                i += 2;

                // Exp
                character.Exp = ReverseAndCombine(saveState[i + 2], saveState[i + 1], saveState[i]);
                i += 3;

                // Status
                byte status = saveState[i];
                character.StatusEffetcs.Darkness = (status & (byte)StatusEnum.Darkness) != 0;
                character.StatusEffetcs.Imp = (status & (byte)StatusEnum.Imp) != 0;
                character.StatusEffetcs.Invisible = (status & (byte)StatusEnum.Invisible) != 0;
                character.StatusEffetcs.Magitek = (status & (byte)StatusEnum.Magitek) != 0;
                character.StatusEffetcs.Poison = (status & (byte)StatusEnum.Poison) != 0;
                character.StatusEffetcs.Stone = (status & (byte)StatusEnum.Stone) != 0;
                character.StatusEffetcs.Wounded = (status & (byte)StatusEnum.Wounded) != 0;
                character.StatusEffetcs.Zombie = (status & (byte)StatusEnum.Zombie) != 0;
                ++i;

                // Float
                character.StatusEffetcs.Float = (saveState[i] < 128) ? false : true;
                ++i;

                // Commands
                for (uint c = 0; c < 4; ++c)
                {
                    byte b = saveState[c + i];
                    CommandsEnum ce = (CommandsEnum)b;
                    if (b == 255)
                    {
                        ce = CommandsEnum.Unassigned;
                    }
                    character.SetCommand(c, ce);
                }
                i += 4;

                // Vigor
                character.Vigor = saveState[i];
                ++i;

                // Speed
                character.Speed = saveState[i];
                ++i;

                // Stamina
                character.Stamina = saveState[i];
                ++i;

                // Magic
                character.Magic = saveState[i];
                ++i;

                // Esper
                character.Esper = (EsperEnum)saveState[i];
                ++i;

                // Weapon
                character.Weapon = saveState[i];
                ++i;

                // Shield
                character.Shield = saveState[i];
                ++i;

                // Helmet
                character.Helmet = saveState[i];
                ++i;

                // Armor
                character.Armor = saveState[i];
                ++i;

                // Relic 1
                character.Relic1 = saveState[i];
                ++i;

                // Relic 2
                character.Relic2 = saveState[i];
                ++i;

                // Magic
                i = offsets.MagicOffset + NEXT_CHARACTER_MAGIC_OFFSET * (int)cEnum;
                for (uint m = 0; m < Enum.GetValues(typeof(MagicEnum)).Length; ++m)
                {
                    character.KnownMagic.SetPercentLearned(m, saveState[i + m]);
                }
            }
        }

        /// <summary>
        /// Load the skills data from the byte array
        /// Skills include SwdTech, Lore, Blitzes, Rage, and Dances
        /// </summary>
        /// <param name="skills">The SkillsDto to populate with the saveState's data</param>
        /// <param name="saveState">The save's byte array</param>
        /// <param name="offsets">Contians the offsets for the data in the save's byte array</param>
        private static void LoadSkills(SkillsDto skills, Byte[] saveState, Offsets offsets)
        {
            // SwdTech
            uint b = saveState[offsets.SwdTechOffset];
            DetermineKnownSkills(skills.KnownSwdTechs, Enum.GetValues(typeof(SwdTechEnum)), b);

            // Lore
            b = ReverseAndCombine(saveState[offsets.LoreOffset + 2], saveState[offsets.LoreOffset + 1], saveState[offsets.LoreOffset]);
            DetermineKnownSkills(skills.KnownLore, Enum.GetValues(typeof(LoreEnum)), b);

            // Blitz
            b = saveState[offsets.BlitzOffset];
            DetermineKnownSkills(skills.KnownBlitzes, Enum.GetValues(typeof(BlitzEnum)), b);

            // Rage
            for (int i = 0; i < RagesContainer.GroupCount; ++i)
            {
                b = saveState[i + offsets.RageOffset];
                DetermineKnownSkills(skills.KnownRages, RagesContainer.GetRageDtoGroupForOffset(i), b);
            }

            // Dance
            b = saveState[offsets.DanceOffset];
            DetermineKnownSkills(skills.KnownDances, Enum.GetValues(typeof(DanceEnum)), b);
        }

        /// <summary>
        /// Load the obtained espers data from the byte array
        /// </summary>
        /// <param name="knownEspers">The KnownSkillsDto-EsperEnum to populate with the saveState's data</param>
        /// <param name="saveState">The save's byte array</param>
        /// <param name="offsets">Contians the offsets for the data in the save's byte array</param>
        private static void LoadEspers(KnownSkillsDto<EsperEnum> knownEspers, Byte[] saveState, Offsets offsets)
        {
            // Espers
            for (int i = 0; i < EsperContainer.GroupCount; ++i)
            {
                uint b = saveState[i + offsets.EsperOffset];
                DetermineKnownSkills(knownEspers, EsperContainer.GetEsperGroupForOffset(i), b);
            }
        }

        /// <summary>
        /// Load the other data from the byte array
        /// Other data include GP, Steps, Number of Saves, Map X/Y Axis, and Airship X/Y Axis
        /// </summary>
        /// <param name="otherStats">The OtherStatsDto to populate with the saveState's data</param>
        /// <param name="saveState">The save's byte array</param>
        /// <param name="offsets">Contians the offsets for the data in the save's byte array</param>
        private static void LoadOtherStats(OtherStatsDto otherStats, Byte[] saveState, Offsets offsets)
        {
            // Gold
            otherStats.GP = ReverseAndCombine(saveState[offsets.GoldOffset + 2], saveState[offsets.GoldOffset + 1], saveState[offsets.GoldOffset]);

            // Steps
            otherStats.Steps = ReverseAndCombine(saveState[offsets.StepsOffset + 2], saveState[offsets.StepsOffset + 1], saveState[offsets.StepsOffset]);

            // Number of Saves
            otherStats.NumberOfSaves = saveState[offsets.NumberOfSaves];
            otherStats.SaveCountRollover = saveState[offsets.NumberOfSaves + 1];

            // Map X/Y Axis
            otherStats.MapX = saveState[offsets.MapXYOffset];
            otherStats.MapY = saveState[offsets.MapXYOffset + 1];

            // Airhisp X/Y Axis
            otherStats.AirshipXAxis = saveState[offsets.AirshipXYOffset];
            otherStats.AirshipYAxis = saveState[offsets.AirshipXYOffset + 1];

            // Airship Visible
            otherStats.IsAirshipVisible = (saveState[offsets.AirshipSettingsOffset] & (1 << 1)) != 0;
        }

        /// <summary>
        /// Load the inventory data from the byte array
        /// </summary>
        /// <param name="inventory">The InventoryDto to populate with the saveState's data</param>
        /// <param name="saveState">The save's byte array</param>
        /// <param name="offsets">Contians the offsets for the data in the save's byte array</param>
        private static void LoadInventory(InventoryDto inventory, Byte[] saveState, Offsets offsets)
        {
            for (int i = 0; i < 255; ++i)
            {
                byte itemId = saveState[i + offsets.InventoryItemIdOffset];
                byte count = saveState[i + offsets.InventoryItemCountOffset];
                inventory.SetItem(i, itemId, count);
            }
        }

        /// <summary>
        /// Save the character data to the save's byte array
        /// </summary>
        /// <param name="characters">The CharactersContainer to retrieve information from</param>
        /// <param name="saveState">The save's byte array</param>
        /// <param name="offsets">Contians the offsets for the data in the save's byte array</param>
        private static void SaveCharacters(CharacterContainer characters, byte[] saveState, Offsets offsets)
        {
            foreach (CharacterEnum cEnum in Enum.GetValues(typeof(CharacterEnum)))
            {
                CharacterDto character = characters.getCharacter(cEnum);

                // Character Offset
                int i = ((int)cEnum) * NEXT_CHARACTER_OFFSET + offsets.CharacterOffset;

                // Character Id
                i += 2;

                // Name
                StringBuilder name = new StringBuilder();
                byte[] chars = { 0, 0, 0, 0, 0, 0 };
                char[] nameChars = character.Name.ToCharArray();
                for (uint n = 0; n < nameChars.Length; ++n)
                {
                    if (nameChars[n] == '?')
                    {
                        chars[n] = 191;
                    }
                    else
                    {
                        int c = (int)(Convert.ToInt32(nameChars[n]) + 63);
                        if (c > 255)
                            c = 255;
                        if (c < 0)
                            c = 0;
                        chars[n] = (byte)c;
                    }
                }
                for (uint n = 0; n < 6; ++n, ++i)
                {
                    if (nameChars.Length > 0 && n >= nameChars.Length)
                    {
                        chars[n] = 0xFF;
                    }
                    saveState[i] = chars[n];
                }

                // Level
                saveState[i] = (byte)character.Level;
                ++i;

                // Current HP
                SplitToBytes(character.CurrentHP, out saveState[i + 1], out saveState[i]);
                i += 2;

                // Max HP
                SplitToBytes(character.MaxHP, out saveState[i + 1], out saveState[i]);
                i += 2;

                // Current MP
                SplitToBytes(character.CurrentMP, out saveState[i + 1], out saveState[i]);
                i += 2;

                // Max MP
                SplitToBytes(character.MaxMP, out saveState[i + 1], out saveState[i]);
                i += 2;

                // Exp
                SplitToBytes(character.Exp, out saveState[i + 2], out saveState[i + 1], out saveState[i]);
                i += 3;

                // Status
                saveState[i] = character.StatusEffetcs.GetStatusEffectsByte();
                ++i;

                // Float
                saveState[i] = (byte)((character.StatusEffetcs.Float) ? 0xFF : 0x0);
                ++i;

                // Commands
                for (uint c = 0; c < 4; ++c)
                {
                    CommandsEnum e = character.GetCommand(c);
                    if (e == CommandsEnum.Unassigned)
                    {
                        saveState[c + i] = 0xFF;
                    }
                    else
                    {
                        saveState[c + i] = (byte)character.GetCommand(c);
                    }
                }
                i += 4;

                // Vigor
                saveState[i] = (byte)character.Vigor;
                ++i;

                // Speed
                saveState[i] = (byte)character.Speed;
                ++i;

                // Stamina
                saveState[i] = (byte)character.Stamina;
                ++i;

                // Magic
                saveState[i] = (byte)character.Magic;
                ++i;

                // Esper
                saveState[i] = (byte)character.Esper;
                ++i;

                // Weapon
                saveState[i] = (byte)character.Weapon;
                ++i;

                // Shield
                saveState[i] = (byte)character.Shield;
                ++i;

                // Helmet
                saveState[i] = (byte)character.Helmet;
                ++i;

                // Armor
                saveState[i] = (byte)character.Armor;
                ++i;

                // Relic 1
                saveState[i] = (byte)character.Relic1;
                ++i;

                // Relic 2
                saveState[i] = (byte)character.Relic2;
                ++i;

                // Magic
                i = offsets.MagicOffset + NEXT_CHARACTER_MAGIC_OFFSET * (int)cEnum;
                for (uint m = 0; m < Enum.GetValues(typeof(MagicEnum)).Length; ++m)
                {
                    saveState[i + m] = character.KnownMagic.GetPercentLeared(m);
                }
            }
        }

        /// <summary>
        /// Save the skills data to the save's byte array
        /// Skills include SwdTech, Lore, Blitzes, Rage, and Dances
        /// </summary>
        /// <param name="skills">The SkillsDto to retrieve information from</param>
        /// <param name="saveState">The save's byte array</param>
        /// <param name="offsets">Contians the offsets for the data in the save's byte array</param>
        private static void SaveSkills(SkillsDto skills, byte[] saveState, Offsets offsets)
        {
            // SwdTech
            saveState[offsets.SwdTechOffset] = BitConverter.GetBytes(FlagsToInt(skills.KnownSwdTechs, Enum.GetValues(typeof(SwdTechEnum))))[0];

            // Lore
            uint flaggedInt = FlagsToInt(skills.KnownLore, Enum.GetValues(typeof(LoreEnum)));
            SplitToBytes(flaggedInt, out saveState[offsets.LoreOffset + 2], out saveState[offsets.LoreOffset + 1], out saveState[offsets.LoreOffset]);

            // Blitz
            saveState[offsets.BlitzOffset] = BitConverter.GetBytes(FlagsToInt(skills.KnownBlitzes, Enum.GetValues(typeof(BlitzEnum))))[0];

            // Rage
            for (int i = 0; i < RagesContainer.GroupCount; ++i)
            {
                saveState[i + offsets.RageOffset] = BitConverter.GetBytes(FlagsToInt(skills.KnownRages, RagesContainer.GetRageDtoGroupForOffset(i)))[0];
            }

            // Dances
            saveState[offsets.DanceOffset] = BitConverter.GetBytes(FlagsToInt(skills.KnownDances, Enum.GetValues(typeof(DanceEnum))))[0];
        }

        /// <summary>
        /// Save the obtained espers data to the save's byte array
        /// </summary>
        /// <param name="knownEspers">The KnownSkillsDto-EsperEnum to retrieve information from</param>
        /// <param name="saveState">The save's byte array</param>
        /// <param name="offsets">Contians the offsets for the data in the save's byte array</param>
        private static void SaveEspers(KnownSkillsDto<EsperEnum> knownEspers, byte[] saveState, Offsets offsets)
        {
            // Known Espers
            for (int i = 0; i < EsperContainer.GroupCount; ++i)
            {
                saveState[i + offsets.EsperOffset] = BitConverter.GetBytes(FlagsToInt(knownEspers, EsperContainer.GetEsperGroupForOffset(i)))[0];
            }
        }

        /// <summary>
        /// Save the other data to the save's byte array
        /// Other data inclues GP, Steps, Number of Saves, Map X/Y Axis, and Airship X/Y Axis
        /// </summary>
        /// <param name="otherStats">The OtherStatsDto to retrieve information from</param>
        /// <param name="saveState">The save's byte array</param>
        /// <param name="offsets">Contians the offsets for the data in the save's byte array</param>
        private static void SaveOtherStats(OtherStatsDto otherStats, byte[] saveState, Offsets offsets)
        {
            // GP
            SplitToBytes(otherStats.GP, out saveState[offsets.GoldOffset + 2], out saveState[offsets.GoldOffset + 1], out saveState[offsets.GoldOffset]);

            // Steps
            SplitToBytes(otherStats.Steps, out saveState[offsets.StepsOffset + 2], out saveState[offsets.StepsOffset + 1], out saveState[offsets.StepsOffset]);

            // Number of Saves
            saveState[offsets.NumberOfSaves] = otherStats.NumberOfSaves;
            saveState[offsets.NumberOfSaves + 1] = otherStats.SaveCountRollover;

            // Map X/Y Axis
            saveState[offsets.MapXYOffset] = otherStats.MapX;
            saveState[offsets.MapXYOffset + 1] = otherStats.MapY;

            // Airship X/Y Axis
            saveState[offsets.AirshipXYOffset] = otherStats.AirshipXAxis;
            saveState[offsets.AirshipXYOffset + 1] = otherStats.AirshipYAxis;

            // Airship Visible
            if (otherStats.IsAirshipVisible)
                saveState[offsets.AirshipSettingsOffset] |= 1 << 1;
            else
                saveState[offsets.AirshipSettingsOffset] = (byte)(saveState[offsets.AirshipSettingsOffset] & ~(1 << 1));
        }

        /// <summary>
        /// Save the inventory data to the save's byte array
        /// </summary>
        /// <param name="inventory">The InventoryDto to retrieve information from</param>
        /// <param name="saveState">The save's byte array</param>
        /// <param name="offsets">Contians the offsets for the data in the save's byte array</param>
        private static void SaveInventory(InventoryDto inventory, byte[] saveState, Offsets offsets)
        {
            for (int i = 0; i < 255; ++i)
            {
                saveState[i + offsets.InventoryItemIdOffset] = inventory.GetItemId(i);
                saveState[i + offsets.InventoryItemCountOffset] = inventory.GetCount(i);
            }
        }

        /// <summary>
        /// Helper method for converting skill flags into an int
        /// </summary>
        /// <typeparam name="S">Skill type to save like DanceEnum, LoreEnum, or BlitzEnum</typeparam>
        /// <param name="dto">Known skills of the S type</param>
        /// <param name="flags">Array of flags</param>
        /// <returns>A value which contains the converted flags</returns>
        private static uint FlagsToInt<S>(KnownSkillsDto<S> dto, Array flags) where S : IConvertible
        {
            uint i = 0;
            foreach (S flag in flags)
            {
                if (dto.IsSkillKnown(flag))
                {
                    i |= (uint)flag.GetHashCode();
                }
            }
            return i;
        }

        private static uint FlagsToInt<S>(KnownSkillsDto<S> dto, List<RageDto> rages) where S : IConvertible
        {
            uint i = 0;
            foreach (RageDto rage in rages)
            {
                if (dto.IsSkillKnown((int)rage.Rage))
                {
                    i |= (uint)rage.Flag;
                }
            }
            return i;
        }

        private static uint FlagsToInt<S>(KnownSkillsDto<S> dto, List<EsperDto> espers) where S : IConvertible
        {
            uint i = 0;
            foreach (EsperDto esper in espers)
            {
                if (dto.IsSkillKnown((int)esper.Esper))
                {
                    i |= (uint)esper.Flag;
                }
            }
            return i;
        }

        /// <summary>
        /// Helper method to reverse the given bytes and combine them into a single value
        /// </summary>
        /// <param name="bytes">Bytes to reverse and combine</param>
        /// <returns>The combined, reversed, bytes</returns>
        private static uint ReverseAndCombine(params byte[] bytes)
        {
            uint i = 0;
            foreach (byte b in bytes)
            {
                i = i << 8;
                i += b;
            }
            return i;
        }

        /// <summary>
        /// Helper method whch splits the given value into 2 bytes (the 2 lowest orders) then sets the left and right according to their placement into the save's array
        /// </summary>
        /// <param name="i">The value to retrieve the bytes from</param>
        /// <param name="left">Should be placed at save byte's array index</param>
        /// <param name="right">Should be placed at save byte's array index + 1</param>
        private static void SplitToBytes(uint i, out byte left, out byte right)
        {
            byte[] b = BitConverter.GetBytes(i);
            right = b[0];
            left = b[1];
        }

        /// <summary>
        /// Helper method whch splits the given value into 3 bytes (the 3 lowest orders) then sets the left, middle, and right according to their placement into the save's array
        /// </summary>
        /// <param name="i">The value to retrieve the bytes from</param>
        /// <param name="left">Should be placed at save byte's array index</param>
        /// <param name="middle">Should be placed at save byte's array index + 1</param>
        /// <param name="right">Should be placed at save byte's array index + 2</param>
        private static void SplitToBytes(uint i, out byte left, out byte middle, out byte right)
        {
            byte[] b = BitConverter.GetBytes(i);
            right = b[0];
            middle = b[1];
            left = b[2];
        }

        /// <summary>
        /// Determine the known skills from the given value which is a bit set of flags
        /// </summary>
        /// <param name="dto">KnownSkillsDto to populate</param>
        /// <param name="rages">All possible flags for this flag set</param>
        /// <param name="bits">The value to retrieve the flags from</param>
        private static void DetermineKnownSkills<S>(KnownSkillsDto<S> dto, Array skills, uint bits) where S : IConvertible
        {
            int i = 0;
            foreach (IConvertible skill in skills)
            {
                dto.SetSkillKnown(i, (bits & skill.GetHashCode()) != 0);
                ++i;
            }
        }

        private static void DetermineKnownSkills(KnownSkillsDto<IConvertible> knownRages, List<RageDto> rages, uint bits)
        {
            foreach (RageDto rage in rages)
            {
                knownRages.SetSkillKnown((int)rage.Rage, (bits & rage.Flag) != 0);
            }
        }

        private static void DetermineKnownSkills(KnownSkillsDto<EsperEnum> knownEspers, List<EsperDto> espers, uint bits)
        {
            foreach (EsperDto esper in espers)
            {
                knownEspers.SetSkillKnown((int)esper.Esper, (bits & esper.Flag) != 0);
            }
        }

        private static bool TryFindDynamicOffset(SaveDto saveDto, out int offset, string nameOfTerra = "TERRA")
        {
            offset = -1;
            saveDto.NameOfTerra = Interaction.InputBox("Terra's Name", "", nameOfTerra);
            if (saveDto.NameOfTerra != "")
            {
                char[] terraName = saveDto.NameOfTerra.ToCharArray();
                byte[] name = new byte[terraName.Length];
                int i = 0;
                for (; i < terraName.Length; ++i)
                {
                    if (terraName[i] == '?')
                        name[i] = 0xBF;
                    else
                        name[i] = (byte)(terraName[i] + 63);
                }
                try
                {
                    bool found = false;
                    byte[] bytes = saveDto.SaveByteArray;
                    for (i = 0; i < bytes.Length; ++i)
                    {
                        if (name[0] == bytes[i])
                        {
                            found = true;
                            for (int j = 0; j < name.Length && i + j < bytes.Length; ++j)
                            {
                                if (name[j] != bytes[i + j])
                                {
                                    found = false;
                                    break;
                                }
                            }
                            if (found)
                                break;
                        }
                    }
                    if (found)
                    {
                        offset = i;
                        return true;
                    }
                }
                catch
                {
                }
            }
            MessageBox.Show(
                "Failed to find offset based off the given name.\n" +
                "The name is case sensitive.\n" +
                "If saving and Terra's name changed use Terra's previous name.");
            saveDto.NameOfTerra = "";
            return false;
        }
    }
}
