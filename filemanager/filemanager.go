package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)

	if err != nil {
		return nil, errors.New("could not open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	// scanner reads line by line and returns boolean
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("reading the file content failed")
	}

	file.Close()
	return lines, nil
}

func WriteJson(data interface{}, filePath string) error {
	file, err := os.Create(filePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to convert data")
	}

	file.Close()
	return nil
}
