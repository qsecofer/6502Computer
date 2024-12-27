package cpu

func (CPU *CPU) ANDImmediate(operands []byte) {
	CPU.A = CPU.A & operands[0]
	CPU.Z = CPU.A == 0
	CPU.N = CPU.A&0x80 != 0
	CPU.PC += uint16(len(operands) + 1)
}
