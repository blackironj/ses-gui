package screen

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/blackironj/ses-gui/screen/channel"
	"github.com/blackironj/ses-gui/screen/component"
)

func MainView(w fyne.Window) fyne.CanvasObject {
	tmplListTitle := component.MakeTemplateListTitle()
	templList := component.MakeTemplateList(w)
	uploadTemplBtn := component.MakeUploadBtn(w)

	leftSide := container.NewBorder(tmplListTitle, uploadTemplBtn, nil, nil, templList)

	selectedTmplateLabel := widget.NewLabel("")
	emailSendFormTitle := component.MakeSendEmailTitle()
	emailSendForm := component.MakeSendEmailForm(w)
	selectedTemplateIndicatior := component.MakeSelectedTemplateIndicator(selectedTmplateLabel)

	emailVarBox := component.MakeEmailVarBox()
	addEmailVarBtn := component.MakeAddEmailVarBtn(w, emailVarBox)
	addEmailVarView := container.NewBorder(addEmailVarBtn, nil, nil, nil, container.NewVScroll(emailVarBox))

	rightSide := container.NewVSplit(
		container.NewVBox(
			emailSendFormTitle,
			selectedTemplateIndicatior,
			emailSendForm,
		),
		addEmailVarView,
	)

	go channel.RefreshView(templList, emailVarBox, selectedTmplateLabel)

	return container.NewHSplit(leftSide, rightSide)
}
