package vm

import "fmt"

func InitVM() *ApplicationMemory {
	fmt.Println("InitVM")
	memory := new(ApplicationMemory)
	memory.memory_space = new([0x10000]byte)
	memory.graphics_memory = memory.memory_space[0xa000:0xafa0]
	fmt.Println("VM inited")
	return memory
}

func (memory *ApplicationMemory) LoadProgram(program []byte, startAddr uint16) {
	fmt.Println("Load Program")
	fmt.Println(program)
	if int(startAddr) + len(program) > len(memory.memory_space) {
		panic("Program can't fit in memory")
	}
	copy(memory.memory_space[startAddr:int(startAddr) + len(program)], program)
	fmt.Println(program)
	fmt.Println(memory.memory_space[startAddr:int(startAddr) + len(program)])
}

func (memory *ApplicationMemory) RunProgram(program []byte) {
	fmt.Println("Run Program")
	fmt.Println(program[5:])
	if program[0] != byte('G') || program[1] != '3' || program[2] != '2' {
		panic("This isn't a GASM program")
	}
	start := load2BytesIntoInt16(program[3:5])
	memory.LoadProgram(program[5:], start)
	memory.Execute(start, uint16(len(program)))
}

func (memory *ApplicationMemory) PrintScreen() {
	fmt.Println("Print Screen")
	fmt.Println(memory.graphics_memory)
	for col := 0; col < len(memory.graphics_memory)/50; col++ {
		for row := col * 50; row < (col + 1) * 50; row++ {
			fmt.Print(string(memory.graphics_memory[row] & 255))
		}
		fmt.Print("\n")
	}
}

func (memory *ApplicationMemory) Execute(execAddr uint16, length uint16) {
	fmt.Println("Execute")
	instructionAddr := execAddr
	fmt.Println(execAddr)
	fmt.Println((*memory).memory_space[execAddr:execAddr+length])
	for {
		fmt.Println((*memory).memory_space[instructionAddr])
		switch (*memory).memory_space[instructionAddr] {
			case 0x02: // LDX
				instructionAddr++
				memory.register_X[0] = memory.memory_space[instructionAddr]
				instructionAddr++
				memory.register_X[1] = memory.memory_space[instructionAddr]
			case 0x01: // LDA
				instructionAddr++
				memory.register_A = memory.memory_space[instructionAddr]
				fmt.Println("LDA: ", memory.register_A)
			case 0x03: // STA
				instructionAddr++
				register := memory.memory_space[instructionAddr]
				value := [2]byte{0, 0}
				switch register {
					case A:
						value = [2]byte{0, memory.register_A}
					case B:
						value = [2]byte{0, memory.register_B}
					case D:
						value = memory.getD()
					case X:
						value = memory.register_X
					case Y:
						value = memory.register_Y
				}
				index := load2BytesIntoInt16(value[:])
				fmt.Println("value ", value, index)
				fmt.Println("register ", memory.register_A)
				memory.memory_space[index] = memory.register_A
			case 0x04: // END
				instructionAddr = load2BytesIntoInt16(memory.memory_space[instructionAddr+1:instructionAddr+3])
				if instructionAddr > execAddr + length ||
					instructionAddr < execAddr {
					//panic("INVALID MEMORY ACCESS")
				}
				return
			default:
				panic("Unrecognized instruction: " + string((*memory).memory_space[instructionAddr]))
		}
		instructionAddr++
	}
}

func load2BytesIntoInt16(input []byte) uint16 {
	out := uint16(input[0])
	out <<= 8
	out += uint16(input[1])
	return out
}

func Int16IntoTwoBytes(input uint16, output []byte) {
	output[1] = byte(input & 255)
	input >>= 8
	output[0] = byte(input)
}
