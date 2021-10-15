package component

import (
	"log"

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
				log.Println("download button is clicked")
			}
			btns[1].(*widget.Button).OnTapped = func() {
				log.Println("delete button is clicked : ", id)
				repo.Instance().Delete(id)
				refreshReqChan <- struct{}{}
			}
		})
	list.OnSelected = func(id widget.ListItemID) {
		log.Println("selected item : ", id)
	}
	list.OnUnselected = func(id widget.ListItemID) {
		log.Println("unselected item : ", id)
	}

	return list
}
