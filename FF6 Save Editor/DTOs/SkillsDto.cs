using FF6_Save_Editor.Enums;
using System;

namespace FF6_Save_Editor.DTOs
{
    /// <summary>
    /// Defines the skills that are currently known including Blitzes, Lore, SwdTechs, Rages, and Dances
    /// </summary>
    class SkillsDto
    {
        public readonly KnownSkillsDto<BlitzEnum> KnownBlitzes;
        public readonly KnownSkillsDto<LoreEnum> KnownLore;
        public readonly KnownSkillsDto<SwdTechEnum> KnownSwdTechs;
        public readonly KnownSkillsDto<IConvertible> KnownRages;
        public readonly KnownSkillsDto<DanceEnum> KnownDances;

        public SkillsDto()
        {
            this.KnownBlitzes = new KnownSkillsDto<BlitzEnum>(typeof(BlitzEnum));
            this.KnownLore = new KnownSkillsDto<LoreEnum>(typeof(LoreEnum));
            this.KnownSwdTechs = new KnownSkillsDto<SwdTechEnum>(typeof(SwdTechEnum));
            this.KnownRages = new KnownSkillsDto<IConvertible>(typeof(RageEnum));
            this.KnownDances = new KnownSkillsDto<DanceEnum>(typeof(DanceEnum));
        }
    }
}
