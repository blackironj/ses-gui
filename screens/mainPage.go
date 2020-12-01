package screens

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	"github.com/blackironj/ses-gui/router"
)

type MainView struct {
	fyne.Widget
}

func NewMainPage(navigator router.Navigator, window fyne.Window) (router.Page, error) {
	uploadBtn := makeUploadViewBtn(window)
	uploadBtnWithListLabel := widget.NewVBox(
		uploadBtn,
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			widget.NewIcon(theme.DocumentIcon()),
			widget.NewLabelWithStyle("Template List", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})),
		widget.NewSeparator(),
	)

	templateList := makeTemplateList()

	left := fyne.NewContainerWithLayout(
		layout.NewBorderLayout(uploadBtnWithListLabel, nil, templateList, nil),
		uploadBtnWithListLabel,
		templateList,
	)

	currSelected := widget.NewHBox(
		fyne.NewContainerWithLayout(
			layout.NewHBoxLayout(),
			widget.NewIcon(theme.ConfirmIcon()),
			widget.NewLabelWithStyle("Selected :", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		),
		&currSelectedTemplate,
	)

	mid := widget.NewVBox(
		currSelected,
		makeDownloadBtn(window),
		makeDeleteBtn(window),
	)

	right := makeSendEmailForm(window)

	content := widget.NewHBox(
		left,
		widget.NewSeparator(),
		mid,
		widget.NewSeparator(),
		right,
	)

	return &MainView{
		content,
	}, nil
}

func (page *MainView) BeforeDestroy() {

}
