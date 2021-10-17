package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/blackironj/ses-gui/screen"
)

func main() {
	a := app.NewWithID("xyz.blackironj.ses-gui")
	w := a.NewWindow("SES-GUI")

	mainView := screen.MainView(w)
	w.SetContent(mainView)
	screen.AskForAccessToAWS(w, a)

	w.Resize(fyne.NewSize(600, 400))

	w.ShowAndRun()
}
