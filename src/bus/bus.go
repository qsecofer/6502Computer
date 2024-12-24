package bus

import "fmt"

type AddressRange struct {
	Start, End uint16
}

type DisplayData struct {
	Line1 string
	Line2 string
}

type Chip interface {
	Write(address uint16, data byte)
	Read(address uint16) byte
	RespondsTo(address uint16) bool
	ReadDisplay() DisplayData
}

type Bus struct {
	chips []Chip
}

func New() *Bus {
	return &Bus{}
}

func (b *Bus) Attach(chip Chip) {
	b.chips = append(b.chips, chip)
}

func (b *Bus) Write(address uint16, data byte) {
	for _, chip := range b.chips {
		if chip.RespondsTo(address) {
			chip.Write(address, data)
		}
	}
}

func (b *Bus) Read(address uint16) byte {
	for _, chip := range b.chips {
		if chip.RespondsTo(address) {
			return chip.Read(address)
		}
	}
	return 0
}

func (b *Bus) Dump(start uint16, total int) string {
	data := ""
	for i := 0; i < total; i++ {
		data += fmt.Sprintf("%02X ", b.Read(start+uint16(i)))
		if i%16 == 15 {
			data += "\n"
		}
	}
	return data
}

func (b *Bus) ReadDisplay(address uint16) DisplayData {
	for _, chip := range b.chips {
		if chip.RespondsTo(address) {
			return chip.ReadDisplay()
		}
	}
	return DisplayData{
		Line1: "",
		Line2: "",
	}
}
