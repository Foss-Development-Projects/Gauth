package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Button struct {
}

func main() {
	a := app.New()
	w := a.NewWindow("Gauth")
	button := widget.NewButtonWithIcon("Update", theme.HomeIcon(), func() {
		log.Println("Updated")
	})
	w.SetContent(widget.NewLabel("Gauth"))
	w.SetContent(button)
	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}
