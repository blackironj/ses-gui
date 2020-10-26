package screens

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func makeSendEmailForm() *widget.Box {
	from := widget.NewEntry()
	from.SetPlaceHolder("test@example.com")
	to := widget.NewEntry()
	to.SetPlaceHolder("test@example.com")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "From", Widget: from},
			{Text: "To", Widget: to},
		},
		OnSubmit: func() {
			//TODO: write a confirm logic
		},
	}

	return widget.NewVBox(
		widget.NewLabelWithStyle("Send a email", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		form,
	)
}
