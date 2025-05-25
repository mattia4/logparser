package file

import (
	"encoding/json"
	"fmt"
	"io"
	pm "logparser/parser/models"
	"os"
	"path/filepath"
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

func InjectLogsDataJS(entries pm.Result, outputDir string) error {
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}

	whatToWrite := "const logsData = %s;\n"
	jsContent := fmt.Sprintf(whatToWrite, string(data))
	return os.WriteFile("logs_output/logs_data.js", []byte(jsContent), 0644)
}

func InjectLogsDataJSON(entries pm.Result, outputDir string) error {
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}

	outputPath := filepath.Join(outputDir, "logs_data.json")

	return os.WriteFile(outputPath, data, 0644)
}

func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
