package channel

import "fyne.io/fyne/v2"

var (
	RefreshTemplateListReq = make(chan struct{}, 1)
	RefreshEmailVarListReq = make(chan struct{}, 1)
)

func RefreshView(templListView, emailVarListView fyne.CanvasObject) {
	for {
		select {
		case <-RefreshTemplateListReq:
			templListView.Refresh()
		case <-RefreshEmailVarListReq:
			emailVarListView.Refresh()
		}
	}
}
