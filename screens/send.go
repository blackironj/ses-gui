package screens

import (
	"errors"
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	uuid "github.com/satori/go.uuid"

	"github.com/blackironj/ses-gui/ses"
)

type orderLabelWithId struct {
	orderLabel *widget.Label
	id         string
}

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

	templateItemBox := widget.NewVBox()
	orderLabels := []orderLabelWithId{}

	return widget.NewVBox(
		widget.NewLabelWithStyle("Send a email", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		form,
		makeAddTemplateDataBtnWithEntry(templateItemBox, orderLabels),
		templateItemBox,
	)
}

func updateNumLabels(numberLabels []orderLabelWithId) {
	for i := range numberLabels {
		numberLabels[i].orderLabel.Text = strconv.Itoa(i+1) + ". "
	}
}

func makeAddTemplateDataBtnWithEntry(templateItemBox *widget.Box, orderLabels []orderLabelWithId) *widget.Button {
	addTemplateDataBtnWithEntry := &widget.Button{
		Alignment: widget.ButtonAlignLeading,
		Icon:      theme.ContentAddIcon(),
		Text:      "Add template data",
		OnTapped: func() {
			u1 := uuid.NewV4()
			label := widget.NewLabel("")

			orderLabels = append(orderLabels, orderLabelWithId{
				id:         u1.String(),
				orderLabel: label,
			})

			entry := widget.NewEntry()
			delBtn := &widget.Button{
				Icon: theme.ContentRemoveIcon(),
				OnTapped: func() {
					id := u1.String()
					for i, v := range orderLabels {
						if id == v.id {
							orderLabels = append(orderLabels[:i], orderLabels[i+1:]...)
							updateNumLabels(orderLabels)

							templateItemBox.Children = append(templateItemBox.Children[:i], templateItemBox.Children[i+1:]...)
							break
						}
					}
					templateItemBox.Refresh()
				},
			}

			newItem := widget.NewHBox(label, entry, delBtn)
			updateNumLabels(orderLabels)

			templateItemBox.Append(newItem)
		},
	}
	return addTemplateDataBtnWithEntry
}
