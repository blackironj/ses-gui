package screens

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"runtime"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	"github.com/blackironj/ses-gui/ses"
	"github.com/mitchellh/go-homedir"
)

const (
	downloadDir = "Downloads"
)

func makeDownloadBtn(window fyne.Window) *widget.Button {
	downloadwBtn := widget.NewButtonWithIcon("Download", theme.MoveDownIcon(),
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

			homdir, err := homedir.Dir()
			if err != nil {
				fyne.LogError("fail to get a homedir", err)
				return
			}
			downPath := filepath.Join(homdir, downloadDir)

			writeErr := ioutil.WriteFile(
				filepath.Join(downPath, *output.Template.TemplateName+".html"),
				[]byte(*output.Template.HtmlPart), 0644)
			if writeErr != nil {
				dialog.ShowError(errors.New("Fail to save a template file"), window)
				fyne.LogError("fail to save a file", writeErr)
			}

			infoWin := dialog.NewConfirm("Success", fmt.Sprintf("download path : %s", downPath),
				func(response bool) {
					if response {
						openDir(downPath)
					}
				}, window)

			infoWin.SetDismissText("Close")
			infoWin.SetConfirmText("Open download path")
			infoWin.Show()
		})
	return downloadwBtn
}

func openDir(path string) {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open", path}
	case "windows":
		args = []string{"cmd", "/c", "start", path}
	default:
		args = []string{"xdg-open", path}
	}
	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		fyne.LogError("open in directory", err)
	}
}
