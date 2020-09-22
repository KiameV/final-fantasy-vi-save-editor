using FF6_Save_Editor.Enums;
using System;
using System.Linq;

namespace FF6_Save_Editor
{
    partial class SaveEditorForm
    {
        /// <summary>
        /// Required designer variable.
        /// </summary>
        private System.ComponentModel.IContainer components = null;

        /// <summary>
        /// Clean up any resources being used.
        /// </summary>
        /// <param name="disposing">true if managed resources should be disposed; otherwise, false.</param>
        protected override void Dispose(bool disposing)
        {
            if (disposing && (components != null))
            {
                components.Dispose();
            }
            base.Dispose(disposing);
        }

        #region Windows Form Designer generated code

        /// <summary>
        /// Required method for Designer support - do not modify
        /// the contents of this method with the code editor.
        /// </summary>
        private void InitializeComponent()
        {
            this.components = new System.ComponentModel.Container();
            this.menuStrip1 = new System.Windows.Forms.MenuStrip();
            this.fileToolStripMenuItem = new System.Windows.Forms.ToolStripMenuItem();
            this.toolStripMenuItem1 = new System.Windows.Forms.ToolStripMenuItem();
            this.slot1ToolStripMenuItem = new System.Windows.Forms.ToolStripMenuItem();
            this.slot2ToolStripMenuItem = new System.Windows.Forms.ToolStripMenuItem();
            this.slot3ToolStripMenuItem = new System.Windows.Forms.ToolStripMenuItem();
            this.openToolStripMenuItem = new System.Windows.Forms.ToolStripMenuItem();
            this.sRMToolStripMenuItem = new System.Windows.Forms.ToolStripMenuItem();
            this.snes9xToolStripMenuItem = new System.Windows.Forms.ToolStripMenuItem();
            this.snes9xV16ToolStripMenuItem = new System.Windows.Forms.ToolStripMenuItem();
            this.toolStripSeparator1 = new System.Windows.Forms.ToolStripSeparator();
            this.saveToolStripMenuItem = new System.Windows.Forms.ToolStripMenuItem();
            this.saveAsToolStripMenuItem = new System.Windows.Forms.ToolStripMenuItem();
            this.toolStripSeparator2 = new System.Windows.Forms.ToolStripSeparator();
            this.exitToolStripMenuItem = new System.Windows.Forms.ToolStripMenuItem();
            this.mainTabs = new System.Windows.Forms.TabControl();
            this.tabPage1 = new System.Windows.Forms.TabPage();
            this.tableLayoutPanel2 = new System.Windows.Forms.TableLayoutPanel();
            this.characterTabs = new System.Windows.Forms.TabControl();
            this.tabPage5 = new System.Windows.Forms.TabPage();
            this.tableLayoutPanel25 = new System.Windows.Forms.TableLayoutPanel();
            this.tableLayoutPanel3 = new System.Windows.Forms.TableLayoutPanel();
            this.groupBox34 = new System.Windows.Forms.GroupBox();
            this.nameTB = new System.Windows.Forms.TextBox();
            this.groupBox13 = new System.Windows.Forms.GroupBox();
            this.vigor = new System.Windows.Forms.NumericUpDown();
            this.groupBox14 = new System.Windows.Forms.GroupBox();
            this.stamina = new System.Windows.Forms.NumericUpDown();
            this.groupBox15 = new System.Windows.Forms.GroupBox();
            this.speed = new System.Windows.Forms.NumericUpDown();
            this.groupBox16 = new System.Windows.Forms.GroupBox();
            this.magic = new System.Windows.Forms.NumericUpDown();
            this.groupBox12 = new System.Windows.Forms.GroupBox();
            this.currentMP = new System.Windows.Forms.NumericUpDown();
            this.groupBox33 = new System.Windows.Forms.GroupBox();
            this.maxMP = new System.Windows.Forms.NumericUpDown();
            this.groupBox32 = new System.Windows.Forms.GroupBox();
            this.maxHP = new System.Windows.Forms.NumericUpDown();
            this.groupBox11 = new System.Windows.Forms.GroupBox();
            this.currentHp = new System.Windows.Forms.NumericUpDown();
            this.groupBox17 = new System.Windows.Forms.GroupBox();
            this.exp = new System.Windows.Forms.NumericUpDown();
            this.groupBox36 = new System.Windows.Forms.GroupBox();
            this.level = new System.Windows.Forms.NumericUpDown();
            this.experienceTB = new System.Windows.Forms.RichTextBox();
            this.tabPage6 = new System.Windows.Forms.TabPage();
            this.tableLayoutPanel7 = new System.Windows.Forms.TableLayoutPanel();
            this.groupBox18 = new System.Windows.Forms.GroupBox();
            this.magicGrid = new System.Windows.Forms.TableLayoutPanel();
            this.magicHelpText = new System.Windows.Forms.TextBox();
            this.tabPage7 = new System.Windows.Forms.TabPage();
            this.tableLayoutPanel13 = new System.Windows.Forms.TableLayoutPanel();
            this.tableLayoutPanel8 = new System.Windows.Forms.TableLayoutPanel();
            this.groupBox25 = new System.Windows.Forms.GroupBox();
            this.groupBox23 = new System.Windows.Forms.GroupBox();
            this.groupBox21 = new System.Windows.Forms.GroupBox();
            this.groupBox19 = new System.Windows.Forms.GroupBox();
            this.groupBox29 = new System.Windows.Forms.GroupBox();
            this.groupBox27 = new System.Windows.Forms.GroupBox();
            this.tableLayoutPanel14 = new System.Windows.Forms.TableLayoutPanel();
            this.groupBox28 = new System.Windows.Forms.GroupBox();
            this.weaponShieldRTB = new System.Windows.Forms.RichTextBox();
            this.groupBox30 = new System.Windows.Forms.GroupBox();
            this.helmetArmorRTB = new System.Windows.Forms.RichTextBox();
            this.groupBox31 = new System.Windows.Forms.GroupBox();
            this.relicsRTB = new System.Windows.Forms.RichTextBox();
            this.tabPage8 = new System.Windows.Forms.TabPage();
            this.tableLayoutPanel9 = new System.Windows.Forms.TableLayoutPanel();
            this.groupBox20 = new System.Windows.Forms.GroupBox();
            this.command1 = new System.Windows.Forms.ComboBox();
            this.groupBox22 = new System.Windows.Forms.GroupBox();
            this.command2 = new System.Windows.Forms.ComboBox();
            this.groupBox24 = new System.Windows.Forms.GroupBox();
            this.command3 = new System.Windows.Forms.ComboBox();
            this.groupBox26 = new System.Windows.Forms.GroupBox();
            this.command4 = new System.Windows.Forms.ComboBox();
            this.tabPage9 = new System.Windows.Forms.TabPage();
            this.tableLayoutPanel12 = new System.Windows.Forms.TableLayoutPanel();
            this.statusEffectsLB = new System.Windows.Forms.CheckedListBox();
            this.tableLayoutPanel4 = new System.Windows.Forms.TableLayoutPanel();
            this.groupBox10 = new System.Windows.Forms.GroupBox();
            this.characterComboBox = new System.Windows.Forms.ComboBox();
            this.tabPage2 = new System.Windows.Forms.TabPage();
            this.tableLayoutPanel5 = new System.Windows.Forms.TableLayoutPanel();
            this.inventoryGrid = new System.Windows.Forms.DataGridView();
            this.tableLayoutPanel17 = new System.Windows.Forms.TableLayoutPanel();
            this.groupBox4 = new System.Windows.Forms.GroupBox();
            this.tableLayoutPanel19 = new System.Windows.Forms.TableLayoutPanel();
            this.resultByName = new System.Windows.Forms.Label();
            this.findByName = new System.Windows.Forms.TextBox();
            this.inventoryRTB = new System.Windows.Forms.RichTextBox();
            this.groupBox3 = new System.Windows.Forms.GroupBox();
            this.tableLayoutPanel18 = new System.Windows.Forms.TableLayoutPanel();
            this.resultByHex = new System.Windows.Forms.Label();
            this.tabPage3 = new System.Windows.Forms.TabPage();
            this.tableLayoutPanel6 = new System.Windows.Forms.TableLayoutPanel();
            this.tableLayoutPanel23 = new System.Windows.Forms.TableLayoutPanel();
            this.groupBox5 = new System.Windows.Forms.GroupBox();
            this.danceLB = new System.Windows.Forms.CheckedListBox();
            this.groupBox35 = new System.Windows.Forms.GroupBox();
            this.swdtechLB = new System.Windows.Forms.CheckedListBox();
            this.groupBox37 = new System.Windows.Forms.GroupBox();
            this.blitzLB = new System.Windows.Forms.CheckedListBox();
            this.tableLayoutPanel24 = new System.Windows.Forms.TableLayoutPanel();
            this.groupBox39 = new System.Windows.Forms.GroupBox();
            this.tableLayoutPanel11 = new System.Windows.Forms.TableLayoutPanel();
            this.rageLB = new System.Windows.Forms.CheckedListBox();
            this.tableLayoutPanel16 = new System.Windows.Forms.TableLayoutPanel();
            this.rageDeselectAllButton = new System.Windows.Forms.Button();
            this.rageSelectAllButton = new System.Windows.Forms.Button();
            this.groupBox38 = new System.Windows.Forms.GroupBox();
            this.tableLayoutPanel10 = new System.Windows.Forms.TableLayoutPanel();
            this.loreLB = new System.Windows.Forms.CheckedListBox();
            this.tableLayoutPanel15 = new System.Windows.Forms.TableLayoutPanel();
            this.loreDeselectAllButton = new System.Windows.Forms.Button();
            this.loreSelectAllButton = new System.Windows.Forms.Button();
            this.tabPage10 = new System.Windows.Forms.TabPage();
            this.tableLayoutPanel20 = new System.Windows.Forms.TableLayoutPanel();
            this.tableLayoutPanel21 = new System.Windows.Forms.TableLayoutPanel();
            this.espersLB = new System.Windows.Forms.CheckedListBox();
            this.tableLayoutPanel22 = new System.Windows.Forms.TableLayoutPanel();
            this.esperDeselectAll = new System.Windows.Forms.Button();
            this.esperSelectAll = new System.Windows.Forms.Button();
            this.tabPage4 = new System.Windows.Forms.TabPage();
            this.tableLayoutPanel1 = new System.Windows.Forms.TableLayoutPanel();
            this.groupBoxAS = new System.Windows.Forms.GroupBox();
            this.airshipXAxis = new System.Windows.Forms.NumericUpDown();
            this.groupBox6 = new System.Windows.Forms.GroupBox();
            this.numberOfSaves = new System.Windows.Forms.NumericUpDown();
            this.groupBox2 = new System.Windows.Forms.GroupBox();
            this.steps = new System.Windows.Forms.NumericUpDown();
            this.groupBox1 = new System.Windows.Forms.GroupBox();
            this.gold = new System.Windows.Forms.NumericUpDown();
            this.groupBox7 = new System.Windows.Forms.GroupBox();
            this.saveCountRollover = new System.Windows.Forms.NumericUpDown();
            this.groupBox9 = new System.Windows.Forms.GroupBox();
            this.mapYAxis = new System.Windows.Forms.NumericUpDown();
            this.groupBox8 = new System.Windows.Forms.GroupBox();
            this.mapXAxis = new System.Windows.Forms.NumericUpDown();
            this.groupBoxASX = new System.Windows.Forms.GroupBox();
            this.airshipYAxis = new System.Windows.Forms.NumericUpDown();
            this.isAirshipVisible = new System.Windows.Forms.CheckBox();
            this.toolTip1 = new System.Windows.Forms.ToolTip(this.components);
            this.armorId = new FF6_Save_Editor.UI.HexNumericUpDown();
            this.helmitId = new FF6_Save_Editor.UI.HexNumericUpDown();
            this.shieldId = new FF6_Save_Editor.UI.HexNumericUpDown();
            this.weaponId = new FF6_Save_Editor.UI.HexNumericUpDown();
            this.relic2Id = new FF6_Save_Editor.UI.HexNumericUpDown();
            this.relic1Id = new FF6_Save_Editor.UI.HexNumericUpDown();
            this.findByHex = new FF6_Save_Editor.UI.HexNumericUpDown();
            this.menuStrip1.SuspendLayout();
            this.mainTabs.SuspendLayout();
            this.tabPage1.SuspendLayout();
            this.tableLayoutPanel2.SuspendLayout();
            this.characterTabs.SuspendLayout();
            this.tabPage5.SuspendLayout();
            this.tableLayoutPanel25.SuspendLayout();
            this.tableLayoutPanel3.SuspendLayout();
            this.groupBox34.SuspendLayout();
            this.groupBox13.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.vigor)).BeginInit();
            this.groupBox14.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.stamina)).BeginInit();
            this.groupBox15.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.speed)).BeginInit();
            this.groupBox16.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.magic)).BeginInit();
            this.groupBox12.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.currentMP)).BeginInit();
            this.groupBox33.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.maxMP)).BeginInit();
            this.groupBox32.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.maxHP)).BeginInit();
            this.groupBox11.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.currentHp)).BeginInit();
            this.groupBox17.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.exp)).BeginInit();
            this.groupBox36.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.level)).BeginInit();
            this.tabPage6.SuspendLayout();
            this.tableLayoutPanel7.SuspendLayout();
            this.groupBox18.SuspendLayout();
            this.tabPage7.SuspendLayout();
            this.tableLayoutPanel13.SuspendLayout();
            this.tableLayoutPanel8.SuspendLayout();
            this.groupBox25.SuspendLayout();
            this.groupBox23.SuspendLayout();
            this.groupBox21.SuspendLayout();
            this.groupBox19.SuspendLayout();
            this.groupBox29.SuspendLayout();
            this.groupBox27.SuspendLayout();
            this.tableLayoutPanel14.SuspendLayout();
            this.groupBox28.SuspendLayout();
            this.groupBox30.SuspendLayout();
            this.groupBox31.SuspendLayout();
            this.tabPage8.SuspendLayout();
            this.tableLayoutPanel9.SuspendLayout();
            this.groupBox20.SuspendLayout();
            this.groupBox22.SuspendLayout();
            this.groupBox24.SuspendLayout();
            this.groupBox26.SuspendLayout();
            this.tabPage9.SuspendLayout();
            this.tableLayoutPanel12.SuspendLayout();
            this.tableLayoutPanel4.SuspendLayout();
            this.groupBox10.SuspendLayout();
            this.tabPage2.SuspendLayout();
            this.tableLayoutPanel5.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.inventoryGrid)).BeginInit();
            this.tableLayoutPanel17.SuspendLayout();
            this.groupBox4.SuspendLayout();
            this.tableLayoutPanel19.SuspendLayout();
            this.groupBox3.SuspendLayout();
            this.tableLayoutPanel18.SuspendLayout();
            this.tabPage3.SuspendLayout();
            this.tableLayoutPanel6.SuspendLayout();
            this.tableLayoutPanel23.SuspendLayout();
            this.groupBox5.SuspendLayout();
            this.groupBox35.SuspendLayout();
            this.groupBox37.SuspendLayout();
            this.tableLayoutPanel24.SuspendLayout();
            this.groupBox39.SuspendLayout();
            this.tableLayoutPanel11.SuspendLayout();
            this.tableLayoutPanel16.SuspendLayout();
            this.groupBox38.SuspendLayout();
            this.tableLayoutPanel10.SuspendLayout();
            this.tableLayoutPanel15.SuspendLayout();
            this.tabPage10.SuspendLayout();
            this.tableLayoutPanel20.SuspendLayout();
            this.tableLayoutPanel21.SuspendLayout();
            this.tableLayoutPanel22.SuspendLayout();
            this.tabPage4.SuspendLayout();
            this.tableLayoutPanel1.SuspendLayout();
            this.groupBoxAS.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.airshipXAxis)).BeginInit();
            this.groupBox6.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.numberOfSaves)).BeginInit();
            this.groupBox2.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.steps)).BeginInit();
            this.groupBox1.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.gold)).BeginInit();
            this.groupBox7.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.saveCountRollover)).BeginInit();
            this.groupBox9.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.mapYAxis)).BeginInit();
            this.groupBox8.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.mapXAxis)).BeginInit();
            this.groupBoxASX.SuspendLayout();
            ((System.ComponentModel.ISupportInitialize)(this.airshipYAxis)).BeginInit();
            ((System.ComponentModel.ISupportInitialize)(this.armorId)).BeginInit();
            ((System.ComponentModel.ISupportInitialize)(this.helmitId)).BeginInit();
            ((System.ComponentModel.ISupportInitialize)(this.shieldId)).BeginInit();
            ((System.ComponentModel.ISupportInitialize)(this.weaponId)).BeginInit();
            ((System.ComponentModel.ISupportInitialize)(this.relic2Id)).BeginInit();
            ((System.ComponentModel.ISupportInitialize)(this.relic1Id)).BeginInit();
            ((System.ComponentModel.ISupportInitialize)(this.findByHex)).BeginInit();
            this.SuspendLayout();
            // 
            // menuStrip1
            // 
            this.menuStrip1.Items.AddRange(new System.Windows.Forms.ToolStripItem[] {
            this.fileToolStripMenuItem});
            this.menuStrip1.Location = new System.Drawing.Point(0, 0);
            this.menuStrip1.Name = "menuStrip1";
            this.menuStrip1.Size = new System.Drawing.Size(384, 24);
            this.menuStrip1.TabIndex = 0;
            this.menuStrip1.Text = "menuStrip1";
            // 
            // fileToolStripMenuItem
            // 
            this.fileToolStripMenuItem.DropDownItems.AddRange(new System.Windows.Forms.ToolStripItem[] {
            this.toolStripMenuItem1,
            this.openToolStripMenuItem,
            this.toolStripSeparator1,
            this.saveToolStripMenuItem,
            this.saveAsToolStripMenuItem,
            this.toolStripSeparator2,
            this.exitToolStripMenuItem});
            this.fileToolStripMenuItem.Name = "fileToolStripMenuItem";
            this.fileToolStripMenuItem.Size = new System.Drawing.Size(37, 20);
            this.fileToolStripMenuItem.Text = "File";
            // 
            // toolStripMenuItem1
            // 
            this.toolStripMenuItem1.DropDownItems.AddRange(new System.Windows.Forms.ToolStripItem[] {
            this.slot1ToolStripMenuItem,
            this.slot2ToolStripMenuItem,
            this.slot3ToolStripMenuItem});
            this.toolStripMenuItem1.Name = "toolStripMenuItem1";
            this.toolStripMenuItem1.Size = new System.Drawing.Size(159, 22);
            this.toolStripMenuItem1.Text = "Open SRM";
            // 
            // slot1ToolStripMenuItem
            // 
            this.slot1ToolStripMenuItem.Name = "slot1ToolStripMenuItem";
            this.slot1ToolStripMenuItem.Size = new System.Drawing.Size(103, 22);
            this.slot1ToolStripMenuItem.Text = "Slot 1";
            this.slot1ToolStripMenuItem.Click += new System.EventHandler(this.openSRM1_Click);
            // 
            // slot2ToolStripMenuItem
            // 
            this.slot2ToolStripMenuItem.Name = "slot2ToolStripMenuItem";
            this.slot2ToolStripMenuItem.Size = new System.Drawing.Size(103, 22);
            this.slot2ToolStripMenuItem.Text = "Slot 2";
            this.slot2ToolStripMenuItem.Click += new System.EventHandler(this.openSRM2_Click);
            // 
            // slot3ToolStripMenuItem
            // 
            this.slot3ToolStripMenuItem.Name = "slot3ToolStripMenuItem";
            this.slot3ToolStripMenuItem.Size = new System.Drawing.Size(103, 22);
            this.slot3ToolStripMenuItem.Text = "Slot 3";
            this.slot3ToolStripMenuItem.Click += new System.EventHandler(this.openSRM3_Click);
            // 
            // openToolStripMenuItem
            // 
            this.openToolStripMenuItem.DropDownItems.AddRange(new System.Windows.Forms.ToolStripItem[] {
            this.sRMToolStripMenuItem,
            this.snes9xToolStripMenuItem,
            this.snes9xV16ToolStripMenuItem});
            this.openToolStripMenuItem.Name = "openToolStripMenuItem";
            this.openToolStripMenuItem.Size = new System.Drawing.Size(159, 22);
            this.openToolStripMenuItem.Text = "Open Save State";
            // 
            // sRMToolStripMenuItem
            // 
            this.sRMToolStripMenuItem.Name = "sRMToolStripMenuItem";
            this.sRMToolStripMenuItem.Size = new System.Drawing.Size(134, 22);
            this.sRMToolStripMenuItem.Text = "ZNES";
            this.sRMToolStripMenuItem.Click += new System.EventHandler(this.openZnesSaveState_Click);
            // 
            // snes9xToolStripMenuItem
            // 
            this.snes9xToolStripMenuItem.Name = "snes9xToolStripMenuItem";
            this.snes9xToolStripMenuItem.Size = new System.Drawing.Size(134, 22);
            this.snes9xToolStripMenuItem.Text = "Snes9x v1.5";
            this.snes9xToolStripMenuItem.Click += new System.EventHandler(this.openSnes9xSaveStatev15_Click);
            // 
            // snes9xV16ToolStripMenuItem
            // 
            this.snes9xV16ToolStripMenuItem.Name = "snes9xV16ToolStripMenuItem";
            this.snes9xV16ToolStripMenuItem.Size = new System.Drawing.Size(134, 22);
            this.snes9xV16ToolStripMenuItem.Text = "Snes9x v1.6";
            this.snes9xV16ToolStripMenuItem.Click += new System.EventHandler(this.openSnes9xSaveStatev16_Click);
            // 
            // toolStripSeparator1
            // 
            this.toolStripSeparator1.Name = "toolStripSeparator1";
            this.toolStripSeparator1.Size = new System.Drawing.Size(156, 6);
            // 
            // saveToolStripMenuItem
            // 
            this.saveToolStripMenuItem.Name = "saveToolStripMenuItem";
            this.saveToolStripMenuItem.Size = new System.Drawing.Size(159, 22);
            this.saveToolStripMenuItem.Text = "Save";
            this.saveToolStripMenuItem.Click += new System.EventHandler(this.save_click);
            // 
            // saveAsToolStripMenuItem
            // 
            this.saveAsToolStripMenuItem.Name = "saveAsToolStripMenuItem";
            this.saveAsToolStripMenuItem.Size = new System.Drawing.Size(159, 22);
            this.saveAsToolStripMenuItem.Text = "Save As...";
            this.saveAsToolStripMenuItem.Click += new System.EventHandler(this.saveAs_Click);
            // 
            // toolStripSeparator2
            // 
            this.toolStripSeparator2.Name = "toolStripSeparator2";
            this.toolStripSeparator2.Size = new System.Drawing.Size(156, 6);
            // 
            // exitToolStripMenuItem
            // 
            this.exitToolStripMenuItem.Name = "exitToolStripMenuItem";
            this.exitToolStripMenuItem.Size = new System.Drawing.Size(159, 22);
            this.exitToolStripMenuItem.Text = "Exit";
            this.exitToolStripMenuItem.Click += new System.EventHandler(this.exit_Click);
            // 
            // mainTabs
            // 
            this.mainTabs.Controls.Add(this.tabPage1);
            this.mainTabs.Controls.Add(this.tabPage2);
            this.mainTabs.Controls.Add(this.tabPage3);
            this.mainTabs.Controls.Add(this.tabPage10);
            this.mainTabs.Controls.Add(this.tabPage4);
            this.mainTabs.Dock = System.Windows.Forms.DockStyle.Fill;
            this.mainTabs.Enabled = false;
            this.mainTabs.Location = new System.Drawing.Point(0, 24);
            this.mainTabs.Name = "mainTabs";
            this.mainTabs.SelectedIndex = 0;
            this.mainTabs.Size = new System.Drawing.Size(384, 487);
            this.mainTabs.TabIndex = 0;
            this.mainTabs.Visible = false;
            // 
            // tabPage1
            // 
            this.tabPage1.Controls.Add(this.tableLayoutPanel2);
            this.tabPage1.Location = new System.Drawing.Point(4, 22);
            this.tabPage1.Name = "tabPage1";
            this.tabPage1.Size = new System.Drawing.Size(376, 461);
            this.tabPage1.TabIndex = 0;
            this.tabPage1.Text = "Characters";
            this.tabPage1.UseVisualStyleBackColor = true;
            // 
            // tableLayoutPanel2
            // 
            this.tableLayoutPanel2.ColumnCount = 1;
            this.tableLayoutPanel2.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle());
            this.tableLayoutPanel2.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Absolute, 20F));
            this.tableLayoutPanel2.Controls.Add(this.characterTabs, 0, 1);
            this.tableLayoutPanel2.Controls.Add(this.tableLayoutPanel4, 0, 0);
            this.tableLayoutPanel2.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel2.Location = new System.Drawing.Point(0, 0);
            this.tableLayoutPanel2.Name = "tableLayoutPanel2";
            this.tableLayoutPanel2.RowCount = 2;
            this.tableLayoutPanel2.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel2.RowStyles.Add(new System.Windows.Forms.RowStyle());
            this.tableLayoutPanel2.Size = new System.Drawing.Size(376, 461);
            this.tableLayoutPanel2.TabIndex = 0;
            // 
            // characterTabs
            // 
            this.characterTabs.Controls.Add(this.tabPage5);
            this.characterTabs.Controls.Add(this.tabPage6);
            this.characterTabs.Controls.Add(this.tabPage7);
            this.characterTabs.Controls.Add(this.tabPage8);
            this.characterTabs.Controls.Add(this.tabPage9);
            this.characterTabs.Dock = System.Windows.Forms.DockStyle.Fill;
            this.characterTabs.Location = new System.Drawing.Point(3, 53);
            this.characterTabs.Name = "characterTabs";
            this.characterTabs.SelectedIndex = 0;
            this.characterTabs.Size = new System.Drawing.Size(370, 405);
            this.characterTabs.TabIndex = 2;
            // 
            // tabPage5
            // 
            this.tabPage5.Controls.Add(this.tableLayoutPanel25);
            this.tabPage5.Location = new System.Drawing.Point(4, 22);
            this.tabPage5.Name = "tabPage5";
            this.tabPage5.Padding = new System.Windows.Forms.Padding(3);
            this.tabPage5.Size = new System.Drawing.Size(362, 379);
            this.tabPage5.TabIndex = 0;
            this.tabPage5.Text = "Stats";
            this.tabPage5.UseVisualStyleBackColor = true;
            // 
            // tableLayoutPanel25
            // 
            this.tableLayoutPanel25.ColumnCount = 2;
            this.tableLayoutPanel25.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Absolute, 200F));
            this.tableLayoutPanel25.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel25.Controls.Add(this.tableLayoutPanel3, 0, 0);
            this.tableLayoutPanel25.Controls.Add(this.experienceTB, 1, 0);
            this.tableLayoutPanel25.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel25.Location = new System.Drawing.Point(3, 3);
            this.tableLayoutPanel25.Name = "tableLayoutPanel25";
            this.tableLayoutPanel25.RowCount = 1;
            this.tableLayoutPanel25.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel25.Size = new System.Drawing.Size(356, 373);
            this.tableLayoutPanel25.TabIndex = 1;
            // 
            // tableLayoutPanel3
            // 
            this.tableLayoutPanel3.ColumnCount = 2;
            this.tableLayoutPanel3.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Absolute, 100F));
            this.tableLayoutPanel3.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Absolute, 100F));
            this.tableLayoutPanel3.Controls.Add(this.groupBox34, 0, 0);
            this.tableLayoutPanel3.Controls.Add(this.groupBox13, 0, 4);
            this.tableLayoutPanel3.Controls.Add(this.groupBox14, 1, 4);
            this.tableLayoutPanel3.Controls.Add(this.groupBox15, 0, 5);
            this.tableLayoutPanel3.Controls.Add(this.groupBox16, 1, 5);
            this.tableLayoutPanel3.Controls.Add(this.groupBox12, 0, 3);
            this.tableLayoutPanel3.Controls.Add(this.groupBox33, 1, 3);
            this.tableLayoutPanel3.Controls.Add(this.groupBox32, 1, 2);
            this.tableLayoutPanel3.Controls.Add(this.groupBox11, 0, 2);
            this.tableLayoutPanel3.Controls.Add(this.groupBox17, 1, 1);
            this.tableLayoutPanel3.Controls.Add(this.groupBox36, 0, 1);
            this.tableLayoutPanel3.Location = new System.Drawing.Point(3, 3);
            this.tableLayoutPanel3.Name = "tableLayoutPanel3";
            this.tableLayoutPanel3.RowCount = 10;
            this.tableLayoutPanel3.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel3.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel3.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel3.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel3.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel3.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel3.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel3.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel3.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel3.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel3.Size = new System.Drawing.Size(194, 367);
            this.tableLayoutPanel3.TabIndex = 0;
            // 
            // groupBox34
            // 
            this.groupBox34.Controls.Add(this.nameTB);
            this.groupBox34.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox34.Location = new System.Drawing.Point(3, 3);
            this.groupBox34.Name = "groupBox34";
            this.groupBox34.Size = new System.Drawing.Size(94, 44);
            this.groupBox34.TabIndex = 3;
            this.groupBox34.TabStop = false;
            this.groupBox34.Text = "Name";
            // 
            // nameTB
            // 
            this.nameTB.Dock = System.Windows.Forms.DockStyle.Fill;
            this.nameTB.Location = new System.Drawing.Point(3, 16);
            this.nameTB.MaxLength = 6;
            this.nameTB.Name = "nameTB";
            this.nameTB.Size = new System.Drawing.Size(88, 20);
            this.nameTB.TabIndex = 0;
            // 
            // groupBox13
            // 
            this.groupBox13.Controls.Add(this.vigor);
            this.groupBox13.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox13.Location = new System.Drawing.Point(3, 203);
            this.groupBox13.Name = "groupBox13";
            this.groupBox13.Size = new System.Drawing.Size(94, 44);
            this.groupBox13.TabIndex = 10;
            this.groupBox13.TabStop = false;
            this.groupBox13.Text = "Vigor";
            // 
            // vigor
            // 
            this.vigor.Dock = System.Windows.Forms.DockStyle.Fill;
            this.vigor.Location = new System.Drawing.Point(3, 16);
            this.vigor.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.vigor.Name = "vigor";
            this.vigor.Size = new System.Drawing.Size(88, 20);
            this.vigor.TabIndex = 2;
            // 
            // groupBox14
            // 
            this.groupBox14.Controls.Add(this.stamina);
            this.groupBox14.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox14.Location = new System.Drawing.Point(103, 203);
            this.groupBox14.Name = "groupBox14";
            this.groupBox14.Size = new System.Drawing.Size(94, 44);
            this.groupBox14.TabIndex = 11;
            this.groupBox14.TabStop = false;
            this.groupBox14.Text = "Stamina";
            // 
            // stamina
            // 
            this.stamina.Dock = System.Windows.Forms.DockStyle.Fill;
            this.stamina.Location = new System.Drawing.Point(3, 16);
            this.stamina.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.stamina.Name = "stamina";
            this.stamina.Size = new System.Drawing.Size(88, 20);
            this.stamina.TabIndex = 2;
            // 
            // groupBox15
            // 
            this.groupBox15.Controls.Add(this.speed);
            this.groupBox15.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox15.Location = new System.Drawing.Point(3, 253);
            this.groupBox15.Name = "groupBox15";
            this.groupBox15.Size = new System.Drawing.Size(94, 44);
            this.groupBox15.TabIndex = 12;
            this.groupBox15.TabStop = false;
            this.groupBox15.Text = "Speed";
            // 
            // speed
            // 
            this.speed.Dock = System.Windows.Forms.DockStyle.Fill;
            this.speed.Location = new System.Drawing.Point(3, 16);
            this.speed.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.speed.Name = "speed";
            this.speed.Size = new System.Drawing.Size(88, 20);
            this.speed.TabIndex = 2;
            // 
            // groupBox16
            // 
            this.groupBox16.Controls.Add(this.magic);
            this.groupBox16.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox16.Location = new System.Drawing.Point(103, 253);
            this.groupBox16.Name = "groupBox16";
            this.groupBox16.Size = new System.Drawing.Size(94, 44);
            this.groupBox16.TabIndex = 13;
            this.groupBox16.TabStop = false;
            this.groupBox16.Text = "Magic";
            // 
            // magic
            // 
            this.magic.Dock = System.Windows.Forms.DockStyle.Fill;
            this.magic.Location = new System.Drawing.Point(3, 16);
            this.magic.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.magic.Name = "magic";
            this.magic.Size = new System.Drawing.Size(88, 20);
            this.magic.TabIndex = 2;
            // 
            // groupBox12
            // 
            this.groupBox12.Controls.Add(this.currentMP);
            this.groupBox12.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox12.Location = new System.Drawing.Point(3, 153);
            this.groupBox12.Name = "groupBox12";
            this.groupBox12.Size = new System.Drawing.Size(94, 44);
            this.groupBox12.TabIndex = 8;
            this.groupBox12.TabStop = false;
            this.groupBox12.Text = "Current MP";
            // 
            // currentMP
            // 
            this.currentMP.Dock = System.Windows.Forms.DockStyle.Fill;
            this.currentMP.Location = new System.Drawing.Point(3, 16);
            this.currentMP.Margin = new System.Windows.Forms.Padding(10, 0, 10, 0);
            this.currentMP.Maximum = new decimal(new int[] {
            16777216,
            0,
            0,
            0});
            this.currentMP.Name = "currentMP";
            this.currentMP.Size = new System.Drawing.Size(88, 20);
            this.currentMP.TabIndex = 10;
            // 
            // groupBox33
            // 
            this.groupBox33.Controls.Add(this.maxMP);
            this.groupBox33.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox33.Location = new System.Drawing.Point(103, 153);
            this.groupBox33.Name = "groupBox33";
            this.groupBox33.Size = new System.Drawing.Size(94, 44);
            this.groupBox33.TabIndex = 9;
            this.groupBox33.TabStop = false;
            this.groupBox33.Text = "Max MP";
            // 
            // maxMP
            // 
            this.maxMP.Dock = System.Windows.Forms.DockStyle.Fill;
            this.maxMP.Location = new System.Drawing.Point(3, 16);
            this.maxMP.Margin = new System.Windows.Forms.Padding(10, 0, 10, 0);
            this.maxMP.Maximum = new decimal(new int[] {
            16777216,
            0,
            0,
            0});
            this.maxMP.Name = "maxMP";
            this.maxMP.Size = new System.Drawing.Size(88, 20);
            this.maxMP.TabIndex = 3;
            // 
            // groupBox32
            // 
            this.groupBox32.Controls.Add(this.maxHP);
            this.groupBox32.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox32.Location = new System.Drawing.Point(103, 103);
            this.groupBox32.Name = "groupBox32";
            this.groupBox32.Size = new System.Drawing.Size(94, 44);
            this.groupBox32.TabIndex = 7;
            this.groupBox32.TabStop = false;
            this.groupBox32.Text = "Max HP";
            // 
            // maxHP
            // 
            this.maxHP.Dock = System.Windows.Forms.DockStyle.Fill;
            this.maxHP.Location = new System.Drawing.Point(3, 16);
            this.maxHP.Margin = new System.Windows.Forms.Padding(10, 0, 10, 0);
            this.maxHP.Maximum = new decimal(new int[] {
            16777216,
            0,
            0,
            0});
            this.maxHP.Name = "maxHP";
            this.maxHP.Size = new System.Drawing.Size(88, 20);
            this.maxHP.TabIndex = 2;
            // 
            // groupBox11
            // 
            this.groupBox11.Controls.Add(this.currentHp);
            this.groupBox11.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox11.Location = new System.Drawing.Point(3, 103);
            this.groupBox11.Name = "groupBox11";
            this.groupBox11.Size = new System.Drawing.Size(94, 44);
            this.groupBox11.TabIndex = 6;
            this.groupBox11.TabStop = false;
            this.groupBox11.Text = "Current HP";
            // 
            // currentHp
            // 
            this.currentHp.Dock = System.Windows.Forms.DockStyle.Fill;
            this.currentHp.Location = new System.Drawing.Point(3, 16);
            this.currentHp.Margin = new System.Windows.Forms.Padding(10, 0, 10, 0);
            this.currentHp.Maximum = new decimal(new int[] {
            16777216,
            0,
            0,
            0});
            this.currentHp.Name = "currentHp";
            this.currentHp.Size = new System.Drawing.Size(88, 20);
            this.currentHp.TabIndex = 10;
            // 
            // groupBox17
            // 
            this.groupBox17.Controls.Add(this.exp);
            this.groupBox17.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox17.Location = new System.Drawing.Point(103, 53);
            this.groupBox17.Name = "groupBox17";
            this.groupBox17.Size = new System.Drawing.Size(94, 44);
            this.groupBox17.TabIndex = 5;
            this.groupBox17.TabStop = false;
            this.groupBox17.Text = "Exp";
            // 
            // exp
            // 
            this.exp.Dock = System.Windows.Forms.DockStyle.Fill;
            this.exp.Location = new System.Drawing.Point(3, 16);
            this.exp.Maximum = new decimal(new int[] {
            16777216,
            0,
            0,
            0});
            this.exp.Name = "exp";
            this.exp.Size = new System.Drawing.Size(88, 20);
            this.exp.TabIndex = 2;
            // 
            // groupBox36
            // 
            this.groupBox36.Controls.Add(this.level);
            this.groupBox36.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox36.Location = new System.Drawing.Point(3, 53);
            this.groupBox36.Name = "groupBox36";
            this.groupBox36.Size = new System.Drawing.Size(94, 44);
            this.groupBox36.TabIndex = 4;
            this.groupBox36.TabStop = false;
            this.groupBox36.Text = "Level";
            // 
            // level
            // 
            this.level.Dock = System.Windows.Forms.DockStyle.Fill;
            this.level.Location = new System.Drawing.Point(3, 16);
            this.level.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.level.Name = "level";
            this.level.Size = new System.Drawing.Size(88, 20);
            this.level.TabIndex = 2;
            // 
            // experienceTB
            // 
            this.experienceTB.Dock = System.Windows.Forms.DockStyle.Fill;
            this.experienceTB.Location = new System.Drawing.Point(203, 3);
            this.experienceTB.Name = "experienceTB";
            this.experienceTB.ReadOnly = true;
            this.experienceTB.Size = new System.Drawing.Size(150, 367);
            this.experienceTB.TabIndex = 1;
            this.experienceTB.Text = "";
            // 
            // tabPage6
            // 
            this.tabPage6.Controls.Add(this.tableLayoutPanel7);
            this.tabPage6.Location = new System.Drawing.Point(4, 22);
            this.tabPage6.Name = "tabPage6";
            this.tabPage6.Padding = new System.Windows.Forms.Padding(3);
            this.tabPage6.Size = new System.Drawing.Size(362, 379);
            this.tabPage6.TabIndex = 1;
            this.tabPage6.Text = "Magic";
            this.tabPage6.UseVisualStyleBackColor = true;
            // 
            // tableLayoutPanel7
            // 
            this.tableLayoutPanel7.ColumnCount = 2;
            this.tableLayoutPanel7.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Absolute, 150F));
            this.tableLayoutPanel7.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel7.Controls.Add(this.groupBox18, 0, 0);
            this.tableLayoutPanel7.Controls.Add(this.magicHelpText, 1, 0);
            this.tableLayoutPanel7.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel7.Location = new System.Drawing.Point(3, 3);
            this.tableLayoutPanel7.Name = "tableLayoutPanel7";
            this.tableLayoutPanel7.RowCount = 1;
            this.tableLayoutPanel7.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel7.Size = new System.Drawing.Size(356, 373);
            this.tableLayoutPanel7.TabIndex = 2;
            // 
            // groupBox18
            // 
            this.groupBox18.Controls.Add(this.magicGrid);
            this.groupBox18.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox18.Location = new System.Drawing.Point(3, 3);
            this.groupBox18.Name = "groupBox18";
            this.groupBox18.Size = new System.Drawing.Size(144, 367);
            this.groupBox18.TabIndex = 1;
            this.groupBox18.TabStop = false;
            this.groupBox18.Text = "Known Spells";
            // 
            // magicGrid
            // 
            this.magicGrid.AutoScroll = true;
            this.magicGrid.ColumnCount = 1;
            this.magicGrid.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.magicGrid.Dock = System.Windows.Forms.DockStyle.Fill;
            this.magicGrid.Location = new System.Drawing.Point(3, 16);
            this.magicGrid.Name = "magicGrid";
            this.magicGrid.RowCount = 55;
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.magicGrid.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.magicGrid.Size = new System.Drawing.Size(138, 348);
            this.magicGrid.TabIndex = 0;
            // 
            // magicHelpText
            // 
            this.magicHelpText.Dock = System.Windows.Forms.DockStyle.Fill;
            this.magicHelpText.Location = new System.Drawing.Point(153, 3);
            this.magicHelpText.Multiline = true;
            this.magicHelpText.Name = "magicHelpText";
            this.magicHelpText.ReadOnly = true;
            this.magicHelpText.Size = new System.Drawing.Size(200, 367);
            this.magicHelpText.TabIndex = 2;
            // 
            // tabPage7
            // 
            this.tabPage7.Controls.Add(this.tableLayoutPanel13);
            this.tabPage7.Location = new System.Drawing.Point(4, 22);
            this.tabPage7.Name = "tabPage7";
            this.tabPage7.Padding = new System.Windows.Forms.Padding(3);
            this.tabPage7.Size = new System.Drawing.Size(362, 379);
            this.tabPage7.TabIndex = 2;
            this.tabPage7.Text = "Equipment";
            this.tabPage7.UseVisualStyleBackColor = true;
            // 
            // tableLayoutPanel13
            // 
            this.tableLayoutPanel13.ColumnCount = 2;
            this.tableLayoutPanel13.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Absolute, 150F));
            this.tableLayoutPanel13.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel13.Controls.Add(this.tableLayoutPanel8, 0, 0);
            this.tableLayoutPanel13.Controls.Add(this.tableLayoutPanel14, 1, 0);
            this.tableLayoutPanel13.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel13.Location = new System.Drawing.Point(3, 3);
            this.tableLayoutPanel13.Name = "tableLayoutPanel13";
            this.tableLayoutPanel13.RowCount = 1;
            this.tableLayoutPanel13.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel13.Size = new System.Drawing.Size(356, 373);
            this.tableLayoutPanel13.TabIndex = 1;
            // 
            // tableLayoutPanel8
            // 
            this.tableLayoutPanel8.ColumnCount = 1;
            this.tableLayoutPanel8.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Absolute, 150F));
            this.tableLayoutPanel8.Controls.Add(this.groupBox25, 0, 3);
            this.tableLayoutPanel8.Controls.Add(this.groupBox23, 0, 2);
            this.tableLayoutPanel8.Controls.Add(this.groupBox21, 0, 1);
            this.tableLayoutPanel8.Controls.Add(this.groupBox19, 0, 0);
            this.tableLayoutPanel8.Controls.Add(this.groupBox29, 0, 6);
            this.tableLayoutPanel8.Controls.Add(this.groupBox27, 0, 5);
            this.tableLayoutPanel8.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel8.Location = new System.Drawing.Point(3, 3);
            this.tableLayoutPanel8.Name = "tableLayoutPanel8";
            this.tableLayoutPanel8.RowCount = 8;
            this.tableLayoutPanel8.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel8.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel8.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel8.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel8.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 20F));
            this.tableLayoutPanel8.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel8.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel8.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel8.Size = new System.Drawing.Size(144, 367);
            this.tableLayoutPanel8.TabIndex = 0;
            // 
            // groupBox25
            // 
            this.groupBox25.Controls.Add(this.armorId);
            this.groupBox25.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox25.Location = new System.Drawing.Point(3, 153);
            this.groupBox25.Name = "groupBox25";
            this.groupBox25.Size = new System.Drawing.Size(144, 44);
            this.groupBox25.TabIndex = 4;
            this.groupBox25.TabStop = false;
            this.groupBox25.Text = "Armor Id";
            // 
            // groupBox23
            // 
            this.groupBox23.Controls.Add(this.helmitId);
            this.groupBox23.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox23.Location = new System.Drawing.Point(3, 103);
            this.groupBox23.Name = "groupBox23";
            this.groupBox23.Size = new System.Drawing.Size(144, 44);
            this.groupBox23.TabIndex = 3;
            this.groupBox23.TabStop = false;
            this.groupBox23.Text = "Helmet Id";
            // 
            // groupBox21
            // 
            this.groupBox21.Controls.Add(this.shieldId);
            this.groupBox21.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox21.Location = new System.Drawing.Point(3, 53);
            this.groupBox21.Name = "groupBox21";
            this.groupBox21.Size = new System.Drawing.Size(144, 44);
            this.groupBox21.TabIndex = 2;
            this.groupBox21.TabStop = false;
            this.groupBox21.Text = "Shield Id";
            // 
            // groupBox19
            // 
            this.groupBox19.Controls.Add(this.weaponId);
            this.groupBox19.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox19.Location = new System.Drawing.Point(3, 3);
            this.groupBox19.Name = "groupBox19";
            this.groupBox19.Size = new System.Drawing.Size(144, 44);
            this.groupBox19.TabIndex = 1;
            this.groupBox19.TabStop = false;
            this.groupBox19.Text = "Weapon Id";
            // 
            // groupBox29
            // 
            this.groupBox29.Controls.Add(this.relic2Id);
            this.groupBox29.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox29.Location = new System.Drawing.Point(3, 273);
            this.groupBox29.Name = "groupBox29";
            this.groupBox29.Size = new System.Drawing.Size(144, 44);
            this.groupBox29.TabIndex = 6;
            this.groupBox29.TabStop = false;
            this.groupBox29.Text = "Relic 2 Id";
            // 
            // groupBox27
            // 
            this.groupBox27.Controls.Add(this.relic1Id);
            this.groupBox27.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox27.Location = new System.Drawing.Point(3, 223);
            this.groupBox27.Name = "groupBox27";
            this.groupBox27.Size = new System.Drawing.Size(144, 44);
            this.groupBox27.TabIndex = 5;
            this.groupBox27.TabStop = false;
            this.groupBox27.Text = "Relic 1 Id";
            // 
            // tableLayoutPanel14
            // 
            this.tableLayoutPanel14.ColumnCount = 1;
            this.tableLayoutPanel14.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel14.Controls.Add(this.groupBox28, 0, 0);
            this.tableLayoutPanel14.Controls.Add(this.groupBox30, 0, 1);
            this.tableLayoutPanel14.Controls.Add(this.groupBox31, 0, 2);
            this.tableLayoutPanel14.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel14.Location = new System.Drawing.Point(153, 3);
            this.tableLayoutPanel14.Name = "tableLayoutPanel14";
            this.tableLayoutPanel14.RowCount = 3;
            this.tableLayoutPanel14.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel14.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel14.RowStyles.Add(new System.Windows.Forms.RowStyle());
            this.tableLayoutPanel14.Size = new System.Drawing.Size(200, 367);
            this.tableLayoutPanel14.TabIndex = 1;
            // 
            // groupBox28
            // 
            this.groupBox28.Controls.Add(this.weaponShieldRTB);
            this.groupBox28.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox28.Location = new System.Drawing.Point(3, 3);
            this.groupBox28.Name = "groupBox28";
            this.groupBox28.Size = new System.Drawing.Size(194, 124);
            this.groupBox28.TabIndex = 0;
            this.groupBox28.TabStop = false;
            this.groupBox28.Text = "Weapons/Shields";
            // 
            // weaponShieldRTB
            // 
            this.weaponShieldRTB.Dock = System.Windows.Forms.DockStyle.Fill;
            this.weaponShieldRTB.Location = new System.Drawing.Point(3, 16);
            this.weaponShieldRTB.Name = "weaponShieldRTB";
            this.weaponShieldRTB.ReadOnly = true;
            this.weaponShieldRTB.Size = new System.Drawing.Size(188, 105);
            this.weaponShieldRTB.TabIndex = 1;
            this.weaponShieldRTB.Text = "";
            // 
            // groupBox30
            // 
            this.groupBox30.Controls.Add(this.helmetArmorRTB);
            this.groupBox30.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox30.Location = new System.Drawing.Point(3, 133);
            this.groupBox30.Name = "groupBox30";
            this.groupBox30.Size = new System.Drawing.Size(194, 124);
            this.groupBox30.TabIndex = 1;
            this.groupBox30.TabStop = false;
            this.groupBox30.Text = "Helmets/Armor";
            // 
            // helmetArmorRTB
            // 
            this.helmetArmorRTB.Dock = System.Windows.Forms.DockStyle.Fill;
            this.helmetArmorRTB.Location = new System.Drawing.Point(3, 16);
            this.helmetArmorRTB.Name = "helmetArmorRTB";
            this.helmetArmorRTB.ReadOnly = true;
            this.helmetArmorRTB.Size = new System.Drawing.Size(188, 105);
            this.helmetArmorRTB.TabIndex = 0;
            this.helmetArmorRTB.Text = "";
            // 
            // groupBox31
            // 
            this.groupBox31.Controls.Add(this.relicsRTB);
            this.groupBox31.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox31.Location = new System.Drawing.Point(3, 263);
            this.groupBox31.Name = "groupBox31";
            this.groupBox31.Size = new System.Drawing.Size(194, 101);
            this.groupBox31.TabIndex = 2;
            this.groupBox31.TabStop = false;
            this.groupBox31.Text = "Relics";
            // 
            // relicsRTB
            // 
            this.relicsRTB.Dock = System.Windows.Forms.DockStyle.Fill;
            this.relicsRTB.Location = new System.Drawing.Point(3, 16);
            this.relicsRTB.Name = "relicsRTB";
            this.relicsRTB.ReadOnly = true;
            this.relicsRTB.Size = new System.Drawing.Size(188, 82);
            this.relicsRTB.TabIndex = 1;
            this.relicsRTB.Text = "";
            // 
            // tabPage8
            // 
            this.tabPage8.Controls.Add(this.tableLayoutPanel9);
            this.tabPage8.Location = new System.Drawing.Point(4, 22);
            this.tabPage8.Name = "tabPage8";
            this.tabPage8.Padding = new System.Windows.Forms.Padding(3);
            this.tabPage8.Size = new System.Drawing.Size(362, 379);
            this.tabPage8.TabIndex = 3;
            this.tabPage8.Text = "Commands";
            this.tabPage8.UseVisualStyleBackColor = true;
            // 
            // tableLayoutPanel9
            // 
            this.tableLayoutPanel9.ColumnCount = 2;
            this.tableLayoutPanel9.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Absolute, 150F));
            this.tableLayoutPanel9.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel9.Controls.Add(this.groupBox20, 0, 0);
            this.tableLayoutPanel9.Controls.Add(this.groupBox22, 0, 1);
            this.tableLayoutPanel9.Controls.Add(this.groupBox24, 0, 2);
            this.tableLayoutPanel9.Controls.Add(this.groupBox26, 0, 3);
            this.tableLayoutPanel9.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel9.Location = new System.Drawing.Point(3, 3);
            this.tableLayoutPanel9.Name = "tableLayoutPanel9";
            this.tableLayoutPanel9.RowCount = 5;
            this.tableLayoutPanel9.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel9.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel9.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel9.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel9.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel9.Size = new System.Drawing.Size(356, 373);
            this.tableLayoutPanel9.TabIndex = 0;
            // 
            // groupBox20
            // 
            this.groupBox20.Controls.Add(this.command1);
            this.groupBox20.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox20.Location = new System.Drawing.Point(3, 3);
            this.groupBox20.Name = "groupBox20";
            this.groupBox20.Size = new System.Drawing.Size(144, 44);
            this.groupBox20.TabIndex = 1;
            this.groupBox20.TabStop = false;
            this.groupBox20.Text = "Command 1";
            // 
            // command1
            // 
            this.command1.Dock = System.Windows.Forms.DockStyle.Fill;
            this.command1.FormattingEnabled = true;
            this.command1.Items.AddRange(new object[] {
            FF6_Save_Editor.Enums.CommandsEnum.Fight,
            FF6_Save_Editor.Enums.CommandsEnum.Item,
            FF6_Save_Editor.Enums.CommandsEnum.Magic,
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic))),
            FF6_Save_Editor.Enums.CommandsEnum.Revert,
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Revert))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Revert))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Revert))),
            FF6_Save_Editor.Enums.CommandsEnum.Throw,
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Revert | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            FF6_Save_Editor.Enums.CommandsEnum.Rage,
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Revert | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Throw | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Revert | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage)))});
            this.command1.Location = new System.Drawing.Point(3, 16);
            this.command1.Margin = new System.Windows.Forms.Padding(4);
            this.command1.Name = "command1";
            this.command1.Size = new System.Drawing.Size(138, 21);
            this.command1.TabIndex = 0;
            // 
            // groupBox22
            // 
            this.groupBox22.Controls.Add(this.command2);
            this.groupBox22.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox22.Location = new System.Drawing.Point(3, 53);
            this.groupBox22.Name = "groupBox22";
            this.groupBox22.Size = new System.Drawing.Size(144, 44);
            this.groupBox22.TabIndex = 2;
            this.groupBox22.TabStop = false;
            this.groupBox22.Text = "Command 2";
            // 
            // command2
            // 
            this.command2.Dock = System.Windows.Forms.DockStyle.Fill;
            this.command2.FormattingEnabled = true;
            this.command2.Items.AddRange(new object[] {
            FF6_Save_Editor.Enums.CommandsEnum.Fight,
            FF6_Save_Editor.Enums.CommandsEnum.Item,
            FF6_Save_Editor.Enums.CommandsEnum.Magic,
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic))),
            FF6_Save_Editor.Enums.CommandsEnum.Revert,
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Revert))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Revert))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Revert))),
            FF6_Save_Editor.Enums.CommandsEnum.Throw,
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Revert | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            FF6_Save_Editor.Enums.CommandsEnum.Rage,
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Revert | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Throw | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Revert | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage)))});
            this.command2.Location = new System.Drawing.Point(3, 16);
            this.command2.Margin = new System.Windows.Forms.Padding(4);
            this.command2.Name = "command2";
            this.command2.Size = new System.Drawing.Size(138, 21);
            this.command2.TabIndex = 2;
            // 
            // groupBox24
            // 
            this.groupBox24.Controls.Add(this.command3);
            this.groupBox24.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox24.Location = new System.Drawing.Point(3, 103);
            this.groupBox24.Name = "groupBox24";
            this.groupBox24.Size = new System.Drawing.Size(144, 44);
            this.groupBox24.TabIndex = 3;
            this.groupBox24.TabStop = false;
            this.groupBox24.Text = "Command 3";
            // 
            // command3
            // 
            this.command3.Dock = System.Windows.Forms.DockStyle.Fill;
            this.command3.FormattingEnabled = true;
            this.command3.Items.AddRange(new object[] {
            FF6_Save_Editor.Enums.CommandsEnum.Fight,
            FF6_Save_Editor.Enums.CommandsEnum.Item,
            FF6_Save_Editor.Enums.CommandsEnum.Magic,
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic))),
            FF6_Save_Editor.Enums.CommandsEnum.Revert,
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Revert))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Revert))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Revert))),
            FF6_Save_Editor.Enums.CommandsEnum.Throw,
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Revert | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            FF6_Save_Editor.Enums.CommandsEnum.Rage,
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Revert | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Throw | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Revert | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage)))});
            this.command3.Location = new System.Drawing.Point(3, 16);
            this.command3.Margin = new System.Windows.Forms.Padding(4);
            this.command3.Name = "command3";
            this.command3.Size = new System.Drawing.Size(138, 21);
            this.command3.TabIndex = 4;
            // 
            // groupBox26
            // 
            this.groupBox26.Controls.Add(this.command4);
            this.groupBox26.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox26.Location = new System.Drawing.Point(3, 153);
            this.groupBox26.Name = "groupBox26";
            this.groupBox26.Size = new System.Drawing.Size(144, 44);
            this.groupBox26.TabIndex = 4;
            this.groupBox26.TabStop = false;
            this.groupBox26.Text = "Command 4";
            // 
            // command4
            // 
            this.command4.Dock = System.Windows.Forms.DockStyle.Fill;
            this.command4.FormattingEnabled = true;
            this.command4.Items.AddRange(new object[] {
            FF6_Save_Editor.Enums.CommandsEnum.Fight,
            FF6_Save_Editor.Enums.CommandsEnum.Item,
            FF6_Save_Editor.Enums.CommandsEnum.Magic,
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic))),
            FF6_Save_Editor.Enums.CommandsEnum.Revert,
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Revert))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Revert))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Revert))),
            FF6_Save_Editor.Enums.CommandsEnum.Throw,
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Revert | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw))),
            FF6_Save_Editor.Enums.CommandsEnum.Rage,
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Revert | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((FF6_Save_Editor.Enums.CommandsEnum.Throw | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Magic) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)(((FF6_Save_Editor.Enums.CommandsEnum.Revert | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Item | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage))),
            ((FF6_Save_Editor.Enums.CommandsEnum)((((FF6_Save_Editor.Enums.CommandsEnum.Magic | FF6_Save_Editor.Enums.CommandsEnum.Revert) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Throw) 
                        | FF6_Save_Editor.Enums.CommandsEnum.Rage)))});
            this.command4.Location = new System.Drawing.Point(3, 16);
            this.command4.Margin = new System.Windows.Forms.Padding(4);
            this.command4.Name = "command4";
            this.command4.Size = new System.Drawing.Size(138, 21);
            this.command4.TabIndex = 6;
            // 
            // tabPage9
            // 
            this.tabPage9.Controls.Add(this.tableLayoutPanel12);
            this.tabPage9.Location = new System.Drawing.Point(4, 22);
            this.tabPage9.Name = "tabPage9";
            this.tabPage9.Padding = new System.Windows.Forms.Padding(3);
            this.tabPage9.Size = new System.Drawing.Size(362, 379);
            this.tabPage9.TabIndex = 4;
            this.tabPage9.Text = "Status Effects";
            this.tabPage9.UseVisualStyleBackColor = true;
            // 
            // tableLayoutPanel12
            // 
            this.tableLayoutPanel12.ColumnCount = 2;
            this.tableLayoutPanel12.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Absolute, 100F));
            this.tableLayoutPanel12.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel12.Controls.Add(this.statusEffectsLB, 0, 0);
            this.tableLayoutPanel12.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel12.Location = new System.Drawing.Point(3, 3);
            this.tableLayoutPanel12.Name = "tableLayoutPanel12";
            this.tableLayoutPanel12.RowCount = 2;
            this.tableLayoutPanel12.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 150F));
            this.tableLayoutPanel12.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel12.Size = new System.Drawing.Size(356, 373);
            this.tableLayoutPanel12.TabIndex = 0;
            // 
            // statusEffectsLB
            // 
            this.statusEffectsLB.Dock = System.Windows.Forms.DockStyle.Fill;
            this.statusEffectsLB.FormattingEnabled = true;
            this.statusEffectsLB.Items.AddRange(new object[] {
            FF6_Save_Editor.Enums.StatusEnum.Darkness,
            FF6_Save_Editor.Enums.StatusEnum.Zombie,
            FF6_Save_Editor.Enums.StatusEnum.Poison,
            FF6_Save_Editor.Enums.StatusEnum.Magitek,
            FF6_Save_Editor.Enums.StatusEnum.Invisible,
            FF6_Save_Editor.Enums.StatusEnum.Imp,
            FF6_Save_Editor.Enums.StatusEnum.Stone,
            FF6_Save_Editor.Enums.StatusEnum.Wounded,
            FF6_Save_Editor.Enums.StatusEnum.Float});
            this.statusEffectsLB.Location = new System.Drawing.Point(3, 3);
            this.statusEffectsLB.Name = "statusEffectsLB";
            this.statusEffectsLB.Size = new System.Drawing.Size(94, 144);
            this.statusEffectsLB.TabIndex = 0;
            this.statusEffectsLB.SelectedIndexChanged += new System.EventHandler(this.listBox_SelectedIndexChanged);
            // 
            // tableLayoutPanel4
            // 
            this.tableLayoutPanel4.ColumnCount = 2;
            this.tableLayoutPanel4.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Absolute, 100F));
            this.tableLayoutPanel4.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel4.Controls.Add(this.groupBox10, 0, 0);
            this.tableLayoutPanel4.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel4.Location = new System.Drawing.Point(3, 3);
            this.tableLayoutPanel4.Name = "tableLayoutPanel4";
            this.tableLayoutPanel4.RowCount = 1;
            this.tableLayoutPanel4.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel4.Size = new System.Drawing.Size(370, 44);
            this.tableLayoutPanel4.TabIndex = 2;
            // 
            // groupBox10
            // 
            this.groupBox10.Controls.Add(this.characterComboBox);
            this.groupBox10.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox10.Location = new System.Drawing.Point(3, 3);
            this.groupBox10.Name = "groupBox10";
            this.groupBox10.Size = new System.Drawing.Size(94, 38);
            this.groupBox10.TabIndex = 1;
            this.groupBox10.TabStop = false;
            this.groupBox10.Text = "Character";
            // 
            // characterComboBox
            // 
            this.characterComboBox.Dock = System.Windows.Forms.DockStyle.Fill;
            this.characterComboBox.FormattingEnabled = true;
            this.characterComboBox.Items.AddRange(new object[] {
            FF6_Save_Editor.Enums.CharacterEnum.Terra,
            FF6_Save_Editor.Enums.CharacterEnum.Locke,
            FF6_Save_Editor.Enums.CharacterEnum.Cyan,
            FF6_Save_Editor.Enums.CharacterEnum.Shadow,
            FF6_Save_Editor.Enums.CharacterEnum.Edgar,
            FF6_Save_Editor.Enums.CharacterEnum.Sabin,
            FF6_Save_Editor.Enums.CharacterEnum.Celes,
            FF6_Save_Editor.Enums.CharacterEnum.Strago,
            FF6_Save_Editor.Enums.CharacterEnum.Relm,
            FF6_Save_Editor.Enums.CharacterEnum.Setzer,
            FF6_Save_Editor.Enums.CharacterEnum.Mog,
            FF6_Save_Editor.Enums.CharacterEnum.Gau,
            FF6_Save_Editor.Enums.CharacterEnum.Gogo,
            FF6_Save_Editor.Enums.CharacterEnum.Umaru,
            FF6_Save_Editor.Enums.CharacterEnum.WedgeLeo,
            FF6_Save_Editor.Enums.CharacterEnum.Vicks});
            this.characterComboBox.Location = new System.Drawing.Point(3, 16);
            this.characterComboBox.Name = "characterComboBox";
            this.characterComboBox.Size = new System.Drawing.Size(88, 21);
            this.characterComboBox.TabIndex = 1;
            this.characterComboBox.SelectedIndexChanged += new System.EventHandler(this.characterComboBox_CursorChanged);
            // 
            // tabPage2
            // 
            this.tabPage2.Controls.Add(this.tableLayoutPanel5);
            this.tabPage2.Location = new System.Drawing.Point(4, 22);
            this.tabPage2.Name = "tabPage2";
            this.tabPage2.Size = new System.Drawing.Size(376, 461);
            this.tabPage2.TabIndex = 0;
            this.tabPage2.Text = "Inventory";
            this.tabPage2.UseVisualStyleBackColor = true;
            // 
            // tableLayoutPanel5
            // 
            this.tableLayoutPanel5.ColumnCount = 2;
            this.tableLayoutPanel5.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 60.10638F));
            this.tableLayoutPanel5.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 39.89362F));
            this.tableLayoutPanel5.Controls.Add(this.inventoryGrid, 0, 0);
            this.tableLayoutPanel5.Controls.Add(this.tableLayoutPanel17, 1, 0);
            this.tableLayoutPanel5.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel5.Location = new System.Drawing.Point(0, 0);
            this.tableLayoutPanel5.Name = "tableLayoutPanel5";
            this.tableLayoutPanel5.RowCount = 1;
            this.tableLayoutPanel5.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel5.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 461F));
            this.tableLayoutPanel5.Size = new System.Drawing.Size(376, 461);
            this.tableLayoutPanel5.TabIndex = 2;
            // 
            // inventoryGrid
            // 
            this.inventoryGrid.AllowUserToAddRows = false;
            this.inventoryGrid.AllowUserToDeleteRows = false;
            this.inventoryGrid.ColumnHeadersHeightSizeMode = System.Windows.Forms.DataGridViewColumnHeadersHeightSizeMode.AutoSize;
            this.inventoryGrid.Dock = System.Windows.Forms.DockStyle.Fill;
            this.inventoryGrid.Location = new System.Drawing.Point(3, 3);
            this.inventoryGrid.Name = "inventoryGrid";
            this.inventoryGrid.Size = new System.Drawing.Size(219, 455);
            this.inventoryGrid.TabIndex = 2;
            // 
            // tableLayoutPanel17
            // 
            this.tableLayoutPanel17.ColumnCount = 1;
            this.tableLayoutPanel17.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel17.Controls.Add(this.groupBox4, 0, 1);
            this.tableLayoutPanel17.Controls.Add(this.inventoryRTB, 0, 2);
            this.tableLayoutPanel17.Controls.Add(this.groupBox3, 0, 0);
            this.tableLayoutPanel17.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel17.Location = new System.Drawing.Point(228, 3);
            this.tableLayoutPanel17.Name = "tableLayoutPanel17";
            this.tableLayoutPanel17.RowCount = 3;
            this.tableLayoutPanel17.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel17.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel17.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel17.Size = new System.Drawing.Size(145, 455);
            this.tableLayoutPanel17.TabIndex = 3;
            // 
            // groupBox4
            // 
            this.groupBox4.Controls.Add(this.tableLayoutPanel19);
            this.groupBox4.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox4.Location = new System.Drawing.Point(3, 53);
            this.groupBox4.Name = "groupBox4";
            this.groupBox4.Size = new System.Drawing.Size(139, 44);
            this.groupBox4.TabIndex = 3;
            this.groupBox4.TabStop = false;
            this.groupBox4.Text = "Find By Name";
            // 
            // tableLayoutPanel19
            // 
            this.tableLayoutPanel19.ColumnCount = 2;
            this.tableLayoutPanel19.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 66.91729F));
            this.tableLayoutPanel19.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 33.08271F));
            this.tableLayoutPanel19.Controls.Add(this.resultByName, 1, 0);
            this.tableLayoutPanel19.Controls.Add(this.findByName, 0, 0);
            this.tableLayoutPanel19.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel19.Location = new System.Drawing.Point(3, 16);
            this.tableLayoutPanel19.Name = "tableLayoutPanel19";
            this.tableLayoutPanel19.RowCount = 1;
            this.tableLayoutPanel19.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel19.Size = new System.Drawing.Size(133, 25);
            this.tableLayoutPanel19.TabIndex = 0;
            // 
            // resultByName
            // 
            this.resultByName.AutoSize = true;
            this.resultByName.Dock = System.Windows.Forms.DockStyle.Fill;
            this.resultByName.Location = new System.Drawing.Point(90, 6);
            this.resultByName.Margin = new System.Windows.Forms.Padding(2, 6, 6, 6);
            this.resultByName.Name = "resultByName";
            this.resultByName.Size = new System.Drawing.Size(37, 13);
            this.resultByName.TabIndex = 2;
            // 
            // findByName
            // 
            this.findByName.Dock = System.Windows.Forms.DockStyle.Fill;
            this.findByName.Location = new System.Drawing.Point(3, 3);
            this.findByName.MaxLength = 20;
            this.findByName.Name = "findByName";
            this.findByName.Size = new System.Drawing.Size(82, 20);
            this.findByName.TabIndex = 0;
            this.findByName.TextChanged += new System.EventHandler(this.findByName_TextChanged);
            // 
            // inventoryRTB
            // 
            this.inventoryRTB.Dock = System.Windows.Forms.DockStyle.Fill;
            this.inventoryRTB.Location = new System.Drawing.Point(3, 103);
            this.inventoryRTB.Name = "inventoryRTB";
            this.inventoryRTB.ReadOnly = true;
            this.inventoryRTB.Size = new System.Drawing.Size(139, 349);
            this.inventoryRTB.TabIndex = 1;
            this.inventoryRTB.Text = "";
            // 
            // groupBox3
            // 
            this.groupBox3.Controls.Add(this.tableLayoutPanel18);
            this.groupBox3.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox3.Location = new System.Drawing.Point(3, 3);
            this.groupBox3.Name = "groupBox3";
            this.groupBox3.Size = new System.Drawing.Size(139, 44);
            this.groupBox3.TabIndex = 2;
            this.groupBox3.TabStop = false;
            this.groupBox3.Text = "Find By Hex";
            // 
            // tableLayoutPanel18
            // 
            this.tableLayoutPanel18.ColumnCount = 2;
            this.tableLayoutPanel18.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel18.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel18.Controls.Add(this.findByHex, 0, 0);
            this.tableLayoutPanel18.Controls.Add(this.resultByHex, 1, 0);
            this.tableLayoutPanel18.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel18.Location = new System.Drawing.Point(3, 16);
            this.tableLayoutPanel18.Name = "tableLayoutPanel18";
            this.tableLayoutPanel18.RowCount = 1;
            this.tableLayoutPanel18.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel18.Size = new System.Drawing.Size(133, 25);
            this.tableLayoutPanel18.TabIndex = 0;
            // 
            // resultByHex
            // 
            this.resultByHex.AutoSize = true;
            this.resultByHex.Dock = System.Windows.Forms.DockStyle.Fill;
            this.resultByHex.Location = new System.Drawing.Point(68, 6);
            this.resultByHex.Margin = new System.Windows.Forms.Padding(2, 6, 6, 6);
            this.resultByHex.Name = "resultByHex";
            this.resultByHex.Size = new System.Drawing.Size(59, 13);
            this.resultByHex.TabIndex = 1;
            // 
            // tabPage3
            // 
            this.tabPage3.Controls.Add(this.tableLayoutPanel6);
            this.tabPage3.Location = new System.Drawing.Point(4, 22);
            this.tabPage3.Name = "tabPage3";
            this.tabPage3.Padding = new System.Windows.Forms.Padding(3);
            this.tabPage3.Size = new System.Drawing.Size(376, 461);
            this.tabPage3.TabIndex = 1;
            this.tabPage3.Text = "Skills";
            this.tabPage3.UseVisualStyleBackColor = true;
            // 
            // tableLayoutPanel6
            // 
            this.tableLayoutPanel6.ColumnCount = 1;
            this.tableLayoutPanel6.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel6.Controls.Add(this.tableLayoutPanel23, 0, 0);
            this.tableLayoutPanel6.Controls.Add(this.tableLayoutPanel24, 0, 1);
            this.tableLayoutPanel6.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel6.Location = new System.Drawing.Point(3, 3);
            this.tableLayoutPanel6.Name = "tableLayoutPanel6";
            this.tableLayoutPanel6.RowCount = 2;
            this.tableLayoutPanel6.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 160F));
            this.tableLayoutPanel6.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel6.Size = new System.Drawing.Size(370, 455);
            this.tableLayoutPanel6.TabIndex = 0;
            // 
            // tableLayoutPanel23
            // 
            this.tableLayoutPanel23.ColumnCount = 3;
            this.tableLayoutPanel23.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 33F));
            this.tableLayoutPanel23.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 33F));
            this.tableLayoutPanel23.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 34F));
            this.tableLayoutPanel23.Controls.Add(this.groupBox5, 2, 0);
            this.tableLayoutPanel23.Controls.Add(this.groupBox35, 0, 0);
            this.tableLayoutPanel23.Controls.Add(this.groupBox37, 1, 0);
            this.tableLayoutPanel23.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel23.Location = new System.Drawing.Point(3, 3);
            this.tableLayoutPanel23.Name = "tableLayoutPanel23";
            this.tableLayoutPanel23.RowCount = 1;
            this.tableLayoutPanel23.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel23.Size = new System.Drawing.Size(364, 154);
            this.tableLayoutPanel23.TabIndex = 1;
            // 
            // groupBox5
            // 
            this.groupBox5.Controls.Add(this.danceLB);
            this.groupBox5.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox5.Location = new System.Drawing.Point(243, 3);
            this.groupBox5.Name = "groupBox5";
            this.groupBox5.Size = new System.Drawing.Size(118, 148);
            this.groupBox5.TabIndex = 2;
            this.groupBox5.TabStop = false;
            this.groupBox5.Text = "Dance";
            // 
            // danceLB
            // 
            this.danceLB.Dock = System.Windows.Forms.DockStyle.Fill;
            this.danceLB.FormattingEnabled = true;
            this.danceLB.Items.AddRange(new object[] {
            FF6_Save_Editor.Enums.DanceEnum.WindSong,
            FF6_Save_Editor.Enums.DanceEnum.ForestSuite,
            FF6_Save_Editor.Enums.DanceEnum.DesertAria,
            FF6_Save_Editor.Enums.DanceEnum.LoveSonata,
            FF6_Save_Editor.Enums.DanceEnum.EarthBlues,
            FF6_Save_Editor.Enums.DanceEnum.WaterRondo,
            FF6_Save_Editor.Enums.DanceEnum.DuskRequium,
            FF6_Save_Editor.Enums.DanceEnum.SnowmanJazz});
            this.danceLB.Location = new System.Drawing.Point(3, 16);
            this.danceLB.Name = "danceLB";
            this.danceLB.Size = new System.Drawing.Size(112, 129);
            this.danceLB.TabIndex = 1;
            this.danceLB.SelectedIndexChanged += new System.EventHandler(this.listBox_SelectedIndexChanged);
            // 
            // groupBox35
            // 
            this.groupBox35.Controls.Add(this.swdtechLB);
            this.groupBox35.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox35.Location = new System.Drawing.Point(3, 3);
            this.groupBox35.Name = "groupBox35";
            this.groupBox35.Size = new System.Drawing.Size(114, 148);
            this.groupBox35.TabIndex = 0;
            this.groupBox35.TabStop = false;
            this.groupBox35.Text = "SwdTech";
            // 
            // swdtechLB
            // 
            this.swdtechLB.Dock = System.Windows.Forms.DockStyle.Fill;
            this.swdtechLB.FormattingEnabled = true;
            this.swdtechLB.Items.AddRange(new object[] {
            FF6_Save_Editor.Enums.SwdTechEnum.Dispatch,
            FF6_Save_Editor.Enums.SwdTechEnum.Retort,
            FF6_Save_Editor.Enums.SwdTechEnum.Slash,
            FF6_Save_Editor.Enums.SwdTechEnum.QuadraSlam,
            FF6_Save_Editor.Enums.SwdTechEnum.Empowerer,
            FF6_Save_Editor.Enums.SwdTechEnum.Stunner,
            FF6_Save_Editor.Enums.SwdTechEnum.QuadraSlice,
            FF6_Save_Editor.Enums.SwdTechEnum.Cleave});
            this.swdtechLB.Location = new System.Drawing.Point(3, 16);
            this.swdtechLB.Name = "swdtechLB";
            this.swdtechLB.Size = new System.Drawing.Size(108, 129);
            this.swdtechLB.TabIndex = 0;
            this.swdtechLB.SelectedIndexChanged += new System.EventHandler(this.listBox_SelectedIndexChanged);
            // 
            // groupBox37
            // 
            this.groupBox37.Controls.Add(this.blitzLB);
            this.groupBox37.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox37.Location = new System.Drawing.Point(123, 3);
            this.groupBox37.Name = "groupBox37";
            this.groupBox37.Size = new System.Drawing.Size(114, 148);
            this.groupBox37.TabIndex = 1;
            this.groupBox37.TabStop = false;
            this.groupBox37.Text = "Blitz";
            // 
            // blitzLB
            // 
            this.blitzLB.Dock = System.Windows.Forms.DockStyle.Fill;
            this.blitzLB.FormattingEnabled = true;
            this.blitzLB.Items.AddRange(new object[] {
            FF6_Save_Editor.Enums.BlitzEnum.Pummel,
            FF6_Save_Editor.Enums.BlitzEnum.Surplex,
            FF6_Save_Editor.Enums.BlitzEnum.AuraBolt,
            FF6_Save_Editor.Enums.BlitzEnum.FireDance,
            FF6_Save_Editor.Enums.BlitzEnum.Mantra,
            FF6_Save_Editor.Enums.BlitzEnum.AirBlade,
            FF6_Save_Editor.Enums.BlitzEnum.Spiraler,
            FF6_Save_Editor.Enums.BlitzEnum.BumRush});
            this.blitzLB.Location = new System.Drawing.Point(3, 16);
            this.blitzLB.Name = "blitzLB";
            this.blitzLB.Size = new System.Drawing.Size(108, 129);
            this.blitzLB.TabIndex = 1;
            this.blitzLB.SelectedIndexChanged += new System.EventHandler(this.listBox_SelectedIndexChanged);
            // 
            // tableLayoutPanel24
            // 
            this.tableLayoutPanel24.ColumnCount = 2;
            this.tableLayoutPanel24.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel24.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel24.Controls.Add(this.groupBox39, 1, 0);
            this.tableLayoutPanel24.Controls.Add(this.groupBox38, 0, 0);
            this.tableLayoutPanel24.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel24.Location = new System.Drawing.Point(3, 163);
            this.tableLayoutPanel24.Name = "tableLayoutPanel24";
            this.tableLayoutPanel24.RowCount = 1;
            this.tableLayoutPanel24.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel24.Size = new System.Drawing.Size(364, 289);
            this.tableLayoutPanel24.TabIndex = 2;
            // 
            // groupBox39
            // 
            this.groupBox39.Controls.Add(this.tableLayoutPanel11);
            this.groupBox39.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox39.Location = new System.Drawing.Point(185, 3);
            this.groupBox39.Name = "groupBox39";
            this.groupBox39.Size = new System.Drawing.Size(176, 283);
            this.groupBox39.TabIndex = 6;
            this.groupBox39.TabStop = false;
            this.groupBox39.Text = "Rage";
            // 
            // tableLayoutPanel11
            // 
            this.tableLayoutPanel11.ColumnCount = 1;
            this.tableLayoutPanel11.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel11.Controls.Add(this.rageLB, 0, 0);
            this.tableLayoutPanel11.Controls.Add(this.tableLayoutPanel16, 0, 1);
            this.tableLayoutPanel11.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel11.Location = new System.Drawing.Point(3, 16);
            this.tableLayoutPanel11.Name = "tableLayoutPanel11";
            this.tableLayoutPanel11.RowCount = 2;
            this.tableLayoutPanel11.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel11.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 30F));
            this.tableLayoutPanel11.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 20F));
            this.tableLayoutPanel11.Size = new System.Drawing.Size(170, 264);
            this.tableLayoutPanel11.TabIndex = 1;
            // 
            // rageLB
            // 
            this.rageLB.Dock = System.Windows.Forms.DockStyle.Fill;
            this.rageLB.FormattingEnabled = true;
            this.rageLB.Location = new System.Drawing.Point(3, 3);
            this.rageLB.Name = "rageLB";
            this.rageLB.Size = new System.Drawing.Size(164, 228);
            this.rageLB.TabIndex = 3;
            this.rageLB.SelectedIndexChanged += new System.EventHandler(this.listBox_SelectedIndexChanged);
            // 
            // tableLayoutPanel16
            // 
            this.tableLayoutPanel16.ColumnCount = 2;
            this.tableLayoutPanel16.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel16.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel16.Controls.Add(this.rageDeselectAllButton, 1, 0);
            this.tableLayoutPanel16.Controls.Add(this.rageSelectAllButton, 0, 0);
            this.tableLayoutPanel16.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel16.Location = new System.Drawing.Point(3, 237);
            this.tableLayoutPanel16.Name = "tableLayoutPanel16";
            this.tableLayoutPanel16.RowCount = 1;
            this.tableLayoutPanel16.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel16.Size = new System.Drawing.Size(164, 24);
            this.tableLayoutPanel16.TabIndex = 4;
            // 
            // rageDeselectAllButton
            // 
            this.rageDeselectAllButton.Anchor = ((System.Windows.Forms.AnchorStyles)((((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.rageDeselectAllButton.Location = new System.Drawing.Point(84, 2);
            this.rageDeselectAllButton.Margin = new System.Windows.Forms.Padding(2);
            this.rageDeselectAllButton.Name = "rageDeselectAllButton";
            this.rageDeselectAllButton.Size = new System.Drawing.Size(78, 20);
            this.rageDeselectAllButton.TabIndex = 8;
            this.rageDeselectAllButton.Text = "Deselect All";
            this.rageDeselectAllButton.UseVisualStyleBackColor = true;
            this.rageDeselectAllButton.Click += new System.EventHandler(this.DeselectAllRage);
            // 
            // rageSelectAllButton
            // 
            this.rageSelectAllButton.Anchor = ((System.Windows.Forms.AnchorStyles)((((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.rageSelectAllButton.Location = new System.Drawing.Point(2, 2);
            this.rageSelectAllButton.Margin = new System.Windows.Forms.Padding(2);
            this.rageSelectAllButton.Name = "rageSelectAllButton";
            this.rageSelectAllButton.Size = new System.Drawing.Size(78, 20);
            this.rageSelectAllButton.TabIndex = 7;
            this.rageSelectAllButton.Text = "Select All";
            this.rageSelectAllButton.UseVisualStyleBackColor = true;
            this.rageSelectAllButton.Click += new System.EventHandler(this.SelectAllRage);
            // 
            // groupBox38
            // 
            this.groupBox38.Controls.Add(this.tableLayoutPanel10);
            this.groupBox38.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox38.Location = new System.Drawing.Point(3, 3);
            this.groupBox38.Name = "groupBox38";
            this.groupBox38.Size = new System.Drawing.Size(176, 283);
            this.groupBox38.TabIndex = 3;
            this.groupBox38.TabStop = false;
            this.groupBox38.Text = "Lore";
            // 
            // tableLayoutPanel10
            // 
            this.tableLayoutPanel10.ColumnCount = 1;
            this.tableLayoutPanel10.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel10.Controls.Add(this.loreLB, 0, 0);
            this.tableLayoutPanel10.Controls.Add(this.tableLayoutPanel15, 0, 1);
            this.tableLayoutPanel10.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel10.Location = new System.Drawing.Point(3, 16);
            this.tableLayoutPanel10.Name = "tableLayoutPanel10";
            this.tableLayoutPanel10.RowCount = 2;
            this.tableLayoutPanel10.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel10.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 30F));
            this.tableLayoutPanel10.Size = new System.Drawing.Size(170, 264);
            this.tableLayoutPanel10.TabIndex = 0;
            // 
            // loreLB
            // 
            this.loreLB.Dock = System.Windows.Forms.DockStyle.Fill;
            this.loreLB.FormattingEnabled = true;
            this.loreLB.Location = new System.Drawing.Point(3, 3);
            this.loreLB.Name = "loreLB";
            this.loreLB.Size = new System.Drawing.Size(164, 228);
            this.loreLB.TabIndex = 1;
            this.loreLB.SelectedIndexChanged += new System.EventHandler(this.listBox_SelectedIndexChanged);
            // 
            // tableLayoutPanel15
            // 
            this.tableLayoutPanel15.ColumnCount = 2;
            this.tableLayoutPanel15.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel15.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel15.Controls.Add(this.loreDeselectAllButton, 1, 0);
            this.tableLayoutPanel15.Controls.Add(this.loreSelectAllButton, 0, 0);
            this.tableLayoutPanel15.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel15.Location = new System.Drawing.Point(3, 237);
            this.tableLayoutPanel15.Name = "tableLayoutPanel15";
            this.tableLayoutPanel15.RowCount = 1;
            this.tableLayoutPanel15.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel15.Size = new System.Drawing.Size(164, 24);
            this.tableLayoutPanel15.TabIndex = 2;
            // 
            // loreDeselectAllButton
            // 
            this.loreDeselectAllButton.Anchor = ((System.Windows.Forms.AnchorStyles)((((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.loreDeselectAllButton.Location = new System.Drawing.Point(84, 2);
            this.loreDeselectAllButton.Margin = new System.Windows.Forms.Padding(2);
            this.loreDeselectAllButton.Name = "loreDeselectAllButton";
            this.loreDeselectAllButton.Size = new System.Drawing.Size(78, 20);
            this.loreDeselectAllButton.TabIndex = 5;
            this.loreDeselectAllButton.Text = "Deselect All";
            this.loreDeselectAllButton.UseVisualStyleBackColor = true;
            this.loreDeselectAllButton.Click += new System.EventHandler(this.DeselectAllLore);
            // 
            // loreSelectAllButton
            // 
            this.loreSelectAllButton.Anchor = ((System.Windows.Forms.AnchorStyles)((((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.loreSelectAllButton.Location = new System.Drawing.Point(2, 2);
            this.loreSelectAllButton.Margin = new System.Windows.Forms.Padding(2);
            this.loreSelectAllButton.Name = "loreSelectAllButton";
            this.loreSelectAllButton.Size = new System.Drawing.Size(78, 20);
            this.loreSelectAllButton.TabIndex = 4;
            this.loreSelectAllButton.Text = "Select All";
            this.loreSelectAllButton.UseVisualStyleBackColor = true;
            this.loreSelectAllButton.Click += new System.EventHandler(this.SelectAllLore);
            // 
            // tabPage10
            // 
            this.tabPage10.Controls.Add(this.tableLayoutPanel20);
            this.tabPage10.Location = new System.Drawing.Point(4, 22);
            this.tabPage10.Name = "tabPage10";
            this.tabPage10.Padding = new System.Windows.Forms.Padding(3);
            this.tabPage10.Size = new System.Drawing.Size(376, 461);
            this.tabPage10.TabIndex = 3;
            this.tabPage10.Text = "Espers";
            this.tabPage10.UseVisualStyleBackColor = true;
            // 
            // tableLayoutPanel20
            // 
            this.tableLayoutPanel20.ColumnCount = 2;
            this.tableLayoutPanel20.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Absolute, 175F));
            this.tableLayoutPanel20.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel20.Controls.Add(this.tableLayoutPanel21, 0, 0);
            this.tableLayoutPanel20.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel20.Location = new System.Drawing.Point(3, 3);
            this.tableLayoutPanel20.Name = "tableLayoutPanel20";
            this.tableLayoutPanel20.RowCount = 1;
            this.tableLayoutPanel20.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel20.Size = new System.Drawing.Size(370, 455);
            this.tableLayoutPanel20.TabIndex = 0;
            // 
            // tableLayoutPanel21
            // 
            this.tableLayoutPanel21.ColumnCount = 1;
            this.tableLayoutPanel21.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel21.Controls.Add(this.espersLB, 0, 0);
            this.tableLayoutPanel21.Controls.Add(this.tableLayoutPanel22, 0, 1);
            this.tableLayoutPanel21.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel21.Location = new System.Drawing.Point(3, 3);
            this.tableLayoutPanel21.Name = "tableLayoutPanel21";
            this.tableLayoutPanel21.RowCount = 2;
            this.tableLayoutPanel21.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel21.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 30F));
            this.tableLayoutPanel21.Size = new System.Drawing.Size(169, 449);
            this.tableLayoutPanel21.TabIndex = 2;
            // 
            // espersLB
            // 
            this.espersLB.Dock = System.Windows.Forms.DockStyle.Fill;
            this.espersLB.FormattingEnabled = true;
            this.espersLB.Location = new System.Drawing.Point(3, 3);
            this.espersLB.Name = "espersLB";
            this.espersLB.Size = new System.Drawing.Size(163, 413);
            this.espersLB.TabIndex = 1;
            this.espersLB.SelectedIndexChanged += new System.EventHandler(this.listBox_SelectedIndexChanged);
            // 
            // tableLayoutPanel22
            // 
            this.tableLayoutPanel22.ColumnCount = 2;
            this.tableLayoutPanel22.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel22.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel22.Controls.Add(this.esperDeselectAll, 0, 0);
            this.tableLayoutPanel22.Controls.Add(this.esperSelectAll, 0, 0);
            this.tableLayoutPanel22.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel22.Location = new System.Drawing.Point(3, 422);
            this.tableLayoutPanel22.Name = "tableLayoutPanel22";
            this.tableLayoutPanel22.RowCount = 1;
            this.tableLayoutPanel22.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 50F));
            this.tableLayoutPanel22.Size = new System.Drawing.Size(163, 24);
            this.tableLayoutPanel22.TabIndex = 2;
            // 
            // esperDeselectAll
            // 
            this.esperDeselectAll.Anchor = ((System.Windows.Forms.AnchorStyles)((((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.esperDeselectAll.Location = new System.Drawing.Point(83, 2);
            this.esperDeselectAll.Margin = new System.Windows.Forms.Padding(2);
            this.esperDeselectAll.Name = "esperDeselectAll";
            this.esperDeselectAll.Size = new System.Drawing.Size(78, 20);
            this.esperDeselectAll.TabIndex = 4;
            this.esperDeselectAll.Text = "Deselect All";
            this.esperDeselectAll.UseVisualStyleBackColor = true;
            this.esperDeselectAll.Click += new System.EventHandler(this.DeselectAllEspers);
            // 
            // esperSelectAll
            // 
            this.esperSelectAll.Anchor = ((System.Windows.Forms.AnchorStyles)((((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left) 
            | System.Windows.Forms.AnchorStyles.Right)));
            this.esperSelectAll.Location = new System.Drawing.Point(2, 2);
            this.esperSelectAll.Margin = new System.Windows.Forms.Padding(2);
            this.esperSelectAll.Name = "esperSelectAll";
            this.esperSelectAll.Size = new System.Drawing.Size(77, 20);
            this.esperSelectAll.TabIndex = 3;
            this.esperSelectAll.Text = "Select All";
            this.esperSelectAll.UseVisualStyleBackColor = true;
            this.esperSelectAll.Click += new System.EventHandler(this.SelectAllEspers);
            // 
            // tabPage4
            // 
            this.tabPage4.Controls.Add(this.tableLayoutPanel1);
            this.tabPage4.Location = new System.Drawing.Point(4, 22);
            this.tabPage4.Name = "tabPage4";
            this.tabPage4.Padding = new System.Windows.Forms.Padding(3);
            this.tabPage4.Size = new System.Drawing.Size(376, 461);
            this.tabPage4.TabIndex = 2;
            this.tabPage4.Text = "Other";
            this.tabPage4.UseVisualStyleBackColor = true;
            // 
            // tableLayoutPanel1
            // 
            this.tableLayoutPanel1.ColumnCount = 3;
            this.tableLayoutPanel1.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Absolute, 150F));
            this.tableLayoutPanel1.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Absolute, 150F));
            this.tableLayoutPanel1.ColumnStyles.Add(new System.Windows.Forms.ColumnStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel1.Controls.Add(this.groupBoxAS, 0, 4);
            this.tableLayoutPanel1.Controls.Add(this.groupBox6, 0, 2);
            this.tableLayoutPanel1.Controls.Add(this.groupBox2, 0, 1);
            this.tableLayoutPanel1.Controls.Add(this.groupBox1, 0, 0);
            this.tableLayoutPanel1.Controls.Add(this.groupBox7, 1, 2);
            this.tableLayoutPanel1.Controls.Add(this.groupBox9, 1, 3);
            this.tableLayoutPanel1.Controls.Add(this.groupBox8, 0, 3);
            this.tableLayoutPanel1.Controls.Add(this.groupBoxASX, 1, 4);
            this.tableLayoutPanel1.Controls.Add(this.isAirshipVisible, 0, 5);
            this.tableLayoutPanel1.Dock = System.Windows.Forms.DockStyle.Fill;
            this.tableLayoutPanel1.Location = new System.Drawing.Point(3, 3);
            this.tableLayoutPanel1.Name = "tableLayoutPanel1";
            this.tableLayoutPanel1.RowCount = 10;
            this.tableLayoutPanel1.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel1.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel1.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel1.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel1.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel1.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel1.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel1.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel1.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Absolute, 50F));
            this.tableLayoutPanel1.RowStyles.Add(new System.Windows.Forms.RowStyle(System.Windows.Forms.SizeType.Percent, 100F));
            this.tableLayoutPanel1.Size = new System.Drawing.Size(370, 455);
            this.tableLayoutPanel1.TabIndex = 0;
            // 
            // groupBoxAS
            // 
            this.groupBoxAS.Controls.Add(this.airshipXAxis);
            this.groupBoxAS.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBoxAS.Location = new System.Drawing.Point(3, 203);
            this.groupBoxAS.Name = "groupBoxAS";
            this.groupBoxAS.Size = new System.Drawing.Size(144, 44);
            this.groupBoxAS.TabIndex = 6;
            this.groupBoxAS.TabStop = false;
            this.groupBoxAS.Text = "Airship X Axis";
            // 
            // airshipXAxis
            // 
            this.airshipXAxis.Dock = System.Windows.Forms.DockStyle.Fill;
            this.airshipXAxis.Location = new System.Drawing.Point(3, 16);
            this.airshipXAxis.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.airshipXAxis.Name = "airshipXAxis";
            this.airshipXAxis.Size = new System.Drawing.Size(138, 20);
            this.airshipXAxis.TabIndex = 5;
            // 
            // groupBox6
            // 
            this.groupBox6.Controls.Add(this.numberOfSaves);
            this.groupBox6.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox6.Location = new System.Drawing.Point(3, 103);
            this.groupBox6.Name = "groupBox6";
            this.groupBox6.Size = new System.Drawing.Size(144, 44);
            this.groupBox6.TabIndex = 3;
            this.groupBox6.TabStop = false;
            this.groupBox6.Text = "Number of Saves";
            // 
            // numberOfSaves
            // 
            this.numberOfSaves.Dock = System.Windows.Forms.DockStyle.Fill;
            this.numberOfSaves.Location = new System.Drawing.Point(3, 16);
            this.numberOfSaves.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.numberOfSaves.Name = "numberOfSaves";
            this.numberOfSaves.Size = new System.Drawing.Size(138, 20);
            this.numberOfSaves.TabIndex = 3;
            // 
            // groupBox2
            // 
            this.groupBox2.Controls.Add(this.steps);
            this.groupBox2.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox2.Location = new System.Drawing.Point(3, 53);
            this.groupBox2.Name = "groupBox2";
            this.groupBox2.Size = new System.Drawing.Size(144, 44);
            this.groupBox2.TabIndex = 2;
            this.groupBox2.TabStop = false;
            this.groupBox2.Text = "Steps";
            // 
            // steps
            // 
            this.steps.Dock = System.Windows.Forms.DockStyle.Fill;
            this.steps.Location = new System.Drawing.Point(3, 16);
            this.steps.Maximum = new decimal(new int[] {
            16777216,
            0,
            0,
            0});
            this.steps.Name = "steps";
            this.steps.Size = new System.Drawing.Size(138, 20);
            this.steps.TabIndex = 2;
            // 
            // groupBox1
            // 
            this.groupBox1.Controls.Add(this.gold);
            this.groupBox1.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox1.Location = new System.Drawing.Point(3, 3);
            this.groupBox1.Name = "groupBox1";
            this.groupBox1.Size = new System.Drawing.Size(144, 44);
            this.groupBox1.TabIndex = 1;
            this.groupBox1.TabStop = false;
            this.groupBox1.Text = "GP";
            // 
            // gold
            // 
            this.gold.Dock = System.Windows.Forms.DockStyle.Fill;
            this.gold.Location = new System.Drawing.Point(3, 16);
            this.gold.Maximum = new decimal(new int[] {
            16777216,
            0,
            0,
            0});
            this.gold.Name = "gold";
            this.gold.Size = new System.Drawing.Size(138, 20);
            this.gold.TabIndex = 1;
            // 
            // groupBox7
            // 
            this.groupBox7.Controls.Add(this.saveCountRollover);
            this.groupBox7.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox7.Location = new System.Drawing.Point(153, 103);
            this.groupBox7.Name = "groupBox7";
            this.groupBox7.Size = new System.Drawing.Size(144, 44);
            this.groupBox7.TabIndex = 4;
            this.groupBox7.TabStop = false;
            this.groupBox7.Text = "Save Count Rollover";
            // 
            // saveCountRollover
            // 
            this.saveCountRollover.Dock = System.Windows.Forms.DockStyle.Fill;
            this.saveCountRollover.Location = new System.Drawing.Point(3, 16);
            this.saveCountRollover.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.saveCountRollover.Name = "saveCountRollover";
            this.saveCountRollover.Size = new System.Drawing.Size(138, 20);
            this.saveCountRollover.TabIndex = 4;
            // 
            // groupBox9
            // 
            this.groupBox9.Controls.Add(this.mapYAxis);
            this.groupBox9.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox9.Location = new System.Drawing.Point(153, 153);
            this.groupBox9.Name = "groupBox9";
            this.groupBox9.Size = new System.Drawing.Size(144, 44);
            this.groupBox9.TabIndex = 6;
            this.groupBox9.TabStop = false;
            this.groupBox9.Text = "Map Y Axis";
            // 
            // mapYAxis
            // 
            this.mapYAxis.Dock = System.Windows.Forms.DockStyle.Fill;
            this.mapYAxis.Location = new System.Drawing.Point(3, 16);
            this.mapYAxis.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.mapYAxis.Name = "mapYAxis";
            this.mapYAxis.Size = new System.Drawing.Size(138, 20);
            this.mapYAxis.TabIndex = 6;
            // 
            // groupBox8
            // 
            this.groupBox8.Controls.Add(this.mapXAxis);
            this.groupBox8.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBox8.Location = new System.Drawing.Point(3, 153);
            this.groupBox8.Name = "groupBox8";
            this.groupBox8.Size = new System.Drawing.Size(144, 44);
            this.groupBox8.TabIndex = 5;
            this.groupBox8.TabStop = false;
            this.groupBox8.Text = "Map X Axis";
            // 
            // mapXAxis
            // 
            this.mapXAxis.Dock = System.Windows.Forms.DockStyle.Fill;
            this.mapXAxis.Location = new System.Drawing.Point(3, 16);
            this.mapXAxis.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.mapXAxis.Name = "mapXAxis";
            this.mapXAxis.Size = new System.Drawing.Size(138, 20);
            this.mapXAxis.TabIndex = 5;
            // 
            // groupBoxASX
            // 
            this.groupBoxASX.Controls.Add(this.airshipYAxis);
            this.groupBoxASX.Dock = System.Windows.Forms.DockStyle.Fill;
            this.groupBoxASX.Location = new System.Drawing.Point(153, 203);
            this.groupBoxASX.Name = "groupBoxASX";
            this.groupBoxASX.Size = new System.Drawing.Size(144, 44);
            this.groupBoxASX.TabIndex = 7;
            this.groupBoxASX.TabStop = false;
            this.groupBoxASX.Text = "Airship Y Axis";
            // 
            // airshipYAxis
            // 
            this.airshipYAxis.Dock = System.Windows.Forms.DockStyle.Fill;
            this.airshipYAxis.Location = new System.Drawing.Point(3, 16);
            this.airshipYAxis.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.airshipYAxis.Name = "airshipYAxis";
            this.airshipYAxis.Size = new System.Drawing.Size(138, 20);
            this.airshipYAxis.TabIndex = 5;
            // 
            // isAirshipVisible
            // 
            this.isAirshipVisible.Anchor = ((System.Windows.Forms.AnchorStyles)(((System.Windows.Forms.AnchorStyles.Top | System.Windows.Forms.AnchorStyles.Bottom) 
            | System.Windows.Forms.AnchorStyles.Left)));
            this.isAirshipVisible.AutoSize = true;
            this.isAirshipVisible.Location = new System.Drawing.Point(3, 253);
            this.isAirshipVisible.Name = "isAirshipVisible";
            this.isAirshipVisible.Size = new System.Drawing.Size(101, 44);
            this.isAirshipVisible.TabIndex = 0;
            this.isAirshipVisible.Text = "Is Airship Visible";
            this.isAirshipVisible.UseVisualStyleBackColor = true;
            // 
            // armorId
            // 
            this.armorId.Dock = System.Windows.Forms.DockStyle.Fill;
            this.armorId.Hexadecimal = true;
            this.armorId.HexLength = 2;
            this.armorId.Location = new System.Drawing.Point(3, 16);
            this.armorId.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.armorId.Name = "armorId";
            this.armorId.Size = new System.Drawing.Size(138, 20);
            this.armorId.TabIndex = 1;
            this.armorId.Value = ((long)(0));
            // 
            // helmitId
            // 
            this.helmitId.Dock = System.Windows.Forms.DockStyle.Fill;
            this.helmitId.Hexadecimal = true;
            this.helmitId.HexLength = 2;
            this.helmitId.Location = new System.Drawing.Point(3, 16);
            this.helmitId.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.helmitId.Name = "helmitId";
            this.helmitId.Size = new System.Drawing.Size(138, 20);
            this.helmitId.TabIndex = 1;
            this.helmitId.Value = ((long)(0));
            // 
            // shieldId
            // 
            this.shieldId.Dock = System.Windows.Forms.DockStyle.Fill;
            this.shieldId.Hexadecimal = true;
            this.shieldId.HexLength = 2;
            this.shieldId.Location = new System.Drawing.Point(3, 16);
            this.shieldId.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.shieldId.Name = "shieldId";
            this.shieldId.Size = new System.Drawing.Size(138, 20);
            this.shieldId.TabIndex = 1;
            this.shieldId.Value = ((long)(0));
            // 
            // weaponId
            // 
            this.weaponId.Dock = System.Windows.Forms.DockStyle.Fill;
            this.weaponId.Hexadecimal = true;
            this.weaponId.HexLength = 2;
            this.weaponId.Location = new System.Drawing.Point(3, 16);
            this.weaponId.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.weaponId.Name = "weaponId";
            this.weaponId.Size = new System.Drawing.Size(138, 20);
            this.weaponId.TabIndex = 0;
            this.weaponId.Value = ((long)(0));
            // 
            // relic2Id
            // 
            this.relic2Id.Dock = System.Windows.Forms.DockStyle.Fill;
            this.relic2Id.Hexadecimal = true;
            this.relic2Id.HexLength = 2;
            this.relic2Id.Location = new System.Drawing.Point(3, 16);
            this.relic2Id.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.relic2Id.Name = "relic2Id";
            this.relic2Id.Size = new System.Drawing.Size(138, 20);
            this.relic2Id.TabIndex = 1;
            this.relic2Id.Value = ((long)(0));
            // 
            // relic1Id
            // 
            this.relic1Id.Dock = System.Windows.Forms.DockStyle.Fill;
            this.relic1Id.Hexadecimal = true;
            this.relic1Id.HexLength = 2;
            this.relic1Id.Location = new System.Drawing.Point(3, 16);
            this.relic1Id.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.relic1Id.Name = "relic1Id";
            this.relic1Id.Size = new System.Drawing.Size(138, 20);
            this.relic1Id.TabIndex = 1;
            this.relic1Id.Value = ((long)(0));
            // 
            // findByHex
            // 
            this.findByHex.Dock = System.Windows.Forms.DockStyle.Fill;
            this.findByHex.Hexadecimal = true;
            this.findByHex.HexLength = 2;
            this.findByHex.Location = new System.Drawing.Point(3, 3);
            this.findByHex.Maximum = new decimal(new int[] {
            255,
            0,
            0,
            0});
            this.findByHex.Name = "findByHex";
            this.findByHex.Size = new System.Drawing.Size(60, 20);
            this.findByHex.TabIndex = 0;
            this.findByHex.Value = ((long)(0));
            this.findByHex.ValueChanged += new System.EventHandler(this.findByHex_ValueChanged);
            // 
            // SaveEditorForm
            // 
            this.AutoScaleDimensions = new System.Drawing.SizeF(6F, 13F);
            this.AutoScaleMode = System.Windows.Forms.AutoScaleMode.Font;
            this.ClientSize = new System.Drawing.Size(384, 511);
            this.Controls.Add(this.mainTabs);
            this.Controls.Add(this.menuStrip1);
            this.MainMenuStrip = this.menuStrip1;
            this.Name = "SaveEditorForm";
            this.Text = "Final Fantasy 3/6 Save State Editor";
            this.menuStrip1.ResumeLayout(false);
            this.menuStrip1.PerformLayout();
            this.mainTabs.ResumeLayout(false);
            this.tabPage1.ResumeLayout(false);
            this.tableLayoutPanel2.ResumeLayout(false);
            this.characterTabs.ResumeLayout(false);
            this.tabPage5.ResumeLayout(false);
            this.tableLayoutPanel25.ResumeLayout(false);
            this.tableLayoutPanel3.ResumeLayout(false);
            this.groupBox34.ResumeLayout(false);
            this.groupBox34.PerformLayout();
            this.groupBox13.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.vigor)).EndInit();
            this.groupBox14.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.stamina)).EndInit();
            this.groupBox15.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.speed)).EndInit();
            this.groupBox16.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.magic)).EndInit();
            this.groupBox12.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.currentMP)).EndInit();
            this.groupBox33.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.maxMP)).EndInit();
            this.groupBox32.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.maxHP)).EndInit();
            this.groupBox11.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.currentHp)).EndInit();
            this.groupBox17.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.exp)).EndInit();
            this.groupBox36.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.level)).EndInit();
            this.tabPage6.ResumeLayout(false);
            this.tableLayoutPanel7.ResumeLayout(false);
            this.tableLayoutPanel7.PerformLayout();
            this.groupBox18.ResumeLayout(false);
            this.tabPage7.ResumeLayout(false);
            this.tableLayoutPanel13.ResumeLayout(false);
            this.tableLayoutPanel8.ResumeLayout(false);
            this.groupBox25.ResumeLayout(false);
            this.groupBox23.ResumeLayout(false);
            this.groupBox21.ResumeLayout(false);
            this.groupBox19.ResumeLayout(false);
            this.groupBox29.ResumeLayout(false);
            this.groupBox27.ResumeLayout(false);
            this.tableLayoutPanel14.ResumeLayout(false);
            this.groupBox28.ResumeLayout(false);
            this.groupBox30.ResumeLayout(false);
            this.groupBox31.ResumeLayout(false);
            this.tabPage8.ResumeLayout(false);
            this.tableLayoutPanel9.ResumeLayout(false);
            this.groupBox20.ResumeLayout(false);
            this.groupBox22.ResumeLayout(false);
            this.groupBox24.ResumeLayout(false);
            this.groupBox26.ResumeLayout(false);
            this.tabPage9.ResumeLayout(false);
            this.tableLayoutPanel12.ResumeLayout(false);
            this.tableLayoutPanel4.ResumeLayout(false);
            this.groupBox10.ResumeLayout(false);
            this.tabPage2.ResumeLayout(false);
            this.tableLayoutPanel5.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.inventoryGrid)).EndInit();
            this.tableLayoutPanel17.ResumeLayout(false);
            this.groupBox4.ResumeLayout(false);
            this.tableLayoutPanel19.ResumeLayout(false);
            this.tableLayoutPanel19.PerformLayout();
            this.groupBox3.ResumeLayout(false);
            this.tableLayoutPanel18.ResumeLayout(false);
            this.tableLayoutPanel18.PerformLayout();
            this.tabPage3.ResumeLayout(false);
            this.tableLayoutPanel6.ResumeLayout(false);
            this.tableLayoutPanel23.ResumeLayout(false);
            this.groupBox5.ResumeLayout(false);
            this.groupBox35.ResumeLayout(false);
            this.groupBox37.ResumeLayout(false);
            this.tableLayoutPanel24.ResumeLayout(false);
            this.groupBox39.ResumeLayout(false);
            this.tableLayoutPanel11.ResumeLayout(false);
            this.tableLayoutPanel16.ResumeLayout(false);
            this.groupBox38.ResumeLayout(false);
            this.tableLayoutPanel10.ResumeLayout(false);
            this.tableLayoutPanel15.ResumeLayout(false);
            this.tabPage10.ResumeLayout(false);
            this.tableLayoutPanel20.ResumeLayout(false);
            this.tableLayoutPanel21.ResumeLayout(false);
            this.tableLayoutPanel22.ResumeLayout(false);
            this.tabPage4.ResumeLayout(false);
            this.tableLayoutPanel1.ResumeLayout(false);
            this.tableLayoutPanel1.PerformLayout();
            this.groupBoxAS.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.airshipXAxis)).EndInit();
            this.groupBox6.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.numberOfSaves)).EndInit();
            this.groupBox2.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.steps)).EndInit();
            this.groupBox1.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.gold)).EndInit();
            this.groupBox7.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.saveCountRollover)).EndInit();
            this.groupBox9.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.mapYAxis)).EndInit();
            this.groupBox8.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.mapXAxis)).EndInit();
            this.groupBoxASX.ResumeLayout(false);
            ((System.ComponentModel.ISupportInitialize)(this.airshipYAxis)).EndInit();
            ((System.ComponentModel.ISupportInitialize)(this.armorId)).EndInit();
            ((System.ComponentModel.ISupportInitialize)(this.helmitId)).EndInit();
            ((System.ComponentModel.ISupportInitialize)(this.shieldId)).EndInit();
            ((System.ComponentModel.ISupportInitialize)(this.weaponId)).EndInit();
            ((System.ComponentModel.ISupportInitialize)(this.relic2Id)).EndInit();
            ((System.ComponentModel.ISupportInitialize)(this.relic1Id)).EndInit();
            ((System.ComponentModel.ISupportInitialize)(this.findByHex)).EndInit();
            this.ResumeLayout(false);
            this.PerformLayout();

        }

        #endregion

        private System.Windows.Forms.MenuStrip menuStrip1;
        private System.Windows.Forms.ToolStripMenuItem fileToolStripMenuItem;
        private System.Windows.Forms.TabControl mainTabs;
        private System.Windows.Forms.TabPage tabPage1;
        private System.Windows.Forms.TabPage tabPage2;
        private System.Windows.Forms.TabPage tabPage3;
        private System.Windows.Forms.TabPage tabPage4;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel1;
        private System.Windows.Forms.GroupBox groupBox2;
        private System.Windows.Forms.GroupBox groupBox1;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel2;
        private System.Windows.Forms.TabControl characterTabs;
        private System.Windows.Forms.TabPage tabPage5;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel3;
        private System.Windows.Forms.TabPage tabPage6;
        private System.Windows.Forms.TabPage tabPage7;
        private System.Windows.Forms.TabPage tabPage8;
        private System.Windows.Forms.TabPage tabPage9;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel4;
        private System.Windows.Forms.GroupBox groupBox10;
        private System.Windows.Forms.ComboBox characterComboBox;
        private System.Windows.Forms.GroupBox groupBox11;
        private System.Windows.Forms.GroupBox groupBox17;
        private System.Windows.Forms.NumericUpDown exp;
        private System.Windows.Forms.GroupBox groupBox16;
        private System.Windows.Forms.NumericUpDown magic;
        private System.Windows.Forms.GroupBox groupBox15;
        private System.Windows.Forms.NumericUpDown speed;
        private System.Windows.Forms.GroupBox groupBox14;
        private System.Windows.Forms.NumericUpDown stamina;
        private System.Windows.Forms.GroupBox groupBox13;
        private System.Windows.Forms.NumericUpDown vigor;
        private System.Windows.Forms.GroupBox groupBox18;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel7;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel8;
        private System.Windows.Forms.GroupBox groupBox25;
        private System.Windows.Forms.GroupBox groupBox23;
        private System.Windows.Forms.GroupBox groupBox21;
        private System.Windows.Forms.GroupBox groupBox19;
        private System.Windows.Forms.GroupBox groupBox29;
        private System.Windows.Forms.GroupBox groupBox27;
        private UI.HexNumericUpDown armorId;
        private UI.HexNumericUpDown helmitId;
        private UI.HexNumericUpDown shieldId;
        private UI.HexNumericUpDown weaponId;
        private UI.HexNumericUpDown relic2Id;
        private UI.HexNumericUpDown relic1Id;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel9;
        private System.Windows.Forms.ComboBox command4;
        private System.Windows.Forms.ComboBox command3;
        private System.Windows.Forms.ComboBox command2;
        private System.Windows.Forms.ComboBox command1;
        private System.Windows.Forms.TableLayoutPanel magicGrid;
        private System.Windows.Forms.GroupBox groupBox20;
        private System.Windows.Forms.GroupBox groupBox22;
        private System.Windows.Forms.GroupBox groupBox24;
        private System.Windows.Forms.GroupBox groupBox26;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel12;
        private System.Windows.Forms.CheckedListBox statusEffectsLB;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel13;
        private System.Windows.Forms.RichTextBox weaponShieldRTB;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel14;
        private System.Windows.Forms.GroupBox groupBox28;
        private System.Windows.Forms.GroupBox groupBox30;
        private System.Windows.Forms.RichTextBox helmetArmorRTB;
        private System.Windows.Forms.GroupBox groupBox31;
        private System.Windows.Forms.RichTextBox relicsRTB;
        private System.Windows.Forms.TextBox magicHelpText;
        private System.Windows.Forms.GroupBox groupBox34;
        private System.Windows.Forms.TextBox nameTB;
        private System.Windows.Forms.GroupBox groupBox12;
        private System.Windows.Forms.NumericUpDown currentMP;
        private System.Windows.Forms.GroupBox groupBox33;
        private System.Windows.Forms.NumericUpDown maxMP;
        private System.Windows.Forms.GroupBox groupBox32;
        private System.Windows.Forms.NumericUpDown maxHP;
        private System.Windows.Forms.NumericUpDown currentHp;
        private System.Windows.Forms.GroupBox groupBox36;
        private System.Windows.Forms.NumericUpDown level;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel5;
        private System.Windows.Forms.RichTextBox inventoryRTB;
        private System.Windows.Forms.DataGridView inventoryGrid;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel6;
        private System.Windows.Forms.GroupBox groupBox39;
        private System.Windows.Forms.GroupBox groupBox38;
        private System.Windows.Forms.GroupBox groupBox37;
        private System.Windows.Forms.CheckedListBox blitzLB;
        private System.Windows.Forms.GroupBox groupBox35;
        private System.Windows.Forms.CheckedListBox swdtechLB;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel11;
        private System.Windows.Forms.CheckedListBox rageLB;
        private System.Windows.Forms.Button rageSelectAllButton;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel10;
        private System.Windows.Forms.CheckedListBox loreLB;
        private System.Windows.Forms.Button loreSelectAllButton;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel16;
        private System.Windows.Forms.Button rageDeselectAllButton;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel15;
        private System.Windows.Forms.Button loreDeselectAllButton;
        private System.Windows.Forms.NumericUpDown steps;
        private System.Windows.Forms.NumericUpDown gold;
        private System.Windows.Forms.ToolStripMenuItem saveToolStripMenuItem;
        private System.Windows.Forms.ToolStripMenuItem saveAsToolStripMenuItem;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel17;
        private System.Windows.Forms.GroupBox groupBox4;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel19;
        private System.Windows.Forms.Label resultByName;
        private System.Windows.Forms.TextBox findByName;
        private System.Windows.Forms.GroupBox groupBox3;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel18;
        private UI.HexNumericUpDown findByHex;
        private System.Windows.Forms.Label resultByHex;
        private System.Windows.Forms.TabPage tabPage10;
        private System.Windows.Forms.ToolStripMenuItem openToolStripMenuItem;
        private System.Windows.Forms.ToolStripMenuItem sRMToolStripMenuItem;
        private System.Windows.Forms.ToolStripMenuItem snes9xToolStripMenuItem;
        private System.Windows.Forms.ToolStripMenuItem toolStripMenuItem1;
        private System.Windows.Forms.ToolStripSeparator toolStripSeparator1;
        private System.Windows.Forms.ToolStripSeparator toolStripSeparator2;
        private System.Windows.Forms.ToolStripMenuItem exitToolStripMenuItem;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel20;
        private System.Windows.Forms.CheckedListBox espersLB;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel21;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel22;
        private System.Windows.Forms.Button esperDeselectAll;
        private System.Windows.Forms.Button esperSelectAll;
        private System.Windows.Forms.ToolStripMenuItem slot1ToolStripMenuItem;
        private System.Windows.Forms.ToolStripMenuItem slot2ToolStripMenuItem;
        private System.Windows.Forms.ToolStripMenuItem slot3ToolStripMenuItem;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel23;
        private System.Windows.Forms.GroupBox groupBox5;
        private System.Windows.Forms.CheckedListBox danceLB;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel24;
        private System.Windows.Forms.GroupBox groupBox6;
        private System.Windows.Forms.NumericUpDown numberOfSaves;
        private System.Windows.Forms.TableLayoutPanel tableLayoutPanel25;
        private System.Windows.Forms.RichTextBox experienceTB;
        private System.Windows.Forms.GroupBox groupBox7;
        private System.Windows.Forms.NumericUpDown saveCountRollover;
        private System.Windows.Forms.ToolTip toolTip1;
        private System.Windows.Forms.GroupBox groupBox8;
        private System.Windows.Forms.NumericUpDown mapXAxis;
        private System.Windows.Forms.GroupBox groupBox9;
        private System.Windows.Forms.NumericUpDown mapYAxis;
        private System.Windows.Forms.GroupBox groupBoxAS;
        private System.Windows.Forms.NumericUpDown airshipXAxis;
        private System.Windows.Forms.GroupBox groupBoxASX;
        private System.Windows.Forms.NumericUpDown airshipYAxis;
        private System.Windows.Forms.CheckBox isAirshipVisible;
        private System.Windows.Forms.ToolStripMenuItem snes9xV16ToolStripMenuItem;
    }
}

