using System;

namespace FF6_Save_Editor.Enums
{
    /// <summary>
    /// Defines each StatusEnum and the byte flag which specifies it
    /// </summary>
    [Flags]
    public enum StatusEnum
    {
        None = 0x0,
        Darkness = 0x1,
        Zombie = 0x2,
        Poison = 0x4,
        Magitek = 0x8,
        Invisible = 0x10,
        Imp = 0x20,
        Stone = 0x40,
        Wounded = 0x80,
        Float = 0x100 // This is in another byte and will not 'flag' with the others, this is here for better code flow to the UI
    }
}
