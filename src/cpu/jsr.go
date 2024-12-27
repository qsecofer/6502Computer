package cpu

func (CPU *CPU) JSR(operands []byte) {
	low := int8(CPU.PC & 0xFF)
	high := int8(CPU.PC >> 8)
	CPU.bus.Write(uint16(CPU.SP), byte(low))
	CPU.SP--
	CPU.bus.Write(uint16(CPU.SP), byte(high))
	CPU.SP--
	address := uint16(operands[0]) | uint16(operands[1])<<8
	CPU.PC = address
}
