package main

import "fmt"
import "flag"
import "./vm"
import "io/ioutil"

func main() {
	program_name := flag.String("program", "test.g32", "The program to run")
	flag.Parse()
	program_data, err := ioutil.ReadFile(*program_name)
	fmt.Println(program_data)
	if err != nil {
		fmt.Println("Can't read the file: " + *program_name)
		panic(err)
	}
	VM := vm.InitVM()
	VM.RunProgram(program_data)
	VM.PrintScreen()


}
