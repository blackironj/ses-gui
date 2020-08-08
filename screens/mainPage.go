package screens

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"github.com/blackironj/ses-gui/router"
)

type MainView struct {
	fyne.Widget
}

func NewMainPage(navigator router.Navigator, window fyne.Window) (router.Page, error) {
	sample := widget.NewMultiLineEntry()
	sample.SetReadOnly(true)
	/*TODO: 1. Get template list from aws-ses
	2. make buttons by template name*/
	var buttons []fyne.CanvasObject
	for i := 1; i <= 10; i++ {
		index := i
		newButton := widget.NewButton(fmt.Sprintf("Button %d", index), func() {
			//TODO: get a template from aws ses
			sample.SetText(fmt.Sprintf("Button %d", index))
			sample.Refresh()
		})
		buttons = append(buttons, newButton)
	}

	uploadBtn := makeUploadViewBtn(window)

	templateBtns := widget.NewVScrollContainer(widget.NewVBox(buttons...))

	left := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(uploadBtn, nil, templateBtns, nil),
		uploadBtn,
		templateBtns,
	)
	right := widget.NewVScrollContainer(sample)

	content := widget.NewHBox(left, right)
	return &MainView{
		content,
	}, nil
}

func (page *MainView) BeforeDestroy() {

}
