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

func (memory *ApplicationMemory) LoadProgram(program []byte, startAddr int16) {
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
	memory.Execute(start, int16(len(program)))
}

func (memory *ApplicationMemory) PrintScreen() {
	fmt.Println("Print Screen")
	for col := 0; col < len(memory.graphics_memory)/50; col++ {
		for row := col * 50; row < (col + 1) * 50; row++ {
			fmt.Print(string(memory.graphics_memory[row] & 255))
		}
		fmt.Print("\n")
	}
}

func (memory *ApplicationMemory) Execute(execAddr int16, length int16) {
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
			case 0x03: // STA
				instructionAddr++
				value := memory.memory_space[instructionAddr]
				memory.memory_space[value] = memory.register_A
			case 0x04: // END
				fmt.Println("here?")
				instructionAddr = load2BytesIntoInt16(memory.memory_space[instructionAddr+1:instructionAddr+2])
				fmt.Println("there?")
				if instructionAddr > execAddr + length ||
					instructionAddr < execAddr {
					panic("INVALID MEMORY ACCESS")
				}
				return
			default:
				panic("Unrecognized instruction: " + string((*memory).memory_space[instructionAddr]))
		}
		instructionAddr++
	}
}

func load2BytesIntoInt16(input []byte) int16 {
	out := int16(input[0])
	out <<= 8
	out += int16(input[1])
	return out
}

func Int16IntoTwoBytes(input int16, output []byte) {
	output[1] = byte(input & 255)
	input >>= 8
	output[0] = byte(input)
}
