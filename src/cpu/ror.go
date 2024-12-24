package cpu

func (cpu *CPU) RORAccumulator(operands []byte) {
	oldCarry := cpu.C
	cpu.CFlag(cpu.A)
	cpu.A = cpu.A >> 1
	if oldCarry {
		cpu.A |= 0x80
	}
	cpu.ZFlag(cpu.A)
	cpu.NFlag(cpu.A)
	cpu.PC += uint16(len(operands) + 1)
}
