package ram

import (
	"computer/src/bus"
)

type ram struct {
	Name   string
	Range  bus.AddressRange
	Memory map[uint16]byte
}

func New(name string, start, end uint16) *ram {
	return &ram{
		Name:   name,
		Range:  bus.AddressRange{Start: start, End: end},
		Memory: make(map[uint16]byte),
	}
}

func (c *ram) Write(address uint16, data byte) {
	c.Memory[address] = data
}

func (c *ram) Read(address uint16) byte {
	if c.Memory == nil {
		c.Memory = make(map[uint16]byte)
	}
	if value, ok := c.Memory[address]; ok {
		return value
	}
	return 0
}

func (c *ram) RespondsTo(address uint16) bool {
	return address >= c.Range.Start && address <= c.Range.End
}
