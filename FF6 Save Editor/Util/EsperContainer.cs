using FF6_Save_Editor.Enums;
using System;
using System.Collections.Generic;

namespace FF6_Save_State_Editor.Util
{
    /// <summary>
    /// Interface for finding Espers
    /// </summary>
    class EsperContainer
    {
        /// <summary>
        /// Contains the EsperEnum, offset, and bit flag for a Rage
        /// </summary>
        public class EsperDto
        {
            /// <summary>
            /// Defines which Esper this is
            /// </summary>
            public EsperEnum Esper { get; private set; }
            /// <summary>
            /// Offset from the beginning of the Espers offsets
            /// </summary>
            public int Offset { get; private set; }
            /// <summary>
            /// Bit flag for determining whether the Esper is available or not
            /// </summary>
            public byte Flag { get; private set; }

            internal EsperDto(EsperEnum esper, int offset, byte flag)
            {
                this.Esper = esper;
                this.Offset = offset;
                this.Flag = flag;
            }
        }

        /// <summary>
        /// List of all Espers
        /// </summary>
        private static List<EsperDto> espers = new List<EsperDto>(Enum.GetValues(typeof(EsperEnum)).Length);

        /// <summary>
        /// Number of Espers
        /// </summary>
        public static int Count { get { return espers.Count; } }

        /// <summary>
        /// A 'group' is the Espers that are defined in a specific byte. This is the total number of groups.
        /// </summary>
        public static int GroupCount { get; private set; }

        /// <summary>
        /// Get the EsperDto for the specified EsperEnum
        /// </summary>
        /// <param name="esper">EsperEnum to get the EsperDto for</param>
        /// <returns>The EsperDto for the specified EsperEnum</returns>
        public static EsperDto GetEsperDto(EsperEnum esper)
        {
            return espers[(int)esper];
        }

        /// <summary>
        /// Get the Espers defined in the byte specified by the offset
        /// </summary>
        /// <param name="offset">The offset from the beginning of the esper's offset</param>
        /// <returns>A list of all Espers which are in the offset group</returns>
        public static List<EsperDto> GetEsperGroupForOffset(int offset)
        {
            List<EsperDto> l = new List<EsperDto>();
            foreach (EsperDto esper in espers)
            {
                if (esper.Offset == offset)
                {
                    l.Add(esper);
                }
                else if (esper.Offset > offset)
                {
                    break;
                }
            }
            return l;
        }

        /// <summary>
        /// Create an EsperDto for each Esper
        /// </summary>
        static EsperContainer()
        {
            // Define each Esper by group
            espers.Add(new EsperDto(EsperEnum.Ramuh, 0, 0x1));
            espers.Add(new EsperDto(EsperEnum.Ifrit, 0, 0x2));
            espers.Add(new EsperDto(EsperEnum.Shiva, 0, 0x4));
            espers.Add(new EsperDto(EsperEnum.Siren, 0, 0x8));
            espers.Add(new EsperDto(EsperEnum.Terrato, 0, 0x10));
            espers.Add(new EsperDto(EsperEnum.Maduin, 0, 0x20));
            espers.Add(new EsperDto(EsperEnum.Shoat, 0, 0x40));
            espers.Add(new EsperDto(EsperEnum.Bismark, 0, 0x80));

            espers.Add(new EsperDto(EsperEnum.Stray, 1, 0x1));
            espers.Add(new EsperDto(EsperEnum.Palidor, 1, 0x2));
            espers.Add(new EsperDto(EsperEnum.Tritoch, 1, 0x4));
            espers.Add(new EsperDto(EsperEnum.Odin, 1, 0x8));
            espers.Add(new EsperDto(EsperEnum.Raiden, 1, 0x10));
            espers.Add(new EsperDto(EsperEnum.Bahamut, 1, 0x20));
            espers.Add(new EsperDto(EsperEnum.Alexandr, 1, 0x40));
            espers.Add(new EsperDto(EsperEnum.Crusader, 1, 0x80));

            espers.Add(new EsperDto(EsperEnum.Ragnarok, 2, 0x1));
            espers.Add(new EsperDto(EsperEnum.Kirin, 2, 0x2));
            espers.Add(new EsperDto(EsperEnum.Zoneseek, 2, 0x4));
            espers.Add(new EsperDto(EsperEnum.Carbunkl, 2, 0x8));
            espers.Add(new EsperDto(EsperEnum.Phantom, 2, 0x10));
            espers.Add(new EsperDto(EsperEnum.Sraphim, 2, 0x20));
            espers.Add(new EsperDto(EsperEnum.Golem, 2, 0x40));
            espers.Add(new EsperDto(EsperEnum.Unicorn, 2, 0x80));

            espers.Add(new EsperDto(EsperEnum.Fenrir, 3, 0x1));
            espers.Add(new EsperDto(EsperEnum.Startlet, 3, 0x2));
            espers.Add(new EsperDto(EsperEnum.Phoenix, 3, 0x4));

            // Determine the number of groups
            int currentGroup = -1;
            int groupCount = 0;
            foreach (EsperDto esper in espers)
            {
                if (currentGroup != esper.Offset)
                {
                    currentGroup = esper.Offset;
                    ++groupCount;
                }
            }
            GroupCount = groupCount;
        }
    }
}
