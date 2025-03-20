package cpu

const (
	ADC_Immediate  = 0x69
	ADC_Zeropage   = 0x65
	ADC_Zeropage_x = 0x75
	ADC_Absolute   = 0x6D
	ADC_Absolute_x = 0x7D
	ADC_Absolute_y = 0x79
	ADC_Indirect_x = 0x61
	ADC_Indirect_y = 0x71

	AND_Immediate  = 0x29
	AND_Zeropage   = 0x25
	AND_Zeropage_x = 0x35
	AND_Absolute   = 0x2D
	AND_Absolute_x = 0x3D
	AND_Absolute_y = 0x39
	AND_Indirect_x = 0x21
	AND_Indirect_y = 0x31

	ASL_Accumulator = 0x0A

	BNE = 0xD0

	BEQ = 0xF0

	BMI_Relative = 0x30

	BIT_Absolute = 0x2C

	CLC = 0x18

	CLD = 0xD8

	CLI = 0x58

	INX = 0xE8

	JMP_Absolute = 0x4C

	JSR_Absolute = 0x20

	LDA_Immediate  = 0xA9
	LDA_Zeropage   = 0xA5
	LDA_Zeropage_x = 0xB5
	LDA_Absolute   = 0xAD
	LDA_Absolute_x = 0xBD
	LDA_Absolute_y = 0xB9
	LDA_Indirect_x = 0xA1
	LDA_Indirect_y = 0xB1

	LDX_Immediate  = 0xA2
	LDX_Zeropage   = 0xA6
	LDA_Zeropage_y = 0xB6
	LDX_Absolute   = 0xAE
	LDX_Absolute_y = 0xBE

	LDY_Immediate = 0xA0

	LSR_Accumulator = 0x4A

	NOP = 0xEA

	ORA_Immediate = 0x09

	PHA = 0x48

	PLA = 0x68

	ROR_Accumulator = 0x6A

	RTI = 0x40

	RTS = 0x60

	SED = 0xF8

	SEI = 0x78

	STA_Absolute = 0x8D
	STA_Zeropage = 0x85

	STY_Absolute = 0x8C

	TXS = 0x9A
)
