package screens

import (
	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"

	"github.com/blackironj/ses-gui/ses"
)

var listTemplate widget.Box
var currSelectedTemplate widget.Label

func makeTemplateList() *container.Scroll {
	updateTemplateList()
	return container.NewVScroll(&listTemplate)
}

func updateTemplateList() {
	templates, err := ses.ListSEStemplates()
	if err != nil {
		listTemplate.Children = []fyne.CanvasObject{widget.NewLabel("Fail to access aws-ses")}
		return
	}

	if len(templates) == 0 {
		listTemplate.Children = []fyne.CanvasObject{widget.NewLabel("No templates")}
	}

	btns := make([]fyne.CanvasObject, 0, 20)
	for _, data := range templates {
		templName := *data.Name
		btn := widget.NewButton(templName, func() {
			currSelectedTemplate.SetText(templName)
			currSelectedTemplate.Refresh()
		})
		btns = append(btns, btn)
	}
	listTemplate.Children = btns

	listTemplate.Refresh()
}
