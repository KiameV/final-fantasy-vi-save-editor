using FF6_Save_Editor.Enums;
using System;

namespace FF6_Save_Editor.DTOs
{
    /// <summary>
    /// Defines all information specific to a Character
    /// </summary>
    class CharacterDto
    {
        /// <summary>
        /// Specifies which Character this is
        /// </summary>
        public readonly CharacterEnum Character;
        /// <summary>
        /// The commands which are available for the character in combat
        /// </summary>
        private readonly CommandsEnum[] commands = new CommandsEnum[4];
        /// <summary>
        /// All Magic known by the character
        /// </summary>
        public readonly KnownMagicDto KnownMagic = new KnownMagicDto();
        /// <summary>
        /// Status effects currently applied to the character
        /// </summary>
        public readonly StatusEffectsDto StatusEffetcs = new StatusEffectsDto();
        /// <summary>
        /// The Character's custom name
        /// </summary>
        public string Name { get; set; }
        /// <summary>
        /// The Character's current level
        /// </summary>
        public uint Level { get; set; }
        /// <summary>
        /// The Character's current HP
        /// </summary>
        public uint CurrentHP { get; set; }
        /// <summary>
        /// The Character's max HP
        /// </summary>
        public uint MaxHP { get; set; }
        /// <summary>
        /// The Character's current MP
        /// </summary>
        public uint CurrentMP { get; set; }
        /// <summary>
        /// The Character's max MP
        /// </summary>
        public uint MaxMP { get; set; }
        /// <summary>
        /// The Character's vigor
        /// </summary>
        public uint Vigor { get; set; }
        /// <summary>
        /// The Character's stamina
        /// </summary>
        public uint Stamina { get; set; }
        /// <summary>
        /// The Character's speed
        /// </summary>
        public uint Speed { get; set; }
        /// <summary>
        /// The Character's magic
        /// </summary>
        public uint Magic { get; set; }
        /// <summary>
        /// The Character's experience
        /// </summary>
        public uint Exp { get; set; }
        /// <summary>
        /// The Esper which is currently equiped
        /// </summary>
        public EsperEnum Esper { get; set; }
        /// <summary>
        /// The id of the Weapon which is currently equiped
        /// </summary>
        public uint Weapon { get; set; }
        /// <summary>
        /// The id of the Shield which is currently equiped
        /// </summary>
        public uint Shield { get; set; }
        /// <summary>
        /// The id of the Helmet which is currently equiped
        /// </summary>
        public uint Helmet { get; set; }
        /// <summary>
        /// The id of the Armor which is currently equiped
        /// </summary>
        public uint Armor { get; set; }
        /// <summary>
        /// The id of the Relic which is currently equiped in slot 1
        /// </summary>
        public uint Relic1 { get; set; }
        /// <summary>
        /// The id of the Relic which is currently equiped in slot 2
        /// </summary>
        public uint Relic2 { get; set; }

        /// <summary>
        /// Constructor
        /// </summary>
        /// <param name="character">The CharacterEnum for this character</param>
        public CharacterDto (CharacterEnum character)
        {
            this.Character = character;
        }

        /// <summary>
        /// Get the command at the specified index.
        /// </summary>
        /// <param name="i">Valid indexes are 0, 1, 2, and 3.</param>
        /// <returns>The command at the specified index</returns>
        public CommandsEnum GetCommand(uint i)
        {
            if (i >= this.commands.Length)
            {
                throw new IndexOutOfRangeException(i + " is larger than commands array.");
            }
            return this.commands[i];
        }

        /// <summary>
        /// Set the command at the specified index.
        /// </summary>
        /// <param name="i">Valid indexes are 0, 1, 2, and 3.</param>
        /// <param name="command">The command to set</param>
        public void SetCommand(uint i, CommandsEnum command)
        {
            if (i >= this.commands.Length)
            {
                throw new IndexOutOfRangeException(i + " is larger than commands array.");
            }
            this.commands[i] = command;
        }
    }
}
