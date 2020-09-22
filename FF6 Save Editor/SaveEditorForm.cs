using FF6_Save_Editor.DTOs;
using FF6_Save_Editor.Enums;
using FF6_Save_Editor.UI;
using FF6_Save_Editor.Util;
using FF6_Save_State_Editor.Enums;
using FF6_Save_State_Editor.Util;
using System;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text.RegularExpressions;
using System.Windows.Forms;

namespace FF6_Save_Editor
{
    public partial class SaveEditorForm : Form
    {
        /// <summary>
        /// SaveDto for for the opened file
        /// </summary>
        private SaveDto saveDto = new SaveDto();
        /// <summary>
        /// Array of known magic. This gets populated when a character is selected from the UI as this information is character-specific.
        /// </summary>
        private NumericUpDown[] percentKnownMagicUpDown;
        /// <summary>
        /// The currently selected character
        /// </summary>
        private CharacterDto selectedCharacter = null;

        public SaveEditorForm()
        {
            InitializeComponent();

            // Create a NumericUpDown for each knowable magic
            Array magic = Enum.GetValues(typeof(MagicEnum));
            magicGrid.RowCount = magic.Length + 1;
            percentKnownMagicUpDown = new NumericUpDown[magic.Length];
            this.magicGrid.ResumeLayout();
            foreach (MagicEnum m in magic)
            {
                NumericUpDown nud = new NumericUpDown();
                nud.Dock = System.Windows.Forms.DockStyle.Fill;
                nud.Location = new System.Drawing.Point(10, 0);
                nud.Margin = new System.Windows.Forms.Padding(10, 0, 10, 0);
                nud.Name = "nud" + m.ToString();
                nud.Size = new System.Drawing.Size(130, 20);
                nud.TabIndex = 0;
                nud.Minimum = 0;
                nud.Maximum = 255;
                nud.Increment = 1;
                percentKnownMagicUpDown[(uint)m] = nud;

                GroupBox gb = new GroupBox();
                gb.Controls.Add(nud);
                gb.Dock = System.Windows.Forms.DockStyle.Fill;
                gb.Location = new System.Drawing.Point(3, 223);
                gb.Name = "groupBox" + m.ToString();
                gb.Size = new System.Drawing.Size(144, 44);
                gb.TabIndex = 8;
                gb.TabStop = false;
                gb.Text = m.ToString();
                gb.Margin = new System.Windows.Forms.Padding(3, 3, 50, 3);

                this.magicGrid.Controls.Add(gb, 0, (int)m);

                nud.SuspendLayout();
                gb.SuspendLayout();
            }
            this.magicGrid.SuspendLayout();

            // Create the inventor data table
            DataTable dt = new DataTable();
            dt.Columns.Add("Count", typeof(string));
            dt.Columns.Add("Item Id", typeof(string));
            for (uint i = 0; i < 255; ++i)
            {
                dt.LoadDataRow(new String[] { "", "" }, true);
            }
            this.inventoryGrid.DataSource = dt;

            this.inventoryGrid.Columns[0].Width = 75;
            this.inventoryGrid.Columns[1].Width = 75;

            // Create all the help text
            this.WriteHelpText();

            // Populate the Rage ListBox with items
            foreach (IConvertible rage in Enum.GetValues(typeof(RageEnum)))
            {
                rageLB.Items.Add(rage);
            }

            // Populate the Lore ListBox with items
            foreach (LoreEnum lore in Enum.GetValues(typeof(LoreEnum)))
            {
                loreLB.Items.Add(lore);
            }

            // Populate the Esper ListBox with items
            foreach (EsperEnum esper in Enum.GetValues(typeof(EsperEnum)))
            {
                espersLB.Items.Add(esper);
            }

            // Default the search fields
            resultByHex.Text = ItemContainer.GetMatchFor(0);
        }

        /// <summary>
        /// Called when a different character is selected
        /// </summary>
        private void characterComboBox_CursorChanged(object sender, EventArgs e)
        {
            if (this.characterComboBox.SelectedItem == null)
            {
                this.characterComboBox.SelectedIndex = 0;
            }
            else
            {
                this.SaveSelectedCharacterInformation();

                this.selectedCharacter = this.saveDto.Characters.getCharacter(
                        (CharacterEnum)this.characterComboBox.SelectedItem);

                this.LoadCharacterInformation(selectedCharacter);
            }
        }

        /// <summary>
        /// Load the Character's information from SaveDto - call when a new character is selected
        /// </summary>
        /// <param name="selectedCharacter">The newly selected character</param>
        private void LoadCharacterInformation(CharacterDto selectedCharacter)
        {
            nameTB.Text = selectedCharacter.Name;
            level.Value = selectedCharacter.Level;

            currentHp.Value = selectedCharacter.CurrentHP;
            maxHP.Value = selectedCharacter.MaxHP;

            currentMP.Value = selectedCharacter.CurrentMP;
            maxMP.Value = selectedCharacter.MaxMP;

            vigor.Value = selectedCharacter.Vigor;
            stamina.Value = selectedCharacter.Stamina;
            speed.Value = selectedCharacter.Speed;
            magic.Value = selectedCharacter.Magic;
            exp.Value = selectedCharacter.Exp;

            foreach (MagicEnum magic in Enum.GetValues(typeof(MagicEnum)))
            {
                percentKnownMagicUpDown[(uint)magic].Value = selectedCharacter.KnownMagic.GetPercentLeared(magic);
            }

            weaponId.Value = selectedCharacter.Weapon;
            shieldId.Value = selectedCharacter.Shield;
            helmitId.Value = selectedCharacter.Helmet;
            armorId.Value = selectedCharacter.Armor;
            relic1Id.Value = selectedCharacter.Relic1;
            relic2Id.Value = selectedCharacter.Relic2;

            command1.SelectedItem = selectedCharacter.GetCommand(0);
            command2.SelectedItem = selectedCharacter.GetCommand(1);
            command3.SelectedItem = selectedCharacter.GetCommand(2);
            command4.SelectedItem = selectedCharacter.GetCommand(3);

            foreach (StatusEnum status in selectedCharacter.StatusEffetcs.GetStatusEffects())
            {
                for (int i = 0; i < statusEffectsLB.Items.Count; ++i)
                {
                    string item = statusEffectsLB.Items[i].ToString();
                    if (item.Equals(status.ToString()))
                    {
                        statusEffectsLB.SetItemChecked(i, true);
                    }
                }
            }
        }

        /// <summary>
        /// Save the selected character's information to SaveDto
        /// </summary>
        private void SaveSelectedCharacterInformation()
        {
            if (this.selectedCharacter != null)
            {
                this.selectedCharacter.Name = nameTB.Text.Trim();
                this.selectedCharacter.Level = (uint)level.Value;

                this.selectedCharacter.CurrentHP = (uint)currentHp.Value;
                this.selectedCharacter.MaxHP = (uint)maxHP.Value;

                this.selectedCharacter.CurrentMP = (uint)currentMP.Value;
                this.selectedCharacter.MaxMP = (uint)maxMP.Value;

                this.selectedCharacter.Vigor = (uint)vigor.Value;
                this.selectedCharacter.Stamina = (uint)stamina.Value;
                this.selectedCharacter.Speed = (uint)speed.Value;
                this.selectedCharacter.Magic = (uint)magic.Value;
                this.selectedCharacter.Exp = (uint)exp.Value;

                foreach (MagicEnum magic in Enum.GetValues(typeof(MagicEnum)))
                {
                    this.selectedCharacter.KnownMagic.SetPercentLearned(magic, (byte)percentKnownMagicUpDown[(uint)magic].Value);
                }

                this.selectedCharacter.Weapon = (uint)weaponId.Value;
                this.selectedCharacter.Shield = (uint)shieldId.Value;
                this.selectedCharacter.Helmet = (uint)helmitId.Value;
                this.selectedCharacter.Armor = (uint)armorId.Value;
                this.selectedCharacter.Relic1 = (uint)relic1Id.Value;
                this.selectedCharacter.Relic2 = (uint)relic2Id.Value;

                this.selectedCharacter.SetCommand(0, (CommandsEnum)command1.SelectedItem);
                this.selectedCharacter.SetCommand(1, (CommandsEnum)command2.SelectedItem);
                this.selectedCharacter.SetCommand(2, (CommandsEnum)command3.SelectedItem);
                this.selectedCharacter.SetCommand(3, (CommandsEnum)command4.SelectedItem);

                this.selectedCharacter.StatusEffetcs.SetStatusEffects(statusEffectsLB.CheckedItems.Cast<Object>());
            }
        }

        /// <summary>
        /// Called to load a save file
        /// </summary>
        /// <param name="saveFileType">The type of save file to load</param>
        private void LoadFile(SaveFileTypeEnum saveFileType)
        {
            if (HexFileUtil.FindHexFile(saveFileType, System.IO.FileAccess.Read))
            {
                Byte[] b;
                if (HexFileUtil.OpenFile(saveFileType, out b))
                {
                    // Default selected character to no character selected
                    this.selectedCharacter = null;

                    // Create a new SaveDto and populate it
                    this.saveDto = new SaveDto();
                    this.saveDto.SetSaveState(saveFileType, b);
                    SaveStateIOUtil.LoadSaveStateInformation(this.saveDto);

                    // Update the UI with the new data
                    // Update the Inventory tab
                    LoadInventoryToGridView(this.saveDto.Inventory);
                    // Update the Skills tab
                    LoadSkills(this.saveDto.Skills);
                    // Update the Espers tab
                    LoadEspers(this.saveDto.KnownEspers);
                    // Update the Other Stats tab
                    LoadOtherStats(this.saveDto.OtherStats);

                    // Default the selected character selection to none
                    characterComboBox_CursorChanged(null, null);

                    // Make the selection visible if this is the first time a save file has been loaded
                    this.mainTabs.Visible = true;
                    this.mainTabs.Enabled = true;
                    this.UpdateFormName();
                }
            }
        }

        /// <summary>
        /// Updates the form's name to reflect which/what kind of file is open/being edited
        /// </summary>
        private void UpdateFormName()
        {
            System.Text.StringBuilder sb = new System.Text.StringBuilder("FF6 Editor: ");
            if (this.saveDto.SaveFileType == SaveFileTypeEnum.SRMSlot1)
                sb.Append("Slot 1 ");
            else if (this.saveDto.SaveFileType == SaveFileTypeEnum.SRMSlot2)
                sb.Append("Slot 2 ");
            else if (this.saveDto.SaveFileType == SaveFileTypeEnum.SRMSlot3)
                sb.Append("Slot 3 ");
            else if (this.saveDto.SaveFileType == SaveFileTypeEnum.ZnesSaveState)
                sb.Append("ZNES ");
            else if (this.saveDto.SaveFileType == SaveFileTypeEnum.ZnesSaveState)
                sb.Append("Snes9x ");
            sb.Append("\"");
            sb.Append(System.IO.Path.GetFileName(HexFileUtil.SelectedFile));
            sb.Append("\"");
            this.Text = sb.ToString();
        }

        /// <summary>
        /// Update the Skills UI tab
        /// </summary>
        /// <param name="skills">Defines the known skills</param>
        private void LoadSkills(SkillsDto skills)
        {
            // Blitz
            for (int i = 0; i < skills.KnownBlitzes.Count; ++i)
            {
                this.blitzLB.SetItemChecked(i, skills.KnownBlitzes.IsSkillKnown(i));
            }

            // SwdTech
            for (int i = 0; i < skills.KnownSwdTechs.Count; ++i)
            {
                this.swdtechLB.SetItemChecked(i, skills.KnownSwdTechs.IsSkillKnown(i));
            }

            // Dance
            for (int i = 0; i < skills.KnownDances.Count; ++i)
            {
                this.danceLB.SetItemChecked(i, skills.KnownDances.IsSkillKnown(i));
            }

            // Lore
            for (int i = 0; i < skills.KnownLore.Count; ++i)
            {
                this.loreLB.SetItemChecked(i, skills.KnownLore.IsSkillKnown(i));
            }

            // Rage
            for (int i = 0; i < skills.KnownRages.Count; ++i)
            {
                this.rageLB.SetItemChecked(i, skills.KnownRages.IsSkillKnown(i));
            }
        }

        /// <summary>
        /// Update the Espers UI tab
        /// </summary>
        /// <param name="knownEspers">Defines the known espers</param>
        private void LoadEspers(KnownSkillsDto<EsperEnum> espers)
        {
            for (int i = 0; i < espers.Count; ++i)
            {
                this.espersLB.SetItemChecked(i, espers.IsSkillKnown(i));
            }
        }

        /// <summary>
        /// Update the Other Stats UI tab
        /// </summary>
        /// <param name="otherStats">Defines the other stats</param>
        private void LoadOtherStats(OtherStatsDto otherStats)
        {
            this.gold.Value = otherStats.GP;
            this.steps.Value = otherStats.Steps;
            this.numberOfSaves.Value = otherStats.NumberOfSaves;
            this.saveCountRollover.Value = otherStats.SaveCountRollover;
            this.mapXAxis.Value = otherStats.MapX;
            this.mapYAxis.Value = otherStats.MapY;
            this.airshipXAxis.Value = otherStats.AirshipXAxis;
            this.airshipYAxis.Value = otherStats.AirshipYAxis;
            this.isAirshipVisible.Checked = otherStats.IsAirshipVisible;
        }

        /// <summary>
        /// Update the Inventory UI tab
        /// </summary>
        /// <param name="inventory">Defines the inventory</param>
        private void LoadInventoryToGridView(InventoryDto inventory)
        {
            for (int i = 0; i < 255; ++i)
            {
                var row = this.inventoryGrid.Rows[i];
                row.Cells["Count"].Value = inventory.GetCount(i);
                row.Cells["Item Id"].Value = inventory.GetItemId(i).ToString("X2");
            }
        }

        /// <summary>
        /// Append the given text to the given RichTextBox in the specified style
        /// </summary>
        /// <param name="rtb">The RichTextBox to append the text to</param>
        /// <param name="style">The style the appended text should be in</param>
        /// <param name="text">The text to append</param>
        private void AppendText(RichTextBox rtb, FontStyle style, string text)
        {
            rtb.SelectionFont = new Font(rtb.Font, style);
            rtb.AppendText(text);
        }

        /// <summary>
        /// Called when the user selects Save As from the UI
        /// </summary>
        private void saveAs_Click(object sender, EventArgs e)
        {
            string oldFile = HexFileUtil.SelectedFile;
            if (HexFileUtil.FindHexFile(this.saveDto.SaveFileType, System.IO.FileAccess.Write))
            {
                save_click(null, null);
            }
        }

        /// <summary>
        /// Called when the user selects Save from the UI
        /// </summary>
        private void save_click(object sender, EventArgs e)
        {
            SaveSelectedCharacterInformation();
            SaveSkills(this.saveDto.Skills);
            SaveEspers(this.saveDto.KnownEspers);
            SaveOtherStats(this.saveDto.OtherStats);
            SaveInventory(this.saveDto.Inventory);

            SaveStateIOUtil.SaveSaveStateInformation(this.saveDto);
            if (HexFileUtil.SaveFile(this.saveDto.SaveFileType, this.saveDto.SaveByteArray))
            {
                this.UpdateFormName();
            }
        }

        /// <summary>
        /// Save the data on the Inventory UI tab
        /// </summary>
        /// <param name="inventory">Place to save the inventory to</param>
        private void SaveInventory(InventoryDto inventory)
        {
            for (int i = 0; i < 255; ++i)
            {
                var row = this.inventoryGrid.Rows[i];
                uint itemId = Convert.ToUInt32(row.Cells["Item Id"].Value.ToString(), 16);
                uint count = Convert.ToUInt32(row.Cells["Count"].Value.ToString(), 10);
                inventory.SetItem(i, (byte)itemId, (byte)count);
            }
        }

        /// <summary>
        /// Save the data on the Skills UI tab
        /// </summary>
        /// <param name="skills">Place to save the skills to</param>
        private void SaveSkills(SkillsDto skills)
        {
            // Blitz
            for (int i = 0; i < skills.KnownBlitzes.Count; ++i)
            {
                skills.KnownBlitzes.SetSkillKnown(i, this.blitzLB.GetItemChecked(i));
            }

            // SwdTech
            for (int i = 0; i < skills.KnownSwdTechs.Count; ++i)
            {
                skills.KnownSwdTechs.SetSkillKnown(i, this.swdtechLB.GetItemChecked(i));
            }

            // Blitz
            for (int i = 0; i < skills.KnownDances.Count; ++i)
            {
                skills.KnownDances.SetSkillKnown(i, this.danceLB.GetItemChecked(i));
            }

            // Lore
            for (int i = 0; i < skills.KnownLore.Count; ++i)
            {
                skills.KnownLore.SetSkillKnown(i, this.loreLB.GetItemChecked(i));
            }

            // Rage
            for (int i = 0; i < skills.KnownRages.Count; ++i)
            {
                skills.KnownRages.SetSkillKnown(i, this.rageLB.GetItemChecked(i));
            }
        }

        /// <summary>
        /// Save the data on the Espers UI tab
        /// </summary>
        /// <param name="espers">Place to save the espers to</param>
        private void SaveEspers(KnownSkillsDto<EsperEnum> espers)
        {
            for (int i = 0; i < espers.Count; ++i)
            {
                espers.SetSkillKnown(i, this.espersLB.GetItemChecked(i));
            }
        }

        /// <summary>
        /// Save the data on the Other Stats UI tab
        /// </summary>
        /// <param name="otherStats">Place to save the other stats to</param>
        private void SaveOtherStats(OtherStatsDto otherStats)
        {
            otherStats.GP = (uint)this.gold.Value;
            otherStats.Steps = (uint)this.steps.Value;
            otherStats.NumberOfSaves = (byte)this.numberOfSaves.Value;
            otherStats.SaveCountRollover = (byte)this.saveCountRollover.Value;
            otherStats.MapX = (byte)this.mapXAxis.Value;
            otherStats.MapY = (byte)this.mapYAxis.Value;
            otherStats.AirshipXAxis = (byte)this.airshipXAxis.Value;
            otherStats.AirshipYAxis = (byte)this.airshipYAxis.Value;
            otherStats.IsAirshipVisible = this.isAirshipVisible.Checked;
        }

        /// <summary>
        /// Called when a ListBox's selection changes
        /// </summary>
        /// <param name="sender">The ListBox which caused this call</param>
        private void listBox_SelectedIndexChanged(object sender, EventArgs e)
        {
            CheckedListBox lb = (CheckedListBox)sender;
            int i = lb.SelectedIndex;
            if (i != -1)
            {
                bool isChecked = lb.GetItemChecked(lb.SelectedIndex);
                lb.SetItemChecked(i, !isChecked);
                lb.ClearSelected();
            }
        }

        /// <summary>
        /// Called when the search-by-name is used
        /// </summary>
        private void findByName_TextChanged(object sender, EventArgs e)
        {
            resultByName.Text = ItemContainer.GetMatchFor(((TextBox)sender).Text);
        }

        /// <summary>
        /// Called when the search-by-hex-value is used
        /// </summary>
        private void findByHex_ValueChanged(object sender, EventArgs e)
        {
            resultByHex.Text = ItemContainer.GetMatchFor(Convert.ToByte(((HexNumericUpDown)sender).Value));
        }

        /// <summary>
        /// Helper method to write Help text to different text fields in the UI
        /// </summary>
        private void WriteHelpText()
        {
            this.magicHelpText.Text =
                "Values:" + Environment.NewLine +
                " 0 - 100 is the 'percent learned'" + Environment.NewLine +
                " 255 means the spell is learned and can be used.";

            AppendText(this.weaponShieldRTB, FontStyle.Regular, "FF - Empty" + Environment.NewLine);
            AppendText(this.helmetArmorRTB, FontStyle.Regular, "FF - Empty" + Environment.NewLine);
            AppendText(this.relicsRTB, FontStyle.Regular, "FF - Empty" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular, "FF - Empty" + Environment.NewLine);

            AppendText(this.weaponShieldRTB, FontStyle.Bold,
                "Dirk" + Environment.NewLine);
            AppendText(this.weaponShieldRTB, FontStyle.Regular,
                "00 - Dirk" + Environment.NewLine +
                "01 - MithrilKnife" + Environment.NewLine +
                "02 - Guardian" + Environment.NewLine +
                "03 - Air Lancet" + Environment.NewLine +
                "04 - ThiefKnife" + Environment.NewLine +
                "05 - Assassin" + Environment.NewLine +
                "06 - Man Eater" + Environment.NewLine +
                "07 - SwordBreaker" + Environment.NewLine +
                "08 - Graedus" + Environment.NewLine +
                "09 - ValiantKnife" + Environment.NewLine);

            AppendText(this.weaponShieldRTB, FontStyle.Bold,
                "Sword" + Environment.NewLine);
            AppendText(this.weaponShieldRTB, FontStyle.Regular,
                "0A - MithrilBlade" + Environment.NewLine +
                "0B - RegalCutlass" + Environment.NewLine +
                "0C - Rune Edge" + Environment.NewLine +
                "0D - Flame Sabre" + Environment.NewLine +
                "0E - Blizzard" + Environment.NewLine +
                "0F - ThunderBlade" + Environment.NewLine +
                "10 - Epee" + Environment.NewLine +
                "11 - Break Blade" + Environment.NewLine +
                "12 - Drainer" + Environment.NewLine +
                "13 - Enhancer" + Environment.NewLine +
                "14 - Crystal" + Environment.NewLine +
                "15 - Falchion" + Environment.NewLine +
                "16 - Soul Sabre" + Environment.NewLine +
                "17 - Ogre Nix" + Environment.NewLine +
                "18 - Excalibur" + Environment.NewLine +
                "19 - Scimiter" + Environment.NewLine +
                "1A - Illumina" + Environment.NewLine +
                "1B - Ragnarok" + Environment.NewLine +
                "1C - Atma Weapon" + Environment.NewLine);

            AppendText(this.weaponShieldRTB, FontStyle.Bold,
                "Lance" + Environment.NewLine);
            AppendText(this.weaponShieldRTB, FontStyle.Regular,
                "1D - Mithril Pike" + Environment.NewLine +
                "1E - Trident" + Environment.NewLine +
                "1F - Stout Spear" + Environment.NewLine +
                "20 - Partisan" + Environment.NewLine +
                "21 - Pearl Lance" + Environment.NewLine +
                "22 - Gold Lance" + Environment.NewLine +
                "23 - Aura Lance" + Environment.NewLine +
                "24 - Imp Halberd" + Environment.NewLine);

            AppendText(this.weaponShieldRTB, FontStyle.Bold,
                "Dirk" + Environment.NewLine);
            AppendText(this.weaponShieldRTB, FontStyle.Regular,
                "25 - Imperial" + Environment.NewLine +
                "26 - Kodachi" + Environment.NewLine +
                "27 - Blossom" + Environment.NewLine +
                "28 - Hardened" + Environment.NewLine +
                "29 - Striker" + Environment.NewLine +
                "2A - Stunner" + Environment.NewLine);

            AppendText(this.weaponShieldRTB, FontStyle.Bold,
                "Knife" + Environment.NewLine);
            AppendText(this.weaponShieldRTB, FontStyle.Regular,
                "2B - Ashura" + Environment.NewLine +
                "2C - Kotetsu" + Environment.NewLine +
                "2D - Forged" + Environment.NewLine +
                "2E - Tempest" + Environment.NewLine +
                "2F - Murasame" + Environment.NewLine +
                "30 - Aura" + Environment.NewLine +
                "31 - Strato" + Environment.NewLine +
                "32 - Sky Render" + Environment.NewLine);

            AppendText(this.weaponShieldRTB, FontStyle.Bold,
                "Rod" + Environment.NewLine);
            AppendText(this.weaponShieldRTB, FontStyle.Regular,
                "33 - Heal Rod" + Environment.NewLine +
                "34 - Mithril Rod" + Environment.NewLine +
                "35 - Fire Rod" + Environment.NewLine +
                "36 - Ice Rod" + Environment.NewLine +
                "37 - Thunder Rod" + Environment.NewLine +
                "38 - Poison Rod" + Environment.NewLine +
                "39 - Pearl Rod" + Environment.NewLine +
                "3A - Gravity Rod" + Environment.NewLine +
                "3B - Punisher" + Environment.NewLine +
                "3C - Magus Rod" + Environment.NewLine);

            AppendText(this.weaponShieldRTB, FontStyle.Bold,
                "Brush" + Environment.NewLine);
            AppendText(this.weaponShieldRTB, FontStyle.Regular,
                "3D - Chocobo Brsh" + Environment.NewLine +
                "3E - DaVinci Brsh" + Environment.NewLine +
                "3F - Magical Brsh" + Environment.NewLine +
                "40 - Rainbow Brsh" + Environment.NewLine);

            AppendText(this.weaponShieldRTB, FontStyle.Bold,
                "Stars" + Environment.NewLine);
            AppendText(this.weaponShieldRTB, FontStyle.Regular,
                "41 - Shuriken" + Environment.NewLine +
                "42 - Ninja Star" + Environment.NewLine +
                "43 - Tack Star" + Environment.NewLine);

            AppendText(this.weaponShieldRTB, FontStyle.Bold,
                "Special" + Environment.NewLine);
            AppendText(this.weaponShieldRTB, FontStyle.Regular,
                "44 - Flail" + Environment.NewLine +
                "45 - Full Moon" + Environment.NewLine +
                "46 - Morning Star" + Environment.NewLine +
                "47 - Boomerang" + Environment.NewLine +
                "48 - Rising Sun" + Environment.NewLine +
                "49 - Hawk Eye" + Environment.NewLine +
                "4A - Bone Club" + Environment.NewLine +
                "4B - Sniper" + Environment.NewLine +
                "4C - Wing Edge" + Environment.NewLine);

            AppendText(this.weaponShieldRTB, FontStyle.Bold,
                "Gambler" + Environment.NewLine);
            AppendText(this.weaponShieldRTB, FontStyle.Regular,
                "4D - Cards" + Environment.NewLine +
                "4E - Darts" + Environment.NewLine +
                "4F - Doom Darts" + Environment.NewLine +
                "50 - Trump" + Environment.NewLine +
                "51 - Dice" + Environment.NewLine +
                "52 - Fixed Dice" + Environment.NewLine);

            AppendText(this.weaponShieldRTB, FontStyle.Bold,
                "Claw" + Environment.NewLine);
            AppendText(this.weaponShieldRTB, FontStyle.Regular,
                "53 - MetalKnuckle" + Environment.NewLine +
                "54 - Mithril Claw" + Environment.NewLine +
                "55 - Kaiser" + Environment.NewLine +
                "56 - Poison Claw" + Environment.NewLine +
                "57 - Fire Knuckle" + Environment.NewLine +
                "58 - Dragon Claw" + Environment.NewLine +
                "59 - Tiger Fangs" + Environment.NewLine);

            AppendText(this.weaponShieldRTB, FontStyle.Bold,
                "Shield" + Environment.NewLine);
            AppendText(this.weaponShieldRTB, FontStyle.Regular,
                "5A - Buckler" + Environment.NewLine +
                "5B - Heavy Shld" + Environment.NewLine +
                "5C - Mithril Shld" + Environment.NewLine +
                "5D - Gold Shld" + Environment.NewLine +
                "5E - Aegis Shld" + Environment.NewLine +
                "5F - Diamond Shld" + Environment.NewLine +
                "60 - Flame Shld" + Environment.NewLine +
                "61 - Ice Shld" + Environment.NewLine +
                "62 - Thunder Shld" + Environment.NewLine +
                "63 - Crystal Shld" + Environment.NewLine +
                "64 - Genji Shld" + Environment.NewLine +
                "65 - TortoiseShld" + Environment.NewLine +
                "66 - Cursed Shld" + Environment.NewLine +
                "67 - Paladin Shld" + Environment.NewLine +
                "68 - Force Shld" + Environment.NewLine);

            AppendText(this.helmetArmorRTB, FontStyle.Bold,
                "Helmet" + Environment.NewLine);
            AppendText(this.helmetArmorRTB, FontStyle.Regular,
                "69 - Leather Hat" + Environment.NewLine +
                "6A - Hair Band" + Environment.NewLine +
                "6B - Plumed Hat" + Environment.NewLine +
                "6C - Beret" + Environment.NewLine +
                "6D - Magus Hat" + Environment.NewLine +
                "6E - Bandana" + Environment.NewLine +
                "6F - Iron Helmet" + Environment.NewLine +
                "70 - Coronet" + Environment.NewLine +
                "71 - Bard's Hat" + Environment.NewLine +
                "72 - Green Beret" + Environment.NewLine +
                "73 - Head Band" + Environment.NewLine +
                "74 - Mithril Helm" + Environment.NewLine +
                "75 - Tiara" + Environment.NewLine +
                "76 - Gold Helmet" + Environment.NewLine +
                "77 - Tiger Mask" + Environment.NewLine +
                "78 - Red Hat" + Environment.NewLine +
                "79 - Mystery Veil" + Environment.NewLine +
                "7A - Circlet" + Environment.NewLine +
                "7B - Regal Crown" + Environment.NewLine +
                "7C - Diamond Helm" + Environment.NewLine +
                "7D - Dark Hood" + Environment.NewLine +
                "7E - Crystal Helm" + Environment.NewLine +
                "7F - Oath Veil" + Environment.NewLine +
                "80 - Cat Hood" + Environment.NewLine +
                "81 - Genji Helmet" + Environment.NewLine +
                "82 - Thornlet" + Environment.NewLine +
                "83 - Titanium" + Environment.NewLine);

            AppendText(this.helmetArmorRTB, FontStyle.Bold,
                "Armor" + Environment.NewLine);
            AppendText(this.helmetArmorRTB, FontStyle.Regular,
                "84 - LeatherArmor" + Environment.NewLine +
                "85 - Cotton Robe" + Environment.NewLine +
                "86 - Kung Fu Suit" + Environment.NewLine +
                "87 - Iron Armor" + Environment.NewLine +
                "88 - Silk Robe" + Environment.NewLine +
                "89 - Mithril Vest" + Environment.NewLine +
                "8A - Ninja Gear" + Environment.NewLine +
                "8B - White Dress" + Environment.NewLine +
                "8C - Mithril Mail" + Environment.NewLine +
                "8D - Gaia Gear" + Environment.NewLine +
                "8E - Mirage Dress" + Environment.NewLine +
                "8F - Gold Armor" + Environment.NewLine +
                "90 - Power Sash" + Environment.NewLine +
                "91 - Light Robe" + Environment.NewLine +
                "92 - Diamond Vest" + Environment.NewLine +
                "93 - Red Jacket" + Environment.NewLine +
                "94 - Force Armor" + Environment.NewLine +
                "95 - DiamondArmor" + Environment.NewLine +
                "96 - Dark Gear" + Environment.NewLine +
                "97 - Tao Robe" + Environment.NewLine +
                "98 - Crystal Mail" + Environment.NewLine +
                "99 - Czarina Gown" + Environment.NewLine +
                "9A - Genji Armor" + Environment.NewLine +
                "9B - Imp's Armor" + Environment.NewLine +
                "9C - Minerva" + Environment.NewLine +
                "9D - Tabby Suit" + Environment.NewLine +
                "9E - Chocobo Suit" + Environment.NewLine +
                "9F - Moogle Suit" + Environment.NewLine +
                "A0 - Nutkin Suit" + Environment.NewLine +
                "A1 - BehemethSuit" + Environment.NewLine +
                "A2 - Snow Muffler" + Environment.NewLine);

            AppendText(this.relicsRTB, FontStyle.Bold,
                "Relic" + Environment.NewLine);
            AppendText(this.relicsRTB, FontStyle.Regular,
                "B0 - Goggles" + Environment.NewLine +
                "B1 - Star Pendant" + Environment.NewLine +
                "B2 - Peace Ring" + Environment.NewLine +
                "B3 - Amulet" + Environment.NewLine +
                "B4 - White Cape" + Environment.NewLine +
                "B5 - Jewel Ring" + Environment.NewLine +
                "B6 - Fair Ring" + Environment.NewLine +
                "B7 - Barrier Ring" + Environment.NewLine +
                "B8 - MithrilGlove" + Environment.NewLine +
                "B9 - Guard Ring" + Environment.NewLine +
                "BA - RunningShoes" + Environment.NewLine +
                "BB - Wall Ring" + Environment.NewLine +
                "BC - Cherub Down" + Environment.NewLine +
                "BD - Cure Ring" + Environment.NewLine +
                "BE - True Knight" + Environment.NewLine +
                "BF - DragoonBoots" + Environment.NewLine +
                "C0 - Zephyr Cape" + Environment.NewLine +
                "C1 - Czarina Ring" + Environment.NewLine +
                "C2 - Cursed Cing" + Environment.NewLine +
                "C3 - Earrings" + Environment.NewLine +
                "C4 - Atlas Armlet" + Environment.NewLine +
                "C5 - BlizzardRing" + Environment.NewLine +
                "C6 - Rage Ring" + Environment.NewLine +
                "C7 - Sneak Ring" + Environment.NewLine +
                "C8 - Pod Bracelet" + Environment.NewLine +
                "C9 - Hero Ring" + Environment.NewLine +
                "CA - Ribbon" + Environment.NewLine +
                "CB - Muscle Belt" + Environment.NewLine +
                "CC - Crystal Orb" + Environment.NewLine +
                "CD - Gold Hairpin" + Environment.NewLine +
                "CE - Economizer" + Environment.NewLine +
                "CF - Thief Glove" + Environment.NewLine +
                "D0 - Gauntlet" + Environment.NewLine +
                "D1 - Genji Glove" + Environment.NewLine +
                "D2 - Hyper Wrist" + Environment.NewLine +
                "D3 - Offering" + Environment.NewLine +
                "D4 - Beads" + Environment.NewLine +
                "D5 - Black Belt" + Environment.NewLine +
                "D6 - Coin Toss" + Environment.NewLine +
                "D7 - FakeMustache" + Environment.NewLine +
                "D8 - Gem Box" + Environment.NewLine +
                "D9 - Dragon Horn" + Environment.NewLine +
                "DA - Merit Award" + Environment.NewLine +
                "DB - Momento Ring" + Environment.NewLine +
                "DC - Safety Bit" + Environment.NewLine +
                "DD - Relic Ring" + Environment.NewLine +
                "DE - Moogle Charm" + Environment.NewLine +
                "DF - Charm Bangle" + Environment.NewLine +
                "E0 - Marvel Shoes" + Environment.NewLine +
                "E1 - Back Gaurd" + Environment.NewLine +
                "E2 - Gale Hairpin" + Environment.NewLine +
                "E3 - Sniper Sight" + Environment.NewLine +
                "E4 - Exp.Egg" + Environment.NewLine +
                "E5 - Tintinabar" + Environment.NewLine +
                "E6 - Sprint Shoes" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Dirk" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "00 - Dirk" + Environment.NewLine +
                "01 - MithrilKnife" + Environment.NewLine +
                "02 - Guardian" + Environment.NewLine +
                "03 - Air Lancet" + Environment.NewLine +
                "04 - ThiefKnife" + Environment.NewLine +
                "05 - Assassin" + Environment.NewLine +
                "06 - Man Eater" + Environment.NewLine +
                "07 - SwordBreaker" + Environment.NewLine +
                "08 - Graedus" + Environment.NewLine +
                "09 - ValiantKnife" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Sword" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "0A - MithrilBlade" + Environment.NewLine +
                "0B - RegalCutlass" + Environment.NewLine +
                "0C - Rune Edge" + Environment.NewLine +
                "0D - Flame Sabre" + Environment.NewLine +
                "0E - Blizzard" + Environment.NewLine +
                "0F - ThunderBlade" + Environment.NewLine +
                "10 - Epee" + Environment.NewLine +
                "11 - Break Blade" + Environment.NewLine +
                "12 - Drainer" + Environment.NewLine +
                "13 - Enhancer" + Environment.NewLine +
                "14 - Crystal" + Environment.NewLine +
                "15 - Falchion" + Environment.NewLine +
                "16 - Soul Sabre" + Environment.NewLine +
                "17 - Ogre Nix" + Environment.NewLine +
                "18 - Excalibur" + Environment.NewLine +
                "19 - Scimiter" + Environment.NewLine +
                "1A - Illumina" + Environment.NewLine +
                "1B - Ragnarok" + Environment.NewLine +
                "1C - Atma Weapon" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Lance" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "1D - Mithril Pike" + Environment.NewLine +
                "1E - Trident" + Environment.NewLine +
                "1F - Stout Spear" + Environment.NewLine +
                "20 - Partisan" + Environment.NewLine +
                "21 - Pearl Lance" + Environment.NewLine +
                "22 - Gold Lance" + Environment.NewLine +
                "23 - Aura Lance" + Environment.NewLine +
                "24 - Imp Halberd" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Dirk" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "25 - Imperial" + Environment.NewLine +
                "26 - Kodachi" + Environment.NewLine +
                "27 - Blossom" + Environment.NewLine +
                "28 - Hardened" + Environment.NewLine +
                "29 - Striker" + Environment.NewLine +
                "2A - Stunner" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Knife" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "2B - Ashura" + Environment.NewLine +
                "2C - Kotetsu" + Environment.NewLine +
                "2D - Forged" + Environment.NewLine +
                "2E - Tempest" + Environment.NewLine +
                "2F - Murasame" + Environment.NewLine +
                "30 - Aura" + Environment.NewLine +
                "31 - Strato" + Environment.NewLine +
                "32 - Sky Render" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Rod" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "33 - Heal Rod" + Environment.NewLine +
                "34 - Mithril Rod" + Environment.NewLine +
                "35 - Fire Rod" + Environment.NewLine +
                "36 - Ice Rod" + Environment.NewLine +
                "37 - Thunder Rod" + Environment.NewLine +
                "38 - Poison Rod" + Environment.NewLine +
                "39 - Pearl Rod" + Environment.NewLine +
                "3A - Gravity Rod" + Environment.NewLine +
                "3B - Punisher" + Environment.NewLine +
                "3C - Magus Rod" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Brush" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "3D - Chocobo Brsh" + Environment.NewLine +
                "3E - DaVinci Brsh" + Environment.NewLine +
                "3F - Magical Brsh" + Environment.NewLine +
                "40 - Rainbow Brsh" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Stars" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "41 - Shuriken" + Environment.NewLine +
                "42 - Ninja Star" + Environment.NewLine +
                "43 - Tack Star" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Special" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "44 - Flail" + Environment.NewLine +
                "45 - Full Moon" + Environment.NewLine +
                "46 - Morning Star" + Environment.NewLine +
                "47 - Boomerang" + Environment.NewLine +
                "48 - Rising Sun" + Environment.NewLine +
                "49 - Hawk Eye" + Environment.NewLine +
                "4A - Bone Club" + Environment.NewLine +
                "4B - Sniper" + Environment.NewLine +
                "4C - Wing Edge" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Gambler" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "4D - Cards" + Environment.NewLine +
                "4E - Darts" + Environment.NewLine +
                "4F - Doom Darts" + Environment.NewLine +
                "50 - Trump" + Environment.NewLine +
                "51 - Dice" + Environment.NewLine +
                "52 - Fixed Dice" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Claw" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "53 - MetalKnuckle" + Environment.NewLine +
                "54 - Mithril Claw" + Environment.NewLine +
                "55 - Kaiser" + Environment.NewLine +
                "56 - Poison Claw" + Environment.NewLine +
                "57 - Fire Knuckle" + Environment.NewLine +
                "58 - Dragon Claw" + Environment.NewLine +
                "59 - Tiger Fangs" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Shield" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "5A - Buckler" + Environment.NewLine +
                "5B - Heavy Shld" + Environment.NewLine +
                "5C - Mithril Shld" + Environment.NewLine +
                "5D - Gold Shld" + Environment.NewLine +
                "5E - Aegis Shld" + Environment.NewLine +
                "5F - Diamond Shld" + Environment.NewLine +
                "60 - Flame Shld" + Environment.NewLine +
                "61 - Ice Shld" + Environment.NewLine +
                "62 - Thunder Shld" + Environment.NewLine +
                "63 - Crystal Shld" + Environment.NewLine +
                "64 - Genji Shld" + Environment.NewLine +
                "65 - TortoiseShld" + Environment.NewLine +
                "66 - Cursed Shld" + Environment.NewLine +
                "67 - Paladin Shld" + Environment.NewLine +
                "68 - Force Shld" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Helmet" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "69 - Leather Hat" + Environment.NewLine +
                "6A - Hair Band" + Environment.NewLine +
                "6B - Plumed Hat" + Environment.NewLine +
                "6C - Beret" + Environment.NewLine +
                "6D - Magus Hat" + Environment.NewLine +
                "6E - Bandana" + Environment.NewLine +
                "6F - Iron Helmet" + Environment.NewLine +
                "70 - Coronet" + Environment.NewLine +
                "71 - Bard's Hat" + Environment.NewLine +
                "72 - Green Beret" + Environment.NewLine +
                "73 - Head Band" + Environment.NewLine +
                "74 - Mithril Helm" + Environment.NewLine +
                "75 - Tiara" + Environment.NewLine +
                "76 - Gold Helmet" + Environment.NewLine +
                "77 - Tiger Mask" + Environment.NewLine +
                "78 - Red Hat" + Environment.NewLine +
                "79 - Mystery Veil" + Environment.NewLine +
                "7A - Circlet" + Environment.NewLine +
                "7B - Regal Crown" + Environment.NewLine +
                "7C - Diamond Helm" + Environment.NewLine +
                "7D - Dark Hood" + Environment.NewLine +
                "7E - Crystal Helm" + Environment.NewLine +
                "7F - Oath Veil" + Environment.NewLine +
                "80 - Cat Hood" + Environment.NewLine +
                "81 - Genji Helmet" + Environment.NewLine +
                "82 - Thornlet" + Environment.NewLine +
                "83 - Titanium" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Armor" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "84 - LeatherArmor" + Environment.NewLine +
                "85 - Cotton Robe" + Environment.NewLine +
                "86 - Kung Fu Suit" + Environment.NewLine +
                "87 - Iron Armor" + Environment.NewLine +
                "88 - Silk Robe" + Environment.NewLine +
                "89 - Mithril Vest" + Environment.NewLine +
                "8A - Ninja Gear" + Environment.NewLine +
                "8B - White Dress" + Environment.NewLine +
                "8C - Mithril Mail" + Environment.NewLine +
                "8D - Gaia Gear" + Environment.NewLine +
                "8E - Mirage Dress" + Environment.NewLine +
                "8F - Gold Armor" + Environment.NewLine +
                "90 - Power Sash" + Environment.NewLine +
                "91 - Light Robe" + Environment.NewLine +
                "92 - Diamond Vest" + Environment.NewLine +
                "93 - Red Jacket" + Environment.NewLine +
                "94 - Force Armor" + Environment.NewLine +
                "95 - DiamondArmor" + Environment.NewLine +
                "96 - Dark Gear" + Environment.NewLine +
                "97 - Tao Robe" + Environment.NewLine +
                "98 - Crystal Mail" + Environment.NewLine +
                "99 - Czarina Gown" + Environment.NewLine +
                "9A - Genji Armor" + Environment.NewLine +
                "9B - Imp's Armor" + Environment.NewLine +
                "9C - Minerva" + Environment.NewLine +
                "9D - Tabby Suit" + Environment.NewLine +
                "9E - Chocobo Suit" + Environment.NewLine +
                "9F - Moogle Suit" + Environment.NewLine +
                "A0 - Nutkin Suit" + Environment.NewLine +
                "A1 - BehemethSuit" + Environment.NewLine +
                "A2 - Snow Muffler" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Relic" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "B0 - Goggles" + Environment.NewLine +
                "B1 - Star Pendant" + Environment.NewLine +
                "B2 - Peace Ring" + Environment.NewLine +
                "B3 - Amulet" + Environment.NewLine +
                "B4 - White Cape" + Environment.NewLine +
                "B5 - Jewel Ring" + Environment.NewLine +
                "B6 - Fair Ring" + Environment.NewLine +
                "B7 - Barrier Ring" + Environment.NewLine +
                "B8 - MithrilGlove" + Environment.NewLine +
                "B9 - Guard Ring" + Environment.NewLine +
                "BA - RunningShoes" + Environment.NewLine +
                "BB - Wall Ring" + Environment.NewLine +
                "BC - Cherub Down" + Environment.NewLine +
                "BD - Cure Ring" + Environment.NewLine +
                "BE - True Knight" + Environment.NewLine +
                "BF - DragoonBoots" + Environment.NewLine +
                "C0 - Zephyr Cape" + Environment.NewLine +
                "C1 - Czarina Ring" + Environment.NewLine +
                "C2 - Cursed Cing" + Environment.NewLine +
                "C3 - Earrings" + Environment.NewLine +
                "C4 - Atlas Armlet" + Environment.NewLine +
                "C5 - BlizzardRing" + Environment.NewLine +
                "C6 - Rage Ring" + Environment.NewLine +
                "C7 - Sneak Ring" + Environment.NewLine +
                "C8 - Pod Bracelet" + Environment.NewLine +
                "C9 - Hero Ring" + Environment.NewLine +
                "CA - Ribbon" + Environment.NewLine +
                "CB - Muscle Belt" + Environment.NewLine +
                "CC - Crystal Orb" + Environment.NewLine +
                "CD - Gold Hairpin" + Environment.NewLine +
                "CE - Economizer" + Environment.NewLine +
                "CF - Thief Glove" + Environment.NewLine +
                "D0 - Gauntlet" + Environment.NewLine +
                "D1 - Genji Glove" + Environment.NewLine +
                "D2 - Hyper Wrist" + Environment.NewLine +
                "D3 - Offering" + Environment.NewLine +
                "D4 - Beads" + Environment.NewLine +
                "D5 - Black Belt" + Environment.NewLine +
                "D6 - Coin Toss" + Environment.NewLine +
                "D7 - FakeMustache" + Environment.NewLine +
                "D8 - Gem Box" + Environment.NewLine +
                "D9 - Dragon Horn" + Environment.NewLine +
                "DA - Merit Award" + Environment.NewLine +
                "DB - Momento Ring" + Environment.NewLine +
                "DC - Safety Bit" + Environment.NewLine +
                "DD - Relic Ring" + Environment.NewLine +
                "DE - Moogle Charm" + Environment.NewLine +
                "DF - Charm Bangle" + Environment.NewLine +
                "E0 - Marvel Shoes" + Environment.NewLine +
                "E1 - Back Gaurd" + Environment.NewLine +
                "E2 - Gale Hairpin" + Environment.NewLine +
                "E3 - Sniper Sight" + Environment.NewLine +
                "E4 - Exp.Egg" + Environment.NewLine +
                "E5 - Tintinabar" + Environment.NewLine +
                "E6 - Sprint Shoes" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Tool" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "A3 - NoiseBlaster" + Environment.NewLine +
                "A4 - Bio Blaster" + Environment.NewLine +
                "A5 - Flash" + Environment.NewLine +
                "A6 - Chain Saw" + Environment.NewLine +
                "A7 - Debilitator" + Environment.NewLine +
                "A8 - Drill" + Environment.NewLine +
                "A9 - Air Anchor" + Environment.NewLine +
                "AA - AutoCrossbow" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Skean" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "AB - Fire Skean" + Environment.NewLine +
                "AC - Water Edge" + Environment.NewLine +
                "AD - Bolt Edge" + Environment.NewLine +
                "AE - Inviz Edge" + Environment.NewLine +
                "AF - Shadow Edge" + Environment.NewLine);

            AppendText(this.inventoryRTB, FontStyle.Bold,
                "Miscellaneous" + Environment.NewLine);
            AppendText(this.inventoryRTB, FontStyle.Regular,
                "E7 - Rename Card" + Environment.NewLine +
                "E8 - Tonic" + Environment.NewLine +
                "E9 - Potion" + Environment.NewLine +
                "EA - X - Potion" + Environment.NewLine +
                "EB - Tincture" + Environment.NewLine +
                "EC - Ether" + Environment.NewLine +
                "ED - X - Ether" + Environment.NewLine +
                "EE - Elixir" + Environment.NewLine +
                "EF - Megalixir" + Environment.NewLine +
                "F0 - Fenix Down" + Environment.NewLine +
                "F1 - Revivify" + Environment.NewLine +
                "F2 - Antidote" + Environment.NewLine +
                "F3 - Eydrop" + Environment.NewLine +
                "F4 - Soft" + Environment.NewLine +
                "F5 - Remedy" + Environment.NewLine +
                "F6 - Sleeping Bag" + Environment.NewLine +
                "F7 - Tent" + Environment.NewLine +
                "F8 - Green Cherry" + Environment.NewLine +
                "F9 - Magicite" + Environment.NewLine +
                "FA - Super Ball" + Environment.NewLine +
                "FB - Echo Screen" + Environment.NewLine +
                "FC - Smoke Bomb" + Environment.NewLine +
                "FD - Warp Stone" + Environment.NewLine +
                "FE - Dried Meat" + Environment.NewLine);

            experienceTB.Text =
                "Level - Experience" + Environment.NewLine + 
                "01 - 0" + Environment.NewLine +
                "02 - 32" + Environment.NewLine +
                "03 - 96" + Environment.NewLine +
                "04 - 208" + Environment.NewLine +
                "05 - 400" + Environment.NewLine +
                "06 - 672" + Environment.NewLine +
                "07 - 1056" + Environment.NewLine +
                "08 - 1552" + Environment.NewLine +
                "09 - 2184" + Environment.NewLine +
                "10 - 2976" + Environment.NewLine +
                "11 - 3936" + Environment.NewLine +
                "12 - 5080" + Environment.NewLine +
                "13 - 6432" + Environment.NewLine +
                "14 - 7992" + Environment.NewLine +
                "15 - 9784" + Environment.NewLine +
                "16 - 11840" + Environment.NewLine +
                "17 - 14152" + Environment.NewLine +
                "18 - 16736" + Environment.NewLine +
                "19 - 19616" + Environment.NewLine +
                "20 - 22832" + Environment.NewLine +
                "21 - 26360" + Environment.NewLine +
                "22 - 30232" + Environment.NewLine +
                "23 - 24456" + Environment.NewLine +
                "24 - 39056" + Environment.NewLine +
                "25 - 44072" + Environment.NewLine +
                "26 - 49464" + Environment.NewLine +
                "27 - 55288" + Environment.NewLine +
                "28 - 61568" + Environment.NewLine +
                "29 - 68304" + Environment.NewLine +
                "30 - 75496" + Environment.NewLine +
                "31 - 93184" + Environment.NewLine +
                "32 - 91384" + Environment.NewLine +
                "33 - 100083" + Environment.NewLine +
                "34 - 108344" + Environment.NewLine +
                "35 - 119136" + Environment.NewLine +
                "36 - 129504" + Environment.NewLine +
                "37 - 140464" + Environment.NewLine +
                "38 - 152008" + Environment.NewLine +
                "39 - 164184" + Environment.NewLine +
                "40 - 176976" + Environment.NewLine +
                "41 - 190416" + Environment.NewLine +
                "42 - 204520" + Environment.NewLine +
                "43 - 219320" + Environment.NewLine +
                "44 - 234808" + Environment.NewLine +
                "45 - 251000" + Environment.NewLine +
                "46 - 267936" + Environment.NewLine +
                "47 - 285600" + Environment.NewLine +
                "48 - 304040" + Environment.NewLine +
                "49 - 323248" + Environment.NewLine +
                "50 - 343248" + Environment.NewLine +
                "51 - 364064" + Environment.NewLine +
                "52 - 385696" + Environment.NewLine +
                "53 - 408160" + Environment.NewLine +
                "54 - 431488" + Environment.NewLine +
                "55 - 455680" + Environment.NewLine +
                "56 - 480776" + Environment.NewLine +
                "57 - 506760" + Environment.NewLine +
                "58 - 533680" + Environment.NewLine +
                "59 - 561528" + Environment.NewLine +
                "60 - 590320" + Environment.NewLine +
                "61 - 620096" + Environment.NewLine +
                "62 - 650840" + Environment.NewLine +
                "63 - 682600" + Environment.NewLine +
                "64 - 715368" + Environment.NewLine +
                "65 - 749160" + Environment.NewLine +
                "66 - 784016" + Environment.NewLine +
                "67 - 819920" + Environment.NewLine +
                "68 - 856920" + Environment.NewLine +
                "69 - 895016" + Environment.NewLine +
                "70 - 934208" + Environment.NewLine +
                "71 - 974536" + Environment.NewLine +
                "72 - 1016000" + Environment.NewLine +
                "73 - 1058640" + Environment.NewLine +
                "74 - 1102456" + Environment.NewLine +
                "75 - 1147456" + Environment.NewLine +
                "76 - 1193648" + Environment.NewLine +
                "77 - 1241080" + Environment.NewLine +
                "78 - 1289744" + Environment.NewLine +
                "79 - 1339672" + Environment.NewLine +
                "80 - 1390872" + Environment.NewLine +
                "81 - 1443368" + Environment.NewLine +
                "82 - 1497160" + Environment.NewLine +
                "83 - 1553364" + Environment.NewLine +
                "84 - 1608712" + Environment.NewLine +
                "85 - 1666512" + Environment.NewLine +
                "86 - 1725688" + Environment.NewLine +
                "87 - 1786240" + Environment.NewLine +
                "88 - 1848184" + Environment.NewLine +
                "89 - 1911552" + Environment.NewLine +
                "90 - 1976352" + Environment.NewLine +
                "91 - 2042608" + Environment.NewLine +
                "92 - 2110320" + Environment.NewLine +
                "93 - 2179504" + Environment.NewLine +
                "94 - 2250192" + Environment.NewLine +
                "95 - 2322392" + Environment.NewLine +
                "96 - 2396128" + Environment.NewLine +
                "97 - 2471400" + Environment.NewLine +
                "98 - 2548224" + Environment.NewLine +
                "99 - 2637112";

            toolTip1.SetToolTip(this.saveCountRollover, "Counts the number of times the 'Number of Saves' has passed 255.\nExample: 2 in this field means 512 + the current value in 'Number of Saves'.");
        }

        /// <summary>
        /// Called when the user selectes Exit from the UI
        /// </summary>
        private void exit_Click(object sender, EventArgs e)
        {
            this.Close();
        }

        /// <summary>
        /// Called when the user selectes Open SRM 1 from the UI
        /// </summary>
        private void openSRM1_Click(object sender, EventArgs e)
        {
            this.LoadFile(SaveFileTypeEnum.SRMSlot1);
        }

        /// <summary>
        /// Called when the user selectes Open SRM 2 from the UI
        /// </summary>
        private void openSRM2_Click(object sender, EventArgs e)
        {
            this.LoadFile(SaveFileTypeEnum.SRMSlot2);
        }

        /// <summary>
        /// Called when the user selectes Open SRM 3 from the UI
        /// </summary>
        private void openSRM3_Click(object sender, EventArgs e)
        {
            this.LoadFile(SaveFileTypeEnum.SRMSlot3);
        }

        /// <summary>
        /// Called when the user selectes Open ZSNES from the UI
        /// </summary>
        private void openZnesSaveState_Click(object sender, EventArgs e)
        {
            this.LoadFile(SaveFileTypeEnum.ZnesSaveState);
        }

        /// <summary>
        /// Called when the user selectes Open Snes9x v1.5 from the UI
        /// </summary>
        private void openSnes9xSaveStatev15_Click(object sender, EventArgs e)
        {
            this.LoadFile(SaveFileTypeEnum.Snes9xSaveState15);
        }

        /// <summary>
        /// Called when the user selectes Open Snes9x v1.6 Offset 1 from the UI
        /// </summary>
        private void openSnes9xSaveStatev16_Click(object sender, EventArgs e)
        {
            this.LoadFile(SaveFileTypeEnum.Snes9xSaveState16);
        }

        /// <summary>
        /// Called when the user selectes the Select All Lore button
        /// </summary>
        private void SelectAllLore(object sender, EventArgs e)
        {
            SetAllChecks(this.loreLB, true);
        }

        /// <summary>
        /// Called when the user selectes the Deselect All Lore button
        /// </summary>
        private void DeselectAllLore(object sender, EventArgs e)
        {
            SetAllChecks(this.loreLB, false);
        }

        /// <summary>
        /// Called when the user selectes the Select All Rage button
        /// </summary>
        private void SelectAllRage(object sender, EventArgs e)
        {
            SetAllChecks(this.rageLB, true);
        }

        /// <summary>
        /// Called when the user selectes the Deselect All Rage button
        /// </summary>
        private void DeselectAllRage(object sender, EventArgs e)
        {
            SetAllChecks(this.rageLB, false);
        }

        /// <summary>
        /// Called when the user selectes the Select All Espers button
        /// </summary>
        private void SelectAllEspers(object sender, EventArgs e)
        {
            SetAllChecks(this.espersLB, true);
        }

        /// <summary>
        /// Called when the user selectes the Deselect All Espers button
        /// </summary>
        private void DeselectAllEspers(object sender, EventArgs e)
        {
            SetAllChecks(this.espersLB, false);
        }

        /// <summary>
        /// Helper method to set all check boxes in a CheckedListBox to checked or unchecked
        /// </summary>
        /// <param name="clb">The CheckedListBox to perform the action on</param>
        /// <param name="isChecked">True to check all boxes. False to uncheck all boxes.</param>
        private void SetAllChecks(CheckedListBox clb, bool isChecked)
        {
            for (int i = 0; i < clb.Items.Count; ++i)
            {
                clb.SetItemChecked(i, isChecked);
            }
        }
    }
}
