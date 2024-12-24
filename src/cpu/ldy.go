package cpu

func (cpu *CPU) LDYImmediate(operands []byte) {
	cpu.Y = operands[0]
	cpu.ZFlag(cpu.Y)
	cpu.NFlag(cpu.Y)
	cpu.PC += uint16(len(operands) + 1)
}
