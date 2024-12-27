package cpu

func (cpu *CPU) BEQ(operands []byte) {
	if cpu.Z {
		cpu.PC += uint16(int8(operands[0]))
	}
	cpu.PC += uint16(len(operands) + 1)
}
