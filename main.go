package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/blackironj/ses-gui/screen"
)

func main() {
	a := app.NewWithID("xyz.blackironj.ses-gui")
	w := a.NewWindow("SES-GUI")

	hello := widget.NewLabel("Hello SES GUI!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	screen.AskForAccessToAWS(w, a)

	w.Resize(fyne.NewSize(600, 400))

	w.ShowAndRun()
}
