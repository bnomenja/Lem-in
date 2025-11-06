package main

import (
	"fmt"
	"os"

	"tired/functions"
)

// main reads the input file, validates the farm, finds optimal paths, and simulates ant movements.
func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run . 'fileName.txt'")
		return
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	farm, err := functions.ValidateFormat(string(data))
	if err != nil {
		fmt.Println("ERROR: invalid data format,", err)
		return
	}

	paths, assigned := functions.Suurballe(&farm)
	if paths == nil {
		fmt.Println("not solvable")
		return
	}

	functions.MooveAnts(paths, farm.Antnumber, string(data), assigned)

}
