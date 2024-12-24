package cpu

func (cpu *CPU) STAAbsolute(operands []byte) {
	address := uint16(operands[0]) | uint16(operands[1])<<8
	cpu.bus.Write(address, cpu.A)
	cpu.PC += uint16(len(operands) + 1)
}
