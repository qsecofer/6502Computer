package uiregister

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Register struct {
	value byte
	Disp  [8]fyne.CanvasObject
}

func NewRegister() *Register {
	reg := &Register{}
	reg.value = 0
	for i := 0; i < 8; i++ {
		reg.Disp[i] = canvas.NewCircle(color.Black)
		reg.Disp[i].Resize(fyne.NewSize(8, 8))
		reg.Disp[i].Move(fyne.NewPos(float32(10+10*(7-i)), float32(10)))
	}
	reg.SetValue(1)
	return reg
}

func (r *Register) SetValue(val byte) {
	r.value = val
	for i := 0; i < 8; i++ {
		if val&(1<<uint(i)) != 0 {
			r.Disp[i].(*canvas.Circle).FillColor = color.RGBA{255, 0, 0, 255}
		} else {
			r.Disp[i].(*canvas.Circle).FillColor = color.RGBA{155, 155, 155, 255}
		}
		r.Disp[i].Refresh()
	}
}
