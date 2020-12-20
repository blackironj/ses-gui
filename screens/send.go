package screens

import (
	"errors"
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	uuid "github.com/satori/go.uuid"

	"github.com/blackironj/ses-gui/ses"
)

const entryCapacitorSize = 5

type entryWithId struct {
	keyEntry *widget.Entry
	valEntry *widget.Entry
	id       string
}

func makeSendEmailForm(window fyne.Window) *widget.Box {
	from := widget.NewEntry()
	from.SetPlaceHolder("test@example.com")
	to := widget.NewEntry()
	to.SetPlaceHolder("test@example.com")

	entries := make([]*entryWithId, 0, entryCapacitorSize)

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "From", Widget: from},
			{Text: "To", Widget: to},
		},
		OnSubmit: func() {
			dataMap := make(map[string]interface{})

			for _, data := range entries {
				dataMap[data.keyEntry.Text] = data.valEntry.Text
			}

			err := ses.SendEmailWithTemplate(from.Text, to.Text, currSelectedTemplate.Text, dataMap)
			if err != nil {
				dialog.ShowError(errors.New("Fail to send"), window)
				fyne.LogError("fail to send a email", err)
				return
			}

			dialog.ShowInformation("Success", fmt.Sprintf("sending email to %s", to.Text), window)
		},
	}

	templateItemBox := widget.NewVBox()

	addTemplDataBtn := &widget.Button{
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

			entries = append(entries, &entryWithId{
				id:       u1.String(),
				keyEntry: keyData,
				valEntry: valData,
			})

			delBtn := &widget.Button{
				Icon: theme.ContentRemoveIcon(),
				OnTapped: func() {
					id := u1.String()
					for i, v := range entries {
						if id == v.id {
							entries = append(entries[:i], entries[i+1:]...)
							templateItemBox.Children = append(templateItemBox.Children[:i], templateItemBox.Children[i+1:]...)
							break
						}
					}
					templateItemBox.Refresh()
				},
			}

			newItem := widget.NewHBox(label, keyData, valData, delBtn)
			templateItemBox.Append(newItem)
		},
	}

	return widget.NewVBox(
		widget.NewLabelWithStyle("Send a email", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		form,
		addTemplDataBtn,
		templateItemBox,
	)
}
