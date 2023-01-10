package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Button struct {
}

func main() {
	a := app.New()
	w := a.NewWindow("Gauth")
	button := widget.NewButton("Update", func() {
		log.Println("Updated")
	})

	button.Resize(fyne.NewSize(100, 20))
	w.SetContent(widget.NewLabel("Gauth"))
	w.SetContent(container.NewCenter(button))
	w.Resize(fyne.NewSize(600, 600))

	w.ShowAndRun()
}
