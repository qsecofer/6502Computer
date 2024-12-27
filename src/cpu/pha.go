package cpu

func (cpu *CPU) PHA(operands []byte) {
	cpu.bus.Write(0x100+uint16(cpu.SP), cpu.A)
	cpu.SP--
	cpu.PC += uint16(len(operands) + 1)
}
