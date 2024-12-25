package computer

import (
	"computer/src/bus"
	"computer/src/cpu"
	"computer/src/hitachidisplay"
	"computer/src/ram"
	"computer/src/rom"
)

type Computer struct {
	Bus *bus.Bus
	Cpu *cpu.CPU
	HD  *hitachidisplay.HitachiDisplay
}

func New() *Computer {
	c := &Computer{}
	c.Bus = bus.New()
	c.Cpu = cpu.New(c.Bus)

	c.Bus.Attach(ram.New("RAM", 0x0000, 0x5FFF))
	c.Bus.Attach(rom.New("ROM", 0x8000, 0xFFFF))

	c.HD = hitachidisplay.New("HITACHI", 0x6000, 0x6003)
	c.Bus.Attach(c.HD)

	return c
}
