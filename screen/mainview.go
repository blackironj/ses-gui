package screen

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"github.com/blackironj/ses-gui/screen/component"
)

func MainView() fyne.CanvasObject {
	refreshReqChan := make(chan struct{})

	listTab := component.MakeListTab(refreshReqChan)
	uploadBtn := component.MakeUploadBtn(refreshReqChan)

	go refreshList(refreshReqChan, listTab)

	return container.NewBorder(nil, uploadBtn, nil, nil, listTab)
}

func refreshList(refreshReqChan <-chan struct{}, listCanvas fyne.CanvasObject) {
	for {
		<-refreshReqChan
		listCanvas.Refresh()
	}
}
