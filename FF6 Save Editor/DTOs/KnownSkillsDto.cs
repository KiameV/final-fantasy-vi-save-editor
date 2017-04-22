using System;

namespace FF6_Save_Editor.DTOs
{
    class KnownSkillsDto<S> where S : IConvertible
    {
        class KnownSkillDto
        {
            public S Skill { get; private set; }
            public bool isKnown { get; set; }

            public KnownSkillDto(S skill)
            {
                this.Skill = skill;
                this.isKnown = false;
            }

            public override string ToString()
            {
                return Skill.ToString() + " - " + isKnown.ToString();
            }
        }

        private readonly KnownSkillDto[] knownSkills;

        public int Count { get { return this.knownSkills.Length; } }

        public KnownSkillsDto(S[] skills)
        {
            knownSkills = new KnownSkillDto[skills.Length];
            uint i = 0;
            foreach (S s in skills)
            {
                knownSkills[i] = new KnownSkillDto(s);
                ++i;
            }
        }

        public KnownSkillsDto(Type skillsEnum)
        {
            Array skills = Enum.GetValues(skillsEnum);
            knownSkills = new KnownSkillDto[skills.Length];
            uint i = 0;
            foreach (S s in skills)
            {
                knownSkills[i] = new KnownSkillDto(s);
                ++i;
            }
        }

        public bool IsSkillKnown(int i)
        {
            if (i > this.knownSkills.Length)
            {
                throw new IndexOutOfRangeException(i + " is too big for knownSkills. Max index is " + (this.knownSkills.Length - 1));
            }
            return this.knownSkills[i].isKnown;
        }

        public bool IsSkillKnown(S s)
        {
            foreach (KnownSkillDto skill in knownSkills)
            {
                if (skill.Skill.Equals(s))
                {
                    return skill.isKnown;
                }
            }
            return false;
        }

        public void SetSkillKnown(int i, bool isKnown)
        {
            if (i > this.knownSkills.Length)
            {
                throw new IndexOutOfRangeException(i + " is too big for knownSkills. Max index is " + (this.knownSkills.Length - 1));
            }
            this.knownSkills[i].isKnown = isKnown;
        }
    }
}
