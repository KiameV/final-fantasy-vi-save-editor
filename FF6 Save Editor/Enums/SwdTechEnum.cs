using System;

namespace FF6_Save_Editor.Enums
{
    /// <summary>
    /// Defines each SwdTech and the byte flag which specifies it
    /// </summary>
    [Flags]
    enum SwdTechEnum
    {
        Dispatch = 0x1,
        Retort = 0x2,
        Slash = 0x4,
        QuadraSlam = 0x8,
        Empowerer = 0x10,
        Stunner = 0x20,
        QuadraSlice = 0x40,
        Cleave = 0x80
    }
}
