package screens

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/storage"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"

	sessdk "github.com/aws/aws-sdk-go/service/ses"

	"github.com/blackironj/ses-gui/ses"
)

func makeUploadViewBtn(window fyne.Window) *widget.Button {
	uploadViewBtn := widget.NewButtonWithIcon("Upload", theme.ContentAddIcon(),
		func() {
			templateName := widget.NewEntry()
			subject := widget.NewEntry()

			filePath := widget.NewEntry()
			findPathBtn := makeFindHTMLbtn(window, filePath)
			path := widget.NewHBox(filePath, findPathBtn)

			contents := widget.NewForm(
				widget.NewFormItem("Template name", templateName),
				widget.NewFormItem("Subject", subject),
				widget.NewFormItem("Path", path),
			)

			dialog.ShowCustomConfirm("Upload a Template", "upload", "cancel", contents,
				func(b bool) {
					if !b {
						return
					}

					if filepath.Ext(filePath.Text) != ".html" {
						err := errors.New("Please load a html file")
						dialog.ShowError(err, window)
						return
					}

					htmlFile, readFileErr := ioutil.ReadFile(filePath.Text)
					if readFileErr != nil {
						fyne.LogError("Failed to read a file", readFileErr)

						err := errors.New("Failed to read a file")
						dialog.ShowError(err, window)
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
						fyne.LogError("Fail to upload", uploadErr)

						err := errors.New("Fail to upload")
						dialog.ShowError(err, window)
					} else {
						dialog.ShowInformation("Information", "Success to upload", window)
						updateTemplateList()
					}
				}, window)
		})
	return uploadViewBtn
}

func makeFindHTMLbtn(window fyne.Window, filePath *widget.Entry) *widget.Button {
	findBtn := widget.NewButtonWithIcon("find", theme.SearchIcon(),
		func() {
			fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
				if err == nil && reader == nil {
					return
				}
				if err != nil {
					dialog.ShowError(err, window)
					return
				}
				path := getHTMLpath(reader)
				filePath.SetText(path)
			}, window)

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
		fyne.LogError("Failed to close stream", err)
	}
	return
}
