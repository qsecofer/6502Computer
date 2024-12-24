package dsp

import (
	"computer/src/bus"
	"fmt"
)

type dsp struct {
	Name   string
	Range  bus.AddressRange
	Memory map[uint16]byte
}

func New(name string, start, end uint16) *dsp {
	return &dsp{
		Name:   name,
		Range:  bus.AddressRange{Start: start, End: end},
		Memory: make(map[uint16]byte),
	}
}

func (dsp *dsp) Write(address uint16, data byte) {
	dsp.Memory[address] = data
	fmt.Printf("%s\n", string(data))
}

func (dsp *dsp) Read(address uint16) byte {
	if dsp.Memory == nil {
		dsp.Memory = make(map[uint16]byte)
	}
	if value, ok := dsp.Memory[address]; ok {
		return value
	}
	return 0
}

func (dsp *dsp) RespondsTo(address uint16) bool {
	return address >= dsp.Range.Start && address <= dsp.Range.End
}
