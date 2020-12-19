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

type labelAndEntryWithId struct {
	orderLabel *widget.Label
	keyEntry   *widget.Entry
	valEntry   *widget.Entry
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
	labelAndEntry := []labelAndEntryWithId{}

	return widget.NewVBox(
		widget.NewLabelWithStyle("Send a email", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		form,
		makeAddTemplateDataBtnWithEntry(templateItemBox, labelAndEntry),
		templateItemBox,
	)
}

func updateNumLabels(numberLabels []labelAndEntryWithId) {
	for i := range numberLabels {
		numberLabels[i].orderLabel.Text = strconv.Itoa(i+1) + ". "
	}
}

func makeAddTemplateDataBtnWithEntry(templateItemBox *widget.Box, orderlabelAndDataEntry []labelAndEntryWithId) *widget.Button {
	addTemplateDataBtnWithEntry := &widget.Button{
		Alignment: widget.ButtonAlignLeading,
		Icon:      theme.ContentAddIcon(),
		Text:      "Add template data",
		OnTapped: func() {
			u1 := uuid.NewV4()
			label := widget.NewLabel("")

			keyData := widget.NewEntry()
			keyData.SetPlaceHolder("key")

			valData := widget.NewEntry()
			valData.SetPlaceHolder("val")

			orderlabelAndDataEntry = append(orderlabelAndDataEntry, labelAndEntryWithId{
				id:         u1.String(),
				keyEntry:   keyData,
				valEntry:   valData,
				orderLabel: label,
			})

			delBtn := &widget.Button{
				Icon: theme.ContentRemoveIcon(),
				OnTapped: func() {
					id := u1.String()
					for i, v := range orderlabelAndDataEntry {
						if id == v.id {
							orderlabelAndDataEntry = append(orderlabelAndDataEntry[:i], orderlabelAndDataEntry[i+1:]...)
							updateNumLabels(orderlabelAndDataEntry)

							templateItemBox.Children = append(templateItemBox.Children[:i], templateItemBox.Children[i+1:]...)
							break
						}
					}
					templateItemBox.Refresh()
				},
			}

			newItem := widget.NewHBox(label, keyData, valData, delBtn)
			updateNumLabels(orderlabelAndDataEntry)

			templateItemBox.Append(newItem)
		},
	}
	return addTemplateDataBtnWithEntry
}
