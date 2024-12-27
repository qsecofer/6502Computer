package cpu

func (cpu *CPU) JMPAbsolute(operands []byte) {
	address := uint16(operands[0]) | uint16(operands[1])<<8
	cpu.PC = address
}
