package screens

import (
	"errors"
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/blackironj/ses-gui/ses"
)

func makeDeleteBtn(window fyne.Window) *widget.Button {
	delBtn := widget.NewButtonWithIcon("Delete a Template", theme.ContentClearIcon(),
		func() {
			if currSelectedTemplate.Text == "" {
				dialog.ShowInformation("Warning", "Please select a template first", window)
				return
			}

			deletionConfirmCallback := func(response bool) {
				if !response {
					return
				}

				err := ses.DeleteSEStemplate(&currSelectedTemplate.Text)
				if err != nil {
					dialog.ShowError(errors.New("Fail to delete"), window)
					fyne.LogError("fail to delete", err)
				} else {
					dialog.ShowInformation("Information", "Success to delete", window)
					updateTemplateList()
				}
			}

			cnf := dialog.NewConfirm("Confirmation", fmt.Sprintf("Are you sure to delete \"%s\"", currSelectedTemplate.Text), deletionConfirmCallback, window)
			cnf.SetDismissText("No")
			cnf.SetConfirmText("Yes")
			cnf.Show()
		})
	return delBtn
}
