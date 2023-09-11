package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var (
	a = app.New()
	w = a.NewWindow("Library")
)

func MakeUI() {
	w.Resize(fyne.NewSize(600, 300))
	w.CenterOnScreen()
	w.SetContent(makeTabs())
	w.SetFixedSize(true)

	w.ShowAndRun()
}
