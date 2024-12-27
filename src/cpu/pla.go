package cpu

func (cpu *CPU) PLA(operands []byte) {
	cpu.SP++
	cpu.A = cpu.bus.Read(uint16(cpu.SP))
	cpu.Z = cpu.A == 0
	cpu.N = cpu.A&0x80 > 0
	cpu.PC += uint16(len(operands) + 1)
}
