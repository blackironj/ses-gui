package screens

import (
	"errors"

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

			err := ses.DeleteSEStemplate(&currSelectedTemplate.Text)
			if err != nil {
				dialog.ShowError(errors.New("Fail to delete"), window)
				fyne.LogError("fail to delete", err)
			} else {
				updateTemplateList()
			}
		})
	return delBtn
}