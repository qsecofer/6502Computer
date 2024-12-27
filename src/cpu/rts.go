package cpu

func (cpu *CPU) RTS(operands []byte) {
	cpu.PC = 3 + uint16(cpu.bus.Read(uint16(cpu.SP)+2)) | uint16(cpu.bus.Read(uint16(cpu.SP)+1))<<8
	cpu.SP += 2
}
