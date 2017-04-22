namespace FF6_Save_Editor.DTOs
{
    class InventoryDto
    {
        /// <summary>
        /// Index 0: Count
        /// Index 1: Item Id
        /// </summary>
        private byte[][] items = new byte[255][];

        public InventoryDto()
        {
            for (uint i = 0; i < 255; ++i)
            {
                this.items[i] = new byte[2] { 0xFF, 0xFF };
            }
        }

        public byte GetCount(int index)
        {
            return this.items[index][0];
        }

        public byte GetItemId(int index)
        {
            return this.items[index][1];
        }

        public void SetItem(int index, byte itemId)
        {
            this.items[index][1] = itemId;
        }

        public void SetItem(int index, byte itemId, byte count)
        {
            this.items[index][1] = itemId;
            this.SetCount(index, count);
        }

        public void SetCount(int index, byte count)
        {
            this.items[index][0] = count;
        }
    }
}
