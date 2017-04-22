using FF6_Save_Editor.Enums;
using System.Collections.Generic;
using System;
using System.Windows.Forms;

namespace FF6_Save_Editor.DTOs
{
    /// <summary>
    /// Defines the different status effects
    /// </summary>
    class StatusEffectsDto
    {
        /// <summary>
        /// Whether a character has Darkness
        /// </summary>
        public bool Darkness { get; set; }
        /// <summary>
        /// Whether a character has Zombie
        /// </summary>
        public bool Zombie { get; set; }
        /// <summary>
        /// Whether a character has Poison
        /// </summary>
        public bool Poison { get; set; }
        /// <summary>
        /// Whether a character is in a Magitek suit
        /// </summary>
        public bool Magitek { get; set; }
        /// <summary>
        /// Whether a character is Invisible
        /// </summary>
        public bool Invisible { get; set; }
        /// <summary>
        /// Whether a character is in Imp form
        /// </summary>
        public bool Imp { get; set; }
        /// <summary>
        /// Whether a character is Petrified
        /// </summary>
        public bool Stone { get; set; }
        /// <summary>
        /// Whether a character is Wounded
        /// </summary>
        public bool Wounded { get; set; }
        /// <summary>
        /// Whether a character has Float
        /// </summary>
        public bool Float { get; set; }

        /// <summary>
        /// Get the byte which represents the Character's current statuses
        /// </summary>
        /// <returns>The byte which represents the Character's current statuses</returns>
        public byte GetStatusEffectsByte()
        {
            byte b = 0;
            if (this.Darkness)
            {
                b |= (int)StatusEnum.Darkness;
            }
            if (this.Zombie)
            {
                b |= (int)StatusEnum.Zombie;
            }
            if (this.Poison)
            {
                b |= (int)StatusEnum.Poison;
            }
            if (this.Magitek)
            {
                b |= (int)StatusEnum.Magitek;
            }
            if (this.Invisible)
            {
                b |= (int)StatusEnum.Invisible;
            }
            if (this.Imp)
            {
                b |= (int)StatusEnum.Imp;
            }
            if (this.Stone)
            {
                b |= (int)StatusEnum.Stone;
            }
            if (this.Wounded)
            {
                b |= (int)StatusEnum.Wounded;
            }
            return b;
        }

        /// <summary>
        /// Gets a list of StatusEnums which specify which status effects a character has
        /// </summary>
        /// <returns>A list of StatusEnums which specify which status effects a character has</returns>
        public List<StatusEnum> GetStatusEffects()
        {
            List<StatusEnum> statusEffects = new List<StatusEnum>();
            if (this.Darkness)
            {
                statusEffects.Add(StatusEnum.Darkness);
            }
            if (this.Zombie)
            {
                statusEffects.Add(StatusEnum.Zombie);
            }
            if (this.Poison)
            {
                statusEffects.Add(StatusEnum.Poison);
            }
            if (this.Magitek)
            {
                statusEffects.Add(StatusEnum.Magitek);
            }
            if (this.Invisible)
            {
                statusEffects.Add(StatusEnum.Invisible);
            }
            if (this.Imp)
            {
                statusEffects.Add(StatusEnum.Imp);
            }
            if (this.Stone)
            {
                statusEffects.Add(StatusEnum.Stone);
            }
            if (this.Wounded)
            {
                statusEffects.Add(StatusEnum.Wounded);
            }
            if (this.Float)
            {
                statusEffects.Add(StatusEnum.Float);
            }
            return statusEffects;
        }

        /// <summary>
        /// Sets the status effects based off the given group of status effects
        /// </summary>
        /// <param name="statusEffects">Group of status effects</param>
        public void SetStatusEffects(IEnumerable<object> statusEffects)
        {
            List<object> l = new List<object>(statusEffects);
            this.Darkness = (l.Contains(StatusEnum.Darkness)) ? true : false;
            this.Zombie = (l.Contains(StatusEnum.Zombie)) ? true : false;
            this.Poison = (l.Contains(StatusEnum.Poison)) ? true : false;
            this.Magitek = (l.Contains(StatusEnum.Magitek)) ? true : false;
            this.Invisible = (l.Contains(StatusEnum.Invisible)) ? true : false;
            this.Imp = (l.Contains(StatusEnum.Imp)) ? true : false;
            this.Stone = (l.Contains(StatusEnum.Stone)) ? true : false;
            this.Wounded = (l.Contains(StatusEnum.Wounded)) ? true : false;
            this.Float = (l.Contains(StatusEnum.Float)) ? true : false;
        }
    }
}
