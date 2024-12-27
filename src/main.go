package main

import (
	"computer/src/computer"
	"computer/src/computerui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	win := app.NewWindow("Computer")
	win.Resize(fyne.NewSize(750, 600))

	computer := computer.New()
	computerUI := computerui.New(computer, win.Canvas())
	win.SetContent(computerUI.Build())

	win.ShowAndRun()
}
