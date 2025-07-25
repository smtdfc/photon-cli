package helpers

import (
	"log"
	"os"
	"path/filepath"
)

func GetCWD() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func DirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || !os.IsNotExist(err)
}

func EnsureDirAndResolve(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	dir := filepath.Dir(absPath)

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return "", err
	}

	return absPath, nil
}
