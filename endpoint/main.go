package main

import (
	"GoReloaded"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("The input format should be [go run maing.go <inputfile-name> <outputfile-name>]")
		return
	}

	inputPath := os.Args[1]
	outputPath := os.Args[2]

	fmt.Println(GoReloaded.FileToFileHandler(inputPath, outputPath))

}