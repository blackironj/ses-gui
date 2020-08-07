package screens

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"

	"github.com/blackironj/ses-gui/router"
	"github.com/blackironj/ses-gui/ses"
)

var regionList = []string{
	"us-east-1",
	"us-east-2",
	"us-west-2",
	"ap-south-1",
	"ap-northeast-2",
	"ap-southeast-1",
	"ap-southeast-2",
	"ap-northeast-1",
	"ca-central-1",
	"eu-central-1",
	"eu-west-1",
	"eu-west-2",
	"sa-east-1",
}

type LoginView struct {
	fyne.Widget
}

func NewLoginPage(navigator router.Navigator) (router.Page, error) {
	errLabel := widget.NewLabel("")

	accessKeyID := widget.NewEntry()
	scretKey := widget.NewPasswordEntry()
	regions := widget.NewSelectEntry(regionList)

	form := widget.NewForm(
		widget.NewFormItem("Access Key ID", accessKeyID),
		widget.NewFormItem("Secret Key", scretKey),
		widget.NewFormItem("Region", regions),
	)

	content := widget.NewVBox(
		layout.NewSpacer(),
		widget.NewLabelWithStyle("Welcome to the SES-GUI", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		layout.NewSpacer(),
		widget.NewHBox(layout.NewSpacer(), form, layout.NewSpacer()),
		errLabel,
	)

	form.OnSubmit = func() {
		if accessKeyID.Text == "" || scretKey.Text == "" || regions.Text == "" {
			errLabel.SetText("Key and region must be set")
		} else {
			var err error
			ses.AwsSession, err = ses.NewSession(accessKeyID.Text, scretKey.Text, regions.Text)
			if err != nil {
				errLabel.SetText("Can not make a session")
				return
			}

			_, err = ses.ListSEStemplates(10)
			if err != nil {
				errLabel.SetText("Can not access. Please check keys")
				return
			}

			//go to list view
			err = navigator.Push(router.ListPath, "")
			if err != nil {
				errLabel.SetText(err.Error())
			}
		}
	}

	return &LoginView{content}, nil
}

func (page *LoginView) BeforeDestroy() {

}
