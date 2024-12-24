package cpu

func (cpu *CPU) CLD(operands []byte) {
	cpu.D = false
	cpu.PC += uint16(len(operands) + 1)
}
