package main

import (
	"fmt"
	"os"

	"tired/functions"
)

// main reads the input file, validates the farm, finds optimal paths, and simulates ant movements.
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . 'fileName.txt'")
		return
	}

	fileName := os.Args[1]

	if !functions.IsValidFile(fileName) {
		fmt.Println("only text file are allowed")
		return
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	farm, err := functions.ValidateFormat(string(data))
	if err != nil {
		fmt.Println("ERROR: invalid data format,", err)
		return
	}

	paths, assigned := functions.GetPathsAndDistribute(&farm)
	if paths == nil {
		fmt.Println("ERROR: this ant farm cannot be solved")
		return
	}

	functions.MooveAnts(paths, farm.Antnumber, string(data), assigned)
}
