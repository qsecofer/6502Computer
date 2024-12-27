package cpu

func (cpu *CPU) JSR(operands []byte) {
	low := int8(cpu.PC & 0xFF)
	high := int8(cpu.PC >> 8)
	cpu.bus.Write(uint16(cpu.SP), byte(low))
	cpu.SP--
	cpu.bus.Write(uint16(cpu.SP), byte(high))
	cpu.SP--
	address := uint16(operands[0]) | uint16(operands[1])<<8
	cpu.PC = address
}
