package models

import "html/template"

type LogEntry struct {
	Date    string
	Time    string
	Level   string
	Tag     string
	PID     string
	Message string
}

type TemplateData struct {
	LogsJSON template.JS
}
