package main

import (
	"bufio"
	logger "logparser/logger"
	p "logparser/parser"
	pm "logparser/parser/models"
	_ "logparser/parser/types"
	ct "logparser/template/common"
	fu "logparser/utils/file"
	"path/filepath"
)

func main() {
	logger := logger.NewLogger(true)

	htmlTP := ct.GetHtmlTemplatePath()
	cssTP := ct.GetCssTemplatePath()
	jsTP := ct.GetJsTemplatePath()

	var logs []pm.LogResult
	var logsResults = pm.Result{}

	file := fu.Open("logs.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parsers := pm.RegisteredParsers
		parsedEntry := p.ParseLogLine(line, parsers)
		logs = append(logs, parsedEntry)
		cols := parsedEntry.Cols

		logsResults.LogResult = logs
		logsResults.Cols = cols
	}

	htmlFile := fu.Read(htmlTP)
	htmlTemplate := fu.TemplateNewParse("log", htmlFile)

	fu.CreateDirOrError(ct.GetOutputDir(), func(err error) {
		logger.Fatal("Error in creating directory %s: %v", ct.GetOutputDir(), err)
	})

	logsHtmlOut := fu.CreateFile(filepath.Join(ct.GetOutputDir(), "logs.html"))
	defer logsHtmlOut.Close()

	htmlTemplate.Execute(logsHtmlOut, nil)

	data := fu.EncodeJsonOrError(logsResults, func(err error) {
		logger.Error("Parsing error: %v", err)
	})

	fu.WriteFileOrError(filepath.Join(ct.GetOutputDir(), "logs_data.json"), data, 0644, func(err error) {
		logger.Error("Error in writing file: %v", err)
	})

	fu.CopyFileOrError(cssTP, filepath.Join(ct.GetOutputDir(), "template.css"), func(err error) {
		logger.Fatal("Error in creating template.css: %v", err)
	})

	fu.CopyFileOrError(jsTP, filepath.Join(ct.GetOutputDir(), "template.js"), func(err error) {
		logger.Fatal("Error in creating template.js: %v", err)
	})
}
