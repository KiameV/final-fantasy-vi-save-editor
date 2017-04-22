namespace FF6_Save_Editor.DTOs
{
    /// <summary>
    /// Defines misc stats
    /// </summary>
    class OtherStatsDto
    {
        public uint GP { get; set; }
        public uint Steps { get; set; }
        public uint PlayerXAxis { get; set; }
        public uint PlayerYAxis { get; set; }
        public string CurrentMap { get; set; }
        public uint SaveCount { get; set; }
        public byte NumberOfSaves { get; set; }
        public byte SaveCountRollover { get; set; }
        public byte MapX { get; set; }
        public byte MapY { get; set; }
        public byte AirshipXAxis { get; set; }
        public byte AirshipYAxis { get; set; }
        public bool IsAirshipVisible { get; set; }

        public uint SlotId { get; private set; }

        public void SetSlotId(uint slotId)
        {
            this.SlotId = slotId;
        }
    }
}
