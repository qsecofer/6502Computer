package uidisplay7

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

const (
	SEGMENTS  = 8
	THICKNESS = float32(2)
	LENGTH    = float32(12)
)

var off = color.RGBA{255, 0, 0, 115}
var on = color.RGBA{255, 0, 0, 255}

var segmentPositions = [SEGMENTS]fyne.Position{
	{X: THICKNESS, Y: 0},
	{X: THICKNESS + LENGTH, Y: THICKNESS},
	{X: THICKNESS + LENGTH, Y: THICKNESS*2 + LENGTH},
	{X: THICKNESS, Y: THICKNESS*2 + LENGTH*2},
	{X: 0, Y: THICKNESS*2 + LENGTH},
	{X: 0, Y: THICKNESS},
	{X: THICKNESS, Y: THICKNESS + LENGTH},
}

type SevenSegment struct {
	segments [SEGMENTS]*canvas.Rectangle
}

func NewSevenSegment() *SevenSegment {
	s := &SevenSegment{}

	for i := 0; i < SEGMENTS; i++ {
		segment := canvas.NewRectangle(off)
		segment.Resize(fyne.NewSize(
			LENGTH,
			THICKNESS,
		))

		if i == 1 || i == 2 || i == 4 || i == 5 {
			segment.Resize(fyne.NewSize(THICKNESS, LENGTH))
		}

		segment.Move(segmentPositions[i])
		s.segments[i] = segment
	}

	return s
}

func (s *SevenSegment) SetNumber(n int) {
	segments := [16][SEGMENTS]bool{
		{true, true, true, true, true, true, false},     // 0
		{false, true, true, false, false, false, false}, // 1
		{true, true, false, true, true, false, true},    // 2
		{true, true, true, true, false, false, true},    // 3
		{false, true, true, false, false, true, true},   // 4
		{true, false, true, true, false, true, true},    // 5
		{false, false, true, true, true, true, true},    // 6
		{true, true, true, false, false, false, false},  // 7
		{true, true, true, true, true, true, true},      // 8
		{true, true, true, true, false, true, true},     // 9
		{true, true, true, false, true, true, true},     // A
		{false, false, true, true, true, true, true},    // b
		{true, false, false, true, true, true, false},   // C
		{false, true, true, true, true, false, true},    // d
		{true, false, false, true, true, true, true},    // E
		{true, false, false, false, true, true, true},   // F
	}

	for i, segment := range s.segments {
		if segments[n][i] {
			segment.FillColor = on
		} else {
			segment.FillColor = off
		}
	}
}

func (s *SevenSegment) Build() fyne.CanvasObject {
	dsp := container.NewWithoutLayout(
		s.segments[0],
		s.segments[1],
		s.segments[2],
		s.segments[3],
		s.segments[4],
		s.segments[5],
		s.segments[6],
	)
	return dsp
}

func (s *SevenSegment) Refresh() {
	for _, segment := range s.segments {
		segment.Refresh()
	}
}
