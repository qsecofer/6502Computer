package uisevensegmentndidgist

import (
	"computer/src/uidisplay7"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type SevenSegmentNDigits struct {
	segments []*uidisplay7.SevenSegment
	value    uint64
	digits   int
}

func NewSevenSegmentNDigits(digits int) *SevenSegmentNDigits {
	s := &SevenSegmentNDigits{
		segments: make([]*uidisplay7.SevenSegment, digits),
		digits:   digits,
	}
	s.value = 0

	for i := 0; i < digits; i++ {
		s.segments[i] = uidisplay7.NewSevenSegment()
	}
	s.SetNumber(s.value)
	return s
}

func (s *SevenSegmentNDigits) SetNumber(n uint64) {
	s.value = n
	for i := 0; i < s.digits; i++ {
		digitValue := (s.value / uint64(1<<uint(4*i))) % 16
		s.segments[s.digits-1-i].SetNumber(int(digitValue))
		// s.segments
	}
}

func (s *SevenSegmentNDigits) Build() fyne.CanvasObject {
	background := canvas.NewRectangle(color.Black)
	background.Resize(fyne.NewSize(25*float32(s.digits), 50))
	cnt := container.NewWithoutLayout()
	cnt.Add(background)
	for i := 0; i < s.digits; i++ {
		segment := s.segments[i].Build()
		segment.Move(fyne.NewPos(float32(5+i*25), 10))
		cnt.Add(segment)
	}
	return cnt
}

func (s *SevenSegmentNDigits) Refresh() {
	for i := 0; i < s.digits; i++ {
		s.segments[i].Refresh()
	}
}
