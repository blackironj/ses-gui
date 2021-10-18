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

	emailVarList := component.MakeEmailVarList(w)
	addEmailVarBtn := component.MakeAddEmailVarBtn(w, emailVarList)

	go channel.RefreshView(templList, emailVarList)

	leftSide := container.NewBorder(nil, uploadTemplBtn, nil, nil, templList)
	rightSide := container.NewBorder(addEmailVarBtn, nil, nil, nil, container.NewVScroll(emailVarList))

	return container.NewHSplit(leftSide, rightSide)
}
