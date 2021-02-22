using FF6_Save_Editor.Enums;
using System;
using System.Collections.Generic;

namespace FF6_Save_State_Editor.Util
{
    /// <summary>
    /// Interface for retrieving Rages
    /// </summary>
    public static class RagesContainer
    {
        /// <summary>
        /// Contains the RageEnum, offset, and bit flag for a Rage
        /// </summary>
        public class RageDto
        {
            /// <summary>
            /// Specifies which Rage this is
            /// </summary>
            public RageEnum Rage { get; private set; }
            /// <summary>
            /// Offset from the beginning of the Rage offsets
            /// </summary>
            public int Offset { get; private set; }
            /// <summary>
            /// Bit flag for determining whether the Rage is available or not
            /// </summary>
            public byte Flag { get; private set; }

            internal RageDto(RageEnum rage, int offset, byte flag)
            {
                this.Rage = rage;
                this.Offset = offset;
                this.Flag = flag;
            }
        }

        /// <summary>
        /// A list of all Rages
        /// </summary>
        private static List<RageDto> rages = new List<RageDto>(Enum.GetValues(typeof(RageEnum)).Length);

        /// <summary>
        /// Number of all possible rages
        /// </summary>
        public static int Count { get { return rages.Count; } }

        /// <summary>
        /// A 'group' is the Rages that are defined in a specific byte. This is the total number of groups.
        /// </summary>
        public static int GroupCount { get; private set; }

        /// <summary>
        /// Get the RageDTO for the specified RageEnum
        /// </summary>
        /// <param name="rage">The rage to get the RageDTO for</param>
        /// <returns>The RageDTO that defines the given RageEnum</returns>
        public static RageDto GetRageDto(RageEnum rage)
        {
            return rages[(int)rage];
        }

        /// <summary>
        /// Get the group of Rages for the specified offset
        /// </summary>
        /// <param name="offset">The offset for the rages to get</param>
        /// <returns>The group of Rages for the specified offset</returns>
        public static List<RageDto> GetRageDtoGroupForOffset(int offset)
        {
            List<RageDto> l = new List<RageDto>(8);
            foreach (RageDto rage in rages)
            {
                if (rage.Offset == offset)
                {
                    l.Add(rage);
                }
                else if (rage.Offset > offset)
                {
                    break;
                }
            }
            return l;
        }

        /// <summary>
        /// Create a RageDTO for each Rage and populate the rages list with each RageDTO
        /// </summary>
        static RagesContainer()
        {
            // Create each RageDTO in the order so they are in their proper group
            rages.Add(new RageDto(RageEnum.Guard, 0, 0x1));
            rages.Add(new RageDto(RageEnum.Soldier, 0, 0x2));
            rages.Add(new RageDto(RageEnum.Templar, 0, 0x4));
            rages.Add(new RageDto(RageEnum.Ninja, 0, 0x8));
            rages.Add(new RageDto(RageEnum.Samurai, 0, 0x10));
            rages.Add(new RageDto(RageEnum.Orog, 0, 0x20));
            rages.Add(new RageDto(RageEnum.MagRoader, 0, 0x40));
            rages.Add(new RageDto(RageEnum.Retainer, 0, 0x80));

            rages.Add(new RageDto(RageEnum.Hazer, 1, 0x1));
            rages.Add(new RageDto(RageEnum.Dahling, 1, 0x2));
            rages.Add(new RageDto(RageEnum.RainMan, 1, 0x4));
            rages.Add(new RageDto(RageEnum.Brawler, 1, 0x8));
            rages.Add(new RageDto(RageEnum.Apokryphos, 1, 0x10));
            rages.Add(new RageDto(RageEnum.DarkForce, 1, 0x20));
            rages.Add(new RageDto(RageEnum.Whisper, 1, 0x40));
            rages.Add(new RageDto(RageEnum.OverMind, 1, 0x80));

            rages.Add(new RageDto(RageEnum.Osteosaur, 2, 0x1));
            rages.Add(new RageDto(RageEnum.Commander, 2, 0x2));
            rages.Add(new RageDto(RageEnum.Rhodox, 2, 0x4));
            rages.Add(new RageDto(RageEnum.WereRat, 2, 0x8));
            rages.Add(new RageDto(RageEnum.Ursus, 2, 0x10));
            rages.Add(new RageDto(RageEnum.Rhinotaur, 2, 0x20));
            rages.Add(new RageDto(RageEnum.Steroidite, 2, 0x40));
            rages.Add(new RageDto(RageEnum.Leafer, 2, 0x80));

            rages.Add(new RageDto(RageEnum.StrayCat, 3, 0x1));
            rages.Add(new RageDto(RageEnum.Lobo, 3, 0x2));
            rages.Add(new RageDto(RageEnum.Doberman, 3, 0x4));
            rages.Add(new RageDto(RageEnum.Vomammoth, 3, 0x8));
            rages.Add(new RageDto(RageEnum.Fidor, 3, 0x10));
            rages.Add(new RageDto(RageEnum.Baskervor, 3, 0x20));
            rages.Add(new RageDto(RageEnum.Suriander, 3, 0x40));
            rages.Add(new RageDto(RageEnum.Chimera, 3, 0x80));

            rages.Add(new RageDto(RageEnum.Behemoth, 4, 0x1));
            rages.Add(new RageDto(RageEnum.Mesosaur, 4, 0x2));
            rages.Add(new RageDto(RageEnum.Pterodon, 4, 0x4));
            rages.Add(new RageDto(RageEnum.FossilFang, 4, 0x8));
            rages.Add(new RageDto(RageEnum.WhiteDrgn, 4, 0x10));
            rages.Add(new RageDto(RageEnum.DoomDrgn, 4, 0x20));
            rages.Add(new RageDto(RageEnum.Brachosaur, 4, 0x40));
            rages.Add(new RageDto(RageEnum.Tyranosaur, 4, 0x80));

            rages.Add(new RageDto(RageEnum.DarkWind, 5, 0x1));
            rages.Add(new RageDto(RageEnum.Beakor, 5, 0x2));
            rages.Add(new RageDto(RageEnum.Vulture, 5, 0x4));
            rages.Add(new RageDto(RageEnum.Harpy, 5, 0x8));
            rages.Add(new RageDto(RageEnum.HermitCrab, 5, 0x10));
            rages.Add(new RageDto(RageEnum.Trapper, 5, 0x20));
            rages.Add(new RageDto(RageEnum.Hornet, 5, 0x40));
            rages.Add(new RageDto(RageEnum.Crasshoppr, 5, 0x80));

            rages.Add(new RageDto(RageEnum.DeltaBug, 6, 0x1));
            rages.Add(new RageDto(RageEnum.Gilomantis, 6, 0x2));
            rages.Add(new RageDto(RageEnum.Trilium, 6, 0x4));
            rages.Add(new RageDto(RageEnum.Nightshade, 6, 0x8));
            rages.Add(new RageDto(RageEnum.Tumbleweed, 6, 0x10));
            rages.Add(new RageDto(RageEnum.Bloompire, 6, 0x20));
            rages.Add(new RageDto(RageEnum.Trilobiter, 6, 0x40));
            rages.Add(new RageDto(RageEnum.Siegfried, 6, 0x80));

            rages.Add(new RageDto(RageEnum.Nautiloid, 7, 0x1));
            rages.Add(new RageDto(RageEnum.Exocite, 7, 0x2));
            rages.Add(new RageDto(RageEnum.Anguiform, 7, 0x4));
            rages.Add(new RageDto(RageEnum.ReachFrog, 7, 0x8));
            rages.Add(new RageDto(RageEnum.Lizard, 7, 0x10));
            rages.Add(new RageDto(RageEnum.Chickenlip, 7, 0x20));
            rages.Add(new RageDto(RageEnum.Hoover, 7, 0x40));
            rages.Add(new RageDto(RageEnum.Rider, 7, 0x80));

            rages.Add(new RageDto(RageEnum.Chupon, 8, 0x1));
            rages.Add(new RageDto(RageEnum.Pipsqueak, 8, 0x2));
            rages.Add(new RageDto(RageEnum.MTekArmor, 8, 0x4));
            rages.Add(new RageDto(RageEnum.SkyArmor, 8, 0x8));
            rages.Add(new RageDto(RageEnum.Telstar, 8, 0x10));
            rages.Add(new RageDto(RageEnum.LethalWpn, 8, 0x20));
            rages.Add(new RageDto(RageEnum.Vaporite, 8, 0x40));
            rages.Add(new RageDto(RageEnum.Flan, 8, 0x80));

            rages.Add(new RageDto(RageEnum.Ing, 9, 0x1));
            rages.Add(new RageDto(RageEnum.Humpty, 9, 0x2));
            rages.Add(new RageDto(RageEnum.Brainpan, 9, 0x4));
            rages.Add(new RageDto(RageEnum.Cruller, 9, 0x8));
            rages.Add(new RageDto(RageEnum.Cactrot, 9, 0x10));
            rages.Add(new RageDto(RageEnum.RepoMan, 9, 0x20));
            rages.Add(new RageDto(RageEnum.Harvester, 9, 0x40));
            rages.Add(new RageDto(RageEnum.Bomb, 9, 0x80));

            rages.Add(new RageDto(RageEnum.StillLife, 10, 0x1));
            rages.Add(new RageDto(RageEnum.BoxedSet, 10, 0x2));
            rages.Add(new RageDto(RageEnum.SlamDancer, 10, 0x4));
            rages.Add(new RageDto(RageEnum.HadesGigas, 10, 0x8));
            rages.Add(new RageDto(RageEnum.Pug, 10, 0x10));
            rages.Add(new RageDto(RageEnum.MagicUrn, 10, 0x20));
            rages.Add(new RageDto(RageEnum.Mover, 10, 0x40));
            rages.Add(new RageDto(RageEnum.Figaliz, 10, 0x80));

            rages.Add(new RageDto(RageEnum.Buffalax, 11, 0x1));
            rages.Add(new RageDto(RageEnum.Aspik, 11, 0x2));
            rages.Add(new RageDto(RageEnum.Ghost, 11, 0x4));
            rages.Add(new RageDto(RageEnum.Crawler, 11, 0x8));
            rages.Add(new RageDto(RageEnum.SandRay, 11, 0x10));
            rages.Add(new RageDto(RageEnum.Areneid, 11, 0x20));
            rages.Add(new RageDto(RageEnum.Actaneon, 11, 0x40));
            rages.Add(new RageDto(RageEnum.SandHorse, 11, 0x80));

            rages.Add(new RageDto(RageEnum.DarkSide, 12, 0x1));
            rages.Add(new RageDto(RageEnum.MadOscar, 12, 0x2));
            rages.Add(new RageDto(RageEnum.Crawly, 12, 0x4));
            rages.Add(new RageDto(RageEnum.Bleary, 12, 0x8));
            rages.Add(new RageDto(RageEnum.Marshal, 12, 0x10));
            rages.Add(new RageDto(RageEnum.Trooper, 12, 0x20));
            rages.Add(new RageDto(RageEnum.General, 12, 0x40));
            rages.Add(new RageDto(RageEnum.Covert, 12, 0x80));

            rages.Add(new RageDto(RageEnum.Ogor, 13, 0x1));
            rages.Add(new RageDto(RageEnum.Warlock, 13, 0x2));
            rages.Add(new RageDto(RageEnum.Madam, 13, 0x4));
            rages.Add(new RageDto(RageEnum.Joker, 13, 0x8));
            rages.Add(new RageDto(RageEnum.IronFist, 13, 0x10));
            rages.Add(new RageDto(RageEnum.Goblin, 13, 0x20));
            rages.Add(new RageDto(RageEnum.Apparite, 13, 0x40));
            rages.Add(new RageDto(RageEnum.PowerDemon, 13, 0x80));

            rages.Add(new RageDto(RageEnum.Displayer, 14, 0x1));
            rages.Add(new RageDto(RageEnum.VectorPup, 14, 0x2));
            rages.Add(new RageDto(RageEnum.Peepers, 14, 0x4));
            rages.Add(new RageDto(RageEnum.SewerRat, 14, 0x8));
            rages.Add(new RageDto(RageEnum.Slatter, 14, 0x10));
            rages.Add(new RageDto(RageEnum.Rhinox, 14, 0x20));
            rages.Add(new RageDto(RageEnum.Rhobite, 14, 0x40));
            rages.Add(new RageDto(RageEnum.WildCat, 14, 0x80));

            rages.Add(new RageDto(RageEnum.RedFang, 15, 0x1));
            rages.Add(new RageDto(RageEnum.BountyMan, 15, 0x2));
            rages.Add(new RageDto(RageEnum.Tusker, 15, 0x4));
            rages.Add(new RageDto(RageEnum.Ralph, 15, 0x8));
            rages.Add(new RageDto(RageEnum.Chitonid, 15, 0x10));
            rages.Add(new RageDto(RageEnum.WartPuck, 15, 0x20));
            rages.Add(new RageDto(RageEnum.Rhyos, 15, 0x40));
            rages.Add(new RageDto(RageEnum.SrBehemoth, 15, 0x80));

            rages.Add(new RageDto(RageEnum.Vectaur, 16, 0x1));
            rages.Add(new RageDto(RageEnum.Wyvern, 16, 0x2));
            rages.Add(new RageDto(RageEnum.Zombone, 16, 0x4));
            rages.Add(new RageDto(RageEnum.Dragon, 16, 0x8));
            rages.Add(new RageDto(RageEnum.Brontaur, 16, 0x10));
            rages.Add(new RageDto(RageEnum.Allosaurus, 16, 0x20));
            rages.Add(new RageDto(RageEnum.Cirpius, 16, 0x40));
            rages.Add(new RageDto(RageEnum.Sprinter, 16, 0x80));

            rages.Add(new RageDto(RageEnum.Gobbler, 17, 0x1));
            rages.Add(new RageDto(RageEnum.Harpai, 17, 0x2));
            rages.Add(new RageDto(RageEnum.Gloomshell, 17, 0x4));
            rages.Add(new RageDto(RageEnum.Drop, 17, 0x8));
            rages.Add(new RageDto(RageEnum.MindCandy, 17, 0x10));
            rages.Add(new RageDto(RageEnum.WeedFeeder, 17, 0x20));
            rages.Add(new RageDto(RageEnum.Luridan, 17, 0x40));
            rages.Add(new RageDto(RageEnum.ToeCutter, 17, 0x80));

            rages.Add(new RageDto(RageEnum.OverGrunk, 18, 0x1));
            rages.Add(new RageDto(RageEnum.Exoray, 18, 0x2));
            rages.Add(new RageDto(RageEnum.Crusher, 18, 0x4));
            rages.Add(new RageDto(RageEnum.Uroburos, 18, 0x8));
            rages.Add(new RageDto(RageEnum.Primordite, 18, 0x10));
            rages.Add(new RageDto(RageEnum.SkyCap, 18, 0x20));
            rages.Add(new RageDto(RageEnum.Cephaler, 18, 0x40));
            rages.Add(new RageDto(RageEnum.Maliga, 18, 0x80));

            rages.Add(new RageDto(RageEnum.GiganToad, 19, 0x1));
            rages.Add(new RageDto(RageEnum.Geckorex, 19, 0x2));
            rages.Add(new RageDto(RageEnum.Cluck, 19, 0x4));
            rages.Add(new RageDto(RageEnum.LandWorm, 19, 0x8));
            rages.Add(new RageDto(RageEnum.TestRider, 19, 0x10));
            rages.Add(new RageDto(RageEnum.PlutoArmor, 19, 0x20));
            rages.Add(new RageDto(RageEnum.TombThumb, 19, 0x40));
            rages.Add(new RageDto(RageEnum.HeavyArmor, 19, 0x80));

            rages.Add(new RageDto(RageEnum.Chaser, 20, 0x1));
            rages.Add(new RageDto(RageEnum.Scullion, 20, 0x2));
            rages.Add(new RageDto(RageEnum.Poplium, 20, 0x4));
            rages.Add(new RageDto(RageEnum.Intangir, 20, 0x8));
            rages.Add(new RageDto(RageEnum.Misfit, 20, 0x10));
            rages.Add(new RageDto(RageEnum.Eland, 20, 0x20));
            rages.Add(new RageDto(RageEnum.Enuo, 20, 0x40));
            rages.Add(new RageDto(RageEnum.DeepEye, 20, 0x80));

            rages.Add(new RageDto(RageEnum.GreaseMonk, 21, 0x1));
            rages.Add(new RageDto(RageEnum.NeckHunter, 21, 0x2));
            rages.Add(new RageDto(RageEnum.Grenadeb, 21, 0x4));
            rages.Add(new RageDto(RageEnum.Critic, 21, 0x8));
            rages.Add(new RageDto(RageEnum.PanDora, 21, 0x10));
            rages.Add(new RageDto(RageEnum.SoulDancer, 21, 0x20));
            rages.Add(new RageDto(RageEnum.Gigantos, 21, 0x40));
            rages.Add(new RageDto(RageEnum.MagRoaderA, 21, 0x80));

            rages.Add(new RageDto(RageEnum.SpekTor, 22, 0x1));
            rages.Add(new RageDto(RageEnum.Parasite, 22, 0x2));
            rages.Add(new RageDto(RageEnum.EarthGuard, 22, 0x4));
            rages.Add(new RageDto(RageEnum.Coelecite, 22, 0x8));
            rages.Add(new RageDto(RageEnum.Anemone, 22, 0x10));
            rages.Add(new RageDto(RageEnum.Hipocampus, 22, 0x20));
            rages.Add(new RageDto(RageEnum.Spectre, 22, 0x40));
            rages.Add(new RageDto(RageEnum.EvilOscar, 22, 0x80));

            rages.Add(new RageDto(RageEnum.Slurm, 23, 0x1));
            rages.Add(new RageDto(RageEnum.Latimeria, 23, 0x2));
            rages.Add(new RageDto(RageEnum.StillGoing, 23, 0x4));
            rages.Add(new RageDto(RageEnum.AlloVer, 23, 0x8));
            rages.Add(new RageDto(RageEnum.Phase, 23, 0x10));
            rages.Add(new RageDto(RageEnum.Outsider, 23, 0x20));
            rages.Add(new RageDto(RageEnum.Barbe, 23, 0x40));
            rages.Add(new RageDto(RageEnum.Parasoul, 23, 0x80));

            rages.Add(new RageDto(RageEnum.PmStalker, 24, 0x1));
            rages.Add(new RageDto(RageEnum.Hemophyte, 24, 0x2));
            rages.Add(new RageDto(RageEnum.SpForces, 24, 0x4));
            rages.Add(new RageDto(RageEnum.Nohrabbit, 24, 0x8));
            rages.Add(new RageDto(RageEnum.Wizard, 24, 0x10));
            rages.Add(new RageDto(RageEnum.Scrapper, 24, 0x20));
            rages.Add(new RageDto(RageEnum.Ceritops, 24, 0x40));
            rages.Add(new RageDto(RageEnum.Commando, 24, 0x80));

            rages.Add(new RageDto(RageEnum.Opinicus, 25, 0x1));
            rages.Add(new RageDto(RageEnum.Poppers, 25, 0x2));
            rages.Add(new RageDto(RageEnum.Lunaris, 25, 0x4));
            rages.Add(new RageDto(RageEnum.Garm, 25, 0x8));
            rages.Add(new RageDto(RageEnum.Vindr, 25, 0x10));
            rages.Add(new RageDto(RageEnum.Kiwak, 25, 0x20));
            rages.Add(new RageDto(RageEnum.Nastidon, 25, 0x40));
            rages.Add(new RageDto(RageEnum.Rinn, 25, 0x80));

            rages.Add(new RageDto(RageEnum.Insecare, 26, 0x1));
            rages.Add(new RageDto(RageEnum.Vermin, 26, 0x2));
            rages.Add(new RageDto(RageEnum.Mantodea, 26, 0x4));
            rages.Add(new RageDto(RageEnum.Bogy, 26, 0x8));
            rages.Add(new RageDto(RageEnum.Prussian, 26, 0x10));
            rages.Add(new RageDto(RageEnum.BlackDrgn, 26, 0x20));
            rages.Add(new RageDto(RageEnum.Adamanchyt, 26, 0x40));
            rages.Add(new RageDto(RageEnum.Dante, 26, 0x80));

            rages.Add(new RageDto(RageEnum.WireyDrgn, 27, 0x1));
            rages.Add(new RageDto(RageEnum.Dueller, 27, 0x2));
            rages.Add(new RageDto(RageEnum.Psycot, 27, 0x4));
            rages.Add(new RageDto(RageEnum.Muus, 27, 0x8));
            rages.Add(new RageDto(RageEnum.Karkass, 27, 0x10));
            rages.Add(new RageDto(RageEnum.Punisher, 27, 0x20));
            rages.Add(new RageDto(RageEnum.Balloon, 27, 0x40));
            rages.Add(new RageDto(RageEnum.Gabbldegak, 27, 0x80));

            rages.Add(new RageDto(RageEnum.GtBehemoth, 28, 0x1));
            rages.Add(new RageDto(RageEnum.Scorpion, 28, 0x2));
            rages.Add(new RageDto(RageEnum.ChaosDrgn, 28, 0x4));
            rages.Add(new RageDto(RageEnum.SpitFire, 28, 0x8));
            rages.Add(new RageDto(RageEnum.Vectagoyle, 28, 0x10));
            rages.Add(new RageDto(RageEnum.Lick, 28, 0x20));
            rages.Add(new RageDto(RageEnum.Osprey, 28, 0x40));
            rages.Add(new RageDto(RageEnum.MagRoaderB, 28, 0x80));

            rages.Add(new RageDto(RageEnum.Bug, 29, 0x1));
            rages.Add(new RageDto(RageEnum.SeaFlower, 29, 0x2));
            rages.Add(new RageDto(RageEnum.Fortis, 29, 0x4));
            rages.Add(new RageDto(RageEnum.Abolisher, 29, 0x8));
            rages.Add(new RageDto(RageEnum.Aquila, 29, 0x10));
            rages.Add(new RageDto(RageEnum.Junk, 29, 0x20));
            rages.Add(new RageDto(RageEnum.Mandrake, 29, 0x40));
            rages.Add(new RageDto(RageEnum.FirstClass, 29, 0x80));

            rages.Add(new RageDto(RageEnum.TapDancer, 30, 0x1));
            rages.Add(new RageDto(RageEnum.Necromancer, 30, 0x2));
            rages.Add(new RageDto(RageEnum.Borras, 30, 0x4));
            rages.Add(new RageDto(RageEnum.MagRoaderC, 4, 0x8));
            rages.Add(new RageDto(RageEnum.WildRat, 30, 0x10));
            rages.Add(new RageDto(RageEnum.GoldBear, 30, 0x20));
            rages.Add(new RageDto(RageEnum.Innoc, 30, 0x40));
            rages.Add(new RageDto(RageEnum.Trixter, 30, 0x80));

            rages.Add(new RageDto(RageEnum.RedWolf, 31, 0x1));
            rages.Add(new RageDto(RageEnum.Didalos, 31, 0x2));
            rages.Add(new RageDto(RageEnum.Woolly, 31, 0x4));
            rages.Add(new RageDto(RageEnum.Veteran, 31, 0x8));
            rages.Add(new RageDto(RageEnum.SkyBase, 31, 0x10));
            rages.Add(new RageDto(RageEnum.IronHitman, 31, 0x20));
            rages.Add(new RageDto(RageEnum.Io, 31, 0x40));

            // Determine the number of rage groups
            int currentGroup = -1;
            int groupCount = 0;
            foreach (RageDto rage in rages)
            {
                if (currentGroup != rage.Offset)
                {
                    currentGroup = rage.Offset;
                    ++groupCount;
                }
            }
            GroupCount = groupCount;
        }
    }
}
