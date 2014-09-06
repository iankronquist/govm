package assembler

import (
	"io/ioutil"
	"strings"
	"os"
	"encoding/hex"
	"fmt"
	"../../vm"
)

func Assemble(input_file_name *string, ouput_file_name *string, origin *string) {
	input_file, err := ioutil.ReadFile(*input_file_name)
	if err != nil {
		panic(err)
	}
	//err = ioutil.WriteFile(*ouput_file_name, append([]byte("G32"), output...), 0644)
	out_file, err := os.Create(*ouput_file_name)
	if err != nil {
		panic(err)
	}
	out_file.Write([]byte("G32"))
	out_file.Write([]byte{byte(0x03), byte(0xe8)})
	out_file.Write([]byte{byte(0), byte(0)})
	Parse(input_file, out_file)
}

func Parse(input_file []byte, output_file *os.File) {
	ch := make(chan string)
	labelTable := make(map[string]byte)

	go Tokenize(string(input_file), ch)
	for token := range ch {
		// All labels end with a :
		if strings.HasSuffix(token, ":") {
			current_pos, err := output_file.Seek(0, os.SEEK_CUR)
			if err != nil {
				panic(err)
			}
			labelTable[token[:len(token)-1]] = byte(current_pos)
		}
		interpret(token, ch, output_file, &labelTable)
	}
}

func interpret(token string, ch chan string, output *os.File, labelTable *map[string]byte) {
	fmt.Print(token, " ")
	for instruction_iter := range AllInstructions {
		instruction := AllInstructions[instruction_iter]
		if token != instruction.mneumonic {
			continue
		} else {
			// TODO: store in a buffer and make this one big write
			output.Write([]byte{instruction.repr})
			for iter := 0; iter < instruction.expects_args; iter++ {
				op := <-ch
				fmt.Print(op, " ")
				op_as_bytes, err := parseOp(op, labelTable)
				if err != nil {
					panic("Error parsing ops near: " + instruction.mneumonic)
				}
				if len(op_as_bytes) > instruction.arg_sizes {
					panic("Operator " + op + " too long near: " + instruction.mneumonic)
				}
				// TODO: check for error
				output.Write(op_as_bytes)
			}
		}
	}
	fmt.Print("\n")
}


func parseOp(op string, labelTable *map[string]byte) ([]byte, error) {
	if len(op) > 1 && op[:2] == "0x" {
		op_byte, err := hex.DecodeString(op[2:])
		return op_byte, err
	}
	switch op {
		case "X":
			return []byte{byte(vm.X)}, nil
		case "Y":
			return []byte{byte(vm.Y)}, nil
		case "A":
			return []byte{byte(vm.A)}, nil
		case "B":
			return []byte{byte(vm.B)}, nil
		case "D":
			return []byte{byte(vm.D)}, nil
	}
	pos, ok := (*labelTable)[op]
	if ok {
		return []byte{byte(pos)}, nil
	}
	panic("Unrecognized symbol!" + op)
}
