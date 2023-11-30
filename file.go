package tool

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func FileGetDirName(targetPath string) string {
	if len(targetPath) < 1 {
		targetPath, _ = os.Getwd()
	}
	return filepath.Base(targetPath)
}

func FileReadToBytes(targetPath string) ([]byte, error) {
	file, err := os.Open(targetPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	return content, err
}

func FileReadToString(targetPath string) (string, error) {
	content, err := FileReadToBytes(targetPath)
	return string(content), err
}
