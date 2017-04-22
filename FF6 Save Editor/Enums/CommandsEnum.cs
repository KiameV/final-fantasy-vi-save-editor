using System;

namespace FF6_Save_Editor.Enums
{
    /// <summary>
    /// Defines each Command and the byte flag which specifies it
    /// </summary>
    [Flags]
    enum CommandsEnum
    {
        Fight = 0x0,
        Item = 0x1,
        Magic = 0x2,
        Morph = 0x3,
        Revert = 0x4,
        Steal = 0x5,
        Capture = 0x6,
        SwdTech = 0x7,
        Throw = 0x8,
        Tools = 0x9,
        Blitz = 0x0A,
        Runic = 0x0B,
        Lore = 0x0C,
        Sketch = 0x0D,
        Control = 0x0E,
        Slot = 0x0F,
        Rage = 0x10,
        Leap = 0x11,
        Mimic = 0x12,
        Dance = 0x13,
        Row = 0x14,
        Def = 0x15,
        Jump = 0x16,
        XMagic = 0x17,
        GPRain = 0x18,
        Summon = 0x19,
        Health = 0x1A,
        Shock = 0x1B,
        Possess = 0x1C,
        Magitek = 0x1D,
        Blank = 0x1E,
        Unassigned = 0xFF
    }
}
