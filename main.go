package main

import (
	"bufio"
	"encoding/json"
	"html/template"
	"logparser/models"
	"logparser/utils"
	"os"
)

func main() {
	file, err := os.Open("logs.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var logs []models.LogEntry
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		entry := utils.ParseLogLine(line)

		if entry != nil {
			logs = append(logs, *entry)
		}
	}

	htmlFile, err := os.ReadFile("template.html")
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("log").Parse(string(htmlFile))
	if err != nil {
		panic(err)
	}

	jsonBytes, err := json.Marshal(logs)
	if err != nil {
		panic(err)
	}

	rows := models.TemplateData{
		LogsJSON: template.JS(jsonBytes),
	}

	out, err := os.Create("logs.html")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	tmpl.Execute(out, rows)
}
