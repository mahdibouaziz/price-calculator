package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		message := fmt.Sprintf("failed to open file %v", filePath)
		return nil, errors.New(message)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		message := fmt.Sprintf("error scanning the content of the file %v", filePath)
		return nil, errors.New(message)
	}

	return lines, nil
}
