/*
 *  Celestial Math.
 *  Copyright (c) Jan Kampherbeek.
 *  Celestial Math is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package frontend

import (
	"fyne.io/fyne/v2"
	"log"
	"sync"
)

type GuiMgr struct {
	App    fyne.App
	Log    *log.Logger
	window fyne.Window
	views  map[string]fyne.CanvasObject
}

var (
	gmInstance *GuiMgr
	gmOnce     sync.Once
)

func NewGuiMgr(app fyne.App, window fyne.Window) *GuiMgr {

	gmOnce.Do(func() {
		gmInstance = &GuiMgr{
			App:    app,
			window: window,
			views:  make(map[string]fyne.CanvasObject),
		}
	})
	return gmInstance
}

func GetGuiMgr() *GuiMgr {
	if gmInstance == nil {
		panic("Gui manager not initialized")
	}
	return gmInstance
}

func (gm *GuiMgr) Register(name string, view fyne.CanvasObject) {
	gm.views[name] = view
}

func (gm *GuiMgr) Show(name string) {
	if view, ok := gm.views[name]; ok {
		gm.window.SetContent(view)
		gm.window.Show()
	}
}

func (gm *GuiMgr) Refresh(name string) {
	if view, ok := gm.views[name]; ok {
		view.Refresh()
	}
}
