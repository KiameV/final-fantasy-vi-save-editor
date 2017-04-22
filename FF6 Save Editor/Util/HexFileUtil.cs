using FF6_Save_State_Editor.Enums;
using System;
using System.IO;
using System.IO.Compression;
using System.Windows.Forms;

namespace FF6_Save_Editor.Util
{
    /// <summary>
    /// Utility for opening and saving hex-based files
    /// </summary>
    class HexFileUtil
    {
        /// <summary>
        /// The selected file. This is set using FindHexFile.
        /// </summary>
        public static string SelectedFile { get; private set; }

        /// <summary>
        /// Will prompt the user to find a file
        /// </summary>
        /// <param name="saveFileType">Specifies what type of file to find. ie ZNES, Snes9x, or SRM</param>
        /// <param name="fileAccess">Specifies the action that will be taken. ie Read or Write</param>
        /// <returns></returns>
        public static bool FindHexFile(SaveFileTypeEnum saveFileType, FileAccess fileAccess)
        {
            if (fileAccess == FileAccess.Read)
                return OpenDialog(saveFileType);
            return SaveDialog(saveFileType);
        }

        /// <summary>
        /// Open and parse the selected file. Must use FindHexFile first.
        /// </summary>
        /// <param name="saveFileType">Specifies what type of file to find. ie ZNES, Snes9x, or SRM</param>
        /// <param name="byteStream">The byte array to write the opened file to</param>
        /// <returns>True if the parsing was successful. False if there was a problem.</returns>
        public static bool OpenFile(SaveFileTypeEnum saveFileType, out Byte[] byteStream)
        {
            if (saveFileType == SaveFileTypeEnum.Snes9xSaveState)
            {
                return OpenGzipHexFile(out byteStream);
            }
            return OpenHexFile(out byteStream);
        }

        /// <summary>
        /// Save the given byte array to the file
        /// </summary>
        /// <param name="saveFileType">Specifies what type of file to find. ie ZNES, Snes9x, or SRM</param>
        /// <param name="byteStream">The byte array to write to the file</param>
        /// <returns>True if the save was successful. False if there was a problem.</returns>
        public static bool SaveFile(SaveFileTypeEnum saveFileType, Byte[] byteStream)
        {
            if (saveFileType == SaveFileTypeEnum.Snes9xSaveState)
            {
                return SaveGzipHexFile(byteStream);
            }
            return SaveHexFile(byteStream);
        }

        /// <summary>
        /// Create an open dialog window
        /// </summary>
        /// <param name="saveFileType">Specifies what type of file to find. ie ZNES, Snes9x, or SRM</param>
        /// <returns>True if a file was selected. False if the user canceled the operation (did not selected a file).</returns>
        private static bool OpenDialog(SaveFileTypeEnum saveFileType)
        {
            OpenFileDialog ofd = new OpenFileDialog();
            SetFileDialogFields(ofd, saveFileType);

            if (ofd.ShowDialog() == DialogResult.Cancel)
            {
                return false;
            }
            SelectedFile = Path.GetFullPath(ofd.FileName);
            return true;
        }

        /// <summary>
        /// Create a Save Dialog
        /// </summary>
        /// <param name="saveFileType">Specifies what type of file to find. ie ZNES, Snes9x, or SRM</param>
        /// <returns>True if a file was selected. False if the user canceled the operation (did not selected a file).</returns>
        private static bool SaveDialog(SaveFileTypeEnum saveFileType)
        {
            SaveFileDialog sfd = new SaveFileDialog();
            SetFileDialogFields(sfd, saveFileType);

            sfd.RestoreDirectory = true;
            sfd.InitialDirectory = Path.GetDirectoryName(SelectedFile);
            sfd.FileName = SelectedFile;

            if (sfd.ShowDialog() == DialogResult.Cancel)
            {
                return false;
            }

            SelectedFile = Path.GetFullPath(sfd.FileName);
            return true;
        }

        /// <summary>
        /// Helper method for populating the Open/Save Dialog window's title, default file types, and filters.
        /// </summary>
        /// <param name="dialog">The Dialog window which to populate</param>
        /// <param name="saveFileType">Specifies what type of file to find. ie ZNES, Snes9x, or SRM</param>
        private static void SetFileDialogFields(FileDialog dialog, SaveFileTypeEnum saveFileType)
        {
            switch(saveFileType)
            {
                case SaveFileTypeEnum.Snes9xSaveState:
                    dialog.Title = "Select the Save State";
                    dialog.DefaultExt = "*.000";
                    dialog.Filter = "Snes9x Save State Files|*.0*";
                    break;
                case SaveFileTypeEnum.SRMSlot1:
                case SaveFileTypeEnum.SRMSlot2:
                case SaveFileTypeEnum.SRMSlot3:
                    dialog.Title = "Select the SRM";
                    dialog.DefaultExt = "*.srm";
                    dialog.Filter = "SRM Files|*.srm";
                    break;
                case SaveFileTypeEnum.ZnesSaveState:
                    dialog.Title = "Select the Save State";
                    dialog.DefaultExt = "*.zst";
                    dialog.Filter = "ZST Files|*.zs*";
                    break;
            }
        }

        /// <summary>
        /// Helper method for opening and reading a hex file
        /// </summary>
        /// <param name="byteStream">Byte array to write the file to</param>
        /// <returns>True if the operation was a success. False if there was a problem.</returns>
        private static bool OpenHexFile(out Byte[] byteStream)
        {
            if (SelectedFile == null)
            {
                throw new Exception("No file was selected.");
            }

            try
            {
                // get selected filename and open it with binarystream
                using (FileStream fs = File.Open(SelectedFile, FileMode.Open, FileAccess.Read, FileShare.Read))
                {
                    using (BinaryReader binReader = new BinaryReader(fs))
                    {
                        // get file length and alloc array
                        long lfileLength = binReader.BaseStream.Length;
                        byteStream = new Byte[lfileLength];

                        // read every byte of file's content
                        for (long lIdx = 0; lIdx < lfileLength; lIdx++)
                        {
                            byteStream[lIdx] = binReader.ReadByte();
                        }
                    }
                }
                return true;
            }
            catch (Exception e)
            {
                MessageBox.Show("Unable to load file: " + e.Message);
                SelectedFile = null;
                byteStream = null;
                return false;
            }
        }

        /// <summary>
        /// Helper method for opening and reading a gzip'ed hex file
        /// </summary>
        /// <param name="byteStream">Byte array to write the file to</param>
        /// <returns>True if the operation was a success. False if there was a problem.</returns>
        private static bool OpenGzipHexFile(out Byte[] byteStream)
        {
            if (SelectedFile == null)
            {
                throw new Exception("No file was selected.");
            }

            try
            {
                using (FileStream fs = File.Open(SelectedFile, FileMode.Open, FileAccess.Read, FileShare.Read))
                {
                    using (var decompressedMs = new MemoryStream())
                    {
                        using (var gzs = new BufferedStream(new GZipStream(fs, CompressionMode.Decompress), 4096))
                        {
                            gzs.CopyTo(decompressedMs);
                        }
                        byteStream = decompressedMs.ToArray();
                    }
                }
                return true;
            }
            catch (Exception e)
            {
                MessageBox.Show("Unable to load file: " + e.Message);
                SelectedFile = null;
                byteStream = null;
                return false;
            }
        }

        /// <summary>
        /// Helper method for saving a byte array to a hex file
        /// </summary>
        /// <param name="byteStream">Byte array to write to the file</param>
        /// <returns>True if the operation was a success. False if there was a problem.</returns>
        private static bool SaveHexFile(Byte[] byteStream)
        {
            if (SelectedFile == null)
            {
                throw new Exception("No file was selected.");
            }

            try
            {
                using (FileStream fs = File.Open(SelectedFile, FileMode.OpenOrCreate, FileAccess.Write))
                {
                    using (BinaryWriter writer = new BinaryWriter(fs))
                    {
                        writer.Write(byteStream);
                    }
                }
                return true;
            }
            catch
            {
                return false;
            }
        }

        /// <summary>
        /// Helper method for saving a byte array to a gzip'ed hex file
        /// </summary>
        /// <param name="byteStream">Byte array to write to the file</param>
        /// <returns>True if the operation was a success. False if there was a problem.</returns>
        private static bool SaveGzipHexFile(Byte[] byteStream)
        {
            if (SelectedFile == null)
            {
                throw new Exception("No file was selected.");
            }

            try
            {
                using (FileStream fs = File.Open(SelectedFile, FileMode.OpenOrCreate, FileAccess.Write))
                {
                    using (var gzip = new GZipStream(fs, CompressionMode.Compress))
                    {//http://stackoverflow.com/questions/27997161/decompress-by-gzip-failed
                        gzip.Write(byteStream, 0, byteStream.Length);
                    }
                }
                return true;
            }
            catch
            {
                return false;
            }
        }
    }
}
