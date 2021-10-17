package channel

import "fyne.io/fyne/v2"

var RefreshReq = make(chan struct{}, 1)

func RefreshTemplateList(listCanvas fyne.CanvasObject) {
	for {
		<-RefreshReq
		listCanvas.Refresh()
	}
}
