package models

import (
	"regexp"
)

type LogEntry struct {
	RawLine     string
	FormatTag   string
	Date        *string
	Time        *string
	Pid         *string
	Tid         *string
	Level       *string
	Tag         *string
	Message     *string
	OtherFields map[string]string
}

type AccessLogEntry struct {
	RawLine    string
	Site       string
	IPAddress  string
	Date       string
	Time       string
	Message    string
	StatusCode string
}

type LogResult struct {
	RawLine    string
	FormatTag  string
	ParsedData interface{}
	Cols       []ColTemplate
}

type Result struct {
	LogResult []LogResult
	Cols      []ColTemplate
}

type ColTemplate struct {
	Name  string
	Value string
}

type LogParser struct {
	Name    string
	Regex   *regexp.Regexp
	ParseFn func(matches []string, rawLine string) LogResult
}

var RegisteredParsers []LogParser
