using System;

namespace FF6_Save_Editor.Enums
{
    /// <summary>
    /// Defines each Lore and the byte flag which specifies it
    /// </summary>
    [Flags]
    enum LoreEnum
    {
        Condemned = 0x0,
        Roulette = 0x1,
        CleanSweep = 0x2,
        AquaRake = 0x4,
        Aero = 0x8,
        BlowFish = 0x10,
        BigGuard = 0x20,
        Revenge = 0x40,
        PearlWind = 0x80,
        L5Doom = 0x100,
        L4Flare = 0x200,
        L3Muddle = 0x400,
        Reflect = 0x800,
        LPearl = 0x1000,
        StepMine = 0x2000,
        ForceField = 0x4000,
        Dischord = 0x8000,
        SourMouth = 0x10000,
        PepUp = 0x20000,
        Rippler = 0x40000,
        Stone = 0x80000,
        Quasar = 0x100000,
        GrandTrain = 0x200000,
        Exploder = 0x400000
    }
}
