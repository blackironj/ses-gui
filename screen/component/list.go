package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/blackironj/ses-gui/repo"
)

func MakeListTab(refreshReqChan chan<- struct{}) fyne.CanvasObject {
	list := widget.NewList(
		func() int {
			return repo.Instance().Len()
		},
		func() fyne.CanvasObject {
			buttonBox := container.NewHBox(
				widget.NewButtonWithIcon("", theme.MoveDownIcon(), nil),
				widget.NewButtonWithIcon("", theme.DeleteIcon(), nil),
			)
			return container.NewBorder(nil, nil, nil, buttonBox, widget.NewLabel("template name"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0].(*widget.Label).SetText(repo.Instance().Get(id))

			btns := item.(*fyne.Container).Objects[1].(*fyne.Container).Objects
			btns[0].(*widget.Button).OnTapped = func() {
				//TODO: download a email template from S3
			}
			btns[1].(*widget.Button).OnTapped = func() {
				//TODO: delete a email template in S3
				repo.Instance().Delete(id)
				refreshReqChan <- struct{}{}
			}
		})
	list.OnSelected = func(id widget.ListItemID) {
	}
	list.OnUnselected = func(id widget.ListItemID) {
	}
	return list
}

func MakeUploadBtn(refreshReqChan chan<- struct{}) *widget.Button {
	return widget.NewButtonWithIcon("Upload", theme.ContentAddIcon(), func() {
		/*TODO: upload a Email template to S3
		implement `searching file` UI
		*/
		refreshReqChan <- struct{}{}
	})
}
