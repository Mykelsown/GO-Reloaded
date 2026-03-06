package function

import (
	"bufio"
	"fmt"
	"os"
)

func FileToFileHandler(inputFile, outputFile string) error {

	inputFilePath, InputErr := os.Open(inputFile)
	if InputErr != nil {
		return fmt.Errorf("there was an error opening the input file: %w", InputErr)
	}
	defer inputFilePath.Close()

	outputFilePath, OutputErr := os.Create(outputFile)
	if OutputErr != nil {
		return fmt.Errorf("there was an error creating the output file: %w", OutputErr)
	}
	defer outputFilePath.Close()

	bufferedInputReader := bufio.NewScanner(inputFilePath)
	bufferedOutputWriter := bufio.NewWriter(outputFilePath)

	for bufferedInputReader.Scan() {
		unrefinedText := bufferedInputReader.Text()
		refinedText := MyLogic(unrefinedText)
		if _, err := bufferedOutputWriter.WriteString(refinedText + "\n"); err != nil {
			return fmt.Errorf("error writing to output file: %w", err)
		}
	}

	if nonEofErr := bufferedInputReader.Err(); nonEofErr != nil {
		return fmt.Errorf("an error was encountered while reading text: %w", nonEofErr)
	}

	if err := bufferedOutputWriter.Flush(); err != nil {
		return fmt.Errorf("error flushing output buffer: %w", err)
	}

	return nil
}
