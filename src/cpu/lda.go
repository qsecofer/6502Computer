package cpu

func (cpu *CPU) LDAImmediate(operands []byte) {
	cpu.A = operands[0]
	cpu.ZFlag(cpu.A)
	cpu.NFlag(cpu.A)
	cpu.PC += uint16(len(operands) + 1)
}
