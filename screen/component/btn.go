package component

import (
	"strconv"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/blackironj/ses-gui/repo"
)

var testData int

func MakeUploadBtn(refreshReqChan chan<- struct{}) *widget.Button {
	return widget.NewButtonWithIcon("Upload", theme.ContentAddIcon(), func() {
		//TODO: upload a Email template to S3
		repo.Instance().Append("newTemplate" + strconv.Itoa(testData))
		testData++
		refreshReqChan <- struct{}{}
	})
}
