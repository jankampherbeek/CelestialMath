package frontend

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewHomeView(gm *GuiMgr) fyne.CanvasObject {
	//	toolBar := gm.CreateToolBar()

	buttons := createButtons()
	ariesGlyph := '\uE000'
	taurusGlyph := '\uE001'
	tempLabel := fmt.Sprintf("Right with glyphs %c %c", ariesGlyph, taurusGlyph)
	mainPart := container.NewHSplit(
		widget.NewLabel("Left panel"),
		widget.NewLabel(tempLabel),
	)
	content := container.NewBorder(buttons, nil, nil, mainPart)

	return content
}

func createButtons() *fyne.Container {
	btnHelp := widget.NewButton("Help", func() {})
	btnCancel := widget.NewButton("Cancel", func() {})
	btnExit := widget.NewButton("Exit", func() {})
	return container.New(layout.NewHBoxLayout(), layout.NewSpacer(), btnHelp, btnCancel, btnExit)
}
