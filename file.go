package tool

import (
	"os"
	"path/filepath"
)

func FileGetDirName(targetPath string) string {
	if len(targetPath) < 1 {
		targetPath, _ = os.Getwd()
	}
	return filepath.Base(targetPath)
}
