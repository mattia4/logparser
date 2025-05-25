package main

import (
	"bufio"
	"html/template"
	logger "logparser/logger"
	p "logparser/parser"
	pm "logparser/parser/models"
	_ "logparser/parser/types"
	common_template "logparser/template/common"
	file_utils "logparser/utils/file"
	"os"
	"path/filepath"
)

func main() {
	logger := logger.NewLogger(true)

	log_html_template_path := common_template.GetHtmlTemplatePath()
	log_css_template_path := common_template.GetCssTemplatePath()
	log_js_template_path := common_template.GetJsTemplatePath()

	file, err := os.Open("logs.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var logs []pm.LogResult
	var logsResults = pm.Result{}

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

	htmlFile, err := os.ReadFile(log_html_template_path)
	if err != nil {
		panic(err)
	}

	htmlTemplate, err := template.New("log").Parse(string(htmlFile))
	if err != nil {
		panic(err)
	}

	if err := file_utils.CreateDir(common_template.GetOutputDir()); err != nil {
		logger.Fatal("Error in creating directory %s: %v", common_template.GetOutputDir(), err)
	}

	logsHtmlOut, err := os.Create(filepath.Join(common_template.GetOutputDir(), "logs.html"))
	if err != nil {
		panic(err)
	}
	defer logsHtmlOut.Close()

	htmlTemplate.Execute(logsHtmlOut, nil)

	file_utils.InjectLogsDataJSON(logsResults, common_template.GetOutputDir())

	err = file_utils.CopyFile(log_css_template_path, filepath.Join(common_template.GetOutputDir(), "template.css"))
	if err != nil {
		logger.Fatal("Error in creating template.css: %v", err)
	}

	err = file_utils.CopyFile(log_js_template_path, filepath.Join(common_template.GetOutputDir(), "template.js"))
	if err != nil {
		logger.Fatal("Error in creating template.js: %v", err)
	}
}
