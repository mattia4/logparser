package utils

import (
	"io"
	"os"
)

func CopyFile(src, dest string) error {
	if _, err := os.Stat(dest); err == nil {
		return nil
	}

	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
}

func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
