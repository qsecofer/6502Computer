package hitachidisplay

import (
	"computer/src/bus"
	"fmt"
	"strings"
)

const (
	NR_OFF_LINES = 2
	NR_CHARS     = 40

	RS = 0x20
	RW = 0x40
	E  = 0x80
)

type Command byte

const (
	CmdNoComand Command = iota
	CmdClearDisplay
	CmdReturnHome
	CmdEntryModeSet
	CmdDisplayControl
	CmdCursorShift
	CmdFunctionSet
	CmdSetCGRAMAddr
	CmdSetDDRAMAddr
)

type DispDataPusher interface {
	PushData(displaydata DisplayData)
}

type DisplayData struct {
	Line1 string
	Line2 string
}

type HitachiDisplay struct {
	DispDataPusher DispDataPusher
	DisplayData    DisplayData

	Name  string
	Range bus.AddressRange

	DisplayRam [NR_OFF_LINES * NR_CHARS]string

	buffer        byte
	bufferaddress uint16

	control        byte
	controladdress uint16

	ddra     uint16
	ddradata byte
	ddrb     uint16
	ddrbdata byte

	addrescounter uint8

	registeredPackages []DispDataPusher
}

func New(name string, start, end uint16) *HitachiDisplay {
	hd := &HitachiDisplay{
		Name:               name,
		Range:              bus.AddressRange{Start: start, End: end},
		DisplayRam:         [NR_OFF_LINES * NR_CHARS]string{},
		addrescounter:      0,
		registeredPackages: []DispDataPusher{},
		DisplayData:        DisplayData{},
	}
	hd.bufferaddress = hd.Range.Start
	hd.controladdress = hd.Range.Start + 1
	hd.ddra = hd.Range.Start + 3
	hd.ddrb = hd.Range.Start + 2
	return hd
}

func (hd *HitachiDisplay) executeCmd(data byte) {
	if data == 0 {
		return
	}
	if isExecuteControlCommand(data) {
		commands := map[Command]func(){
			CmdClearDisplay:   hd.clearDisplay,
			CmdReturnHome:     hd.returnHome,
			CmdEntryModeSet:   hd.entryModeSet,
			CmdDisplayControl: hd.displayControl,
			CmdCursorShift:    hd.cursorShift,
			CmdFunctionSet:    hd.functionSet,
			CmdSetCGRAMAddr:   hd.setCGRAMAddr,
			CmdSetDDRAMAddr:   hd.setDDRAMAddr,
		}

		cmd := getcmdindex(hd.buffer)
		if commandFunc, exists := commands[Command(cmd)]; exists {
			commandFunc()
		}
	}
	if isSendChar(data) {
		hd.DisplayRam[hd.addrescounter] = string(hd.buffer)
		hd.addrescounter = (hd.addrescounter + 1) % (NR_OFF_LINES * NR_CHARS)
		hd.PushData(hd.DisplayRam[hd.addrescounter])
	}
}

func (hd *HitachiDisplay) Read(address uint16) byte {
	readFuncs := map[uint16]func() byte{
		hd.bufferaddress:  func() byte { return hd.buffer },
		hd.controladdress: func() byte { return hd.control },
		hd.ddra:           func() byte { return hd.ddradata },
		hd.ddrb:           func() byte { return hd.ddrbdata },
	}
	if readFunc, exists := readFuncs[address]; exists {
		return readFunc()
	}
	return 0
}

func (hd *HitachiDisplay) RespondsTo(address uint16) bool {
	return address >= hd.Range.Start && address <= hd.Range.End
}

func (hd *HitachiDisplay) Write(address uint16, data byte) {
	writeFuncs := map[uint16]func(byte){
		hd.bufferaddress: func(data byte) { hd.buffer = data },
		hd.controladdress: func(data byte) {
			hd.control = data
			hd.executeCmd(data)
		},
		hd.ddra: func(data byte) { hd.ddradata = data },
		hd.ddrb: func(data byte) { hd.ddrbdata = data },
	}
	if writeFunc, exists := writeFuncs[address]; exists {
		writeFunc(data)
	}
}

func getcmdindex(cmd uint8) int8 {
	if cmd&0x80 != 0 {
		return 8
	}
	for i := 1; i < 8; i++ {
		if (cmd<<uint(i))&0x80 != 0 {
			return int8(8 - i)
		}
	}
	return 0
}

// DisplayDataPusher interface
func (hd *HitachiDisplay) PushData(data string) {
	var displayData DisplayData
	displayData.Line1 = strings.Join(hd.DisplayRam[:NR_CHARS], "")
	displayData.Line2 = strings.Join(hd.DisplayRam[NR_CHARS:NR_OFF_LINES*NR_CHARS], "")
	for _, dpp := range hd.registeredPackages {
		dpp.PushData(displayData)
	}
}

func (hd *HitachiDisplay) RegisterPackage(dpp DispDataPusher) {
	hd.registeredPackages = append(hd.registeredPackages, dpp)
}

// Helper functions

func isExecuteControlCommand(cmd uint8) bool {
	return isEset(cmd) && !isRSset(cmd) && !isRWset(cmd)
}

func isSendChar(cmd uint8) bool {
	return isEset(cmd) && isRSset(cmd) && !isRWset(cmd)
}

func isRSset(cmd uint8) bool {
	return cmd&RS != 0
}

func isRWset(cmd uint8) bool {
	return cmd&RW != 0
}

func isEset(cmd uint8) bool {
	return cmd&E != 0
}

// Hitachi control functions

func (hd *HitachiDisplay) clearDisplay() {
	for i := 0; i < NR_OFF_LINES*NR_CHARS; i++ {
		hd.DisplayRam[i] = ""
	}
	hd.PushData(hd.DisplayRam[hd.addrescounter])
}

func (hd *HitachiDisplay) cursorShift() {
	fmt.Println("Cursor Shift")
}

func (hd *HitachiDisplay) displayControl() {
	fmt.Println("Display Control")
}

func (hd *HitachiDisplay) entryModeSet() {
	fmt.Println("Entry Mode Set")
}

func (hd *HitachiDisplay) functionSet() {
	fmt.Println("Function Set")
}

func (hd *HitachiDisplay) returnHome() {
	fmt.Println("Return Home")
}

func (hd *HitachiDisplay) setCGRAMAddr() {
	fmt.Println("Set CGRAM Addr")
}

func (hd *HitachiDisplay) setDDRAMAddr() {
	fmt.Println("Set DDRAM Addr")
}
