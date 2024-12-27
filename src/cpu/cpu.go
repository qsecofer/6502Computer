package cpu

import (
	"computer/src/bus"
	"fmt"
	"os"
)

type CPU struct {
	PC  uint16
	SP  uint8
	A   uint8
	X   uint8
	Y   uint8
	N   bool
	V   bool
	B   bool
	D   bool
	I   bool
	Z   bool
	C   bool
	bus *bus.Bus
}

func New(bus *bus.Bus) *CPU {
	cpu := &CPU{
		PC: 0x8000,
		SP: 0xFF,
		A:  0x0,
		X:  0x0,
		Y:  0x0,
		N:  false, V: false, B: false, D: false, I: false, Z: false, C: false,
		bus: bus,
	}
	return cpu
}

func (c *CPU) getOperands(bus *bus.Bus) []byte {
	opcode := bus.Read(c.PC)
	operands := make([]byte, 0)

	for i := 0; i < Opcodes[opcode].Operands; i++ {
		operands = append(operands, bus.Read(c.PC+uint16(i+1)))
	}

	return operands
}

func (cpu *CPU) Debug(bus *bus.Bus) string {
	opcode := bus.Read(cpu.PC)
	mnemonic := Opcodes[opcode].Mnemonic
	var output string
	if mnemonic == "" {
		fmt.Printf("Unknown opcode: $%02X\n", opcode)
		os.Exit(-1)
	}
	// fmt.Printf("%s", mnemonic)
	output = mnemonic
	if Opcodes[opcode].Mode == Immediate {
		output += fmt.Sprintf(" #$%02X    ", bus.Read(cpu.PC+1))
	}
	if Opcodes[opcode].Mode == Implied {
		output += ("         ")
	}
	if Opcodes[opcode].Mode == Absolute {
		output += (fmt.Sprintf(" $%02X%02X   ", bus.Read(cpu.PC+2), bus.Read(cpu.PC+1)))
	}
	if Opcodes[opcode].Mode == Accumulator {
		output += (" A       ")
	}
	return output
}

func (cpu *CPU) ExecuteInstruction(bus *bus.Bus) {
	opcode := bus.Read(cpu.PC)
	operands := cpu.getOperands(bus)
	Opcodes[opcode].Fn(cpu, operands)
}

func (cpu *CPU) PC2Stack() {
	cpu.bus.Write(0x100+uint16(cpu.SP), byte(cpu.PC>>8))
	cpu.SP--
	cpu.bus.Write(0x100+uint16(cpu.SP), byte(cpu.PC))
	cpu.SP--
}

func (cpu *CPU) Stack2PC() {
	cpu.SP++
	cpu.PC = uint16(cpu.bus.Read(0x100 + uint16(cpu.SP)))
	cpu.SP++
	cpu.PC |= uint16(cpu.bus.Read(0x100+uint16(cpu.SP))) << 8
}

func (cpu *CPU) NMI() {
	cpu.D = false
	cpu.PC2Stack()

	low := cpu.bus.Read(0xFFFA)
	high := cpu.bus.Read(0xFFFB)
	cpu.PC = uint16(high)<<8 | uint16(low)

}

func (cpu *CPU) IRQ() {
	if cpu.I {
		return
	}
	cpu.D = false
	cpu.PC2Stack()

	cpu.I = false
	cpu.B = false
	low := cpu.bus.Read(0xFFFE)
	high := cpu.bus.Read(0xFFFF)
	cpu.PC = uint16(high)<<8 | uint16(low)
}

func (cpu *CPU) Flags2Byte() byte {
	var flags byte
	if cpu.N {
		flags |= 0x80
	}
	if cpu.V {
		flags |= 0x40
	}
	if cpu.B {
		flags |= 0x10
	}
	if cpu.D {
		flags |= 0x08
	}
	if cpu.I {
		flags |= 0x04
	}
	if cpu.Z {
		flags |= 0x02
	}
	if cpu.C {
		flags |= 0x01
	}
	return flags
}

func (cpu *CPU) Byte2Flags(flags byte) {
	cpu.N = flags&0x80 != 0
	cpu.V = flags&0x40 != 0
	cpu.B = flags&0x10 != 0
	cpu.D = flags&0x08 != 0
	cpu.I = flags&0x04 != 0
	cpu.Z = flags&0x02 != 0
	cpu.C = flags&0x01 != 0
}

func (c *CPU) String() string {
	return fmt.Sprintf("PC: 0x%04X  "+
		"SP: 0x%02X    "+
		"A: 0x%02X %08b  "+
		"X: 0x%02X %08b  "+
		"Y: 0x%02X    "+
		"N: %v V: %v B: %v D: %v I: %v Z: %v C: %v",
		c.PC, c.SP, c.A, c.A, c.X, c.X, c.Y, c.N, c.V, c.B, c.D, c.I, c.Z, c.C)
}
