package cpu

func (cpu *CPU) INX(operands []byte) {
	cpu.X++
	cpu.Z = cpu.X == 0
	cpu.N = cpu.X&0x80 != 0
	cpu.PC += uint16(len(operands) + 1)
}
