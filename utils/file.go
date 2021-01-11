package utils

import (
	"io/ioutil"
	"os"
)

func GetFile(filepath string) *os.File {
	f, err := os.Open(filepath)
	if err != nil {
		return nil
	}
	return f
}

func WriteToFile(filepath string, data []byte) error {
	writeErr := ioutil.WriteFile(filepath, data, 0644)
	if writeErr != nil {
		return writeErr
	}
	return nil
}
