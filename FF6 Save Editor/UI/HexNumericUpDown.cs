using System;

namespace FF6_Save_Editor.UI
{
    /// <summary>
    /// A modified NumbericUpDown UI component which limits the number of displayed values to a defined limit (HexLength) defaulted to 2.
    /// </summary>
    public class HexNumericUpDown : System.Windows.Forms.NumericUpDown
    {
        public HexNumericUpDown()
        {
            this.Hexadecimal = true;
            this.HexLength = 2;
        }

        protected override void ValidateEditText()
        {
            if (base.UserEdit)
            {
                base.ValidateEditText();
            }
        }

        protected override void UpdateEditText()
        {
            Text = System.Convert.ToInt64(base.Value).ToString("X" + HexLength);
        }

        [System.ComponentModel.DefaultValue(4)]
        public int HexLength { get; set; }

        public new Int64 Value
        {
            get { return System.Convert.ToInt64(base.Value); }
            set { base.Value = System.Convert.ToDecimal(value); }
        }
    }
}
