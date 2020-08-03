package main

import (
	"fyne.io/fyne/app"

	"github.com/blackironj/ses-gui/router"
	"github.com/blackironj/ses-gui/screens"
)

const (
	LoginPath = "/login"
)

func main() {
	var cfg router.RouterConfig
	cfg.Route(LoginPath, func(navigator router.Navigator, ctx interface{}) (router.Page, error) {
		return screens.NewLoginPage(navigator)
	})

	cfg.InitialPath(LoginPath)

	myApp := app.New()
	myWindow := myApp.NewWindow("SES-gui")
	router, err := cfg.Build()
	panicIfErr(err)

	myWindow.SetContent(router)
	myWindow.ShowAndRun()
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
