package frontend

import (
	"fmt"
	"fyne.io/fyne/v2"
)

// CreateMenu defines the main menu.
func CreateMenu(gm *GuiMgr) *fyne.MainMenu {

	menuGeneral := createMenuGeneral()
	menuDateTime := createMenuDateTime()
	menuHelp := createMenuHelp()
	mainMenu := fyne.NewMainMenu(menuGeneral, menuDateTime, menuHelp)
	return mainMenu
}

func createMenuGeneral() *fyne.Menu {

	closeMenuItem := fyne.NewMenuItem("Close", func() {
		fmt.Println("Close clicked.")
	})
	return fyne.NewMenu("General", closeMenuItem)
}

func createMenuDateTime() *fyne.Menu {

	jdMenuItem := fyne.NewMenuItem("Julian Day", func() {
		fmt.Println("Julian Day clicked.")
	})
	doeMenuItem := fyne.NewMenuItem("Date of Easter", func() {
		fmt.Println("Day of Easter clicked.")
	})
	return fyne.NewMenu("Date and time", jdMenuItem, doeMenuItem)
}

func createMenuHelp() *fyne.Menu {

	aboutMenuItem := fyne.NewMenuItem("About", func() {
		fmt.Println("About clicked.")
	})
	whatsNewItem := fyne.NewMenuItem("What's new?", func() {
		fmt.Println("What's New clicked.")
	})
	return fyne.NewMenu("Help", aboutMenuItem, whatsNewItem)
}
