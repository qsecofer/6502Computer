package cpu

import "fmt"

func (cpu *CPU) STYAbsolute(operands []byte) {
	address := uint16(operands[0]) | uint16(operands[1])<<8
	fmt.Printf("STY $%04X\n", address)
	fmt.Printf("Y: $%02X\n", cpu.Y)
	cpu.bus.Write(address, cpu.Y)
	cpu.PC += uint16(len(operands) + 1)
}
