package cpu

func (CPU *CPU) PHA(operands []byte) {
	CPU.bus.Write(uint16(CPU.SP), CPU.A)
	CPU.SP--
	CPU.PC += uint16(len(operands) + 1)
}
