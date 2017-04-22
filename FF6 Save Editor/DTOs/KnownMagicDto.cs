using FF6_Save_Editor.Enums;
using System;

namespace FF6_Save_Editor.DTOs
{
    class KnownMagicDto
    {
        private class MagicDto
        {
            public MagicEnum Magic { get; private set; }
            public byte PercentKnown { get; set; }

            public MagicDto(MagicEnum magic)
            {
                this.Magic = magic;
            }

            public override string ToString()
            {
                return Magic.ToString() + " - " + PercentKnown.ToString();
            }
        }

        private readonly MagicDto[] knownMagic;

        public KnownMagicDto()
        {
            Array magics = Enum.GetValues(typeof(MagicEnum));
            knownMagic = new MagicDto[magics.Length];
            uint i = 0;
            foreach (MagicEnum m in magics)
            {
                knownMagic[i] = new MagicDto(m);
                ++i;
            }
        }

        public byte GetPercentLeared(MagicEnum magic)
        {
            return this.knownMagic[(uint)magic].PercentKnown;
        }

        public byte GetPercentLeared(uint i)
        {
            if (i > this.knownMagic.Length)
            {
                throw new IndexOutOfRangeException(i + " is too big for knownSkills. Max index is " + (this.knownMagic.Length - 1));
            }
            return this.knownMagic[i].PercentKnown;
        }

        public void SetPercentLearned(MagicEnum magic, byte percentKnown)
        {
            this.knownMagic[(uint)magic].PercentKnown = percentKnown;
        }

        public void SetPercentLearned(uint i, byte percentKnown)
        {
            if (i > this.knownMagic.Length)
            {
                throw new IndexOutOfRangeException(i + " is too big for knownSkills. Max index is " + (this.knownMagic.Length - 1));
            }
            this.knownMagic[i].PercentKnown = percentKnown;
        }
    }
}
