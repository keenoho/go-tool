package tool

import (
	"io"
	"os"
)

func FileReadToBytes(targetPath string) ([]byte, error) {
	file, err := os.Open(targetPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	return content, err
}

func FileReadToString(targetPath string) (string, error) {
	content, err := FileReadToBytes(targetPath)
	return string(content), err
}
