package channel

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"github.com/blackironj/ses-gui/repo"
)

var (
	RefreshTemplateListReq      = make(chan struct{}, 1)
	RefreshEmailVarListReq      = make(chan struct{}, 1)
	RefreshCurrSelectedLabelReq = make(chan struct{}, 1)
)

func RefreshView(templListView, emailVarListView, currTemplateLabel fyne.CanvasObject) {
	for {
		select {
		case <-RefreshTemplateListReq:
			templListView.Refresh()
		case <-RefreshEmailVarListReq:
			emailVarListView.Refresh()
		case <-RefreshCurrSelectedLabelReq:
			currTemplateLabel.(*widget.Label).Text = repo.TemplateList().CurrSelectedTemplate()
			currTemplateLabel.Refresh()
		}
	}
}
