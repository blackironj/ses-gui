package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"

	"github.com/blackironj/ses-gui/router"
	"github.com/blackironj/ses-gui/screens"
)

func main() {
	var cfg router.RouterConfig
	cfg.Route(router.LoginPath, func(navigator router.Navigator, ctx interface{}) (router.Page, error) {
		return screens.NewLoginPage(navigator)
	})

	cfg.Route(router.MainPath, func(navigator router.Navigator, ctx interface{}) (router.Page, error) {
		return screens.NewMainPage(navigator)
	})

	cfg.InitialPath(router.LoginPath)

	myApp := app.New()
	myWindow := myApp.NewWindow("SES-gui")

	router, err := cfg.Build()
	panicIfErr(err)

	myWindow.SetContent(router)
	myWindow.Resize(fyne.NewSize(640, 480))
	myWindow.ShowAndRun()
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
