package screens

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	"github.com/blackironj/ses-gui/ses"
	"github.com/mitchellh/go-homedir"
)

func makeDownloadBtn(window fyne.Window) *widget.Button {
	downloadwBtn := widget.NewButtonWithIcon("Download a Template", theme.MoveDownIcon(),
		func() {
			if currSelectedTemplate.Text == "" {
				dialog.ShowInformation("Warning", "Please select a template first", window)
				return
			}

			output, err := ses.GetSEStemplate(&currSelectedTemplate.Text)
			if err != nil {
				dialog.ShowError(errors.New("Fail to download"), window)
				fyne.LogError("fail to get a template", err)
				return
			}

			/*TODO: user should be able to enter a download path
			currently, a downloaded template is saved at home dir
			*/
			downPath, err := homedir.Dir()
			if err != nil {
				fyne.LogError("fail to get a homedir", err)
				return
			}
			downPath = filepath.Join(downPath, *output.Template.TemplateName+".html")

			writeErr := ioutil.WriteFile(downPath, []byte(*output.Template.HtmlPart), 0644)
			if writeErr != nil {
				dialog.ShowError(errors.New("Fail to save a template file"), window)
				fyne.LogError("fail to save a file", writeErr)
			}
			dialog.ShowInformation("Success", fmt.Sprintf("save path : %s", downPath), window)
		})
	return downloadwBtn
}
