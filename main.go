package main

import (
	"ffvi_editor/ui/characters"
	"os"

	"github.com/therecipe/qt/widgets"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)
	
	var window = widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(600, 500)
	window.SetWindowTitle("Final Fantasy VI Editor")

	mb := widgets.NewQMenuBar(nil)
	window.SetMenuBar(mb)
	menu := mb.AddMenu2("File")
	menu = menu

	tb := widgets.NewQTabWidget(nil)
	cTab := tb.AddTab(characters.CreateCharactersLayout(window), "Characters")
	cTab = cTab

	window.SetCentralWidget(tb)

	window.Show()

	widgets.QApplication_Exec()
}
