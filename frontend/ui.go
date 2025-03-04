package frontend

import "fyne.io/fyne/v2"

func MakeUI(app fyne.App) {
	mainWindow := app.NewWindow("Celestial Math")
	guiMgr := NewGuiMgr(app, mainWindow)
	mainWindow.Resize(fyne.NewSize(1200, 900))
	mainWindow.SetMaster()
	mainWindow.SetMainMenu(CreateMenu(guiMgr))
	homeView := NewHomeView(guiMgr)
	guiMgr.Register("home", homeView)
	mainWindow.SetContent(homeView)

	mainWindow.ShowAndRun()
}
