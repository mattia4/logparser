package file

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	pm "logparser/parser/models"
	"os"
)

func InjectLogsDataJS(entries pm.Result, outputDir string) error {
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}

	whatToWrite := "const logsData = %s;\n"
	jsContent := fmt.Sprintf(whatToWrite, string(data))
	return os.WriteFile("logs_output/logs_data.js", []byte(jsContent), 0644)
}

func EncodeJsonOrError(entries pm.Result, errorHandler func(err error)) []byte {
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		errorHandler(err)
	}
	return data
}

func Open(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

func Read(path string) []byte {
	htmlFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return htmlFile
}

func CreateFile(filepath string) *os.File {
	output, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	return output
}

func CreateDirOrError(path string, errorHandler func(err error)) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		errorHandler(err)
	}
}

func WriteFileOrError(name string, data []byte, perm os.FileMode, errorHandler func(err error)) {
	if err := os.WriteFile(name, data, perm); err != nil {
		errorHandler(err)
	}
}

func CopyFileOrError(src string, dest string, errorHandler func(er error)) {
	if _, err := os.Stat(dest); err == nil {
		return
	}

	sourceFile, err := os.Open(src)
	if err != nil {
		errorHandler(err)
		return
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		errorHandler(err)
		return
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, sourceFile); err != nil {
		errorHandler(fmt.Errorf("error while copy '%s' a '%s': %w", src, dest, err))
		return
	}
}

func TemplateNewParse(tmplName string, file []byte) *template.Template {
	template, err := template.New(tmplName).Parse(string(file))
	if err != nil {
		panic(err)
	}
	return template
}
