package screen

import "fyne.io/fyne/v2"

func RefreshList(refreshReqChan <-chan struct{}, listCanvas fyne.CanvasObject) {
	for {
		<-refreshReqChan
		listCanvas.Refresh()
	}
}
