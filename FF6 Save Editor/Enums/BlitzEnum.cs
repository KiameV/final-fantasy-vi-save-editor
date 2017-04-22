using System;

namespace FF6_Save_Editor.Enums
{
    /// <summary>
    /// Defines each Blitz and the byte flag which specifies it
    /// </summary>
    [Flags]
    enum BlitzEnum
    {
        Pummel = 0x1,
        Surplex = 0x2,
        AuraBolt = 0x4,
        FireDance = 0x8,
        Mantra = 0x10,
        AirBlade = 0x20,
        Spiraler = 0x40,
        BumRush = 0x80
    }
}
