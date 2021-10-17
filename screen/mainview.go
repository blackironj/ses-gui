package screen

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"github.com/blackironj/ses-gui/screen/channel"
	"github.com/blackironj/ses-gui/screen/component"
)

func MainView() fyne.CanvasObject {
	listTab := component.MakeListTab()
	uploadBtn := component.MakeUploadBtn()

	go channel.RefreshTemplateList(listTab)

	return container.NewBorder(nil, uploadBtn, nil, nil, listTab)
}
