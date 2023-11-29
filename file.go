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

func FileReadOne(targetpath string) (string, error) {
	file, err := os.Open(targetpath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	return string(content), err
}
