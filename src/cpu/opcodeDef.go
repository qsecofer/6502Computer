package cpu

type AddressMode int

const (
	Implied AddressMode = iota
	Accumulator
	Immediate
	ZeroPage
	ZeroPageX
	ZeroPageY
	Relative
	Absolute
	AbsoluteX
	AbsoluteY
	Indirect
	IndexedIndirect
	IndirectIndexed
)

type Instruction struct {
	Mode     AddressMode
	Mnemonic string
	Operands int
	Fn       func(cpu *CPU, operans []byte)
}

var Opcodes = map[byte]Instruction{
	CLD:             {Mode: Implied, Mnemonic: "CLD", Operands: 0, Fn: func(cpu *CPU, operands []byte) { cpu.CLD(operands) }},
	CLI:             {Mode: Implied, Mnemonic: "CLI", Operands: 0, Fn: func(cpu *CPU, operands []byte) { cpu.CLI(operands) }},
	JMP_Absolute:    {Mode: Absolute, Mnemonic: "JMP", Operands: 2, Fn: func(cpu *CPU, operands []byte) { cpu.JMPAbsolute(operands) }},
	LDA_Immediate:   {Mode: Immediate, Mnemonic: "LDA", Operands: 1, Fn: func(cpu *CPU, operands []byte) { cpu.LDAImmediate(operands) }},
	LDX_Immediate:   {Mode: Immediate, Mnemonic: "LDX", Operands: 1, Fn: func(cpu *CPU, operands []byte) { cpu.LDXImmediate(operands) }},
	LDY_Immediate:   {Mode: Immediate, Mnemonic: "LDY", Operands: 1, Fn: func(cpu *CPU, operands []byte) { cpu.LDYImmediate(operands) }},
	ROR_Accumulator: {Mode: Accumulator, Mnemonic: "ROR", Operands: 0, Fn: func(cpu *CPU, operands []byte) { cpu.RORAccumulator(operands) }},
	STA_Absolute:    {Mode: Absolute, Mnemonic: "STA", Operands: 2, Fn: func(cpu *CPU, operands []byte) { cpu.STAAbsolute(operands) }},
	STY_Absolute:    {Mode: Absolute, Mnemonic: "STY", Operands: 2, Fn: func(cpu *CPU, operands []byte) { cpu.STYAbsolute(operands) }},
	NOP:             {Mode: Implied, Mnemonic: "NOP", Operands: 0, Fn: func(cpu *CPU, operands []byte) { cpu.NOP(operands) }},
	TXS:             {Mode: Implied, Mnemonic: "TXS", Operands: 0, Fn: func(cpu *CPU, operands []byte) { cpu.TXS(operands) }},
}
