package cpu

import "fmt"

func (cpu *CPU) RTI(operands []byte) {
	cpu.Stack2PC()
	fmt.Printf("PC: %02X\n", cpu.PC)
}
