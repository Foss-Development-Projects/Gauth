package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	icon, _ := fyne.LoadResourceFromPath("D:/PROGRAMMING/Production/Projects/Go/gauth/src/assets/images/appicon.png")
	a.SetIcon(icon)
	w := a.NewWindow("Gauth")
	usernameInput := widget.NewMultiLineEntry()
	passwordInput := widget.NewMultiLineEntry()
	usernameLabel := widget.NewLabel("Username")
	passwordLabel := widget.NewLabel("Password")
	form := &widget.Form{
		Items: []*widget.FormItem{
			{
				Text:   usernameLabel.Text,
				Widget: usernameInput,
			},
			{
				Text:   passwordLabel.Text,
				Widget: passwordInput,
			},
		},
		OnSubmit: func() {
			log.Println("UserName Submitted" + usernameInput.Text)
			log.Println("PassWord Submitted" + passwordInput.Text)
			w.Close()

		},
	}
	usernameInput.Resize(fyne.NewSize(400, 30))
	passwordInput.Resize(fyne.NewSize(400, 30))
	w.SetContent(widget.NewLabel("Gauth"))
	w.SetContent(container.NewCenter(form))
	w.Resize(fyne.NewSize(600, 600))

	w.ShowAndRun()
}
