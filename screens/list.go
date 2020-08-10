package screens

import (
	"fmt"

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

	btns := make([]fyne.CanvasObject, 0, 20)
	for _, data := range templates {
		name := *data.Name
		btn := widget.NewButton(name, func() {
			currSelectedTemplate.SetText(fmt.Sprintf("Selected Template : %s", name))
			currSelectedTemplate.Refresh()
		})
		btns = append(btns, btn)
	}
	listTemplate.Children = btns

	listTemplate.Refresh()
}
