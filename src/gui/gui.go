package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type Gui struct {
	Win   fyne.Window
	Title binding.String
}

func (g *Gui) MakeGUI() fyne.CanvasObject {

	debugStep := widget.NewLabelWithData(g.Title)

	return widget.NewLabel(debugStep.Text)
}
