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

func MakeSelectedTemplateIndicator(currTemplate *widget.Label) *fyne.Container {
	return container.NewHBox(
		widget.NewIcon(theme.ConfirmIcon()),
		widget.NewLabelWithStyle("Selected :", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		currTemplate,
	)
}

func MakeSendEmailTitle() *fyne.Container {
	return container.NewCenter(
		container.NewHBox(
			widget.NewIcon(theme.MailSendIcon()),
			widget.NewLabelWithStyle("Send a email", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		),
	)
}

func MakeSendEmailForm(w fyne.Window) *widget.Form {
	from, to := widget.NewEntry(), widget.NewEntry()
	from.SetPlaceHolder("from@example.com")
	to.SetPlaceHolder("to@example.com")

	return &widget.Form{
		Items: []*widget.FormItem{
			{Text: "From", Widget: from},
			{Text: "To", Widget: to},
		},
		OnSubmit: func() {
			//TODO: send a email with template using AWS-SES
		},
	}
}

func MakeEmailVarBox() *fyne.Container {
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

			emailVarListBox.Add(container.NewBorder(nil, nil, widget.NewIcon(theme.ConfirmIcon()), delBtn, entryBox))
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
