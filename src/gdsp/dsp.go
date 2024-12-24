package gdsp

import (
	"computer/src/bus"
	"fmt"
)

type gdsp struct {
	Name   string
	Range  bus.AddressRange
	Memory map[uint16]byte
}

func New(name string, start, end uint16) *gdsp {
	gdsp := &gdsp{
		Name:   name,
		Range:  bus.AddressRange{Start: start, End: end},
		Memory: make(map[uint16]byte),
	}
	return gdsp
}

func (gdsp *gdsp) Write(address uint16, data byte) {
	gdsp.Memory[address] = data
	fmt.Printf("%s\n", string(data))
}

func (gdsp *gdsp) Read(address uint16) byte {
	if gdsp.Memory == nil {
		gdsp.Memory = make(map[uint16]byte)
	}
	if value, ok := gdsp.Memory[address]; ok {
		return value
	}
	return 0
}

func (gdsp *gdsp) RespondsTo(address uint16) bool {
	return address >= gdsp.Range.Start && address <= gdsp.Range.End
}
