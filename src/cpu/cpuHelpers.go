package cpu

func (cpu *CPU) ZFlag(value byte) {
	cpu.Z = value == 0
}

func (cpu *CPU) NFlag(value byte) {
	cpu.N = value&0x80 != 0
}

func (cpu *CPU) CFlag(value byte) {
	cpu.C = value&0x01 != 0
}
