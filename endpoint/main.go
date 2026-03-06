package main

import (
	"GoReloaded/Eugene"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("The input format should be: go run main.go <inputfile-name> <outputfile-name>")
		return
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	fmt.Println(function.FileToFileHandler(inputPath, outputPath))

}