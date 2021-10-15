package screen

import (
	"errors"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

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

func AskForAccessToAWS(w fyne.Window, a fyne.App) {
	accesskey := widget.NewEntryWithData(binding.BindPreferenceString("login.accesskey", a.Preferences()))
	secretkey := widget.NewPasswordEntry()
	regions := widget.NewSelectEntry(regionList)

	dialog.ShowForm("Access to AWS", "Login", "Clear",
		[]*widget.FormItem{
			widget.NewFormItem("Access Key", accesskey),
			widget.NewFormItem("Secret Key", secretkey),
			widget.NewFormItem("Region", regions),
		}, func(ok bool) {
			if !ok {
				a.Preferences().SetString("login.accesskey", "")
				AskForAccessToAWS(w, a)
				return
			}
			accessToAWS(accesskey.Text, secretkey.Text, regions.Text, w, a)
		}, w)
	if accesskey.Text == "" {
		w.Canvas().Focus(accesskey)
	} else {
		w.Canvas().Focus(secretkey)
	}
}

func accessToAWS(accesskey, secretkey, region string, w fyne.Window, a fyne.App) {
	awsSess, err := ses.NewSession(accesskey, secretkey, region)
	if err != nil {
		log.Println(err)
		showError(errors.New("fail to access to AWS"), w, a)
		return
	}

	_, err = ses.ListSEStemplates()
	if err != nil {
		log.Println(err)
		showError(errors.New("cannot access to template list, please check your keys"), w, a)
		return
	}
	ses.AwsSession = awsSess
	//TODO: save template list metadata
}

func showError(err error, w fyne.Window, a fyne.App) {
	d := dialog.NewError(err, w)
	d.SetOnClosed(func() {
		AskForAccessToAWS(w, a)
	})
	d.Show()
}
