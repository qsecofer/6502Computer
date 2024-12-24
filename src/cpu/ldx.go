package cpu

func (cpu *CPU) LDXImmediate(operands []byte) {
	cpu.X = operands[0]
	cpu.ZFlag(cpu.X)
	cpu.NFlag(cpu.X)
	cpu.PC += uint16(len(operands) + 1)
}
