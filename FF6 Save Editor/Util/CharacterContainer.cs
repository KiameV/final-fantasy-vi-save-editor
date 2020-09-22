using FF6_Save_Editor.Enums;
using System;

namespace FF6_Save_Editor.DTOs
{
    /// <summary>
    /// Interface for finding CharacterDtos
    /// </summary>
    class CharacterContainer
    {
        /// <summary>
        /// Contains all characters
        /// </summary>
        private CharacterDto[] characters;

        /// <summary>
        /// Define a CharacterDto for each CharacterEnum
        /// </summary>
        public CharacterContainer()
        {
            Array characterEnums = Enum.GetValues(typeof(CharacterEnum));
            this.characters = new CharacterDto[characterEnums.Length];
            foreach (CharacterEnum character in characterEnums)
            {
                this.characters[(uint)character] = new CharacterDto(character);
            }
        }

        /// <summary>
        /// Get the character for the specified offset
        /// </summary>
        /// <param name="offset">Offset for the character to get</param>
        /// <returns>The character at the specified offset.</returns>
        /// <exception cref="IndexOutOfRangeException">If the offset is invalid (too small/large)</exception>
        public CharacterDto getCharacter(byte offset)
        {
            if (offset > characters.Length)
            {
                throw new IndexOutOfRangeException(offset + " is larger than characters. Max index is " + (this.characters.Length - 1));
            }
            return this.characters[offset];
        }

        /// <summary>
        /// Get the CharacterDto for the specified CharacterEnum
        /// </summary>
        /// <param name="character">CharacterEnum to get the CharacterDto of</param>
        /// <returns>The CharacterDto for the specified CharacterEnum</returns>
        public CharacterDto getCharacter(CharacterEnum character)
        {
            return this.characters[(uint)character];
        }

        public CharacterDto this[int i] => this.characters[i];

        public int Count => this.characters.Length;
    }
}
