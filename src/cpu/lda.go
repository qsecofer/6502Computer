package cpu

func (cpu *CPU) LDAImmediate(operands []byte) {
	cpu.A = operands[0]
	cpu.ZFlag(cpu.A)
	cpu.NFlag(cpu.A)
	cpu.PC += uint16(len(operands) + 1)
}

func (cpu *CPU) LDAAbsolute(operands []byte) {
	address := uint16(operands[0]) | uint16(operands[1])<<8
	cpu.A = cpu.bus.Read(address)
	cpu.ZFlag(cpu.A)
	cpu.NFlag(cpu.A)
	cpu.PC += uint16(len(operands) + 1)
}

func (cpu *CPU) LDAAbsoluteX(operands []byte) {
	address := uint16(operands[0]) | uint16(operands[1])<<8
	address += uint16(cpu.X)
	cpu.A = cpu.bus.Read(address)
	cpu.ZFlag(cpu.A)
	cpu.NFlag(cpu.A)
	cpu.PC += uint16(len(operands) + 1)
}
