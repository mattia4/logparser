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
	common_template "logparser/template/common"
	logparser_utils "logparser/utils"
	"os"
	"path/filepath"
)

func main() {
	logger := logger.NewLogger(true)

	file, err := os.Open("access.log.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var logs []pm.LogResult
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parsers := pm.RegisteredParsers
		parsedEntry := p.ParseLogLine(line, parsers)
		logs = append(logs, parsedEntry)

	}

	log_html_template_path,
		log_css_template_path,
		log_js_template_path := logparser_utils.HandleTemplatePath(logs[0].FormatTag)

	htmlFile, err := os.ReadFile(log_html_template_path)
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

	if err := logparser_utils.CreateDir(common_template.GetOutputDir()); err != nil {
		logger.Fatal("Error in creating directory %s: %v", common_template.GetOutputDir(), err)
	}

	logsHtmlOut, err := os.Create(filepath.Join(common_template.GetOutputDir(), "logs.html"))
	if err != nil {
		panic(err)
	}
	defer logsHtmlOut.Close()

	htmlTemplate.Execute(logsHtmlOut, rows)

	err = logparser_utils.CopyFile(log_css_template_path, filepath.Join(common_template.GetOutputDir(), "template.css"))
	if err != nil {
		logger.Fatal("Error in creating template.css: %v", err)
	}

	err = logparser_utils.CopyFile(log_js_template_path, filepath.Join(common_template.GetOutputDir(), "template.js"))
	if err != nil {
		logger.Fatal("Error in creating template.js: %v", err)
	}
}
