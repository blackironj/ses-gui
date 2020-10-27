package screens

import (
	"errors"
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"

	"github.com/blackironj/ses-gui/ses"
)

func makeSendEmailForm(window fyne.Window) *widget.Box {
	from := widget.NewEntry()
	from.SetPlaceHolder("test@example.com")
	to := widget.NewEntry()
	to.SetPlaceHolder("test@example.com")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "From", Widget: from},
			{Text: "To", Widget: to},
		},
		OnSubmit: func() {
			//TODO: if template data exist in email-template, datas must be input from user

			err := ses.SendEmailWithTemplate(from.Text, to.Text, currSelectedTemplate.Text)
			if err != nil {
				dialog.ShowError(errors.New("Fail to send"), window)
				fyne.LogError("fail to send a email", err)
				return
			}

			dialog.ShowInformation("Success", fmt.Sprintf("sending email to %s", to.Text), window)
		},
	}

	return widget.NewVBox(
		widget.NewLabelWithStyle("Send a email", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		form,
	)
}
