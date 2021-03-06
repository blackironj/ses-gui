package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"

	"github.com/blackironj/ses-gui/router"
	"github.com/blackironj/ses-gui/screens"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())

	myWindow := myApp.NewWindow("SES-gui")

	var cfg router.RouterConfig
	cfg.Route(router.LoginPath, func(navigator router.Navigator, ctx interface{}) (router.Page, error) {
		return screens.NewLoginPage(navigator)
	})

	cfg.Route(router.MainPath, func(navigator router.Navigator, ctx interface{}) (router.Page, error) {
		return screens.NewMainPage(navigator, myWindow)
	})

	cfg.InitialPath(router.LoginPath)

	router, err := cfg.Build()
	panicIfErr(err)

	myWindow.SetContent(router)
	myWindow.Resize(fyne.NewSize(600, 400))
	myWindow.ShowAndRun()
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
