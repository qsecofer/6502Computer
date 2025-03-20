package cpu

func (cpu *CPU) RTS(operands []byte) {
	cpu.Stack2PC()
	cpu.PC += 3
}
