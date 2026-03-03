package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("The input format should be [go run maing.go <inputfile-name> <outputfile-name>]")
		return
	}

	input, _ := os.Open(os.Args[1])
	defer input.Close()
	output, _ := os.Create(os.Args[2])
	defer output.Close()
	mainInKey := bufio.NewScanner(input)
	// mainOut := bufio.NewWriter(output)
	mainIn := mainInKey.Text()
	fmt.Printf("%v", mainIn)
}