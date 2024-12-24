package cpu

func (cpu *CPU) TXS(operands []byte) {
	cpu.SP = cpu.X
	cpu.PC += uint16(len(operands) + 1)
}
