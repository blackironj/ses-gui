package screens

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"

	"github.com/blackironj/ses-gui/ses"
)

var listTemplate widget.Box
var currSelectedTemplate widget.Label

func makeTemplateList() *widget.ScrollContainer {
	updateTemplateList()
	return widget.NewVScrollContainer(&listTemplate)
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

	for _, data := range templates {
		btn := widget.NewButton(*data.Name, func() {
			currSelectedTemplate.SetText(*data.Name)
			currSelectedTemplate.Refresh()
		})
		listTemplate.Children = append(listTemplate.Children, btn)
	}
	listTemplate.Refresh()
}
