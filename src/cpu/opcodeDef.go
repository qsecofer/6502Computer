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
	AND_Immediate:   {Mode: Immediate, Mnemonic: "AND", Operands: 1, Fn: func(cpu *CPU, operands []byte) { cpu.ANDImmediate(operands) }},
	BEQ:             {Mode: Relative, Mnemonic: "BEQ", Operands: 1, Fn: func(cpu *CPU, operands []byte) { cpu.BEQ(operands) }},
	BNE:             {Mode: Relative, Mnemonic: "BNE", Operands: 1, Fn: func(cpu *CPU, operands []byte) { cpu.BNE(operands) }},
	CLD:             {Mode: Implied, Mnemonic: "CLD", Operands: 0, Fn: func(cpu *CPU, operands []byte) { cpu.CLD(operands) }},
	CLI:             {Mode: Implied, Mnemonic: "CLI", Operands: 0, Fn: func(cpu *CPU, operands []byte) { cpu.CLI(operands) }},
	INX:             {Mode: Implied, Mnemonic: "INX", Operands: 0, Fn: func(cpu *CPU, operands []byte) { cpu.INX(operands) }},
	JMP_Absolute:    {Mode: Absolute, Mnemonic: "JMP", Operands: 2, Fn: func(cpu *CPU, operands []byte) { cpu.JMPAbsolute(operands) }},
	JSR_Absolute:    {Mode: Absolute, Mnemonic: "JSR", Operands: 2, Fn: func(cpu *CPU, operands []byte) { cpu.JSR(operands) }},
	LDA_Immediate:   {Mode: Immediate, Mnemonic: "LDA", Operands: 1, Fn: func(cpu *CPU, operands []byte) { cpu.LDAImmediate(operands) }},
	LDA_Absolute:    {Mode: Absolute, Mnemonic: "LDA", Operands: 2, Fn: func(cpu *CPU, operands []byte) { cpu.LDAAbsolute(operands) }},
	LDA_Absolute_x:  {Mode: AbsoluteX, Mnemonic: "LDA$,x", Operands: 2, Fn: func(cpu *CPU, operands []byte) { cpu.LDAAbsoluteX(operands) }},
	LDX_Immediate:   {Mode: Immediate, Mnemonic: "LDX", Operands: 1, Fn: func(cpu *CPU, operands []byte) { cpu.LDXImmediate(operands) }},
	LDY_Immediate:   {Mode: Immediate, Mnemonic: "LDY", Operands: 1, Fn: func(cpu *CPU, operands []byte) { cpu.LDYImmediate(operands) }},
	ORA_Immediate:   {Mode: Immediate, Mnemonic: "ORA", Operands: 1, Fn: func(cpu *CPU, operands []byte) { cpu.ORAImmediate(operands) }},
	PHA:             {Mode: Implied, Mnemonic: "PHA", Operands: 0, Fn: func(cpu *CPU, operands []byte) { cpu.PHA(operands) }},
	PLA:             {Mode: Implied, Mnemonic: "PLA", Operands: 0, Fn: func(cpu *CPU, operands []byte) { cpu.PLA(operands) }},
	ROR_Accumulator: {Mode: Accumulator, Mnemonic: "ROR", Operands: 0, Fn: func(cpu *CPU, operands []byte) { cpu.RORAccumulator(operands) }},
	RTI:             {Mode: Implied, Mnemonic: "RTI", Operands: 0, Fn: func(cpu *CPU, operands []byte) { cpu.RTI(operands) }},
	RTS:             {Mode: Implied, Mnemonic: "RTS", Operands: 0, Fn: func(cpu *CPU, operands []byte) { cpu.RTS(operands) }},
	STA_Absolute:    {Mode: Absolute, Mnemonic: "STA", Operands: 2, Fn: func(cpu *CPU, operands []byte) { cpu.STAAbsolute(operands) }},
	STY_Absolute:    {Mode: Absolute, Mnemonic: "STY", Operands: 2, Fn: func(cpu *CPU, operands []byte) { cpu.STYAbsolute(operands) }},
	NOP:             {Mode: Implied, Mnemonic: "NOP", Operands: 0, Fn: func(cpu *CPU, operands []byte) { cpu.NOP(operands) }},
	TXS:             {Mode: Implied, Mnemonic: "TXS", Operands: 0, Fn: func(cpu *CPU, operands []byte) { cpu.TXS(operands) }},
}
