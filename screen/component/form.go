package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	uuid "github.com/satori/go.uuid"

	"github.com/blackironj/ses-gui/repo"
	"github.com/blackironj/ses-gui/screen/channel"
)

func MakeEmailVarList(w fyne.Window) *fyne.Container {
	return container.NewVBox()
}

func MakeAddEmailVarBtn(w fyne.Window, emailVarListBox *fyne.Container) *widget.Button {
	return widget.NewButtonWithIcon("Add email variable", theme.ContentAddIcon(),
		func() {
			id := uuid.NewV4().String()

			keyEntry := widget.NewEntry()
			keyEntry.SetPlaceHolder("key")
			valEntry := widget.NewEntry()
			valEntry.SetPlaceHolder("val")
			entryBox := container.NewVBox(keyEntry, valEntry)

			delBtn := makeDelEmailVarBtn(id, emailVarListBox)

			emailVarListBox.Add(container.NewBorder(nil, nil, nil, delBtn, entryBox))
			repo.EmailVarList().Append(id, keyEntry, valEntry)

			channel.RefreshEmailVarListReq <- struct{}{}
		})
}

func makeDelEmailVarBtn(id string, emailVarListBox *fyne.Container) *widget.Button {
	return widget.NewButtonWithIcon("", theme.ContentRemoveIcon(),
		func() {
			varList := repo.EmailVarList().List()
			for i, v := range varList {
				if id == v.Id {
					repo.EmailVarList().Delete(i)
					emailVarListBox.Remove(emailVarListBox.Objects[i])
					break
				}
			}
			channel.RefreshEmailVarListReq <- struct{}{}
		})
}
