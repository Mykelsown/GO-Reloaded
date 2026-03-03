package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open(os.Args[1])
	defer input.Close()
	mainInKey := bufio.NewScanner(input)
	output, _ := os.Create(os.Args[2])
	defer output.Close()
	// mainOut := bufio.NewWriter(output)
	mainIn := mainInKey.Text()

	fmt.Printf("%v", mainIn)
}