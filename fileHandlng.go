package GoReloaded

import (
	"bufio"
	"fmt"
	"os"
)

func FileToFileHandler(inputFile, outputFile string) error {

	inputFilePath, InputErr := os.Open(inputFile)
	if InputErr != nil {
		fmt.Errorf("There was an error opening the input file: %w", InputErr)
	}
	defer inputFilePath.Close()
	outputFilePath, OutputErr := os.Create(outputFile)
	if OutputErr != nil {
		fmt.Errorf("There was an error creating the output file: %w", OutputErr)
	}
	defer outputFilePath.Close()

	bufferedInput_reader := bufio.NewScanner(inputFilePath)
	bufferedOutput_writer := bufio.NewWriter(outputFilePath)

	for bufferedInput_reader.Scan() {
		if nonEofErr := bufferedInput_reader.Err(); nonEofErr != nil {
			fmt.Errorf("An error was encounter while reading text: %w", nonEofErr)
		}

		line := bufferedInput_reader.Text()
		transformedText := MyLogic(line)
		bufferedOutput_writer.WriteString(transformedText + "\n")
	}
	
	defer bufferedOutput_writer.Flush()

	return nil
}
