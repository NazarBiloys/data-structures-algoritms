package service

import (
	"fmt"
	"os"
)

func MakeFile(fileName string) *os.File {
	file, err := os.Create("/logs/" + fileName)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return nil
	}

	return file
}

func WriteIntoFile(line string, file *os.File) {
	_, err := fmt.Fprintln(file, line)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
}
