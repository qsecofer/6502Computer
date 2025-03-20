package cpu

func (cpu *CPU) LSRAccumulator(operands []byte) {
	cpu.C = cpu.A&0x01 == 0x01
	cpu.A >>= 1
	cpu.Z = cpu.A == 0
	cpu.N = false
}
