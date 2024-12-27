package cpu

func (cpu *CPU) ANDImmediate(operands []byte) {
	cpu.A = cpu.A & operands[0]
	cpu.Z = cpu.A == 0
	cpu.N = cpu.A&0x80 != 0
	cpu.PC += uint16(len(operands) + 1)
}
