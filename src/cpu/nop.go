package cpu

func (cpu *CPU) NOP(operands []byte) {
	cpu.PC += uint16(len(operands) + 1)
}
