package screen

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"github.com/blackironj/ses-gui/screen/channel"
	"github.com/blackironj/ses-gui/screen/component"
)

func MainView(w fyne.Window) fyne.CanvasObject {
	templList := component.MakeTemplateList(w)
	uploadTemplBtn := component.MakeUploadBtn(w)

	leftSide := container.NewBorder(nil, uploadTemplBtn, nil, nil, templList)

	emailSendFormTitle := component.MakeSendEmailTitle()
	emailSendForm := component.MakeSendEmailForm(w)
	emailVarBox := component.MakeEmailVarBox()
	addEmailVarBtn := component.MakeAddEmailVarBtn(w, emailVarBox)
	addEmailVarView := container.NewBorder(addEmailVarBtn, nil, nil, nil, container.NewVScroll(emailVarBox))

	rightSide := container.NewVSplit(
		container.NewVBox(emailSendFormTitle, emailSendForm),
		addEmailVarView,
	)

	go channel.RefreshView(templList, emailVarBox)

	return container.NewHSplit(leftSide, rightSide)
}
