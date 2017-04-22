using System;

namespace FF6_Save_Editor.Enums
{
    /// <summary>
    /// Defines each Dance and the byte flag which specifies it
    /// </summary>
    [Flags]
    enum DanceEnum
    {
        WindSong = 0x1,
        ForestSuite = 0x2,
        DesertAria = 0x4,
        LoveSonata = 0x8,
        EarthBlues = 0x10,
        WaterRondo = 0x20,
        DuskRequium = 0x40,
        SnowmanJazz = 0x80
    }
}
