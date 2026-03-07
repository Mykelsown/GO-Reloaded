package main

import (
	"GoReloaded/functions"
	"fmt"
	"os"
)

// ----------------------
// @MYKELSOWN 😀🔀😀🐀🌀
// ----------------------

func main() {
	// Checks if the user input is a valid input, if it is not, it will print out the correct format of the input then exist the function
	if len(os.Args) < 3 {
		fmt.Println("The input format should be: go run main.go <inputfile-name> <outputfile-name>")
		return
	}

	inputPath := os.Args[1] // stores the input file name as a string
	outputPath := os.Args[2] // stores the output file name as a string

	// Prints out error found from:
	// Opening file
	// Reading file
	// Writing to file 
	fmt.Println(function.FileToFileHandler(inputPath, outputPath)) 

}