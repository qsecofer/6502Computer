package cpu

func (cpu *CPU) JSR(operands []byte) {
	cpu.PC2Stack()
	address := uint16(operands[0]) | uint16(operands[1])<<8
	cpu.PC = address
}
