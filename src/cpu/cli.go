package cpu

func (cpu *CPU) CLI(operands []byte) {
	cpu.I = false
	cpu.PC += uint16(len(operands) + 1)
}
