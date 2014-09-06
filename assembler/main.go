package main

import (
	"flag"
	"./assembler"
)

func main() {
	input_file_name := flag.String("input", "test.gasm",  "The input gasm file")
	output_file_name := flag.String("output", "output.g32", "The ouput")
	//origin := flag.String("origin", "origin", "origin")
	origin := ""
	assembler.Assemble(input_file_name, output_file_name, &origin)
}
