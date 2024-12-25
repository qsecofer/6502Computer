package computerui

import (
	"computer/src/computer"
	"computer/src/hitachidisplay"
	"computer/src/uiregister"
	"computer/src/uisevensegmentndidgist"
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const MAXSIZE = 5

type debugData struct {
	Label string
}

type debugDataList struct {
	Elements []debugData
}

func (dd *debugDataList) DeleteLast() {
	if len(dd.Elements) > 0 {
		dd.Elements = dd.Elements[:len(dd.Elements)-1]
	}
}

func (dd *debugDataList) InsertFirst(newElement debugData) {
	if len(dd.Elements) == MAXSIZE {
		dd.DeleteLast()
	}
	dd.Elements = append([]debugData{newElement}, dd.Elements...)
}

type ComputerUI struct {
	computer  *computer.Computer
	dbg       *widget.RichText
	debugList *debugDataList
	pcdsp     *uisevensegmentndidgist.SevenSegmentNDigits
	spdsp     *uisevensegmentndidgist.SevenSegmentNDigits
	acdsp     *uisevensegmentndidgist.SevenSegmentNDigits
	xdsp      *uisevensegmentndidgist.SevenSegmentNDigits
	ydsp      *uisevensegmentndidgist.SevenSegmentNDigits
	acleds    *uiregister.Register
	xleds     *uiregister.Register
	yleds     *uiregister.Register
	flags     *uiregister.Register
	dumpdsp   *widget.RichText

	hdline1 *widget.Label
	hdline2 *widget.Label
}

func New(computer *computer.Computer, canvas fyne.Canvas) *ComputerUI {
	ui := &ComputerUI{
		computer:  computer,
		dbg:       widget.NewRichText(),
		debugList: &debugDataList{Elements: make([]debugData, 0, MAXSIZE)},
		pcdsp:     uisevensegmentndidgist.NewSevenSegmentNDigits(4),
		spdsp:     uisevensegmentndidgist.NewSevenSegmentNDigits(2),
		acdsp:     uisevensegmentndidgist.NewSevenSegmentNDigits(2),
		xdsp:      uisevensegmentndidgist.NewSevenSegmentNDigits(2),
		ydsp:      uisevensegmentndidgist.NewSevenSegmentNDigits(2),
		acleds:    uiregister.NewRegister(),
		xleds:     uiregister.NewRegister(),
		yleds:     uiregister.NewRegister(),
		flags:     uiregister.NewRegister(),
		dumpdsp:   widget.NewRichText(),
		hdline1:   widget.NewLabel("Line 1"),
		hdline2:   widget.NewLabel("Line 2"),
	}
	ui.debugList.InsertFirst(debugData{Label: ui.computer.Cpu.Debug(ui.computer.Bus)})
	ui.setValues()
	canvas.SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		ui.HandleKeyEvent(keyEvent)
	})
	computer.HD.RegisterPackage(ui)
	return ui
}

func (ui *ComputerUI) memoryDump() {
	from := uint16(0x6000)
	dump := ui.computer.Bus.Dump(from, 0xF)
	parts := strings.Split(dump, " ")
	dump = fmt.Sprintf("\n%04X: ", from)
	for i, part := range parts {
		dump += part + " "
		if i%16 == 15 {
			dump += "\n"
			dump += fmt.Sprintf("%04X: ", from+uint16(i+1))
		}
	}
	ui.dumpdsp.ParseMarkdown(dump)
}

func (ui *ComputerUI) setValues() {
	ui.dbg.ParseMarkdown("")
	for _, element := range ui.debugList.Elements {
		ui.dbg.AppendMarkdown(element.Label + "\n")
	}
	ui.pcdsp.SetNumber(uint64(ui.computer.Cpu.PC))
	ui.spdsp.SetNumber(uint64(ui.computer.Cpu.SP))
	ui.acdsp.SetNumber(uint64(ui.computer.Cpu.A))
	ui.xdsp.SetNumber(uint64(ui.computer.Cpu.X))
	ui.ydsp.SetNumber(uint64(ui.computer.Cpu.Y))
	ui.acleds.SetValue(ui.computer.Cpu.A)
	ui.xleds.SetValue(ui.computer.Cpu.X)
	ui.yleds.SetValue(ui.computer.Cpu.Y)
	ui.flags.SetValue(ui.computer.Cpu.Flags2Byte())
	ui.memoryDump()
}

func (ui *ComputerUI) Build() fyne.CanvasObject {

	btnStep := ui.executeButton()
	btnStep.Resize(fyne.NewSize(100, 50))
	btnStep.Move(fyne.NewPos(10, 15))

	ui.dbg.Move(fyne.NewPos(10, 70))

	dspPC := ui.pcdsp.Build()
	dspPC.Move(fyne.NewPos(120, 15))

	dspSP := ui.spdsp.Build()
	dspSP.Move(fyne.NewPos(230, 15))

	dspAC := ui.acdsp.Build()
	dspAC.Move(fyne.NewPos(310, 15))

	dspX := ui.xdsp.Build()
	dspX.Move(fyne.NewPos(375, 15))

	dspY := ui.ydsp.Build()
	dspY.Move(fyne.NewPos(440, 15))

	dispRegAc := container.NewWithoutLayout(ui.acleds.Disp[:]...)
	dispRegAc.Move(fyne.NewPos(500, 8))

	dispRegX := container.NewWithoutLayout(ui.xleds.Disp[:]...)
	dispRegX.Move(fyne.NewPos(500, 26))

	dispRegY := container.NewWithoutLayout(ui.yleds.Disp[:]...)
	dispRegY.Move(fyne.NewPos(500, 44))

	flagsLabel := canvas.NewText("NV  BDIZC", color.Black)
	flagsLabel.Move(fyne.NewPos(610, 15))
	flagsLabel.TextStyle.Monospace = true

	dispFlags := container.NewWithoutLayout(ui.flags.Disp[:]...)
	dispFlags.Move(fyne.NewPos(600, 25))

	dispDump := ui.dumpdsp
	dispDump.Move(fyne.NewPos(120, 70))

	dispLine1 := ui.hdline1
	dispLine1.Move(fyne.NewPos(10, 200))

	dispLine2 := ui.hdline2
	dispLine2.Move(fyne.NewPos(10, 220))

	return container.NewWithoutLayout(
		append([]fyne.CanvasObject{
			ui.dbg, btnStep, dspPC, dspSP, dspAC, dspX, dspY, dispDump, dispLine1, dispLine2},
			dispRegAc, dispRegX, dispRegY, flagsLabel, dispFlags)...)

}

func (ui *ComputerUI) executeButton() fyne.CanvasObject {
	btn := widget.NewButton("Step", func() {
		ui.computer.Cpu.ExecuteInstruction(ui.computer.Bus)
		ui.Update()
	})
	return btn
}

func (ui *ComputerUI) HandleKeyEvent(keyEvent *fyne.KeyEvent) {
}

func (ui *ComputerUI) Update() {
	ui.debugList.InsertFirst(debugData{Label: ui.computer.Cpu.Debug(ui.computer.Bus)})
	ui.setValues()
}

func (ui *ComputerUI) PushData(data hitachidisplay.DisplayData) {
	fmt.Println("PushData", data)
	ui.hdline1.SetText(data.Line1)
	ui.hdline1.Refresh()
	ui.hdline2.SetText(data.Line2)
	ui.hdline2.Refresh()
}
