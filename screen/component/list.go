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
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	sessdk "github.com/aws/aws-sdk-go/service/ses"
	"github.com/mitchellh/go-homedir"

	"github.com/blackironj/ses-gui/repo"
	"github.com/blackironj/ses-gui/screen/channel"
	"github.com/blackironj/ses-gui/ses"
)

const (
	_downloadDir = "Downloads"
)

func MakeTemplateListTitle() *fyne.Container {
	return container.NewCenter(
		container.NewHBox(
			widget.NewIcon(theme.DocumentIcon()),
			widget.NewLabelWithStyle("Template List", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		),
	)
}

func MakeTemplateList(w fyne.Window) *widget.List {
	list := widget.NewList(
		func() int {
			return repo.TemplateList().Len()
		},
		func() fyne.CanvasObject {
			buttonBox := container.NewHBox(
				widget.NewButtonWithIcon("", theme.MoveDownIcon(), nil),
				widget.NewButtonWithIcon("", theme.DeleteIcon(), nil),
			)
			return container.NewBorder(nil, nil, nil, buttonBox, widget.NewLabel("template name"))
		},
		func(itemID widget.ListItemID, item fyne.CanvasObject) {
			templateName := repo.TemplateList().Get(itemID)
			item.(*fyne.Container).Objects[0].(*widget.Label).SetText(templateName)

			btns := item.(*fyne.Container).Objects[1].(*fyne.Container).Objects
			btns[0].(*widget.Button).OnTapped = func() {
				downloadToLocal(w, templateName)
			}
			btns[1].(*widget.Button).OnTapped = func() {
				deleteFromS3(w, templateName, itemID)
			}
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		templName := repo.TemplateList().Get(id)
		repo.TemplateList().SetCurrSelectedTemplate(templName)
		channel.RefreshCurrSelectedLabelReq <- struct{}{}
	}
	return list
}

func deleteFromS3(w fyne.Window, templateName string, itemID int) {
	deletionConfirmCallback := func(response bool) {
		if !response {
			return
		}

		err := ses.DeleteSEStemplate(&templateName)
		if err != nil {
			log.Println("fail to delete a template: ", err)
			dialog.ShowError(errors.New("fail to delete\n"+err.Error()), w)
			return
		}
		repo.TemplateList().Delete(itemID)
		channel.RefreshTemplateListReq <- struct{}{}
		dialog.ShowInformation("Information", "success to delete", w)
	}

	cnf := dialog.NewConfirm("Confirmation",
		fmt.Sprintf("Are you sure to delete \"%s\"", templateName),
		deletionConfirmCallback, w)
	cnf.SetDismissText("No")
	cnf.SetConfirmText("Yes")
	cnf.Show()
}

func downloadToLocal(w fyne.Window, templateName string) {
	output, err := ses.GetSEStemplate(&templateName)
	if err != nil {
		log.Println("fail to get a template: ", err)
		dialog.ShowError(errors.New("fail to download\n"+err.Error()), w)
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
		log.Println("fail to save a file: ", writeErr)
		dialog.ShowError(errors.New("fail to save a template file\n"+writeErr.Error()), w)
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

func MakeUploadBtn(w fyne.Window) *widget.Button {
	return widget.NewButtonWithIcon("Upload", theme.ContentAddIcon(), func() {
		templateName := widget.NewEntry()
		subject := widget.NewEntry()

		filePath := widget.NewEntry()
		findPathBtn := makeFindHTMLbtn(w, filePath)
		path := container.NewBorder(nil, nil, nil, findPathBtn, filePath)

		contents := widget.NewForm(
			widget.NewFormItem("Template name", templateName),
			widget.NewFormItem("Subject", subject),
			widget.NewFormItem("Path", path),
		)

		uploadForm := dialog.NewCustomConfirm("Upload a Template", "upload", "cancel", contents,
			func(ok bool) {
				if !ok {
					return
				}

				if filepath.Ext(filePath.Text) != ".html" {
					dialog.ShowError(errors.New("please load a html file"), w)
					return
				}

				htmlFile, readFileErr := ioutil.ReadFile(filePath.Text)
				if readFileErr != nil {
					log.Println("failed to read a file: ", readFileErr)
					dialog.ShowError(errors.New("failed to read a file\n"+readFileErr.Error()), w)
					return
				}

				contents := string(htmlFile)
				inputTemplate := &sessdk.Template{
					HtmlPart:     &contents,
					TemplateName: &templateName.Text,
					SubjectPart:  &subject.Text,
				}

				uploadErr := ses.UploadSEStemplate(inputTemplate)
				if uploadErr != nil {
					log.Println("fail to upload: ", uploadErr)
					dialog.ShowError(errors.New("fail to upload\n"+uploadErr.Error()), w)
					return
				}
				dialog.ShowInformation("Information", "Success to upload", w)
				repo.TemplateList().Append(templateName.Text)
				channel.RefreshTemplateListReq <- struct{}{}
			}, w)

		uploadForm.Resize(fyne.NewSize(400, 250))
		uploadForm.Show()
	})
}

func makeFindHTMLbtn(w fyne.Window, filePath *widget.Entry) *widget.Button {
	findBtn := widget.NewButtonWithIcon("find", theme.SearchIcon(),
		func() {
			fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err == nil && reader == nil {
					return
				}
				if err != nil {
					dialog.ShowError(err, w)
					return
				}
				path := getHTMLpath(reader)
				filePath.SetText(path)
			}, w)

			fd.SetFilter(storage.NewExtensionFileFilter([]string{".html"}))
			fd.Show()
		},
	)
	return findBtn
}

func getHTMLpath(f fyne.URIReadCloser) (path string) {
	if f == nil {
		return
	}

	ext := f.URI().Extension()
	if ext == ".html" {
		path = f.URI().String()
		scheme := f.URI().Scheme() + "://"

		return path[len(scheme):]
	}
	err := f.Close()
	if err != nil {
		log.Println("failed to close stream: ", err)
	}
	return
}
