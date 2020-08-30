package screens

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"github.com/blackironj/ses-gui/router"
)

type MainView struct {
	fyne.Widget
}

func NewMainPage(navigator router.Navigator, window fyne.Window) (router.Page, error) {
	uploadBtn := makeUploadViewBtn(window)
	templateList := makeTemplateList()

	left := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(uploadBtn, nil, templateList, nil),
		uploadBtn,
		templateList,
	)

	currSelected := widget.NewHBox(
		widget.NewLabel("Selected : "),
		&currSelectedTemplate,
	)

	downloadBtn := makeDownloadBtn(window)
	deteleteBtn := makeDeleteBtn(window)

	right := widget.NewVBox(
		currSelected,
		downloadBtn,
		deteleteBtn,
	)

	content := widget.NewHBox(left, right)
	return &MainView{
		content,
	}, nil
}

func (page *MainView) BeforeDestroy() {

}
