package vm

type ApplicationMemory struct {
	register_A byte
	register_B byte
	register_X [2]byte
	register_Y [2]byte
	startAddress int16
	execAddress int16
	memory_space *[0x10000]byte // 64kb
	graphics_memory []byte
}

func (appMem *ApplicationMemory) getD() [2]byte {
	return [2]byte{appMem.register_A, appMem.register_B}
}

func (appMem *ApplicationMemory) setD(value [2]byte) {
	appMem.register_A = value[0]
	appMem.register_B = value[1]
}
