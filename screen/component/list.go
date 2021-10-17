package component

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/mitchellh/go-homedir"

	"github.com/blackironj/ses-gui/repo"
	"github.com/blackironj/ses-gui/screen/channel"
	"github.com/blackironj/ses-gui/ses"
)

const (
	_downloadDir = "Downloads"
)

func MakeListTab(w fyne.Window) fyne.CanvasObject {
	list := widget.NewList(
		func() int {
			return repo.Instance().Len()
		},
		func() fyne.CanvasObject {
			buttonBox := container.NewHBox(
				widget.NewButtonWithIcon("", theme.MoveDownIcon(), nil),
				widget.NewButtonWithIcon("", theme.DeleteIcon(), nil),
			)
			return container.NewBorder(nil, nil, nil, buttonBox, widget.NewLabel("template name"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			templateName := repo.Instance().Get(id)
			item.(*fyne.Container).Objects[0].(*widget.Label).SetText(templateName)

			btns := item.(*fyne.Container).Objects[1].(*fyne.Container).Objects
			btns[0].(*widget.Button).OnTapped = func() {
				downloadToLocal(w, &templateName)
			}
			btns[1].(*widget.Button).OnTapped = func() {
				//TODO: delete a email template in S3
				repo.Instance().Delete(id)
				channel.RefreshReq <- struct{}{}
			}
		})
	list.OnSelected = func(id widget.ListItemID) {
	}
	list.OnUnselected = func(id widget.ListItemID) {
	}
	return list
}

func downloadToLocal(w fyne.Window, templateName *string) {
	output, err := ses.GetSEStemplate(templateName)
	if err != nil {
		dialog.ShowError(errors.New("fail to download"), w)
		log.Println("fail to get a template: ", err)
		return
	}

	homdir, err := homedir.Dir()
	if err != nil {
		log.Println("fail to get a homedir: ", err)
		return
	}
	downPath := filepath.Join(homdir, _downloadDir)

	writeErr := ioutil.WriteFile(
		filepath.Join(downPath, *output.Template.TemplateName+".html"),
		[]byte(*output.Template.HtmlPart), 0644)
	if writeErr != nil {
		dialog.ShowError(errors.New("fail to save a template file"), w)
		log.Println("fail to save a file: ", writeErr)
	}

	infoWin := dialog.NewConfirm("Success", fmt.Sprintf("download path : %s", downPath),
		func(response bool) {
			if response {
				openDir(downPath)
			}
		}, w)

	infoWin.SetDismissText("Close")
	infoWin.SetConfirmText("Open download path")
	infoWin.Show()
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
		log.Println("open in directory: ", err)
	}
}

func MakeUploadBtn() *widget.Button {
	return widget.NewButtonWithIcon("Upload", theme.ContentAddIcon(), func() {
		/*TODO: upload a Email template to S3
		implement `searching file` UI
		*/
		channel.RefreshReq <- struct{}{}
	})
}
