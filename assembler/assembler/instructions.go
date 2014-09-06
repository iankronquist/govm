package assembler


type Instruction struct {
	mneumonic string
	repr byte
	expects_args int
	arg_sizes int // Should be 1, 2, or 0
}

var AllInstructions = []Instruction{
	Instruction{
		mneumonic: "LDA",
		repr: byte(0x01),
		expects_args: 1,
		arg_sizes: 1,
	},
	Instruction{
		mneumonic: "LDX",
		repr: byte(0x02),
		expects_args: 1,
		arg_sizes: 2,
	},
	Instruction{
		mneumonic: "STA",
		repr: byte(0x03),
		expects_args: 1,
		arg_sizes: 1,
	},
	Instruction{
		mneumonic: "END",
		repr: byte(0x04),
		expects_args: 1,
		arg_sizes: 1,
	},
}

