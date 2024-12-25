package rom

import (
	"computer/src/bus"
	"fmt"
	"os"
)

type rom struct {
	Name   string
	Range  bus.AddressRange
	Memory map[uint16]byte
}

func New(name string, start, end uint16) *rom {
	rom := &rom{
		Name:   name,
		Range:  bus.AddressRange{Start: start, End: end},
		Memory: make(map[uint16]byte),
	}
	rom.LoadROM("./rom/bios.bin")
	return rom
}

func (rom *rom) Write(address uint16, data byte) {
	fmt.Printf("Can't write to ROM\n")
}

func (rom *rom) Read(address uint16) byte {
	if value, ok := rom.Memory[address-rom.Range.Start]; ok {
		return value
	}
	return 0
}

func (rom *rom) RespondsTo(address uint16) bool {
	return address >= rom.Range.Start && address <= rom.Range.End
}

func (rom *rom) LoadROM(filepath string) error {
	var MemorySize = 0x8000
	romData, err := os.ReadFile(filepath)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	if len(romData) > MemorySize {
		fmt.Printf("ROM size exceeds memory bounds")
		return fmt.Errorf("ROM size exceeds memory bounds")
	}

	for i, bin := range romData {
		rom.Memory[uint16(i)] = bin
	}
	return nil
}

func (rom *rom) Dump(start int) {

	end := start + 0xF

	for i := start; i < end; i++ {
		if i%16 == 0 {
			fmt.Printf("addr: 0x%04X\t", start)
		}
		fmt.Printf("0x%02X ", rom.Memory[uint16(i)])
	}
	println()
}
