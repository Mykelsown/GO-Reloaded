package function

import (
	"bufio"
	"fmt"
	"os"
)

func FileToFileHandler(inputFile, outputFile string) error {
	// This opens the file where the unrefined text, sentences are.
	inputFilePath, InputErr := os.Open(inputFile)
	// Handles the error encountered while opening file
	if InputErr != nil {
		return fmt.Errorf("there was an error opening the input file: %w", InputErr)
	}
	defer inputFilePath.Close() // close the open file immmediately after the function returns.

	// This create/truncate a file where all the modified/refined sentences will stay
	outputFilePath, OutputErr := os.Create(outputFile)
	// Handles the error encountered when creating the file
	if OutputErr != nil {
		return fmt.Errorf("there was an error creating the output file: %w", OutputErr)
	}
	defer outputFilePath.Close() // closes the file after function returns

	bufferedInputReader := bufio.NewScanner(inputFilePath) // This read the text inside of the file and stores it as a buffered character, where the max limit of character it can store is 64kb
	bufferedOutputWriter := bufio.NewWriter(outputFilePath) // this create a space for writing to the create file

	for bufferedInputReader.Scan() { // This reads data in the input file line by line, until EOF
		unrefinedText := bufferedInputReader.Text() // returns a string of text of each line in the input file
		refinedText := MyLogic(unrefinedText)
		// Error handling if there is a problem writing to the space created in the output file.
		if _, err := bufferedOutputWriter.WriteString(refinedText + "\n"); err != nil {
			return fmt.Errorf("error writing to output file: %w", err)
		}
	}

	// Handling the error that isn't EOF while sanning 
	if nonEofErr := bufferedInputReader.Err(); nonEofErr != nil {
		return fmt.Errorf("an error was encountered while reading text: %w", nonEofErr)
	}

	// This is a crucial part of after writing to the file, this ensures all pending data is being transfered totally.
	if err := bufferedOutputWriter.Flush(); err != nil {
		return fmt.Errorf("error flushing output buffer: %w", err)
	}

	return nil // returns nil if no error found
}
