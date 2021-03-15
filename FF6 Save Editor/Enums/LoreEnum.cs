using System;

namespace FF6_Save_Editor.Enums
{
    /// <summary>
    /// Defines each Lore and the byte flag which specifies it
    /// </summary>
    [Flags]
    enum LoreEnum
    {
        Condemned = 0x1,
        Roulette = 0x2,
        CleanSweep = 0x4,
        AquaRake = 0x8,
        Aero = 0x10,
        BlowFish = 0x20,
        BigGuard = 0x40,
        Revenge = 0x80,
        PearlWind = 0x100,
        L5Doom = 0x200,
        L4Flare = 0x400,
        L3Muddle = 0x800,
        Reflect = 0x1000,
        LPearl = 0x2000,
        StepMine = 0x4000,
        ForceField = 0x8000,
        Dischord = 0x10000,
        SourMouth = 0x20000,
        PepUp = 0x40000,
        Rippler = 0x80000,
        Stone = 0x100000,
        Quasar = 0x200000,
        GrandTrain = 0x400000,
        Exploder = 0x800000
    }
}
