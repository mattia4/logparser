package main

import (
	"bufio"
	"encoding/json"
	"html/template"
	logger "logparser/logger"
	"logparser/models"
	p "logparser/parser"
	pm "logparser/parser/models"
	_ "logparser/parser/types"
	logparser_template "logparser/template"
	logparser_utils "logparser/utils"
	"os"
	"path/filepath"
)

func main() {
	logger := logger.NewLogger(true)

	file, err := os.Open("logs.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var logs []pm.LogEntry
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parsers := pm.RegisteredParsers
		parsedEntry := p.ParseLogLine(line, parsers)
		logs = append(logs, parsedEntry)

	}

	htmlFile, err := os.ReadFile(logparser_template.GetHtmlTemplatePath())
	if err != nil {
		panic(err)
	}

	htmlTemplate, err := template.New("log").Parse(string(htmlFile))
	if err != nil {
		panic(err)
	}

	htmlBytes, err := json.Marshal(logs)
	if err != nil {
		panic(err)
	}

	rows := models.TemplateData{
		LogsJSON: template.JS(htmlBytes),
	}

	if err := logparser_utils.CreateDir(logparser_template.GetOutputDir()); err != nil {
		logger.Fatal("Error in creating directory %s: %v", logparser_template.GetOutputDir(), err)
	}

	logsHtmlOut, err := os.Create(filepath.Join(logparser_template.GetOutputDir(), "logs.html"))
	if err != nil {
		panic(err)
	}
	defer logsHtmlOut.Close()

	htmlTemplate.Execute(logsHtmlOut, rows)

	err = logparser_utils.CopyFile(logparser_template.GetCssTemplatePath(), filepath.Join(logparser_template.GetOutputDir(), "template.css"))
	if err != nil {
		logger.Fatal("Error in creating template.css: %v", err)
	}

	err = logparser_utils.CopyFile(logparser_template.GetJsTemplatePath(), filepath.Join(logparser_template.GetOutputDir(), "template.js"))
	if err != nil {
		logger.Fatal("Error in creating template.js: %v", err)
	}
}
